package tx

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SuggestFees sets EIP-1559 tip and fee cap on auth and sets the nonce
func SuggestFeesAndGetNonce(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts) error {
	tip, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		return err
	}
	nonce, err := client.NonceAt(ctx, auth.From, nil)
	if err != nil {
		return err
	}
	base, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}
	auth.GasTipCap = tip
	auth.GasFeeCap = new(big.Int).Add(base, tip)
	auth.Nonce = new(big.Int).SetUint64(nonce)
	return nil
}

// EstimateGas returns a gas estimation for the given payload
func EstimateGas(
	ctx context.Context,
	client *ethclient.Client,
	auth *bind.TransactOpts,
	to common.Address,
	data []byte,
) (uint64, error) {
	msg := ethereum.CallMsg{
		From:      auth.From,
		To:        &to,
		GasFeeCap: auth.GasFeeCap,
		GasTipCap: auth.GasTipCap,
		Data:      data,
	}
	return client.EstimateGas(ctx, msg)
}

// BumpFee multiplies the GasFeeCap and GasTipCap by factor
func BumpFee(auth *bind.TransactOpts, factor float64, client *ethclient.Client, ctx context.Context) {
	feeCap := auth.GasFeeCap
	tipCap := auth.GasTipCap

	mul := big.NewFloat(1 + factor)
	feeF := new(big.Float).Mul(mul, new(big.Float).SetInt(feeCap))
	tipF := new(big.Float).Mul(mul, new(big.Float).SetInt(tipCap))

	newFee, _ := feeF.Int(nil)
	newTip, _ := tipF.Int(nil)

	auth.GasFeeCap = newFee
	auth.GasTipCap = newTip
}

func SendWithRetry(
	ctx context.Context,
	client *ethclient.Client,
	auth *bind.TransactOpts,
	txFunc func(*bind.TransactOpts) (*types.Transaction, error),
	maxRetries int,
	bumpFactor float64,
	waitTimeout time.Duration,
) (*types.Receipt, error) {
	err := SuggestFeesAndGetNonce(ctx, client, auth)
	if err != nil {
		return nil, fmt.Errorf("failed to suggest fees: %w", err)
	}

	var lastErr error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		fmt.Printf("GasFeeCap: %s, GasTipCap: %s\n", auth.GasFeeCap.String(), auth.GasTipCap.String())

		rec, tx, err := waitMinedWithTimeout(ctx, client, auth, txFunc, waitTimeout)
		if err == nil && rec.Status == types.ReceiptStatusSuccessful {
			return rec, nil
		}

		if rec != nil && rec.Status == types.ReceiptStatusFailed {
			// Try to decode revert reason:
			reason, derr := getRevertReason(ctx, client, tx, auth.From)
			if derr != nil {
				lastErr = fmt.Errorf("tx reverted, reason decode failed: %w", derr)
			} else {
				lastErr = fmt.Errorf("tx reverted: %s", reason)
			}
			break
		}

		if err != nil {
			lastErr = err
		}

		fmt.Printf("Transaction failed: %s\n", lastErr.Error())

		if attempt < maxRetries && strings.Contains(lastErr.Error(), "timeout") {
			BumpFee(auth, bumpFactor, client, ctx)
			continue
		}
		if attempt < maxRetries && strings.Contains(lastErr.Error(), "underpriced") {
			BumpFee(auth, bumpFactor, client, ctx)
			continue
		}
		break
	}

	return nil, fmt.Errorf("tx failed after %d attempts: %w", maxRetries+1, lastErr)
}

func waitMinedWithTimeout(
	ctx context.Context,
	client *ethclient.Client,
	auth *bind.TransactOpts,
	txFunc func(*bind.TransactOpts) (*types.Transaction, error),
	timeout time.Duration,
) (*types.Receipt, *types.Transaction, error) {
	tx, err := txFunc(auth)
	if err != nil {
		return nil, nil, err
	}

	tmr := time.After(timeout)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-tmr:
			return nil, tx, fmt.Errorf("timeout waiting for tx %s", tx.Hash().Hex())
		case <-ticker.C:
			rec, err := client.TransactionReceipt(ctx, tx.Hash())
			if err == nil && rec != nil {
				return rec, tx, nil
			}
		}
	}
}

// getRevertReason simulates the failed tx to extract the revert reason.
func getRevertReason(ctx context.Context, client *ethclient.Client, tx *types.Transaction, from common.Address) (string, error) {
	msg := ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	// simulate at latest block
	data, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return "", fmt.Errorf("CallContract error: %w", err)
	}
	if len(data) < 4+32+32 {
		return "", fmt.Errorf("Revert data too short: %s", hex.EncodeToString(data))
	}

	// Solidity revert reason is a string: abi.encodeWithSignature("Error(string)")
	reasonLen := new(big.Int).SetBytes(data[4+32 : 4+32+32]).Int64()
	reasonBytes := data[4+32+32 : 4+32+32+reasonLen]
	return string(reasonBytes), nil
}
