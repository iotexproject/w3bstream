// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package projectreward

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

// ProjectRewardMetaData contains all meta data concerning the ProjectReward contract.
var ProjectRewardMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardAmountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"name\":\"RewardTokenSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_project\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"project\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"rewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"}],\"name\":\"setRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProjectRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use ProjectRewardMetaData.ABI instead.
var ProjectRewardABI = ProjectRewardMetaData.ABI

// ProjectReward is an auto generated Go binding around an Ethereum contract.
type ProjectReward struct {
	ProjectRewardCaller     // Read-only binding to the contract
	ProjectRewardTransactor // Write-only binding to the contract
	ProjectRewardFilterer   // Log filterer for contract events
}

// ProjectRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProjectRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProjectRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProjectRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProjectRewardSession struct {
	Contract     *ProjectReward    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProjectRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProjectRewardCallerSession struct {
	Contract *ProjectRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ProjectRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProjectRewardTransactorSession struct {
	Contract     *ProjectRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ProjectRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProjectRewardRaw struct {
	Contract *ProjectReward // Generic contract binding to access the raw methods on
}

// ProjectRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProjectRewardCallerRaw struct {
	Contract *ProjectRewardCaller // Generic read-only contract binding to access the raw methods on
}

// ProjectRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProjectRewardTransactorRaw struct {
	Contract *ProjectRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProjectReward creates a new instance of ProjectReward, bound to a specific deployed contract.
func NewProjectReward(address common.Address, backend bind.ContractBackend) (*ProjectReward, error) {
	contract, err := bindProjectReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProjectReward{ProjectRewardCaller: ProjectRewardCaller{contract: contract}, ProjectRewardTransactor: ProjectRewardTransactor{contract: contract}, ProjectRewardFilterer: ProjectRewardFilterer{contract: contract}}, nil
}

// NewProjectRewardCaller creates a new read-only instance of ProjectReward, bound to a specific deployed contract.
func NewProjectRewardCaller(address common.Address, caller bind.ContractCaller) (*ProjectRewardCaller, error) {
	contract, err := bindProjectReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectRewardCaller{contract: contract}, nil
}

// NewProjectRewardTransactor creates a new write-only instance of ProjectReward, bound to a specific deployed contract.
func NewProjectRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*ProjectRewardTransactor, error) {
	contract, err := bindProjectReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectRewardTransactor{contract: contract}, nil
}

// NewProjectRewardFilterer creates a new log filterer instance of ProjectReward, bound to a specific deployed contract.
func NewProjectRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*ProjectRewardFilterer, error) {
	contract, err := bindProjectReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProjectRewardFilterer{contract: contract}, nil
}

// bindProjectReward binds a generic wrapper to an already deployed contract.
func bindProjectReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProjectRewardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProjectReward *ProjectRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProjectReward.Contract.ProjectRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProjectReward *ProjectRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectReward.Contract.ProjectRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProjectReward *ProjectRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProjectReward.Contract.ProjectRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProjectReward *ProjectRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProjectReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProjectReward *ProjectRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProjectReward *ProjectRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProjectReward.Contract.contract.Transact(opts, method, params...)
}

