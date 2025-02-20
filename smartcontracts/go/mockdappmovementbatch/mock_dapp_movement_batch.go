// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockdappmovementbatch

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

// MockdappmovementbatchMetaData contains all meta data concerning the Mockdappmovementbatch contract.
var MockdappmovementbatchMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CustomError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"bytesToBytes32Array\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deviceTick\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"errorType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bitmap\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"isBitSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prover\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_taskIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"process\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_errorType\",\"type\":\"uint8\"}],\"name\":\"setErrorType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MockdappmovementbatchABI is the input ABI used to generate the binding from.
// Deprecated: Use MockdappmovementbatchMetaData.ABI instead.
var MockdappmovementbatchABI = MockdappmovementbatchMetaData.ABI

// Mockdappmovementbatch is an auto generated Go binding around an Ethereum contract.
type Mockdappmovementbatch struct {
	MockdappmovementbatchCaller     // Read-only binding to the contract
	MockdappmovementbatchTransactor // Write-only binding to the contract
	MockdappmovementbatchFilterer   // Log filterer for contract events
}

// MockdappmovementbatchCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockdappmovementbatchCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockdappmovementbatchTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockdappmovementbatchTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockdappmovementbatchFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockdappmovementbatchFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockdappmovementbatchSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockdappmovementbatchSession struct {
	Contract     *Mockdappmovementbatch // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MockdappmovementbatchCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockdappmovementbatchCallerSession struct {
	Contract *MockdappmovementbatchCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// MockdappmovementbatchTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockdappmovementbatchTransactorSession struct {
	Contract     *MockdappmovementbatchTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// MockdappmovementbatchRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockdappmovementbatchRaw struct {
	Contract *Mockdappmovementbatch // Generic contract binding to access the raw methods on
}

// MockdappmovementbatchCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockdappmovementbatchCallerRaw struct {
	Contract *MockdappmovementbatchCaller // Generic read-only contract binding to access the raw methods on
}

// MockdappmovementbatchTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockdappmovementbatchTransactorRaw struct {
	Contract *MockdappmovementbatchTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockdappmovementbatch creates a new instance of Mockdappmovementbatch, bound to a specific deployed contract.
func NewMockdappmovementbatch(address common.Address, backend bind.ContractBackend) (*Mockdappmovementbatch, error) {
	contract, err := bindMockdappmovementbatch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mockdappmovementbatch{MockdappmovementbatchCaller: MockdappmovementbatchCaller{contract: contract}, MockdappmovementbatchTransactor: MockdappmovementbatchTransactor{contract: contract}, MockdappmovementbatchFilterer: MockdappmovementbatchFilterer{contract: contract}}, nil
}

// NewMockdappmovementbatchCaller creates a new read-only instance of Mockdappmovementbatch, bound to a specific deployed contract.
func NewMockdappmovementbatchCaller(address common.Address, caller bind.ContractCaller) (*MockdappmovementbatchCaller, error) {
	contract, err := bindMockdappmovementbatch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockdappmovementbatchCaller{contract: contract}, nil
}

// NewMockdappmovementbatchTransactor creates a new write-only instance of Mockdappmovementbatch, bound to a specific deployed contract.
func NewMockdappmovementbatchTransactor(address common.Address, transactor bind.ContractTransactor) (*MockdappmovementbatchTransactor, error) {
	contract, err := bindMockdappmovementbatch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockdappmovementbatchTransactor{contract: contract}, nil
}

// NewMockdappmovementbatchFilterer creates a new log filterer instance of Mockdappmovementbatch, bound to a specific deployed contract.
func NewMockdappmovementbatchFilterer(address common.Address, filterer bind.ContractFilterer) (*MockdappmovementbatchFilterer, error) {
	contract, err := bindMockdappmovementbatch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockdappmovementbatchFilterer{contract: contract}, nil
}

// bindMockdappmovementbatch binds a generic wrapper to an already deployed contract.
func bindMockdappmovementbatch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockdappmovementbatchMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mockdappmovementbatch *MockdappmovementbatchRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mockdappmovementbatch.Contract.MockdappmovementbatchCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mockdappmovementbatch *MockdappmovementbatchRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mockdappmovementbatch.Contract.MockdappmovementbatchTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mockdappmovementbatch *MockdappmovementbatchRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mockdappmovementbatch.Contract.MockdappmovementbatchTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mockdappmovementbatch *MockdappmovementbatchCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mockdappmovementbatch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mockdappmovementbatch *MockdappmovementbatchTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mockdappmovementbatch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mockdappmovementbatch *MockdappmovementbatchTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mockdappmovementbatch.Contract.contract.Transact(opts, method, params...)
}

