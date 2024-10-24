// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package taskmanager

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

// TaskAssignment is an auto generated low-level Go binding around an user-defined struct.
type TaskAssignment struct {
	ProjectId *big.Int
	TaskId    [32]byte
	Hash      [32]byte
	Signature []byte
	Prover    common.Address
}

// TaskmanagerMetaData contains all meta data concerning the Taskmanager contract.
var TaskmanagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"taskId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"TaskAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"taskId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"TaskSettled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"taskId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"internalType\":\"structTaskAssignment[]\",\"name\":\"taskAssignments\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"assign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"taskId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"internalType\":\"structTaskAssignment\",\"name\":\"assignment\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"assign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"debits\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_debits\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_projectReward\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proverStore\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"projectReward\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverStore\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"taskId\",\"type\":\"bytes32\"}],\"name\":\"recall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardForProver\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardForSequencer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"taskId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TaskmanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskmanagerMetaData.ABI instead.
var TaskmanagerABI = TaskmanagerMetaData.ABI

// Taskmanager is an auto generated Go binding around an Ethereum contract.
type Taskmanager struct {
	TaskmanagerCaller     // Read-only binding to the contract
	TaskmanagerTransactor // Write-only binding to the contract
	TaskmanagerFilterer   // Log filterer for contract events
}

// TaskmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskmanagerSession struct {
	Contract     *Taskmanager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskmanagerCallerSession struct {
	Contract *TaskmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TaskmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskmanagerTransactorSession struct {
	Contract     *TaskmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TaskmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskmanagerRaw struct {
	Contract *Taskmanager // Generic contract binding to access the raw methods on
}

// TaskmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskmanagerCallerRaw struct {
	Contract *TaskmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// TaskmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskmanagerTransactorRaw struct {
	Contract *TaskmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaskmanager creates a new instance of Taskmanager, bound to a specific deployed contract.
func NewTaskmanager(address common.Address, backend bind.ContractBackend) (*Taskmanager, error) {
	contract, err := bindTaskmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Taskmanager{TaskmanagerCaller: TaskmanagerCaller{contract: contract}, TaskmanagerTransactor: TaskmanagerTransactor{contract: contract}, TaskmanagerFilterer: TaskmanagerFilterer{contract: contract}}, nil
}

// NewTaskmanagerCaller creates a new read-only instance of Taskmanager, bound to a specific deployed contract.
func NewTaskmanagerCaller(address common.Address, caller bind.ContractCaller) (*TaskmanagerCaller, error) {
	contract, err := bindTaskmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskmanagerCaller{contract: contract}, nil
}

// NewTaskmanagerTransactor creates a new write-only instance of Taskmanager, bound to a specific deployed contract.
func NewTaskmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskmanagerTransactor, error) {
	contract, err := bindTaskmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskmanagerTransactor{contract: contract}, nil
}

// NewTaskmanagerFilterer creates a new log filterer instance of Taskmanager, bound to a specific deployed contract.
func NewTaskmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskmanagerFilterer, error) {
	contract, err := bindTaskmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskmanagerFilterer{contract: contract}, nil
}

// bindTaskmanager binds a generic wrapper to an already deployed contract.
func bindTaskmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaskmanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Taskmanager *TaskmanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Taskmanager.Contract.TaskmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Taskmanager *TaskmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Taskmanager.Contract.TaskmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Taskmanager *TaskmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Taskmanager.Contract.TaskmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Taskmanager *TaskmanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Taskmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Taskmanager *TaskmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Taskmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Taskmanager *TaskmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Taskmanager.Contract.contract.Transact(opts, method, params...)
}

