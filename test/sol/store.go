// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sol

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

// SolMetaData contains all meta data concerning the Sol contract.
var SolMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"getProof\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_proof\",\"type\":\"string\"}],\"name\":\"setProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061030f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806315336f801461003b5780635c5d625e146100f6575b600080fd5b6100f46004803603602081101561005157600080fd5b810190808035906020019064010000000081111561006e57600080fd5b82018360208201111561008057600080fd5b803590602001918460018302840111640100000000831117156100a257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610179565b005b6100fe610193565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561013e578082015181840152602081019050610123565b50505050905090810190601f16801561016b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b806000908051906020019061018f929190610235565b5050565b606060008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561022b5780601f106102005761010080835404028352916020019161022b565b820191906000526020600020905b81548152906001019060200180831161020e57829003601f168201915b5050505050905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061027657805160ff19168380011785556102a4565b828001600101855582156102a4579182015b828111156102a3578251825591602001919060010190610288565b5b5090506102b191906102b5565b5090565b6102d791905b808211156102d35760008160009055506001016102bb565b5090565b9056fea265627a7a72315820120470cf11d076eebbf0d105cc884ccc350a2bbe8ad73dc6eb7f4d6d1580c76664736f6c63430005110032",
}

// SolABI is the input ABI used to generate the binding from.
// Deprecated: Use SolMetaData.ABI instead.
var SolABI = SolMetaData.ABI

// SolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolMetaData.Bin instead.
var SolBin = SolMetaData.Bin

// DeploySol deploys a new Ethereum contract, binding an instance of Sol to it.
func DeploySol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sol, error) {
	parsed, err := SolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sol{SolCaller: SolCaller{contract: contract}, SolTransactor: SolTransactor{contract: contract}, SolFilterer: SolFilterer{contract: contract}}, nil
}

// Sol is an auto generated Go binding around an Ethereum contract.
type Sol struct {
	SolCaller     // Read-only binding to the contract
	SolTransactor // Write-only binding to the contract
	SolFilterer   // Log filterer for contract events
}

// SolCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolSession struct {
	Contract     *Sol              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolCallerSession struct {
	Contract *SolCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolTransactorSession struct {
	Contract     *SolTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolRaw struct {
	Contract *Sol // Generic contract binding to access the raw methods on
}

// SolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolCallerRaw struct {
	Contract *SolCaller // Generic read-only contract binding to access the raw methods on
}

// SolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolTransactorRaw struct {
	Contract *SolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSol creates a new instance of Sol, bound to a specific deployed contract.
func NewSol(address common.Address, backend bind.ContractBackend) (*Sol, error) {
	contract, err := bindSol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sol{SolCaller: SolCaller{contract: contract}, SolTransactor: SolTransactor{contract: contract}, SolFilterer: SolFilterer{contract: contract}}, nil
}

// NewSolCaller creates a new read-only instance of Sol, bound to a specific deployed contract.
func NewSolCaller(address common.Address, caller bind.ContractCaller) (*SolCaller, error) {
	contract, err := bindSol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolCaller{contract: contract}, nil
}

// NewSolTransactor creates a new write-only instance of Sol, bound to a specific deployed contract.
func NewSolTransactor(address common.Address, transactor bind.ContractTransactor) (*SolTransactor, error) {
	contract, err := bindSol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolTransactor{contract: contract}, nil
}

// NewSolFilterer creates a new log filterer instance of Sol, bound to a specific deployed contract.
func NewSolFilterer(address common.Address, filterer bind.ContractFilterer) (*SolFilterer, error) {
	contract, err := bindSol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolFilterer{contract: contract}, nil
}

// bindSol binds a generic wrapper to an already deployed contract.
func bindSol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sol *SolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sol.Contract.SolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sol *SolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sol.Contract.SolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sol *SolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sol.Contract.SolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sol *SolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sol *SolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sol *SolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sol.Contract.contract.Transact(opts, method, params...)
}

// GetProof is a free data retrieval call binding the contract method 0x5c5d625e.
//
// Solidity: function getProof() view returns(string)
func (_Sol *SolCaller) GetProof(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Sol.contract.Call(opts, &out, "getProof")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x5c5d625e.
//
// Solidity: function getProof() view returns(string)
func (_Sol *SolSession) GetProof() (string, error) {
	return _Sol.Contract.GetProof(&_Sol.CallOpts)
}

// GetProof is a free data retrieval call binding the contract method 0x5c5d625e.
//
// Solidity: function getProof() view returns(string)
func (_Sol *SolCallerSession) GetProof() (string, error) {
	return _Sol.Contract.GetProof(&_Sol.CallOpts)
}

// SetProof is a paid mutator transaction binding the contract method 0x15336f80.
//
// Solidity: function setProof(string _proof) returns()
func (_Sol *SolTransactor) SetProof(opts *bind.TransactOpts, _proof string) (*types.Transaction, error) {
	return _Sol.contract.Transact(opts, "setProof", _proof)
}

// SetProof is a paid mutator transaction binding the contract method 0x15336f80.
//
// Solidity: function setProof(string _proof) returns()
func (_Sol *SolSession) SetProof(_proof string) (*types.Transaction, error) {
	return _Sol.Contract.SetProof(&_Sol.TransactOpts, _proof)
}

// SetProof is a paid mutator transaction binding the contract method 0x15336f80.
//
// Solidity: function setProof(string _proof) returns()
func (_Sol *SolTransactorSession) SetProof(_proof string) (*types.Transaction, error) {
	return _Sol.Contract.SetProof(&_Sol.TransactOpts, _proof)
}
