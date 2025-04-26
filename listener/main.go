package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/joho/godotenv"

	"entropy-listener/abi"
)

type CircuitInput struct {
	Secret     string `json:"secret"`
	Seed       string `json:"seed"`
	Entropy    string `json:"entropy"`
	Commitment string `json:"commitment"`
}

type Config struct {
	RPCURL            string `json:"rpc_url"`
	ContractAddress   string `json:"contract_address"`
	EntropyRolePK     string `json:"entropy_role_private_key"`
	PayoutAddress     string `json:"payout_address"`
	ChainID           string `json:"chain_id"`
	ConnectionRetries int    `json:"connection_retries"`
}

type CommitmentJson struct {
	Secret     string `json:"secret"`
	Commitment string `json:"commitment"`
}

var bn128FieldPrime, _ = new(big.Int).SetString(
	"21888242871839275222246405745257275088548364400416034343698204186575808495617", 10,
)

func initialize() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	config := Config{
		RPCURL:          os.Getenv("RPC_URL"),
		ContractAddress: os.Getenv("CONTRACT_ADDRESS"),
		EntropyRolePK:   os.Getenv("ENTROPY_ROLE_PRIVATE_KEY"),
		PayoutAddress:   os.Getenv("PAYOUT_ADDRESS"),
		ChainID:         os.Getenv("CHAIN_ID"),
	}

	if config.RPCURL == "" || config.ContractAddress == "" || config.EntropyRolePK == "" || config.PayoutAddress == "" || config.ChainID == "" {
		return config, fmt.Errorf("missing required environment variables")
	}

	return config, nil
}

