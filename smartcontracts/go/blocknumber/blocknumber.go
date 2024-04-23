// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blocknumber

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

// BlocknumberMetaData contains all meta data concerning the Blocknumber contract.
var BlocknumberMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"blockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BlocknumberABI is the input ABI used to generate the binding from.
// Deprecated: Use BlocknumberMetaData.ABI instead.
var BlocknumberABI = BlocknumberMetaData.ABI

// Blocknumber is an auto generated Go binding around an Ethereum contract.
type Blocknumber struct {
	BlocknumberCaller     // Read-only binding to the contract
	BlocknumberTransactor // Write-only binding to the contract
	BlocknumberFilterer   // Log filterer for contract events
}

// BlocknumberCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlocknumberCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlocknumberTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlocknumberTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlocknumberFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlocknumberFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlocknumberSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlocknumberSession struct {
	Contract     *Blocknumber      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlocknumberCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlocknumberCallerSession struct {
	Contract *BlocknumberCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BlocknumberTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlocknumberTransactorSession struct {
	Contract     *BlocknumberTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BlocknumberRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlocknumberRaw struct {
	Contract *Blocknumber // Generic contract binding to access the raw methods on
}

// BlocknumberCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlocknumberCallerRaw struct {
	Contract *BlocknumberCaller // Generic read-only contract binding to access the raw methods on
}

// BlocknumberTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlocknumberTransactorRaw struct {
	Contract *BlocknumberTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlocknumber creates a new instance of Blocknumber, bound to a specific deployed contract.
func NewBlocknumber(address common.Address, backend bind.ContractBackend) (*Blocknumber, error) {
	contract, err := bindBlocknumber(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blocknumber{BlocknumberCaller: BlocknumberCaller{contract: contract}, BlocknumberTransactor: BlocknumberTransactor{contract: contract}, BlocknumberFilterer: BlocknumberFilterer{contract: contract}}, nil
}

// NewBlocknumberCaller creates a new read-only instance of Blocknumber, bound to a specific deployed contract.
func NewBlocknumberCaller(address common.Address, caller bind.ContractCaller) (*BlocknumberCaller, error) {
	contract, err := bindBlocknumber(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlocknumberCaller{contract: contract}, nil
}

// NewBlocknumberTransactor creates a new write-only instance of Blocknumber, bound to a specific deployed contract.
func NewBlocknumberTransactor(address common.Address, transactor bind.ContractTransactor) (*BlocknumberTransactor, error) {
	contract, err := bindBlocknumber(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlocknumberTransactor{contract: contract}, nil
}

// NewBlocknumberFilterer creates a new log filterer instance of Blocknumber, bound to a specific deployed contract.
func NewBlocknumberFilterer(address common.Address, filterer bind.ContractFilterer) (*BlocknumberFilterer, error) {
	contract, err := bindBlocknumber(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlocknumberFilterer{contract: contract}, nil
}

// bindBlocknumber binds a generic wrapper to an already deployed contract.
func bindBlocknumber(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlocknumberMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blocknumber *BlocknumberRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blocknumber.Contract.BlocknumberCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blocknumber *BlocknumberRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blocknumber.Contract.BlocknumberTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blocknumber *BlocknumberRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blocknumber.Contract.BlocknumberTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blocknumber *BlocknumberCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blocknumber.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blocknumber *BlocknumberTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blocknumber.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blocknumber *BlocknumberTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blocknumber.Contract.contract.Transact(opts, method, params...)
}

// BlockNumber is a free data retrieval call binding the contract method 0x57e871e7.
//
// Solidity: function blockNumber() view returns(uint256)
func (_Blocknumber *BlocknumberCaller) BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blocknumber.contract.Call(opts, &out, "blockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockNumber is a free data retrieval call binding the contract method 0x57e871e7.
//
// Solidity: function blockNumber() view returns(uint256)
func (_Blocknumber *BlocknumberSession) BlockNumber() (*big.Int, error) {
	return _Blocknumber.Contract.BlockNumber(&_Blocknumber.CallOpts)
}

// BlockNumber is a free data retrieval call binding the contract method 0x57e871e7.
//
// Solidity: function blockNumber() view returns(uint256)
func (_Blocknumber *BlocknumberCallerSession) BlockNumber() (*big.Int, error) {
	return _Blocknumber.Contract.BlockNumber(&_Blocknumber.CallOpts)
}