// BytesToBytes32Array is a free data retrieval call binding the contract method 0x4debc578.
//
// Solidity: function bytesToBytes32Array(bytes data) pure returns(bytes32[])
func (_Mockdappmovementbatch *MockdappmovementbatchCaller) BytesToBytes32Array(opts *bind.CallOpts, data []byte) ([][32]byte, error) {
	var out []interface{}
	err := _Mockdappmovementbatch.contract.Call(opts, &out, "bytesToBytes32Array", data)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// BytesToBytes32Array is a free data retrieval call binding the contract method 0x4debc578.
//
// Solidity: function bytesToBytes32Array(bytes data) pure returns(bytes32[])
func (_Mockdappmovementbatch *MockdappmovementbatchSession) BytesToBytes32Array(data []byte) ([][32]byte, error) {
	return _Mockdappmovementbatch.Contract.BytesToBytes32Array(&_Mockdappmovementbatch.CallOpts, data)
}

// BytesToBytes32Array is a free data retrieval call binding the contract method 0x4debc578.
//
// Solidity: function bytesToBytes32Array(bytes data) pure returns(bytes32[])
func (_Mockdappmovementbatch *MockdappmovementbatchCallerSession) BytesToBytes32Array(data []byte) ([][32]byte, error) {
	return _Mockdappmovementbatch.Contract.BytesToBytes32Array(&_Mockdappmovementbatch.CallOpts, data)
}

// DeviceTick is a free data retrieval call binding the contract method 0x0e3a9e39.
//
// Solidity: function deviceTick(address ) view returns(uint64)
func (_Mockdappmovementbatch *MockdappmovementbatchCaller) DeviceTick(opts *bind.CallOpts, arg0 common.Address) (uint64, error) {
	var out []interface{}
	err := _Mockdappmovementbatch.contract.Call(opts, &out, "deviceTick", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DeviceTick is a free data retrieval call binding the contract method 0x0e3a9e39.
//
// Solidity: function deviceTick(address ) view returns(uint64)
func (_Mockdappmovementbatch *MockdappmovementbatchSession) DeviceTick(arg0 common.Address) (uint64, error) {
	return _Mockdappmovementbatch.Contract.DeviceTick(&_Mockdappmovementbatch.CallOpts, arg0)
}

// DeviceTick is a free data retrieval call binding the contract method 0x0e3a9e39.
//
// Solidity: function deviceTick(address ) view returns(uint64)
func (_Mockdappmovementbatch *MockdappmovementbatchCallerSession) DeviceTick(arg0 common.Address) (uint64, error) {
	return _Mockdappmovementbatch.Contract.DeviceTick(&_Mockdappmovementbatch.CallOpts, arg0)
}

// ErrorType is a free data retrieval call binding the contract method 0x5a4c2ceb.
//
// Solidity: function errorType() view returns(uint8)
func (_Mockdappmovementbatch *MockdappmovementbatchCaller) ErrorType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Mockdappmovementbatch.contract.Call(opts, &out, "errorType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ErrorType is a free data retrieval call binding the contract method 0x5a4c2ceb.
//
// Solidity: function errorType() view returns(uint8)
func (_Mockdappmovementbatch *MockdappmovementbatchSession) ErrorType() (uint8, error) {
	return _Mockdappmovementbatch.Contract.ErrorType(&_Mockdappmovementbatch.CallOpts)
}

// ErrorType is a free data retrieval call binding the contract method 0x5a4c2ceb.
//
// Solidity: function errorType() view returns(uint8)
func (_Mockdappmovementbatch *MockdappmovementbatchCallerSession) ErrorType() (uint8, error) {
	return _Mockdappmovementbatch.Contract.ErrorType(&_Mockdappmovementbatch.CallOpts)
}

// IsBitSet is a free data retrieval call binding the contract method 0xbaba9f0b.
//
// Solidity: function isBitSet(bytes32 bitmap, uint256 n) pure returns(bool)
func (_Mockdappmovementbatch *MockdappmovementbatchCaller) IsBitSet(opts *bind.CallOpts, bitmap [32]byte, n *big.Int) (bool, error) {
	var out []interface{}
	err := _Mockdappmovementbatch.contract.Call(opts, &out, "isBitSet", bitmap, n)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBitSet is a free data retrieval call binding the contract method 0xbaba9f0b.
//
// Solidity: function isBitSet(bytes32 bitmap, uint256 n) pure returns(bool)
func (_Mockdappmovementbatch *MockdappmovementbatchSession) IsBitSet(bitmap [32]byte, n *big.Int) (bool, error) {
	return _Mockdappmovementbatch.Contract.IsBitSet(&_Mockdappmovementbatch.CallOpts, bitmap, n)
}

// IsBitSet is a free data retrieval call binding the contract method 0xbaba9f0b.
//
// Solidity: function isBitSet(bytes32 bitmap, uint256 n) pure returns(bool)
func (_Mockdappmovementbatch *MockdappmovementbatchCallerSession) IsBitSet(bitmap [32]byte, n *big.Int) (bool, error) {
	return _Mockdappmovementbatch.Contract.IsBitSet(&_Mockdappmovementbatch.CallOpts, bitmap, n)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Mockdappmovementbatch *MockdappmovementbatchCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mockdappmovementbatch.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Mockdappmovementbatch *MockdappmovementbatchSession) Verifier() (common.Address, error) {
	return _Mockdappmovementbatch.Contract.Verifier(&_Mockdappmovementbatch.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Mockdappmovementbatch *MockdappmovementbatchCallerSession) Verifier() (common.Address, error) {
	return _Mockdappmovementbatch.Contract.Verifier(&_Mockdappmovementbatch.CallOpts)
}

// Process is a paid mutator transaction binding the contract method 0x108d9184.
//
// Solidity: function process(address _prover, uint256 _projectId, bytes32[] _taskIds, bytes _data) returns()
func (_Mockdappmovementbatch *MockdappmovementbatchTransactor) Process(opts *bind.TransactOpts, _prover common.Address, _projectId *big.Int, _taskIds [][32]byte, _data []byte) (*types.Transaction, error) {
	return _Mockdappmovementbatch.contract.Transact(opts, "process", _prover, _projectId, _taskIds, _data)
}

// Process is a paid mutator transaction binding the contract method 0x108d9184.
//
// Solidity: function process(address _prover, uint256 _projectId, bytes32[] _taskIds, bytes _data) returns()
func (_Mockdappmovementbatch *MockdappmovementbatchSession) Process(_prover common.Address, _projectId *big.Int, _taskIds [][32]byte, _data []byte) (*types.Transaction, error) {
	return _Mockdappmovementbatch.Contract.Process(&_Mockdappmovementbatch.TransactOpts, _prover, _projectId, _taskIds, _data)
}

// Process is a paid mutator transaction binding the contract method 0x108d9184.
//
// Solidity: function process(address _prover, uint256 _projectId, bytes32[] _taskIds, bytes _data) returns()
func (_Mockdappmovementbatch *MockdappmovementbatchTransactorSession) Process(_prover common.Address, _projectId *big.Int, _taskIds [][32]byte, _data []byte) (*types.Transaction, error) {
	return _Mockdappmovementbatch.Contract.Process(&_Mockdappmovementbatch.TransactOpts, _prover, _projectId, _taskIds, _data)
}

// SetErrorType is a paid mutator transaction binding the contract method 0xdafe3e3a.
//
// Solidity: function setErrorType(uint8 _errorType) returns()
func (_Mockdappmovementbatch *MockdappmovementbatchTransactor) SetErrorType(opts *bind.TransactOpts, _errorType uint8) (*types.Transaction, error) {
	return _Mockdappmovementbatch.contract.Transact(opts, "setErrorType", _errorType)
}

// SetErrorType is a paid mutator transaction binding the contract method 0xdafe3e3a.
//
// Solidity: function setErrorType(uint8 _errorType) returns()
func (_Mockdappmovementbatch *MockdappmovementbatchSession) SetErrorType(_errorType uint8) (*types.Transaction, error) {
	return _Mockdappmovementbatch.Contract.SetErrorType(&_Mockdappmovementbatch.TransactOpts, _errorType)
}

// SetErrorType is a paid mutator transaction binding the contract method 0xdafe3e3a.
//
// Solidity: function setErrorType(uint8 _errorType) returns()
func (_Mockdappmovementbatch *MockdappmovementbatchTransactorSession) SetErrorType(_errorType uint8) (*types.Transaction, error) {
	return _Mockdappmovementbatch.Contract.SetErrorType(&_Mockdappmovementbatch.TransactOpts, _errorType)
}
