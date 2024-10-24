// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package prover

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

// ProverMetaData contains all meta data concerning the Prover contract.
var ProverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"BeneficiarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"ProverPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"ProverResumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"ratio\",\"type\":\"uint16\"}],\"name\":\"RebateRatioSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"typ\",\"type\":\"uint256\"}],\"name\":\"VMTypeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"typ\",\"type\":\"uint256\"}],\"name\":\"VMTypeDeleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_type\",\"type\":\"uint256\"}],\"name\":\"addVMType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prover\",\"type\":\"address\"}],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"changeBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_type\",\"type\":\"uint256\"}],\"name\":\"delVMType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prover\",\"type\":\"address\"}],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prover\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_type\",\"type\":\"uint256\"}],\"name\":\"isVMTypeSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prover\",\"type\":\"address\"}],\"name\":\"rebateRatio\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_ratio\",\"type\":\"uint16\"}],\"name\":\"setRebateRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProverABI is the input ABI used to generate the binding from.
// Deprecated: Use ProverMetaData.ABI instead.
var ProverABI = ProverMetaData.ABI

// Prover is an auto generated Go binding around an Ethereum contract.
type Prover struct {
	ProverCaller     // Read-only binding to the contract
	ProverTransactor // Write-only binding to the contract
	ProverFilterer   // Log filterer for contract events
}

// ProverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProverSession struct {
	Contract     *Prover           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProverCallerSession struct {
	Contract *ProverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProverTransactorSession struct {
	Contract     *ProverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProverRaw struct {
	Contract *Prover // Generic contract binding to access the raw methods on
}

// ProverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProverCallerRaw struct {
	Contract *ProverCaller // Generic read-only contract binding to access the raw methods on
}

// ProverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProverTransactorRaw struct {
	Contract *ProverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProver creates a new instance of Prover, bound to a specific deployed contract.
func NewProver(address common.Address, backend bind.ContractBackend) (*Prover, error) {
	contract, err := bindProver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Prover{ProverCaller: ProverCaller{contract: contract}, ProverTransactor: ProverTransactor{contract: contract}, ProverFilterer: ProverFilterer{contract: contract}}, nil
}

// NewProverCaller creates a new read-only instance of Prover, bound to a specific deployed contract.
func NewProverCaller(address common.Address, caller bind.ContractCaller) (*ProverCaller, error) {
	contract, err := bindProver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProverCaller{contract: contract}, nil
}

// NewProverTransactor creates a new write-only instance of Prover, bound to a specific deployed contract.
func NewProverTransactor(address common.Address, transactor bind.ContractTransactor) (*ProverTransactor, error) {
	contract, err := bindProver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProverTransactor{contract: contract}, nil
}

// NewProverFilterer creates a new log filterer instance of Prover, bound to a specific deployed contract.
func NewProverFilterer(address common.Address, filterer bind.ContractFilterer) (*ProverFilterer, error) {
	contract, err := bindProver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProverFilterer{contract: contract}, nil
}

// bindProver binds a generic wrapper to an already deployed contract.
func bindProver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Prover *ProverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Prover.Contract.ProverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Prover *ProverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prover.Contract.ProverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Prover *ProverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Prover.Contract.ProverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Prover *ProverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Prover.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Prover *ProverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prover.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Prover *ProverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Prover.Contract.contract.Transact(opts, method, params...)
}

// Beneficiary is a free data retrieval call binding the contract method 0x81008568.
//
// Solidity: function beneficiary(address _prover) view returns(address)
func (_Prover *ProverCaller) Beneficiary(opts *bind.CallOpts, _prover common.Address) (common.Address, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "beneficiary", _prover)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x81008568.
//
// Solidity: function beneficiary(address _prover) view returns(address)
func (_Prover *ProverSession) Beneficiary(_prover common.Address) (common.Address, error) {
	return _Prover.Contract.Beneficiary(&_Prover.CallOpts, _prover)
}

