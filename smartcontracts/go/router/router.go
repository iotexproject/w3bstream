// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package router

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

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"}],\"name\":\"DappBound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"DappUnbound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"router\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"DataProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dapp\",\"type\":\"address\"}],\"name\":\"bindDapp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dapp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fleetManagement\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fleetManagement\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_projectStore\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"projectStore\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_proverId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_clientId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_taskId\",\"type\":\"bytes32\"}],\"name\":\"route\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_add\",\"type\":\"address\"}],\"name\":\"setTaskManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"unbindDapp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// Dapp is a free data retrieval call binding the contract method 0x1bf8131f.
//
// Solidity: function dapp(uint256 ) view returns(address)
func (_Router *RouterCaller) Dapp(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "dapp", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dapp is a free data retrieval call binding the contract method 0x1bf8131f.
//
// Solidity: function dapp(uint256 ) view returns(address)
func (_Router *RouterSession) Dapp(arg0 *big.Int) (common.Address, error) {
	return _Router.Contract.Dapp(&_Router.CallOpts, arg0)
}

// Dapp is a free data retrieval call binding the contract method 0x1bf8131f.
//
// Solidity: function dapp(uint256 ) view returns(address)
func (_Router *RouterCallerSession) Dapp(arg0 *big.Int) (common.Address, error) {
	return _Router.Contract.Dapp(&_Router.CallOpts, arg0)
}

// FleetManagement is a free data retrieval call binding the contract method 0x53ef0542.
//
// Solidity: function fleetManagement() view returns(address)
func (_Router *RouterCaller) FleetManagement(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "fleetManagement")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FleetManagement is a free data retrieval call binding the contract method 0x53ef0542.
//
// Solidity: function fleetManagement() view returns(address)
func (_Router *RouterSession) FleetManagement() (common.Address, error) {
	return _Router.Contract.FleetManagement(&_Router.CallOpts)
}

// FleetManagement is a free data retrieval call binding the contract method 0x53ef0542.
//
// Solidity: function fleetManagement() view returns(address)
func (_Router *RouterCallerSession) FleetManagement() (common.Address, error) {
	return _Router.Contract.FleetManagement(&_Router.CallOpts)
}

// ProjectStore is a free data retrieval call binding the contract method 0xa0fadaaa.
//
// Solidity: function projectStore() view returns(address)
func (_Router *RouterCaller) ProjectStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "projectStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProjectStore is a free data retrieval call binding the contract method 0xa0fadaaa.
//
// Solidity: function projectStore() view returns(address)
func (_Router *RouterSession) ProjectStore() (common.Address, error) {
	return _Router.Contract.ProjectStore(&_Router.CallOpts)
}

// ProjectStore is a free data retrieval call binding the contract method 0xa0fadaaa.
//
// Solidity: function projectStore() view returns(address)
func (_Router *RouterCallerSession) ProjectStore() (common.Address, error) {
	return _Router.Contract.ProjectStore(&_Router.CallOpts)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_Router *RouterCaller) TaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "taskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_Router *RouterSession) TaskManager() (common.Address, error) {
	return _Router.Contract.TaskManager(&_Router.CallOpts)
}

// TaskManager is a free data retrieval call binding the contract method 0xa50a640e.
//
// Solidity: function taskManager() view returns(address)
func (_Router *RouterCallerSession) TaskManager() (common.Address, error) {
	return _Router.Contract.TaskManager(&_Router.CallOpts)
}

// BindDapp is a paid mutator transaction binding the contract method 0x85a7d275.
//
// Solidity: function bindDapp(uint256 _projectId, address _dapp) returns()
func (_Router *RouterTransactor) BindDapp(opts *bind.TransactOpts, _projectId *big.Int, _dapp common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "bindDapp", _projectId, _dapp)
}

// BindDapp is a paid mutator transaction binding the contract method 0x85a7d275.
//
// Solidity: function bindDapp(uint256 _projectId, address _dapp) returns()
func (_Router *RouterSession) BindDapp(_projectId *big.Int, _dapp common.Address) (*types.Transaction, error) {
	return _Router.Contract.BindDapp(&_Router.TransactOpts, _projectId, _dapp)
}

