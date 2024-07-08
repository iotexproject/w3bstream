// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package projectclient

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

// ProjectclientMetaData contains all meta data concerning the Projectclient contract.
var ProjectclientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"Approve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"Unapprove\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_devices\",\"type\":\"address[]\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_device\",\"type\":\"address\"}],\"name\":\"approved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ioIDRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_w3bstreamProject\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ioIDRegistry\",\"outputs\":[{\"internalType\":\"contractIioIDRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_devices\",\"type\":\"address[]\"}],\"name\":\"unapprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"w3bstreamProject\",\"outputs\":[{\"internalType\":\"contractIW3bstreamProject\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ProjectclientABI is the input ABI used to generate the binding from.
// Deprecated: Use ProjectclientMetaData.ABI instead.
var ProjectclientABI = ProjectclientMetaData.ABI

// Projectclient is an auto generated Go binding around an Ethereum contract.
type Projectclient struct {
	ProjectclientCaller     // Read-only binding to the contract
	ProjectclientTransactor // Write-only binding to the contract
	ProjectclientFilterer   // Log filterer for contract events
}

// ProjectclientCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProjectclientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectclientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProjectclientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectclientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProjectclientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectclientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProjectclientSession struct {
	Contract     *Projectclient    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProjectclientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProjectclientCallerSession struct {
	Contract *ProjectclientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ProjectclientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProjectclientTransactorSession struct {
	Contract     *ProjectclientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ProjectclientRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProjectclientRaw struct {
	Contract *Projectclient // Generic contract binding to access the raw methods on
}

// ProjectclientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProjectclientCallerRaw struct {
	Contract *ProjectclientCaller // Generic read-only contract binding to access the raw methods on
}

// ProjectclientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProjectclientTransactorRaw struct {
	Contract *ProjectclientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProjectclient creates a new instance of Projectclient, bound to a specific deployed contract.
func NewProjectclient(address common.Address, backend bind.ContractBackend) (*Projectclient, error) {
	contract, err := bindProjectclient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Projectclient{ProjectclientCaller: ProjectclientCaller{contract: contract}, ProjectclientTransactor: ProjectclientTransactor{contract: contract}, ProjectclientFilterer: ProjectclientFilterer{contract: contract}}, nil
}

// NewProjectclientCaller creates a new read-only instance of Projectclient, bound to a specific deployed contract.
func NewProjectclientCaller(address common.Address, caller bind.ContractCaller) (*ProjectclientCaller, error) {
	contract, err := bindProjectclient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectclientCaller{contract: contract}, nil
}

// NewProjectclientTransactor creates a new write-only instance of Projectclient, bound to a specific deployed contract.
func NewProjectclientTransactor(address common.Address, transactor bind.ContractTransactor) (*ProjectclientTransactor, error) {
	contract, err := bindProjectclient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectclientTransactor{contract: contract}, nil
}

// NewProjectclientFilterer creates a new log filterer instance of Projectclient, bound to a specific deployed contract.
func NewProjectclientFilterer(address common.Address, filterer bind.ContractFilterer) (*ProjectclientFilterer, error) {
	contract, err := bindProjectclient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProjectclientFilterer{contract: contract}, nil
}

// bindProjectclient binds a generic wrapper to an already deployed contract.
func bindProjectclient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProjectclientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Projectclient *ProjectclientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Projectclient.Contract.ProjectclientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Projectclient *ProjectclientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Projectclient.Contract.ProjectclientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Projectclient *ProjectclientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Projectclient.Contract.ProjectclientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Projectclient *ProjectclientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Projectclient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Projectclient *ProjectclientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Projectclient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Projectclient *ProjectclientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Projectclient.Contract.contract.Transact(opts, method, params...)
}

