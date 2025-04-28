package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	WSRPCURL          string
	HTTPRPCURL        string
	ContractAddress   string
	FulfillerPK       string
	PayoutAddress     string
	ChainID           int64
	ConnectionRetries int
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
	cfg := Config{
		WSRPCURL:          os.Getenv("WS_RPC_URL"),
		HTTPRPCURL:        os.Getenv("HTTP_RPC_URL"),
		ContractAddress:   os.Getenv("CONTRACT_ADDRESS"),
		FulfillerPK:       os.Getenv("FULFILLER_PK"),
		PayoutAddress:     os.Getenv("PAYOUT_ADDRESS"),
		ChainID:           chainID,
		ConnectionRetries: retries,
	}
	if cfg.WSRPCURL == "" || cfg.HTTPRPCURL == "" || cfg.ContractAddress == "" ||
		cfg.FulfillerPK == "" || cfg.PayoutAddress == "" {
		return cfg, fmt.Errorf("missing required environment variables")
	}
	return cfg, nil
}