// BindDapp is a paid mutator transaction binding the contract method 0x85a7d275.
//
// Solidity: function bindDapp(uint256 _projectId, address _dapp) returns()
func (_Router *RouterTransactorSession) BindDapp(_projectId *big.Int, _dapp common.Address) (*types.Transaction, error) {
	return _Router.Contract.BindDapp(&_Router.TransactOpts, _projectId, _dapp)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _fleetManagement, address _projectStore) returns()
func (_Router *RouterTransactor) Initialize(opts *bind.TransactOpts, _fleetManagement common.Address, _projectStore common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "initialize", _fleetManagement, _projectStore)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _fleetManagement, address _projectStore) returns()
func (_Router *RouterSession) Initialize(_fleetManagement common.Address, _projectStore common.Address) (*types.Transaction, error) {
	return _Router.Contract.Initialize(&_Router.TransactOpts, _fleetManagement, _projectStore)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _fleetManagement, address _projectStore) returns()
func (_Router *RouterTransactorSession) Initialize(_fleetManagement common.Address, _projectStore common.Address) (*types.Transaction, error) {
	return _Router.Contract.Initialize(&_Router.TransactOpts, _fleetManagement, _projectStore)
}

// Route is a paid mutator transaction binding the contract method 0x0cb7f523.
//
// Solidity: function route(uint256 _projectId, uint256 _proverId, string _clientId, bytes _data, bytes32 _taskId) returns()
func (_Router *RouterTransactor) Route(opts *bind.TransactOpts, _projectId *big.Int, _proverId *big.Int, _clientId string, _data []byte, _taskId [32]byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "route", _projectId, _proverId, _clientId, _data, _taskId)
}

// Route is a paid mutator transaction binding the contract method 0x0cb7f523.
//
// Solidity: function route(uint256 _projectId, uint256 _proverId, string _clientId, bytes _data, bytes32 _taskId) returns()
func (_Router *RouterSession) Route(_projectId *big.Int, _proverId *big.Int, _clientId string, _data []byte, _taskId [32]byte) (*types.Transaction, error) {
	return _Router.Contract.Route(&_Router.TransactOpts, _projectId, _proverId, _clientId, _data, _taskId)
}

// Route is a paid mutator transaction binding the contract method 0x0cb7f523.
//
// Solidity: function route(uint256 _projectId, uint256 _proverId, string _clientId, bytes _data, bytes32 _taskId) returns()
func (_Router *RouterTransactorSession) Route(_projectId *big.Int, _proverId *big.Int, _clientId string, _data []byte, _taskId [32]byte) (*types.Transaction, error) {
	return _Router.Contract.Route(&_Router.TransactOpts, _projectId, _proverId, _clientId, _data, _taskId)
}

// SetTaskManager is a paid mutator transaction binding the contract method 0x327d0a60.
//
// Solidity: function setTaskManager(address _add) returns()
func (_Router *RouterTransactor) SetTaskManager(opts *bind.TransactOpts, _add common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setTaskManager", _add)
}

// SetTaskManager is a paid mutator transaction binding the contract method 0x327d0a60.
//
// Solidity: function setTaskManager(address _add) returns()
func (_Router *RouterSession) SetTaskManager(_add common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetTaskManager(&_Router.TransactOpts, _add)
}

// SetTaskManager is a paid mutator transaction binding the contract method 0x327d0a60.
//
// Solidity: function setTaskManager(address _add) returns()
func (_Router *RouterTransactorSession) SetTaskManager(_add common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetTaskManager(&_Router.TransactOpts, _add)
}

// UnbindDapp is a paid mutator transaction binding the contract method 0xd869758c.
//
// Solidity: function unbindDapp(uint256 _projectId) returns()
func (_Router *RouterTransactor) UnbindDapp(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "unbindDapp", _projectId)
}

// UnbindDapp is a paid mutator transaction binding the contract method 0xd869758c.
//
// Solidity: function unbindDapp(uint256 _projectId) returns()
func (_Router *RouterSession) UnbindDapp(_projectId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.UnbindDapp(&_Router.TransactOpts, _projectId)
}