// IsPaused is a free data retrieval call binding the contract method 0xbdf2a43c.
//
// Solidity: function isPaused(uint256 _projectId) view returns(bool)
func (_ProjectReward *ProjectRewardCaller) IsPaused(opts *bind.CallOpts, _projectId *big.Int) (bool, error) {
	var out []interface{}
	err := _ProjectReward.contract.Call(opts, &out, "isPaused", _projectId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xbdf2a43c.
//
// Solidity: function isPaused(uint256 _projectId) view returns(bool)
func (_ProjectReward *ProjectRewardSession) IsPaused(_projectId *big.Int) (bool, error) {
	return _ProjectReward.Contract.IsPaused(&_ProjectReward.CallOpts, _projectId)
}

// IsPaused is a free data retrieval call binding the contract method 0xbdf2a43c.
//
// Solidity: function isPaused(uint256 _projectId) view returns(bool)
func (_ProjectReward *ProjectRewardCallerSession) IsPaused(_projectId *big.Int) (bool, error) {
	return _ProjectReward.Contract.IsPaused(&_ProjectReward.CallOpts, _projectId)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ProjectReward *ProjectRewardCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProjectReward.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ProjectReward *ProjectRewardSession) Operator() (common.Address, error) {
	return _ProjectReward.Contract.Operator(&_ProjectReward.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ProjectReward *ProjectRewardCallerSession) Operator() (common.Address, error) {
	return _ProjectReward.Contract.Operator(&_ProjectReward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProjectReward *ProjectRewardCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProjectReward.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProjectReward *ProjectRewardSession) Owner() (common.Address, error) {
	return _ProjectReward.Contract.Owner(&_ProjectReward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProjectReward *ProjectRewardCallerSession) Owner() (common.Address, error) {
	return _ProjectReward.Contract.Owner(&_ProjectReward.CallOpts)
}

// Project is a free data retrieval call binding the contract method 0xf60ca60d.
//
// Solidity: function project() view returns(address)
func (_ProjectReward *ProjectRewardCaller) Project(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProjectReward.contract.Call(opts, &out, "project")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Project is a free data retrieval call binding the contract method 0xf60ca60d.
//
// Solidity: function project() view returns(address)
func (_ProjectReward *ProjectRewardSession) Project() (common.Address, error) {
	return _ProjectReward.Contract.Project(&_ProjectReward.CallOpts)
}

// Project is a free data retrieval call binding the contract method 0xf60ca60d.
//
// Solidity: function project() view returns(address)
func (_ProjectReward *ProjectRewardCallerSession) Project() (common.Address, error) {
	return _ProjectReward.Contract.Project(&_ProjectReward.CallOpts)
}

// RewardAmount is a free data retrieval call binding the contract method 0xc85e4d49.
//
// Solidity: function rewardAmount(address owner, uint256 id) view returns(uint256)
func (_ProjectReward *ProjectRewardCaller) RewardAmount(opts *bind.CallOpts, owner common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ProjectReward.contract.Call(opts, &out, "rewardAmount", owner, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardAmount is a free data retrieval call binding the contract method 0xc85e4d49.
//
// Solidity: function rewardAmount(address owner, uint256 id) view returns(uint256)
func (_ProjectReward *ProjectRewardSession) RewardAmount(owner common.Address, id *big.Int) (*big.Int, error) {
	return _ProjectReward.Contract.RewardAmount(&_ProjectReward.CallOpts, owner, id)
}

// RewardAmount is a free data retrieval call binding the contract method 0xc85e4d49.
//
// Solidity: function rewardAmount(address owner, uint256 id) view returns(uint256)
func (_ProjectReward *ProjectRewardCallerSession) RewardAmount(owner common.Address, id *big.Int) (*big.Int, error) {
	return _ProjectReward.Contract.RewardAmount(&_ProjectReward.CallOpts, owner, id)
}

// RewardToken is a free data retrieval call binding the contract method 0x509b6c3f.
//
// Solidity: function rewardToken(uint256 _id) view returns(address)
func (_ProjectReward *ProjectRewardCaller) RewardToken(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ProjectReward.contract.Call(opts, &out, "rewardToken", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0x509b6c3f.
//
// Solidity: function rewardToken(uint256 _id) view returns(address)
func (_ProjectReward *ProjectRewardSession) RewardToken(_id *big.Int) (common.Address, error) {
	return _ProjectReward.Contract.RewardToken(&_ProjectReward.CallOpts, _id)
}

// RewardToken is a free data retrieval call binding the contract method 0x509b6c3f.
//
// Solidity: function rewardToken(uint256 _id) view returns(address)
func (_ProjectReward *ProjectRewardCallerSession) RewardToken(_id *big.Int) (common.Address, error) {
	return _ProjectReward.Contract.RewardToken(&_ProjectReward.CallOpts, _id)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _project) returns()
func (_ProjectReward *ProjectRewardTransactor) Initialize(opts *bind.TransactOpts, _project common.Address) (*types.Transaction, error) {
	return _ProjectReward.contract.Transact(opts, "initialize", _project)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _project) returns()
func (_ProjectReward *ProjectRewardSession) Initialize(_project common.Address) (*types.Transaction, error) {
	return _ProjectReward.Contract.Initialize(&_ProjectReward.TransactOpts, _project)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _project) returns()
func (_ProjectReward *ProjectRewardTransactorSession) Initialize(_project common.Address) (*types.Transaction, error) {
	return _ProjectReward.Contract.Initialize(&_ProjectReward.TransactOpts, _project)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProjectReward *ProjectRewardTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProjectReward.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProjectReward *ProjectRewardSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProjectReward.Contract.RenounceOwnership(&_ProjectReward.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProjectReward *ProjectRewardTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProjectReward.Contract.RenounceOwnership(&_ProjectReward.TransactOpts)
}

// SetReward is a paid mutator transaction binding the contract method 0xa47bd496.
//
// Solidity: function setReward(uint256 _id, uint256 _amount) returns()
func (_ProjectReward *ProjectRewardTransactor) SetReward(opts *bind.TransactOpts, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ProjectReward.contract.Transact(opts, "setReward", _id, _amount)
}

// SetReward is a paid mutator transaction binding the contract method 0xa47bd496.
//
// Solidity: function setReward(uint256 _id, uint256 _amount) returns()
func (_ProjectReward *ProjectRewardSession) SetReward(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ProjectReward.Contract.SetReward(&_ProjectReward.TransactOpts, _id, _amount)
}

// SetReward is a paid mutator transaction binding the contract method 0xa47bd496.
//
// Solidity: function setReward(uint256 _id, uint256 _amount) returns()
func (_ProjectReward *ProjectRewardTransactorSession) SetReward(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ProjectReward.Contract.SetReward(&_ProjectReward.TransactOpts, _id, _amount)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0xbdd13d4a.
//
// Solidity: function setRewardToken(uint256 _id, address _rewardToken) returns()
func (_ProjectReward *ProjectRewardTransactor) SetRewardToken(opts *bind.TransactOpts, _id *big.Int, _rewardToken common.Address) (*types.Transaction, error) {
	return _ProjectReward.contract.Transact(opts, "setRewardToken", _id, _rewardToken)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0xbdd13d4a.
//
// Solidity: function setRewardToken(uint256 _id, address _rewardToken) returns()
func (_ProjectReward *ProjectRewardSession) SetRewardToken(_id *big.Int, _rewardToken common.Address) (*types.Transaction, error) {
	return _ProjectReward.Contract.SetRewardToken(&_ProjectReward.TransactOpts, _id, _rewardToken)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0xbdd13d4a.
//
// Solidity: function setRewardToken(uint256 _id, address _rewardToken) returns()
func (_ProjectReward *ProjectRewardTransactorSession) SetRewardToken(_id *big.Int, _rewardToken common.Address) (*types.Transaction, error) {
	return _ProjectReward.Contract.SetRewardToken(&_ProjectReward.TransactOpts, _id, _rewardToken)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProjectReward *ProjectRewardTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProjectReward.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProjectReward *ProjectRewardSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProjectReward.Contract.TransferOwnership(&_ProjectReward.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProjectReward *ProjectRewardTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProjectReward.Contract.TransferOwnership(&_ProjectReward.TransactOpts, newOwner)
}

// ProjectRewardInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ProjectReward contract.
type ProjectRewardInitializedIterator struct {
	Event *ProjectRewardInitialized // Event containing the contract specifics and raw log

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
func (it *ProjectRewardInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectRewardInitialized)
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
		it.Event = new(ProjectRewardInitialized)
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
func (it *ProjectRewardInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectRewardInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectRewardInitialized represents a Initialized event raised by the ProjectReward contract.
type ProjectRewardInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ProjectReward *ProjectRewardFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProjectRewardInitializedIterator, error) {

	logs, sub, err := _ProjectReward.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProjectRewardInitializedIterator{contract: _ProjectReward.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ProjectReward *ProjectRewardFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProjectRewardInitialized) (event.Subscription, error) {

	logs, sub, err := _ProjectReward.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectRewardInitialized)
				if err := _ProjectReward.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ProjectReward *ProjectRewardFilterer) ParseInitialized(log types.Log) (*ProjectRewardInitialized, error) {
	event := new(ProjectRewardInitialized)
	if err := _ProjectReward.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectRewardOperatorSetIterator is returned from FilterOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorSet events raised by the ProjectReward contract.
type ProjectRewardOperatorSetIterator struct {
	Event *ProjectRewardOperatorSet // Event containing the contract specifics and raw log

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
func (it *ProjectRewardOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectRewardOperatorSet)
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
		it.Event = new(ProjectRewardOperatorSet)
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
func (it *ProjectRewardOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectRewardOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectRewardOperatorSet represents a OperatorSet event raised by the ProjectReward contract.
type ProjectRewardOperatorSet struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSet is a free log retrieval operation binding the contract event 0x99d737e0adf2c449d71890b86772885ec7959b152ddb265f76325b6e68e105d3.
//
// Solidity: event OperatorSet(address indexed operator)
func (_ProjectReward *ProjectRewardFilterer) FilterOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ProjectRewardOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ProjectReward.contract.FilterLogs(opts, "OperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ProjectRewardOperatorSetIterator{contract: _ProjectReward.contract, event: "OperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSet is a free log subscription operation binding the contract event 0x99d737e0adf2c449d71890b86772885ec7959b152ddb265f76325b6e68e105d3.
//
// Solidity: event OperatorSet(address indexed operator)
func (_ProjectReward *ProjectRewardFilterer) WatchOperatorSet(opts *bind.WatchOpts, sink chan<- *ProjectRewardOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ProjectReward.contract.WatchLogs(opts, "OperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectRewardOperatorSet)
				if err := _ProjectReward.contract.UnpackLog(event, "OperatorSet", log); err != nil {
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

// ParseOperatorSet is a log parse operation binding the contract event 0x99d737e0adf2c449d71890b86772885ec7959b152ddb265f76325b6e68e105d3.
//
// Solidity: event OperatorSet(address indexed operator)
func (_ProjectReward *ProjectRewardFilterer) ParseOperatorSet(log types.Log) (*ProjectRewardOperatorSet, error) {
	event := new(ProjectRewardOperatorSet)
	if err := _ProjectReward.contract.UnpackLog(event, "OperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectRewardOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProjectReward contract.
type ProjectRewardOwnershipTransferredIterator struct {
	Event *ProjectRewardOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProjectRewardOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectRewardOwnershipTransferred)
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
		it.Event = new(ProjectRewardOwnershipTransferred)
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
func (it *ProjectRewardOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectRewardOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectRewardOwnershipTransferred represents a OwnershipTransferred event raised by the ProjectReward contract.
type ProjectRewardOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProjectReward *ProjectRewardFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProjectRewardOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProjectReward.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProjectRewardOwnershipTransferredIterator{contract: _ProjectReward.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProjectReward *ProjectRewardFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProjectRewardOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProjectReward.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectRewardOwnershipTransferred)
				if err := _ProjectReward.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ProjectReward *ProjectRewardFilterer) ParseOwnershipTransferred(log types.Log) (*ProjectRewardOwnershipTransferred, error) {
	event := new(ProjectRewardOwnershipTransferred)
	if err := _ProjectReward.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectRewardRewardAmountSetIterator is returned from FilterRewardAmountSet and is used to iterate over the raw logs and unpacked data for RewardAmountSet events raised by the ProjectReward contract.
type ProjectRewardRewardAmountSetIterator struct {
	Event *ProjectRewardRewardAmountSet // Event containing the contract specifics and raw log

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
func (it *ProjectRewardRewardAmountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectRewardRewardAmountSet)
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
		it.Event = new(ProjectRewardRewardAmountSet)
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
func (it *ProjectRewardRewardAmountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectRewardRewardAmountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectRewardRewardAmountSet represents a RewardAmountSet event raised by the ProjectReward contract.
type ProjectRewardRewardAmountSet struct {
	Owner  common.Address
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAmountSet is a free log retrieval operation binding the contract event 0x6b81fd16c9c40197030f5597ecf0a7ef2f1003a250b4c96dd30b1ac71562191d.
//
// Solidity: event RewardAmountSet(address indexed owner, uint256 indexed id, uint256 amount)
func (_ProjectReward *ProjectRewardFilterer) FilterRewardAmountSet(opts *bind.FilterOpts, owner []common.Address, id []*big.Int) (*ProjectRewardRewardAmountSetIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ProjectReward.contract.FilterLogs(opts, "RewardAmountSet", ownerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ProjectRewardRewardAmountSetIterator{contract: _ProjectReward.contract, event: "RewardAmountSet", logs: logs, sub: sub}, nil
}

// WatchRewardAmountSet is a free log subscription operation binding the contract event 0x6b81fd16c9c40197030f5597ecf0a7ef2f1003a250b4c96dd30b1ac71562191d.
//
// Solidity: event RewardAmountSet(address indexed owner, uint256 indexed id, uint256 amount)
func (_ProjectReward *ProjectRewardFilterer) WatchRewardAmountSet(opts *bind.WatchOpts, sink chan<- *ProjectRewardRewardAmountSet, owner []common.Address, id []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ProjectReward.contract.WatchLogs(opts, "RewardAmountSet", ownerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectRewardRewardAmountSet)
				if err := _ProjectReward.contract.UnpackLog(event, "RewardAmountSet", log); err != nil {
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

// ParseRewardAmountSet is a log parse operation binding the contract event 0x6b81fd16c9c40197030f5597ecf0a7ef2f1003a250b4c96dd30b1ac71562191d.
//
// Solidity: event RewardAmountSet(address indexed owner, uint256 indexed id, uint256 amount)
func (_ProjectReward *ProjectRewardFilterer) ParseRewardAmountSet(log types.Log) (*ProjectRewardRewardAmountSet, error) {
	event := new(ProjectRewardRewardAmountSet)
	if err := _ProjectReward.contract.UnpackLog(event, "RewardAmountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectRewardRewardTokenSetIterator is returned from FilterRewardTokenSet and is used to iterate over the raw logs and unpacked data for RewardTokenSet events raised by the ProjectReward contract.
type ProjectRewardRewardTokenSetIterator struct {
	Event *ProjectRewardRewardTokenSet // Event containing the contract specifics and raw log

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
func (it *ProjectRewardRewardTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectRewardRewardTokenSet)
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
		it.Event = new(ProjectRewardRewardTokenSet)
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
func (it *ProjectRewardRewardTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectRewardRewardTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectRewardRewardTokenSet represents a RewardTokenSet event raised by the ProjectReward contract.
type ProjectRewardRewardTokenSet struct {
	Id          *big.Int
	RewardToken common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardTokenSet is a free log retrieval operation binding the contract event 0xd0c0770d6fa3022a42b8328f70c23439b2ac9eae5c54cd50e3da1bdaa600bdf5.
//
// Solidity: event RewardTokenSet(uint256 indexed id, address indexed rewardToken)
func (_ProjectReward *ProjectRewardFilterer) FilterRewardTokenSet(opts *bind.FilterOpts, id []*big.Int, rewardToken []common.Address) (*ProjectRewardRewardTokenSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _ProjectReward.contract.FilterLogs(opts, "RewardTokenSet", idRule, rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &ProjectRewardRewardTokenSetIterator{contract: _ProjectReward.contract, event: "RewardTokenSet", logs: logs, sub: sub}, nil
}

// WatchRewardTokenSet is a free log subscription operation binding the contract event 0xd0c0770d6fa3022a42b8328f70c23439b2ac9eae5c54cd50e3da1bdaa600bdf5.
//
// Solidity: event RewardTokenSet(uint256 indexed id, address indexed rewardToken)
func (_ProjectReward *ProjectRewardFilterer) WatchRewardTokenSet(opts *bind.WatchOpts, sink chan<- *ProjectRewardRewardTokenSet, id []*big.Int, rewardToken []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _ProjectReward.contract.WatchLogs(opts, "RewardTokenSet", idRule, rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectRewardRewardTokenSet)
				if err := _ProjectReward.contract.UnpackLog(event, "RewardTokenSet", log); err != nil {
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

// ParseRewardTokenSet is a log parse operation binding the contract event 0xd0c0770d6fa3022a42b8328f70c23439b2ac9eae5c54cd50e3da1bdaa600bdf5.
//
// Solidity: event RewardTokenSet(uint256 indexed id, address indexed rewardToken)
func (_ProjectReward *ProjectRewardFilterer) ParseRewardTokenSet(log types.Log) (*ProjectRewardRewardTokenSet, error) {
	event := new(ProjectRewardRewardTokenSet)
	if err := _ProjectReward.contract.UnpackLog(event, "RewardTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
