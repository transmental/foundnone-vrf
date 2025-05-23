// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FoundnoneVRFMetaData contains all meta data concerning the FoundnoneVRF contract.
var FoundnoneVRFMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adminRole\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitmentAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitmentInUse\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommitmentBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFeePercentage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequester\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSeedOrBlockHashUnavailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestAlreadyFulfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestNotFulfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestStillValid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPercentage\",\"type\":\"uint256\"}],\"name\":\"ContractFeePercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ContractFeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"RequestFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[24]\",\"name\":\"proof\",\"type\":\"uint256[24]\"},{\"indexed\":false,\"internalType\":\"uint256[3]\",\"name\":\"publicInputs\",\"type\":\"uint256[3]\"}],\"name\":\"RequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardReceiverBalanceWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePaid\",\"type\":\"uint256\"}],\"name\":\"RngRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"commitmentBlockSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"commitmentInUse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractFeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"entropies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getEntropy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardReceiver\",\"type\":\"address\"}],\"name\":\"getRewardReceiverBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"refundUnfulfilledRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestBlockSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestFeePaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestRng\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requesters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardReceiverBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_commitment\",\"type\":\"uint256\"}],\"name\":\"setCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPercentage\",\"type\":\"uint256\"}],\"name\":\"setContractFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"setRequestFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[24]\",\"name\":\"_proof\",\"type\":\"uint256[24]\"},{\"internalType\":\"uint256[3]\",\"name\":\"_publicInputs\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewardReceiver\",\"type\":\"address\"}],\"name\":\"submitEntropy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[24]\",\"name\":\"_proof\",\"type\":\"uint256[24]\"},{\"internalType\":\"uint256[3]\",\"name\":\"_pubSignals\",\"type\":\"uint256[3]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawContractFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawRewardReceiverBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FoundnoneVRFABI is the input ABI used to generate the binding from.
// Deprecated: Use FoundnoneVRFMetaData.ABI instead.
var FoundnoneVRFABI = FoundnoneVRFMetaData.ABI

// FoundnoneVRF is an auto generated Go binding around an Ethereum contract.
type FoundnoneVRF struct {
	FoundnoneVRFCaller     // Read-only binding to the contract
	FoundnoneVRFTransactor // Write-only binding to the contract
	FoundnoneVRFFilterer   // Log filterer for contract events
}