// UnbindDapp is a paid mutator transaction binding the contract method 0xd869758c.
//
// Solidity: function unbindDapp(uint256 _projectId) returns()
func (_Router *RouterTransactorSession) UnbindDapp(_projectId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.UnbindDapp(&_Router.TransactOpts, _projectId)
}

// RouterDappBoundIterator is returned from FilterDappBound and is used to iterate over the raw logs and unpacked data for DappBound events raised by the Router contract.
type RouterDappBoundIterator struct {
	Event *RouterDappBound // Event containing the contract specifics and raw log

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
func (it *RouterDappBoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterDappBound)
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
		it.Event = new(RouterDappBound)
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
func (it *RouterDappBoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterDappBoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterDappBound represents a DappBound event raised by the Router contract.
type RouterDappBound struct {
	ProjectId *big.Int
	Operator  common.Address
	Dapp      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDappBound is a free log retrieval operation binding the contract event 0xf121fc55c0fd19e108d2d5642aff2967949fb708d9b985093c530a8a1fb97778.
//
// Solidity: event DappBound(uint256 indexed projectId, address indexed operator, address dapp)
func (_Router *RouterFilterer) FilterDappBound(opts *bind.FilterOpts, projectId []*big.Int, operator []common.Address) (*RouterDappBoundIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "DappBound", projectIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &RouterDappBoundIterator{contract: _Router.contract, event: "DappBound", logs: logs, sub: sub}, nil
}

// WatchDappBound is a free log subscription operation binding the contract event 0xf121fc55c0fd19e108d2d5642aff2967949fb708d9b985093c530a8a1fb97778.
//
// Solidity: event DappBound(uint256 indexed projectId, address indexed operator, address dapp)
func (_Router *RouterFilterer) WatchDappBound(opts *bind.WatchOpts, sink chan<- *RouterDappBound, projectId []*big.Int, operator []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "DappBound", projectIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterDappBound)
				if err := _Router.contract.UnpackLog(event, "DappBound", log); err != nil {
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

// ParseDappBound is a log parse operation binding the contract event 0xf121fc55c0fd19e108d2d5642aff2967949fb708d9b985093c530a8a1fb97778.
//
// Solidity: event DappBound(uint256 indexed projectId, address indexed operator, address dapp)
func (_Router *RouterFilterer) ParseDappBound(log types.Log) (*RouterDappBound, error) {
	event := new(RouterDappBound)
	if err := _Router.contract.UnpackLog(event, "DappBound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterDappUnboundIterator is returned from FilterDappUnbound and is used to iterate over the raw logs and unpacked data for DappUnbound events raised by the Router contract.
type RouterDappUnboundIterator struct {
	Event *RouterDappUnbound // Event containing the contract specifics and raw log

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
func (it *RouterDappUnboundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterDappUnbound)
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
		it.Event = new(RouterDappUnbound)
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
func (it *RouterDappUnboundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterDappUnboundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterDappUnbound represents a DappUnbound event raised by the Router contract.
type RouterDappUnbound struct {
	ProjectId *big.Int
	Operator  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDappUnbound is a free log retrieval operation binding the contract event 0x7019ee8601397d5c4fe244404e2428a9c0b0a4d8679186133186cc01376ee9f1.
//
// Solidity: event DappUnbound(uint256 indexed projectId, address indexed operator)
func (_Router *RouterFilterer) FilterDappUnbound(opts *bind.FilterOpts, projectId []*big.Int, operator []common.Address) (*RouterDappUnboundIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "DappUnbound", projectIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &RouterDappUnboundIterator{contract: _Router.contract, event: "DappUnbound", logs: logs, sub: sub}, nil
}

// WatchDappUnbound is a free log subscription operation binding the contract event 0x7019ee8601397d5c4fe244404e2428a9c0b0a4d8679186133186cc01376ee9f1.
//
// Solidity: event DappUnbound(uint256 indexed projectId, address indexed operator)
func (_Router *RouterFilterer) WatchDappUnbound(opts *bind.WatchOpts, sink chan<- *RouterDappUnbound, projectId []*big.Int, operator []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "DappUnbound", projectIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterDappUnbound)
				if err := _Router.contract.UnpackLog(event, "DappUnbound", log); err != nil {
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

// ParseDappUnbound is a log parse operation binding the contract event 0x7019ee8601397d5c4fe244404e2428a9c0b0a4d8679186133186cc01376ee9f1.
//
// Solidity: event DappUnbound(uint256 indexed projectId, address indexed operator)
func (_Router *RouterFilterer) ParseDappUnbound(log types.Log) (*RouterDappUnbound, error) {
	event := new(RouterDappUnbound)
	if err := _Router.contract.UnpackLog(event, "DappUnbound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterDataProcessedIterator is returned from FilterDataProcessed and is used to iterate over the raw logs and unpacked data for DataProcessed events raised by the Router contract.
type RouterDataProcessedIterator struct {
	Event *RouterDataProcessed // Event containing the contract specifics and raw log

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
func (it *RouterDataProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterDataProcessed)
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
		it.Event = new(RouterDataProcessed)
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
func (it *RouterDataProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterDataProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterDataProcessed represents a DataProcessed event raised by the Router contract.
type RouterDataProcessed struct {
	ProjectId *big.Int
	Router    *big.Int
	Operator  common.Address
	Success   bool
	Error     []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDataProcessed is a free log retrieval operation binding the contract event 0x588bd93a32d237b3f0595fb64f6e8d7f74c1a5a893cd1094db2a31ac584480f7.
//
// Solidity: event DataProcessed(uint256 indexed projectId, uint256 indexed router, address indexed operator, bool success, bytes error)
func (_Router *RouterFilterer) FilterDataProcessed(opts *bind.FilterOpts, projectId []*big.Int, router []*big.Int, operator []common.Address) (*RouterDataProcessedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "DataProcessed", projectIdRule, routerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &RouterDataProcessedIterator{contract: _Router.contract, event: "DataProcessed", logs: logs, sub: sub}, nil
}

// WatchDataProcessed is a free log subscription operation binding the contract event 0x588bd93a32d237b3f0595fb64f6e8d7f74c1a5a893cd1094db2a31ac584480f7.
//
// Solidity: event DataProcessed(uint256 indexed projectId, uint256 indexed router, address indexed operator, bool success, bytes error)
func (_Router *RouterFilterer) WatchDataProcessed(opts *bind.WatchOpts, sink chan<- *RouterDataProcessed, projectId []*big.Int, router []*big.Int, operator []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "DataProcessed", projectIdRule, routerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterDataProcessed)
				if err := _Router.contract.UnpackLog(event, "DataProcessed", log); err != nil {
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

// ParseDataProcessed is a log parse operation binding the contract event 0x588bd93a32d237b3f0595fb64f6e8d7f74c1a5a893cd1094db2a31ac584480f7.
//
// Solidity: event DataProcessed(uint256 indexed projectId, uint256 indexed router, address indexed operator, bool success, bytes error)
func (_Router *RouterFilterer) ParseDataProcessed(log types.Log) (*RouterDataProcessed, error) {
	event := new(RouterDataProcessed)
	if err := _Router.contract.UnpackLog(event, "DataProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Router contract.
type RouterInitializedIterator struct {
	Event *RouterInitialized // Event containing the contract specifics and raw log

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
func (it *RouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterInitialized)
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
		it.Event = new(RouterInitialized)
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
func (it *RouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterInitialized represents a Initialized event raised by the Router contract.
type RouterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Router *RouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*RouterInitializedIterator, error) {

	logs, sub, err := _Router.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RouterInitializedIterator{contract: _Router.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Router *RouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RouterInitialized) (event.Subscription, error) {

	logs, sub, err := _Router.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterInitialized)
				if err := _Router.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Router *RouterFilterer) ParseInitialized(log types.Log) (*RouterInitialized, error) {
	event := new(RouterInitialized)
	if err := _Router.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
