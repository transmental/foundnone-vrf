package main

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"foundnone-vrf/abi"
	"foundnone-vrf/cli"
	"foundnone-vrf/commitment"
	"foundnone-vrf/config"
	"foundnone-vrf/handler"
	"foundnone-vrf/kmswallet"
	"foundnone-vrf/relayer"
	"foundnone-vrf/tx"
)

func main() {
	// gather command cli flags
	gather := flag.Bool("gather", false, "Sweep KMS balances and exit")
	dbDSN := flag.String("db", "", "Postgres DSN")
	rpcURL := flag.String("rpc", "", "Ethereum HTTP RPC URL")
	chainID := flag.Int64("chain", 0, "Chain ID for transactions")
	threshold := flag.Float64("threshold", 1e-5, "Minimum ETH balance to sweep")
	toAddr := flag.String("to", "", "Destination address")
	kmsKey := flag.String("kms_key", "", "Must include KMS Key")
	kmsRegion := flag.String("kms_region", "", "Must include KMS Region")

	// collect command cli flags
	collect := flag.Bool("collect", false, "collect VRF request fulfillment earnings and exit")
	key := flag.String("key", "", "private key for collecting fees")
	address := flag.String("address", "", "contract address to collect fees from")
	chainId := flag.Int64("chain_id", 0, "Chain ID for collecting fees")
	rpcUrl := flag.String("rpc_url", "", "Ethereum HTTP RPC URL for given chain ID")

	flag.Parse()
	switch {
	case *gather:
		if err := cli.Gather(
			*dbDSN, *rpcURL, *chainID, *threshold, *toAddr, *kmsKey, *kmsRegion,
		); err != nil {
			log.Fatalf("gather error: %v", err)
		}
	case *collect:
		if err := cli.CollectFees(
			*key, *address, *chainId, *rpcUrl,
		); err != nil {
			log.Fatalf("collect error: %v", err)
		}
	default:
		if err := run(context.Background()); err != nil {
			log.Fatal(err)
		}
	}
}

func run(ctx context.Context) error {
	// 1. Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	// 2. Dial both WS and HTTP RPCs
	ws, httpc, err := dialClients(cfg)
	if err != nil {
		return err
	}

	// 3. Prepare the “main” EOA auth and contract binding
	auth, contract, err := prepareAuthAndContract(cfg, httpc)
	if err != nil {
		return err
	}
	log.Printf("initialized with main EOA addr: %s", auth.From)

	// 4. If KMS mode, initialize your pool and signer function
	var pool *kmswallet.AccountPool
	var signerFn func(*big.Int) bind.SignerFn
	if cfg.WalletMode == config.WalletModeKMS {
		pool, signerFn, err = initKmsPool(ctx, cfg, httpc, contract, auth)
		if err != nil {
			return err
		}
	}

	// 5. For non-KMS, do a one-off commitment setup and get secret/comm
	var secret, comm *big.Int
	if cfg.WalletMode != config.WalletModeKMS {
		secret, comm, err = ensureCommitment(ctx, httpc, auth, contract, cfg)
		if err != nil {
			return err
		}
	}

	// 6. Hand off to the event subscription loop
	return subscribeLoop(
		ctx, ws, httpc,
		auth, contract,
		secret, comm,
		cfg, pool, signerFn,
	)
}