// FoundnoneVRFCaller is an auto generated read-only Go binding around an Ethereum contract.
type FoundnoneVRFCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoundnoneVRFTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FoundnoneVRFTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoundnoneVRFFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FoundnoneVRFFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoundnoneVRFSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FoundnoneVRFSession struct {
	Contract     *FoundnoneVRF     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FoundnoneVRFCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FoundnoneVRFCallerSession struct {
	Contract *FoundnoneVRFCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FoundnoneVRFTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FoundnoneVRFTransactorSession struct {
	Contract     *FoundnoneVRFTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FoundnoneVRFRaw is an auto generated low-level Go binding around an Ethereum contract.
type FoundnoneVRFRaw struct {
	Contract *FoundnoneVRF // Generic contract binding to access the raw methods on
}

// FoundnoneVRFCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FoundnoneVRFCallerRaw struct {
	Contract *FoundnoneVRFCaller // Generic read-only contract binding to access the raw methods on
}

// FoundnoneVRFTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FoundnoneVRFTransactorRaw struct {
	Contract *FoundnoneVRFTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFoundnoneVRF creates a new instance of FoundnoneVRF, bound to a specific deployed contract.
func NewFoundnoneVRF(address common.Address, backend bind.ContractBackend) (*FoundnoneVRF, error) {
	contract, err := bindFoundnoneVRF(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FoundnoneVRF{FoundnoneVRFCaller: FoundnoneVRFCaller{contract: contract}, FoundnoneVRFTransactor: FoundnoneVRFTransactor{contract: contract}, FoundnoneVRFFilterer: FoundnoneVRFFilterer{contract: contract}}, nil
}

// NewFoundnoneVRFCaller creates a new read-only instance of FoundnoneVRF, bound to a specific deployed contract.
func NewFoundnoneVRFCaller(address common.Address, caller bind.ContractCaller) (*FoundnoneVRFCaller, error) {
	contract, err := bindFoundnoneVRF(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FoundnoneVRFCaller{contract: contract}, nil
}

// NewFoundnoneVRFTransactor creates a new write-only instance of FoundnoneVRF, bound to a specific deployed contract.
func NewFoundnoneVRFTransactor(address common.Address, transactor bind.ContractTransactor) (*FoundnoneVRFTransactor, error) {
	contract, err := bindFoundnoneVRF(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FoundnoneVRFTransactor{contract: contract}, nil
}

// NewFoundnoneVRFFilterer creates a new log filterer instance of FoundnoneVRF, bound to a specific deployed contract.
func NewFoundnoneVRFFilterer(address common.Address, filterer bind.ContractFilterer) (*FoundnoneVRFFilterer, error) {
	contract, err := bindFoundnoneVRF(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FoundnoneVRFFilterer{contract: contract}, nil
}

// bindFoundnoneVRF binds a generic wrapper to an already deployed contract.
func bindFoundnoneVRF(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FoundnoneVRFMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FoundnoneVRF *FoundnoneVRFRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FoundnoneVRF.Contract.FoundnoneVRFCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FoundnoneVRF *FoundnoneVRFRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.FoundnoneVRFTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FoundnoneVRF *FoundnoneVRFRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.FoundnoneVRFTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FoundnoneVRF *FoundnoneVRFCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FoundnoneVRF.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FoundnoneVRF *FoundnoneVRFTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FoundnoneVRF *FoundnoneVRFTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_FoundnoneVRF *FoundnoneVRFCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_FoundnoneVRF *FoundnoneVRFSession) ADMINROLE() ([32]byte, error) {
	return _FoundnoneVRF.Contract.ADMINROLE(&_FoundnoneVRF.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) ADMINROLE() ([32]byte, error) {
	return _FoundnoneVRF.Contract.ADMINROLE(&_FoundnoneVRF.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FoundnoneVRF *FoundnoneVRFCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FoundnoneVRF *FoundnoneVRFSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FoundnoneVRF.Contract.DEFAULTADMINROLE(&_FoundnoneVRF.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FoundnoneVRF.Contract.DEFAULTADMINROLE(&_FoundnoneVRF.CallOpts)
}

// CommitmentBlockSet is a free data retrieval call binding the contract method 0x1ad6bdd6.
//
// Solidity: function commitmentBlockSet(address ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCaller) CommitmentBlockSet(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "commitmentBlockSet", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitmentBlockSet is a free data retrieval call binding the contract method 0x1ad6bdd6.
//
// Solidity: function commitmentBlockSet(address ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFSession) CommitmentBlockSet(arg0 common.Address) (*big.Int, error) {
	return _FoundnoneVRF.Contract.CommitmentBlockSet(&_FoundnoneVRF.CallOpts, arg0)
}

// CommitmentBlockSet is a free data retrieval call binding the contract method 0x1ad6bdd6.
//
// Solidity: function commitmentBlockSet(address ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) CommitmentBlockSet(arg0 common.Address) (*big.Int, error) {
	return _FoundnoneVRF.Contract.CommitmentBlockSet(&_FoundnoneVRF.CallOpts, arg0)
}

// CommitmentInUse is a free data retrieval call binding the contract method 0x21b4dae0.
//
// Solidity: function commitmentInUse(uint256 ) view returns(bool)
func (_FoundnoneVRF *FoundnoneVRFCaller) CommitmentInUse(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "commitmentInUse", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CommitmentInUse is a free data retrieval call binding the contract method 0x21b4dae0.
//
// Solidity: function commitmentInUse(uint256 ) view returns(bool)
func (_FoundnoneVRF *FoundnoneVRFSession) CommitmentInUse(arg0 *big.Int) (bool, error) {
	return _FoundnoneVRF.Contract.CommitmentInUse(&_FoundnoneVRF.CallOpts, arg0)
}

// CommitmentInUse is a free data retrieval call binding the contract method 0x21b4dae0.
//
// Solidity: function commitmentInUse(uint256 ) view returns(bool)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) CommitmentInUse(arg0 *big.Int) (bool, error) {
	return _FoundnoneVRF.Contract.CommitmentInUse(&_FoundnoneVRF.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0xe8fcf723.
//
// Solidity: function commitments(address ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCaller) Commitments(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0xe8fcf723.
//
// Solidity: function commitments(address ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFSession) Commitments(arg0 common.Address) (*big.Int, error) {
	return _FoundnoneVRF.Contract.Commitments(&_FoundnoneVRF.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0xe8fcf723.
//
// Solidity: function commitments(address ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) Commitments(arg0 common.Address) (*big.Int, error) {
	return _FoundnoneVRF.Contract.Commitments(&_FoundnoneVRF.CallOpts, arg0)
}

// ContractFeeBalance is a free data retrieval call binding the contract method 0xb160f9cc.
//
// Solidity: function contractFeeBalance() view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCaller) ContractFeeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "contractFeeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractFeeBalance is a free data retrieval call binding the contract method 0xb160f9cc.
//
// Solidity: function contractFeeBalance() view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFSession) ContractFeeBalance() (*big.Int, error) {
	return _FoundnoneVRF.Contract.ContractFeeBalance(&_FoundnoneVRF.CallOpts)
}

// ContractFeeBalance is a free data retrieval call binding the contract method 0xb160f9cc.
//
// Solidity: function contractFeeBalance() view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) ContractFeeBalance() (*big.Int, error) {
	return _FoundnoneVRF.Contract.ContractFeeBalance(&_FoundnoneVRF.CallOpts)
}

// ContractFeePercentage is a free data retrieval call binding the contract method 0xdbd8987c.
//
// Solidity: function contractFeePercentage() view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCaller) ContractFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "contractFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractFeePercentage is a free data retrieval call binding the contract method 0xdbd8987c.
//
// Solidity: function contractFeePercentage() view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFSession) ContractFeePercentage() (*big.Int, error) {
	return _FoundnoneVRF.Contract.ContractFeePercentage(&_FoundnoneVRF.CallOpts)
}

// ContractFeePercentage is a free data retrieval call binding the contract method 0xdbd8987c.
//
// Solidity: function contractFeePercentage() view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) ContractFeePercentage() (*big.Int, error) {
	return _FoundnoneVRF.Contract.ContractFeePercentage(&_FoundnoneVRF.CallOpts)
}

// Entropies is a free data retrieval call binding the contract method 0xe2ac9d5b.
//
// Solidity: function entropies(uint256 ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCaller) Entropies(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "entropies", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Entropies is a free data retrieval call binding the contract method 0xe2ac9d5b.
//
// Solidity: function entropies(uint256 ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFSession) Entropies(arg0 *big.Int) (*big.Int, error) {
	return _FoundnoneVRF.Contract.Entropies(&_FoundnoneVRF.CallOpts, arg0)
}

// Entropies is a free data retrieval call binding the contract method 0xe2ac9d5b.
//
// Solidity: function entropies(uint256 ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) Entropies(arg0 *big.Int) (*big.Int, error) {
	return _FoundnoneVRF.Contract.Entropies(&_FoundnoneVRF.CallOpts, arg0)
}

// GetEntropy is a free data retrieval call binding the contract method 0x74e6a469.
//
// Solidity: function getEntropy(uint256 _requestId) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCaller) GetEntropy(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "getEntropy", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEntropy is a free data retrieval call binding the contract method 0x74e6a469.
//
// Solidity: function getEntropy(uint256 _requestId) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFSession) GetEntropy(_requestId *big.Int) (*big.Int, error) {
	return _FoundnoneVRF.Contract.GetEntropy(&_FoundnoneVRF.CallOpts, _requestId)
}

// GetEntropy is a free data retrieval call binding the contract method 0x74e6a469.
//
// Solidity: function getEntropy(uint256 _requestId) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) GetEntropy(_requestId *big.Int) (*big.Int, error) {
	return _FoundnoneVRF.Contract.GetEntropy(&_FoundnoneVRF.CallOpts, _requestId)
}

// GetRewardReceiverBalance is a free data retrieval call binding the contract method 0x64b0c523.
//
// Solidity: function getRewardReceiverBalance(address _rewardReceiver) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCaller) GetRewardReceiverBalance(opts *bind.CallOpts, _rewardReceiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "getRewardReceiverBalance", _rewardReceiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardReceiverBalance is a free data retrieval call binding the contract method 0x64b0c523.
//
// Solidity: function getRewardReceiverBalance(address _rewardReceiver) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFSession) GetRewardReceiverBalance(_rewardReceiver common.Address) (*big.Int, error) {
	return _FoundnoneVRF.Contract.GetRewardReceiverBalance(&_FoundnoneVRF.CallOpts, _rewardReceiver)
}

