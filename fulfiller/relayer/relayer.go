package relayer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RelayerResponse struct {
	Success bool   `json:"success"`
	TxHash  string `json:"txHash"`
}

var NetworkNameByChainId = map[int64]string{
	1:     "ETHEREUM",
	33111: "CURTIS",
	8453:  "BASE",
	42161: "ARBITRUM",
	84532: "BASE_SEPOLIA",
}

type RelayerDto struct {
	ContractAddress string `json:"contractAddress"`
	Args            []any  `json:"args"`
	FuncName        string `json:"funcName"`
	Value           string `json:"value"`
	Network         string `json:"network"`
	SkipSimulation  bool   `json:"skipSimulation"`
}

func Relay(ctx context.Context, relayerUrl string, contractAddress string, args []any, funcName string, value string, chainId int64) (*RelayerResponse, error) {
	var body = RelayerDto{
		ContractAddress: contractAddress,
		Args:            args,
		FuncName:        funcName,
		Value:           value,
		Network:         NetworkNameByChainId[chainId],
		SkipSimulation:  true,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal relayer request body: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", relayerUrl+"/relay", io.NopCloser(bytes.NewReader(bodyBytes)))
	if err != nil {
		return nil, fmt.Errorf("failed to create request to relayer: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed to send request to relayer: %w", err)
	}
	var relayerResponse RelayerResponse
	if err := json.NewDecoder(resp.Body).Decode(&relayerResponse); err != nil {
		return nil, fmt.Errorf("failed to decode relayer response: %w", err)
	}
	if !relayerResponse.Success {
		return nil, fmt.Errorf("relayer response indicates failure: %s", relayerResponse.TxHash)
	}

	log.Printf("Relayer response: %s", relayerResponse.TxHash)
	return &relayerResponse, nil
}