// initKmsPool encapsulates steps 4–after for KMS mode:
func initKmsPool(
	ctx context.Context,
	cfg config.Config,
	httpc *ethclient.Client,
	contract *abi.Abi,
	auth *bind.TransactOpts,
) (*kmswallet.AccountPool, func(*big.Int) bind.SignerFn, error) {

	log.Println("🔄 initializing KMS-backed pool…")

	// 4a. Open Postgres
	db, err := sql.Open("postgres", cfg.PGConnString)
	if err != nil {
		return nil, nil, fmt.Errorf("postgres open: %w", err)
	}
	defer db.Close()

	vault, err := kmswallet.NewKeyVault(db, cfg.KMSKeyID, cfg.KMSRegion)
	if err != nil {
		return nil, nil, fmt.Errorf("key vault init: %w", err)
	}

	// 4b. Ensure we have enough keys in the vault
	keys, commitments, secrets, err := vault.LoadAllKeys()
	if err != nil {
		return nil, nil, fmt.Errorf("load vault keys: %w", err)
	}
	if len(keys) < cfg.MaxAccounts {
		if err := topUpVault(vault, cfg.MaxAccounts-len(keys)); err != nil {
			return nil, nil, err
		}
		keys, commitments, secrets, err = vault.LoadAllKeys()
		if err != nil {
			return nil, nil, fmt.Errorf("reload vault keys: %w", err)
		}
	}

	// 4c. Build the AccountPool and its per-account signerFn
	addresses, keyMap := buildKeyMap(keys, commitments, secrets)
	pool := kmswallet.NewAccountPool(addresses, secrets, commitments)
	signerFn := buildPoolSignerFn(keyMap, cfg.ChainID)

	// 4d. On startup, fund low-balance accounts and set their commitment
	if err := bootstrapPoolAccounts(ctx, httpc, contract, cfg, pool, keyMap, auth, vault); err != nil {
		return nil, nil, err
	}

	return pool, signerFn, nil
}

// topUpVault generates + stores N new keys
func topUpVault(
	vault *kmswallet.KeyVault,
	needed int,
) error {
	log.Printf("Vault short by %d keys, generating…", needed)
	gen := func(priv *ecdsa.PrivateKey) (string, string, error) {
		secret, commitment, err := commitment.Generate()
		return secret.String(), commitment.String(), err
	}
	return vault.GenerateAndStoreKeys(needed, gen)
}

// buildKeyMap returns the slice of addresses and a map[address]privKey
func buildKeyMap(
	keys []*ecdsa.PrivateKey,
	commitments, secrets []string,
) ([]common.Address, map[common.Address]*ecdsa.PrivateKey) {
	addrs := make([]common.Address, len(keys))
	m := make(map[common.Address]*ecdsa.PrivateKey, len(keys))
	for i, priv := range keys {
		addr := crypto.PubkeyToAddress(priv.PublicKey)
		addrs[i] = addr
		m[addr] = priv
	}
	return addrs, m
}

// buildPoolSignerFn returns a bind.SignerFn that just wraps NewKeyedTransactor…
func buildPoolSignerFn(
	keyMap map[common.Address]*ecdsa.PrivateKey,
	chainID int64,
) func(*big.Int) bind.SignerFn {
	return func(chain *big.Int) bind.SignerFn {
		return func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			priv, ok := keyMap[addr]
			if !ok {
				return nil, fmt.Errorf("unknown account %s", addr.Hex())
			}
			opts, err := bind.NewKeyedTransactorWithChainID(priv, chain)
			if err != nil {
				return nil, err
			}
			return opts.Signer(opts.From, tx)
		}
	}
}