// Debits is a free data retrieval call binding the contract method 0x6f0f11e5.
//
// Solidity: function debits() view returns(address)
func (_Taskmanager *TaskmanagerCaller) Debits(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Taskmanager.contract.Call(opts, &out, "debits")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Debits is a free data retrieval call binding the contract method 0x6f0f11e5.
//
// Solidity: function debits() view returns(address)
func (_Taskmanager *TaskmanagerSession) Debits() (common.Address, error) {
	return _Taskmanager.Contract.Debits(&_Taskmanager.CallOpts)
}

// Debits is a free data retrieval call binding the contract method 0x6f0f11e5.
//
// Solidity: function debits() view returns(address)
func (_Taskmanager *TaskmanagerCallerSession) Debits() (common.Address, error) {
	return _Taskmanager.Contract.Debits(&_Taskmanager.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Taskmanager *TaskmanagerCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Taskmanager.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Taskmanager *TaskmanagerSession) Operators(arg0 common.Address) (bool, error) {
	return _Taskmanager.Contract.Operators(&_Taskmanager.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Taskmanager *TaskmanagerCallerSession) Operators(arg0 common.Address) (bool, error) {
	return _Taskmanager.Contract.Operators(&_Taskmanager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Taskmanager *TaskmanagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Taskmanager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Taskmanager *TaskmanagerSession) Owner() (common.Address, error) {
	return _Taskmanager.Contract.Owner(&_Taskmanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Taskmanager *TaskmanagerCallerSession) Owner() (common.Address, error) {
	return _Taskmanager.Contract.Owner(&_Taskmanager.CallOpts)
}

// ProjectReward is a free data retrieval call binding the contract method 0xa6095890.
//
// Solidity: function projectReward() view returns(address)
func (_Taskmanager *TaskmanagerCaller) ProjectReward(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Taskmanager.contract.Call(opts, &out, "projectReward")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProjectReward is a free data retrieval call binding the contract method 0xa6095890.
//
// Solidity: function projectReward() view returns(address)
func (_Taskmanager *TaskmanagerSession) ProjectReward() (common.Address, error) {
	return _Taskmanager.Contract.ProjectReward(&_Taskmanager.CallOpts)
}

// ProjectReward is a free data retrieval call binding the contract method 0xa6095890.
//
// Solidity: function projectReward() view returns(address)
func (_Taskmanager *TaskmanagerCallerSession) ProjectReward() (common.Address, error) {
	return _Taskmanager.Contract.ProjectReward(&_Taskmanager.CallOpts)
}

// ProverStore is a free data retrieval call binding the contract method 0x79b851f6.
//
// Solidity: function proverStore() view returns(address)
func (_Taskmanager *TaskmanagerCaller) ProverStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Taskmanager.contract.Call(opts, &out, "proverStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverStore is a free data retrieval call binding the contract method 0x79b851f6.
//
// Solidity: function proverStore() view returns(address)
func (_Taskmanager *TaskmanagerSession) ProverStore() (common.Address, error) {
	return _Taskmanager.Contract.ProverStore(&_Taskmanager.CallOpts)
}

// ProverStore is a free data retrieval call binding the contract method 0x79b851f6.
//
// Solidity: function proverStore() view returns(address)
func (_Taskmanager *TaskmanagerCallerSession) ProverStore() (common.Address, error) {
	return _Taskmanager.Contract.ProverStore(&_Taskmanager.CallOpts)
}

// Records is a free data retrieval call binding the contract method 0x63c7a7c8.
//
// Solidity: function records(uint256 , bytes32 ) view returns(bytes32 hash, address owner, address sequencer, address prover, uint256 rewardForProver, uint256 rewardForSequencer, uint256 deadline, bool settled)
func (_Taskmanager *TaskmanagerCaller) Records(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	Hash               [32]byte
	Owner              common.Address
	Sequencer          common.Address
	Prover             common.Address
	RewardForProver    *big.Int
	RewardForSequencer *big.Int
	Deadline           *big.Int
	Settled            bool
}, error) {
	var out []interface{}
	err := _Taskmanager.contract.Call(opts, &out, "records", arg0, arg1)

	outstruct := new(struct {
		Hash               [32]byte
		Owner              common.Address
		Sequencer          common.Address
		Prover             common.Address
		RewardForProver    *big.Int
		RewardForSequencer *big.Int
		Deadline           *big.Int
		Settled            bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Sequencer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Prover = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.RewardForProver = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RewardForSequencer = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Settled = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Records is a free data retrieval call binding the contract method 0x63c7a7c8.
//
// Solidity: function records(uint256 , bytes32 ) view returns(bytes32 hash, address owner, address sequencer, address prover, uint256 rewardForProver, uint256 rewardForSequencer, uint256 deadline, bool settled)
func (_Taskmanager *TaskmanagerSession) Records(arg0 *big.Int, arg1 [32]byte) (struct {
	Hash               [32]byte
	Owner              common.Address
	Sequencer          common.Address
	Prover             common.Address
	RewardForProver    *big.Int
	RewardForSequencer *big.Int
	Deadline           *big.Int
	Settled            bool
}, error) {
	return _Taskmanager.Contract.Records(&_Taskmanager.CallOpts, arg0, arg1)
}

// Records is a free data retrieval call binding the contract method 0x63c7a7c8.
//
// Solidity: function records(uint256 , bytes32 ) view returns(bytes32 hash, address owner, address sequencer, address prover, uint256 rewardForProver, uint256 rewardForSequencer, uint256 deadline, bool settled)
func (_Taskmanager *TaskmanagerCallerSession) Records(arg0 *big.Int, arg1 [32]byte) (struct {
	Hash               [32]byte
	Owner              common.Address
	Sequencer          common.Address
	Prover             common.Address
	RewardForProver    *big.Int
	RewardForSequencer *big.Int
	Deadline           *big.Int
	Settled            bool
}, error) {
	return _Taskmanager.Contract.Records(&_Taskmanager.CallOpts, arg0, arg1)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_Taskmanager *TaskmanagerTransactor) AddOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Taskmanager.contract.Transact(opts, "addOperator", operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_Taskmanager *TaskmanagerSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _Taskmanager.Contract.AddOperator(&_Taskmanager.TransactOpts, operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_Taskmanager *TaskmanagerTransactorSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _Taskmanager.Contract.AddOperator(&_Taskmanager.TransactOpts, operator)
}

// Assign is a paid mutator transaction binding the contract method 0x11ba23ab.
//
// Solidity: function assign((uint256,bytes32,bytes32,bytes,address)[] taskAssignments, address sequencer, uint256 deadline) returns()
func (_Taskmanager *TaskmanagerTransactor) Assign(opts *bind.TransactOpts, taskAssignments []TaskAssignment, sequencer common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Taskmanager.contract.Transact(opts, "assign", taskAssignments, sequencer, deadline)
}

// Assign is a paid mutator transaction binding the contract method 0x11ba23ab.
//
// Solidity: function assign((uint256,bytes32,bytes32,bytes,address)[] taskAssignments, address sequencer, uint256 deadline) returns()
func (_Taskmanager *TaskmanagerSession) Assign(taskAssignments []TaskAssignment, sequencer common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Taskmanager.Contract.Assign(&_Taskmanager.TransactOpts, taskAssignments, sequencer, deadline)
}

// Assign is a paid mutator transaction binding the contract method 0x11ba23ab.
//
// Solidity: function assign((uint256,bytes32,bytes32,bytes,address)[] taskAssignments, address sequencer, uint256 deadline) returns()
func (_Taskmanager *TaskmanagerTransactorSession) Assign(taskAssignments []TaskAssignment, sequencer common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Taskmanager.Contract.Assign(&_Taskmanager.TransactOpts, taskAssignments, sequencer, deadline)
}

// Assign0 is a paid mutator transaction binding the contract method 0x79365ddb.
//
// Solidity: function assign((uint256,bytes32,bytes32,bytes,address) assignment, address sequencer, uint256 deadline) returns()
func (_Taskmanager *TaskmanagerTransactor) Assign0(opts *bind.TransactOpts, assignment TaskAssignment, sequencer common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Taskmanager.contract.Transact(opts, "assign0", assignment, sequencer, deadline)
}

// Assign0 is a paid mutator transaction binding the contract method 0x79365ddb.
//
// Solidity: function assign((uint256,bytes32,bytes32,bytes,address) assignment, address sequencer, uint256 deadline) returns()
func (_Taskmanager *TaskmanagerSession) Assign0(assignment TaskAssignment, sequencer common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Taskmanager.Contract.Assign0(&_Taskmanager.TransactOpts, assignment, sequencer, deadline)
}

// Assign0 is a paid mutator transaction binding the contract method 0x79365ddb.
//
// Solidity: function assign((uint256,bytes32,bytes32,bytes,address) assignment, address sequencer, uint256 deadline) returns()
func (_Taskmanager *TaskmanagerTransactorSession) Assign0(assignment TaskAssignment, sequencer common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Taskmanager.Contract.Assign0(&_Taskmanager.TransactOpts, assignment, sequencer, deadline)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _debits, address _projectReward, address _proverStore) returns()
func (_Taskmanager *TaskmanagerTransactor) Initialize(opts *bind.TransactOpts, _debits common.Address, _projectReward common.Address, _proverStore common.Address) (*types.Transaction, error) {
	return _Taskmanager.contract.Transact(opts, "initialize", _debits, _projectReward, _proverStore)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _debits, address _projectReward, address _proverStore) returns()
func (_Taskmanager *TaskmanagerSession) Initialize(_debits common.Address, _projectReward common.Address, _proverStore common.Address) (*types.Transaction, error) {
	return _Taskmanager.Contract.Initialize(&_Taskmanager.TransactOpts, _debits, _projectReward, _proverStore)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _debits, address _projectReward, address _proverStore) returns()
func (_Taskmanager *TaskmanagerTransactorSession) Initialize(_debits common.Address, _projectReward common.Address, _proverStore common.Address) (*types.Transaction, error) {
	return _Taskmanager.Contract.Initialize(&_Taskmanager.TransactOpts, _debits, _projectReward, _proverStore)
}

// Recall is a paid mutator transaction binding the contract method 0x585407e0.
//
// Solidity: function recall(uint256 projectId, bytes32 taskId) returns()
func (_Taskmanager *TaskmanagerTransactor) Recall(opts *bind.TransactOpts, projectId *big.Int, taskId [32]byte) (*types.Transaction, error) {
	return _Taskmanager.contract.Transact(opts, "recall", projectId, taskId)
}

// Recall is a paid mutator transaction binding the contract method 0x585407e0.
//
// Solidity: function recall(uint256 projectId, bytes32 taskId) returns()
func (_Taskmanager *TaskmanagerSession) Recall(projectId *big.Int, taskId [32]byte) (*types.Transaction, error) {
	return _Taskmanager.Contract.Recall(&_Taskmanager.TransactOpts, projectId, taskId)
}

// Recall is a paid mutator transaction binding the contract method 0x585407e0.
//
// Solidity: function recall(uint256 projectId, bytes32 taskId) returns()
func (_Taskmanager *TaskmanagerTransactorSession) Recall(projectId *big.Int, taskId [32]byte) (*types.Transaction, error) {
	return _Taskmanager.Contract.Recall(&_Taskmanager.TransactOpts, projectId, taskId)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_Taskmanager *TaskmanagerTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Taskmanager.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_Taskmanager *TaskmanagerSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _Taskmanager.Contract.RemoveOperator(&_Taskmanager.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_Taskmanager *TaskmanagerTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _Taskmanager.Contract.RemoveOperator(&_Taskmanager.TransactOpts, operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Taskmanager *TaskmanagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Taskmanager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Taskmanager *TaskmanagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Taskmanager.Contract.RenounceOwnership(&_Taskmanager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Taskmanager *TaskmanagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Taskmanager.Contract.RenounceOwnership(&_Taskmanager.TransactOpts)
}

// Settle is a paid mutator transaction binding the contract method 0x05c42eeb.
//
// Solidity: function settle(uint256 projectId, bytes32 taskId, address prover) returns()
func (_Taskmanager *TaskmanagerTransactor) Settle(opts *bind.TransactOpts, projectId *big.Int, taskId [32]byte, prover common.Address) (*types.Transaction, error) {
	return _Taskmanager.contract.Transact(opts, "settle", projectId, taskId, prover)
}

// Settle is a paid mutator transaction binding the contract method 0x05c42eeb.
//
// Solidity: function settle(uint256 projectId, bytes32 taskId, address prover) returns()
func (_Taskmanager *TaskmanagerSession) Settle(projectId *big.Int, taskId [32]byte, prover common.Address) (*types.Transaction, error) {
	return _Taskmanager.Contract.Settle(&_Taskmanager.TransactOpts, projectId, taskId, prover)
}

// Settle is a paid mutator transaction binding the contract method 0x05c42eeb.
//
// Solidity: function settle(uint256 projectId, bytes32 taskId, address prover) returns()
func (_Taskmanager *TaskmanagerTransactorSession) Settle(projectId *big.Int, taskId [32]byte, prover common.Address) (*types.Transaction, error) {
	return _Taskmanager.Contract.Settle(&_Taskmanager.TransactOpts, projectId, taskId, prover)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Taskmanager *TaskmanagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Taskmanager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Taskmanager *TaskmanagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Taskmanager.Contract.TransferOwnership(&_Taskmanager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Taskmanager *TaskmanagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Taskmanager.Contract.TransferOwnership(&_Taskmanager.TransactOpts, newOwner)
}

// TaskmanagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Taskmanager contract.
type TaskmanagerInitializedIterator struct {
	Event *TaskmanagerInitialized // Event containing the contract specifics and raw log

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
func (it *TaskmanagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskmanagerInitialized)
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
		it.Event = new(TaskmanagerInitialized)
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
func (it *TaskmanagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskmanagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskmanagerInitialized represents a Initialized event raised by the Taskmanager contract.
type TaskmanagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Taskmanager *TaskmanagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaskmanagerInitializedIterator, error) {

	logs, sub, err := _Taskmanager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaskmanagerInitializedIterator{contract: _Taskmanager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Taskmanager *TaskmanagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaskmanagerInitialized) (event.Subscription, error) {

	logs, sub, err := _Taskmanager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskmanagerInitialized)
				if err := _Taskmanager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Taskmanager *TaskmanagerFilterer) ParseInitialized(log types.Log) (*TaskmanagerInitialized, error) {
	event := new(TaskmanagerInitialized)
	if err := _Taskmanager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskmanagerOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the Taskmanager contract.
type TaskmanagerOperatorAddedIterator struct {
	Event *TaskmanagerOperatorAdded // Event containing the contract specifics and raw log

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
func (it *TaskmanagerOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskmanagerOperatorAdded)
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
		it.Event = new(TaskmanagerOperatorAdded)
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
func (it *TaskmanagerOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskmanagerOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskmanagerOperatorAdded represents a OperatorAdded event raised by the Taskmanager contract.
type TaskmanagerOperatorAdded struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address operator)
func (_Taskmanager *TaskmanagerFilterer) FilterOperatorAdded(opts *bind.FilterOpts) (*TaskmanagerOperatorAddedIterator, error) {

	logs, sub, err := _Taskmanager.contract.FilterLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return &TaskmanagerOperatorAddedIterator{contract: _Taskmanager.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address operator)
func (_Taskmanager *TaskmanagerFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *TaskmanagerOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _Taskmanager.contract.WatchLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskmanagerOperatorAdded)
				if err := _Taskmanager.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address operator)
func (_Taskmanager *TaskmanagerFilterer) ParseOperatorAdded(log types.Log) (*TaskmanagerOperatorAdded, error) {
	event := new(TaskmanagerOperatorAdded)
	if err := _Taskmanager.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskmanagerOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the Taskmanager contract.
type TaskmanagerOperatorRemovedIterator struct {
	Event *TaskmanagerOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *TaskmanagerOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskmanagerOperatorRemoved)
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
		it.Event = new(TaskmanagerOperatorRemoved)
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
func (it *TaskmanagerOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskmanagerOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskmanagerOperatorRemoved represents a OperatorRemoved event raised by the Taskmanager contract.
type TaskmanagerOperatorRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address operator)
func (_Taskmanager *TaskmanagerFilterer) FilterOperatorRemoved(opts *bind.FilterOpts) (*TaskmanagerOperatorRemovedIterator, error) {

	logs, sub, err := _Taskmanager.contract.FilterLogs(opts, "OperatorRemoved")
	if err != nil {
		return nil, err
	}
	return &TaskmanagerOperatorRemovedIterator{contract: _Taskmanager.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address operator)
func (_Taskmanager *TaskmanagerFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *TaskmanagerOperatorRemoved) (event.Subscription, error) {

	logs, sub, err := _Taskmanager.contract.WatchLogs(opts, "OperatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskmanagerOperatorRemoved)
				if err := _Taskmanager.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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

// ParseOperatorRemoved is a log parse operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address operator)
func (_Taskmanager *TaskmanagerFilterer) ParseOperatorRemoved(log types.Log) (*TaskmanagerOperatorRemoved, error) {
	event := new(TaskmanagerOperatorRemoved)
	if err := _Taskmanager.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskmanagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Taskmanager contract.
type TaskmanagerOwnershipTransferredIterator struct {
	Event *TaskmanagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaskmanagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskmanagerOwnershipTransferred)
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
		it.Event = new(TaskmanagerOwnershipTransferred)
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
func (it *TaskmanagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskmanagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskmanagerOwnershipTransferred represents a OwnershipTransferred event raised by the Taskmanager contract.
type TaskmanagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Taskmanager *TaskmanagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaskmanagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Taskmanager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaskmanagerOwnershipTransferredIterator{contract: _Taskmanager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Taskmanager *TaskmanagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaskmanagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Taskmanager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskmanagerOwnershipTransferred)
				if err := _Taskmanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Taskmanager *TaskmanagerFilterer) ParseOwnershipTransferred(log types.Log) (*TaskmanagerOwnershipTransferred, error) {
	event := new(TaskmanagerOwnershipTransferred)
	if err := _Taskmanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskmanagerTaskAssignedIterator is returned from FilterTaskAssigned and is used to iterate over the raw logs and unpacked data for TaskAssigned events raised by the Taskmanager contract.
type TaskmanagerTaskAssignedIterator struct {
	Event *TaskmanagerTaskAssigned // Event containing the contract specifics and raw log

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
func (it *TaskmanagerTaskAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskmanagerTaskAssigned)
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
		it.Event = new(TaskmanagerTaskAssigned)
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
func (it *TaskmanagerTaskAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskmanagerTaskAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskmanagerTaskAssigned represents a TaskAssigned event raised by the Taskmanager contract.
type TaskmanagerTaskAssigned struct {
	ProjectId *big.Int
	TaskId    [32]byte
	Prover    common.Address
	Deadline  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskAssigned is a free log retrieval operation binding the contract event 0x45fbdf2605d3db829aa556ebcbe72c28b04d5ac864de9f5352854b761feb0e2b.
//
// Solidity: event TaskAssigned(uint256 indexed projectId, bytes32 indexed taskId, address indexed prover, uint256 deadline)
func (_Taskmanager *TaskmanagerFilterer) FilterTaskAssigned(opts *bind.FilterOpts, projectId []*big.Int, taskId [][32]byte, prover []common.Address) (*TaskmanagerTaskAssignedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _Taskmanager.contract.FilterLogs(opts, "TaskAssigned", projectIdRule, taskIdRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &TaskmanagerTaskAssignedIterator{contract: _Taskmanager.contract, event: "TaskAssigned", logs: logs, sub: sub}, nil
}

// WatchTaskAssigned is a free log subscription operation binding the contract event 0x45fbdf2605d3db829aa556ebcbe72c28b04d5ac864de9f5352854b761feb0e2b.
//
// Solidity: event TaskAssigned(uint256 indexed projectId, bytes32 indexed taskId, address indexed prover, uint256 deadline)
func (_Taskmanager *TaskmanagerFilterer) WatchTaskAssigned(opts *bind.WatchOpts, sink chan<- *TaskmanagerTaskAssigned, projectId []*big.Int, taskId [][32]byte, prover []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _Taskmanager.contract.WatchLogs(opts, "TaskAssigned", projectIdRule, taskIdRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskmanagerTaskAssigned)
				if err := _Taskmanager.contract.UnpackLog(event, "TaskAssigned", log); err != nil {
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

// ParseTaskAssigned is a log parse operation binding the contract event 0x45fbdf2605d3db829aa556ebcbe72c28b04d5ac864de9f5352854b761feb0e2b.
//
// Solidity: event TaskAssigned(uint256 indexed projectId, bytes32 indexed taskId, address indexed prover, uint256 deadline)
func (_Taskmanager *TaskmanagerFilterer) ParseTaskAssigned(log types.Log) (*TaskmanagerTaskAssigned, error) {
	event := new(TaskmanagerTaskAssigned)
	if err := _Taskmanager.contract.UnpackLog(event, "TaskAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskmanagerTaskSettledIterator is returned from FilterTaskSettled and is used to iterate over the raw logs and unpacked data for TaskSettled events raised by the Taskmanager contract.
type TaskmanagerTaskSettledIterator struct {
	Event *TaskmanagerTaskSettled // Event containing the contract specifics and raw log

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
func (it *TaskmanagerTaskSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskmanagerTaskSettled)
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
		it.Event = new(TaskmanagerTaskSettled)
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
func (it *TaskmanagerTaskSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskmanagerTaskSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskmanagerTaskSettled represents a TaskSettled event raised by the Taskmanager contract.
type TaskmanagerTaskSettled struct {
	ProjectId *big.Int
	TaskId    [32]byte
	Prover    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskSettled is a free log retrieval operation binding the contract event 0x76776b8a3b3fbc171162894ab4103f74c531c25c7e505e21bcb78b52b62392b6.
//
// Solidity: event TaskSettled(uint256 indexed projectId, bytes32 indexed taskId, address prover)
func (_Taskmanager *TaskmanagerFilterer) FilterTaskSettled(opts *bind.FilterOpts, projectId []*big.Int, taskId [][32]byte) (*TaskmanagerTaskSettledIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _Taskmanager.contract.FilterLogs(opts, "TaskSettled", projectIdRule, taskIdRule)
	if err != nil {
		return nil, err
	}
	return &TaskmanagerTaskSettledIterator{contract: _Taskmanager.contract, event: "TaskSettled", logs: logs, sub: sub}, nil
}

// WatchTaskSettled is a free log subscription operation binding the contract event 0x76776b8a3b3fbc171162894ab4103f74c531c25c7e505e21bcb78b52b62392b6.
//
// Solidity: event TaskSettled(uint256 indexed projectId, bytes32 indexed taskId, address prover)
func (_Taskmanager *TaskmanagerFilterer) WatchTaskSettled(opts *bind.WatchOpts, sink chan<- *TaskmanagerTaskSettled, projectId []*big.Int, taskId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _Taskmanager.contract.WatchLogs(opts, "TaskSettled", projectIdRule, taskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskmanagerTaskSettled)
				if err := _Taskmanager.contract.UnpackLog(event, "TaskSettled", log); err != nil {
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

// ParseTaskSettled is a log parse operation binding the contract event 0x76776b8a3b3fbc171162894ab4103f74c531c25c7e505e21bcb78b52b62392b6.
//
// Solidity: event TaskSettled(uint256 indexed projectId, bytes32 indexed taskId, address prover)
func (_Taskmanager *TaskmanagerFilterer) ParseTaskSettled(log types.Log) (*TaskmanagerTaskSettled, error) {
	event := new(TaskmanagerTaskSettled)
	if err := _Taskmanager.contract.UnpackLog(event, "TaskSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
