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

// AbiMetaData contains all meta data concerning the Abi contract.
var AbiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adminRole\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitmentAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitmentInUse\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommitmentBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFeePercentage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequester\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSeedOrBlockHashUnavailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestAlreadyFulfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestNotFulfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestStillValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequesterNotWhitelisted\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPercentage\",\"type\":\"uint256\"}],\"name\":\"ContractFeePercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ContractFeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"RequestFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"callbackSuccess\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[24]\",\"name\":\"proof\",\"type\":\"uint256[24]\"},{\"indexed\":false,\"internalType\":\"uint256[3]\",\"name\":\"publicInputs\",\"type\":\"uint256[3]\"}],\"name\":\"RequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RequestRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardReceiverBalanceWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePaid\",\"type\":\"uint256\"}],\"name\":\"RngRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"commitmentBlockSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"commitmentInUse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractFeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"entropies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getEntropy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardReceiver\",\"type\":\"address\"}],\"name\":\"getRewardReceiverBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"refundUnfulfilledRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"request\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"callbackAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"requestBlockSet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callbackAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"}],\"name\":\"requestRng\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardReceiverBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_commitment\",\"type\":\"uint256\"}],\"name\":\"setCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPercentage\",\"type\":\"uint256\"}],\"name\":\"setContractFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"setRequestFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_whitelistedRequester\",\"type\":\"address\"}],\"name\":\"setWhitelistedRequester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[24]\",\"name\":\"_proof\",\"type\":\"uint256[24]\"},{\"internalType\":\"uint256[3]\",\"name\":\"_publicInputs\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewardReceiver\",\"type\":\"address\"}],\"name\":\"submitEntropy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[24]\",\"name\":\"_proof\",\"type\":\"uint256[24]\"},{\"internalType\":\"uint256[3]\",\"name\":\"_pubSignals\",\"type\":\"uint256[3]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistedRequester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawContractFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawRewardReceiverBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AbiABI is the input ABI used to generate the binding from.
// Deprecated: Use AbiMetaData.ABI instead.
var AbiABI = AbiMetaData.ABI

// Abi is an auto generated Go binding around an Ethereum contract.
type Abi struct {
	AbiCaller     // Read-only binding to the contract
	AbiTransactor // Write-only binding to the contract
	AbiFilterer   // Log filterer for contract events
}

// AbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiSession struct {
	Contract     *Abi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiCallerSession struct {
	Contract *AbiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiTransactorSession struct {
	Contract     *AbiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiRaw struct {
	Contract *Abi // Generic contract binding to access the raw methods on
}

// AbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiCallerRaw struct {
	Contract *AbiCaller // Generic read-only contract binding to access the raw methods on
}

// AbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiTransactorRaw struct {
	Contract *AbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbi creates a new instance of Abi, bound to a specific deployed contract.
func NewAbi(address common.Address, backend bind.ContractBackend) (*Abi, error) {
	contract, err := bindAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abi{AbiCaller: AbiCaller{contract: contract}, AbiTransactor: AbiTransactor{contract: contract}, AbiFilterer: AbiFilterer{contract: contract}}, nil
}

// NewAbiCaller creates a new read-only instance of Abi, bound to a specific deployed contract.
func NewAbiCaller(address common.Address, caller bind.ContractCaller) (*AbiCaller, error) {
	contract, err := bindAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiCaller{contract: contract}, nil
}

// NewAbiTransactor creates a new write-only instance of Abi, bound to a specific deployed contract.
func NewAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiTransactor, error) {
	contract, err := bindAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiTransactor{contract: contract}, nil
}

// NewAbiFilterer creates a new log filterer instance of Abi, bound to a specific deployed contract.
func NewAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiFilterer, error) {
	contract, err := bindAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiFilterer{contract: contract}, nil
}

