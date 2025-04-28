package prover

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
)

type CircuitInput struct {
	Secret     string `json:"secret"`
	Seed       string `json:"seed"`
	Entropy    string `json:"entropy"`
	Commitment string `json:"commitment"`
}

func Run(input CircuitInput) ([24]*big.Int, [3]*big.Int, error) {
	var zeroProof [24]*big.Int
	var zeroPub [3]*big.Int

	// 1) marshal input
	payload, err := json.Marshal(input)
	if err != nil {
		return zeroProof, zeroPub, fmt.Errorf("marshal input: %w", err)
	}

	// 2) call prover server
	resp, err := http.Post("http://localhost:3000/prove", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return zeroProof, zeroPub, fmt.Errorf("POST to prover: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return zeroProof, zeroPub, fmt.Errorf("prover server status %d", resp.StatusCode)
	}

	// 3) decode JSON into exported field
	var result struct {
		Calldata []string `json:"calldata"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return zeroProof, zeroPub, fmt.Errorf("decode response: %w", err)
	}

	if len(result.Calldata) != 27 {
		return zeroProof, zeroPub, fmt.Errorf("expected 27 values, got %d", len(result.Calldata))
	}

	// 4) parse proof (first 24) and public (last 3) with base=0
	var proof [24]*big.Int
	for i := range 24 {
		v := new(big.Int)
		if _, ok := v.SetString(result.Calldata[i], 0); !ok {
			return zeroProof, zeroPub, fmt.Errorf("invalid proof element %d: %s", i, result.Calldata[i])
		}
		proof[i] = v
	}

	var pub [3]*big.Int
	for i := range 3 {
		v := new(big.Int)
		if _, ok := v.SetString(result.Calldata[24+i], 0); !ok {
			return zeroProof, zeroPub, fmt.Errorf("invalid public input %d: %s", i, result.Calldata[24+i])
		}
		pub[i] = v
	}
	return proof, pub, nil
}
