package main

import (
	"context"
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
	"foundnone-vrf/commitment"
	"foundnone-vrf/config"
	"foundnone-vrf/handler"
	"foundnone-vrf/tx"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	ws, httpc, err := dialClients(cfg)
	if err != nil {
		return err
	}

	auth, contract, contractAddr, err := prepareAuthAndContract(cfg, httpc)
	if err != nil {
		return err
	}

	secret, comm, err := ensureCommitment(ctx, httpc, auth, contract)
	if err != nil {
		return err
	}

	payoutAddr := common.HexToAddress(cfg.PayoutAddress)
	return subscribeLoop(ctx, ws, httpc, auth, contract, secret, comm, contractAddr, payoutAddr, cfg.ConnectionRetries)
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

func prepareAuthAndContract(cfg config.Config, httpc *ethclient.Client) (*bind.TransactOpts, *abi.FoundnoneVRF, common.Address, error) {
	key, err := crypto.HexToECDSA(strings.TrimPrefix(cfg.FulfillerPK, "0x"))
	if err != nil {
		return nil, nil, common.Address{}, fmt.Errorf("parse private key: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(cfg.ChainID))
	if err != nil {
		return nil, nil, common.Address{}, fmt.Errorf("new transactor: %w", err)
	}

	contractAddr := common.HexToAddress(cfg.ContractAddress)
	contract, err := abi.NewFoundnoneVRF(contractAddr, httpc)
	if err != nil {
		return nil, nil, contractAddr, fmt.Errorf("contract binding: %w", err)
	}
	return auth, contract, contractAddr, nil
}

func ensureCommitment(ctx context.Context, httpc *ethclient.Client, auth *bind.TransactOpts, contract *abi.FoundnoneVRF) (*big.Int, *big.Int, error) {
	secret, comm, err := commitment.Load("zk/commitment.json")
	if err == nil {
		return secret, comm, nil
	}

	secret, comm, err = commitment.Generate()
	if err != nil {
		return nil, nil, fmt.Errorf("generate commitment: %w", err)
	}

	_, err = tx.SendWithRetry(ctx, httpc, auth, func(a *bind.TransactOpts) (*types.Transaction, error) {
		return contract.SetCommitment(a, comm)
	}, 5, 0.12, 30*time.Second)
	if err != nil {
		return nil, nil, fmt.Errorf("submit commitment transaction: %w", err)
	}

	if err := commitment.Save("zk/commitment.json", secret, comm); err != nil {
		return nil, nil, fmt.Errorf("save commitment: %w", err)
	}
	return secret, comm, nil
}

func subscribeLoop(
	ctx context.Context,
	ws *ethclient.Client,
	httpc *ethclient.Client,
	auth *bind.TransactOpts,
	contract *abi.FoundnoneVRF,
	secret, comm *big.Int,
	contractAddr, payoutAddr common.Address,
	retries int,
) error {
	query := ethereum.FilterQuery{Addresses: []common.Address{contractAddr}}
	logs := make(chan types.Log)
	sub, err := ws.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("subscribe logs: %w", err)
	}

	fmt.Printf("🔄 subscribing to events on %s\n", contractAddr.Hex())

	for {
		select {
		case subErr := <-sub.Err():
			log.Printf("subscription error: %v", subErr)
			sub.Unsubscribe()

			var newSub ethereum.Subscription
			for i := 1; i <= retries; i++ {
				log.Printf("reconnect attempt %d/%d…", i, retries)
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
				return fmt.Errorf("could not reconnect after %d attempts", retries)
			}

		case vLog := <-logs:
			event, err := contract.ParseRngRequested(vLog)
			if err != nil {
				continue
			}
			if err := handler.HandleEvent(ctx, httpc, contract, auth, event, secret, comm, payoutAddr, contractAddr); err != nil {
				log.Printf("HandleEvent error: %v", err)
			}
		}
	}
}