// Beneficiary is a free data retrieval call binding the contract method 0x81008568.
//
// Solidity: function beneficiary(address _prover) view returns(address)
func (_Prover *ProverCallerSession) Beneficiary(_prover common.Address) (common.Address, error) {
	return _Prover.Contract.Beneficiary(&_Prover.CallOpts, _prover)
}

// IsPaused is a free data retrieval call binding the contract method 0x5b14f183.
//
// Solidity: function isPaused(address _prover) view returns(bool)
func (_Prover *ProverCaller) IsPaused(opts *bind.CallOpts, _prover common.Address) (bool, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "isPaused", _prover)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0x5b14f183.
//
// Solidity: function isPaused(address _prover) view returns(bool)
func (_Prover *ProverSession) IsPaused(_prover common.Address) (bool, error) {
	return _Prover.Contract.IsPaused(&_Prover.CallOpts, _prover)
}

// IsPaused is a free data retrieval call binding the contract method 0x5b14f183.
//
// Solidity: function isPaused(address _prover) view returns(bool)
func (_Prover *ProverCallerSession) IsPaused(_prover common.Address) (bool, error) {
	return _Prover.Contract.IsPaused(&_Prover.CallOpts, _prover)
}

// IsVMTypeSupported is a free data retrieval call binding the contract method 0x17490cbc.
//
// Solidity: function isVMTypeSupported(address _prover, uint256 _type) view returns(bool)
func (_Prover *ProverCaller) IsVMTypeSupported(opts *bind.CallOpts, _prover common.Address, _type *big.Int) (bool, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "isVMTypeSupported", _prover, _type)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVMTypeSupported is a free data retrieval call binding the contract method 0x17490cbc.
//
// Solidity: function isVMTypeSupported(address _prover, uint256 _type) view returns(bool)
func (_Prover *ProverSession) IsVMTypeSupported(_prover common.Address, _type *big.Int) (bool, error) {
	return _Prover.Contract.IsVMTypeSupported(&_Prover.CallOpts, _prover, _type)
}

// IsVMTypeSupported is a free data retrieval call binding the contract method 0x17490cbc.
//
// Solidity: function isVMTypeSupported(address _prover, uint256 _type) view returns(bool)
func (_Prover *ProverCallerSession) IsVMTypeSupported(_prover common.Address, _type *big.Int) (bool, error) {
	return _Prover.Contract.IsVMTypeSupported(&_Prover.CallOpts, _prover, _type)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Prover *ProverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Prover *ProverSession) Owner() (common.Address, error) {
	return _Prover.Contract.Owner(&_Prover.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Prover *ProverCallerSession) Owner() (common.Address, error) {
	return _Prover.Contract.Owner(&_Prover.CallOpts)
}

// RebateRatio is a free data retrieval call binding the contract method 0xff763ceb.
//
// Solidity: function rebateRatio(address _prover) view returns(uint16)
func (_Prover *ProverCaller) RebateRatio(opts *bind.CallOpts, _prover common.Address) (uint16, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "rebateRatio", _prover)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RebateRatio is a free data retrieval call binding the contract method 0xff763ceb.
//
// Solidity: function rebateRatio(address _prover) view returns(uint16)
func (_Prover *ProverSession) RebateRatio(_prover common.Address) (uint16, error) {
	return _Prover.Contract.RebateRatio(&_Prover.CallOpts, _prover)
}

// RebateRatio is a free data retrieval call binding the contract method 0xff763ceb.
//
// Solidity: function rebateRatio(address _prover) view returns(uint16)
func (_Prover *ProverCallerSession) RebateRatio(_prover common.Address) (uint16, error) {
	return _Prover.Contract.RebateRatio(&_Prover.CallOpts, _prover)
}

// AddVMType is a paid mutator transaction binding the contract method 0x298cad7a.
//
// Solidity: function addVMType(uint256 _type) returns()
func (_Prover *ProverTransactor) AddVMType(opts *bind.TransactOpts, _type *big.Int) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "addVMType", _type)
}

// AddVMType is a paid mutator transaction binding the contract method 0x298cad7a.
//
// Solidity: function addVMType(uint256 _type) returns()
func (_Prover *ProverSession) AddVMType(_type *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.AddVMType(&_Prover.TransactOpts, _type)
}

