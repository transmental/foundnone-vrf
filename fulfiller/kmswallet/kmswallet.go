package kmswallet

import (
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"database/sql"
	"encoding/asn1"
	"fmt"
	"math/big"
	"sync"

	_ "github.com/lib/pq"

	"github.com/btcsuite/btcd/btcec/v2"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type EncryptedKey struct {
	ID         int
	Address    string
	KMSKeyID   string
	Region     string
	Ciphertext []byte
	Commitment []byte
	Secret     []byte
}

type KeyVault struct {
	db        *sql.DB
	kmsClient *kms.KMS
	kmsKeyID  string
	region    string
}

func NewKeyVault(db *sql.DB, kmsKeyID, region string) (*KeyVault, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return nil, err
	}
	client := kms.New(sess)
	return &KeyVault{
		db:        db,
		kmsClient: client,
		kmsKeyID:  kmsKeyID,
		region:    region,
	}, nil
}

// Encrypt and store a new private key, commitment, and secret
func (kv *KeyVault) EncryptAndStoreKey(priv *ecdsa.PrivateKey, commitment, secret string) error {
	privBytes := crypto.FromECDSA(priv)
	encOut, err := kv.kmsClient.Encrypt(&kms.EncryptInput{
		KeyId:     &kv.kmsKeyID,
		Plaintext: privBytes,
	})
	if err != nil {
		return fmt.Errorf("KMS encrypt: %w", err)
	}
	address := crypto.PubkeyToAddress(priv.PublicKey).Hex()
	// log all data going into the DB
	fmt.Printf("[KMS DEBUG] Storing key for address %s with commitment %s and secret %s and ciphertext %s\n", address, commitment, secret, encOut.CiphertextBlob)
	_, err = kv.db.Exec(
		`INSERT INTO kms_eth_keys (address, kms_key_id, region, ciphertext, commitment, secret) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (address) DO NOTHING`,
		address, kv.kmsKeyID, kv.region, encOut.CiphertextBlob, commitment, secret,
	)
	return err
}

func UpdateCommitmentAndSecretInDbForAddress(kv *KeyVault, address common.Address, commitment, secret string) error {
	_, err := kv.db.Exec(
		`UPDATE kms_eth_keys SET commitment = $1, secret = $2 WHERE address = $3`,
		commitment, secret, address.Hex(),
	)
	if err != nil {
		return fmt.Errorf("update commitment and secret in DB: %w", err)
	}
	fmt.Printf("[KMS DEBUG] Updated commitment and secret for address %s\n", address)
	return nil
}

// Load all keys from the vault
func (kv *KeyVault) LoadAllKeys() ([]*ecdsa.PrivateKey, []string, []string, error) {
	rows, err := kv.db.Query(`SELECT address, ciphertext, commitment, secret FROM kms_eth_keys WHERE kms_key_id = $1 AND region = $2`, kv.kmsKeyID, kv.region)
	if err != nil {
		return nil, nil, nil, err
	}
	defer rows.Close()
	var keys []*ecdsa.PrivateKey
	var commitments []string
	var secrets []string
	for rows.Next() {
		var address string
		var ciphertext, commitment, secret string
		if err := rows.Scan(&address, &ciphertext, &commitment, &secret); err != nil {
			return nil, nil, nil, err
		}
		priv, err := kv.DecryptKey([]byte(ciphertext))
		if err != nil {
			return nil, nil, nil, fmt.Errorf("decrypt key for %s: %w", address, err)
		}
		keys = append(keys, priv)
		commitments = append(commitments, commitment)
		secrets = append(secrets, secret)
	}
	return keys, commitments, secrets, nil
}

// Decrypt a key
func (kv *KeyVault) DecryptKey(ciphertext []byte) (*ecdsa.PrivateKey, error) {
	out, err := kv.kmsClient.Decrypt(&kms.DecryptInput{
		CiphertextBlob: ciphertext,
		KeyId:          &kv.kmsKeyID,
	})
	if err != nil {
		return nil, fmt.Errorf("KMS decrypt: %w", err)
	}
	priv, err := crypto.ToECDSA(out.Plaintext)
	if err != nil {
		return nil, fmt.Errorf("toECDSA: %w", err)
	}
	return priv, nil
}

// Generate and store N new keys
func (kv *KeyVault) GenerateAndStoreKeys(n int, commitmentGen func(*ecdsa.PrivateKey) (string, string, error)) error {
	for range n {
		priv, err := crypto.GenerateKey()
		if err != nil {
			return err
		}
		commitment, secret, err := commitmentGen(priv)
		if err != nil {
			return err
		}
		if err := kv.EncryptAndStoreKey(priv, commitment, secret); err != nil {
			return err
		}
	}
	return nil
}

type KMSWallet struct {
	KMSKeyID  string
	Address   common.Address
	mu        sync.Mutex
	kmsClient *kms.KMS
	pubKey    *ecdsa.PublicKey
}