// GetRewardReceiverBalance is a free data retrieval call binding the contract method 0x64b0c523.
//
// Solidity: function getRewardReceiverBalance(address _rewardReceiver) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) GetRewardReceiverBalance(_rewardReceiver common.Address) (*big.Int, error) {
	return _FoundnoneVRF.Contract.GetRewardReceiverBalance(&_FoundnoneVRF.CallOpts, _rewardReceiver)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FoundnoneVRF *FoundnoneVRFCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FoundnoneVRF *FoundnoneVRFSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FoundnoneVRF.Contract.GetRoleAdmin(&_FoundnoneVRF.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FoundnoneVRF.Contract.GetRoleAdmin(&_FoundnoneVRF.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FoundnoneVRF *FoundnoneVRFCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FoundnoneVRF *FoundnoneVRFSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FoundnoneVRF.Contract.HasRole(&_FoundnoneVRF.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FoundnoneVRF.Contract.HasRole(&_FoundnoneVRF.CallOpts, role, account)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCaller) NextRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "nextRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFSession) NextRequestId() (*big.Int, error) {
	return _FoundnoneVRF.Contract.NextRequestId(&_FoundnoneVRF.CallOpts)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) NextRequestId() (*big.Int, error) {
	return _FoundnoneVRF.Contract.NextRequestId(&_FoundnoneVRF.CallOpts)
}

// RequestBlockSet is a free data retrieval call binding the contract method 0x5face0f9.
//
// Solidity: function requestBlockSet(uint256 ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCaller) RequestBlockSet(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "requestBlockSet", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestBlockSet is a free data retrieval call binding the contract method 0x5face0f9.
//
// Solidity: function requestBlockSet(uint256 ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFSession) RequestBlockSet(arg0 *big.Int) (*big.Int, error) {
	return _FoundnoneVRF.Contract.RequestBlockSet(&_FoundnoneVRF.CallOpts, arg0)
}

// RequestBlockSet is a free data retrieval call binding the contract method 0x5face0f9.
//
// Solidity: function requestBlockSet(uint256 ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) RequestBlockSet(arg0 *big.Int) (*big.Int, error) {
	return _FoundnoneVRF.Contract.RequestBlockSet(&_FoundnoneVRF.CallOpts, arg0)
}

// RequestFee is a free data retrieval call binding the contract method 0xeb2e578b.
//
// Solidity: function requestFee() view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCaller) RequestFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "requestFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestFee is a free data retrieval call binding the contract method 0xeb2e578b.
//
// Solidity: function requestFee() view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFSession) RequestFee() (*big.Int, error) {
	return _FoundnoneVRF.Contract.RequestFee(&_FoundnoneVRF.CallOpts)
}

// RequestFee is a free data retrieval call binding the contract method 0xeb2e578b.
//
// Solidity: function requestFee() view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) RequestFee() (*big.Int, error) {
	return _FoundnoneVRF.Contract.RequestFee(&_FoundnoneVRF.CallOpts)
}

// RequestFeePaid is a free data retrieval call binding the contract method 0xef7cc992.
//
// Solidity: function requestFeePaid(uint256 ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCaller) RequestFeePaid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "requestFeePaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestFeePaid is a free data retrieval call binding the contract method 0xef7cc992.
//
// Solidity: function requestFeePaid(uint256 ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFSession) RequestFeePaid(arg0 *big.Int) (*big.Int, error) {
	return _FoundnoneVRF.Contract.RequestFeePaid(&_FoundnoneVRF.CallOpts, arg0)
}

// RequestFeePaid is a free data retrieval call binding the contract method 0xef7cc992.
//
// Solidity: function requestFeePaid(uint256 ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) RequestFeePaid(arg0 *big.Int) (*big.Int, error) {
	return _FoundnoneVRF.Contract.RequestFeePaid(&_FoundnoneVRF.CallOpts, arg0)
}

// Requesters is a free data retrieval call binding the contract method 0x5e1a6c17.
//
// Solidity: function requesters(uint256 ) view returns(address)
func (_FoundnoneVRF *FoundnoneVRFCaller) Requesters(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "requesters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Requesters is a free data retrieval call binding the contract method 0x5e1a6c17.
//
// Solidity: function requesters(uint256 ) view returns(address)
func (_FoundnoneVRF *FoundnoneVRFSession) Requesters(arg0 *big.Int) (common.Address, error) {
	return _FoundnoneVRF.Contract.Requesters(&_FoundnoneVRF.CallOpts, arg0)
}

// Requesters is a free data retrieval call binding the contract method 0x5e1a6c17.
//
// Solidity: function requesters(uint256 ) view returns(address)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) Requesters(arg0 *big.Int) (common.Address, error) {
	return _FoundnoneVRF.Contract.Requesters(&_FoundnoneVRF.CallOpts, arg0)
}

// RewardReceiverBalance is a free data retrieval call binding the contract method 0xbdfd7eb3.
//
// Solidity: function rewardReceiverBalance(address ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCaller) RewardReceiverBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "rewardReceiverBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardReceiverBalance is a free data retrieval call binding the contract method 0xbdfd7eb3.
//
// Solidity: function rewardReceiverBalance(address ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFSession) RewardReceiverBalance(arg0 common.Address) (*big.Int, error) {
	return _FoundnoneVRF.Contract.RewardReceiverBalance(&_FoundnoneVRF.CallOpts, arg0)
}