// bootstrapPoolAccounts does exactly what your loop did: fund + setCommitment
func bootstrapPoolAccounts(
	ctx context.Context,
	httpc *ethclient.Client,
	contract *abi.Abi,
	cfg config.Config,
	pool *kmswallet.AccountPool,
	keyMap map[common.Address]*ecdsa.PrivateKey,
	auth *bind.TransactOpts,
	vault *kmswallet.KeyVault,
) error {
	for i, slot := range pool.Accounts {
		log.Printf("--- pool[%d] %s ---", i, slot.Address.Hex())

		// Fund if below threshold
		bal, err := httpc.BalanceAt(ctx, slot.Address, nil)
		if err != nil {
			log.Printf("[WARN] balance check failed: %v", err)
			continue
		}
		if bal.Cmp(cfg.PoolMinGasWei) < 0 {
			_, err := tx.SendETH(ctx, httpc, auth, slot.Address, cfg.PoolRefillAmountWei, big.NewInt(cfg.ChainID))
			if err != nil {
				log.Printf("[WARN] funding failed: %v", err)
				continue
			}
			// wait confirmation
			for j := 0; j < 60; j++ {
				time.Sleep(time.Second)
				cb, _ := httpc.BalanceAt(ctx, slot.Address, nil)
				if cb.Cmp(cfg.PoolMinGasWei) >= 0 {
					break
				}
			}
		}

		// Set commitment if mismatch
		onChain, err := contract.Commitments(&bind.CallOpts{Context: ctx}, slot.Address)
		if err != nil {
			log.Printf("[WARN] commitment check failed: %v", err)
			continue
		}
		want := new(big.Int)
		want.SetString(slot.Commitment, 0)
		if onChain.Cmp(want) != 0 {
			poolAuth, _ := bind.NewKeyedTransactorWithChainID(keyMap[slot.Address], big.NewInt(cfg.ChainID))
			newSecret, newComm, err := commitment.Generate()
			if err != nil {
				log.Printf("[WARN] commitment generation failed: %v", err)
				continue
			}
			slot.Secret = newSecret.String()
			slot.Commitment = newComm.String()
			_, err = tx.SendWithRetry(ctx, httpc, poolAuth,
				func(a *bind.TransactOpts) (*types.Transaction, error) {
					return contract.SetCommitment(a, newComm)
				},
				5, 0.2, 30*time.Second,
			)
			if err != nil {
				log.Printf("[WARN] setCommitment failed: %v", err)
			}

			// update the commitment and secret in the db
			if err := kmswallet.UpdateCommitmentAndSecretInDbForAddress(vault, slot.Address, slot.Commitment, slot.Secret); err != nil {
				log.Printf("[WARN] update commitment/secret in db failed: %v", err)
			} else {
				log.Printf("Updated commitment for %s: %s", slot.Address.Hex(), slot.Commitment)
			}
		}
	}
	return nil
}

func dialClients(cfg config.Config) (*ethclient.Client, *ethclient.Client, error) {
	ws, err := ethclient.Dial(cfg.WSRPCURL)
	if err != nil {
		return nil, nil, fmt.Errorf("dial wsrpc: %w", err)
	}
	httpc, err := ethclient.Dial(cfg.HTTPRPCURL)
	if err != nil {
		return nil, nil, fmt.Errorf("dial httprpc: %w", err)
	}
	return ws, httpc, nil
}

