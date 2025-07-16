package config

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// WalletMode: "single" or "kms"
// KMSKeyID: AWS KMS key id (if using KMS)
// KMSRegion: AWS region for KMS
// MaxAccounts: max number of KMS-derived accounts
// FundGatherEnabled: enable gather funds feature
type WalletMode string

const (
	WalletModeSingle WalletMode = "single"
	WalletModeKMS    WalletMode = "kms"
)

type Config struct {
	WSRPCURL                     string
	HTTPRPCURL                   string
	ContractAddress              string
	FulfillerPK                  string
	PayoutAddress                string
	ChainID                      int64
	ConnectionRetries            int
	RelayerURL                   string
	RelayerConcurrencyLimit      int
	WhitelistedCallbackAddresses []string
	MaxCallbackGasLimit          uint32
	WalletMode                   WalletMode
	KMSKeyID                     string
	KMSRegion                    string
	MaxAccounts                  int
	PoolMinGasWei                *big.Int // new: min gas threshold for pool accounts
	PoolRefillAmountWei          *big.Int // new: refill amount for pool accounts
	PGConnString                 string   // Postgres connection string for KMS wallet
}

func LoadConfig() (Config, error) {
	// Pool min gas threshold and refill amount (default values)
	poolMinGasWei := big.NewInt(0)
	poolMinGasWei.SetString("100000000000000", 10) // 0.005 ETH default
	if v := os.Getenv("POOL_MIN_GAS_WEI"); v != "" {
		if _, ok := poolMinGasWei.SetString(v, 10); !ok {
			return Config{}, fmt.Errorf("invalid POOL_MIN_GAS_WEI")
		}
	}
	poolRefillAmountWei := big.NewInt(0)
	poolRefillAmountWei.SetString("1000000000000000", 10) // 0.01 ETH default
	if v := os.Getenv("POOL_REFILL_AMOUNT_WEI"); v != "" {
		if _, ok := poolRefillAmountWei.SetString(v, 10); !ok {
			return Config{}, fmt.Errorf("invalid POOL_REFILL_AMOUNT_WEI")
		}
	}
	godotenv.Load(".env")
	walletMode := WalletMode(os.Getenv("WALLET_MODE"))
	if walletMode == "" {
		walletMode = WalletModeSingle
	}
	kmsKeyID := os.Getenv("KMS_KEY_ID")
	kmsRegion := os.Getenv("KMS_REGION")
	maxAccounts := 5
	if v := os.Getenv("MAX_ACCOUNTS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			maxAccounts = n
		}
	}
	chainID, err := strconv.ParseInt(os.Getenv("CHAIN_ID"), 10, 64)
	if err != nil {
		return Config{}, fmt.Errorf("invalid CHAIN_ID: %w", err)
	}
	retries := 5
	if v := os.Getenv("CONNECTION_RETRIES"); v != "" {
		if r, err := strconv.Atoi(v); err == nil {
			retries = r
		}
	}
	relayerConcurrencyLimit := 10
	if v := os.Getenv("RELAY_CONCURRENCY_LIMIT"); v != "" {
		if r, err := strconv.Atoi(v); err == nil {
			relayerConcurrencyLimit = r
		}
	}

	whitelistedCallbackAddresses := []string{}
	if v := os.Getenv("WHITELISTED_CALLBACK_ADDRESSES"); v != "" {
		// split the addresses by comma and lowercase, there should be no spaces to worry about here.
		// If there are spaces between the addresses, fix the env var.
		addresses := strings.Split(strings.ToLower(v), ",")
		whitelistedCallbackAddresses = append(whitelistedCallbackAddresses, addresses...)
	}

	maxCallbackGasLimit := uint32(100000)
	if v := os.Getenv("MAX_CALLBACK_GAS_LIMIT"); v != "" {
		if g, err := strconv.ParseUint(v, 10, 32); err == nil {
			maxCallbackGasLimit = uint32(g)
		} else {
			return Config{}, fmt.Errorf("invalid MAX_CALLBACK_GAS_LIMIT: %w", err)
		}
	}

	cfg := Config{
		WSRPCURL:                     os.Getenv("WS_RPC_URL"),
		HTTPRPCURL:                   os.Getenv("HTTP_RPC_URL"),
		ContractAddress:              os.Getenv("CONTRACT_ADDRESS"),
		FulfillerPK:                  os.Getenv("FULFILLER_PK"),
		PayoutAddress:                os.Getenv("PAYOUT_ADDRESS"),
		PGConnString:                 os.Getenv("PG_CONN_STRING"),
		ChainID:                      chainID,
		ConnectionRetries:            retries,
		RelayerURL:                   os.Getenv("RELAYER_URL"),
		RelayerConcurrencyLimit:      relayerConcurrencyLimit,
		WhitelistedCallbackAddresses: whitelistedCallbackAddresses,
		MaxCallbackGasLimit:          maxCallbackGasLimit,
		WalletMode:                   walletMode,
		KMSKeyID:                     kmsKeyID,
		KMSRegion:                    kmsRegion,
		MaxAccounts:                  maxAccounts,
		PoolMinGasWei:                poolMinGasWei,
		PoolRefillAmountWei:          poolRefillAmountWei,
	}
	if cfg.WSRPCURL == "" || cfg.HTTPRPCURL == "" || cfg.ContractAddress == "" ||
		cfg.FulfillerPK == "" || cfg.PayoutAddress == "" {
		return cfg, fmt.Errorf("missing required environment variables")
	}
	return cfg, nil
}