// RewardReceiverBalance is a free data retrieval call binding the contract method 0xbdfd7eb3.
//
// Solidity: function rewardReceiverBalance(address ) view returns(uint256)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) RewardReceiverBalance(arg0 common.Address) (*big.Int, error) {
	return _FoundnoneVRF.Contract.RewardReceiverBalance(&_FoundnoneVRF.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FoundnoneVRF *FoundnoneVRFCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FoundnoneVRF *FoundnoneVRFSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FoundnoneVRF.Contract.SupportsInterface(&_FoundnoneVRF.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FoundnoneVRF.Contract.SupportsInterface(&_FoundnoneVRF.CallOpts, interfaceId)
}

// VerifyProof is a free data retrieval call binding the contract method 0x1d5803fe.
//
// Solidity: function verifyProof(uint256[24] _proof, uint256[3] _pubSignals) view returns(bool)
func (_FoundnoneVRF *FoundnoneVRFCaller) VerifyProof(opts *bind.CallOpts, _proof [24]*big.Int, _pubSignals [3]*big.Int) (bool, error) {
	var out []interface{}
	err := _FoundnoneVRF.contract.Call(opts, &out, "verifyProof", _proof, _pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x1d5803fe.
//
// Solidity: function verifyProof(uint256[24] _proof, uint256[3] _pubSignals) view returns(bool)
func (_FoundnoneVRF *FoundnoneVRFSession) VerifyProof(_proof [24]*big.Int, _pubSignals [3]*big.Int) (bool, error) {
	return _FoundnoneVRF.Contract.VerifyProof(&_FoundnoneVRF.CallOpts, _proof, _pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0x1d5803fe.
//
// Solidity: function verifyProof(uint256[24] _proof, uint256[3] _pubSignals) view returns(bool)
func (_FoundnoneVRF *FoundnoneVRFCallerSession) VerifyProof(_proof [24]*big.Int, _pubSignals [3]*big.Int) (bool, error) {
	return _FoundnoneVRF.Contract.VerifyProof(&_FoundnoneVRF.CallOpts, _proof, _pubSignals)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FoundnoneVRF.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FoundnoneVRF *FoundnoneVRFSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.GrantRole(&_FoundnoneVRF.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.GrantRole(&_FoundnoneVRF.TransactOpts, role, account)
}

// RefundUnfulfilledRequest is a paid mutator transaction binding the contract method 0xca6f30a4.
//
// Solidity: function refundUnfulfilledRequest(uint256 _requestId) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactor) RefundUnfulfilledRequest(opts *bind.TransactOpts, _requestId *big.Int) (*types.Transaction, error) {
	return _FoundnoneVRF.contract.Transact(opts, "refundUnfulfilledRequest", _requestId)
}

// RefundUnfulfilledRequest is a paid mutator transaction binding the contract method 0xca6f30a4.
//
// Solidity: function refundUnfulfilledRequest(uint256 _requestId) returns()
func (_FoundnoneVRF *FoundnoneVRFSession) RefundUnfulfilledRequest(_requestId *big.Int) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.RefundUnfulfilledRequest(&_FoundnoneVRF.TransactOpts, _requestId)
}

// RefundUnfulfilledRequest is a paid mutator transaction binding the contract method 0xca6f30a4.
//
// Solidity: function refundUnfulfilledRequest(uint256 _requestId) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactorSession) RefundUnfulfilledRequest(_requestId *big.Int) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.RefundUnfulfilledRequest(&_FoundnoneVRF.TransactOpts, _requestId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FoundnoneVRF.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_FoundnoneVRF *FoundnoneVRFSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.RenounceRole(&_FoundnoneVRF.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.RenounceRole(&_FoundnoneVRF.TransactOpts, role, callerConfirmation)
}

// RequestRng is a paid mutator transaction binding the contract method 0x79663e49.
//
// Solidity: function requestRng() payable returns()
func (_FoundnoneVRF *FoundnoneVRFTransactor) RequestRng(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FoundnoneVRF.contract.Transact(opts, "requestRng")
}

// RequestRng is a paid mutator transaction binding the contract method 0x79663e49.
//
// Solidity: function requestRng() payable returns()
func (_FoundnoneVRF *FoundnoneVRFSession) RequestRng() (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.RequestRng(&_FoundnoneVRF.TransactOpts)
}

// RequestRng is a paid mutator transaction binding the contract method 0x79663e49.
//
// Solidity: function requestRng() payable returns()
func (_FoundnoneVRF *FoundnoneVRFTransactorSession) RequestRng() (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.RequestRng(&_FoundnoneVRF.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FoundnoneVRF.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FoundnoneVRF *FoundnoneVRFSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.RevokeRole(&_FoundnoneVRF.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.RevokeRole(&_FoundnoneVRF.TransactOpts, role, account)
}

// SetCommitment is a paid mutator transaction binding the contract method 0xbe9cd052.
//
// Solidity: function setCommitment(uint256 _commitment) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactor) SetCommitment(opts *bind.TransactOpts, _commitment *big.Int) (*types.Transaction, error) {
	return _FoundnoneVRF.contract.Transact(opts, "setCommitment", _commitment)
}

// SetCommitment is a paid mutator transaction binding the contract method 0xbe9cd052.
//
// Solidity: function setCommitment(uint256 _commitment) returns()
func (_FoundnoneVRF *FoundnoneVRFSession) SetCommitment(_commitment *big.Int) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.SetCommitment(&_FoundnoneVRF.TransactOpts, _commitment)
}

// SetCommitment is a paid mutator transaction binding the contract method 0xbe9cd052.
//
// Solidity: function setCommitment(uint256 _commitment) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactorSession) SetCommitment(_commitment *big.Int) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.SetCommitment(&_FoundnoneVRF.TransactOpts, _commitment)
}

// SetContractFeePercentage is a paid mutator transaction binding the contract method 0xa6504f8c.
//
// Solidity: function setContractFeePercentage(uint256 _newPercentage) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactor) SetContractFeePercentage(opts *bind.TransactOpts, _newPercentage *big.Int) (*types.Transaction, error) {
	return _FoundnoneVRF.contract.Transact(opts, "setContractFeePercentage", _newPercentage)
}

// SetContractFeePercentage is a paid mutator transaction binding the contract method 0xa6504f8c.
//
// Solidity: function setContractFeePercentage(uint256 _newPercentage) returns()
func (_FoundnoneVRF *FoundnoneVRFSession) SetContractFeePercentage(_newPercentage *big.Int) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.SetContractFeePercentage(&_FoundnoneVRF.TransactOpts, _newPercentage)
}

// SetContractFeePercentage is a paid mutator transaction binding the contract method 0xa6504f8c.
//
// Solidity: function setContractFeePercentage(uint256 _newPercentage) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactorSession) SetContractFeePercentage(_newPercentage *big.Int) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.SetContractFeePercentage(&_FoundnoneVRF.TransactOpts, _newPercentage)
}

// SetRequestFee is a paid mutator transaction binding the contract method 0xffb9c43f.
//
// Solidity: function setRequestFee(uint256 _newFee) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactor) SetRequestFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _FoundnoneVRF.contract.Transact(opts, "setRequestFee", _newFee)
}

