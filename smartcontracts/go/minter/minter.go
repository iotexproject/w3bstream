// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package minter

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

// BlockInfo is an auto generated low-level Go binding around an user-defined struct.
type BlockInfo struct {
	Meta       [4]byte
	Prevhash   [32]byte
	MerkleRoot [32]byte
	Difficulty [4]byte
	Nonce      [8]byte
}

// Sequencer is an auto generated low-level Go binding around an user-defined struct.
type Sequencer struct {
	Addr        common.Address
	Operator    common.Address
	Beneficiary common.Address
}

// TaskAssignment is an auto generated low-level Go binding around an user-defined struct.
type TaskAssignment struct {
	ProjectId *big.Int
	TaskId    [32]byte
	Prover    common.Address
}

// MinterMetaData contains all meta data concerning the Minter contract.
var MinterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"difficulty\",\"type\":\"bytes4\"}],\"name\":\"DifficultySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"TargetDurationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"}],\"name\":\"TaskAllowanceSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adhocDifficulty\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentDifficulty\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"contractIDAO\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDAO\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"contractITaskManager\",\"name\":\"_tm\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"meta\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"prevhash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"difficulty\",\"type\":\"bytes4\"},{\"internalType\":\"bytes8\",\"name\":\"nonce\",\"type\":\"bytes8\"}],\"internalType\":\"structBlockInfo\",\"name\":\"blockinfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structSequencer\",\"name\":\"coinbase\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"taskId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"internalType\":\"structTaskAssignment[]\",\"name\":\"assignments\",\"type\":\"tuple[]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"difficulty\",\"type\":\"bytes4\"}],\"name\":\"setAdhocDifficulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setTargetDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"}],\"name\":\"setTaskAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tm\",\"outputs\":[{\"internalType\":\"contractITaskManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MinterABI is the input ABI used to generate the binding from.
// Deprecated: Use MinterMetaData.ABI instead.
var MinterABI = MinterMetaData.ABI

// Minter is an auto generated Go binding around an Ethereum contract.
type Minter struct {
	MinterCaller     // Read-only binding to the contract
	MinterTransactor // Write-only binding to the contract
	MinterFilterer   // Log filterer for contract events
}

// MinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinterSession struct {
	Contract     *Minter           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinterCallerSession struct {
	Contract *MinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinterTransactorSession struct {
	Contract     *MinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinterRaw struct {
	Contract *Minter // Generic contract binding to access the raw methods on
}

// MinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinterCallerRaw struct {
	Contract *MinterCaller // Generic read-only contract binding to access the raw methods on
}

// MinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinterTransactorRaw struct {
	Contract *MinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinter creates a new instance of Minter, bound to a specific deployed contract.
func NewMinter(address common.Address, backend bind.ContractBackend) (*Minter, error) {
	contract, err := bindMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Minter{MinterCaller: MinterCaller{contract: contract}, MinterTransactor: MinterTransactor{contract: contract}, MinterFilterer: MinterFilterer{contract: contract}}, nil
}

// NewMinterCaller creates a new read-only instance of Minter, bound to a specific deployed contract.
func NewMinterCaller(address common.Address, caller bind.ContractCaller) (*MinterCaller, error) {
	contract, err := bindMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinterCaller{contract: contract}, nil
}

// NewMinterTransactor creates a new write-only instance of Minter, bound to a specific deployed contract.
func NewMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*MinterTransactor, error) {
	contract, err := bindMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinterTransactor{contract: contract}, nil
}

// NewMinterFilterer creates a new log filterer instance of Minter, bound to a specific deployed contract.
func NewMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*MinterFilterer, error) {
	contract, err := bindMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinterFilterer{contract: contract}, nil
}

// bindMinter binds a generic wrapper to an already deployed contract.
func bindMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MinterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Minter *MinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Minter.Contract.MinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Minter *MinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Minter.Contract.MinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Minter *MinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Minter.Contract.MinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Minter *MinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Minter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Minter *MinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Minter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Minter *MinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Minter.Contract.contract.Transact(opts, method, params...)
}