func NewKMSWallet(kmsKeyID string, maxAccounts int, region string) (*KMSWallet, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return nil, err
	}
	client := kms.New(sess)
	w := &KMSWallet{
		KMSKeyID:  kmsKeyID,
		kmsClient: client,
	}
	if err := w.deriveAddressFromKMS(); err != nil {
		return nil, err
	}
	return w, nil
}

// deriveAddressFromKMS fetches the public key from KMS and computes the Ethereum address
func (w *KMSWallet) deriveAddressFromKMS() error {
	out, err := w.kmsClient.GetPublicKey(&kms.GetPublicKeyInput{
		KeyId: &w.KMSKeyID,
	})
	if err != nil {
		return fmt.Errorf("KMS GetPublicKey: %w", err)
	}
	pubKey, err := parseKMSPublicKey(out.PublicKey)
	if err != nil {
		return fmt.Errorf("parse KMS pubkey: %w", err)
	}
	w.pubKey = pubKey
	w.Address = crypto.PubkeyToAddress(*pubKey)
	return nil
}

// GetAddress returns the KMS-derived Ethereum address
func (w *KMSWallet) GetAddress() common.Address {
	return w.Address
}

// KMS-based SignerFn for go-ethereum
func (w *KMSWallet) SignerFn(chainID *big.Int) bind.SignerFn {
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != w.Address {
			return nil, fmt.Errorf("signer: address mismatch")
		}
		h := tx.Hash()
		sig, err := w.kmsSignDigest(h.Bytes())
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(types.NewEIP155Signer(chainID), sig)
	}
}

// GatherFunds sweeps all ETH from the KMS account to the given address
func (w *KMSWallet) GatherFunds(ctx context.Context, client bind.ContractBackend, to common.Address) error {
	// TODO: Query balance, build and sign tx using KMS, send tx
	fmt.Printf("Would sweep funds from %s to %s\n", w.Address.Hex(), to.Hex())
	return nil
}

// --- Helpers ---

// parseKMSPublicKey parses a DER-encoded KMS public key to ecdsa.PublicKey
func parseKMSPublicKey(derBytes []byte) (*ecdsa.PublicKey, error) {
	fmt.Printf("[KMS DEBUG] DER bytes: %x\n", derBytes)
	// Try Go x509 first
	pub, err := x509.ParsePKIXPublicKey(derBytes)
	if err == nil {
		ecPub, ok := pub.(*ecdsa.PublicKey)
		if ok && ecPub.Curve == crypto.S256() {
			return ecPub, nil
		}
	}
	// Fallback: extract EC point from DER and use btcec
	var spki struct {
		Algo struct {
			Algorithm  asn1.ObjectIdentifier
			Parameters asn1.RawValue
		}
		BitString asn1.BitString
	}
	if _, err := asn1.Unmarshal(derBytes, &spki); err != nil {
		return nil, fmt.Errorf("failed to unmarshal DER: %w", err)
	}
	pubkeyBytes := spki.BitString.Bytes
	btcecPub, err := btcec.ParsePubKey(pubkeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse secp256k1 pubkey with btcec: %w", err)
	}
	ecdsaPub := btcecPub.ToECDSA()
	return ecdsaPub, nil
}

// kmsSignDigest signs a digest using AWS KMS and returns a 65-byte Ethereum signature
func (w *KMSWallet) kmsSignDigest(digest []byte) ([]byte, error) {
	// KMS expects a 32-byte digest
	if len(digest) != 32 {
		return nil, fmt.Errorf("digest must be 32 bytes")
	}
	out, err := w.kmsClient.Sign(&kms.SignInput{
		KeyId:            &w.KMSKeyID,
		Message:          digest,
		MessageType:      aws.String(kms.MessageTypeDigest),
		SigningAlgorithm: aws.String(kms.SigningAlgorithmSpecEcdsaSha256),
	})
	if err != nil {
		return nil, fmt.Errorf("KMS sign: %w", err)
	}
	// KMS returns ASN.1 DER encoded signature
	var parsed struct{ R, S *big.Int }
	_, err = asn1.Unmarshal(out.Signature, &parsed)
	if err != nil {
		return nil, fmt.Errorf("asn1 unmarshal: %w", err)
	}
	r, s := parsed.R, parsed.S
	// Ethereum requires [R || S || V] where V is recovery id
	// V is not provided by KMS, so we must recover it
	sig := make([]byte, 65)
	copy(sig[0:32], r.Bytes())
	copy(sig[32:64], s.Bytes())
	// Recovery ID (V): try both 0 and 1
	pub := w.pubKey
	for v := byte(0); v < 2; v++ {
		sig[64] = v
		recovered, err := crypto.SigToPub(digest, sig)
		if err == nil && recovered.X.Cmp(pub.X) == 0 && recovered.Y.Cmp(pub.Y) == 0 {
			return sig, nil
		}
	}
	return nil, fmt.Errorf("could not recover public key from signature")
}