func main() {

	ctx := context.Background()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	config, err := initialize()
	if err != nil {
		log.Fatalf("❌ Error initializing environment variables: %v", err)
	}

	contractAddr := common.HexToAddress(config.ContractAddress)

	payoutAddr := common.HexToAddress(config.PayoutAddress)

	client, err := ethclient.Dial(config.RPCURL)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	keyBytes, err := hex.DecodeString(strings.TrimPrefix(config.EntropyRolePK, "0x"))
	if err != nil {
		log.Fatalf("❌ Failed to decode private key: %v", err)
	}

	privateKey, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		log.Fatalf("❌ Invalid ECDSA private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(84532))
	if err != nil {
		log.Fatalf("❌ Failed to create transactor: %v", err)
	}

	contract, err := abi.NewFoundnoneVRF(contractAddr, client)
	if err != nil {
		log.Fatalf("Failed to bind contract: %v", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Log subscription failed: %v", err)
	}

	fmt.Println("Listening for VrfRequested events...")

	var secret, commitment *big.Int

	// attempt to load the existing secret and commitment from the json file
	commitmentFile, err := os.ReadFile("zk/commitment.json")
	if err == nil {
		var commitmentJson CommitmentJson
		err = json.Unmarshal(commitmentFile, &commitmentJson)
		if err != nil {
			log.Fatalf("❌ Failed to unmarshal commitment JSON: %v", err)
		}
		secret, _ = new(big.Int).SetString(commitmentJson.Secret, 10)
		commitment, _ = new(big.Int).SetString(commitmentJson.Commitment, 10)
	} else {
		log.Printf("❌ Failed to read commitment file: %v", err)
	}

	if secret == nil || commitment == nil {
		secret, commitment, err = generateAndSubmitCommitment(
			contract,
			client,
			auth,
		)
		if err != nil {
			log.Fatalf("❌ Failed to generate and submit commitment: %v", err)
		}
		log.Printf("✅ Generated new commitment: %s", commitment.String())
	}

	for {
		select {
		case subErr := <-sub.Err():
			log.Printf("❌ subscription error: %v", subErr)
			sub.Unsubscribe()

			var newSub ethereum.Subscription
			var dialErr, subErr2 error

			for i := 0; i < config.ConnectionRetries; i++ {
				log.Printf("🔄 reconnect attempt %d/%d…", i+1, config.ConnectionRetries)
				client, dialErr = ethclient.Dial(config.RPCURL)
				if dialErr != nil {
					log.Printf("❌ dial failed: %v", dialErr)
					time.Sleep(time.Second * 1)
					continue
				}

				newSub, subErr2 = client.SubscribeFilterLogs(context.Background(), query, logs)
				if subErr2 != nil {
					log.Printf("❌ subscribe failed: %v", subErr2)
					time.Sleep(time.Second * 1)
					continue
				}

				log.Println("✅ reconnected & resubscribed")
				sub = newSub
				break
			}

			if sub == nil {
				log.Fatalf("❌ could not reconnect after %d attempts (last dial error: %v, last subscribe error: %v)",
					config.ConnectionRetries, dialErr, subErr2)
			}
		case vLog := <-logs:
			event, err := contract.ParseVrfRequested(vLog)
			if err != nil {
				continue
			}

			fmt.Printf("🧩 New requestId: %v from %s\n", event.RequestId, event.Requester.Hex())

			raw := crypto.Keccak256(
				common.LeftPadBytes(event.RequestId.Bytes(), 32),
				common.LeftPadBytes(new(big.Int).SetUint64(vLog.BlockNumber).Bytes(), 32),
			)

			// turn into a big.Int and reduce modulo the bn128 field so that it fits in the poseidon hash
			seed := new(big.Int).SetBytes(raw)
			seed.Mod(seed, bn128FieldPrime)

			entropy, err := poseidon.Hash([]*big.Int{secret, seed})
			if err != nil {
				log.Fatalf("❌ Poseidon hashing failed: %v", err)
			}

			createFolderErr := os.MkdirAll("zk", os.ModePerm)
			if createFolderErr != nil {
				log.Fatalf("❌ Failed to create zk directory: %v", err)
			}

			input := CircuitInput{
				Secret:     secret.String(),
				Seed:       seed.String(),
				Entropy:    entropy.String(),
				Commitment: commitment.String(),
			}

			proofHex, pubHex := runSnarkJsPipeline(input)

			proofArr := [24]*big.Int{}
			for i, h := range proofHex {
				v := new(big.Int)
				if _, ok := v.SetString(h, 0); !ok {
					log.Fatalf("bad proof hex @%d: %s", i, h)
				}
				proofArr[i] = v
			}
			pubArr := [3]*big.Int{}
			for i, h := range pubHex {
				v := new(big.Int)
				if _, ok := v.SetString(h, 0); !ok {
					log.Fatalf("bad pubSignal hex @%d: %s", i, h)
				}
				pubArr[i] = v
			}

			// useful if you need to inspect the output of snarkjs export soliditycalldata
			// signals := []any{
			// 	proofHex, pubHex,
			// }
			// writeJson("zk/proofAndSignals.json", signals)

			// optionally, verify the proof on-chain before sending the transaction
			// callOpts := &bind.CallOpts{Context: ctx}
			// ok, err := contract.VerifyProof(
			// 	callOpts,
			// 	proofArr,
			// 	pubArr,
			// )
			// if err != nil {
			// 	log.Fatalf("❌ verifyProof call failed: %v", err)
			// }
			// if !ok {
			// 	log.Fatalf("❌ proof was invalid—no need to send tx")
			// }

			// we generate a new commitment here and pass to the contract
			// SubmitEntropy assigns our next commitment hash to this so that
			// we don't have to call SetCommitment again
			// the current commitment (used to generate this entropy)
			// is at index 2 of the public signal array
			secret, commitment = generateCommitment()

			commitmentJson := CommitmentJson{
				Secret:     secret.String(),
				Commitment: commitment.String(),
			}
			writeJson("zk/commitment.json", commitmentJson)

			println("🔑 New commitment: ", commitment.String())

			tx, err := contract.SubmitEntropy(
				auth,
				proofArr,
				pubArr,
				event.RequestId,
				payoutAddr,
				commitment,
			)

			if err != nil {
				log.Printf("❌ TX failed: %v", err)
				continue
			}

			receipt, err := bind.WaitMined(ctx, client, tx)
			if err != nil {
				log.Fatalf("❌ submitEntropy tx mining error: %v", err)
			}
			if receipt.Status != types.ReceiptStatusSuccessful {
				log.Fatalf("❌ submitEntropy reverted (status %d)", receipt.Status)
			}
			log.Printf("✅ submitEntropy succeeded: %s", tx.Hash().Hex())

		}
	}
}

func writeJson(path string, data any) {
	f, _ := os.Create(path)
	defer f.Close()
	json.NewEncoder(f).Encode(data)
}

// runSnarkJsPipeline does:
//  1. snarkjs plonk fullprove zk/input.json zk/vrf_js/vrf.wasm zk/vrf_final.zkey zk/proof.json zk/public.json
//  2. snarkjs zkey export soliditycalldata zk/public.json zk/proof.json --plonk
//
// and returns the two string slices (proof[24], publicSignals[3]).
func runSnarkJsPipeline(input CircuitInput) (proofHex []string, pubHex []string) {
	dir, err := os.MkdirTemp("", "zktemp-*")
	if err != nil {
		log.Fatalf("❌ make temp dir: %v", err)
	}
	// clean up the temp dir when done
	defer os.RemoveAll(dir)

	inFile := filepath.Join(dir, "input.json")
	proofF := filepath.Join(dir, "proof.json")
	publicF := filepath.Join(dir, "public.json")

	f, err := os.Create(inFile)
	if err != nil {
		log.Fatalf("❌ create input.json: %v", err)
	}
	if err := json.NewEncoder(f).Encode(input); err != nil {
		log.Fatalf("❌ encode input.json: %v", err)
	}
	f.Close()

	cmd := exec.Command(
		"snarkjs", "plonk", "fullprove",
		inFile,
		"zk/vrf_js/vrf.wasm",
		"zk/vrf_final.zkey",
		proofF,
		publicF,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ snarkjs fullprove failed: %v", err)
	}

	buf := &bytes.Buffer{}
	cmd = exec.Command(
		"snarkjs", "zkey", "export", "soliditycalldata",
		publicF, proofF, "--plonk",
	)
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ export soliditycalldata: %v", err)
	}

	// split the two JSON arrays at the boundary "]["
	raw := strings.TrimSpace(buf.String())
	raw = strings.TrimSuffix(raw, ";")
	const sep = "]["
	i := strings.Index(raw, sep)
	if i < 0 {
		log.Fatalf("❌ unexpected calldata format: %q", raw)
	}
	proofPart := raw[:i+1]
	publicPart := raw[i+1:]

	if err := json.Unmarshal([]byte(proofPart), &proofHex); err != nil {
		log.Fatalf("❌ parsing proof array: %v\n%s", err, proofPart)
	}
	if err := json.Unmarshal([]byte(publicPart), &pubHex); err != nil {
		log.Fatalf("❌ parsing publicSignals: %v\n%s", err, publicPart)
	}

	if len(proofHex) != 24 || len(pubHex) < 2 {
		log.Fatalf("❌ bad lengths: proof=%d, pub=%d", len(proofHex), len(pubHex))
	}
	return
}

func generateCommitment() (*big.Int, *big.Int) {
	secret, _ := rand.Int(rand.Reader, bn128FieldPrime)

	zero := big.NewInt(0)
	commitment, err := poseidon.Hash([]*big.Int{secret, zero})
	if err != nil {
		log.Fatalf("poseidon commitment failed: %v", err)
	}

	return secret, commitment
}

func generateAndSubmitCommitment(
	contract *abi.FoundnoneVRF,
	client *ethclient.Client,
	auth *bind.TransactOpts,
) (*big.Int, *big.Int, error) {

	secret, commitmentBig := generateCommitment()

	tx, err := contract.SetCommitment(auth, commitmentBig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to send SetCommitment txn: %w", err)
	}

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("tx mining error: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return nil, nil, fmt.Errorf("commitment tx reverted (status %d)", receipt.Status)
	}

	return secret, commitmentBig, nil
}