// AdhocDifficulty is a free data retrieval call binding the contract method 0xeaac9a77.
//
// Solidity: function adhocDifficulty() view returns(bytes4)
func (_Minter *MinterCaller) AdhocDifficulty(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "adhocDifficulty")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// AdhocDifficulty is a free data retrieval call binding the contract method 0xeaac9a77.
//
// Solidity: function adhocDifficulty() view returns(bytes4)
func (_Minter *MinterSession) AdhocDifficulty() ([4]byte, error) {
	return _Minter.Contract.AdhocDifficulty(&_Minter.CallOpts)
}

// AdhocDifficulty is a free data retrieval call binding the contract method 0xeaac9a77.
//
// Solidity: function adhocDifficulty() view returns(bytes4)
func (_Minter *MinterCallerSession) AdhocDifficulty() ([4]byte, error) {
	return _Minter.Contract.AdhocDifficulty(&_Minter.CallOpts)
}

// CurrentDifficulty is a free data retrieval call binding the contract method 0x5c062d6c.
//
// Solidity: function currentDifficulty() view returns(bytes4)
func (_Minter *MinterCaller) CurrentDifficulty(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "currentDifficulty")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// CurrentDifficulty is a free data retrieval call binding the contract method 0x5c062d6c.
//
// Solidity: function currentDifficulty() view returns(bytes4)
func (_Minter *MinterSession) CurrentDifficulty() ([4]byte, error) {
	return _Minter.Contract.CurrentDifficulty(&_Minter.CallOpts)
}

