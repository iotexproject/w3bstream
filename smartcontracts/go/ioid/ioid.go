// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ioid

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

// IoidMetaData contains all meta data concerning the Ioid contract.
var IoidMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"CreateIoID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"RemoveDIDWallet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"SetMinter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"SetResolver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deviceProject\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_device\",\"type\":\"address\"}],\"name\":\"did\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_walletImplementation\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_device\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"projectDeviceCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_start\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pageSize\",\"type\":\"uint256\"}],\"name\":\"projectIDs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"array\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_device\",\"type\":\"address\"}],\"name\":\"removeDID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"wallet_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"did_\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_did\",\"type\":\"string\"}],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IoidABI is the input ABI used to generate the binding from.
// Deprecated: Use IoidMetaData.ABI instead.
var IoidABI = IoidMetaData.ABI

// Ioid is an auto generated Go binding around an Ethereum contract.
type Ioid struct {
	IoidCaller     // Read-only binding to the contract
	IoidTransactor // Write-only binding to the contract
	IoidFilterer   // Log filterer for contract events
}

// IoidCaller is an auto generated read-only Go binding around an Ethereum contract.
type IoidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IoidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IoidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IoidSession struct {
	Contract     *Ioid             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IoidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IoidCallerSession struct {
	Contract *IoidCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IoidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IoidTransactorSession struct {
	Contract     *IoidTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IoidRaw is an auto generated low-level Go binding around an Ethereum contract.
type IoidRaw struct {
	Contract *Ioid // Generic contract binding to access the raw methods on
}

// IoidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IoidCallerRaw struct {
	Contract *IoidCaller // Generic read-only contract binding to access the raw methods on
}

// IoidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IoidTransactorRaw struct {
	Contract *IoidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIoid creates a new instance of Ioid, bound to a specific deployed contract.
func NewIoid(address common.Address, backend bind.ContractBackend) (*Ioid, error) {
	contract, err := bindIoid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ioid{IoidCaller: IoidCaller{contract: contract}, IoidTransactor: IoidTransactor{contract: contract}, IoidFilterer: IoidFilterer{contract: contract}}, nil
}

// NewIoidCaller creates a new read-only instance of Ioid, bound to a specific deployed contract.
func NewIoidCaller(address common.Address, caller bind.ContractCaller) (*IoidCaller, error) {
	contract, err := bindIoid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IoidCaller{contract: contract}, nil
}

// NewIoidTransactor creates a new write-only instance of Ioid, bound to a specific deployed contract.
func NewIoidTransactor(address common.Address, transactor bind.ContractTransactor) (*IoidTransactor, error) {
	contract, err := bindIoid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IoidTransactor{contract: contract}, nil
}

// NewIoidFilterer creates a new log filterer instance of Ioid, bound to a specific deployed contract.
func NewIoidFilterer(address common.Address, filterer bind.ContractFilterer) (*IoidFilterer, error) {
	contract, err := bindIoid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IoidFilterer{contract: contract}, nil
}

// bindIoid binds a generic wrapper to an already deployed contract.
func bindIoid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IoidMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ioid *IoidRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ioid.Contract.IoidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ioid *IoidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ioid.Contract.IoidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ioid *IoidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ioid.Contract.IoidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ioid *IoidCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ioid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ioid *IoidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ioid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ioid *IoidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ioid.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ioid *IoidCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ioid *IoidSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Ioid.Contract.BalanceOf(&_Ioid.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ioid *IoidCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Ioid.Contract.BalanceOf(&_Ioid.CallOpts, owner)
}

// DeviceProject is a free data retrieval call binding the contract method 0x7ba0ef27.
//
// Solidity: function deviceProject(address ) view returns(uint256)
func (_Ioid *IoidCaller) DeviceProject(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "deviceProject", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeviceProject is a free data retrieval call binding the contract method 0x7ba0ef27.
//
// Solidity: function deviceProject(address ) view returns(uint256)
func (_Ioid *IoidSession) DeviceProject(arg0 common.Address) (*big.Int, error) {
	return _Ioid.Contract.DeviceProject(&_Ioid.CallOpts, arg0)
}

// DeviceProject is a free data retrieval call binding the contract method 0x7ba0ef27.
//
// Solidity: function deviceProject(address ) view returns(uint256)
func (_Ioid *IoidCallerSession) DeviceProject(arg0 common.Address) (*big.Int, error) {
	return _Ioid.Contract.DeviceProject(&_Ioid.CallOpts, arg0)
}

// Did is a free data retrieval call binding the contract method 0xb292c335.
//
// Solidity: function did(address _device) view returns(string)
func (_Ioid *IoidCaller) Did(opts *bind.CallOpts, _device common.Address) (string, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "did", _device)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Did is a free data retrieval call binding the contract method 0xb292c335.
//
// Solidity: function did(address _device) view returns(string)
func (_Ioid *IoidSession) Did(_device common.Address) (string, error) {
	return _Ioid.Contract.Did(&_Ioid.CallOpts, _device)
}

// Did is a free data retrieval call binding the contract method 0xb292c335.
//
// Solidity: function did(address _device) view returns(string)
func (_Ioid *IoidCallerSession) Did(_device common.Address) (string, error) {
	return _Ioid.Contract.Did(&_Ioid.CallOpts, _device)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Ioid *IoidCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Ioid *IoidSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Ioid.Contract.GetApproved(&_Ioid.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Ioid *IoidCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Ioid.Contract.GetApproved(&_Ioid.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Ioid *IoidCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Ioid *IoidSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Ioid.Contract.IsApprovedForAll(&_Ioid.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Ioid *IoidCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Ioid.Contract.IsApprovedForAll(&_Ioid.CallOpts, owner, operator)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Ioid *IoidCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Ioid *IoidSession) Minter() (common.Address, error) {
	return _Ioid.Contract.Minter(&_Ioid.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Ioid *IoidCallerSession) Minter() (common.Address, error) {
	return _Ioid.Contract.Minter(&_Ioid.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ioid *IoidCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ioid *IoidSession) Name() (string, error) {
	return _Ioid.Contract.Name(&_Ioid.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ioid *IoidCallerSession) Name() (string, error) {
	return _Ioid.Contract.Name(&_Ioid.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Ioid *IoidCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Ioid *IoidSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Ioid.Contract.OwnerOf(&_Ioid.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Ioid *IoidCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Ioid.Contract.OwnerOf(&_Ioid.CallOpts, tokenId)
}

// ProjectDeviceCount is a free data retrieval call binding the contract method 0xf62ce247.
//
// Solidity: function projectDeviceCount(uint256 ) view returns(uint256)
func (_Ioid *IoidCaller) ProjectDeviceCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "projectDeviceCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectDeviceCount is a free data retrieval call binding the contract method 0xf62ce247.
//
// Solidity: function projectDeviceCount(uint256 ) view returns(uint256)
func (_Ioid *IoidSession) ProjectDeviceCount(arg0 *big.Int) (*big.Int, error) {
	return _Ioid.Contract.ProjectDeviceCount(&_Ioid.CallOpts, arg0)
}

// ProjectDeviceCount is a free data retrieval call binding the contract method 0xf62ce247.
//
// Solidity: function projectDeviceCount(uint256 ) view returns(uint256)
func (_Ioid *IoidCallerSession) ProjectDeviceCount(arg0 *big.Int) (*big.Int, error) {
	return _Ioid.Contract.ProjectDeviceCount(&_Ioid.CallOpts, arg0)
}

// ProjectIDs is a free data retrieval call binding the contract method 0x95f8243a.
//
// Solidity: function projectIDs(uint256 _projectId, address _start, uint256 _pageSize) view returns(address[] array, address next)
func (_Ioid *IoidCaller) ProjectIDs(opts *bind.CallOpts, _projectId *big.Int, _start common.Address, _pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "projectIDs", _projectId, _start, _pageSize)

	outstruct := new(struct {
		Array []common.Address
		Next  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Array = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Next = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ProjectIDs is a free data retrieval call binding the contract method 0x95f8243a.
//
// Solidity: function projectIDs(uint256 _projectId, address _start, uint256 _pageSize) view returns(address[] array, address next)
func (_Ioid *IoidSession) ProjectIDs(_projectId *big.Int, _start common.Address, _pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _Ioid.Contract.ProjectIDs(&_Ioid.CallOpts, _projectId, _start, _pageSize)
}

// ProjectIDs is a free data retrieval call binding the contract method 0x95f8243a.
//
// Solidity: function projectIDs(uint256 _projectId, address _start, uint256 _pageSize) view returns(address[] array, address next)
func (_Ioid *IoidCallerSession) ProjectIDs(_projectId *big.Int, _start common.Address, _pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _Ioid.Contract.ProjectIDs(&_Ioid.CallOpts, _projectId, _start, _pageSize)
}

// Resolver is a free data retrieval call binding the contract method 0x108eaa4e.
//
// Solidity: function resolver(uint256 _id) view returns(address)
func (_Ioid *IoidCaller) Resolver(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "resolver", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x108eaa4e.
//
// Solidity: function resolver(uint256 _id) view returns(address)
func (_Ioid *IoidSession) Resolver(_id *big.Int) (common.Address, error) {
	return _Ioid.Contract.Resolver(&_Ioid.CallOpts, _id)
}

// Resolver is a free data retrieval call binding the contract method 0x108eaa4e.
//
// Solidity: function resolver(uint256 _id) view returns(address)
func (_Ioid *IoidCallerSession) Resolver(_id *big.Int) (common.Address, error) {
	return _Ioid.Contract.Resolver(&_Ioid.CallOpts, _id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ioid *IoidCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ioid *IoidSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ioid.Contract.SupportsInterface(&_Ioid.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ioid *IoidCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ioid.Contract.SupportsInterface(&_Ioid.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ioid *IoidCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ioid *IoidSession) Symbol() (string, error) {
	return _Ioid.Contract.Symbol(&_Ioid.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ioid *IoidCallerSession) Symbol() (string, error) {
	return _Ioid.Contract.Symbol(&_Ioid.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Ioid *IoidCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Ioid *IoidSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Ioid.Contract.TokenByIndex(&_Ioid.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Ioid *IoidCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Ioid.Contract.TokenByIndex(&_Ioid.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Ioid *IoidCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Ioid *IoidSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Ioid.Contract.TokenOfOwnerByIndex(&_Ioid.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Ioid *IoidCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Ioid.Contract.TokenOfOwnerByIndex(&_Ioid.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Ioid *IoidCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Ioid *IoidSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Ioid.Contract.TokenURI(&_Ioid.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Ioid *IoidCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Ioid.Contract.TokenURI(&_Ioid.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ioid *IoidCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ioid *IoidSession) TotalSupply() (*big.Int, error) {
	return _Ioid.Contract.TotalSupply(&_Ioid.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ioid *IoidCallerSession) TotalSupply() (*big.Int, error) {
	return _Ioid.Contract.TotalSupply(&_Ioid.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0xa2781335.
//
// Solidity: function wallet(uint256 _id) view returns(address wallet_, string did_)
func (_Ioid *IoidCaller) Wallet(opts *bind.CallOpts, _id *big.Int) (struct {
	Wallet common.Address
	Did    string
}, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "wallet", _id)

	outstruct := new(struct {
		Wallet common.Address
		Did    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Wallet = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Did = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Wallet is a free data retrieval call binding the contract method 0xa2781335.
//
// Solidity: function wallet(uint256 _id) view returns(address wallet_, string did_)
func (_Ioid *IoidSession) Wallet(_id *big.Int) (struct {
	Wallet common.Address
	Did    string
}, error) {
	return _Ioid.Contract.Wallet(&_Ioid.CallOpts, _id)
}

// Wallet is a free data retrieval call binding the contract method 0xa2781335.
//
// Solidity: function wallet(uint256 _id) view returns(address wallet_, string did_)
func (_Ioid *IoidCallerSession) Wallet(_id *big.Int) (struct {
	Wallet common.Address
	Did    string
}, error) {
	return _Ioid.Contract.Wallet(&_Ioid.CallOpts, _id)
}

// Wallet0 is a free data retrieval call binding the contract method 0xaf0be257.
//
// Solidity: function wallet(string _did) view returns(address)
func (_Ioid *IoidCaller) Wallet0(opts *bind.CallOpts, _did string) (common.Address, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "wallet0", _did)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet0 is a free data retrieval call binding the contract method 0xaf0be257.
//
// Solidity: function wallet(string _did) view returns(address)
func (_Ioid *IoidSession) Wallet0(_did string) (common.Address, error) {
	return _Ioid.Contract.Wallet0(&_Ioid.CallOpts, _did)
}

// Wallet0 is a free data retrieval call binding the contract method 0xaf0be257.
//
// Solidity: function wallet(string _did) view returns(address)
func (_Ioid *IoidCallerSession) Wallet0(_did string) (common.Address, error) {
	return _Ioid.Contract.Wallet0(&_Ioid.CallOpts, _did)
}

// WalletImplementation is a free data retrieval call binding the contract method 0x8117abc1.
//
// Solidity: function walletImplementation() view returns(address)
func (_Ioid *IoidCaller) WalletImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "walletImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletImplementation is a free data retrieval call binding the contract method 0x8117abc1.
//
// Solidity: function walletImplementation() view returns(address)
func (_Ioid *IoidSession) WalletImplementation() (common.Address, error) {
	return _Ioid.Contract.WalletImplementation(&_Ioid.CallOpts)
}

// WalletImplementation is a free data retrieval call binding the contract method 0x8117abc1.
//
// Solidity: function walletImplementation() view returns(address)
func (_Ioid *IoidCallerSession) WalletImplementation() (common.Address, error) {
	return _Ioid.Contract.WalletImplementation(&_Ioid.CallOpts)
}

// WalletRegistry is a free data retrieval call binding the contract method 0xab7aa6ad.
//
// Solidity: function walletRegistry() view returns(address)
func (_Ioid *IoidCaller) WalletRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ioid.contract.Call(opts, &out, "walletRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletRegistry is a free data retrieval call binding the contract method 0xab7aa6ad.
//
// Solidity: function walletRegistry() view returns(address)
func (_Ioid *IoidSession) WalletRegistry() (common.Address, error) {
	return _Ioid.Contract.WalletRegistry(&_Ioid.CallOpts)
}

// WalletRegistry is a free data retrieval call binding the contract method 0xab7aa6ad.
//
// Solidity: function walletRegistry() view returns(address)
func (_Ioid *IoidCallerSession) WalletRegistry() (common.Address, error) {
	return _Ioid.Contract.WalletRegistry(&_Ioid.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Ioid *IoidTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ioid.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Ioid *IoidSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ioid.Contract.Approve(&_Ioid.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Ioid *IoidTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ioid.Contract.Approve(&_Ioid.TransactOpts, to, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address _minter, address _walletRegistry, address _walletImplementation, string _name, string _symbol) returns()
func (_Ioid *IoidTransactor) Initialize(opts *bind.TransactOpts, _minter common.Address, _walletRegistry common.Address, _walletImplementation common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _Ioid.contract.Transact(opts, "initialize", _minter, _walletRegistry, _walletImplementation, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address _minter, address _walletRegistry, address _walletImplementation, string _name, string _symbol) returns()
func (_Ioid *IoidSession) Initialize(_minter common.Address, _walletRegistry common.Address, _walletImplementation common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _Ioid.Contract.Initialize(&_Ioid.TransactOpts, _minter, _walletRegistry, _walletImplementation, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address _minter, address _walletRegistry, address _walletImplementation, string _name, string _symbol) returns()
func (_Ioid *IoidTransactorSession) Initialize(_minter common.Address, _walletRegistry common.Address, _walletImplementation common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _Ioid.Contract.Initialize(&_Ioid.TransactOpts, _minter, _walletRegistry, _walletImplementation, _name, _symbol)
}

// Mint is a paid mutator transaction binding the contract method 0xda39b3e7.
//
// Solidity: function mint(uint256 _projectId, address _device, address _owner) returns(uint256)
func (_Ioid *IoidTransactor) Mint(opts *bind.TransactOpts, _projectId *big.Int, _device common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Ioid.contract.Transact(opts, "mint", _projectId, _device, _owner)
}

// Mint is a paid mutator transaction binding the contract method 0xda39b3e7.
//
// Solidity: function mint(uint256 _projectId, address _device, address _owner) returns(uint256)
func (_Ioid *IoidSession) Mint(_projectId *big.Int, _device common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Ioid.Contract.Mint(&_Ioid.TransactOpts, _projectId, _device, _owner)
}

// Mint is a paid mutator transaction binding the contract method 0xda39b3e7.
//
// Solidity: function mint(uint256 _projectId, address _device, address _owner) returns(uint256)
func (_Ioid *IoidTransactorSession) Mint(_projectId *big.Int, _device common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Ioid.Contract.Mint(&_Ioid.TransactOpts, _projectId, _device, _owner)
}

// RemoveDID is a paid mutator transaction binding the contract method 0x330c5a0e.
//
// Solidity: function removeDID(address _device) returns()
func (_Ioid *IoidTransactor) RemoveDID(opts *bind.TransactOpts, _device common.Address) (*types.Transaction, error) {
	return _Ioid.contract.Transact(opts, "removeDID", _device)
}

// RemoveDID is a paid mutator transaction binding the contract method 0x330c5a0e.
//
// Solidity: function removeDID(address _device) returns()
func (_Ioid *IoidSession) RemoveDID(_device common.Address) (*types.Transaction, error) {
	return _Ioid.Contract.RemoveDID(&_Ioid.TransactOpts, _device)
}

// RemoveDID is a paid mutator transaction binding the contract method 0x330c5a0e.
//
// Solidity: function removeDID(address _device) returns()
func (_Ioid *IoidTransactorSession) RemoveDID(_device common.Address) (*types.Transaction, error) {
	return _Ioid.Contract.RemoveDID(&_Ioid.TransactOpts, _device)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Ioid *IoidTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ioid.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Ioid *IoidSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ioid.Contract.SafeTransferFrom(&_Ioid.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Ioid *IoidTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ioid.Contract.SafeTransferFrom(&_Ioid.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Ioid *IoidTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Ioid.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Ioid *IoidSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Ioid.Contract.SafeTransferFrom0(&_Ioid.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Ioid *IoidTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Ioid.Contract.SafeTransferFrom0(&_Ioid.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Ioid *IoidTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Ioid.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Ioid *IoidSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Ioid.Contract.SetApprovalForAll(&_Ioid.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Ioid *IoidTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Ioid.Contract.SetApprovalForAll(&_Ioid.TransactOpts, operator, approved)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Ioid *IoidTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _Ioid.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Ioid *IoidSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Ioid.Contract.SetMinter(&_Ioid.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Ioid *IoidTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Ioid.Contract.SetMinter(&_Ioid.TransactOpts, _minter)
}

// SetResolver is a paid mutator transaction binding the contract method 0xbc7b6d62.
//
// Solidity: function setResolver(uint256 _id, address _resolver) returns()
func (_Ioid *IoidTransactor) SetResolver(opts *bind.TransactOpts, _id *big.Int, _resolver common.Address) (*types.Transaction, error) {
	return _Ioid.contract.Transact(opts, "setResolver", _id, _resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0xbc7b6d62.
//
// Solidity: function setResolver(uint256 _id, address _resolver) returns()
func (_Ioid *IoidSession) SetResolver(_id *big.Int, _resolver common.Address) (*types.Transaction, error) {
	return _Ioid.Contract.SetResolver(&_Ioid.TransactOpts, _id, _resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0xbc7b6d62.
//
// Solidity: function setResolver(uint256 _id, address _resolver) returns()
func (_Ioid *IoidTransactorSession) SetResolver(_id *big.Int, _resolver common.Address) (*types.Transaction, error) {
	return _Ioid.Contract.SetResolver(&_Ioid.TransactOpts, _id, _resolver)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Ioid *IoidTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ioid.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Ioid *IoidSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ioid.Contract.TransferFrom(&_Ioid.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Ioid *IoidTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ioid.Contract.TransferFrom(&_Ioid.TransactOpts, from, to, tokenId)
}

// IoidApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ioid contract.
type IoidApprovalIterator struct {
	Event *IoidApproval // Event containing the contract specifics and raw log

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
func (it *IoidApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoidApproval)
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
		it.Event = new(IoidApproval)
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
func (it *IoidApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoidApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoidApproval represents a Approval event raised by the Ioid contract.
type IoidApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Ioid *IoidFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IoidApprovalIterator, error) {

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

	logs, sub, err := _Ioid.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IoidApprovalIterator{contract: _Ioid.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Ioid *IoidFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IoidApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Ioid.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoidApproval)
				if err := _Ioid.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Ioid *IoidFilterer) ParseApproval(log types.Log) (*IoidApproval, error) {
	event := new(IoidApproval)
	if err := _Ioid.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoidApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Ioid contract.
type IoidApprovalForAllIterator struct {
	Event *IoidApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IoidApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoidApprovalForAll)
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
		it.Event = new(IoidApprovalForAll)
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
func (it *IoidApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoidApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoidApprovalForAll represents a ApprovalForAll event raised by the Ioid contract.
type IoidApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Ioid *IoidFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IoidApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Ioid.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IoidApprovalForAllIterator{contract: _Ioid.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Ioid *IoidFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IoidApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Ioid.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoidApprovalForAll)
				if err := _Ioid.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Ioid *IoidFilterer) ParseApprovalForAll(log types.Log) (*IoidApprovalForAll, error) {
	event := new(IoidApprovalForAll)
	if err := _Ioid.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoidCreateIoIDIterator is returned from FilterCreateIoID and is used to iterate over the raw logs and unpacked data for CreateIoID events raised by the Ioid contract.
type IoidCreateIoIDIterator struct {
	Event *IoidCreateIoID // Event containing the contract specifics and raw log

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
func (it *IoidCreateIoIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoidCreateIoID)
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
		it.Event = new(IoidCreateIoID)
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
func (it *IoidCreateIoIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoidCreateIoIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoidCreateIoID represents a CreateIoID event raised by the Ioid contract.
type IoidCreateIoID struct {
	Owner  common.Address
	Id     *big.Int
	Wallet common.Address
	Did    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreateIoID is a free log retrieval operation binding the contract event 0x313a15bccdaa3cc35e31f4e2f6a0c398a1c735a9231dd4684399ea0307373062.
//
// Solidity: event CreateIoID(address indexed owner, uint256 id, address wallet, string did)
func (_Ioid *IoidFilterer) FilterCreateIoID(opts *bind.FilterOpts, owner []common.Address) (*IoidCreateIoIDIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ioid.contract.FilterLogs(opts, "CreateIoID", ownerRule)
	if err != nil {
		return nil, err
	}
	return &IoidCreateIoIDIterator{contract: _Ioid.contract, event: "CreateIoID", logs: logs, sub: sub}, nil
}

// WatchCreateIoID is a free log subscription operation binding the contract event 0x313a15bccdaa3cc35e31f4e2f6a0c398a1c735a9231dd4684399ea0307373062.
//
// Solidity: event CreateIoID(address indexed owner, uint256 id, address wallet, string did)
func (_Ioid *IoidFilterer) WatchCreateIoID(opts *bind.WatchOpts, sink chan<- *IoidCreateIoID, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ioid.contract.WatchLogs(opts, "CreateIoID", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoidCreateIoID)
				if err := _Ioid.contract.UnpackLog(event, "CreateIoID", log); err != nil {
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

// ParseCreateIoID is a log parse operation binding the contract event 0x313a15bccdaa3cc35e31f4e2f6a0c398a1c735a9231dd4684399ea0307373062.
//
// Solidity: event CreateIoID(address indexed owner, uint256 id, address wallet, string did)
func (_Ioid *IoidFilterer) ParseCreateIoID(log types.Log) (*IoidCreateIoID, error) {
	event := new(IoidCreateIoID)
	if err := _Ioid.contract.UnpackLog(event, "CreateIoID", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoidInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Ioid contract.
type IoidInitializedIterator struct {
	Event *IoidInitialized // Event containing the contract specifics and raw log

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
func (it *IoidInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoidInitialized)
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
		it.Event = new(IoidInitialized)
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
func (it *IoidInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoidInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoidInitialized represents a Initialized event raised by the Ioid contract.
type IoidInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ioid *IoidFilterer) FilterInitialized(opts *bind.FilterOpts) (*IoidInitializedIterator, error) {

	logs, sub, err := _Ioid.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IoidInitializedIterator{contract: _Ioid.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ioid *IoidFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IoidInitialized) (event.Subscription, error) {

	logs, sub, err := _Ioid.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoidInitialized)
				if err := _Ioid.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Ioid *IoidFilterer) ParseInitialized(log types.Log) (*IoidInitialized, error) {
	event := new(IoidInitialized)
	if err := _Ioid.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoidRemoveDIDWalletIterator is returned from FilterRemoveDIDWallet and is used to iterate over the raw logs and unpacked data for RemoveDIDWallet events raised by the Ioid contract.
type IoidRemoveDIDWalletIterator struct {
	Event *IoidRemoveDIDWallet // Event containing the contract specifics and raw log

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
func (it *IoidRemoveDIDWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoidRemoveDIDWallet)
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
		it.Event = new(IoidRemoveDIDWallet)
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
func (it *IoidRemoveDIDWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoidRemoveDIDWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoidRemoveDIDWallet represents a RemoveDIDWallet event raised by the Ioid contract.
type IoidRemoveDIDWallet struct {
	Wallet common.Address
	Did    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemoveDIDWallet is a free log retrieval operation binding the contract event 0x5be405c75c3aee195a3c337b5ad1503937083fff6e6a6c29f9135252f379c64b.
//
// Solidity: event RemoveDIDWallet(address indexed wallet, string did)
func (_Ioid *IoidFilterer) FilterRemoveDIDWallet(opts *bind.FilterOpts, wallet []common.Address) (*IoidRemoveDIDWalletIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Ioid.contract.FilterLogs(opts, "RemoveDIDWallet", walletRule)
	if err != nil {
		return nil, err
	}
	return &IoidRemoveDIDWalletIterator{contract: _Ioid.contract, event: "RemoveDIDWallet", logs: logs, sub: sub}, nil
}

// WatchRemoveDIDWallet is a free log subscription operation binding the contract event 0x5be405c75c3aee195a3c337b5ad1503937083fff6e6a6c29f9135252f379c64b.
//
// Solidity: event RemoveDIDWallet(address indexed wallet, string did)
func (_Ioid *IoidFilterer) WatchRemoveDIDWallet(opts *bind.WatchOpts, sink chan<- *IoidRemoveDIDWallet, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Ioid.contract.WatchLogs(opts, "RemoveDIDWallet", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoidRemoveDIDWallet)
				if err := _Ioid.contract.UnpackLog(event, "RemoveDIDWallet", log); err != nil {
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

// ParseRemoveDIDWallet is a log parse operation binding the contract event 0x5be405c75c3aee195a3c337b5ad1503937083fff6e6a6c29f9135252f379c64b.
//
// Solidity: event RemoveDIDWallet(address indexed wallet, string did)
func (_Ioid *IoidFilterer) ParseRemoveDIDWallet(log types.Log) (*IoidRemoveDIDWallet, error) {
	event := new(IoidRemoveDIDWallet)
	if err := _Ioid.contract.UnpackLog(event, "RemoveDIDWallet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoidSetMinterIterator is returned from FilterSetMinter and is used to iterate over the raw logs and unpacked data for SetMinter events raised by the Ioid contract.
type IoidSetMinterIterator struct {
	Event *IoidSetMinter // Event containing the contract specifics and raw log

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
func (it *IoidSetMinterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoidSetMinter)
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
		it.Event = new(IoidSetMinter)
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
func (it *IoidSetMinterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoidSetMinterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoidSetMinter represents a SetMinter event raised by the Ioid contract.
type IoidSetMinter struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMinter is a free log retrieval operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address indexed minter)
func (_Ioid *IoidFilterer) FilterSetMinter(opts *bind.FilterOpts, minter []common.Address) (*IoidSetMinterIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Ioid.contract.FilterLogs(opts, "SetMinter", minterRule)
	if err != nil {
		return nil, err
	}
	return &IoidSetMinterIterator{contract: _Ioid.contract, event: "SetMinter", logs: logs, sub: sub}, nil
}

// WatchSetMinter is a free log subscription operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address indexed minter)
func (_Ioid *IoidFilterer) WatchSetMinter(opts *bind.WatchOpts, sink chan<- *IoidSetMinter, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Ioid.contract.WatchLogs(opts, "SetMinter", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoidSetMinter)
				if err := _Ioid.contract.UnpackLog(event, "SetMinter", log); err != nil {
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

// ParseSetMinter is a log parse operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address indexed minter)
func (_Ioid *IoidFilterer) ParseSetMinter(log types.Log) (*IoidSetMinter, error) {
	event := new(IoidSetMinter)
	if err := _Ioid.contract.UnpackLog(event, "SetMinter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoidSetResolverIterator is returned from FilterSetResolver and is used to iterate over the raw logs and unpacked data for SetResolver events raised by the Ioid contract.
type IoidSetResolverIterator struct {
	Event *IoidSetResolver // Event containing the contract specifics and raw log

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
func (it *IoidSetResolverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoidSetResolver)
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
		it.Event = new(IoidSetResolver)
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
func (it *IoidSetResolverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoidSetResolverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoidSetResolver represents a SetResolver event raised by the Ioid contract.
type IoidSetResolver struct {
	Id       *big.Int
	Resolver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetResolver is a free log retrieval operation binding the contract event 0x8566ab8e572a998d53ba054eff8f4d5e3577cf95e3d2a0cb76ee8b2ca84e5d95.
//
// Solidity: event SetResolver(uint256 id, address indexed resolver)
func (_Ioid *IoidFilterer) FilterSetResolver(opts *bind.FilterOpts, resolver []common.Address) (*IoidSetResolverIterator, error) {

	var resolverRule []interface{}
	for _, resolverItem := range resolver {
		resolverRule = append(resolverRule, resolverItem)
	}

	logs, sub, err := _Ioid.contract.FilterLogs(opts, "SetResolver", resolverRule)
	if err != nil {
		return nil, err
	}
	return &IoidSetResolverIterator{contract: _Ioid.contract, event: "SetResolver", logs: logs, sub: sub}, nil
}

// WatchSetResolver is a free log subscription operation binding the contract event 0x8566ab8e572a998d53ba054eff8f4d5e3577cf95e3d2a0cb76ee8b2ca84e5d95.
//
// Solidity: event SetResolver(uint256 id, address indexed resolver)
func (_Ioid *IoidFilterer) WatchSetResolver(opts *bind.WatchOpts, sink chan<- *IoidSetResolver, resolver []common.Address) (event.Subscription, error) {

	var resolverRule []interface{}
	for _, resolverItem := range resolver {
		resolverRule = append(resolverRule, resolverItem)
	}

	logs, sub, err := _Ioid.contract.WatchLogs(opts, "SetResolver", resolverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoidSetResolver)
				if err := _Ioid.contract.UnpackLog(event, "SetResolver", log); err != nil {
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

// ParseSetResolver is a log parse operation binding the contract event 0x8566ab8e572a998d53ba054eff8f4d5e3577cf95e3d2a0cb76ee8b2ca84e5d95.
//
// Solidity: event SetResolver(uint256 id, address indexed resolver)
func (_Ioid *IoidFilterer) ParseSetResolver(log types.Log) (*IoidSetResolver, error) {
	event := new(IoidSetResolver)
	if err := _Ioid.contract.UnpackLog(event, "SetResolver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoidTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ioid contract.
type IoidTransferIterator struct {
	Event *IoidTransfer // Event containing the contract specifics and raw log

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
func (it *IoidTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoidTransfer)
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
		it.Event = new(IoidTransfer)
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
func (it *IoidTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoidTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoidTransfer represents a Transfer event raised by the Ioid contract.
type IoidTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Ioid *IoidFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IoidTransferIterator, error) {

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

	logs, sub, err := _Ioid.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IoidTransferIterator{contract: _Ioid.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Ioid *IoidFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IoidTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Ioid.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoidTransfer)
				if err := _Ioid.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Ioid *IoidFilterer) ParseTransfer(log types.Log) (*IoidTransfer, error) {
	event := new(IoidTransfer)
	if err := _Ioid.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
