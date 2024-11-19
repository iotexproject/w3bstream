// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockioidregistry

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

// MockioidregistryMetaData contains all meta data concerning the Mockioidregistry contract.
var MockioidregistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ioid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_device\",\"type\":\"address\"}],\"name\":\"bind\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_device\",\"type\":\"address\"}],\"name\":\"deviceTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ioid\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ioID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ioid\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MockioidregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use MockioidregistryMetaData.ABI instead.
var MockioidregistryABI = MockioidregistryMetaData.ABI

// Mockioidregistry is an auto generated Go binding around an Ethereum contract.
type Mockioidregistry struct {
	MockioidregistryCaller     // Read-only binding to the contract
	MockioidregistryTransactor // Write-only binding to the contract
	MockioidregistryFilterer   // Log filterer for contract events
}

// MockioidregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockioidregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockioidregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockioidregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockioidregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockioidregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockioidregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockioidregistrySession struct {
	Contract     *Mockioidregistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockioidregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockioidregistryCallerSession struct {
	Contract *MockioidregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MockioidregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockioidregistryTransactorSession struct {
	Contract     *MockioidregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MockioidregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockioidregistryRaw struct {
	Contract *Mockioidregistry // Generic contract binding to access the raw methods on
}

// MockioidregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockioidregistryCallerRaw struct {
	Contract *MockioidregistryCaller // Generic read-only contract binding to access the raw methods on
}

// MockioidregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockioidregistryTransactorRaw struct {
	Contract *MockioidregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockioidregistry creates a new instance of Mockioidregistry, bound to a specific deployed contract.
func NewMockioidregistry(address common.Address, backend bind.ContractBackend) (*Mockioidregistry, error) {
	contract, err := bindMockioidregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mockioidregistry{MockioidregistryCaller: MockioidregistryCaller{contract: contract}, MockioidregistryTransactor: MockioidregistryTransactor{contract: contract}, MockioidregistryFilterer: MockioidregistryFilterer{contract: contract}}, nil
}

// NewMockioidregistryCaller creates a new read-only instance of Mockioidregistry, bound to a specific deployed contract.
func NewMockioidregistryCaller(address common.Address, caller bind.ContractCaller) (*MockioidregistryCaller, error) {
	contract, err := bindMockioidregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockioidregistryCaller{contract: contract}, nil
}

// NewMockioidregistryTransactor creates a new write-only instance of Mockioidregistry, bound to a specific deployed contract.
func NewMockioidregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MockioidregistryTransactor, error) {
	contract, err := bindMockioidregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockioidregistryTransactor{contract: contract}, nil
}

// NewMockioidregistryFilterer creates a new log filterer instance of Mockioidregistry, bound to a specific deployed contract.
func NewMockioidregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MockioidregistryFilterer, error) {
	contract, err := bindMockioidregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockioidregistryFilterer{contract: contract}, nil
}

// bindMockioidregistry binds a generic wrapper to an already deployed contract.
func bindMockioidregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockioidregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mockioidregistry *MockioidregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mockioidregistry.Contract.MockioidregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mockioidregistry *MockioidregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mockioidregistry.Contract.MockioidregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mockioidregistry *MockioidregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mockioidregistry.Contract.MockioidregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mockioidregistry *MockioidregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mockioidregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mockioidregistry *MockioidregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mockioidregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mockioidregistry *MockioidregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mockioidregistry.Contract.contract.Transact(opts, method, params...)
}