// CurrentDifficulty is a free data retrieval call binding the contract method 0x5c062d6c.
//
// Solidity: function currentDifficulty() view returns(bytes4)
func (_Minter *MinterCallerSession) CurrentDifficulty() ([4]byte, error) {
	return _Minter.Contract.CurrentDifficulty(&_Minter.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Minter *MinterCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Minter *MinterSession) Dao() (common.Address, error) {
	return _Minter.Contract.Dao(&_Minter.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Minter *MinterCallerSession) Dao() (common.Address, error) {
	return _Minter.Contract.Dao(&_Minter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Minter *MinterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Minter *MinterSession) Owner() (common.Address, error) {
	return _Minter.Contract.Owner(&_Minter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Minter *MinterCallerSession) Owner() (common.Address, error) {
	return _Minter.Contract.Owner(&_Minter.CallOpts)
}

// TargetDuration is a free data retrieval call binding the contract method 0xf0cdec63.
//
// Solidity: function targetDuration() view returns(uint256)
func (_Minter *MinterCaller) TargetDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "targetDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetDuration is a free data retrieval call binding the contract method 0xf0cdec63.
//
// Solidity: function targetDuration() view returns(uint256)
func (_Minter *MinterSession) TargetDuration() (*big.Int, error) {
	return _Minter.Contract.TargetDuration(&_Minter.CallOpts)
}

// TargetDuration is a free data retrieval call binding the contract method 0xf0cdec63.
//
// Solidity: function targetDuration() view returns(uint256)
func (_Minter *MinterCallerSession) TargetDuration() (*big.Int, error) {
	return _Minter.Contract.TargetDuration(&_Minter.CallOpts)
}

// TaskAllowance is a free data retrieval call binding the contract method 0xb6c11e51.
//
// Solidity: function taskAllowance() view returns(uint256)
func (_Minter *MinterCaller) TaskAllowance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "taskAllowance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskAllowance is a free data retrieval call binding the contract method 0xb6c11e51.
//
// Solidity: function taskAllowance() view returns(uint256)
func (_Minter *MinterSession) TaskAllowance() (*big.Int, error) {
	return _Minter.Contract.TaskAllowance(&_Minter.CallOpts)
}

// TaskAllowance is a free data retrieval call binding the contract method 0xb6c11e51.
//
// Solidity: function taskAllowance() view returns(uint256)
func (_Minter *MinterCallerSession) TaskAllowance() (*big.Int, error) {
	return _Minter.Contract.TaskAllowance(&_Minter.CallOpts)
}

// Tm is a free data retrieval call binding the contract method 0xcdb314d3.
//
// Solidity: function tm() view returns(address)
func (_Minter *MinterCaller) Tm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "tm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tm is a free data retrieval call binding the contract method 0xcdb314d3.
//
// Solidity: function tm() view returns(address)
func (_Minter *MinterSession) Tm() (common.Address, error) {
	return _Minter.Contract.Tm(&_Minter.CallOpts)
}

// Tm is a free data retrieval call binding the contract method 0xcdb314d3.
//
// Solidity: function tm() view returns(address)
func (_Minter *MinterCallerSession) Tm() (common.Address, error) {
	return _Minter.Contract.Tm(&_Minter.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _dao, address _tm) returns()
func (_Minter *MinterTransactor) Initialize(opts *bind.TransactOpts, _dao common.Address, _tm common.Address) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "initialize", _dao, _tm)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _dao, address _tm) returns()
func (_Minter *MinterSession) Initialize(_dao common.Address, _tm common.Address) (*types.Transaction, error) {
	return _Minter.Contract.Initialize(&_Minter.TransactOpts, _dao, _tm)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _dao, address _tm) returns()
func (_Minter *MinterTransactorSession) Initialize(_dao common.Address, _tm common.Address) (*types.Transaction, error) {
	return _Minter.Contract.Initialize(&_Minter.TransactOpts, _dao, _tm)
}

// Mint is a paid mutator transaction binding the contract method 0x5e83afcb.
//
// Solidity: function mint((bytes4,bytes32,bytes32,bytes4,bytes8) blockinfo, (address,address,address) coinbase, (uint256,bytes32,address)[] assignments) returns()
func (_Minter *MinterTransactor) Mint(opts *bind.TransactOpts, blockinfo BlockInfo, coinbase Sequencer, assignments []TaskAssignment) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "mint", blockinfo, coinbase, assignments)
}

// Mint is a paid mutator transaction binding the contract method 0x5e83afcb.
//
// Solidity: function mint((bytes4,bytes32,bytes32,bytes4,bytes8) blockinfo, (address,address,address) coinbase, (uint256,bytes32,address)[] assignments) returns()
func (_Minter *MinterSession) Mint(blockinfo BlockInfo, coinbase Sequencer, assignments []TaskAssignment) (*types.Transaction, error) {
	return _Minter.Contract.Mint(&_Minter.TransactOpts, blockinfo, coinbase, assignments)
}

// Mint is a paid mutator transaction binding the contract method 0x5e83afcb.
//
// Solidity: function mint((bytes4,bytes32,bytes32,bytes4,bytes8) blockinfo, (address,address,address) coinbase, (uint256,bytes32,address)[] assignments) returns()
func (_Minter *MinterTransactorSession) Mint(blockinfo BlockInfo, coinbase Sequencer, assignments []TaskAssignment) (*types.Transaction, error) {
	return _Minter.Contract.Mint(&_Minter.TransactOpts, blockinfo, coinbase, assignments)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Minter *MinterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Minter *MinterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Minter.Contract.RenounceOwnership(&_Minter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Minter *MinterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Minter.Contract.RenounceOwnership(&_Minter.TransactOpts)
}

// SetAdhocDifficulty is a paid mutator transaction binding the contract method 0xefdcc949.
//
// Solidity: function setAdhocDifficulty(bytes4 difficulty) returns()
func (_Minter *MinterTransactor) SetAdhocDifficulty(opts *bind.TransactOpts, difficulty [4]byte) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setAdhocDifficulty", difficulty)
}

// SetAdhocDifficulty is a paid mutator transaction binding the contract method 0xefdcc949.
//
// Solidity: function setAdhocDifficulty(bytes4 difficulty) returns()
func (_Minter *MinterSession) SetAdhocDifficulty(difficulty [4]byte) (*types.Transaction, error) {
	return _Minter.Contract.SetAdhocDifficulty(&_Minter.TransactOpts, difficulty)
}

// SetAdhocDifficulty is a paid mutator transaction binding the contract method 0xefdcc949.
//
// Solidity: function setAdhocDifficulty(bytes4 difficulty) returns()
func (_Minter *MinterTransactorSession) SetAdhocDifficulty(difficulty [4]byte) (*types.Transaction, error) {
	return _Minter.Contract.SetAdhocDifficulty(&_Minter.TransactOpts, difficulty)
}

// SetTargetDuration is a paid mutator transaction binding the contract method 0xa00171ea.
//
// Solidity: function setTargetDuration(uint256 duration) returns()
func (_Minter *MinterTransactor) SetTargetDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setTargetDuration", duration)
}

