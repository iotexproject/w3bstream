// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package projectdevice

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

// ProjectdeviceMetaData contains all meta data concerning the Projectdevice contract.
var ProjectdeviceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"Approve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"Unapprove\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_devices\",\"type\":\"address[]\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_device\",\"type\":\"address\"}],\"name\":\"approved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ioIDRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_w3bstreamProject\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ioIDRegistry\",\"outputs\":[{\"internalType\":\"contractIioIDRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_devices\",\"type\":\"address[]\"}],\"name\":\"unapprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"w3bstreamProject\",\"outputs\":[{\"internalType\":\"contractIW3bstreamProject\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ProjectdeviceABI is the input ABI used to generate the binding from.
// Deprecated: Use ProjectdeviceMetaData.ABI instead.
var ProjectdeviceABI = ProjectdeviceMetaData.ABI

// Projectdevice is an auto generated Go binding around an Ethereum contract.
type Projectdevice struct {
	ProjectdeviceCaller     // Read-only binding to the contract
	ProjectdeviceTransactor // Write-only binding to the contract
	ProjectdeviceFilterer   // Log filterer for contract events
}

// ProjectdeviceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProjectdeviceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectdeviceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProjectdeviceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectdeviceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProjectdeviceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectdeviceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProjectdeviceSession struct {
	Contract     *Projectdevice    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProjectdeviceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProjectdeviceCallerSession struct {
	Contract *ProjectdeviceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ProjectdeviceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProjectdeviceTransactorSession struct {
	Contract     *ProjectdeviceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ProjectdeviceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProjectdeviceRaw struct {
	Contract *Projectdevice // Generic contract binding to access the raw methods on
}

// ProjectdeviceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProjectdeviceCallerRaw struct {
	Contract *ProjectdeviceCaller // Generic read-only contract binding to access the raw methods on
}

// ProjectdeviceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProjectdeviceTransactorRaw struct {
	Contract *ProjectdeviceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProjectdevice creates a new instance of Projectdevice, bound to a specific deployed contract.
func NewProjectdevice(address common.Address, backend bind.ContractBackend) (*Projectdevice, error) {
	contract, err := bindProjectdevice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Projectdevice{ProjectdeviceCaller: ProjectdeviceCaller{contract: contract}, ProjectdeviceTransactor: ProjectdeviceTransactor{contract: contract}, ProjectdeviceFilterer: ProjectdeviceFilterer{contract: contract}}, nil
}

// NewProjectdeviceCaller creates a new read-only instance of Projectdevice, bound to a specific deployed contract.
func NewProjectdeviceCaller(address common.Address, caller bind.ContractCaller) (*ProjectdeviceCaller, error) {
	contract, err := bindProjectdevice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectdeviceCaller{contract: contract}, nil
}

// NewProjectdeviceTransactor creates a new write-only instance of Projectdevice, bound to a specific deployed contract.
func NewProjectdeviceTransactor(address common.Address, transactor bind.ContractTransactor) (*ProjectdeviceTransactor, error) {
	contract, err := bindProjectdevice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectdeviceTransactor{contract: contract}, nil
}

// NewProjectdeviceFilterer creates a new log filterer instance of Projectdevice, bound to a specific deployed contract.
func NewProjectdeviceFilterer(address common.Address, filterer bind.ContractFilterer) (*ProjectdeviceFilterer, error) {
	contract, err := bindProjectdevice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProjectdeviceFilterer{contract: contract}, nil
}

// bindProjectdevice binds a generic wrapper to an already deployed contract.
func bindProjectdevice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProjectdeviceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Projectdevice *ProjectdeviceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Projectdevice.Contract.ProjectdeviceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Projectdevice *ProjectdeviceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Projectdevice.Contract.ProjectdeviceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Projectdevice *ProjectdeviceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Projectdevice.Contract.ProjectdeviceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Projectdevice *ProjectdeviceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Projectdevice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Projectdevice *ProjectdeviceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Projectdevice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Projectdevice *ProjectdeviceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Projectdevice.Contract.contract.Transact(opts, method, params...)
}

