// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dao

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

// DaoMetaData contains all meta data concerning the Dao contract.
var DaoMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BlockAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DaoABI is the input ABI used to generate the binding from.
// Deprecated: Use DaoMetaData.ABI instead.
var DaoABI = DaoMetaData.ABI

// Dao is an auto generated Go binding around an Ethereum contract.
type Dao struct {
	DaoCaller     // Read-only binding to the contract
	DaoTransactor // Write-only binding to the contract
	DaoFilterer   // Log filterer for contract events
}

// DaoCaller is an auto generated read-only Go binding around an Ethereum contract.
type DaoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DaoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DaoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DaoSession struct {
	Contract     *Dao              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DaoCallerSession struct {
	Contract *DaoCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DaoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DaoTransactorSession struct {
	Contract     *DaoTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaoRaw is an auto generated low-level Go binding around an Ethereum contract.
type DaoRaw struct {
	Contract *Dao // Generic contract binding to access the raw methods on
}

// DaoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DaoCallerRaw struct {
	Contract *DaoCaller // Generic read-only contract binding to access the raw methods on
}

// DaoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DaoTransactorRaw struct {
	Contract *DaoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDao creates a new instance of Dao, bound to a specific deployed contract.
func NewDao(address common.Address, backend bind.ContractBackend) (*Dao, error) {
	contract, err := bindDao(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dao{DaoCaller: DaoCaller{contract: contract}, DaoTransactor: DaoTransactor{contract: contract}, DaoFilterer: DaoFilterer{contract: contract}}, nil
}

// NewDaoCaller creates a new read-only instance of Dao, bound to a specific deployed contract.
func NewDaoCaller(address common.Address, caller bind.ContractCaller) (*DaoCaller, error) {
	contract, err := bindDao(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DaoCaller{contract: contract}, nil
}

// NewDaoTransactor creates a new write-only instance of Dao, bound to a specific deployed contract.
func NewDaoTransactor(address common.Address, transactor bind.ContractTransactor) (*DaoTransactor, error) {
	contract, err := bindDao(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DaoTransactor{contract: contract}, nil
}

// NewDaoFilterer creates a new log filterer instance of Dao, bound to a specific deployed contract.
func NewDaoFilterer(address common.Address, filterer bind.ContractFilterer) (*DaoFilterer, error) {
	contract, err := bindDao(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DaoFilterer{contract: contract}, nil
}

// bindDao binds a generic wrapper to an already deployed contract.
func bindDao(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DaoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dao *DaoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dao.Contract.DaoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dao *DaoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dao.Contract.DaoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dao *DaoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dao.Contract.DaoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dao *DaoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dao.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dao *DaoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dao.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dao *DaoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dao.Contract.contract.Transact(opts, method, params...)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 ) view returns(bytes32 hash, uint256 timestamp)
func (_Dao *DaoCaller) Blocks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Hash      [32]byte
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "blocks", arg0)

	outstruct := new(struct {
		Hash      [32]byte
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 ) view returns(bytes32 hash, uint256 timestamp)
func (_Dao *DaoSession) Blocks(arg0 *big.Int) (struct {
	Hash      [32]byte
	Timestamp *big.Int
}, error) {
	return _Dao.Contract.Blocks(&_Dao.CallOpts, arg0)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 ) view returns(bytes32 hash, uint256 timestamp)
func (_Dao *DaoCallerSession) Blocks(arg0 *big.Int) (struct {
	Hash      [32]byte
	Timestamp *big.Int
}, error) {
	return _Dao.Contract.Blocks(&_Dao.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dao *DaoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dao *DaoSession) Owner() (common.Address, error) {
	return _Dao.Contract.Owner(&_Dao.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dao *DaoCallerSession) Owner() (common.Address, error) {
	return _Dao.Contract.Owner(&_Dao.CallOpts)
}

// Tip is a free data retrieval call binding the contract method 0x2755cd2d.
//
// Solidity: function tip() view returns(uint256, bytes32, uint256)
func (_Dao *DaoCaller) Tip(opts *bind.CallOpts) (*big.Int, [32]byte, *big.Int, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "tip")

	if err != nil {
		return *new(*big.Int), *new([32]byte), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// Tip is a free data retrieval call binding the contract method 0x2755cd2d.
//
// Solidity: function tip() view returns(uint256, bytes32, uint256)
func (_Dao *DaoSession) Tip() (*big.Int, [32]byte, *big.Int, error) {
	return _Dao.Contract.Tip(&_Dao.CallOpts)
}

// Tip is a free data retrieval call binding the contract method 0x2755cd2d.
//
// Solidity: function tip() view returns(uint256, bytes32, uint256)
func (_Dao *DaoCallerSession) Tip() (*big.Int, [32]byte, *big.Int, error) {
	return _Dao.Contract.Tip(&_Dao.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 genesis) returns()
func (_Dao *DaoTransactor) Initialize(opts *bind.TransactOpts, genesis [32]byte) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "initialize", genesis)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 genesis) returns()
func (_Dao *DaoSession) Initialize(genesis [32]byte) (*types.Transaction, error) {
	return _Dao.Contract.Initialize(&_Dao.TransactOpts, genesis)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 genesis) returns()
func (_Dao *DaoTransactorSession) Initialize(genesis [32]byte) (*types.Transaction, error) {
	return _Dao.Contract.Initialize(&_Dao.TransactOpts, genesis)
}

// Mint is a paid mutator transaction binding the contract method 0xe1856ff4.
//
// Solidity: function mint(bytes32 hash, uint256 timestamp) returns()
func (_Dao *DaoTransactor) Mint(opts *bind.TransactOpts, hash [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "mint", hash, timestamp)
}

// Mint is a paid mutator transaction binding the contract method 0xe1856ff4.
//
// Solidity: function mint(bytes32 hash, uint256 timestamp) returns()
func (_Dao *DaoSession) Mint(hash [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.Mint(&_Dao.TransactOpts, hash, timestamp)
}

// Mint is a paid mutator transaction binding the contract method 0xe1856ff4.
//
// Solidity: function mint(bytes32 hash, uint256 timestamp) returns()
func (_Dao *DaoTransactorSession) Mint(hash [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.Mint(&_Dao.TransactOpts, hash, timestamp)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dao *DaoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dao *DaoSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dao.Contract.RenounceOwnership(&_Dao.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dao *DaoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dao.Contract.RenounceOwnership(&_Dao.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dao *DaoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dao *DaoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dao.Contract.TransferOwnership(&_Dao.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dao *DaoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dao.Contract.TransferOwnership(&_Dao.TransactOpts, newOwner)
}

// DaoBlockAddedIterator is returned from FilterBlockAdded and is used to iterate over the raw logs and unpacked data for BlockAdded events raised by the Dao contract.
type DaoBlockAddedIterator struct {
	Event *DaoBlockAdded // Event containing the contract specifics and raw log

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
func (it *DaoBlockAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoBlockAdded)
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
		it.Event = new(DaoBlockAdded)
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
func (it *DaoBlockAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoBlockAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoBlockAdded represents a BlockAdded event raised by the Dao contract.
type DaoBlockAdded struct {
	Num       *big.Int
	Hash      [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockAdded is a free log retrieval operation binding the contract event 0x706b507f20f95b6f483ddbb8f8a800372f4fb66f49c5e9785553b8f90f592f31.
//
// Solidity: event BlockAdded(uint256 indexed num, bytes32 hash, uint256 timestamp)
func (_Dao *DaoFilterer) FilterBlockAdded(opts *bind.FilterOpts, num []*big.Int) (*DaoBlockAddedIterator, error) {

	var numRule []interface{}
	for _, numItem := range num {
		numRule = append(numRule, numItem)
	}

	logs, sub, err := _Dao.contract.FilterLogs(opts, "BlockAdded", numRule)
	if err != nil {
		return nil, err
	}
	return &DaoBlockAddedIterator{contract: _Dao.contract, event: "BlockAdded", logs: logs, sub: sub}, nil
}

// WatchBlockAdded is a free log subscription operation binding the contract event 0x706b507f20f95b6f483ddbb8f8a800372f4fb66f49c5e9785553b8f90f592f31.
//
// Solidity: event BlockAdded(uint256 indexed num, bytes32 hash, uint256 timestamp)
func (_Dao *DaoFilterer) WatchBlockAdded(opts *bind.WatchOpts, sink chan<- *DaoBlockAdded, num []*big.Int) (event.Subscription, error) {

	var numRule []interface{}
	for _, numItem := range num {
		numRule = append(numRule, numItem)
	}

	logs, sub, err := _Dao.contract.WatchLogs(opts, "BlockAdded", numRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoBlockAdded)
				if err := _Dao.contract.UnpackLog(event, "BlockAdded", log); err != nil {
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

// ParseBlockAdded is a log parse operation binding the contract event 0x706b507f20f95b6f483ddbb8f8a800372f4fb66f49c5e9785553b8f90f592f31.
//
// Solidity: event BlockAdded(uint256 indexed num, bytes32 hash, uint256 timestamp)
func (_Dao *DaoFilterer) ParseBlockAdded(log types.Log) (*DaoBlockAdded, error) {
	event := new(DaoBlockAdded)
	if err := _Dao.contract.UnpackLog(event, "BlockAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Dao contract.
type DaoInitializedIterator struct {
	Event *DaoInitialized // Event containing the contract specifics and raw log

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
func (it *DaoInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoInitialized)
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
		it.Event = new(DaoInitialized)
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
func (it *DaoInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoInitialized represents a Initialized event raised by the Dao contract.
type DaoInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dao *DaoFilterer) FilterInitialized(opts *bind.FilterOpts) (*DaoInitializedIterator, error) {

	logs, sub, err := _Dao.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DaoInitializedIterator{contract: _Dao.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dao *DaoFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DaoInitialized) (event.Subscription, error) {

	logs, sub, err := _Dao.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoInitialized)
				if err := _Dao.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Dao *DaoFilterer) ParseInitialized(log types.Log) (*DaoInitialized, error) {
	event := new(DaoInitialized)
	if err := _Dao.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dao contract.
type DaoOwnershipTransferredIterator struct {
	Event *DaoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DaoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoOwnershipTransferred)
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
		it.Event = new(DaoOwnershipTransferred)
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
func (it *DaoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoOwnershipTransferred represents a OwnershipTransferred event raised by the Dao contract.
type DaoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dao *DaoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DaoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dao.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DaoOwnershipTransferredIterator{contract: _Dao.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dao *DaoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DaoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dao.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoOwnershipTransferred)
				if err := _Dao.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dao *DaoFilterer) ParseOwnershipTransferred(log types.Log) (*DaoOwnershipTransferred, error) {
	event := new(DaoOwnershipTransferred)
	if err := _Dao.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