func prepareAuthAndContract(cfg config.Config, httpc *ethclient.Client) (*bind.TransactOpts, *abi.Abi, error) {
	key, err := crypto.HexToECDSA(strings.TrimPrefix(cfg.FulfillerPK, "0x"))
	if err != nil {
		return nil, nil, fmt.Errorf("parse private key: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(cfg.ChainID))
	if err != nil {
		return nil, nil, fmt.Errorf("new transactor: %w", err)
	}
	contractAddr := common.HexToAddress(cfg.ContractAddress)
	contract, err := abi.NewAbi(contractAddr, httpc)
	if err != nil {
		return nil, nil, fmt.Errorf("contract binding: %w", err)
	}
	return auth, contract, nil
}

func ensureCommitment(ctx context.Context, httpc *ethclient.Client, auth *bind.TransactOpts, contract *abi.Abi, cfg config.Config) (*big.Int, *big.Int, error) {
	secret, comm, err := commitment.Generate()
	if err != nil {
		return nil, nil, fmt.Errorf("generate commitment: %w", err)
	}
	switch {
	case cfg.RelayerURL != "":
		log.Println("🔄 relaying commitment to contract via relayer")
		relayerRes, err := relayer.Relay(
			ctx,
			cfg.RelayerURL,
			cfg.ContractAddress,
			[]any{comm.String()},
			"setCommitment",
			"0",
			cfg.ChainID,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("relay commitment: %w", err)
		}
		log.Printf("🔄 relayed commitment transaction: %s", relayerRes.TxHash)
		if err := commitment.Save("zk/commitment.json", secret, comm); err != nil {
			return nil, nil, fmt.Errorf("save commitment: %w", err)
		}
		return secret, comm, nil
	default:
		// log the address of the fulfiller eoa
		log.Printf("🔄 submitting commitment transaction directly from fulfiller EOA: %s to contract: %s on chainId: %v", auth.From.Hex(), cfg.ContractAddress, cfg.ChainID)
		_, err = tx.SendWithRetry(ctx, httpc, auth,
			func(a *bind.TransactOpts) (*types.Transaction, error) {
				return contract.SetCommitment(a, comm)
			},
			5, 0.2, 30*time.Second)

		if err != nil {
			return nil, nil, fmt.Errorf("submit commitment transaction: %w", err)
		}

		if err := commitment.Save("zk/commitment.json", secret, comm); err != nil {
			return nil, nil, fmt.Errorf("save commitment: %w", err)
		}
		return secret, comm, nil
	}
}

func subscribeLoop(
	ctx context.Context,
	ws *ethclient.Client,
	httpc *ethclient.Client,
	auth *bind.TransactOpts,
	contract *abi.Abi,
	secret, comm *big.Int,
	cfg config.Config,
	pool *kmswallet.AccountPool,
	signerFn func(*big.Int) bind.SignerFn,
) error {
	// Use config fields instead of removed variables
	query := ethereum.FilterQuery{Addresses: []common.Address{common.HexToAddress(cfg.ContractAddress)}, ToBlock: nil}
	logs := make(chan types.Log)
	sub, err := ws.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("subscribe logs: %w", err)
	}

	fmt.Printf("🔄 subscribing to events on %s\n", cfg.ContractAddress)

	relayLimiter := make(chan struct{}, cfg.RelayerConcurrencyLimit)

	for {
		select {
		case subErr := <-sub.Err():
			log.Printf("subscription error: %v", subErr)
			sub.Unsubscribe()

			var newSub ethereum.Subscription
			for i := 1; i <= cfg.ConnectionRetries; i++ {
				log.Printf("reconnect attempt %d/%d…", i, cfg.ConnectionRetries)
				newSub, err = ws.SubscribeFilterLogs(ctx, query, logs)
				if err != nil {
					log.Printf("subscribe failed: %v", err)
					time.Sleep(time.Second)
					continue
				}
				log.Println("reconnected & resubscribed")
				sub = newSub
				break
			}
			if sub == nil {
				return fmt.Errorf("could not reconnect after %d attempts", cfg.ConnectionRetries)
			}

		case vLog := <-logs:
			event, err := contract.ParseRngRequested(vLog)
			if err != nil {
				continue
			}
			if !checkCallbackAddressAndGasLimit(event.CallbackAddress, event.CallbackGasLimit, cfg.WhitelistedCallbackAddresses, cfg.MaxCallbackGasLimit) {
				continue
			}
			if cfg.WalletMode == config.WalletModeKMS {
				go func(event *abi.AbiRngRequested) {
					err := handler.HandleEventWithPool(
						ctx, httpc, contract, pool, signerFn, event,
						common.HexToAddress(cfg.PayoutAddress),
						common.HexToAddress(cfg.ContractAddress),
						cfg.ChainID, auth, httpc,
						cfg.PoolMinGasWei, cfg.PoolRefillAmountWei,
					)
					if err != nil {
						log.Printf("HandleEventWithPool error: %v", err)
					}
				}(event)
			} else {
				// Normal PK flow
				if err := handler.HandleEvent(ctx, httpc, contract, auth, event, secret, comm, common.HexToAddress(cfg.PayoutAddress), common.HexToAddress(cfg.ContractAddress), cfg.RelayerURL, relayLimiter, cfg.ChainID); err != nil {
					log.Printf("HandleEvent error: %v", err)
				}
			}
		}
	}
}

func checkCallbackAddressAndGasLimit(
	callbackAddress common.Address,
	gasRequired uint32,
	whitelistedAddresses []string,
	maxGas uint32,
) bool {
	if callbackAddress == (common.Address{}) {
		return true
	}

	if gasRequired > maxGas {
		log.Printf("Callback gas limit %d exceeds maximum allowed %d for address %s", gasRequired, maxGas, callbackAddress.Hex())
		return false
	}

	if len(whitelistedAddresses) == 0 {
		return true
	}

	cb := callbackAddress.Hex()

	for _, addr := range whitelistedAddresses {
		if strings.EqualFold(addr, cb) {
			return true
		}
	}

	log.Printf("Callback address %s is not whitelisted", cb)
	return false
}