// AddVMType is a paid mutator transaction binding the contract method 0x298cad7a.
//
// Solidity: function addVMType(uint256 _type) returns()
func (_Prover *ProverTransactorSession) AddVMType(_type *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.AddVMType(&_Prover.TransactOpts, _type)
}

// ChangeBeneficiary is a paid mutator transaction binding the contract method 0xdc070657.
//
// Solidity: function changeBeneficiary(address _beneficiary) returns()
func (_Prover *ProverTransactor) ChangeBeneficiary(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "changeBeneficiary", _beneficiary)
}

// ChangeBeneficiary is a paid mutator transaction binding the contract method 0xdc070657.
//
// Solidity: function changeBeneficiary(address _beneficiary) returns()
func (_Prover *ProverSession) ChangeBeneficiary(_beneficiary common.Address) (*types.Transaction, error) {
	return _Prover.Contract.ChangeBeneficiary(&_Prover.TransactOpts, _beneficiary)
}

// ChangeBeneficiary is a paid mutator transaction binding the contract method 0xdc070657.
//
// Solidity: function changeBeneficiary(address _beneficiary) returns()
func (_Prover *ProverTransactorSession) ChangeBeneficiary(_beneficiary common.Address) (*types.Transaction, error) {
	return _Prover.Contract.ChangeBeneficiary(&_Prover.TransactOpts, _beneficiary)
}

// DelVMType is a paid mutator transaction binding the contract method 0x0df813dd.
//
// Solidity: function delVMType(uint256 _type) returns()
func (_Prover *ProverTransactor) DelVMType(opts *bind.TransactOpts, _type *big.Int) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "delVMType", _type)
}

// DelVMType is a paid mutator transaction binding the contract method 0x0df813dd.
//
// Solidity: function delVMType(uint256 _type) returns()
func (_Prover *ProverSession) DelVMType(_type *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.DelVMType(&_Prover.TransactOpts, _type)
}

