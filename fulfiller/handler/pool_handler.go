package handler

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"foundnone-vrf/abi"
	commitmentModule "foundnone-vrf/commitment"
	"foundnone-vrf/kmswallet"
	"foundnone-vrf/prover"
	"foundnone-vrf/tx"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iden3/go-iden3-crypto/poseidon"
)

// HandleEventWithPool handles events using a pool of accounts, each with its own secret/commitment
// pkAuth: TransactOpts for the PK account (used for refilling)
// httpClient: *ethclient.Client for sending ETH
// Added: minGas and refillAmount are now configurable
func HandleEventWithPool(
	ctx context.Context,
	client bind.ContractBackend,
	contract *abi.Abi,
	pool *kmswallet.AccountPool,
	signerFn func(*big.Int) bind.SignerFn,
	event *abi.AbiRngRequested,
	payout common.Address,
	contractAddr common.Address,
	chainId int64,
	pkAuth *bind.TransactOpts,
	httpClient *ethclient.Client,
	minGas *big.Int,
	refillAmount *big.Int,
) error {
	slot := pool.GetAvailableAccount()
	defer pool.ReleaseAccount(slot)

	opts := &bind.CallOpts{Pending: true, Context: ctx}
	fulfilled, err := contract.Entropies(opts, event.RequestId)
	if err != nil {
		return fmt.Errorf("check if event is fulfilled: %w", err)
	}
	if fulfilled.Cmp(big.NewInt(0)) != 0 {
		return fmt.Errorf("event already fulfilled: %s", event.RequestId)
	}

	// Use the same RNG/seed logic as the main handler
	raw := crypto.Keccak256(
		common.LeftPadBytes(event.RequestId.Bytes(), 32),
		common.LeftPadBytes(new(big.Int).SetUint64(event.RequestBlockSet.Uint64()).Bytes(), 32),
		common.LeftPadBytes(event.BlockHash[:], 32),
	)
	seed := new(big.Int).Mod(new(big.Int).SetBytes(raw), commitmentModule.BN128FieldPrime)

	secretBig := new(big.Int)
	if _, ok := secretBig.SetString(slot.Secret, 0); !ok {
		return fmt.Errorf("invalid secret format for %s: %s", slot.Address.Hex(), slot.Secret)
	}

	entropy, err := poseidon.Hash([]*big.Int{secretBig, seed})
	if err != nil {
		return err
	}

	proofArr, pubArr, err := prover.Run(prover.CircuitInput{
		Secret:     slot.Secret,
		Seed:       seed.String(),
		Entropy:    entropy.String(),
		Commitment: slot.Commitment,
	})
	if err != nil {
		return err
	}

	// Use and increment per-account nonce for KMS pool
	auth := slot.BuildTransactOpts(signerFn(big.NewInt(chainId)))

	// Log the values before proof generation
	log.Printf("[DEBUG] Generating proof for account %s with commitment: %s", slot.Address.Hex(), slot.Commitment)
	log.Printf("[DEBUG] Seed: %s, Entropy: %s", seed.String(), entropy.String())

	receipt, err := tx.SendWithRetry(
		ctx,
		httpClient, // type assertion for SendWithRetry
		auth,
		func(a *bind.TransactOpts) (*types.Transaction, error) {
			return contract.SubmitEntropy(a, proofArr, pubArr, event.RequestId, payout)
		},
		5,
		0.2,
		30*time.Second,
	)

	if receipt != nil && receipt.Status == types.ReceiptStatusSuccessful {
		fmt.Printf("Successfully fulfilled entropy for request ID %s using account %s\n", event.RequestId, slot.Address.Hex())
		slot.Nonce++ // increment nonce after successful tx
		// --- ENSURE ACCOUNT GAS THRESHOLD & AUTO-REFILL ---
		bal, err := client.(*ethclient.Client).BalanceAt(ctx, slot.Address, nil)
		if err != nil {
			fmt.Printf("[WARN] Could not check balance for %s: %v\n", slot.Address.Hex(), err)
		} else if bal.Cmp(minGas) < 0 {
			fmt.Printf("[WARN] Account %s below gas threshold: %s wei. Refilling...\n", slot.Address.Hex(), bal.String())
			// Pass nil for nonce to let SendETH fetch the latest pending nonce for the shared pkAuth account
			refillReceipt, err := tx.SendETH(ctx, httpClient, pkAuth, slot.Address, refillAmount, big.NewInt(chainId))
			if err != nil {
				fmt.Printf("[ERROR] Failed to auto-refill account %s: %v\n", slot.Address.Hex(), err)
			} else {
				fmt.Printf("[INFO] Auto-refilled account %s with %s wei: %s\n", slot.Address.Hex(), refillAmount.String(), refillReceipt.TxHash.Hex())
			}
		}
	}

	// Robust error handling: log error, do not crash
	if err != nil {
		fmt.Printf("[ERROR] VRF fulfillment failed for request %s: %v\n", event.RequestId, err)
	}
	return nil
}
