// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockheadervalidator

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

// BlockHeader is an auto generated low-level Go binding around an user-defined struct.
type BlockHeader struct {
	Meta       [4]byte
	Prevhash   [32]byte
	MerkleRoot [32]byte
	Nbits      uint32
	Nonce      [8]byte
}

// BlockheadervalidatorMetaData contains all meta data concerning the Blockheadervalidator contract.
var BlockheadervalidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"nbits\",\"type\":\"uint32\"}],\"name\":\"NBitsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"nbits\",\"type\":\"uint32\"}],\"name\":\"setAdhocNBits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"updateDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"meta\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"prevhash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"nbits\",\"type\":\"uint32\"},{\"internalType\":\"bytes8\",\"name\":\"nonce\",\"type\":\"bytes8\"}],\"internalType\":\"structBlockHeader\",\"name\":\"header\",\"type\":\"tuple\"}],\"name\":\"validate\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// BlockheadervalidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockheadervalidatorMetaData.ABI instead.
var BlockheadervalidatorABI = BlockheadervalidatorMetaData.ABI

// Blockheadervalidator is an auto generated Go binding around an Ethereum contract.
type Blockheadervalidator struct {
	BlockheadervalidatorCaller     // Read-only binding to the contract
	BlockheadervalidatorTransactor // Write-only binding to the contract
	BlockheadervalidatorFilterer   // Log filterer for contract events
}

// BlockheadervalidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockheadervalidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockheadervalidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockheadervalidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockheadervalidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockheadervalidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockheadervalidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockheadervalidatorSession struct {
	Contract     *Blockheadervalidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BlockheadervalidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockheadervalidatorCallerSession struct {
	Contract *BlockheadervalidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// BlockheadervalidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockheadervalidatorTransactorSession struct {
	Contract     *BlockheadervalidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// BlockheadervalidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockheadervalidatorRaw struct {
	Contract *Blockheadervalidator // Generic contract binding to access the raw methods on
}

// BlockheadervalidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockheadervalidatorCallerRaw struct {
	Contract *BlockheadervalidatorCaller // Generic read-only contract binding to access the raw methods on
}

// BlockheadervalidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockheadervalidatorTransactorRaw struct {
	Contract *BlockheadervalidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockheadervalidator creates a new instance of Blockheadervalidator, bound to a specific deployed contract.
func NewBlockheadervalidator(address common.Address, backend bind.ContractBackend) (*Blockheadervalidator, error) {
	contract, err := bindBlockheadervalidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blockheadervalidator{BlockheadervalidatorCaller: BlockheadervalidatorCaller{contract: contract}, BlockheadervalidatorTransactor: BlockheadervalidatorTransactor{contract: contract}, BlockheadervalidatorFilterer: BlockheadervalidatorFilterer{contract: contract}}, nil
}

// NewBlockheadervalidatorCaller creates a new read-only instance of Blockheadervalidator, bound to a specific deployed contract.
func NewBlockheadervalidatorCaller(address common.Address, caller bind.ContractCaller) (*BlockheadervalidatorCaller, error) {
	contract, err := bindBlockheadervalidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockheadervalidatorCaller{contract: contract}, nil
}

// NewBlockheadervalidatorTransactor creates a new write-only instance of Blockheadervalidator, bound to a specific deployed contract.
func NewBlockheadervalidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockheadervalidatorTransactor, error) {
	contract, err := bindBlockheadervalidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockheadervalidatorTransactor{contract: contract}, nil
}

// NewBlockheadervalidatorFilterer creates a new log filterer instance of Blockheadervalidator, bound to a specific deployed contract.
func NewBlockheadervalidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockheadervalidatorFilterer, error) {
	contract, err := bindBlockheadervalidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockheadervalidatorFilterer{contract: contract}, nil
}