// DelVMType is a paid mutator transaction binding the contract method 0x0df813dd.
//
// Solidity: function delVMType(uint256 _type) returns()
func (_Prover *ProverTransactorSession) DelVMType(_type *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.DelVMType(&_Prover.TransactOpts, _type)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Prover *ProverTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Prover *ProverSession) Initialize() (*types.Transaction, error) {
	return _Prover.Contract.Initialize(&_Prover.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Prover *ProverTransactorSession) Initialize() (*types.Transaction, error) {
	return _Prover.Contract.Initialize(&_Prover.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Prover *ProverTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Prover *ProverSession) Pause() (*types.Transaction, error) {
	return _Prover.Contract.Pause(&_Prover.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Prover *ProverTransactorSession) Pause() (*types.Transaction, error) {
	return _Prover.Contract.Pause(&_Prover.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Prover *ProverTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Prover *ProverSession) Register() (*types.Transaction, error) {
	return _Prover.Contract.Register(&_Prover.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Prover *ProverTransactorSession) Register() (*types.Transaction, error) {
	return _Prover.Contract.Register(&_Prover.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Prover *ProverTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Prover *ProverSession) RenounceOwnership() (*types.Transaction, error) {
	return _Prover.Contract.RenounceOwnership(&_Prover.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Prover *ProverTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Prover.Contract.RenounceOwnership(&_Prover.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Prover *ProverTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Prover *ProverSession) Resume() (*types.Transaction, error) {
	return _Prover.Contract.Resume(&_Prover.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Prover *ProverTransactorSession) Resume() (*types.Transaction, error) {
	return _Prover.Contract.Resume(&_Prover.TransactOpts)
}

// SetRebateRatio is a paid mutator transaction binding the contract method 0x36d06a28.
//
// Solidity: function setRebateRatio(uint16 _ratio) returns()
func (_Prover *ProverTransactor) SetRebateRatio(opts *bind.TransactOpts, _ratio uint16) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "setRebateRatio", _ratio)
}

// SetRebateRatio is a paid mutator transaction binding the contract method 0x36d06a28.
//
// Solidity: function setRebateRatio(uint16 _ratio) returns()
func (_Prover *ProverSession) SetRebateRatio(_ratio uint16) (*types.Transaction, error) {
	return _Prover.Contract.SetRebateRatio(&_Prover.TransactOpts, _ratio)
}

// SetRebateRatio is a paid mutator transaction binding the contract method 0x36d06a28.
//
// Solidity: function setRebateRatio(uint16 _ratio) returns()
func (_Prover *ProverTransactorSession) SetRebateRatio(_ratio uint16) (*types.Transaction, error) {
	return _Prover.Contract.SetRebateRatio(&_Prover.TransactOpts, _ratio)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Prover *ProverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Prover *ProverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Prover.Contract.TransferOwnership(&_Prover.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Prover *ProverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Prover.Contract.TransferOwnership(&_Prover.TransactOpts, newOwner)
}

// ProverBeneficiarySetIterator is returned from FilterBeneficiarySet and is used to iterate over the raw logs and unpacked data for BeneficiarySet events raised by the Prover contract.
type ProverBeneficiarySetIterator struct {
	Event *ProverBeneficiarySet // Event containing the contract specifics and raw log

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
func (it *ProverBeneficiarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverBeneficiarySet)
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
		it.Event = new(ProverBeneficiarySet)
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
func (it *ProverBeneficiarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverBeneficiarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverBeneficiarySet represents a BeneficiarySet event raised by the Prover contract.
type ProverBeneficiarySet struct {
	Prover      common.Address
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBeneficiarySet is a free log retrieval operation binding the contract event 0x2906d223dc4163733bb374af8641c7e9ae256e2bae53c90e0c9a2be2e611ae44.
//
// Solidity: event BeneficiarySet(address indexed prover, address indexed beneficiary)
func (_Prover *ProverFilterer) FilterBeneficiarySet(opts *bind.FilterOpts, prover []common.Address, beneficiary []common.Address) (*ProverBeneficiarySetIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "BeneficiarySet", proverRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &ProverBeneficiarySetIterator{contract: _Prover.contract, event: "BeneficiarySet", logs: logs, sub: sub}, nil
}

// WatchBeneficiarySet is a free log subscription operation binding the contract event 0x2906d223dc4163733bb374af8641c7e9ae256e2bae53c90e0c9a2be2e611ae44.
//
// Solidity: event BeneficiarySet(address indexed prover, address indexed beneficiary)
func (_Prover *ProverFilterer) WatchBeneficiarySet(opts *bind.WatchOpts, sink chan<- *ProverBeneficiarySet, prover []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "BeneficiarySet", proverRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverBeneficiarySet)
				if err := _Prover.contract.UnpackLog(event, "BeneficiarySet", log); err != nil {
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

// ParseBeneficiarySet is a log parse operation binding the contract event 0x2906d223dc4163733bb374af8641c7e9ae256e2bae53c90e0c9a2be2e611ae44.
//
// Solidity: event BeneficiarySet(address indexed prover, address indexed beneficiary)
func (_Prover *ProverFilterer) ParseBeneficiarySet(log types.Log) (*ProverBeneficiarySet, error) {
	event := new(ProverBeneficiarySet)
	if err := _Prover.contract.UnpackLog(event, "BeneficiarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Prover contract.
type ProverInitializedIterator struct {
	Event *ProverInitialized // Event containing the contract specifics and raw log

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
func (it *ProverInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverInitialized)
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
		it.Event = new(ProverInitialized)
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
func (it *ProverInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverInitialized represents a Initialized event raised by the Prover contract.
type ProverInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Prover *ProverFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProverInitializedIterator, error) {

	logs, sub, err := _Prover.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProverInitializedIterator{contract: _Prover.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Prover *ProverFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProverInitialized) (event.Subscription, error) {

	logs, sub, err := _Prover.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverInitialized)
				if err := _Prover.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Prover *ProverFilterer) ParseInitialized(log types.Log) (*ProverInitialized, error) {
	event := new(ProverInitialized)
	if err := _Prover.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Prover contract.
type ProverOwnershipTransferredIterator struct {
	Event *ProverOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverOwnershipTransferred)
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
		it.Event = new(ProverOwnershipTransferred)
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
func (it *ProverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverOwnershipTransferred represents a OwnershipTransferred event raised by the Prover contract.
type ProverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Prover *ProverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProverOwnershipTransferredIterator{contract: _Prover.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Prover *ProverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverOwnershipTransferred)
				if err := _Prover.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Prover *ProverFilterer) ParseOwnershipTransferred(log types.Log) (*ProverOwnershipTransferred, error) {
	event := new(ProverOwnershipTransferred)
	if err := _Prover.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverProverPausedIterator is returned from FilterProverPaused and is used to iterate over the raw logs and unpacked data for ProverPaused events raised by the Prover contract.
type ProverProverPausedIterator struct {
	Event *ProverProverPaused // Event containing the contract specifics and raw log

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
func (it *ProverProverPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverProverPaused)
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
		it.Event = new(ProverProverPaused)
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
func (it *ProverProverPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverProverPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverProverPaused represents a ProverPaused event raised by the Prover contract.
type ProverProverPaused struct {
	Prover common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProverPaused is a free log retrieval operation binding the contract event 0xb60833f9b39cb3f4532bc36a284e3a48f4f2ff9c4a057a295568315811c3daff.
//
// Solidity: event ProverPaused(address indexed prover)
func (_Prover *ProverFilterer) FilterProverPaused(opts *bind.FilterOpts, prover []common.Address) (*ProverProverPausedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "ProverPaused", proverRule)
	if err != nil {
		return nil, err
	}
	return &ProverProverPausedIterator{contract: _Prover.contract, event: "ProverPaused", logs: logs, sub: sub}, nil
}

// WatchProverPaused is a free log subscription operation binding the contract event 0xb60833f9b39cb3f4532bc36a284e3a48f4f2ff9c4a057a295568315811c3daff.
//
// Solidity: event ProverPaused(address indexed prover)
func (_Prover *ProverFilterer) WatchProverPaused(opts *bind.WatchOpts, sink chan<- *ProverProverPaused, prover []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "ProverPaused", proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverProverPaused)
				if err := _Prover.contract.UnpackLog(event, "ProverPaused", log); err != nil {
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

// ParseProverPaused is a log parse operation binding the contract event 0xb60833f9b39cb3f4532bc36a284e3a48f4f2ff9c4a057a295568315811c3daff.
//
// Solidity: event ProverPaused(address indexed prover)
func (_Prover *ProverFilterer) ParseProverPaused(log types.Log) (*ProverProverPaused, error) {
	event := new(ProverProverPaused)
	if err := _Prover.contract.UnpackLog(event, "ProverPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverProverResumedIterator is returned from FilterProverResumed and is used to iterate over the raw logs and unpacked data for ProverResumed events raised by the Prover contract.
type ProverProverResumedIterator struct {
	Event *ProverProverResumed // Event containing the contract specifics and raw log

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
func (it *ProverProverResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverProverResumed)
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
		it.Event = new(ProverProverResumed)
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
func (it *ProverProverResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverProverResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverProverResumed represents a ProverResumed event raised by the Prover contract.
type ProverProverResumed struct {
	Prover common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProverResumed is a free log retrieval operation binding the contract event 0xfedc704ee832701f3256eda12d9b4abb087fbbdf584108e7aef66b6e07abdaab.
//
// Solidity: event ProverResumed(address indexed prover)
func (_Prover *ProverFilterer) FilterProverResumed(opts *bind.FilterOpts, prover []common.Address) (*ProverProverResumedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "ProverResumed", proverRule)
	if err != nil {
		return nil, err
	}
	return &ProverProverResumedIterator{contract: _Prover.contract, event: "ProverResumed", logs: logs, sub: sub}, nil
}

// WatchProverResumed is a free log subscription operation binding the contract event 0xfedc704ee832701f3256eda12d9b4abb087fbbdf584108e7aef66b6e07abdaab.
//
// Solidity: event ProverResumed(address indexed prover)
func (_Prover *ProverFilterer) WatchProverResumed(opts *bind.WatchOpts, sink chan<- *ProverProverResumed, prover []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "ProverResumed", proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverProverResumed)
				if err := _Prover.contract.UnpackLog(event, "ProverResumed", log); err != nil {
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

// ParseProverResumed is a log parse operation binding the contract event 0xfedc704ee832701f3256eda12d9b4abb087fbbdf584108e7aef66b6e07abdaab.
//
// Solidity: event ProverResumed(address indexed prover)
func (_Prover *ProverFilterer) ParseProverResumed(log types.Log) (*ProverProverResumed, error) {
	event := new(ProverProverResumed)
	if err := _Prover.contract.UnpackLog(event, "ProverResumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverRebateRatioSetIterator is returned from FilterRebateRatioSet and is used to iterate over the raw logs and unpacked data for RebateRatioSet events raised by the Prover contract.
type ProverRebateRatioSetIterator struct {
	Event *ProverRebateRatioSet // Event containing the contract specifics and raw log

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
func (it *ProverRebateRatioSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverRebateRatioSet)
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
		it.Event = new(ProverRebateRatioSet)
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
func (it *ProverRebateRatioSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverRebateRatioSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverRebateRatioSet represents a RebateRatioSet event raised by the Prover contract.
type ProverRebateRatioSet struct {
	Prover common.Address
	Ratio  uint16
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRebateRatioSet is a free log retrieval operation binding the contract event 0xa0f5fcdc71c05b9272e6317da907366e3ebb2ef84f089a53c325c3a9cb8d4cbd.
//
// Solidity: event RebateRatioSet(address indexed prover, uint16 ratio)
func (_Prover *ProverFilterer) FilterRebateRatioSet(opts *bind.FilterOpts, prover []common.Address) (*ProverRebateRatioSetIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "RebateRatioSet", proverRule)
	if err != nil {
		return nil, err
	}
	return &ProverRebateRatioSetIterator{contract: _Prover.contract, event: "RebateRatioSet", logs: logs, sub: sub}, nil
}

// WatchRebateRatioSet is a free log subscription operation binding the contract event 0xa0f5fcdc71c05b9272e6317da907366e3ebb2ef84f089a53c325c3a9cb8d4cbd.
//
// Solidity: event RebateRatioSet(address indexed prover, uint16 ratio)
func (_Prover *ProverFilterer) WatchRebateRatioSet(opts *bind.WatchOpts, sink chan<- *ProverRebateRatioSet, prover []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "RebateRatioSet", proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverRebateRatioSet)
				if err := _Prover.contract.UnpackLog(event, "RebateRatioSet", log); err != nil {
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

// ParseRebateRatioSet is a log parse operation binding the contract event 0xa0f5fcdc71c05b9272e6317da907366e3ebb2ef84f089a53c325c3a9cb8d4cbd.
//
// Solidity: event RebateRatioSet(address indexed prover, uint16 ratio)
func (_Prover *ProverFilterer) ParseRebateRatioSet(log types.Log) (*ProverRebateRatioSet, error) {
	event := new(ProverRebateRatioSet)
	if err := _Prover.contract.UnpackLog(event, "RebateRatioSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverVMTypeAddedIterator is returned from FilterVMTypeAdded and is used to iterate over the raw logs and unpacked data for VMTypeAdded events raised by the Prover contract.
type ProverVMTypeAddedIterator struct {
	Event *ProverVMTypeAdded // Event containing the contract specifics and raw log

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
func (it *ProverVMTypeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverVMTypeAdded)
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
		it.Event = new(ProverVMTypeAdded)
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
func (it *ProverVMTypeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverVMTypeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverVMTypeAdded represents a VMTypeAdded event raised by the Prover contract.
type ProverVMTypeAdded struct {
	Prover common.Address
	Typ    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVMTypeAdded is a free log retrieval operation binding the contract event 0x9584d49ad17d5b2fc346a4680a1be7b2a02bf95d734fd677d9df585e9d52a655.
//
// Solidity: event VMTypeAdded(address indexed prover, uint256 typ)
func (_Prover *ProverFilterer) FilterVMTypeAdded(opts *bind.FilterOpts, prover []common.Address) (*ProverVMTypeAddedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "VMTypeAdded", proverRule)
	if err != nil {
		return nil, err
	}
	return &ProverVMTypeAddedIterator{contract: _Prover.contract, event: "VMTypeAdded", logs: logs, sub: sub}, nil
}

// WatchVMTypeAdded is a free log subscription operation binding the contract event 0x9584d49ad17d5b2fc346a4680a1be7b2a02bf95d734fd677d9df585e9d52a655.
//
// Solidity: event VMTypeAdded(address indexed prover, uint256 typ)
func (_Prover *ProverFilterer) WatchVMTypeAdded(opts *bind.WatchOpts, sink chan<- *ProverVMTypeAdded, prover []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "VMTypeAdded", proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverVMTypeAdded)
				if err := _Prover.contract.UnpackLog(event, "VMTypeAdded", log); err != nil {
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

// ParseVMTypeAdded is a log parse operation binding the contract event 0x9584d49ad17d5b2fc346a4680a1be7b2a02bf95d734fd677d9df585e9d52a655.
//
// Solidity: event VMTypeAdded(address indexed prover, uint256 typ)
func (_Prover *ProverFilterer) ParseVMTypeAdded(log types.Log) (*ProverVMTypeAdded, error) {
	event := new(ProverVMTypeAdded)
	if err := _Prover.contract.UnpackLog(event, "VMTypeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverVMTypeDeletedIterator is returned from FilterVMTypeDeleted and is used to iterate over the raw logs and unpacked data for VMTypeDeleted events raised by the Prover contract.
type ProverVMTypeDeletedIterator struct {
	Event *ProverVMTypeDeleted // Event containing the contract specifics and raw log

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
func (it *ProverVMTypeDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverVMTypeDeleted)
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
		it.Event = new(ProverVMTypeDeleted)
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
func (it *ProverVMTypeDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverVMTypeDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverVMTypeDeleted represents a VMTypeDeleted event raised by the Prover contract.
type ProverVMTypeDeleted struct {
	Prover common.Address
	Typ    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVMTypeDeleted is a free log retrieval operation binding the contract event 0xf1b01f10c033f88aa1a9daf399bdec09f261109268766ce78194396f3ff77ce8.
//
// Solidity: event VMTypeDeleted(address indexed prover, uint256 typ)
func (_Prover *ProverFilterer) FilterVMTypeDeleted(opts *bind.FilterOpts, prover []common.Address) (*ProverVMTypeDeletedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "VMTypeDeleted", proverRule)
	if err != nil {
		return nil, err
	}
	return &ProverVMTypeDeletedIterator{contract: _Prover.contract, event: "VMTypeDeleted", logs: logs, sub: sub}, nil
}

// WatchVMTypeDeleted is a free log subscription operation binding the contract event 0xf1b01f10c033f88aa1a9daf399bdec09f261109268766ce78194396f3ff77ce8.
//
// Solidity: event VMTypeDeleted(address indexed prover, uint256 typ)
func (_Prover *ProverFilterer) WatchVMTypeDeleted(opts *bind.WatchOpts, sink chan<- *ProverVMTypeDeleted, prover []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "VMTypeDeleted", proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverVMTypeDeleted)
				if err := _Prover.contract.UnpackLog(event, "VMTypeDeleted", log); err != nil {
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

// ParseVMTypeDeleted is a log parse operation binding the contract event 0xf1b01f10c033f88aa1a9daf399bdec09f261109268766ce78194396f3ff77ce8.
//
// Solidity: event VMTypeDeleted(address indexed prover, uint256 typ)
func (_Prover *ProverFilterer) ParseVMTypeDeleted(log types.Log) (*ProverVMTypeDeleted, error) {
	event := new(ProverVMTypeDeleted)
	if err := _Prover.contract.UnpackLog(event, "VMTypeDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
