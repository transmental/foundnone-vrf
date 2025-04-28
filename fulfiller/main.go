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
	ctx := context.Background()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	// RPC clients
	ws, err := ethclient.Dial(cfg.WSRPCURL)
	if err != nil {
		log.Fatal(err)
	}
	httpc, err := ethclient.Dial(cfg.HTTPRPCURL)
	if err != nil {
		log.Fatal(err)
	}

	// auth
	key, _ := crypto.HexToECDSA(strings.TrimPrefix(cfg.FulfillerPK, "0x"))
	auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(cfg.ChainID))

	// contract binding
	contractAddr := common.HexToAddress(cfg.ContractAddress)
	contract, err := abi.NewFoundnoneVRF(contractAddr, httpc)
	if err != nil {
		log.Fatal(err)
	}

	// load or generate commitment
	secret, comm, err := commitment.Load("zk/commitment.json")
	if err != nil {
		secret, comm, err = commitment.Generate()
		if err != nil {
			log.Fatal(err)
		}
		// submit initial commitment tx
		_, err = tx.WaitMined(ctx, httpc, auth, func(a *bind.TransactOpts) (*types.Transaction, error) {
			return contract.SetCommitment(a, comm)
		})
		if err != nil {
			log.Fatal(err)
		}
		if err := commitment.Save("zk/commitment.json", secret, comm); err != nil {
			log.Fatal(err)
		}
	}

	// subscribe to events
	query := ethereum.FilterQuery{Addresses: []common.Address{contractAddr}}
	logs := make(chan types.Log)
	sub, err := ws.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		log.Fatal(err)
	}

	payoutAddr := common.HexToAddress(cfg.PayoutAddress)
	fmt.Printf("🔄 subscribing to events on %s\n", cfg.ContractAddress)
	for {
		select {
		case subErr := <-sub.Err():
			log.Printf("❌ subscription error: %v", subErr)
			sub.Unsubscribe()

			var newSub ethereum.Subscription
			for i := 1; i <= cfg.ConnectionRetries; i++ {
				log.Printf("🔄 reconnect attempt %d/%d…", i, cfg.ConnectionRetries)
				newSub, err = ws.SubscribeFilterLogs(ctx, query, logs)
				if err != nil {
					log.Printf("❌ subscribe failed: %v", err)
					time.Sleep(time.Second)
					continue
				}
				log.Println("✅ reconnected & resubscribed")
				sub = newSub
				break
			}
			if sub == nil {
				log.Fatalf("❌ could not reconnect after %d attempts", cfg.ConnectionRetries)
			}

		case vLog := <-logs:
			event, err := contract.ParseRngRequested(vLog)
			if err != nil {
				continue
			}
			if err := handler.HandleEvent(ctx, httpc, contract, auth, event, secret, comm, payoutAddr, contractAddr); err != nil {
				log.Printf("handle error: %v", err)
			}
		}
	}
}