// Approved is a free data retrieval call binding the contract method 0x8253951a.
//
// Solidity: function approved(uint256 _projectId, address _device) view returns(bool)
func (_Projectclient *ProjectclientCaller) Approved(opts *bind.CallOpts, _projectId *big.Int, _device common.Address) (bool, error) {
	var out []interface{}
	err := _Projectclient.contract.Call(opts, &out, "approved", _projectId, _device)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approved is a free data retrieval call binding the contract method 0x8253951a.
//
// Solidity: function approved(uint256 _projectId, address _device) view returns(bool)
func (_Projectclient *ProjectclientSession) Approved(_projectId *big.Int, _device common.Address) (bool, error) {
	return _Projectclient.Contract.Approved(&_Projectclient.CallOpts, _projectId, _device)
}

// Approved is a free data retrieval call binding the contract method 0x8253951a.
//
// Solidity: function approved(uint256 _projectId, address _device) view returns(bool)
func (_Projectclient *ProjectclientCallerSession) Approved(_projectId *big.Int, _device common.Address) (bool, error) {
	return _Projectclient.Contract.Approved(&_Projectclient.CallOpts, _projectId, _device)
}

// IoIDRegistry is a free data retrieval call binding the contract method 0x95cc7086.
//
// Solidity: function ioIDRegistry() view returns(address)
func (_Projectclient *ProjectclientCaller) IoIDRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Projectclient.contract.Call(opts, &out, "ioIDRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IoIDRegistry is a free data retrieval call binding the contract method 0x95cc7086.
//
// Solidity: function ioIDRegistry() view returns(address)
func (_Projectclient *ProjectclientSession) IoIDRegistry() (common.Address, error) {
	return _Projectclient.Contract.IoIDRegistry(&_Projectclient.CallOpts)
}

// IoIDRegistry is a free data retrieval call binding the contract method 0x95cc7086.
//
// Solidity: function ioIDRegistry() view returns(address)
func (_Projectclient *ProjectclientCallerSession) IoIDRegistry() (common.Address, error) {
	return _Projectclient.Contract.IoIDRegistry(&_Projectclient.CallOpts)
}

// W3bstreamProject is a free data retrieval call binding the contract method 0x561f1fc4.
//
// Solidity: function w3bstreamProject() view returns(address)
func (_Projectclient *ProjectclientCaller) W3bstreamProject(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Projectclient.contract.Call(opts, &out, "w3bstreamProject")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// W3bstreamProject is a free data retrieval call binding the contract method 0x561f1fc4.
//
// Solidity: function w3bstreamProject() view returns(address)
func (_Projectclient *ProjectclientSession) W3bstreamProject() (common.Address, error) {
	return _Projectclient.Contract.W3bstreamProject(&_Projectclient.CallOpts)
}

// W3bstreamProject is a free data retrieval call binding the contract method 0x561f1fc4.
//
// Solidity: function w3bstreamProject() view returns(address)
func (_Projectclient *ProjectclientCallerSession) W3bstreamProject() (common.Address, error) {
	return _Projectclient.Contract.W3bstreamProject(&_Projectclient.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xd744d8dc.
//
// Solidity: function approve(uint256 _projectId, address[] _devices) returns()
func (_Projectclient *ProjectclientTransactor) Approve(opts *bind.TransactOpts, _projectId *big.Int, _devices []common.Address) (*types.Transaction, error) {
	return _Projectclient.contract.Transact(opts, "approve", _projectId, _devices)
}

// Approve is a paid mutator transaction binding the contract method 0xd744d8dc.
//
// Solidity: function approve(uint256 _projectId, address[] _devices) returns()
func (_Projectclient *ProjectclientSession) Approve(_projectId *big.Int, _devices []common.Address) (*types.Transaction, error) {
	return _Projectclient.Contract.Approve(&_Projectclient.TransactOpts, _projectId, _devices)
}

// Approve is a paid mutator transaction binding the contract method 0xd744d8dc.
//
// Solidity: function approve(uint256 _projectId, address[] _devices) returns()
func (_Projectclient *ProjectclientTransactorSession) Approve(_projectId *big.Int, _devices []common.Address) (*types.Transaction, error) {
	return _Projectclient.Contract.Approve(&_Projectclient.TransactOpts, _projectId, _devices)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ioIDRegistry, address _w3bstreamProject) returns()
func (_Projectclient *ProjectclientTransactor) Initialize(opts *bind.TransactOpts, _ioIDRegistry common.Address, _w3bstreamProject common.Address) (*types.Transaction, error) {
	return _Projectclient.contract.Transact(opts, "initialize", _ioIDRegistry, _w3bstreamProject)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ioIDRegistry, address _w3bstreamProject) returns()
func (_Projectclient *ProjectclientSession) Initialize(_ioIDRegistry common.Address, _w3bstreamProject common.Address) (*types.Transaction, error) {
	return _Projectclient.Contract.Initialize(&_Projectclient.TransactOpts, _ioIDRegistry, _w3bstreamProject)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ioIDRegistry, address _w3bstreamProject) returns()
func (_Projectclient *ProjectclientTransactorSession) Initialize(_ioIDRegistry common.Address, _w3bstreamProject common.Address) (*types.Transaction, error) {
	return _Projectclient.Contract.Initialize(&_Projectclient.TransactOpts, _ioIDRegistry, _w3bstreamProject)
}

// Unapprove is a paid mutator transaction binding the contract method 0xfa6b352c.
//
// Solidity: function unapprove(uint256 _projectId, address[] _devices) returns()
func (_Projectclient *ProjectclientTransactor) Unapprove(opts *bind.TransactOpts, _projectId *big.Int, _devices []common.Address) (*types.Transaction, error) {
	return _Projectclient.contract.Transact(opts, "unapprove", _projectId, _devices)
}

// Unapprove is a paid mutator transaction binding the contract method 0xfa6b352c.
//
// Solidity: function unapprove(uint256 _projectId, address[] _devices) returns()
func (_Projectclient *ProjectclientSession) Unapprove(_projectId *big.Int, _devices []common.Address) (*types.Transaction, error) {
	return _Projectclient.Contract.Unapprove(&_Projectclient.TransactOpts, _projectId, _devices)
}

// Unapprove is a paid mutator transaction binding the contract method 0xfa6b352c.
//
// Solidity: function unapprove(uint256 _projectId, address[] _devices) returns()
func (_Projectclient *ProjectclientTransactorSession) Unapprove(_projectId *big.Int, _devices []common.Address) (*types.Transaction, error) {
	return _Projectclient.Contract.Unapprove(&_Projectclient.TransactOpts, _projectId, _devices)
}

// ProjectclientApproveIterator is returned from FilterApprove and is used to iterate over the raw logs and unpacked data for Approve events raised by the Projectclient contract.
type ProjectclientApproveIterator struct {
	Event *ProjectclientApprove // Event containing the contract specifics and raw log

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
func (it *ProjectclientApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectclientApprove)
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
		it.Event = new(ProjectclientApprove)
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
func (it *ProjectclientApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectclientApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectclientApprove represents a Approve event raised by the Projectclient contract.
type ProjectclientApprove struct {
	ProjectId *big.Int
	Device    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterApprove is a free log retrieval operation binding the contract event 0x47ad3e4b0f4bdbe3ac08708bbc45053d2ff616911b39e65563fd9bd781909645.
//
// Solidity: event Approve(uint256 projectId, address indexed device)
func (_Projectclient *ProjectclientFilterer) FilterApprove(opts *bind.FilterOpts, device []common.Address) (*ProjectclientApproveIterator, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _Projectclient.contract.FilterLogs(opts, "Approve", deviceRule)
	if err != nil {
		return nil, err
	}
	return &ProjectclientApproveIterator{contract: _Projectclient.contract, event: "Approve", logs: logs, sub: sub}, nil
}

// WatchApprove is a free log subscription operation binding the contract event 0x47ad3e4b0f4bdbe3ac08708bbc45053d2ff616911b39e65563fd9bd781909645.
//
// Solidity: event Approve(uint256 projectId, address indexed device)
func (_Projectclient *ProjectclientFilterer) WatchApprove(opts *bind.WatchOpts, sink chan<- *ProjectclientApprove, device []common.Address) (event.Subscription, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _Projectclient.contract.WatchLogs(opts, "Approve", deviceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectclientApprove)
				if err := _Projectclient.contract.UnpackLog(event, "Approve", log); err != nil {
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

// ParseApprove is a log parse operation binding the contract event 0x47ad3e4b0f4bdbe3ac08708bbc45053d2ff616911b39e65563fd9bd781909645.
//
// Solidity: event Approve(uint256 projectId, address indexed device)
func (_Projectclient *ProjectclientFilterer) ParseApprove(log types.Log) (*ProjectclientApprove, error) {
	event := new(ProjectclientApprove)
	if err := _Projectclient.contract.UnpackLog(event, "Approve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectclientInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Projectclient contract.
type ProjectclientInitializedIterator struct {
	Event *ProjectclientInitialized // Event containing the contract specifics and raw log

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
func (it *ProjectclientInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectclientInitialized)
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
		it.Event = new(ProjectclientInitialized)
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
func (it *ProjectclientInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectclientInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectclientInitialized represents a Initialized event raised by the Projectclient contract.
type ProjectclientInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Projectclient *ProjectclientFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProjectclientInitializedIterator, error) {

	logs, sub, err := _Projectclient.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProjectclientInitializedIterator{contract: _Projectclient.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Projectclient *ProjectclientFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProjectclientInitialized) (event.Subscription, error) {

	logs, sub, err := _Projectclient.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectclientInitialized)
				if err := _Projectclient.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Projectclient *ProjectclientFilterer) ParseInitialized(log types.Log) (*ProjectclientInitialized, error) {
	event := new(ProjectclientInitialized)
	if err := _Projectclient.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectclientUnapproveIterator is returned from FilterUnapprove and is used to iterate over the raw logs and unpacked data for Unapprove events raised by the Projectclient contract.
type ProjectclientUnapproveIterator struct {
	Event *ProjectclientUnapprove // Event containing the contract specifics and raw log

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
func (it *ProjectclientUnapproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectclientUnapprove)
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
		it.Event = new(ProjectclientUnapprove)
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
func (it *ProjectclientUnapproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectclientUnapproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectclientUnapprove represents a Unapprove event raised by the Projectclient contract.
type ProjectclientUnapprove struct {
	ProjectId *big.Int
	Device    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnapprove is a free log retrieval operation binding the contract event 0x958e4c39d74f2f6ff958f2ff19d4cf29448476d9e44e9aab9547f7e30d4cc86b.
//
// Solidity: event Unapprove(uint256 projectId, address indexed device)
func (_Projectclient *ProjectclientFilterer) FilterUnapprove(opts *bind.FilterOpts, device []common.Address) (*ProjectclientUnapproveIterator, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _Projectclient.contract.FilterLogs(opts, "Unapprove", deviceRule)
	if err != nil {
		return nil, err
	}
	return &ProjectclientUnapproveIterator{contract: _Projectclient.contract, event: "Unapprove", logs: logs, sub: sub}, nil
}

// WatchUnapprove is a free log subscription operation binding the contract event 0x958e4c39d74f2f6ff958f2ff19d4cf29448476d9e44e9aab9547f7e30d4cc86b.
//
// Solidity: event Unapprove(uint256 projectId, address indexed device)
func (_Projectclient *ProjectclientFilterer) WatchUnapprove(opts *bind.WatchOpts, sink chan<- *ProjectclientUnapprove, device []common.Address) (event.Subscription, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _Projectclient.contract.WatchLogs(opts, "Unapprove", deviceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectclientUnapprove)
				if err := _Projectclient.contract.UnpackLog(event, "Unapprove", log); err != nil {
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

// ParseUnapprove is a log parse operation binding the contract event 0x958e4c39d74f2f6ff958f2ff19d4cf29448476d9e44e9aab9547f7e30d4cc86b.
//
// Solidity: event Unapprove(uint256 projectId, address indexed device)
func (_Projectclient *ProjectclientFilterer) ParseUnapprove(log types.Log) (*ProjectclientUnapprove, error) {
	event := new(ProjectclientUnapprove)
	if err := _Projectclient.contract.UnpackLog(event, "Unapprove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
