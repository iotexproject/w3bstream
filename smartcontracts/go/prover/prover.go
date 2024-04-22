// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package prover

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

// ProverMetaData contains all meta data concerning the Prover contract.
var ProverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MinterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"typ\",\"type\":\"uint256\"}],\"name\":\"NodeTypeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ProverPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ProverResumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"changeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"nodeType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"ownerOfOperator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"prover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_type\",\"type\":\"uint256\"}],\"name\":\"updateNodeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProverABI is the input ABI used to generate the binding from.
// Deprecated: Use ProverMetaData.ABI instead.
var ProverABI = ProverMetaData.ABI

// Prover is an auto generated Go binding around an Ethereum contract.
type Prover struct {
	ProverCaller     // Read-only binding to the contract
	ProverTransactor // Write-only binding to the contract
	ProverFilterer   // Log filterer for contract events
}

// ProverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProverSession struct {
	Contract     *Prover           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProverCallerSession struct {
	Contract *ProverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProverTransactorSession struct {
	Contract     *ProverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProverRaw struct {
	Contract *Prover // Generic contract binding to access the raw methods on
}

// ProverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProverCallerRaw struct {
	Contract *ProverCaller // Generic read-only contract binding to access the raw methods on
}

// ProverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProverTransactorRaw struct {
	Contract *ProverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProver creates a new instance of Prover, bound to a specific deployed contract.
func NewProver(address common.Address, backend bind.ContractBackend) (*Prover, error) {
	contract, err := bindProver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Prover{ProverCaller: ProverCaller{contract: contract}, ProverTransactor: ProverTransactor{contract: contract}, ProverFilterer: ProverFilterer{contract: contract}}, nil
}

// NewProverCaller creates a new read-only instance of Prover, bound to a specific deployed contract.
func NewProverCaller(address common.Address, caller bind.ContractCaller) (*ProverCaller, error) {
	contract, err := bindProver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProverCaller{contract: contract}, nil
}

// NewProverTransactor creates a new write-only instance of Prover, bound to a specific deployed contract.
func NewProverTransactor(address common.Address, transactor bind.ContractTransactor) (*ProverTransactor, error) {
	contract, err := bindProver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProverTransactor{contract: contract}, nil
}

// NewProverFilterer creates a new log filterer instance of Prover, bound to a specific deployed contract.
func NewProverFilterer(address common.Address, filterer bind.ContractFilterer) (*ProverFilterer, error) {
	contract, err := bindProver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProverFilterer{contract: contract}, nil
}

// bindProver binds a generic wrapper to an already deployed contract.
func bindProver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Prover *ProverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Prover.Contract.ProverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Prover *ProverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prover.Contract.ProverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Prover *ProverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Prover.Contract.ProverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Prover *ProverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Prover.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Prover *ProverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prover.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Prover *ProverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Prover.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Prover *ProverCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Prover *ProverSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Prover.Contract.BalanceOf(&_Prover.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Prover *ProverCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Prover.Contract.BalanceOf(&_Prover.CallOpts, owner)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Prover *ProverCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Prover *ProverSession) Count() (*big.Int, error) {
	return _Prover.Contract.Count(&_Prover.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Prover *ProverCallerSession) Count() (*big.Int, error) {
	return _Prover.Contract.Count(&_Prover.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Prover *ProverCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Prover *ProverSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Prover.Contract.GetApproved(&_Prover.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Prover *ProverCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Prover.Contract.GetApproved(&_Prover.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Prover *ProverCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Prover *ProverSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Prover.Contract.IsApprovedForAll(&_Prover.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Prover *ProverCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Prover.Contract.IsApprovedForAll(&_Prover.CallOpts, owner, operator)
}

// IsPaused is a free data retrieval call binding the contract method 0xbdf2a43c.
//
// Solidity: function isPaused(uint256 _id) view returns(bool)
func (_Prover *ProverCaller) IsPaused(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "isPaused", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xbdf2a43c.
//
// Solidity: function isPaused(uint256 _id) view returns(bool)
func (_Prover *ProverSession) IsPaused(_id *big.Int) (bool, error) {
	return _Prover.Contract.IsPaused(&_Prover.CallOpts, _id)
}

// IsPaused is a free data retrieval call binding the contract method 0xbdf2a43c.
//
// Solidity: function isPaused(uint256 _id) view returns(bool)
func (_Prover *ProverCallerSession) IsPaused(_id *big.Int) (bool, error) {
	return _Prover.Contract.IsPaused(&_Prover.CallOpts, _id)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Prover *ProverCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Prover *ProverSession) Minter() (common.Address, error) {
	return _Prover.Contract.Minter(&_Prover.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Prover *ProverCallerSession) Minter() (common.Address, error) {
	return _Prover.Contract.Minter(&_Prover.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Prover *ProverCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Prover *ProverSession) Name() (string, error) {
	return _Prover.Contract.Name(&_Prover.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Prover *ProverCallerSession) Name() (string, error) {
	return _Prover.Contract.Name(&_Prover.CallOpts)
}

// NodeType is a free data retrieval call binding the contract method 0x1c794b84.
//
// Solidity: function nodeType(uint256 _id) view returns(uint256)
func (_Prover *ProverCaller) NodeType(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "nodeType", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeType is a free data retrieval call binding the contract method 0x1c794b84.
//
// Solidity: function nodeType(uint256 _id) view returns(uint256)
func (_Prover *ProverSession) NodeType(_id *big.Int) (*big.Int, error) {
	return _Prover.Contract.NodeType(&_Prover.CallOpts, _id)
}

// NodeType is a free data retrieval call binding the contract method 0x1c794b84.
//
// Solidity: function nodeType(uint256 _id) view returns(uint256)
func (_Prover *ProverCallerSession) NodeType(_id *big.Int) (*big.Int, error) {
	return _Prover.Contract.NodeType(&_Prover.CallOpts, _id)
}

// Operator is a free data retrieval call binding the contract method 0xab3d047f.
//
// Solidity: function operator(uint256 _id) view returns(address)
func (_Prover *ProverCaller) Operator(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "operator", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0xab3d047f.
//
// Solidity: function operator(uint256 _id) view returns(address)
func (_Prover *ProverSession) Operator(_id *big.Int) (common.Address, error) {
	return _Prover.Contract.Operator(&_Prover.CallOpts, _id)
}

// Operator is a free data retrieval call binding the contract method 0xab3d047f.
//
// Solidity: function operator(uint256 _id) view returns(address)
func (_Prover *ProverCallerSession) Operator(_id *big.Int) (common.Address, error) {
	return _Prover.Contract.Operator(&_Prover.CallOpts, _id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Prover *ProverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Prover *ProverSession) Owner() (common.Address, error) {
	return _Prover.Contract.Owner(&_Prover.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Prover *ProverCallerSession) Owner() (common.Address, error) {
	return _Prover.Contract.Owner(&_Prover.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Prover *ProverCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Prover *ProverSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Prover.Contract.OwnerOf(&_Prover.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Prover *ProverCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Prover.Contract.OwnerOf(&_Prover.CallOpts, tokenId)
}

// OwnerOfOperator is a free data retrieval call binding the contract method 0xaeb44406.
//
// Solidity: function ownerOfOperator(address _operator) view returns(uint256, address)
func (_Prover *ProverCaller) OwnerOfOperator(opts *bind.CallOpts, _operator common.Address) (*big.Int, common.Address, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "ownerOfOperator", _operator)

	if err != nil {
		return *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// OwnerOfOperator is a free data retrieval call binding the contract method 0xaeb44406.
//
// Solidity: function ownerOfOperator(address _operator) view returns(uint256, address)
func (_Prover *ProverSession) OwnerOfOperator(_operator common.Address) (*big.Int, common.Address, error) {
	return _Prover.Contract.OwnerOfOperator(&_Prover.CallOpts, _operator)
}

// OwnerOfOperator is a free data retrieval call binding the contract method 0xaeb44406.
//
// Solidity: function ownerOfOperator(address _operator) view returns(uint256, address)
func (_Prover *ProverCallerSession) OwnerOfOperator(_operator common.Address) (*big.Int, common.Address, error) {
	return _Prover.Contract.OwnerOfOperator(&_Prover.CallOpts, _operator)
}

// Prover is a free data retrieval call binding the contract method 0x2becbd3e.
//
// Solidity: function prover(uint256 _id) view returns(address)
func (_Prover *ProverCaller) Prover(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "prover", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Prover is a free data retrieval call binding the contract method 0x2becbd3e.
//
// Solidity: function prover(uint256 _id) view returns(address)
func (_Prover *ProverSession) Prover(_id *big.Int) (common.Address, error) {
	return _Prover.Contract.Prover(&_Prover.CallOpts, _id)
}

// Prover is a free data retrieval call binding the contract method 0x2becbd3e.
//
// Solidity: function prover(uint256 _id) view returns(address)
func (_Prover *ProverCallerSession) Prover(_id *big.Int) (common.Address, error) {
	return _Prover.Contract.Prover(&_Prover.CallOpts, _id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Prover *ProverCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Prover *ProverSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Prover.Contract.SupportsInterface(&_Prover.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Prover *ProverCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Prover.Contract.SupportsInterface(&_Prover.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Prover *ProverCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Prover *ProverSession) Symbol() (string, error) {
	return _Prover.Contract.Symbol(&_Prover.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Prover *ProverCallerSession) Symbol() (string, error) {
	return _Prover.Contract.Symbol(&_Prover.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Prover *ProverCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Prover *ProverSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Prover.Contract.TokenURI(&_Prover.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Prover *ProverCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Prover.Contract.TokenURI(&_Prover.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Prover *ProverTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Prover *ProverSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.Approve(&_Prover.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Prover *ProverTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.Approve(&_Prover.TransactOpts, to, tokenId)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x7a3bc817.
//
// Solidity: function changeOperator(uint256 _id, address _operator) returns()
func (_Prover *ProverTransactor) ChangeOperator(opts *bind.TransactOpts, _id *big.Int, _operator common.Address) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "changeOperator", _id, _operator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x7a3bc817.
//
// Solidity: function changeOperator(uint256 _id, address _operator) returns()
func (_Prover *ProverSession) ChangeOperator(_id *big.Int, _operator common.Address) (*types.Transaction, error) {
	return _Prover.Contract.ChangeOperator(&_Prover.TransactOpts, _id, _operator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x7a3bc817.
//
// Solidity: function changeOperator(uint256 _id, address _operator) returns()
func (_Prover *ProverTransactorSession) ChangeOperator(_id *big.Int, _operator common.Address) (*types.Transaction, error) {
	return _Prover.Contract.ChangeOperator(&_Prover.TransactOpts, _id, _operator)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_Prover *ProverTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "initialize", _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_Prover *ProverSession) Initialize(_name string, _symbol string) (*types.Transaction, error) {
	return _Prover.Contract.Initialize(&_Prover.TransactOpts, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_Prover *ProverTransactorSession) Initialize(_name string, _symbol string) (*types.Transaction, error) {
	return _Prover.Contract.Initialize(&_Prover.TransactOpts, _name, _symbol)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _account) returns(uint256 id_)
func (_Prover *ProverTransactor) Mint(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "mint", _account)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _account) returns(uint256 id_)
func (_Prover *ProverSession) Mint(_account common.Address) (*types.Transaction, error) {
	return _Prover.Contract.Mint(&_Prover.TransactOpts, _account)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _account) returns(uint256 id_)
func (_Prover *ProverTransactorSession) Mint(_account common.Address) (*types.Transaction, error) {
	return _Prover.Contract.Mint(&_Prover.TransactOpts, _account)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 _id) returns()
func (_Prover *ProverTransactor) Pause(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "pause", _id)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 _id) returns()
func (_Prover *ProverSession) Pause(_id *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.Pause(&_Prover.TransactOpts, _id)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 _id) returns()
func (_Prover *ProverTransactorSession) Pause(_id *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.Pause(&_Prover.TransactOpts, _id)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Prover *ProverTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Prover *ProverSession) RenounceOwnership() (*types.Transaction, error) {
	return _Prover.Contract.RenounceOwnership(&_Prover.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Prover *ProverTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Prover.Contract.RenounceOwnership(&_Prover.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x414000b5.
//
// Solidity: function resume(uint256 _id) returns()
func (_Prover *ProverTransactor) Resume(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "resume", _id)
}

// Resume is a paid mutator transaction binding the contract method 0x414000b5.
//
// Solidity: function resume(uint256 _id) returns()
func (_Prover *ProverSession) Resume(_id *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.Resume(&_Prover.TransactOpts, _id)
}

// Resume is a paid mutator transaction binding the contract method 0x414000b5.
//
// Solidity: function resume(uint256 _id) returns()
func (_Prover *ProverTransactorSession) Resume(_id *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.Resume(&_Prover.TransactOpts, _id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Prover *ProverTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Prover *ProverSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.SafeTransferFrom(&_Prover.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Prover *ProverTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.SafeTransferFrom(&_Prover.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Prover *ProverTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Prover *ProverSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Prover.Contract.SafeTransferFrom0(&_Prover.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Prover *ProverTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Prover.Contract.SafeTransferFrom0(&_Prover.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Prover *ProverTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Prover *ProverSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Prover.Contract.SetApprovalForAll(&_Prover.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Prover *ProverTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Prover.Contract.SetApprovalForAll(&_Prover.TransactOpts, operator, approved)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Prover *ProverTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Prover *ProverSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Prover.Contract.SetMinter(&_Prover.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Prover *ProverTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Prover.Contract.SetMinter(&_Prover.TransactOpts, _minter)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Prover *ProverTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Prover *ProverSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.TransferFrom(&_Prover.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Prover *ProverTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.TransferFrom(&_Prover.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Prover *ProverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Prover *ProverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Prover.Contract.TransferOwnership(&_Prover.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Prover *ProverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Prover.Contract.TransferOwnership(&_Prover.TransactOpts, newOwner)
}

// UpdateNodeType is a paid mutator transaction binding the contract method 0x6a07973f.
//
// Solidity: function updateNodeType(uint256 _id, uint256 _type) returns()
func (_Prover *ProverTransactor) UpdateNodeType(opts *bind.TransactOpts, _id *big.Int, _type *big.Int) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "updateNodeType", _id, _type)
}

// UpdateNodeType is a paid mutator transaction binding the contract method 0x6a07973f.
//
// Solidity: function updateNodeType(uint256 _id, uint256 _type) returns()
func (_Prover *ProverSession) UpdateNodeType(_id *big.Int, _type *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.UpdateNodeType(&_Prover.TransactOpts, _id, _type)
}

// UpdateNodeType is a paid mutator transaction binding the contract method 0x6a07973f.
//
// Solidity: function updateNodeType(uint256 _id, uint256 _type) returns()
func (_Prover *ProverTransactorSession) UpdateNodeType(_id *big.Int, _type *big.Int) (*types.Transaction, error) {
	return _Prover.Contract.UpdateNodeType(&_Prover.TransactOpts, _id, _type)
}

// ProverApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Prover contract.
type ProverApprovalIterator struct {
	Event *ProverApproval // Event containing the contract specifics and raw log

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
func (it *ProverApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverApproval)
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
		it.Event = new(ProverApproval)
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
func (it *ProverApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverApproval represents a Approval event raised by the Prover contract.
type ProverApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Prover *ProverFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ProverApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ProverApprovalIterator{contract: _Prover.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Prover *ProverFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ProverApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverApproval)
				if err := _Prover.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Prover *ProverFilterer) ParseApproval(log types.Log) (*ProverApproval, error) {
	event := new(ProverApproval)
	if err := _Prover.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Prover contract.
type ProverApprovalForAllIterator struct {
	Event *ProverApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ProverApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverApprovalForAll)
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
		it.Event = new(ProverApprovalForAll)
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
func (it *ProverApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverApprovalForAll represents a ApprovalForAll event raised by the Prover contract.
type ProverApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Prover *ProverFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ProverApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ProverApprovalForAllIterator{contract: _Prover.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Prover *ProverFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ProverApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverApprovalForAll)
				if err := _Prover.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Prover *ProverFilterer) ParseApprovalForAll(log types.Log) (*ProverApprovalForAll, error) {
	event := new(ProverApprovalForAll)
	if err := _Prover.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Prover contract.
type ProverInitializedIterator struct {
	Event *ProverInitialized // Event containing the contract specifics and raw log

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
func (it *ProverInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverInitialized)
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
		it.Event = new(ProverInitialized)
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
func (it *ProverInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverInitialized represents a Initialized event raised by the Prover contract.
type ProverInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Prover *ProverFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProverInitializedIterator, error) {

	logs, sub, err := _Prover.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProverInitializedIterator{contract: _Prover.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Prover *ProverFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProverInitialized) (event.Subscription, error) {

	logs, sub, err := _Prover.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverInitialized)
				if err := _Prover.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Prover *ProverFilterer) ParseInitialized(log types.Log) (*ProverInitialized, error) {
	event := new(ProverInitialized)
	if err := _Prover.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverMinterSetIterator is returned from FilterMinterSet and is used to iterate over the raw logs and unpacked data for MinterSet events raised by the Prover contract.
type ProverMinterSetIterator struct {
	Event *ProverMinterSet // Event containing the contract specifics and raw log

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
func (it *ProverMinterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverMinterSet)
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
		it.Event = new(ProverMinterSet)
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
func (it *ProverMinterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverMinterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverMinterSet represents a MinterSet event raised by the Prover contract.
type ProverMinterSet struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterSet is a free log retrieval operation binding the contract event 0x726b590ef91a8c76ad05bbe91a57ef84605276528f49cd47d787f558a4e755b6.
//
// Solidity: event MinterSet(address minter)
func (_Prover *ProverFilterer) FilterMinterSet(opts *bind.FilterOpts) (*ProverMinterSetIterator, error) {

	logs, sub, err := _Prover.contract.FilterLogs(opts, "MinterSet")
	if err != nil {
		return nil, err
	}
	return &ProverMinterSetIterator{contract: _Prover.contract, event: "MinterSet", logs: logs, sub: sub}, nil
}

// WatchMinterSet is a free log subscription operation binding the contract event 0x726b590ef91a8c76ad05bbe91a57ef84605276528f49cd47d787f558a4e755b6.
//
// Solidity: event MinterSet(address minter)
func (_Prover *ProverFilterer) WatchMinterSet(opts *bind.WatchOpts, sink chan<- *ProverMinterSet) (event.Subscription, error) {

	logs, sub, err := _Prover.contract.WatchLogs(opts, "MinterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverMinterSet)
				if err := _Prover.contract.UnpackLog(event, "MinterSet", log); err != nil {
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

// ParseMinterSet is a log parse operation binding the contract event 0x726b590ef91a8c76ad05bbe91a57ef84605276528f49cd47d787f558a4e755b6.
//
// Solidity: event MinterSet(address minter)
func (_Prover *ProverFilterer) ParseMinterSet(log types.Log) (*ProverMinterSet, error) {
	event := new(ProverMinterSet)
	if err := _Prover.contract.UnpackLog(event, "MinterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverNodeTypeUpdatedIterator is returned from FilterNodeTypeUpdated and is used to iterate over the raw logs and unpacked data for NodeTypeUpdated events raised by the Prover contract.
type ProverNodeTypeUpdatedIterator struct {
	Event *ProverNodeTypeUpdated // Event containing the contract specifics and raw log

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
func (it *ProverNodeTypeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverNodeTypeUpdated)
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
		it.Event = new(ProverNodeTypeUpdated)
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
func (it *ProverNodeTypeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverNodeTypeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverNodeTypeUpdated represents a NodeTypeUpdated event raised by the Prover contract.
type ProverNodeTypeUpdated struct {
	Id  *big.Int
	Typ *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNodeTypeUpdated is a free log retrieval operation binding the contract event 0x09e65b7abf1ac020ee57d85addfaffc3466fbaa144c58cf8d736f09de55820ab.
//
// Solidity: event NodeTypeUpdated(uint256 indexed id, uint256 typ)
func (_Prover *ProverFilterer) FilterNodeTypeUpdated(opts *bind.FilterOpts, id []*big.Int) (*ProverNodeTypeUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "NodeTypeUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &ProverNodeTypeUpdatedIterator{contract: _Prover.contract, event: "NodeTypeUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeTypeUpdated is a free log subscription operation binding the contract event 0x09e65b7abf1ac020ee57d85addfaffc3466fbaa144c58cf8d736f09de55820ab.
//
// Solidity: event NodeTypeUpdated(uint256 indexed id, uint256 typ)
func (_Prover *ProverFilterer) WatchNodeTypeUpdated(opts *bind.WatchOpts, sink chan<- *ProverNodeTypeUpdated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "NodeTypeUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverNodeTypeUpdated)
				if err := _Prover.contract.UnpackLog(event, "NodeTypeUpdated", log); err != nil {
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

// ParseNodeTypeUpdated is a log parse operation binding the contract event 0x09e65b7abf1ac020ee57d85addfaffc3466fbaa144c58cf8d736f09de55820ab.
//
// Solidity: event NodeTypeUpdated(uint256 indexed id, uint256 typ)
func (_Prover *ProverFilterer) ParseNodeTypeUpdated(log types.Log) (*ProverNodeTypeUpdated, error) {
	event := new(ProverNodeTypeUpdated)
	if err := _Prover.contract.UnpackLog(event, "NodeTypeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverOperatorSetIterator is returned from FilterOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorSet events raised by the Prover contract.
type ProverOperatorSetIterator struct {
	Event *ProverOperatorSet // Event containing the contract specifics and raw log

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
func (it *ProverOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverOperatorSet)
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
		it.Event = new(ProverOperatorSet)
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
func (it *ProverOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverOperatorSet represents a OperatorSet event raised by the Prover contract.
type ProverOperatorSet struct {
	Id       *big.Int
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSet is a free log retrieval operation binding the contract event 0x712369dba77e7931b9ec3bd57319108256b9f79ea5b5255122e3c06117421593.
//
// Solidity: event OperatorSet(uint256 indexed id, address indexed operator)
func (_Prover *ProverFilterer) FilterOperatorSet(opts *bind.FilterOpts, id []*big.Int, operator []common.Address) (*ProverOperatorSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "OperatorSet", idRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ProverOperatorSetIterator{contract: _Prover.contract, event: "OperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSet is a free log subscription operation binding the contract event 0x712369dba77e7931b9ec3bd57319108256b9f79ea5b5255122e3c06117421593.
//
// Solidity: event OperatorSet(uint256 indexed id, address indexed operator)
func (_Prover *ProverFilterer) WatchOperatorSet(opts *bind.WatchOpts, sink chan<- *ProverOperatorSet, id []*big.Int, operator []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "OperatorSet", idRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverOperatorSet)
				if err := _Prover.contract.UnpackLog(event, "OperatorSet", log); err != nil {
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

// ParseOperatorSet is a log parse operation binding the contract event 0x712369dba77e7931b9ec3bd57319108256b9f79ea5b5255122e3c06117421593.
//
// Solidity: event OperatorSet(uint256 indexed id, address indexed operator)
func (_Prover *ProverFilterer) ParseOperatorSet(log types.Log) (*ProverOperatorSet, error) {
	event := new(ProverOperatorSet)
	if err := _Prover.contract.UnpackLog(event, "OperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Prover contract.
type ProverOwnershipTransferredIterator struct {
	Event *ProverOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverOwnershipTransferred)
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
		it.Event = new(ProverOwnershipTransferred)
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
func (it *ProverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverOwnershipTransferred represents a OwnershipTransferred event raised by the Prover contract.
type ProverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Prover *ProverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProverOwnershipTransferredIterator{contract: _Prover.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Prover *ProverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverOwnershipTransferred)
				if err := _Prover.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Prover *ProverFilterer) ParseOwnershipTransferred(log types.Log) (*ProverOwnershipTransferred, error) {
	event := new(ProverOwnershipTransferred)
	if err := _Prover.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverProverPausedIterator is returned from FilterProverPaused and is used to iterate over the raw logs and unpacked data for ProverPaused events raised by the Prover contract.
type ProverProverPausedIterator struct {
	Event *ProverProverPaused // Event containing the contract specifics and raw log

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
func (it *ProverProverPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverProverPaused)
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
		it.Event = new(ProverProverPaused)
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
func (it *ProverProverPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverProverPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverProverPaused represents a ProverPaused event raised by the Prover contract.
type ProverProverPaused struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProverPaused is a free log retrieval operation binding the contract event 0x09c10a851184c6f4c4f912c821413d9b27d48061ecf90d270551f40a23131a88.
//
// Solidity: event ProverPaused(uint256 indexed id)
func (_Prover *ProverFilterer) FilterProverPaused(opts *bind.FilterOpts, id []*big.Int) (*ProverProverPausedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "ProverPaused", idRule)
	if err != nil {
		return nil, err
	}
	return &ProverProverPausedIterator{contract: _Prover.contract, event: "ProverPaused", logs: logs, sub: sub}, nil
}

// WatchProverPaused is a free log subscription operation binding the contract event 0x09c10a851184c6f4c4f912c821413d9b27d48061ecf90d270551f40a23131a88.
//
// Solidity: event ProverPaused(uint256 indexed id)
func (_Prover *ProverFilterer) WatchProverPaused(opts *bind.WatchOpts, sink chan<- *ProverProverPaused, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "ProverPaused", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverProverPaused)
				if err := _Prover.contract.UnpackLog(event, "ProverPaused", log); err != nil {
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

// ParseProverPaused is a log parse operation binding the contract event 0x09c10a851184c6f4c4f912c821413d9b27d48061ecf90d270551f40a23131a88.
//
// Solidity: event ProverPaused(uint256 indexed id)
func (_Prover *ProverFilterer) ParseProverPaused(log types.Log) (*ProverProverPaused, error) {
	event := new(ProverProverPaused)
	if err := _Prover.contract.UnpackLog(event, "ProverPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverProverResumedIterator is returned from FilterProverResumed and is used to iterate over the raw logs and unpacked data for ProverResumed events raised by the Prover contract.
type ProverProverResumedIterator struct {
	Event *ProverProverResumed // Event containing the contract specifics and raw log

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
func (it *ProverProverResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverProverResumed)
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
		it.Event = new(ProverProverResumed)
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
func (it *ProverProverResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverProverResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverProverResumed represents a ProverResumed event raised by the Prover contract.
type ProverProverResumed struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProverResumed is a free log retrieval operation binding the contract event 0xd5c12038aca4e36d3193c55c06f70eee8f829f1165a9e383c70b00d28e3bfdb9.
//
// Solidity: event ProverResumed(uint256 indexed id)
func (_Prover *ProverFilterer) FilterProverResumed(opts *bind.FilterOpts, id []*big.Int) (*ProverProverResumedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "ProverResumed", idRule)
	if err != nil {
		return nil, err
	}
	return &ProverProverResumedIterator{contract: _Prover.contract, event: "ProverResumed", logs: logs, sub: sub}, nil
}

// WatchProverResumed is a free log subscription operation binding the contract event 0xd5c12038aca4e36d3193c55c06f70eee8f829f1165a9e383c70b00d28e3bfdb9.
//
// Solidity: event ProverResumed(uint256 indexed id)
func (_Prover *ProverFilterer) WatchProverResumed(opts *bind.WatchOpts, sink chan<- *ProverProverResumed, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "ProverResumed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverProverResumed)
				if err := _Prover.contract.UnpackLog(event, "ProverResumed", log); err != nil {
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

// ParseProverResumed is a log parse operation binding the contract event 0xd5c12038aca4e36d3193c55c06f70eee8f829f1165a9e383c70b00d28e3bfdb9.
//
// Solidity: event ProverResumed(uint256 indexed id)
func (_Prover *ProverFilterer) ParseProverResumed(log types.Log) (*ProverProverResumed, error) {
	event := new(ProverProverResumed)
	if err := _Prover.contract.UnpackLog(event, "ProverResumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Prover contract.
type ProverTransferIterator struct {
	Event *ProverTransfer // Event containing the contract specifics and raw log

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
func (it *ProverTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverTransfer)
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
		it.Event = new(ProverTransfer)
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
func (it *ProverTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverTransfer represents a Transfer event raised by the Prover contract.
type ProverTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Prover *ProverFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ProverTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ProverTransferIterator{contract: _Prover.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Prover *ProverFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ProverTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverTransfer)
				if err := _Prover.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Prover *ProverFilterer) ParseTransfer(log types.Log) (*ProverTransfer, error) {
	event := new(ProverTransfer)
	if err := _Prover.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