// bindAbi binds a generic wrapper to an already deployed contract.
func bindAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.AbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Abi *AbiCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Abi *AbiSession) ADMINROLE() ([32]byte, error) {
	return _Abi.Contract.ADMINROLE(&_Abi.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Abi *AbiCallerSession) ADMINROLE() ([32]byte, error) {
	return _Abi.Contract.ADMINROLE(&_Abi.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Abi *AbiCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Abi *AbiSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Abi.Contract.DEFAULTADMINROLE(&_Abi.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Abi *AbiCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Abi.Contract.DEFAULTADMINROLE(&_Abi.CallOpts)
}

// CommitmentBlockSet is a free data retrieval call binding the contract method 0x1ad6bdd6.
//
// Solidity: function commitmentBlockSet(address ) view returns(uint256)
func (_Abi *AbiCaller) CommitmentBlockSet(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "commitmentBlockSet", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitmentBlockSet is a free data retrieval call binding the contract method 0x1ad6bdd6.
//
// Solidity: function commitmentBlockSet(address ) view returns(uint256)
func (_Abi *AbiSession) CommitmentBlockSet(arg0 common.Address) (*big.Int, error) {
	return _Abi.Contract.CommitmentBlockSet(&_Abi.CallOpts, arg0)
}

// CommitmentBlockSet is a free data retrieval call binding the contract method 0x1ad6bdd6.
//
// Solidity: function commitmentBlockSet(address ) view returns(uint256)
func (_Abi *AbiCallerSession) CommitmentBlockSet(arg0 common.Address) (*big.Int, error) {
	return _Abi.Contract.CommitmentBlockSet(&_Abi.CallOpts, arg0)
}

// CommitmentInUse is a free data retrieval call binding the contract method 0x21b4dae0.
//
// Solidity: function commitmentInUse(uint256 ) view returns(bool)
func (_Abi *AbiCaller) CommitmentInUse(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "commitmentInUse", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CommitmentInUse is a free data retrieval call binding the contract method 0x21b4dae0.
//
// Solidity: function commitmentInUse(uint256 ) view returns(bool)
func (_Abi *AbiSession) CommitmentInUse(arg0 *big.Int) (bool, error) {
	return _Abi.Contract.CommitmentInUse(&_Abi.CallOpts, arg0)
}

// CommitmentInUse is a free data retrieval call binding the contract method 0x21b4dae0.
//
// Solidity: function commitmentInUse(uint256 ) view returns(bool)
func (_Abi *AbiCallerSession) CommitmentInUse(arg0 *big.Int) (bool, error) {
	return _Abi.Contract.CommitmentInUse(&_Abi.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0xe8fcf723.
//
// Solidity: function commitments(address ) view returns(uint256)
func (_Abi *AbiCaller) Commitments(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0xe8fcf723.
//
// Solidity: function commitments(address ) view returns(uint256)
func (_Abi *AbiSession) Commitments(arg0 common.Address) (*big.Int, error) {
	return _Abi.Contract.Commitments(&_Abi.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0xe8fcf723.
//
// Solidity: function commitments(address ) view returns(uint256)
func (_Abi *AbiCallerSession) Commitments(arg0 common.Address) (*big.Int, error) {
	return _Abi.Contract.Commitments(&_Abi.CallOpts, arg0)
}

// ContractFeeBalance is a free data retrieval call binding the contract method 0xb160f9cc.
//
// Solidity: function contractFeeBalance() view returns(uint256)
func (_Abi *AbiCaller) ContractFeeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "contractFeeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractFeeBalance is a free data retrieval call binding the contract method 0xb160f9cc.
//
// Solidity: function contractFeeBalance() view returns(uint256)
func (_Abi *AbiSession) ContractFeeBalance() (*big.Int, error) {
	return _Abi.Contract.ContractFeeBalance(&_Abi.CallOpts)
}

// ContractFeeBalance is a free data retrieval call binding the contract method 0xb160f9cc.
//
// Solidity: function contractFeeBalance() view returns(uint256)
func (_Abi *AbiCallerSession) ContractFeeBalance() (*big.Int, error) {
	return _Abi.Contract.ContractFeeBalance(&_Abi.CallOpts)
}

// ContractFeePercentage is a free data retrieval call binding the contract method 0xdbd8987c.
//
// Solidity: function contractFeePercentage() view returns(uint256)
func (_Abi *AbiCaller) ContractFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "contractFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractFeePercentage is a free data retrieval call binding the contract method 0xdbd8987c.
//
// Solidity: function contractFeePercentage() view returns(uint256)
func (_Abi *AbiSession) ContractFeePercentage() (*big.Int, error) {
	return _Abi.Contract.ContractFeePercentage(&_Abi.CallOpts)
}

// ContractFeePercentage is a free data retrieval call binding the contract method 0xdbd8987c.
//
// Solidity: function contractFeePercentage() view returns(uint256)
func (_Abi *AbiCallerSession) ContractFeePercentage() (*big.Int, error) {
	return _Abi.Contract.ContractFeePercentage(&_Abi.CallOpts)
}

// Entropies is a free data retrieval call binding the contract method 0xe2ac9d5b.
//
// Solidity: function entropies(uint256 ) view returns(uint256)
func (_Abi *AbiCaller) Entropies(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "entropies", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Entropies is a free data retrieval call binding the contract method 0xe2ac9d5b.
//
// Solidity: function entropies(uint256 ) view returns(uint256)
func (_Abi *AbiSession) Entropies(arg0 *big.Int) (*big.Int, error) {
	return _Abi.Contract.Entropies(&_Abi.CallOpts, arg0)
}

// Entropies is a free data retrieval call binding the contract method 0xe2ac9d5b.
//
// Solidity: function entropies(uint256 ) view returns(uint256)
func (_Abi *AbiCallerSession) Entropies(arg0 *big.Int) (*big.Int, error) {
	return _Abi.Contract.Entropies(&_Abi.CallOpts, arg0)
}

// GetEntropy is a free data retrieval call binding the contract method 0x74e6a469.
//
// Solidity: function getEntropy(uint256 _requestId) view returns(uint256)
func (_Abi *AbiCaller) GetEntropy(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getEntropy", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEntropy is a free data retrieval call binding the contract method 0x74e6a469.
//
// Solidity: function getEntropy(uint256 _requestId) view returns(uint256)
func (_Abi *AbiSession) GetEntropy(_requestId *big.Int) (*big.Int, error) {
	return _Abi.Contract.GetEntropy(&_Abi.CallOpts, _requestId)
}

// GetEntropy is a free data retrieval call binding the contract method 0x74e6a469.
//
// Solidity: function getEntropy(uint256 _requestId) view returns(uint256)
func (_Abi *AbiCallerSession) GetEntropy(_requestId *big.Int) (*big.Int, error) {
	return _Abi.Contract.GetEntropy(&_Abi.CallOpts, _requestId)
}

// GetRewardReceiverBalance is a free data retrieval call binding the contract method 0x64b0c523.
//
// Solidity: function getRewardReceiverBalance(address _rewardReceiver) view returns(uint256)
func (_Abi *AbiCaller) GetRewardReceiverBalance(opts *bind.CallOpts, _rewardReceiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getRewardReceiverBalance", _rewardReceiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardReceiverBalance is a free data retrieval call binding the contract method 0x64b0c523.
//
// Solidity: function getRewardReceiverBalance(address _rewardReceiver) view returns(uint256)
func (_Abi *AbiSession) GetRewardReceiverBalance(_rewardReceiver common.Address) (*big.Int, error) {
	return _Abi.Contract.GetRewardReceiverBalance(&_Abi.CallOpts, _rewardReceiver)
}

// GetRewardReceiverBalance is a free data retrieval call binding the contract method 0x64b0c523.
//
// Solidity: function getRewardReceiverBalance(address _rewardReceiver) view returns(uint256)
func (_Abi *AbiCallerSession) GetRewardReceiverBalance(_rewardReceiver common.Address) (*big.Int, error) {
	return _Abi.Contract.GetRewardReceiverBalance(&_Abi.CallOpts, _rewardReceiver)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Abi *AbiCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Abi *AbiSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Abi.Contract.GetRoleAdmin(&_Abi.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Abi *AbiCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Abi.Contract.GetRoleAdmin(&_Abi.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Abi *AbiCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Abi *AbiSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Abi.Contract.HasRole(&_Abi.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Abi *AbiCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Abi.Contract.HasRole(&_Abi.CallOpts, role, account)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_Abi *AbiCaller) NextRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "nextRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_Abi *AbiSession) NextRequestId() (*big.Int, error) {
	return _Abi.Contract.NextRequestId(&_Abi.CallOpts)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_Abi *AbiCallerSession) NextRequestId() (*big.Int, error) {
	return _Abi.Contract.NextRequestId(&_Abi.CallOpts)
}

// Request is a free data retrieval call binding the contract method 0xd845a4b3.
//
// Solidity: function request(uint256 ) view returns(address callbackAddress, uint32 callbackGasLimit, uint256 requestBlockSet, uint256 requestFeePaid, address requester)
func (_Abi *AbiCaller) Request(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CallbackAddress  common.Address
	CallbackGasLimit uint32
	RequestBlockSet  *big.Int
	RequestFeePaid   *big.Int
	Requester        common.Address
}, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "request", arg0)

	outstruct := new(struct {
		CallbackAddress  common.Address
		CallbackGasLimit uint32
		RequestBlockSet  *big.Int
		RequestFeePaid   *big.Int
		Requester        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CallbackAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CallbackGasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.RequestBlockSet = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RequestFeePaid = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Requester = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Request is a free data retrieval call binding the contract method 0xd845a4b3.
//
// Solidity: function request(uint256 ) view returns(address callbackAddress, uint32 callbackGasLimit, uint256 requestBlockSet, uint256 requestFeePaid, address requester)
func (_Abi *AbiSession) Request(arg0 *big.Int) (struct {
	CallbackAddress  common.Address
	CallbackGasLimit uint32
	RequestBlockSet  *big.Int
	RequestFeePaid   *big.Int
	Requester        common.Address
}, error) {
	return _Abi.Contract.Request(&_Abi.CallOpts, arg0)
}

// Request is a free data retrieval call binding the contract method 0xd845a4b3.
//
// Solidity: function request(uint256 ) view returns(address callbackAddress, uint32 callbackGasLimit, uint256 requestBlockSet, uint256 requestFeePaid, address requester)
func (_Abi *AbiCallerSession) Request(arg0 *big.Int) (struct {
	CallbackAddress  common.Address
	CallbackGasLimit uint32
	RequestBlockSet  *big.Int
	RequestFeePaid   *big.Int
	Requester        common.Address
}, error) {
	return _Abi.Contract.Request(&_Abi.CallOpts, arg0)
}

// RequestFee is a free data retrieval call binding the contract method 0xeb2e578b.
//
// Solidity: function requestFee() view returns(uint256)
func (_Abi *AbiCaller) RequestFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "requestFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestFee is a free data retrieval call binding the contract method 0xeb2e578b.
//
// Solidity: function requestFee() view returns(uint256)
func (_Abi *AbiSession) RequestFee() (*big.Int, error) {
	return _Abi.Contract.RequestFee(&_Abi.CallOpts)
}

// RequestFee is a free data retrieval call binding the contract method 0xeb2e578b.
//
// Solidity: function requestFee() view returns(uint256)
func (_Abi *AbiCallerSession) RequestFee() (*big.Int, error) {
	return _Abi.Contract.RequestFee(&_Abi.CallOpts)
}

// RewardReceiverBalance is a free data retrieval call binding the contract method 0xbdfd7eb3.
//
// Solidity: function rewardReceiverBalance(address ) view returns(uint256)
func (_Abi *AbiCaller) RewardReceiverBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "rewardReceiverBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardReceiverBalance is a free data retrieval call binding the contract method 0xbdfd7eb3.
//
// Solidity: function rewardReceiverBalance(address ) view returns(uint256)
func (_Abi *AbiSession) RewardReceiverBalance(arg0 common.Address) (*big.Int, error) {
	return _Abi.Contract.RewardReceiverBalance(&_Abi.CallOpts, arg0)
}

// RewardReceiverBalance is a free data retrieval call binding the contract method 0xbdfd7eb3.
//
// Solidity: function rewardReceiverBalance(address ) view returns(uint256)
func (_Abi *AbiCallerSession) RewardReceiverBalance(arg0 common.Address) (*big.Int, error) {
	return _Abi.Contract.RewardReceiverBalance(&_Abi.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Abi *AbiCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Abi *AbiSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Abi.Contract.SupportsInterface(&_Abi.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Abi *AbiCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Abi.Contract.SupportsInterface(&_Abi.CallOpts, interfaceId)
}

// VerifyProof is a free data retrieval call binding the contract method 0x1d5803fe.
//
// Solidity: function verifyProof(uint256[24] _proof, uint256[3] _pubSignals) view returns(bool)
func (_Abi *AbiCaller) VerifyProof(opts *bind.CallOpts, _proof [24]*big.Int, _pubSignals [3]*big.Int) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "verifyProof", _proof, _pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x1d5803fe.
//
// Solidity: function verifyProof(uint256[24] _proof, uint256[3] _pubSignals) view returns(bool)
func (_Abi *AbiSession) VerifyProof(_proof [24]*big.Int, _pubSignals [3]*big.Int) (bool, error) {
	return _Abi.Contract.VerifyProof(&_Abi.CallOpts, _proof, _pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0x1d5803fe.
//
// Solidity: function verifyProof(uint256[24] _proof, uint256[3] _pubSignals) view returns(bool)
func (_Abi *AbiCallerSession) VerifyProof(_proof [24]*big.Int, _pubSignals [3]*big.Int) (bool, error) {
	return _Abi.Contract.VerifyProof(&_Abi.CallOpts, _proof, _pubSignals)
}

// WhitelistedRequester is a free data retrieval call binding the contract method 0x3012baf7.
//
// Solidity: function whitelistedRequester() view returns(address)
func (_Abi *AbiCaller) WhitelistedRequester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "whitelistedRequester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistedRequester is a free data retrieval call binding the contract method 0x3012baf7.
//
// Solidity: function whitelistedRequester() view returns(address)
func (_Abi *AbiSession) WhitelistedRequester() (common.Address, error) {
	return _Abi.Contract.WhitelistedRequester(&_Abi.CallOpts)
}

// WhitelistedRequester is a free data retrieval call binding the contract method 0x3012baf7.
//
// Solidity: function whitelistedRequester() view returns(address)
func (_Abi *AbiCallerSession) WhitelistedRequester() (common.Address, error) {
	return _Abi.Contract.WhitelistedRequester(&_Abi.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Abi *AbiTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Abi *AbiSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Abi.Contract.GrantRole(&_Abi.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Abi *AbiTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Abi.Contract.GrantRole(&_Abi.TransactOpts, role, account)
}

// RefundUnfulfilledRequest is a paid mutator transaction binding the contract method 0xca6f30a4.
//
// Solidity: function refundUnfulfilledRequest(uint256 _requestId) returns()
func (_Abi *AbiTransactor) RefundUnfulfilledRequest(opts *bind.TransactOpts, _requestId *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "refundUnfulfilledRequest", _requestId)
}

// RefundUnfulfilledRequest is a paid mutator transaction binding the contract method 0xca6f30a4.
//
// Solidity: function refundUnfulfilledRequest(uint256 _requestId) returns()
func (_Abi *AbiSession) RefundUnfulfilledRequest(_requestId *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.RefundUnfulfilledRequest(&_Abi.TransactOpts, _requestId)
}

// RefundUnfulfilledRequest is a paid mutator transaction binding the contract method 0xca6f30a4.
//
// Solidity: function refundUnfulfilledRequest(uint256 _requestId) returns()
func (_Abi *AbiTransactorSession) RefundUnfulfilledRequest(_requestId *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.RefundUnfulfilledRequest(&_Abi.TransactOpts, _requestId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Abi *AbiTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Abi *AbiSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Abi.Contract.RenounceRole(&_Abi.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Abi *AbiTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Abi.Contract.RenounceRole(&_Abi.TransactOpts, role, callerConfirmation)
}

// RequestRng is a paid mutator transaction binding the contract method 0x30a3396d.
//
// Solidity: function requestRng(address callbackAddress, uint32 callbackGasLimit) payable returns(uint256)
func (_Abi *AbiTransactor) RequestRng(opts *bind.TransactOpts, callbackAddress common.Address, callbackGasLimit uint32) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "requestRng", callbackAddress, callbackGasLimit)
}

// RequestRng is a paid mutator transaction binding the contract method 0x30a3396d.
//
// Solidity: function requestRng(address callbackAddress, uint32 callbackGasLimit) payable returns(uint256)
func (_Abi *AbiSession) RequestRng(callbackAddress common.Address, callbackGasLimit uint32) (*types.Transaction, error) {
	return _Abi.Contract.RequestRng(&_Abi.TransactOpts, callbackAddress, callbackGasLimit)
}

// RequestRng is a paid mutator transaction binding the contract method 0x30a3396d.
//
// Solidity: function requestRng(address callbackAddress, uint32 callbackGasLimit) payable returns(uint256)
func (_Abi *AbiTransactorSession) RequestRng(callbackAddress common.Address, callbackGasLimit uint32) (*types.Transaction, error) {
	return _Abi.Contract.RequestRng(&_Abi.TransactOpts, callbackAddress, callbackGasLimit)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Abi *AbiTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Abi *AbiSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Abi.Contract.RevokeRole(&_Abi.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Abi *AbiTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Abi.Contract.RevokeRole(&_Abi.TransactOpts, role, account)
}

// SetCommitment is a paid mutator transaction binding the contract method 0xbe9cd052.
//
// Solidity: function setCommitment(uint256 _commitment) returns()
func (_Abi *AbiTransactor) SetCommitment(opts *bind.TransactOpts, _commitment *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setCommitment", _commitment)
}

// SetCommitment is a paid mutator transaction binding the contract method 0xbe9cd052.
//
// Solidity: function setCommitment(uint256 _commitment) returns()
func (_Abi *AbiSession) SetCommitment(_commitment *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetCommitment(&_Abi.TransactOpts, _commitment)
}

// SetCommitment is a paid mutator transaction binding the contract method 0xbe9cd052.
//
// Solidity: function setCommitment(uint256 _commitment) returns()
func (_Abi *AbiTransactorSession) SetCommitment(_commitment *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetCommitment(&_Abi.TransactOpts, _commitment)
}

// SetContractFeePercentage is a paid mutator transaction binding the contract method 0xa6504f8c.
//
// Solidity: function setContractFeePercentage(uint256 _newPercentage) returns()
func (_Abi *AbiTransactor) SetContractFeePercentage(opts *bind.TransactOpts, _newPercentage *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setContractFeePercentage", _newPercentage)
}

// SetContractFeePercentage is a paid mutator transaction binding the contract method 0xa6504f8c.
//
// Solidity: function setContractFeePercentage(uint256 _newPercentage) returns()
func (_Abi *AbiSession) SetContractFeePercentage(_newPercentage *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetContractFeePercentage(&_Abi.TransactOpts, _newPercentage)
}

// SetContractFeePercentage is a paid mutator transaction binding the contract method 0xa6504f8c.
//
// Solidity: function setContractFeePercentage(uint256 _newPercentage) returns()
func (_Abi *AbiTransactorSession) SetContractFeePercentage(_newPercentage *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetContractFeePercentage(&_Abi.TransactOpts, _newPercentage)
}

// SetRequestFee is a paid mutator transaction binding the contract method 0xffb9c43f.
//
// Solidity: function setRequestFee(uint256 _newFee) returns()
func (_Abi *AbiTransactor) SetRequestFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setRequestFee", _newFee)
}

// SetRequestFee is a paid mutator transaction binding the contract method 0xffb9c43f.
//
// Solidity: function setRequestFee(uint256 _newFee) returns()
func (_Abi *AbiSession) SetRequestFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetRequestFee(&_Abi.TransactOpts, _newFee)
}

// SetRequestFee is a paid mutator transaction binding the contract method 0xffb9c43f.
//
// Solidity: function setRequestFee(uint256 _newFee) returns()
func (_Abi *AbiTransactorSession) SetRequestFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetRequestFee(&_Abi.TransactOpts, _newFee)
}

// SetWhitelistedRequester is a paid mutator transaction binding the contract method 0x711824ba.
//
// Solidity: function setWhitelistedRequester(address _whitelistedRequester) returns()
func (_Abi *AbiTransactor) SetWhitelistedRequester(opts *bind.TransactOpts, _whitelistedRequester common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setWhitelistedRequester", _whitelistedRequester)
}

// SetWhitelistedRequester is a paid mutator transaction binding the contract method 0x711824ba.
//
// Solidity: function setWhitelistedRequester(address _whitelistedRequester) returns()
func (_Abi *AbiSession) SetWhitelistedRequester(_whitelistedRequester common.Address) (*types.Transaction, error) {
	return _Abi.Contract.SetWhitelistedRequester(&_Abi.TransactOpts, _whitelistedRequester)
}

// SetWhitelistedRequester is a paid mutator transaction binding the contract method 0x711824ba.
//
// Solidity: function setWhitelistedRequester(address _whitelistedRequester) returns()
func (_Abi *AbiTransactorSession) SetWhitelistedRequester(_whitelistedRequester common.Address) (*types.Transaction, error) {
	return _Abi.Contract.SetWhitelistedRequester(&_Abi.TransactOpts, _whitelistedRequester)
}

// SubmitEntropy is a paid mutator transaction binding the contract method 0x6804fe15.
//
// Solidity: function submitEntropy(uint256[24] _proof, uint256[3] _publicInputs, uint256 _requestId, address _rewardReceiver) returns()
func (_Abi *AbiTransactor) SubmitEntropy(opts *bind.TransactOpts, _proof [24]*big.Int, _publicInputs [3]*big.Int, _requestId *big.Int, _rewardReceiver common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "submitEntropy", _proof, _publicInputs, _requestId, _rewardReceiver)
}

// SubmitEntropy is a paid mutator transaction binding the contract method 0x6804fe15.
//
// Solidity: function submitEntropy(uint256[24] _proof, uint256[3] _publicInputs, uint256 _requestId, address _rewardReceiver) returns()
func (_Abi *AbiSession) SubmitEntropy(_proof [24]*big.Int, _publicInputs [3]*big.Int, _requestId *big.Int, _rewardReceiver common.Address) (*types.Transaction, error) {
	return _Abi.Contract.SubmitEntropy(&_Abi.TransactOpts, _proof, _publicInputs, _requestId, _rewardReceiver)
}

// SubmitEntropy is a paid mutator transaction binding the contract method 0x6804fe15.
//
// Solidity: function submitEntropy(uint256[24] _proof, uint256[3] _publicInputs, uint256 _requestId, address _rewardReceiver) returns()
func (_Abi *AbiTransactorSession) SubmitEntropy(_proof [24]*big.Int, _publicInputs [3]*big.Int, _requestId *big.Int, _rewardReceiver common.Address) (*types.Transaction, error) {
	return _Abi.Contract.SubmitEntropy(&_Abi.TransactOpts, _proof, _publicInputs, _requestId, _rewardReceiver)
}

// WithdrawContractFees is a paid mutator transaction binding the contract method 0xea3de4cb.
//
// Solidity: function withdrawContractFees() returns()
func (_Abi *AbiTransactor) WithdrawContractFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "withdrawContractFees")
}

// WithdrawContractFees is a paid mutator transaction binding the contract method 0xea3de4cb.
//
// Solidity: function withdrawContractFees() returns()
func (_Abi *AbiSession) WithdrawContractFees() (*types.Transaction, error) {
	return _Abi.Contract.WithdrawContractFees(&_Abi.TransactOpts)
}

// WithdrawContractFees is a paid mutator transaction binding the contract method 0xea3de4cb.
//
// Solidity: function withdrawContractFees() returns()
func (_Abi *AbiTransactorSession) WithdrawContractFees() (*types.Transaction, error) {
	return _Abi.Contract.WithdrawContractFees(&_Abi.TransactOpts)
}

// WithdrawRewardReceiverBalance is a paid mutator transaction binding the contract method 0x5e089f99.
//
// Solidity: function withdrawRewardReceiverBalance() returns()
func (_Abi *AbiTransactor) WithdrawRewardReceiverBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "withdrawRewardReceiverBalance")
}

// WithdrawRewardReceiverBalance is a paid mutator transaction binding the contract method 0x5e089f99.
//
// Solidity: function withdrawRewardReceiverBalance() returns()
func (_Abi *AbiSession) WithdrawRewardReceiverBalance() (*types.Transaction, error) {
	return _Abi.Contract.WithdrawRewardReceiverBalance(&_Abi.TransactOpts)
}

// WithdrawRewardReceiverBalance is a paid mutator transaction binding the contract method 0x5e089f99.
//
// Solidity: function withdrawRewardReceiverBalance() returns()
func (_Abi *AbiTransactorSession) WithdrawRewardReceiverBalance() (*types.Transaction, error) {
	return _Abi.Contract.WithdrawRewardReceiverBalance(&_Abi.TransactOpts)
}

// AbiContractFeePercentageUpdatedIterator is returned from FilterContractFeePercentageUpdated and is used to iterate over the raw logs and unpacked data for ContractFeePercentageUpdated events raised by the Abi contract.
type AbiContractFeePercentageUpdatedIterator struct {
	Event *AbiContractFeePercentageUpdated // Event containing the contract specifics and raw log

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
func (it *AbiContractFeePercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiContractFeePercentageUpdated)
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
		it.Event = new(AbiContractFeePercentageUpdated)
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
func (it *AbiContractFeePercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiContractFeePercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiContractFeePercentageUpdated represents a ContractFeePercentageUpdated event raised by the Abi contract.
type AbiContractFeePercentageUpdated struct {
	NewPercentage *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContractFeePercentageUpdated is a free log retrieval operation binding the contract event 0x88f49dcbaed0e4733413f55abc15dc393fc37709d1c423368bcf593e8ca61288.
//
// Solidity: event ContractFeePercentageUpdated(uint256 newPercentage)
func (_Abi *AbiFilterer) FilterContractFeePercentageUpdated(opts *bind.FilterOpts) (*AbiContractFeePercentageUpdatedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "ContractFeePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &AbiContractFeePercentageUpdatedIterator{contract: _Abi.contract, event: "ContractFeePercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchContractFeePercentageUpdated is a free log subscription operation binding the contract event 0x88f49dcbaed0e4733413f55abc15dc393fc37709d1c423368bcf593e8ca61288.
//
// Solidity: event ContractFeePercentageUpdated(uint256 newPercentage)
func (_Abi *AbiFilterer) WatchContractFeePercentageUpdated(opts *bind.WatchOpts, sink chan<- *AbiContractFeePercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "ContractFeePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiContractFeePercentageUpdated)
				if err := _Abi.contract.UnpackLog(event, "ContractFeePercentageUpdated", log); err != nil {
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
func (_Abi *AbiFilterer) ParseContractFeePercentageUpdated(log types.Log) (*AbiContractFeePercentageUpdated, error) {
	event := new(AbiContractFeePercentageUpdated)
	if err := _Abi.contract.UnpackLog(event, "ContractFeePercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiContractFeesWithdrawnIterator is returned from FilterContractFeesWithdrawn and is used to iterate over the raw logs and unpacked data for ContractFeesWithdrawn events raised by the Abi contract.
type AbiContractFeesWithdrawnIterator struct {
	Event *AbiContractFeesWithdrawn // Event containing the contract specifics and raw log

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
func (it *AbiContractFeesWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiContractFeesWithdrawn)
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
		it.Event = new(AbiContractFeesWithdrawn)
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
func (it *AbiContractFeesWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiContractFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiContractFeesWithdrawn represents a ContractFeesWithdrawn event raised by the Abi contract.
type AbiContractFeesWithdrawn struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterContractFeesWithdrawn is a free log retrieval operation binding the contract event 0x3d9264dd79c7dda789bd13ca13e81718ba78de6e6134e52be85fb6208347b013.
//
// Solidity: event ContractFeesWithdrawn(uint256 amount)
func (_Abi *AbiFilterer) FilterContractFeesWithdrawn(opts *bind.FilterOpts) (*AbiContractFeesWithdrawnIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "ContractFeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &AbiContractFeesWithdrawnIterator{contract: _Abi.contract, event: "ContractFeesWithdrawn", logs: logs, sub: sub}, nil
}

// WatchContractFeesWithdrawn is a free log subscription operation binding the contract event 0x3d9264dd79c7dda789bd13ca13e81718ba78de6e6134e52be85fb6208347b013.
//
// Solidity: event ContractFeesWithdrawn(uint256 amount)
func (_Abi *AbiFilterer) WatchContractFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *AbiContractFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "ContractFeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiContractFeesWithdrawn)
				if err := _Abi.contract.UnpackLog(event, "ContractFeesWithdrawn", log); err != nil {
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
func (_Abi *AbiFilterer) ParseContractFeesWithdrawn(log types.Log) (*AbiContractFeesWithdrawn, error) {
	event := new(AbiContractFeesWithdrawn)
	if err := _Abi.contract.UnpackLog(event, "ContractFeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiRequestFeeUpdatedIterator is returned from FilterRequestFeeUpdated and is used to iterate over the raw logs and unpacked data for RequestFeeUpdated events raised by the Abi contract.
type AbiRequestFeeUpdatedIterator struct {
	Event *AbiRequestFeeUpdated // Event containing the contract specifics and raw log

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
func (it *AbiRequestFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiRequestFeeUpdated)
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
		it.Event = new(AbiRequestFeeUpdated)
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
func (it *AbiRequestFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiRequestFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiRequestFeeUpdated represents a RequestFeeUpdated event raised by the Abi contract.
type AbiRequestFeeUpdated struct {
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRequestFeeUpdated is a free log retrieval operation binding the contract event 0x07695b29775442d5e4836f94223aa3460d93774d5cb9e03119815f418e2a61c4.
//
// Solidity: event RequestFeeUpdated(uint256 newFee)
func (_Abi *AbiFilterer) FilterRequestFeeUpdated(opts *bind.FilterOpts) (*AbiRequestFeeUpdatedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "RequestFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &AbiRequestFeeUpdatedIterator{contract: _Abi.contract, event: "RequestFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchRequestFeeUpdated is a free log subscription operation binding the contract event 0x07695b29775442d5e4836f94223aa3460d93774d5cb9e03119815f418e2a61c4.
//
// Solidity: event RequestFeeUpdated(uint256 newFee)
func (_Abi *AbiFilterer) WatchRequestFeeUpdated(opts *bind.WatchOpts, sink chan<- *AbiRequestFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "RequestFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiRequestFeeUpdated)
				if err := _Abi.contract.UnpackLog(event, "RequestFeeUpdated", log); err != nil {
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
func (_Abi *AbiFilterer) ParseRequestFeeUpdated(log types.Log) (*AbiRequestFeeUpdated, error) {
	event := new(AbiRequestFeeUpdated)
	if err := _Abi.contract.UnpackLog(event, "RequestFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiRequestFulfilledIterator is returned from FilterRequestFulfilled and is used to iterate over the raw logs and unpacked data for RequestFulfilled events raised by the Abi contract.
type AbiRequestFulfilledIterator struct {
	Event *AbiRequestFulfilled // Event containing the contract specifics and raw log

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
func (it *AbiRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiRequestFulfilled)
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
		it.Event = new(AbiRequestFulfilled)
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
func (it *AbiRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiRequestFulfilled represents a RequestFulfilled event raised by the Abi contract.
type AbiRequestFulfilled struct {
	RequestId       *big.Int
	CallbackSuccess bool
	RewardReceiver  common.Address
	Proof           [24]*big.Int
	PublicInputs    [3]*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRequestFulfilled is a free log retrieval operation binding the contract event 0xb1424970474328bbbc9fdcd8dfc4212683f9d3e2e2d89aee0d5c306a5399d3df.
//
// Solidity: event RequestFulfilled(uint256 indexed requestId, bool callbackSuccess, address rewardReceiver, uint256[24] proof, uint256[3] publicInputs)
func (_Abi *AbiFilterer) FilterRequestFulfilled(opts *bind.FilterOpts, requestId []*big.Int) (*AbiRequestFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "RequestFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &AbiRequestFulfilledIterator{contract: _Abi.contract, event: "RequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchRequestFulfilled is a free log subscription operation binding the contract event 0xb1424970474328bbbc9fdcd8dfc4212683f9d3e2e2d89aee0d5c306a5399d3df.
//
// Solidity: event RequestFulfilled(uint256 indexed requestId, bool callbackSuccess, address rewardReceiver, uint256[24] proof, uint256[3] publicInputs)
func (_Abi *AbiFilterer) WatchRequestFulfilled(opts *bind.WatchOpts, sink chan<- *AbiRequestFulfilled, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "RequestFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiRequestFulfilled)
				if err := _Abi.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
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

// ParseRequestFulfilled is a log parse operation binding the contract event 0xb1424970474328bbbc9fdcd8dfc4212683f9d3e2e2d89aee0d5c306a5399d3df.
//
// Solidity: event RequestFulfilled(uint256 indexed requestId, bool callbackSuccess, address rewardReceiver, uint256[24] proof, uint256[3] publicInputs)
func (_Abi *AbiFilterer) ParseRequestFulfilled(log types.Log) (*AbiRequestFulfilled, error) {
	event := new(AbiRequestFulfilled)
	if err := _Abi.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiRequestRefundedIterator is returned from FilterRequestRefunded and is used to iterate over the raw logs and unpacked data for RequestRefunded events raised by the Abi contract.
type AbiRequestRefundedIterator struct {
	Event *AbiRequestRefunded // Event containing the contract specifics and raw log

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
func (it *AbiRequestRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiRequestRefunded)
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
		it.Event = new(AbiRequestRefunded)
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
func (it *AbiRequestRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiRequestRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiRequestRefunded represents a RequestRefunded event raised by the Abi contract.
type AbiRequestRefunded struct {
	RequestId *big.Int
	Requester common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestRefunded is a free log retrieval operation binding the contract event 0x86e8ab52e95a11dc603ce8e0a102c802eec02a52c879f8f39378e6222ca8c82c.
//
// Solidity: event RequestRefunded(uint256 indexed requestId, address indexed requester, uint256 amount)
func (_Abi *AbiFilterer) FilterRequestRefunded(opts *bind.FilterOpts, requestId []*big.Int, requester []common.Address) (*AbiRequestRefundedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "RequestRefunded", requestIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &AbiRequestRefundedIterator{contract: _Abi.contract, event: "RequestRefunded", logs: logs, sub: sub}, nil
}

// WatchRequestRefunded is a free log subscription operation binding the contract event 0x86e8ab52e95a11dc603ce8e0a102c802eec02a52c879f8f39378e6222ca8c82c.
//
// Solidity: event RequestRefunded(uint256 indexed requestId, address indexed requester, uint256 amount)
func (_Abi *AbiFilterer) WatchRequestRefunded(opts *bind.WatchOpts, sink chan<- *AbiRequestRefunded, requestId []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "RequestRefunded", requestIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiRequestRefunded)
				if err := _Abi.contract.UnpackLog(event, "RequestRefunded", log); err != nil {
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

// ParseRequestRefunded is a log parse operation binding the contract event 0x86e8ab52e95a11dc603ce8e0a102c802eec02a52c879f8f39378e6222ca8c82c.
//
// Solidity: event RequestRefunded(uint256 indexed requestId, address indexed requester, uint256 amount)
func (_Abi *AbiFilterer) ParseRequestRefunded(log types.Log) (*AbiRequestRefunded, error) {
	event := new(AbiRequestRefunded)
	if err := _Abi.contract.UnpackLog(event, "RequestRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiRewardReceiverBalanceWithdrawnIterator is returned from FilterRewardReceiverBalanceWithdrawn and is used to iterate over the raw logs and unpacked data for RewardReceiverBalanceWithdrawn events raised by the Abi contract.
type AbiRewardReceiverBalanceWithdrawnIterator struct {
	Event *AbiRewardReceiverBalanceWithdrawn // Event containing the contract specifics and raw log

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
func (it *AbiRewardReceiverBalanceWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiRewardReceiverBalanceWithdrawn)
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
		it.Event = new(AbiRewardReceiverBalanceWithdrawn)
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
func (it *AbiRewardReceiverBalanceWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiRewardReceiverBalanceWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiRewardReceiverBalanceWithdrawn represents a RewardReceiverBalanceWithdrawn event raised by the Abi contract.
type AbiRewardReceiverBalanceWithdrawn struct {
	RewardReceiver common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardReceiverBalanceWithdrawn is a free log retrieval operation binding the contract event 0x0b40a73848f6d60e086f98c5ba41e8d98e0f0a47db8ae16597c266e3a5c40cba.
//
// Solidity: event RewardReceiverBalanceWithdrawn(address indexed rewardReceiver, uint256 amount)
func (_Abi *AbiFilterer) FilterRewardReceiverBalanceWithdrawn(opts *bind.FilterOpts, rewardReceiver []common.Address) (*AbiRewardReceiverBalanceWithdrawnIterator, error) {

	var rewardReceiverRule []interface{}
	for _, rewardReceiverItem := range rewardReceiver {
		rewardReceiverRule = append(rewardReceiverRule, rewardReceiverItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "RewardReceiverBalanceWithdrawn", rewardReceiverRule)
	if err != nil {
		return nil, err
	}
	return &AbiRewardReceiverBalanceWithdrawnIterator{contract: _Abi.contract, event: "RewardReceiverBalanceWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRewardReceiverBalanceWithdrawn is a free log subscription operation binding the contract event 0x0b40a73848f6d60e086f98c5ba41e8d98e0f0a47db8ae16597c266e3a5c40cba.
//
// Solidity: event RewardReceiverBalanceWithdrawn(address indexed rewardReceiver, uint256 amount)
func (_Abi *AbiFilterer) WatchRewardReceiverBalanceWithdrawn(opts *bind.WatchOpts, sink chan<- *AbiRewardReceiverBalanceWithdrawn, rewardReceiver []common.Address) (event.Subscription, error) {

	var rewardReceiverRule []interface{}
	for _, rewardReceiverItem := range rewardReceiver {
		rewardReceiverRule = append(rewardReceiverRule, rewardReceiverItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "RewardReceiverBalanceWithdrawn", rewardReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiRewardReceiverBalanceWithdrawn)
				if err := _Abi.contract.UnpackLog(event, "RewardReceiverBalanceWithdrawn", log); err != nil {
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
func (_Abi *AbiFilterer) ParseRewardReceiverBalanceWithdrawn(log types.Log) (*AbiRewardReceiverBalanceWithdrawn, error) {
	event := new(AbiRewardReceiverBalanceWithdrawn)
	if err := _Abi.contract.UnpackLog(event, "RewardReceiverBalanceWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiRngRequestedIterator is returned from FilterRngRequested and is used to iterate over the raw logs and unpacked data for RngRequested events raised by the Abi contract.
type AbiRngRequestedIterator struct {
	Event *AbiRngRequested // Event containing the contract specifics and raw log

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
func (it *AbiRngRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiRngRequested)
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
		it.Event = new(AbiRngRequested)
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
func (it *AbiRngRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiRngRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiRngRequested represents a RngRequested event raised by the Abi contract.
type AbiRngRequested struct {
	RequestId *big.Int
	BlockHash [32]byte
	Requester common.Address
	FeePaid   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRngRequested is a free log retrieval operation binding the contract event 0xebd7660295e4ce331b3cfbc0706171ac8de0c07c415924affdcf42ddc11774dc.
//
// Solidity: event RngRequested(uint256 requestId, bytes32 blockHash, address requester, uint256 feePaid)
func (_Abi *AbiFilterer) FilterRngRequested(opts *bind.FilterOpts) (*AbiRngRequestedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "RngRequested")
	if err != nil {
		return nil, err
	}
	return &AbiRngRequestedIterator{contract: _Abi.contract, event: "RngRequested", logs: logs, sub: sub}, nil
}

// WatchRngRequested is a free log subscription operation binding the contract event 0xebd7660295e4ce331b3cfbc0706171ac8de0c07c415924affdcf42ddc11774dc.
//
// Solidity: event RngRequested(uint256 requestId, bytes32 blockHash, address requester, uint256 feePaid)
func (_Abi *AbiFilterer) WatchRngRequested(opts *bind.WatchOpts, sink chan<- *AbiRngRequested) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "RngRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiRngRequested)
				if err := _Abi.contract.UnpackLog(event, "RngRequested", log); err != nil {
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
func (_Abi *AbiFilterer) ParseRngRequested(log types.Log) (*AbiRngRequested, error) {
	event := new(AbiRngRequested)
	if err := _Abi.contract.UnpackLog(event, "RngRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Abi contract.
type AbiRoleAdminChangedIterator struct {
	Event *AbiRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AbiRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiRoleAdminChanged)
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
		it.Event = new(AbiRoleAdminChanged)
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
func (it *AbiRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiRoleAdminChanged represents a RoleAdminChanged event raised by the Abi contract.
type AbiRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Abi *AbiFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AbiRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Abi.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AbiRoleAdminChangedIterator{contract: _Abi.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Abi *AbiFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AbiRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Abi.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiRoleAdminChanged)
				if err := _Abi.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Abi *AbiFilterer) ParseRoleAdminChanged(log types.Log) (*AbiRoleAdminChanged, error) {
	event := new(AbiRoleAdminChanged)
	if err := _Abi.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Abi contract.
type AbiRoleGrantedIterator struct {
	Event *AbiRoleGranted // Event containing the contract specifics and raw log

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
func (it *AbiRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiRoleGranted)
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
		it.Event = new(AbiRoleGranted)
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
func (it *AbiRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiRoleGranted represents a RoleGranted event raised by the Abi contract.
type AbiRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Abi *AbiFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AbiRoleGrantedIterator, error) {

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

	logs, sub, err := _Abi.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AbiRoleGrantedIterator{contract: _Abi.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Abi *AbiFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AbiRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Abi.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiRoleGranted)
				if err := _Abi.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Abi *AbiFilterer) ParseRoleGranted(log types.Log) (*AbiRoleGranted, error) {
	event := new(AbiRoleGranted)
	if err := _Abi.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Abi contract.
type AbiRoleRevokedIterator struct {
	Event *AbiRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AbiRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiRoleRevoked)
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
		it.Event = new(AbiRoleRevoked)
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
func (it *AbiRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiRoleRevoked represents a RoleRevoked event raised by the Abi contract.
type AbiRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Abi *AbiFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AbiRoleRevokedIterator, error) {

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

	logs, sub, err := _Abi.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AbiRoleRevokedIterator{contract: _Abi.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Abi *AbiFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AbiRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Abi.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiRoleRevoked)
				if err := _Abi.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Abi *AbiFilterer) ParseRoleRevoked(log types.Log) (*AbiRoleRevoked, error) {
	event := new(AbiRoleRevoked)
	if err := _Abi.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
