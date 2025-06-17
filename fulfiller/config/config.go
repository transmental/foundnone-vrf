package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
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
}

func LoadConfig() (Config, error) {
	godotenv.Load(".env")
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
		ChainID:                      chainID,
		ConnectionRetries:            retries,
		RelayerURL:                   os.Getenv("RELAYER_URL"),
		RelayerConcurrencyLimit:      relayerConcurrencyLimit,
		WhitelistedCallbackAddresses: whitelistedCallbackAddresses,
		MaxCallbackGasLimit:          maxCallbackGasLimit,
	}
	if cfg.WSRPCURL == "" || cfg.HTTPRPCURL == "" || cfg.ContractAddress == "" ||
		cfg.FulfillerPK == "" || cfg.PayoutAddress == "" {
		return cfg, fmt.Errorf("missing required environment variables")
	}
	return cfg, nil
}