// SetTargetDuration is a paid mutator transaction binding the contract method 0xa00171ea.
//
// Solidity: function setTargetDuration(uint256 duration) returns()
func (_Minter *MinterSession) SetTargetDuration(duration *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetTargetDuration(&_Minter.TransactOpts, duration)
}

// SetTargetDuration is a paid mutator transaction binding the contract method 0xa00171ea.
//
// Solidity: function setTargetDuration(uint256 duration) returns()
func (_Minter *MinterTransactorSession) SetTargetDuration(duration *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetTargetDuration(&_Minter.TransactOpts, duration)
}

// SetTaskAllowance is a paid mutator transaction binding the contract method 0x92d894ec.
//
// Solidity: function setTaskAllowance(uint256 allowance) returns()
func (_Minter *MinterTransactor) SetTaskAllowance(opts *bind.TransactOpts, allowance *big.Int) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setTaskAllowance", allowance)
}

// SetTaskAllowance is a paid mutator transaction binding the contract method 0x92d894ec.
//
// Solidity: function setTaskAllowance(uint256 allowance) returns()
func (_Minter *MinterSession) SetTaskAllowance(allowance *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetTaskAllowance(&_Minter.TransactOpts, allowance)
}

// SetTaskAllowance is a paid mutator transaction binding the contract method 0x92d894ec.
//
// Solidity: function setTaskAllowance(uint256 allowance) returns()
func (_Minter *MinterTransactorSession) SetTaskAllowance(allowance *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetTaskAllowance(&_Minter.TransactOpts, allowance)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Minter *MinterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Minter *MinterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Minter.Contract.TransferOwnership(&_Minter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Minter *MinterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Minter.Contract.TransferOwnership(&_Minter.TransactOpts, newOwner)
}

// MinterDifficultySetIterator is returned from FilterDifficultySet and is used to iterate over the raw logs and unpacked data for DifficultySet events raised by the Minter contract.
type MinterDifficultySetIterator struct {
	Event *MinterDifficultySet // Event containing the contract specifics and raw log

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
func (it *MinterDifficultySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterDifficultySet)
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
		it.Event = new(MinterDifficultySet)
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
func (it *MinterDifficultySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterDifficultySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterDifficultySet represents a DifficultySet event raised by the Minter contract.
type MinterDifficultySet struct {
	Difficulty [4]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDifficultySet is a free log retrieval operation binding the contract event 0xff051f56d0d099feca55af838a05490b0f309ce78f78c58d46a7c32059b10ecc.
//
// Solidity: event DifficultySet(bytes4 difficulty)
func (_Minter *MinterFilterer) FilterDifficultySet(opts *bind.FilterOpts) (*MinterDifficultySetIterator, error) {

	logs, sub, err := _Minter.contract.FilterLogs(opts, "DifficultySet")
	if err != nil {
		return nil, err
	}
	return &MinterDifficultySetIterator{contract: _Minter.contract, event: "DifficultySet", logs: logs, sub: sub}, nil
}

// WatchDifficultySet is a free log subscription operation binding the contract event 0xff051f56d0d099feca55af838a05490b0f309ce78f78c58d46a7c32059b10ecc.
//
// Solidity: event DifficultySet(bytes4 difficulty)
func (_Minter *MinterFilterer) WatchDifficultySet(opts *bind.WatchOpts, sink chan<- *MinterDifficultySet) (event.Subscription, error) {

	logs, sub, err := _Minter.contract.WatchLogs(opts, "DifficultySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterDifficultySet)
				if err := _Minter.contract.UnpackLog(event, "DifficultySet", log); err != nil {
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

// ParseDifficultySet is a log parse operation binding the contract event 0xff051f56d0d099feca55af838a05490b0f309ce78f78c58d46a7c32059b10ecc.
//
// Solidity: event DifficultySet(bytes4 difficulty)
func (_Minter *MinterFilterer) ParseDifficultySet(log types.Log) (*MinterDifficultySet, error) {
	event := new(MinterDifficultySet)
	if err := _Minter.contract.UnpackLog(event, "DifficultySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Minter contract.
type MinterInitializedIterator struct {
	Event *MinterInitialized // Event containing the contract specifics and raw log

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
func (it *MinterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterInitialized)
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
		it.Event = new(MinterInitialized)
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
func (it *MinterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterInitialized represents a Initialized event raised by the Minter contract.
type MinterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Minter *MinterFilterer) FilterInitialized(opts *bind.FilterOpts) (*MinterInitializedIterator, error) {

	logs, sub, err := _Minter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MinterInitializedIterator{contract: _Minter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Minter *MinterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MinterInitialized) (event.Subscription, error) {

	logs, sub, err := _Minter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterInitialized)
				if err := _Minter.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Minter *MinterFilterer) ParseInitialized(log types.Log) (*MinterInitialized, error) {
	event := new(MinterInitialized)
	if err := _Minter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Minter contract.
type MinterOwnershipTransferredIterator struct {
	Event *MinterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MinterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterOwnershipTransferred)
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
		it.Event = new(MinterOwnershipTransferred)
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
func (it *MinterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterOwnershipTransferred represents a OwnershipTransferred event raised by the Minter contract.
type MinterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Minter *MinterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MinterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Minter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MinterOwnershipTransferredIterator{contract: _Minter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Minter *MinterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MinterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Minter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterOwnershipTransferred)
				if err := _Minter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Minter *MinterFilterer) ParseOwnershipTransferred(log types.Log) (*MinterOwnershipTransferred, error) {
	event := new(MinterOwnershipTransferred)
	if err := _Minter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterTargetDurationSetIterator is returned from FilterTargetDurationSet and is used to iterate over the raw logs and unpacked data for TargetDurationSet events raised by the Minter contract.
type MinterTargetDurationSetIterator struct {
	Event *MinterTargetDurationSet // Event containing the contract specifics and raw log

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
func (it *MinterTargetDurationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterTargetDurationSet)
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
		it.Event = new(MinterTargetDurationSet)
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
func (it *MinterTargetDurationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterTargetDurationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterTargetDurationSet represents a TargetDurationSet event raised by the Minter contract.
type MinterTargetDurationSet struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTargetDurationSet is a free log retrieval operation binding the contract event 0x09d59fc02af479aae6b9ec944a55713b23cd8748872c7621d79b4ddf645950c6.
//
// Solidity: event TargetDurationSet(uint256 duration)
func (_Minter *MinterFilterer) FilterTargetDurationSet(opts *bind.FilterOpts) (*MinterTargetDurationSetIterator, error) {

	logs, sub, err := _Minter.contract.FilterLogs(opts, "TargetDurationSet")
	if err != nil {
		return nil, err
	}
	return &MinterTargetDurationSetIterator{contract: _Minter.contract, event: "TargetDurationSet", logs: logs, sub: sub}, nil
}

// WatchTargetDurationSet is a free log subscription operation binding the contract event 0x09d59fc02af479aae6b9ec944a55713b23cd8748872c7621d79b4ddf645950c6.
//
// Solidity: event TargetDurationSet(uint256 duration)
func (_Minter *MinterFilterer) WatchTargetDurationSet(opts *bind.WatchOpts, sink chan<- *MinterTargetDurationSet) (event.Subscription, error) {

	logs, sub, err := _Minter.contract.WatchLogs(opts, "TargetDurationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterTargetDurationSet)
				if err := _Minter.contract.UnpackLog(event, "TargetDurationSet", log); err != nil {
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

// ParseTargetDurationSet is a log parse operation binding the contract event 0x09d59fc02af479aae6b9ec944a55713b23cd8748872c7621d79b4ddf645950c6.
//
// Solidity: event TargetDurationSet(uint256 duration)
func (_Minter *MinterFilterer) ParseTargetDurationSet(log types.Log) (*MinterTargetDurationSet, error) {
	event := new(MinterTargetDurationSet)
	if err := _Minter.contract.UnpackLog(event, "TargetDurationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterTaskAllowanceSetIterator is returned from FilterTaskAllowanceSet and is used to iterate over the raw logs and unpacked data for TaskAllowanceSet events raised by the Minter contract.
type MinterTaskAllowanceSetIterator struct {
	Event *MinterTaskAllowanceSet // Event containing the contract specifics and raw log

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
func (it *MinterTaskAllowanceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterTaskAllowanceSet)
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
		it.Event = new(MinterTaskAllowanceSet)
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
func (it *MinterTaskAllowanceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterTaskAllowanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterTaskAllowanceSet represents a TaskAllowanceSet event raised by the Minter contract.
type MinterTaskAllowanceSet struct {
	Allowance *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskAllowanceSet is a free log retrieval operation binding the contract event 0x3b703b8b82fb6fcf76d8b6e487a021501299c2cb2050c7c9cce7b47fa7325bf1.
//
// Solidity: event TaskAllowanceSet(uint256 allowance)
func (_Minter *MinterFilterer) FilterTaskAllowanceSet(opts *bind.FilterOpts) (*MinterTaskAllowanceSetIterator, error) {

	logs, sub, err := _Minter.contract.FilterLogs(opts, "TaskAllowanceSet")
	if err != nil {
		return nil, err
	}
	return &MinterTaskAllowanceSetIterator{contract: _Minter.contract, event: "TaskAllowanceSet", logs: logs, sub: sub}, nil
}

// WatchTaskAllowanceSet is a free log subscription operation binding the contract event 0x3b703b8b82fb6fcf76d8b6e487a021501299c2cb2050c7c9cce7b47fa7325bf1.
//
// Solidity: event TaskAllowanceSet(uint256 allowance)
func (_Minter *MinterFilterer) WatchTaskAllowanceSet(opts *bind.WatchOpts, sink chan<- *MinterTaskAllowanceSet) (event.Subscription, error) {

	logs, sub, err := _Minter.contract.WatchLogs(opts, "TaskAllowanceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterTaskAllowanceSet)
				if err := _Minter.contract.UnpackLog(event, "TaskAllowanceSet", log); err != nil {
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

// ParseTaskAllowanceSet is a log parse operation binding the contract event 0x3b703b8b82fb6fcf76d8b6e487a021501299c2cb2050c7c9cce7b47fa7325bf1.
//
// Solidity: event TaskAllowanceSet(uint256 allowance)
func (_Minter *MinterFilterer) ParseTaskAllowanceSet(log types.Log) (*MinterTaskAllowanceSet, error) {
	event := new(MinterTaskAllowanceSet)
	if err := _Minter.contract.UnpackLog(event, "TaskAllowanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