// SetRequestFee is a paid mutator transaction binding the contract method 0xffb9c43f.
//
// Solidity: function setRequestFee(uint256 _newFee) returns()
func (_FoundnoneVRF *FoundnoneVRFSession) SetRequestFee(_newFee *big.Int) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.SetRequestFee(&_FoundnoneVRF.TransactOpts, _newFee)
}

// SetRequestFee is a paid mutator transaction binding the contract method 0xffb9c43f.
//
// Solidity: function setRequestFee(uint256 _newFee) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactorSession) SetRequestFee(_newFee *big.Int) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.SetRequestFee(&_FoundnoneVRF.TransactOpts, _newFee)
}

// SubmitEntropy is a paid mutator transaction binding the contract method 0x6804fe15.
//
// Solidity: function submitEntropy(uint256[24] _proof, uint256[3] _publicInputs, uint256 _requestId, address _rewardReceiver) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactor) SubmitEntropy(opts *bind.TransactOpts, _proof [24]*big.Int, _publicInputs [3]*big.Int, _requestId *big.Int, _rewardReceiver common.Address) (*types.Transaction, error) {
	return _FoundnoneVRF.contract.Transact(opts, "submitEntropy", _proof, _publicInputs, _requestId, _rewardReceiver)
}

// SubmitEntropy is a paid mutator transaction binding the contract method 0x6804fe15.
//
// Solidity: function submitEntropy(uint256[24] _proof, uint256[3] _publicInputs, uint256 _requestId, address _rewardReceiver) returns()
func (_FoundnoneVRF *FoundnoneVRFSession) SubmitEntropy(_proof [24]*big.Int, _publicInputs [3]*big.Int, _requestId *big.Int, _rewardReceiver common.Address) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.SubmitEntropy(&_FoundnoneVRF.TransactOpts, _proof, _publicInputs, _requestId, _rewardReceiver)
}

// SubmitEntropy is a paid mutator transaction binding the contract method 0x6804fe15.
//
// Solidity: function submitEntropy(uint256[24] _proof, uint256[3] _publicInputs, uint256 _requestId, address _rewardReceiver) returns()
func (_FoundnoneVRF *FoundnoneVRFTransactorSession) SubmitEntropy(_proof [24]*big.Int, _publicInputs [3]*big.Int, _requestId *big.Int, _rewardReceiver common.Address) (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.SubmitEntropy(&_FoundnoneVRF.TransactOpts, _proof, _publicInputs, _requestId, _rewardReceiver)
}

// WithdrawContractFees is a paid mutator transaction binding the contract method 0xea3de4cb.
//
// Solidity: function withdrawContractFees() returns()
func (_FoundnoneVRF *FoundnoneVRFTransactor) WithdrawContractFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FoundnoneVRF.contract.Transact(opts, "withdrawContractFees")
}

// WithdrawContractFees is a paid mutator transaction binding the contract method 0xea3de4cb.
//
// Solidity: function withdrawContractFees() returns()
func (_FoundnoneVRF *FoundnoneVRFSession) WithdrawContractFees() (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.WithdrawContractFees(&_FoundnoneVRF.TransactOpts)
}

// WithdrawContractFees is a paid mutator transaction binding the contract method 0xea3de4cb.
//
// Solidity: function withdrawContractFees() returns()
func (_FoundnoneVRF *FoundnoneVRFTransactorSession) WithdrawContractFees() (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.WithdrawContractFees(&_FoundnoneVRF.TransactOpts)
}

// WithdrawRewardReceiverBalance is a paid mutator transaction binding the contract method 0x5e089f99.
//
// Solidity: function withdrawRewardReceiverBalance() returns()
func (_FoundnoneVRF *FoundnoneVRFTransactor) WithdrawRewardReceiverBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FoundnoneVRF.contract.Transact(opts, "withdrawRewardReceiverBalance")
}

// WithdrawRewardReceiverBalance is a paid mutator transaction binding the contract method 0x5e089f99.
//
// Solidity: function withdrawRewardReceiverBalance() returns()
func (_FoundnoneVRF *FoundnoneVRFSession) WithdrawRewardReceiverBalance() (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.WithdrawRewardReceiverBalance(&_FoundnoneVRF.TransactOpts)
}

// WithdrawRewardReceiverBalance is a paid mutator transaction binding the contract method 0x5e089f99.
//
// Solidity: function withdrawRewardReceiverBalance() returns()
func (_FoundnoneVRF *FoundnoneVRFTransactorSession) WithdrawRewardReceiverBalance() (*types.Transaction, error) {
	return _FoundnoneVRF.Contract.WithdrawRewardReceiverBalance(&_FoundnoneVRF.TransactOpts)
}