// Approved is a free data retrieval call binding the contract method 0x8253951a.
//
// Solidity: function approved(uint256 _projectId, address _device) view returns(bool)
func (_Projectdevice *ProjectdeviceCaller) Approved(opts *bind.CallOpts, _projectId *big.Int, _device common.Address) (bool, error) {
	var out []interface{}
	err := _Projectdevice.contract.Call(opts, &out, "approved", _projectId, _device)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approved is a free data retrieval call binding the contract method 0x8253951a.
//
// Solidity: function approved(uint256 _projectId, address _device) view returns(bool)
func (_Projectdevice *ProjectdeviceSession) Approved(_projectId *big.Int, _device common.Address) (bool, error) {
	return _Projectdevice.Contract.Approved(&_Projectdevice.CallOpts, _projectId, _device)
}

// Approved is a free data retrieval call binding the contract method 0x8253951a.
//
// Solidity: function approved(uint256 _projectId, address _device) view returns(bool)
func (_Projectdevice *ProjectdeviceCallerSession) Approved(_projectId *big.Int, _device common.Address) (bool, error) {
	return _Projectdevice.Contract.Approved(&_Projectdevice.CallOpts, _projectId, _device)
}

// IoIDRegistry is a free data retrieval call binding the contract method 0x95cc7086.
//
// Solidity: function ioIDRegistry() view returns(address)
func (_Projectdevice *ProjectdeviceCaller) IoIDRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Projectdevice.contract.Call(opts, &out, "ioIDRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IoIDRegistry is a free data retrieval call binding the contract method 0x95cc7086.
//
// Solidity: function ioIDRegistry() view returns(address)
func (_Projectdevice *ProjectdeviceSession) IoIDRegistry() (common.Address, error) {
	return _Projectdevice.Contract.IoIDRegistry(&_Projectdevice.CallOpts)
}

// IoIDRegistry is a free data retrieval call binding the contract method 0x95cc7086.
//
// Solidity: function ioIDRegistry() view returns(address)
func (_Projectdevice *ProjectdeviceCallerSession) IoIDRegistry() (common.Address, error) {
	return _Projectdevice.Contract.IoIDRegistry(&_Projectdevice.CallOpts)
}

// W3bstreamProject is a free data retrieval call binding the contract method 0x561f1fc4.
//
// Solidity: function w3bstreamProject() view returns(address)
func (_Projectdevice *ProjectdeviceCaller) W3bstreamProject(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Projectdevice.contract.Call(opts, &out, "w3bstreamProject")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// W3bstreamProject is a free data retrieval call binding the contract method 0x561f1fc4.
//
// Solidity: function w3bstreamProject() view returns(address)
func (_Projectdevice *ProjectdeviceSession) W3bstreamProject() (common.Address, error) {
	return _Projectdevice.Contract.W3bstreamProject(&_Projectdevice.CallOpts)
}

// W3bstreamProject is a free data retrieval call binding the contract method 0x561f1fc4.
//
// Solidity: function w3bstreamProject() view returns(address)
func (_Projectdevice *ProjectdeviceCallerSession) W3bstreamProject() (common.Address, error) {
	return _Projectdevice.Contract.W3bstreamProject(&_Projectdevice.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xd744d8dc.
//
// Solidity: function approve(uint256 _projectId, address[] _devices) returns()
func (_Projectdevice *ProjectdeviceTransactor) Approve(opts *bind.TransactOpts, _projectId *big.Int, _devices []common.Address) (*types.Transaction, error) {
	return _Projectdevice.contract.Transact(opts, "approve", _projectId, _devices)
}

// Approve is a paid mutator transaction binding the contract method 0xd744d8dc.
//
// Solidity: function approve(uint256 _projectId, address[] _devices) returns()
func (_Projectdevice *ProjectdeviceSession) Approve(_projectId *big.Int, _devices []common.Address) (*types.Transaction, error) {
	return _Projectdevice.Contract.Approve(&_Projectdevice.TransactOpts, _projectId, _devices)
}

// Approve is a paid mutator transaction binding the contract method 0xd744d8dc.
//
// Solidity: function approve(uint256 _projectId, address[] _devices) returns()
func (_Projectdevice *ProjectdeviceTransactorSession) Approve(_projectId *big.Int, _devices []common.Address) (*types.Transaction, error) {
	return _Projectdevice.Contract.Approve(&_Projectdevice.TransactOpts, _projectId, _devices)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ioIDRegistry, address _w3bstreamProject) returns()
func (_Projectdevice *ProjectdeviceTransactor) Initialize(opts *bind.TransactOpts, _ioIDRegistry common.Address, _w3bstreamProject common.Address) (*types.Transaction, error) {
	return _Projectdevice.contract.Transact(opts, "initialize", _ioIDRegistry, _w3bstreamProject)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ioIDRegistry, address _w3bstreamProject) returns()
func (_Projectdevice *ProjectdeviceSession) Initialize(_ioIDRegistry common.Address, _w3bstreamProject common.Address) (*types.Transaction, error) {
	return _Projectdevice.Contract.Initialize(&_Projectdevice.TransactOpts, _ioIDRegistry, _w3bstreamProject)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ioIDRegistry, address _w3bstreamProject) returns()
func (_Projectdevice *ProjectdeviceTransactorSession) Initialize(_ioIDRegistry common.Address, _w3bstreamProject common.Address) (*types.Transaction, error) {
	return _Projectdevice.Contract.Initialize(&_Projectdevice.TransactOpts, _ioIDRegistry, _w3bstreamProject)
}

// Unapprove is a paid mutator transaction binding the contract method 0xfa6b352c.
//
// Solidity: function unapprove(uint256 _projectId, address[] _devices) returns()
func (_Projectdevice *ProjectdeviceTransactor) Unapprove(opts *bind.TransactOpts, _projectId *big.Int, _devices []common.Address) (*types.Transaction, error) {
	return _Projectdevice.contract.Transact(opts, "unapprove", _projectId, _devices)
}

// Unapprove is a paid mutator transaction binding the contract method 0xfa6b352c.
//
// Solidity: function unapprove(uint256 _projectId, address[] _devices) returns()
func (_Projectdevice *ProjectdeviceSession) Unapprove(_projectId *big.Int, _devices []common.Address) (*types.Transaction, error) {
	return _Projectdevice.Contract.Unapprove(&_Projectdevice.TransactOpts, _projectId, _devices)
}

// Unapprove is a paid mutator transaction binding the contract method 0xfa6b352c.
//
// Solidity: function unapprove(uint256 _projectId, address[] _devices) returns()
func (_Projectdevice *ProjectdeviceTransactorSession) Unapprove(_projectId *big.Int, _devices []common.Address) (*types.Transaction, error) {
	return _Projectdevice.Contract.Unapprove(&_Projectdevice.TransactOpts, _projectId, _devices)
}

// ProjectdeviceApproveIterator is returned from FilterApprove and is used to iterate over the raw logs and unpacked data for Approve events raised by the Projectdevice contract.
type ProjectdeviceApproveIterator struct {
	Event *ProjectdeviceApprove // Event containing the contract specifics and raw log

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
func (it *ProjectdeviceApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectdeviceApprove)
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
		it.Event = new(ProjectdeviceApprove)
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
func (it *ProjectdeviceApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectdeviceApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectdeviceApprove represents a Approve event raised by the Projectdevice contract.
type ProjectdeviceApprove struct {
	ProjectId *big.Int
	Device    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterApprove is a free log retrieval operation binding the contract event 0x47ad3e4b0f4bdbe3ac08708bbc45053d2ff616911b39e65563fd9bd781909645.
//
// Solidity: event Approve(uint256 projectId, address indexed device)
func (_Projectdevice *ProjectdeviceFilterer) FilterApprove(opts *bind.FilterOpts, device []common.Address) (*ProjectdeviceApproveIterator, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _Projectdevice.contract.FilterLogs(opts, "Approve", deviceRule)
	if err != nil {
		return nil, err
	}
	return &ProjectdeviceApproveIterator{contract: _Projectdevice.contract, event: "Approve", logs: logs, sub: sub}, nil
}

// WatchApprove is a free log subscription operation binding the contract event 0x47ad3e4b0f4bdbe3ac08708bbc45053d2ff616911b39e65563fd9bd781909645.
//
// Solidity: event Approve(uint256 projectId, address indexed device)
func (_Projectdevice *ProjectdeviceFilterer) WatchApprove(opts *bind.WatchOpts, sink chan<- *ProjectdeviceApprove, device []common.Address) (event.Subscription, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _Projectdevice.contract.WatchLogs(opts, "Approve", deviceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectdeviceApprove)
				if err := _Projectdevice.contract.UnpackLog(event, "Approve", log); err != nil {
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
func (_Projectdevice *ProjectdeviceFilterer) ParseApprove(log types.Log) (*ProjectdeviceApprove, error) {
	event := new(ProjectdeviceApprove)
	if err := _Projectdevice.contract.UnpackLog(event, "Approve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectdeviceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Projectdevice contract.
type ProjectdeviceInitializedIterator struct {
	Event *ProjectdeviceInitialized // Event containing the contract specifics and raw log

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
func (it *ProjectdeviceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectdeviceInitialized)
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
		it.Event = new(ProjectdeviceInitialized)
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
func (it *ProjectdeviceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectdeviceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectdeviceInitialized represents a Initialized event raised by the Projectdevice contract.
type ProjectdeviceInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Projectdevice *ProjectdeviceFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProjectdeviceInitializedIterator, error) {

	logs, sub, err := _Projectdevice.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProjectdeviceInitializedIterator{contract: _Projectdevice.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Projectdevice *ProjectdeviceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProjectdeviceInitialized) (event.Subscription, error) {

	logs, sub, err := _Projectdevice.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectdeviceInitialized)
				if err := _Projectdevice.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Projectdevice *ProjectdeviceFilterer) ParseInitialized(log types.Log) (*ProjectdeviceInitialized, error) {
	event := new(ProjectdeviceInitialized)
	if err := _Projectdevice.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectdeviceUnapproveIterator is returned from FilterUnapprove and is used to iterate over the raw logs and unpacked data for Unapprove events raised by the Projectdevice contract.
type ProjectdeviceUnapproveIterator struct {
	Event *ProjectdeviceUnapprove // Event containing the contract specifics and raw log

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
func (it *ProjectdeviceUnapproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectdeviceUnapprove)
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
		it.Event = new(ProjectdeviceUnapprove)
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
func (it *ProjectdeviceUnapproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectdeviceUnapproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectdeviceUnapprove represents a Unapprove event raised by the Projectdevice contract.
type ProjectdeviceUnapprove struct {
	ProjectId *big.Int
	Device    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnapprove is a free log retrieval operation binding the contract event 0x958e4c39d74f2f6ff958f2ff19d4cf29448476d9e44e9aab9547f7e30d4cc86b.
//
// Solidity: event Unapprove(uint256 projectId, address indexed device)
func (_Projectdevice *ProjectdeviceFilterer) FilterUnapprove(opts *bind.FilterOpts, device []common.Address) (*ProjectdeviceUnapproveIterator, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _Projectdevice.contract.FilterLogs(opts, "Unapprove", deviceRule)
	if err != nil {
		return nil, err
	}
	return &ProjectdeviceUnapproveIterator{contract: _Projectdevice.contract, event: "Unapprove", logs: logs, sub: sub}, nil
}

// WatchUnapprove is a free log subscription operation binding the contract event 0x958e4c39d74f2f6ff958f2ff19d4cf29448476d9e44e9aab9547f7e30d4cc86b.
//
// Solidity: event Unapprove(uint256 projectId, address indexed device)
func (_Projectdevice *ProjectdeviceFilterer) WatchUnapprove(opts *bind.WatchOpts, sink chan<- *ProjectdeviceUnapprove, device []common.Address) (event.Subscription, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _Projectdevice.contract.WatchLogs(opts, "Unapprove", deviceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectdeviceUnapprove)
				if err := _Projectdevice.contract.UnpackLog(event, "Unapprove", log); err != nil {
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
func (_Projectdevice *ProjectdeviceFilterer) ParseUnapprove(log types.Log) (*ProjectdeviceUnapprove, error) {
	event := new(ProjectdeviceUnapprove)
	if err := _Projectdevice.contract.UnpackLog(event, "Unapprove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
