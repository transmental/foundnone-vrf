package handler

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iden3/go-iden3-crypto/poseidon"

	vrfAbi "foundnone-vrf/abi"
	commitmentModule "foundnone-vrf/commitment"
	"foundnone-vrf/prover"
	"foundnone-vrf/relayer"
	"foundnone-vrf/tx"
)

func HandleEvent(
	ctx context.Context,
	client *ethclient.Client,
	contract *vrfAbi.Abi,
	auth *bind.TransactOpts,
	event *vrfAbi.AbiRngRequested,
	secret, commitment *big.Int,
	payout common.Address,
	contractAddress common.Address,
	relayerUrl string,
	relayLimiter chan struct{},
	chainId int64,
) error {
	opts := &bind.CallOpts{
		Pending: true,
		Context: ctx,
	}
	fulfilled, err := contract.Entropies(opts, event.RequestId)
	if err != nil {
		return fmt.Errorf("check if event is fulfilled: %w", err)
	}
	if fulfilled.Cmp(big.NewInt(0)) != 0 {
		return fmt.Errorf("event already fulfilled: %s", event.RequestId)
	}

	raw := crypto.Keccak256(
		common.LeftPadBytes(event.RequestId.Bytes(), 32),
		common.LeftPadBytes(new(big.Int).SetUint64(event.RequestBlockSet.Uint64()).Bytes(), 32),
		common.LeftPadBytes(event.BlockHash[:], 32),
	)
	seed := new(big.Int).Mod(new(big.Int).SetBytes(raw), commitmentModule.BN128FieldPrime)

	entropy, err := poseidon.Hash([]*big.Int{secret, seed})
	if err != nil {
		return err
	}

	proofArr, pubArr, err := prover.Run(prover.CircuitInput{
		Secret:     secret.String(),
		Seed:       seed.String(),
		Entropy:    entropy.String(),
		Commitment: commitment.String(),
	})
	if err != nil {
		return err
	}

	switch {
	case relayerUrl != "":
		relayLimiter <- struct{}{}
		defer func() { <-relayLimiter }()

		proofHexArr := make([]string, len(proofArr))
		for i, v := range proofArr {
			proofHexArr[i] = hexutil.EncodeBig(v)
		}
		pubHexArr := make([]string, len(pubArr))
		for i, v := range pubArr {
			pubHexArr[i] = hexutil.EncodeBig(v)
		}

		params := []any{proofHexArr, pubHexArr, event.RequestId, payout}

		go func(reqID string, args []any) {
			ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			res, err := relayer.Relay(ctx2, relayerUrl, contractAddress.Hex(), args, "submitEntropy", "0", chainId)
			if err != nil || !res.Success {
				log.Printf("async relay failed for %s: %v", reqID, err)
			}
		}(event.RequestId.String(), params)

		fmt.Printf("Optimistically relayed entropy for request ID %s\n", event.RequestId)
		return nil

	default:
		receipt, err := tx.SendWithRetry(
			ctx,
			client,
			auth,
			func(a *bind.TransactOpts) (*types.Transaction, error) {
				return contract.SubmitEntropy(a, proofArr, pubArr, event.RequestId, payout)
			},
			5,
			0.12,
			30*time.Second,
		)

		if receipt != nil && receipt.Status == types.ReceiptStatusSuccessful {
			fmt.Printf("Successfully fulfilled entropy for request ID %s\n", event.RequestId)
		}

		return err
	}
}