// FoundnoneVRFContractFeePercentageUpdatedIterator is returned from FilterContractFeePercentageUpdated and is used to iterate over the raw logs and unpacked data for ContractFeePercentageUpdated events raised by the FoundnoneVRF contract.
type FoundnoneVRFContractFeePercentageUpdatedIterator struct {
	Event *FoundnoneVRFContractFeePercentageUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FoundnoneVRFContractFeePercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundnoneVRFContractFeePercentageUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FoundnoneVRFContractFeePercentageUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FoundnoneVRFContractFeePercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundnoneVRFContractFeePercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundnoneVRFContractFeePercentageUpdated represents a ContractFeePercentageUpdated event raised by the FoundnoneVRF contract.
type FoundnoneVRFContractFeePercentageUpdated struct {
	NewPercentage *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContractFeePercentageUpdated is a free log retrieval operation binding the contract event 0x88f49dcbaed0e4733413f55abc15dc393fc37709d1c423368bcf593e8ca61288.
//
// Solidity: event ContractFeePercentageUpdated(uint256 newPercentage)
func (_FoundnoneVRF *FoundnoneVRFFilterer) FilterContractFeePercentageUpdated(opts *bind.FilterOpts) (*FoundnoneVRFContractFeePercentageUpdatedIterator, error) {

	logs, sub, err := _FoundnoneVRF.contract.FilterLogs(opts, "ContractFeePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &FoundnoneVRFContractFeePercentageUpdatedIterator{contract: _FoundnoneVRF.contract, event: "ContractFeePercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchContractFeePercentageUpdated is a free log subscription operation binding the contract event 0x88f49dcbaed0e4733413f55abc15dc393fc37709d1c423368bcf593e8ca61288.
//
// Solidity: event ContractFeePercentageUpdated(uint256 newPercentage)
func (_FoundnoneVRF *FoundnoneVRFFilterer) WatchContractFeePercentageUpdated(opts *bind.WatchOpts, sink chan<- *FoundnoneVRFContractFeePercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _FoundnoneVRF.contract.WatchLogs(opts, "ContractFeePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundnoneVRFContractFeePercentageUpdated)
				if err := _FoundnoneVRF.contract.UnpackLog(event, "ContractFeePercentageUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractFeePercentageUpdated is a log parse operation binding the contract event 0x88f49dcbaed0e4733413f55abc15dc393fc37709d1c423368bcf593e8ca61288.
//
// Solidity: event ContractFeePercentageUpdated(uint256 newPercentage)
func (_FoundnoneVRF *FoundnoneVRFFilterer) ParseContractFeePercentageUpdated(log types.Log) (*FoundnoneVRFContractFeePercentageUpdated, error) {
	event := new(FoundnoneVRFContractFeePercentageUpdated)
	if err := _FoundnoneVRF.contract.UnpackLog(event, "ContractFeePercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundnoneVRFContractFeesWithdrawnIterator is returned from FilterContractFeesWithdrawn and is used to iterate over the raw logs and unpacked data for ContractFeesWithdrawn events raised by the FoundnoneVRF contract.
type FoundnoneVRFContractFeesWithdrawnIterator struct {
	Event *FoundnoneVRFContractFeesWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FoundnoneVRFContractFeesWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundnoneVRFContractFeesWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FoundnoneVRFContractFeesWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FoundnoneVRFContractFeesWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundnoneVRFContractFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundnoneVRFContractFeesWithdrawn represents a ContractFeesWithdrawn event raised by the FoundnoneVRF contract.
type FoundnoneVRFContractFeesWithdrawn struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterContractFeesWithdrawn is a free log retrieval operation binding the contract event 0x3d9264dd79c7dda789bd13ca13e81718ba78de6e6134e52be85fb6208347b013.
//
// Solidity: event ContractFeesWithdrawn(uint256 amount)
func (_FoundnoneVRF *FoundnoneVRFFilterer) FilterContractFeesWithdrawn(opts *bind.FilterOpts) (*FoundnoneVRFContractFeesWithdrawnIterator, error) {

	logs, sub, err := _FoundnoneVRF.contract.FilterLogs(opts, "ContractFeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &FoundnoneVRFContractFeesWithdrawnIterator{contract: _FoundnoneVRF.contract, event: "ContractFeesWithdrawn", logs: logs, sub: sub}, nil
}

// WatchContractFeesWithdrawn is a free log subscription operation binding the contract event 0x3d9264dd79c7dda789bd13ca13e81718ba78de6e6134e52be85fb6208347b013.
//
// Solidity: event ContractFeesWithdrawn(uint256 amount)
func (_FoundnoneVRF *FoundnoneVRFFilterer) WatchContractFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *FoundnoneVRFContractFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _FoundnoneVRF.contract.WatchLogs(opts, "ContractFeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundnoneVRFContractFeesWithdrawn)
				if err := _FoundnoneVRF.contract.UnpackLog(event, "ContractFeesWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractFeesWithdrawn is a log parse operation binding the contract event 0x3d9264dd79c7dda789bd13ca13e81718ba78de6e6134e52be85fb6208347b013.
//
// Solidity: event ContractFeesWithdrawn(uint256 amount)
func (_FoundnoneVRF *FoundnoneVRFFilterer) ParseContractFeesWithdrawn(log types.Log) (*FoundnoneVRFContractFeesWithdrawn, error) {
	event := new(FoundnoneVRFContractFeesWithdrawn)
	if err := _FoundnoneVRF.contract.UnpackLog(event, "ContractFeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundnoneVRFRequestFeeUpdatedIterator is returned from FilterRequestFeeUpdated and is used to iterate over the raw logs and unpacked data for RequestFeeUpdated events raised by the FoundnoneVRF contract.
type FoundnoneVRFRequestFeeUpdatedIterator struct {
	Event *FoundnoneVRFRequestFeeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FoundnoneVRFRequestFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundnoneVRFRequestFeeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FoundnoneVRFRequestFeeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FoundnoneVRFRequestFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundnoneVRFRequestFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundnoneVRFRequestFeeUpdated represents a RequestFeeUpdated event raised by the FoundnoneVRF contract.
type FoundnoneVRFRequestFeeUpdated struct {
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRequestFeeUpdated is a free log retrieval operation binding the contract event 0x07695b29775442d5e4836f94223aa3460d93774d5cb9e03119815f418e2a61c4.
//
// Solidity: event RequestFeeUpdated(uint256 newFee)
func (_FoundnoneVRF *FoundnoneVRFFilterer) FilterRequestFeeUpdated(opts *bind.FilterOpts) (*FoundnoneVRFRequestFeeUpdatedIterator, error) {

	logs, sub, err := _FoundnoneVRF.contract.FilterLogs(opts, "RequestFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &FoundnoneVRFRequestFeeUpdatedIterator{contract: _FoundnoneVRF.contract, event: "RequestFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchRequestFeeUpdated is a free log subscription operation binding the contract event 0x07695b29775442d5e4836f94223aa3460d93774d5cb9e03119815f418e2a61c4.
//
// Solidity: event RequestFeeUpdated(uint256 newFee)
func (_FoundnoneVRF *FoundnoneVRFFilterer) WatchRequestFeeUpdated(opts *bind.WatchOpts, sink chan<- *FoundnoneVRFRequestFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _FoundnoneVRF.contract.WatchLogs(opts, "RequestFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundnoneVRFRequestFeeUpdated)
				if err := _FoundnoneVRF.contract.UnpackLog(event, "RequestFeeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestFeeUpdated is a log parse operation binding the contract event 0x07695b29775442d5e4836f94223aa3460d93774d5cb9e03119815f418e2a61c4.
//
// Solidity: event RequestFeeUpdated(uint256 newFee)
func (_FoundnoneVRF *FoundnoneVRFFilterer) ParseRequestFeeUpdated(log types.Log) (*FoundnoneVRFRequestFeeUpdated, error) {
	event := new(FoundnoneVRFRequestFeeUpdated)
	if err := _FoundnoneVRF.contract.UnpackLog(event, "RequestFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundnoneVRFRequestFulfilledIterator is returned from FilterRequestFulfilled and is used to iterate over the raw logs and unpacked data for RequestFulfilled events raised by the FoundnoneVRF contract.
type FoundnoneVRFRequestFulfilledIterator struct {
	Event *FoundnoneVRFRequestFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FoundnoneVRFRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundnoneVRFRequestFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FoundnoneVRFRequestFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FoundnoneVRFRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundnoneVRFRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundnoneVRFRequestFulfilled represents a RequestFulfilled event raised by the FoundnoneVRF contract.
type FoundnoneVRFRequestFulfilled struct {
	RequestId      *big.Int
	RewardReceiver common.Address
	Proof          [24]*big.Int
	PublicInputs   [3]*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRequestFulfilled is a free log retrieval operation binding the contract event 0x11f7f1fb4febf9b59cf5c95f67764124b90fdf34078f8289de9cad7e57c1cfdc.
//
// Solidity: event RequestFulfilled(uint256 indexed requestId, address rewardReceiver, uint256[24] proof, uint256[3] publicInputs)
func (_FoundnoneVRF *FoundnoneVRFFilterer) FilterRequestFulfilled(opts *bind.FilterOpts, requestId []*big.Int) (*FoundnoneVRFRequestFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FoundnoneVRF.contract.FilterLogs(opts, "RequestFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &FoundnoneVRFRequestFulfilledIterator{contract: _FoundnoneVRF.contract, event: "RequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchRequestFulfilled is a free log subscription operation binding the contract event 0x11f7f1fb4febf9b59cf5c95f67764124b90fdf34078f8289de9cad7e57c1cfdc.
//
// Solidity: event RequestFulfilled(uint256 indexed requestId, address rewardReceiver, uint256[24] proof, uint256[3] publicInputs)
func (_FoundnoneVRF *FoundnoneVRFFilterer) WatchRequestFulfilled(opts *bind.WatchOpts, sink chan<- *FoundnoneVRFRequestFulfilled, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FoundnoneVRF.contract.WatchLogs(opts, "RequestFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundnoneVRFRequestFulfilled)
				if err := _FoundnoneVRF.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestFulfilled is a log parse operation binding the contract event 0x11f7f1fb4febf9b59cf5c95f67764124b90fdf34078f8289de9cad7e57c1cfdc.
//
// Solidity: event RequestFulfilled(uint256 indexed requestId, address rewardReceiver, uint256[24] proof, uint256[3] publicInputs)
func (_FoundnoneVRF *FoundnoneVRFFilterer) ParseRequestFulfilled(log types.Log) (*FoundnoneVRFRequestFulfilled, error) {
	event := new(FoundnoneVRFRequestFulfilled)
	if err := _FoundnoneVRF.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundnoneVRFRewardReceiverBalanceWithdrawnIterator is returned from FilterRewardReceiverBalanceWithdrawn and is used to iterate over the raw logs and unpacked data for RewardReceiverBalanceWithdrawn events raised by the FoundnoneVRF contract.
type FoundnoneVRFRewardReceiverBalanceWithdrawnIterator struct {
	Event *FoundnoneVRFRewardReceiverBalanceWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FoundnoneVRFRewardReceiverBalanceWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundnoneVRFRewardReceiverBalanceWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FoundnoneVRFRewardReceiverBalanceWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FoundnoneVRFRewardReceiverBalanceWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundnoneVRFRewardReceiverBalanceWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundnoneVRFRewardReceiverBalanceWithdrawn represents a RewardReceiverBalanceWithdrawn event raised by the FoundnoneVRF contract.
type FoundnoneVRFRewardReceiverBalanceWithdrawn struct {
	RewardReceiver common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardReceiverBalanceWithdrawn is a free log retrieval operation binding the contract event 0x0b40a73848f6d60e086f98c5ba41e8d98e0f0a47db8ae16597c266e3a5c40cba.
//
// Solidity: event RewardReceiverBalanceWithdrawn(address indexed rewardReceiver, uint256 amount)
func (_FoundnoneVRF *FoundnoneVRFFilterer) FilterRewardReceiverBalanceWithdrawn(opts *bind.FilterOpts, rewardReceiver []common.Address) (*FoundnoneVRFRewardReceiverBalanceWithdrawnIterator, error) {

	var rewardReceiverRule []interface{}
	for _, rewardReceiverItem := range rewardReceiver {
		rewardReceiverRule = append(rewardReceiverRule, rewardReceiverItem)
	}

	logs, sub, err := _FoundnoneVRF.contract.FilterLogs(opts, "RewardReceiverBalanceWithdrawn", rewardReceiverRule)
	if err != nil {
		return nil, err
	}
	return &FoundnoneVRFRewardReceiverBalanceWithdrawnIterator{contract: _FoundnoneVRF.contract, event: "RewardReceiverBalanceWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRewardReceiverBalanceWithdrawn is a free log subscription operation binding the contract event 0x0b40a73848f6d60e086f98c5ba41e8d98e0f0a47db8ae16597c266e3a5c40cba.
//
// Solidity: event RewardReceiverBalanceWithdrawn(address indexed rewardReceiver, uint256 amount)
func (_FoundnoneVRF *FoundnoneVRFFilterer) WatchRewardReceiverBalanceWithdrawn(opts *bind.WatchOpts, sink chan<- *FoundnoneVRFRewardReceiverBalanceWithdrawn, rewardReceiver []common.Address) (event.Subscription, error) {

	var rewardReceiverRule []interface{}
	for _, rewardReceiverItem := range rewardReceiver {
		rewardReceiverRule = append(rewardReceiverRule, rewardReceiverItem)
	}

	logs, sub, err := _FoundnoneVRF.contract.WatchLogs(opts, "RewardReceiverBalanceWithdrawn", rewardReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundnoneVRFRewardReceiverBalanceWithdrawn)
				if err := _FoundnoneVRF.contract.UnpackLog(event, "RewardReceiverBalanceWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardReceiverBalanceWithdrawn is a log parse operation binding the contract event 0x0b40a73848f6d60e086f98c5ba41e8d98e0f0a47db8ae16597c266e3a5c40cba.
//
// Solidity: event RewardReceiverBalanceWithdrawn(address indexed rewardReceiver, uint256 amount)
func (_FoundnoneVRF *FoundnoneVRFFilterer) ParseRewardReceiverBalanceWithdrawn(log types.Log) (*FoundnoneVRFRewardReceiverBalanceWithdrawn, error) {
	event := new(FoundnoneVRFRewardReceiverBalanceWithdrawn)
	if err := _FoundnoneVRF.contract.UnpackLog(event, "RewardReceiverBalanceWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundnoneVRFRngRequestedIterator is returned from FilterRngRequested and is used to iterate over the raw logs and unpacked data for RngRequested events raised by the FoundnoneVRF contract.
type FoundnoneVRFRngRequestedIterator struct {
	Event *FoundnoneVRFRngRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FoundnoneVRFRngRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundnoneVRFRngRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FoundnoneVRFRngRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FoundnoneVRFRngRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundnoneVRFRngRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundnoneVRFRngRequested represents a RngRequested event raised by the FoundnoneVRF contract.
type FoundnoneVRFRngRequested struct {
	RequestId *big.Int
	BlockHash [32]byte
	Requester common.Address
	FeePaid   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRngRequested is a free log retrieval operation binding the contract event 0xebd7660295e4ce331b3cfbc0706171ac8de0c07c415924affdcf42ddc11774dc.
//
// Solidity: event RngRequested(uint256 requestId, bytes32 blockHash, address requester, uint256 feePaid)
func (_FoundnoneVRF *FoundnoneVRFFilterer) FilterRngRequested(opts *bind.FilterOpts) (*FoundnoneVRFRngRequestedIterator, error) {

	logs, sub, err := _FoundnoneVRF.contract.FilterLogs(opts, "RngRequested")
	if err != nil {
		return nil, err
	}
	return &FoundnoneVRFRngRequestedIterator{contract: _FoundnoneVRF.contract, event: "RngRequested", logs: logs, sub: sub}, nil
}

// WatchRngRequested is a free log subscription operation binding the contract event 0xebd7660295e4ce331b3cfbc0706171ac8de0c07c415924affdcf42ddc11774dc.
//
// Solidity: event RngRequested(uint256 requestId, bytes32 blockHash, address requester, uint256 feePaid)
func (_FoundnoneVRF *FoundnoneVRFFilterer) WatchRngRequested(opts *bind.WatchOpts, sink chan<- *FoundnoneVRFRngRequested) (event.Subscription, error) {

	logs, sub, err := _FoundnoneVRF.contract.WatchLogs(opts, "RngRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundnoneVRFRngRequested)
				if err := _FoundnoneVRF.contract.UnpackLog(event, "RngRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRngRequested is a log parse operation binding the contract event 0xebd7660295e4ce331b3cfbc0706171ac8de0c07c415924affdcf42ddc11774dc.
//
// Solidity: event RngRequested(uint256 requestId, bytes32 blockHash, address requester, uint256 feePaid)
func (_FoundnoneVRF *FoundnoneVRFFilterer) ParseRngRequested(log types.Log) (*FoundnoneVRFRngRequested, error) {
	event := new(FoundnoneVRFRngRequested)
	if err := _FoundnoneVRF.contract.UnpackLog(event, "RngRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundnoneVRFRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the FoundnoneVRF contract.
type FoundnoneVRFRoleAdminChangedIterator struct {
	Event *FoundnoneVRFRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FoundnoneVRFRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundnoneVRFRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FoundnoneVRFRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FoundnoneVRFRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundnoneVRFRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundnoneVRFRoleAdminChanged represents a RoleAdminChanged event raised by the FoundnoneVRF contract.
type FoundnoneVRFRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FoundnoneVRF *FoundnoneVRFFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FoundnoneVRFRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _FoundnoneVRF.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FoundnoneVRFRoleAdminChangedIterator{contract: _FoundnoneVRF.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FoundnoneVRF *FoundnoneVRFFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FoundnoneVRFRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _FoundnoneVRF.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundnoneVRFRoleAdminChanged)
				if err := _FoundnoneVRF.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FoundnoneVRF *FoundnoneVRFFilterer) ParseRoleAdminChanged(log types.Log) (*FoundnoneVRFRoleAdminChanged, error) {
	event := new(FoundnoneVRFRoleAdminChanged)
	if err := _FoundnoneVRF.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundnoneVRFRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the FoundnoneVRF contract.
type FoundnoneVRFRoleGrantedIterator struct {
	Event *FoundnoneVRFRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FoundnoneVRFRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundnoneVRFRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FoundnoneVRFRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FoundnoneVRFRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundnoneVRFRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundnoneVRFRoleGranted represents a RoleGranted event raised by the FoundnoneVRF contract.
type FoundnoneVRFRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FoundnoneVRF *FoundnoneVRFFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FoundnoneVRFRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FoundnoneVRF.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FoundnoneVRFRoleGrantedIterator{contract: _FoundnoneVRF.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FoundnoneVRF *FoundnoneVRFFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FoundnoneVRFRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FoundnoneVRF.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundnoneVRFRoleGranted)
				if err := _FoundnoneVRF.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FoundnoneVRF *FoundnoneVRFFilterer) ParseRoleGranted(log types.Log) (*FoundnoneVRFRoleGranted, error) {
	event := new(FoundnoneVRFRoleGranted)
	if err := _FoundnoneVRF.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoundnoneVRFRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the FoundnoneVRF contract.
type FoundnoneVRFRoleRevokedIterator struct {
	Event *FoundnoneVRFRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FoundnoneVRFRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundnoneVRFRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FoundnoneVRFRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FoundnoneVRFRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundnoneVRFRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundnoneVRFRoleRevoked represents a RoleRevoked event raised by the FoundnoneVRF contract.
type FoundnoneVRFRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FoundnoneVRF *FoundnoneVRFFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FoundnoneVRFRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FoundnoneVRF.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FoundnoneVRFRoleRevokedIterator{contract: _FoundnoneVRF.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FoundnoneVRF *FoundnoneVRFFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FoundnoneVRFRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FoundnoneVRF.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundnoneVRFRoleRevoked)
				if err := _FoundnoneVRF.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FoundnoneVRF *FoundnoneVRFFilterer) ParseRoleRevoked(log types.Log) (*FoundnoneVRFRoleRevoked, error) {
	event := new(FoundnoneVRFRoleRevoked)
	if err := _FoundnoneVRF.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
