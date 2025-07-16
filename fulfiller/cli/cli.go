package cli

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/lib/pq"

	"foundnone-vrf/abi"
	"foundnone-vrf/kmswallet"
	"foundnone-vrf/tx"
)

func Gather(
	dsn, rpcURL string,
	chainID int64,
	thresholdEth float64,
	toAddr, kmsKey, region string,
) error {
	log.Printf("[Gather] Starting with dsn=%s rpc=%s chain=%d threshold=%g to=%s kmsKey=%s region=%s",
		dsn, rpcURL, chainID, thresholdEth, toAddr, kmsKey, region,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("postgres open: %w", err)
	}
	defer db.Close()

	vault, err := kmswallet.NewKeyVault(db, kmsKey, region)
	if err != nil {
		return fmt.Errorf("new key vault: %w", err)
	}
	keys, _, _, err := vault.LoadAllKeys()
	if err != nil {
		return fmt.Errorf("load keys: %w", err)
	}
	log.Printf("[Gather] Loaded %d key(s) from vault", len(keys))
	if len(keys) == 0 {
		return fmt.Errorf("no keys found in vault; nothing to sweep")
	}

	// 2) Dial Ethereum node
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("dial rpc: %w", err)
	}

	thresholdWei := new(big.Int).Mul(
		big.NewInt(int64(thresholdEth*1e9)),
		big.NewInt(1e9),
	)

	to := common.HexToAddress(toAddr)

	for _, priv := range keys {
		addr := crypto.PubkeyToAddress(priv.PublicKey)
		bal, err := client.BalanceAt(context.Background(), addr, nil)
		if err != nil {
			log.Printf("[%s] balance error: %v", addr.Hex(), err)
			continue
		}

		if bal.Cmp(thresholdWei) < 0 {
			log.Printf("[%s] below threshold, skipping", addr.Hex())
			continue
		}

		// === reserve gas fee ===
		const gasLimit = 50000
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Printf("[%s] gas price error: %v", addr.Hex(), err)
			continue
		}
		fee := new(big.Int).Mul(big.NewInt(gasLimit), gasPrice)
		sendAmount := new(big.Int).Sub(bal, fee)
		if sendAmount.Sign() <= 0 {
			log.Printf("[%s] not enough to cover fee %s, skipping", addr.Hex(), fee)
			continue
		}
		log.Printf("[%s] balance %s, reserving fee %s, sending %s", addr.Hex(), bal, fee, sendAmount)

		auth, err := bind.NewKeyedTransactorWithChainID(priv, big.NewInt(chainID))
		if err != nil {
			log.Printf("[%s] transactor error: %v", addr.Hex(), err)
			continue
		}

		rec, err := tx.SendETH(context.Background(), client, auth, to, sendAmount, big.NewInt(chainID))
		if err != nil {
			log.Printf("[%s] send error: %v", addr.Hex(), err)
		} else {
			log.Printf("[%s] swept %s wei (fee %s) in tx %s", addr.Hex(), sendAmount, fee, rec.TxHash.Hex())
		}
	}
	return nil
}

func CollectFees(
	key, address string,
	chainId int64,
	rpcUrl string,
) error {
	log.Printf("[CollectFees] Starting with key=%s address=%s chainId=%d rpcUrl=%s",
		key, address, chainId, rpcUrl,
	)

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return fmt.Errorf("dial rpc: %w", err)
	}

	privKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return fmt.Errorf("parse private key: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(chainId))
	if err != nil {
		return fmt.Errorf("new transactor: %w", err)
	}

	contractAddr := common.HexToAddress(address)

	contract, err := abi.NewAbi(contractAddr, client)
	if err != nil {
		return fmt.Errorf("contract binding: %w", err)
	}

	tx, err := contract.WithdrawRewardReceiverBalance(auth)
	if err != nil {
		return fmt.Errorf("withdrawRewardReceiverBalance call failed: %w", err)
	}

	log.Printf("[CollectFees] Transaction sent: %s", tx.Hash().Hex())
	return nil
}
