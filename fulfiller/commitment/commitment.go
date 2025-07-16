package commitment

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
	"os"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

var (
	BN128FieldPrime, _ = new(big.Int).SetString(
		"21888242871839275222246405745257275088548364400416034343698204186575808495617", 0,
	)
)

type Commitment struct {
	Secret     *big.Int `json:"secret"`
	Commitment *big.Int `json:"commitment"`
}

func Generate() (*big.Int, *big.Int, error) {
	secret, err := rand.Int(rand.Reader, BN128FieldPrime)
	if err != nil {
		return nil, nil, err
	}
	s := big.NewInt(0).Set(secret)
	zero := big.NewInt(0)

	comm, err := poseidon.Hash([]*big.Int{s, zero})
	if err != nil {
		return nil, nil, err
	}
	log.Printf("generated commitment: %s", comm)
	return s, comm, nil
}

func Save(path string, s, c *big.Int) error {
	data := Commitment{Secret: s, Commitment: c}
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := os.MkdirAll("zk", 0o700); err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o600)
}

func Load(path string) (*big.Int, *big.Int, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}
	var data Commitment
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, nil, err
	}
	return data.Secret, data.Commitment, nil
}
