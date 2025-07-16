package kmswallet

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type AccountSlot struct {
	Address    common.Address
	Secret     string
	Commitment string
	Nonce      uint64
	Mutex      sync.Mutex
}

type AccountPool struct {
	Accounts []*AccountSlot
	idx      int
	mu       sync.Mutex
}

// NewAccountPool initializes the pool with N accounts, each with its own secret/commitment
func NewAccountPool(addresses []common.Address, secrets []string, commitments []string) *AccountPool {
	accounts := make([]*AccountSlot, len(addresses))
	for i := range addresses {
		accounts[i] = &AccountSlot{
			Address:    addresses[i],
			Secret:     secrets[i],
			Commitment: commitments[i],
		}
	}
	return &AccountPool{Accounts: accounts}
}

// GetAvailableAccount returns the next available account (round-robin, locks the slot)
func (p *AccountPool) GetAvailableAccount() *AccountSlot {
	p.mu.Lock()
	slot := p.Accounts[p.idx]
	p.idx = (p.idx + 1) % len(p.Accounts)
	p.mu.Unlock()
	slot.Mutex.Lock()
	return slot
}

// ReleaseAccount unlocks the slot (for explicit use if needed)
func (p *AccountPool) ReleaseAccount(slot *AccountSlot) {
	slot.Mutex.Unlock()
}

// RotateCommitment generates and sets a new secret/commitment for the slot
func (slot *AccountSlot) RotateCommitment(newSecret, newCommitment *big.Int) {
	slot.Secret = newSecret.String()
	slot.Commitment = newCommitment.String()
}

// BuildTransactOpts returns a TransactOpts for this account (KMS or PK based)
func (slot *AccountSlot) BuildTransactOpts(signer bind.SignerFn) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:   slot.Address,
		Signer: signer,
	}
}

// BuildTransactOptsWithNonce returns a TransactOpts for this account with an explicit nonce (KMS or PK based)
func (slot *AccountSlot) BuildTransactOptsWithNonce(signer bind.SignerFn, nonce uint64) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:   slot.Address,
		Signer: signer,
	}
}