// DeviceTokenId is a free data retrieval call binding the contract method 0x4965c7e1.
//
// Solidity: function deviceTokenId(address _device) view returns(uint256)
func (_Mockioidregistry *MockioidregistryCaller) DeviceTokenId(opts *bind.CallOpts, _device common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mockioidregistry.contract.Call(opts, &out, "deviceTokenId", _device)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeviceTokenId is a free data retrieval call binding the contract method 0x4965c7e1.
//
// Solidity: function deviceTokenId(address _device) view returns(uint256)
func (_Mockioidregistry *MockioidregistrySession) DeviceTokenId(_device common.Address) (*big.Int, error) {
	return _Mockioidregistry.Contract.DeviceTokenId(&_Mockioidregistry.CallOpts, _device)
}

// DeviceTokenId is a free data retrieval call binding the contract method 0x4965c7e1.
//
// Solidity: function deviceTokenId(address _device) view returns(uint256)
func (_Mockioidregistry *MockioidregistryCallerSession) DeviceTokenId(_device common.Address) (*big.Int, error) {
	return _Mockioidregistry.Contract.DeviceTokenId(&_Mockioidregistry.CallOpts, _device)
}

// IoID is a free data retrieval call binding the contract method 0xc3b3135e.
//
// Solidity: function ioID() view returns(address)
func (_Mockioidregistry *MockioidregistryCaller) IoID(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mockioidregistry.contract.Call(opts, &out, "ioID")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IoID is a free data retrieval call binding the contract method 0xc3b3135e.
//
// Solidity: function ioID() view returns(address)
func (_Mockioidregistry *MockioidregistrySession) IoID() (common.Address, error) {
	return _Mockioidregistry.Contract.IoID(&_Mockioidregistry.CallOpts)
}

// IoID is a free data retrieval call binding the contract method 0xc3b3135e.
//
// Solidity: function ioID() view returns(address)
func (_Mockioidregistry *MockioidregistryCallerSession) IoID() (common.Address, error) {
	return _Mockioidregistry.Contract.IoID(&_Mockioidregistry.CallOpts)
}

// Ioid is a free data retrieval call binding the contract method 0x50230d20.
//
// Solidity: function ioid() view returns(address)
func (_Mockioidregistry *MockioidregistryCaller) Ioid(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mockioidregistry.contract.Call(opts, &out, "ioid")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ioid is a free data retrieval call binding the contract method 0x50230d20.
//
// Solidity: function ioid() view returns(address)
func (_Mockioidregistry *MockioidregistrySession) Ioid() (common.Address, error) {
	return _Mockioidregistry.Contract.Ioid(&_Mockioidregistry.CallOpts)
}

// Ioid is a free data retrieval call binding the contract method 0x50230d20.
//
// Solidity: function ioid() view returns(address)
func (_Mockioidregistry *MockioidregistryCallerSession) Ioid() (common.Address, error) {
	return _Mockioidregistry.Contract.Ioid(&_Mockioidregistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mockioidregistry *MockioidregistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mockioidregistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mockioidregistry *MockioidregistrySession) Owner() (common.Address, error) {
	return _Mockioidregistry.Contract.Owner(&_Mockioidregistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mockioidregistry *MockioidregistryCallerSession) Owner() (common.Address, error) {
	return _Mockioidregistry.Contract.Owner(&_Mockioidregistry.CallOpts)
}

// Bind is a paid mutator transaction binding the contract method 0x09088532.
//
// Solidity: function bind(uint256 _ioid, address _device) returns()
func (_Mockioidregistry *MockioidregistryTransactor) Bind(opts *bind.TransactOpts, _ioid *big.Int, _device common.Address) (*types.Transaction, error) {
	return _Mockioidregistry.contract.Transact(opts, "bind", _ioid, _device)
}

// Bind is a paid mutator transaction binding the contract method 0x09088532.
//
// Solidity: function bind(uint256 _ioid, address _device) returns()
func (_Mockioidregistry *MockioidregistrySession) Bind(_ioid *big.Int, _device common.Address) (*types.Transaction, error) {
	return _Mockioidregistry.Contract.Bind(&_Mockioidregistry.TransactOpts, _ioid, _device)
}

// Bind is a paid mutator transaction binding the contract method 0x09088532.
//
// Solidity: function bind(uint256 _ioid, address _device) returns()
func (_Mockioidregistry *MockioidregistryTransactorSession) Bind(_ioid *big.Int, _device common.Address) (*types.Transaction, error) {
	return _Mockioidregistry.Contract.Bind(&_Mockioidregistry.TransactOpts, _ioid, _device)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _ioid) returns()
func (_Mockioidregistry *MockioidregistryTransactor) Initialize(opts *bind.TransactOpts, _ioid common.Address) (*types.Transaction, error) {
	return _Mockioidregistry.contract.Transact(opts, "initialize", _ioid)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _ioid) returns()
func (_Mockioidregistry *MockioidregistrySession) Initialize(_ioid common.Address) (*types.Transaction, error) {
	return _Mockioidregistry.Contract.Initialize(&_Mockioidregistry.TransactOpts, _ioid)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _ioid) returns()
func (_Mockioidregistry *MockioidregistryTransactorSession) Initialize(_ioid common.Address) (*types.Transaction, error) {
	return _Mockioidregistry.Contract.Initialize(&_Mockioidregistry.TransactOpts, _ioid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mockioidregistry *MockioidregistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mockioidregistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mockioidregistry *MockioidregistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Mockioidregistry.Contract.RenounceOwnership(&_Mockioidregistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mockioidregistry *MockioidregistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mockioidregistry.Contract.RenounceOwnership(&_Mockioidregistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mockioidregistry *MockioidregistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mockioidregistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mockioidregistry *MockioidregistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mockioidregistry.Contract.TransferOwnership(&_Mockioidregistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mockioidregistry *MockioidregistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mockioidregistry.Contract.TransferOwnership(&_Mockioidregistry.TransactOpts, newOwner)
}

// MockioidregistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Mockioidregistry contract.
type MockioidregistryInitializedIterator struct {
	Event *MockioidregistryInitialized // Event containing the contract specifics and raw log

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
func (it *MockioidregistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockioidregistryInitialized)
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
		it.Event = new(MockioidregistryInitialized)
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
func (it *MockioidregistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockioidregistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockioidregistryInitialized represents a Initialized event raised by the Mockioidregistry contract.
type MockioidregistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mockioidregistry *MockioidregistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*MockioidregistryInitializedIterator, error) {

	logs, sub, err := _Mockioidregistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MockioidregistryInitializedIterator{contract: _Mockioidregistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mockioidregistry *MockioidregistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MockioidregistryInitialized) (event.Subscription, error) {

	logs, sub, err := _Mockioidregistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockioidregistryInitialized)
				if err := _Mockioidregistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Mockioidregistry *MockioidregistryFilterer) ParseInitialized(log types.Log) (*MockioidregistryInitialized, error) {
	event := new(MockioidregistryInitialized)
	if err := _Mockioidregistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockioidregistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mockioidregistry contract.
type MockioidregistryOwnershipTransferredIterator struct {
	Event *MockioidregistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MockioidregistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockioidregistryOwnershipTransferred)
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
		it.Event = new(MockioidregistryOwnershipTransferred)
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
func (it *MockioidregistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockioidregistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockioidregistryOwnershipTransferred represents a OwnershipTransferred event raised by the Mockioidregistry contract.
type MockioidregistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mockioidregistry *MockioidregistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MockioidregistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mockioidregistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MockioidregistryOwnershipTransferredIterator{contract: _Mockioidregistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mockioidregistry *MockioidregistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MockioidregistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mockioidregistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockioidregistryOwnershipTransferred)
				if err := _Mockioidregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mockioidregistry *MockioidregistryFilterer) ParseOwnershipTransferred(log types.Log) (*MockioidregistryOwnershipTransferred, error) {
	event := new(MockioidregistryOwnershipTransferred)
	if err := _Mockioidregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
