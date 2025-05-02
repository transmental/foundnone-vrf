package handler

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iden3/go-iden3-crypto/poseidon"

	vrfAbi "foundnone-vrf/abi"
	commitmentModule "foundnone-vrf/commitment"
	"foundnone-vrf/prover"
	"foundnone-vrf/tx"
)

func HandleEvent(
	ctx context.Context,
	client *ethclient.Client,
	contract *vrfAbi.FoundnoneVRF,
	auth *bind.TransactOpts,
	event *vrfAbi.FoundnoneVRFRngRequested,
	secret, commitment *big.Int,
	payout common.Address,
	contractAddress common.Address,
) error {
	// check if the event is already fulfilled
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
	fmt.Printf("Fulfilling requestID: %s\n", event.RequestId)
	// 1) compute seed
	raw := crypto.Keccak256(
		common.LeftPadBytes(event.RequestId.Bytes(), 32),
		common.LeftPadBytes(new(big.Int).SetUint64(event.Raw.BlockNumber-1).Bytes(), 32),
		common.LeftPadBytes(event.BlockHash[:], 32),
	)
	seed := new(big.Int).Mod(new(big.Int).SetBytes(raw), commitmentModule.BN128FieldPrime)

	// 2) compute entropy
	entropy, err := poseidon.Hash([]*big.Int{secret, seed})
	if err != nil {
		return err
	}

	// 3) run ZK prover
	proofArr, pubArr, err := prover.Run(prover.CircuitInput{
		Secret:     secret.String(),
		Seed:       seed.String(),
		Entropy:    entropy.String(),
		Commitment: commitment.String(),
	})
	if err != nil {
		return err
	}

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
