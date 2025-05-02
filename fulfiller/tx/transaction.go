package tx

import (
	"context"
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

// SuggestFees sets EIP-1559 tip and fee cap on auth
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
	SuggestFeesAndGetNonce(ctx, client, auth)
	var lastErr error
	for attempt := range maxRetries {
		fmt.Printf("GasFeeCap: %s, GasTipCap: %s\n", auth.GasFeeCap.String(), auth.GasTipCap.String())
		rec, err := waitMinedWithTimeout(ctx, client, auth, txFunc, waitTimeout)
		if err == nil {
			return rec, nil
		}
		lastErr = err

		fmt.Printf("Transaction failed: %s\n", err.Error())

		if attempt < maxRetries && strings.Contains(err.Error(), "timeout") {
			BumpFee(auth, bumpFactor, client, ctx)
			continue
		}
		if attempt < maxRetries && strings.Contains(err.Error(), "underpriced") {
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
) (*types.Receipt, error) {
	tx, err := txFunc(auth)
	if err != nil {
		return nil, err
	}
	tmr := time.After(timeout)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-tmr:
			return nil, fmt.Errorf("timeout waiting for tx %s", tx.Hash().Hex())
		case <-ticker.C:
			if rec, err := client.TransactionReceipt(ctx, tx.Hash()); err == nil && rec != nil {
				if rec.Status == types.ReceiptStatusSuccessful {
					return rec, nil
				}
				return rec, fmt.Errorf("reverted (status %d)", rec.Status)
			}
		}
	}
}