// bindBlockheadervalidator binds a generic wrapper to an already deployed contract.
func bindBlockheadervalidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockheadervalidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockheadervalidator *BlockheadervalidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockheadervalidator.Contract.BlockheadervalidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockheadervalidator *BlockheadervalidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockheadervalidator.Contract.BlockheadervalidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockheadervalidator *BlockheadervalidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockheadervalidator.Contract.BlockheadervalidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockheadervalidator *BlockheadervalidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockheadervalidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockheadervalidator *BlockheadervalidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockheadervalidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockheadervalidator *BlockheadervalidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockheadervalidator.Contract.contract.Transact(opts, method, params...)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Blockheadervalidator *BlockheadervalidatorCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blockheadervalidator.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Blockheadervalidator *BlockheadervalidatorSession) Operator() (common.Address, error) {
	return _Blockheadervalidator.Contract.Operator(&_Blockheadervalidator.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Blockheadervalidator *BlockheadervalidatorCallerSession) Operator() (common.Address, error) {
	return _Blockheadervalidator.Contract.Operator(&_Blockheadervalidator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockheadervalidator *BlockheadervalidatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blockheadervalidator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockheadervalidator *BlockheadervalidatorSession) Owner() (common.Address, error) {
	return _Blockheadervalidator.Contract.Owner(&_Blockheadervalidator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockheadervalidator *BlockheadervalidatorCallerSession) Owner() (common.Address, error) {
	return _Blockheadervalidator.Contract.Owner(&_Blockheadervalidator.CallOpts)
}

// Validate is a free data retrieval call binding the contract method 0x39e9bce8.
//
// Solidity: function validate((bytes4,bytes32,bytes32,uint32,bytes8) header) pure returns(bytes)
func (_Blockheadervalidator *BlockheadervalidatorCaller) Validate(opts *bind.CallOpts, header BlockHeader) ([]byte, error) {
	var out []interface{}
	err := _Blockheadervalidator.contract.Call(opts, &out, "validate", header)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Validate is a free data retrieval call binding the contract method 0x39e9bce8.
//
// Solidity: function validate((bytes4,bytes32,bytes32,uint32,bytes8) header) pure returns(bytes)
func (_Blockheadervalidator *BlockheadervalidatorSession) Validate(header BlockHeader) ([]byte, error) {
	return _Blockheadervalidator.Contract.Validate(&_Blockheadervalidator.CallOpts, header)
}

// Validate is a free data retrieval call binding the contract method 0x39e9bce8.
//
// Solidity: function validate((bytes4,bytes32,bytes32,uint32,bytes8) header) pure returns(bytes)
func (_Blockheadervalidator *BlockheadervalidatorCallerSession) Validate(header BlockHeader) ([]byte, error) {
	return _Blockheadervalidator.Contract.Validate(&_Blockheadervalidator.CallOpts, header)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blockheadervalidator *BlockheadervalidatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockheadervalidator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blockheadervalidator *BlockheadervalidatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Blockheadervalidator.Contract.RenounceOwnership(&_Blockheadervalidator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blockheadervalidator *BlockheadervalidatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Blockheadervalidator.Contract.RenounceOwnership(&_Blockheadervalidator.TransactOpts)
}

// SetAdhocNBits is a paid mutator transaction binding the contract method 0xe115953c.
//
// Solidity: function setAdhocNBits(uint32 nbits) returns()
func (_Blockheadervalidator *BlockheadervalidatorTransactor) SetAdhocNBits(opts *bind.TransactOpts, nbits uint32) (*types.Transaction, error) {
	return _Blockheadervalidator.contract.Transact(opts, "setAdhocNBits", nbits)
}

// SetAdhocNBits is a paid mutator transaction binding the contract method 0xe115953c.
//
// Solidity: function setAdhocNBits(uint32 nbits) returns()
func (_Blockheadervalidator *BlockheadervalidatorSession) SetAdhocNBits(nbits uint32) (*types.Transaction, error) {
	return _Blockheadervalidator.Contract.SetAdhocNBits(&_Blockheadervalidator.TransactOpts, nbits)
}

// SetAdhocNBits is a paid mutator transaction binding the contract method 0xe115953c.
//
// Solidity: function setAdhocNBits(uint32 nbits) returns()
func (_Blockheadervalidator *BlockheadervalidatorTransactorSession) SetAdhocNBits(nbits uint32) (*types.Transaction, error) {
	return _Blockheadervalidator.Contract.SetAdhocNBits(&_Blockheadervalidator.TransactOpts, nbits)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Blockheadervalidator *BlockheadervalidatorTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Blockheadervalidator.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Blockheadervalidator *BlockheadervalidatorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Blockheadervalidator.Contract.SetOperator(&_Blockheadervalidator.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Blockheadervalidator *BlockheadervalidatorTransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Blockheadervalidator.Contract.SetOperator(&_Blockheadervalidator.TransactOpts, _operator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockheadervalidator *BlockheadervalidatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Blockheadervalidator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockheadervalidator *BlockheadervalidatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blockheadervalidator.Contract.TransferOwnership(&_Blockheadervalidator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockheadervalidator *BlockheadervalidatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blockheadervalidator.Contract.TransferOwnership(&_Blockheadervalidator.TransactOpts, newOwner)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0x1b50ad09.
//
// Solidity: function updateDuration(uint256 duration) returns()
func (_Blockheadervalidator *BlockheadervalidatorTransactor) UpdateDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Blockheadervalidator.contract.Transact(opts, "updateDuration", duration)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0x1b50ad09.
//
// Solidity: function updateDuration(uint256 duration) returns()
func (_Blockheadervalidator *BlockheadervalidatorSession) UpdateDuration(duration *big.Int) (*types.Transaction, error) {
	return _Blockheadervalidator.Contract.UpdateDuration(&_Blockheadervalidator.TransactOpts, duration)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0x1b50ad09.
//
// Solidity: function updateDuration(uint256 duration) returns()
func (_Blockheadervalidator *BlockheadervalidatorTransactorSession) UpdateDuration(duration *big.Int) (*types.Transaction, error) {
	return _Blockheadervalidator.Contract.UpdateDuration(&_Blockheadervalidator.TransactOpts, duration)
}

// BlockheadervalidatorNBitsSetIterator is returned from FilterNBitsSet and is used to iterate over the raw logs and unpacked data for NBitsSet events raised by the Blockheadervalidator contract.
type BlockheadervalidatorNBitsSetIterator struct {
	Event *BlockheadervalidatorNBitsSet // Event containing the contract specifics and raw log

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
func (it *BlockheadervalidatorNBitsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockheadervalidatorNBitsSet)
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
		it.Event = new(BlockheadervalidatorNBitsSet)
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
func (it *BlockheadervalidatorNBitsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockheadervalidatorNBitsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockheadervalidatorNBitsSet represents a NBitsSet event raised by the Blockheadervalidator contract.
type BlockheadervalidatorNBitsSet struct {
	Nbits uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNBitsSet is a free log retrieval operation binding the contract event 0xe8dd26f51ea2466f2a2a2bad6b1065a8ba4f43f587f210d03bc7ee3c24b25a98.
//
// Solidity: event NBitsSet(uint32 nbits)
func (_Blockheadervalidator *BlockheadervalidatorFilterer) FilterNBitsSet(opts *bind.FilterOpts) (*BlockheadervalidatorNBitsSetIterator, error) {

	logs, sub, err := _Blockheadervalidator.contract.FilterLogs(opts, "NBitsSet")
	if err != nil {
		return nil, err
	}
	return &BlockheadervalidatorNBitsSetIterator{contract: _Blockheadervalidator.contract, event: "NBitsSet", logs: logs, sub: sub}, nil
}

// WatchNBitsSet is a free log subscription operation binding the contract event 0xe8dd26f51ea2466f2a2a2bad6b1065a8ba4f43f587f210d03bc7ee3c24b25a98.
//
// Solidity: event NBitsSet(uint32 nbits)
func (_Blockheadervalidator *BlockheadervalidatorFilterer) WatchNBitsSet(opts *bind.WatchOpts, sink chan<- *BlockheadervalidatorNBitsSet) (event.Subscription, error) {

	logs, sub, err := _Blockheadervalidator.contract.WatchLogs(opts, "NBitsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockheadervalidatorNBitsSet)
				if err := _Blockheadervalidator.contract.UnpackLog(event, "NBitsSet", log); err != nil {
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

// ParseNBitsSet is a log parse operation binding the contract event 0xe8dd26f51ea2466f2a2a2bad6b1065a8ba4f43f587f210d03bc7ee3c24b25a98.
//
// Solidity: event NBitsSet(uint32 nbits)
func (_Blockheadervalidator *BlockheadervalidatorFilterer) ParseNBitsSet(log types.Log) (*BlockheadervalidatorNBitsSet, error) {
	event := new(BlockheadervalidatorNBitsSet)
	if err := _Blockheadervalidator.contract.UnpackLog(event, "NBitsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockheadervalidatorOperatorSetIterator is returned from FilterOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorSet events raised by the Blockheadervalidator contract.
type BlockheadervalidatorOperatorSetIterator struct {
	Event *BlockheadervalidatorOperatorSet // Event containing the contract specifics and raw log

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
func (it *BlockheadervalidatorOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockheadervalidatorOperatorSet)
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
		it.Event = new(BlockheadervalidatorOperatorSet)
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
func (it *BlockheadervalidatorOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockheadervalidatorOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockheadervalidatorOperatorSet represents a OperatorSet event raised by the Blockheadervalidator contract.
type BlockheadervalidatorOperatorSet struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSet is a free log retrieval operation binding the contract event 0x99d737e0adf2c449d71890b86772885ec7959b152ddb265f76325b6e68e105d3.
//
// Solidity: event OperatorSet(address operator)
func (_Blockheadervalidator *BlockheadervalidatorFilterer) FilterOperatorSet(opts *bind.FilterOpts) (*BlockheadervalidatorOperatorSetIterator, error) {

	logs, sub, err := _Blockheadervalidator.contract.FilterLogs(opts, "OperatorSet")
	if err != nil {
		return nil, err
	}
	return &BlockheadervalidatorOperatorSetIterator{contract: _Blockheadervalidator.contract, event: "OperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSet is a free log subscription operation binding the contract event 0x99d737e0adf2c449d71890b86772885ec7959b152ddb265f76325b6e68e105d3.
//
// Solidity: event OperatorSet(address operator)
func (_Blockheadervalidator *BlockheadervalidatorFilterer) WatchOperatorSet(opts *bind.WatchOpts, sink chan<- *BlockheadervalidatorOperatorSet) (event.Subscription, error) {

	logs, sub, err := _Blockheadervalidator.contract.WatchLogs(opts, "OperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockheadervalidatorOperatorSet)
				if err := _Blockheadervalidator.contract.UnpackLog(event, "OperatorSet", log); err != nil {
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
// Solidity: event OperatorSet(address operator)
func (_Blockheadervalidator *BlockheadervalidatorFilterer) ParseOperatorSet(log types.Log) (*BlockheadervalidatorOperatorSet, error) {
	event := new(BlockheadervalidatorOperatorSet)
	if err := _Blockheadervalidator.contract.UnpackLog(event, "OperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockheadervalidatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Blockheadervalidator contract.
type BlockheadervalidatorOwnershipTransferredIterator struct {
	Event *BlockheadervalidatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlockheadervalidatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockheadervalidatorOwnershipTransferred)
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
		it.Event = new(BlockheadervalidatorOwnershipTransferred)
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
func (it *BlockheadervalidatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockheadervalidatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockheadervalidatorOwnershipTransferred represents a OwnershipTransferred event raised by the Blockheadervalidator contract.
type BlockheadervalidatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blockheadervalidator *BlockheadervalidatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlockheadervalidatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blockheadervalidator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlockheadervalidatorOwnershipTransferredIterator{contract: _Blockheadervalidator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blockheadervalidator *BlockheadervalidatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlockheadervalidatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blockheadervalidator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockheadervalidatorOwnershipTransferred)
				if err := _Blockheadervalidator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Blockheadervalidator *BlockheadervalidatorFilterer) ParseOwnershipTransferred(log types.Log) (*BlockheadervalidatorOwnershipTransferred, error) {
	event := new(BlockheadervalidatorOwnershipTransferred)
	if err := _Blockheadervalidator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
