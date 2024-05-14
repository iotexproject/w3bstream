// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// IoIDMetaData contains all meta data concerning the IoID contract.
var IoIDMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"CreateIoID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"RemoveDIDWallet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"SetMinter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deviceProject\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_device\",\"type\":\"address\"}],\"name\":\"did\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_walletImplementation\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_device\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"projectDeviceCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_start\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pageSize\",\"type\":\"uint256\"}],\"name\":\"projectIDs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"array\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_device\",\"type\":\"address\"}],\"name\":\"removeDID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"wallet_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"did_\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_did\",\"type\":\"string\"}],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IoIDABI is the input ABI used to generate the binding from.
// Deprecated: Use IoIDMetaData.ABI instead.
var IoIDABI = IoIDMetaData.ABI

// IoID is an auto generated Go binding around an Ethereum contract.
type IoID struct {
	IoIDCaller     // Read-only binding to the contract
	IoIDTransactor // Write-only binding to the contract
	IoIDFilterer   // Log filterer for contract events
}

// IoIDCaller is an auto generated read-only Go binding around an Ethereum contract.
type IoIDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoIDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IoIDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoIDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IoIDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoIDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IoIDSession struct {
	Contract     *IoID             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IoIDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IoIDCallerSession struct {
	Contract *IoIDCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IoIDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IoIDTransactorSession struct {
	Contract     *IoIDTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IoIDRaw is an auto generated low-level Go binding around an Ethereum contract.
type IoIDRaw struct {
	Contract *IoID // Generic contract binding to access the raw methods on
}

// IoIDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IoIDCallerRaw struct {
	Contract *IoIDCaller // Generic read-only contract binding to access the raw methods on
}

// IoIDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IoIDTransactorRaw struct {
	Contract *IoIDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIoID creates a new instance of IoID, bound to a specific deployed contract.
func NewIoID(address common.Address, backend bind.ContractBackend) (*IoID, error) {
	contract, err := bindIoID(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IoID{IoIDCaller: IoIDCaller{contract: contract}, IoIDTransactor: IoIDTransactor{contract: contract}, IoIDFilterer: IoIDFilterer{contract: contract}}, nil
}

// NewIoIDCaller creates a new read-only instance of IoID, bound to a specific deployed contract.
func NewIoIDCaller(address common.Address, caller bind.ContractCaller) (*IoIDCaller, error) {
	contract, err := bindIoID(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IoIDCaller{contract: contract}, nil
}

// NewIoIDTransactor creates a new write-only instance of IoID, bound to a specific deployed contract.
func NewIoIDTransactor(address common.Address, transactor bind.ContractTransactor) (*IoIDTransactor, error) {
	contract, err := bindIoID(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IoIDTransactor{contract: contract}, nil
}

// NewIoIDFilterer creates a new log filterer instance of IoID, bound to a specific deployed contract.
func NewIoIDFilterer(address common.Address, filterer bind.ContractFilterer) (*IoIDFilterer, error) {
	contract, err := bindIoID(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IoIDFilterer{contract: contract}, nil
}

// bindIoID binds a generic wrapper to an already deployed contract.
func bindIoID(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IoIDMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoID *IoIDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IoID.Contract.IoIDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoID *IoIDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoID.Contract.IoIDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoID *IoIDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoID.Contract.IoIDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoID *IoIDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IoID.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoID *IoIDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoID.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoID *IoIDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoID.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IoID *IoIDCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IoID *IoIDSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IoID.Contract.BalanceOf(&_IoID.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IoID *IoIDCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IoID.Contract.BalanceOf(&_IoID.CallOpts, owner)
}

// DeviceProject is a free data retrieval call binding the contract method 0x7ba0ef27.
//
// Solidity: function deviceProject(address ) view returns(uint256)
func (_IoID *IoIDCaller) DeviceProject(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "deviceProject", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeviceProject is a free data retrieval call binding the contract method 0x7ba0ef27.
//
// Solidity: function deviceProject(address ) view returns(uint256)
func (_IoID *IoIDSession) DeviceProject(arg0 common.Address) (*big.Int, error) {
	return _IoID.Contract.DeviceProject(&_IoID.CallOpts, arg0)
}

// DeviceProject is a free data retrieval call binding the contract method 0x7ba0ef27.
//
// Solidity: function deviceProject(address ) view returns(uint256)
func (_IoID *IoIDCallerSession) DeviceProject(arg0 common.Address) (*big.Int, error) {
	return _IoID.Contract.DeviceProject(&_IoID.CallOpts, arg0)
}

// Did is a free data retrieval call binding the contract method 0xb292c335.
//
// Solidity: function did(address _device) view returns(string)
func (_IoID *IoIDCaller) Did(opts *bind.CallOpts, _device common.Address) (string, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "did", _device)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Did is a free data retrieval call binding the contract method 0xb292c335.
//
// Solidity: function did(address _device) view returns(string)
func (_IoID *IoIDSession) Did(_device common.Address) (string, error) {
	return _IoID.Contract.Did(&_IoID.CallOpts, _device)
}

// Did is a free data retrieval call binding the contract method 0xb292c335.
//
// Solidity: function did(address _device) view returns(string)
func (_IoID *IoIDCallerSession) Did(_device common.Address) (string, error) {
	return _IoID.Contract.Did(&_IoID.CallOpts, _device)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_IoID *IoIDCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_IoID *IoIDSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IoID.Contract.GetApproved(&_IoID.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_IoID *IoIDCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IoID.Contract.GetApproved(&_IoID.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IoID *IoIDCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IoID *IoIDSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IoID.Contract.IsApprovedForAll(&_IoID.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IoID *IoIDCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IoID.Contract.IsApprovedForAll(&_IoID.CallOpts, owner, operator)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_IoID *IoIDCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_IoID *IoIDSession) Minter() (common.Address, error) {
	return _IoID.Contract.Minter(&_IoID.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_IoID *IoIDCallerSession) Minter() (common.Address, error) {
	return _IoID.Contract.Minter(&_IoID.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IoID *IoIDCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IoID *IoIDSession) Name() (string, error) {
	return _IoID.Contract.Name(&_IoID.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IoID *IoIDCallerSession) Name() (string, error) {
	return _IoID.Contract.Name(&_IoID.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IoID *IoIDCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IoID *IoIDSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IoID.Contract.OwnerOf(&_IoID.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IoID *IoIDCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IoID.Contract.OwnerOf(&_IoID.CallOpts, tokenId)
}

// ProjectDeviceCount is a free data retrieval call binding the contract method 0xf62ce247.
//
// Solidity: function projectDeviceCount(uint256 ) view returns(uint256)
func (_IoID *IoIDCaller) ProjectDeviceCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "projectDeviceCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectDeviceCount is a free data retrieval call binding the contract method 0xf62ce247.
//
// Solidity: function projectDeviceCount(uint256 ) view returns(uint256)
func (_IoID *IoIDSession) ProjectDeviceCount(arg0 *big.Int) (*big.Int, error) {
	return _IoID.Contract.ProjectDeviceCount(&_IoID.CallOpts, arg0)
}

// ProjectDeviceCount is a free data retrieval call binding the contract method 0xf62ce247.
//
// Solidity: function projectDeviceCount(uint256 ) view returns(uint256)
func (_IoID *IoIDCallerSession) ProjectDeviceCount(arg0 *big.Int) (*big.Int, error) {
	return _IoID.Contract.ProjectDeviceCount(&_IoID.CallOpts, arg0)
}

// ProjectIDs is a free data retrieval call binding the contract method 0x95f8243a.
//
// Solidity: function projectIDs(uint256 _projectId, address _start, uint256 _pageSize) view returns(address[] array, address next)
func (_IoID *IoIDCaller) ProjectIDs(opts *bind.CallOpts, _projectId *big.Int, _start common.Address, _pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "projectIDs", _projectId, _start, _pageSize)

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
func (_IoID *IoIDSession) ProjectIDs(_projectId *big.Int, _start common.Address, _pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _IoID.Contract.ProjectIDs(&_IoID.CallOpts, _projectId, _start, _pageSize)
}

// ProjectIDs is a free data retrieval call binding the contract method 0x95f8243a.
//
// Solidity: function projectIDs(uint256 _projectId, address _start, uint256 _pageSize) view returns(address[] array, address next)
func (_IoID *IoIDCallerSession) ProjectIDs(_projectId *big.Int, _start common.Address, _pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _IoID.Contract.ProjectIDs(&_IoID.CallOpts, _projectId, _start, _pageSize)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IoID *IoIDCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IoID *IoIDSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IoID.Contract.SupportsInterface(&_IoID.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IoID *IoIDCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IoID.Contract.SupportsInterface(&_IoID.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IoID *IoIDCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IoID *IoIDSession) Symbol() (string, error) {
	return _IoID.Contract.Symbol(&_IoID.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IoID *IoIDCallerSession) Symbol() (string, error) {
	return _IoID.Contract.Symbol(&_IoID.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IoID *IoIDCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IoID *IoIDSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _IoID.Contract.TokenByIndex(&_IoID.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IoID *IoIDCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _IoID.Contract.TokenByIndex(&_IoID.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_IoID *IoIDCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_IoID *IoIDSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IoID.Contract.TokenOfOwnerByIndex(&_IoID.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_IoID *IoIDCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IoID.Contract.TokenOfOwnerByIndex(&_IoID.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IoID *IoIDCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IoID *IoIDSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IoID.Contract.TokenURI(&_IoID.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IoID *IoIDCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IoID.Contract.TokenURI(&_IoID.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IoID *IoIDCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IoID *IoIDSession) TotalSupply() (*big.Int, error) {
	return _IoID.Contract.TotalSupply(&_IoID.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IoID *IoIDCallerSession) TotalSupply() (*big.Int, error) {
	return _IoID.Contract.TotalSupply(&_IoID.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0xa2781335.
//
// Solidity: function wallet(uint256 _id) view returns(address wallet_, string did_)
func (_IoID *IoIDCaller) Wallet(opts *bind.CallOpts, _id *big.Int) (struct {
	Wallet common.Address
	Did    string
}, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "wallet", _id)

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
func (_IoID *IoIDSession) Wallet(_id *big.Int) (struct {
	Wallet common.Address
	Did    string
}, error) {
	return _IoID.Contract.Wallet(&_IoID.CallOpts, _id)
}

// Wallet is a free data retrieval call binding the contract method 0xa2781335.
//
// Solidity: function wallet(uint256 _id) view returns(address wallet_, string did_)
func (_IoID *IoIDCallerSession) Wallet(_id *big.Int) (struct {
	Wallet common.Address
	Did    string
}, error) {
	return _IoID.Contract.Wallet(&_IoID.CallOpts, _id)
}

// Wallet0 is a free data retrieval call binding the contract method 0xaf0be257.
//
// Solidity: function wallet(string _did) view returns(address)
func (_IoID *IoIDCaller) Wallet0(opts *bind.CallOpts, _did string) (common.Address, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "wallet0", _did)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet0 is a free data retrieval call binding the contract method 0xaf0be257.
//
// Solidity: function wallet(string _did) view returns(address)
func (_IoID *IoIDSession) Wallet0(_did string) (common.Address, error) {
	return _IoID.Contract.Wallet0(&_IoID.CallOpts, _did)
}

// Wallet0 is a free data retrieval call binding the contract method 0xaf0be257.
//
// Solidity: function wallet(string _did) view returns(address)
func (_IoID *IoIDCallerSession) Wallet0(_did string) (common.Address, error) {
	return _IoID.Contract.Wallet0(&_IoID.CallOpts, _did)
}

// WalletImplementation is a free data retrieval call binding the contract method 0x8117abc1.
//
// Solidity: function walletImplementation() view returns(address)
func (_IoID *IoIDCaller) WalletImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "walletImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletImplementation is a free data retrieval call binding the contract method 0x8117abc1.
//
// Solidity: function walletImplementation() view returns(address)
func (_IoID *IoIDSession) WalletImplementation() (common.Address, error) {
	return _IoID.Contract.WalletImplementation(&_IoID.CallOpts)
}

// WalletImplementation is a free data retrieval call binding the contract method 0x8117abc1.
//
// Solidity: function walletImplementation() view returns(address)
func (_IoID *IoIDCallerSession) WalletImplementation() (common.Address, error) {
	return _IoID.Contract.WalletImplementation(&_IoID.CallOpts)
}

// WalletRegistry is a free data retrieval call binding the contract method 0xab7aa6ad.
//
// Solidity: function walletRegistry() view returns(address)
func (_IoID *IoIDCaller) WalletRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IoID.contract.Call(opts, &out, "walletRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletRegistry is a free data retrieval call binding the contract method 0xab7aa6ad.
//
// Solidity: function walletRegistry() view returns(address)
func (_IoID *IoIDSession) WalletRegistry() (common.Address, error) {
	return _IoID.Contract.WalletRegistry(&_IoID.CallOpts)
}

// WalletRegistry is a free data retrieval call binding the contract method 0xab7aa6ad.
//
// Solidity: function walletRegistry() view returns(address)
func (_IoID *IoIDCallerSession) WalletRegistry() (common.Address, error) {
	return _IoID.Contract.WalletRegistry(&_IoID.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IoID *IoIDTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IoID.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IoID *IoIDSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IoID.Contract.Approve(&_IoID.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IoID *IoIDTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IoID.Contract.Approve(&_IoID.TransactOpts, to, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address _minter, address _walletRegistry, address _walletImplementation, string _name, string _symbol) returns()
func (_IoID *IoIDTransactor) Initialize(opts *bind.TransactOpts, _minter common.Address, _walletRegistry common.Address, _walletImplementation common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _IoID.contract.Transact(opts, "initialize", _minter, _walletRegistry, _walletImplementation, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address _minter, address _walletRegistry, address _walletImplementation, string _name, string _symbol) returns()
func (_IoID *IoIDSession) Initialize(_minter common.Address, _walletRegistry common.Address, _walletImplementation common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _IoID.Contract.Initialize(&_IoID.TransactOpts, _minter, _walletRegistry, _walletImplementation, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address _minter, address _walletRegistry, address _walletImplementation, string _name, string _symbol) returns()
func (_IoID *IoIDTransactorSession) Initialize(_minter common.Address, _walletRegistry common.Address, _walletImplementation common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _IoID.Contract.Initialize(&_IoID.TransactOpts, _minter, _walletRegistry, _walletImplementation, _name, _symbol)
}

// Mint is a paid mutator transaction binding the contract method 0xda39b3e7.
//
// Solidity: function mint(uint256 _projectId, address _device, address _owner) returns(uint256)
func (_IoID *IoIDTransactor) Mint(opts *bind.TransactOpts, _projectId *big.Int, _device common.Address, _owner common.Address) (*types.Transaction, error) {
	return _IoID.contract.Transact(opts, "mint", _projectId, _device, _owner)
}

// Mint is a paid mutator transaction binding the contract method 0xda39b3e7.
//
// Solidity: function mint(uint256 _projectId, address _device, address _owner) returns(uint256)
func (_IoID *IoIDSession) Mint(_projectId *big.Int, _device common.Address, _owner common.Address) (*types.Transaction, error) {
	return _IoID.Contract.Mint(&_IoID.TransactOpts, _projectId, _device, _owner)
}

// Mint is a paid mutator transaction binding the contract method 0xda39b3e7.
//
// Solidity: function mint(uint256 _projectId, address _device, address _owner) returns(uint256)
func (_IoID *IoIDTransactorSession) Mint(_projectId *big.Int, _device common.Address, _owner common.Address) (*types.Transaction, error) {
	return _IoID.Contract.Mint(&_IoID.TransactOpts, _projectId, _device, _owner)
}

// RemoveDID is a paid mutator transaction binding the contract method 0x330c5a0e.
//
// Solidity: function removeDID(address _device) returns()
func (_IoID *IoIDTransactor) RemoveDID(opts *bind.TransactOpts, _device common.Address) (*types.Transaction, error) {
	return _IoID.contract.Transact(opts, "removeDID", _device)
}

// RemoveDID is a paid mutator transaction binding the contract method 0x330c5a0e.
//
// Solidity: function removeDID(address _device) returns()
func (_IoID *IoIDSession) RemoveDID(_device common.Address) (*types.Transaction, error) {
	return _IoID.Contract.RemoveDID(&_IoID.TransactOpts, _device)
}

// RemoveDID is a paid mutator transaction binding the contract method 0x330c5a0e.
//
// Solidity: function removeDID(address _device) returns()
func (_IoID *IoIDTransactorSession) RemoveDID(_device common.Address) (*types.Transaction, error) {
	return _IoID.Contract.RemoveDID(&_IoID.TransactOpts, _device)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IoID *IoIDTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IoID.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IoID *IoIDSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IoID.Contract.SafeTransferFrom(&_IoID.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IoID *IoIDTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IoID.Contract.SafeTransferFrom(&_IoID.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IoID *IoIDTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IoID.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IoID *IoIDSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IoID.Contract.SafeTransferFrom0(&_IoID.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IoID *IoIDTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IoID.Contract.SafeTransferFrom0(&_IoID.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IoID *IoIDTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _IoID.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IoID *IoIDSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _IoID.Contract.SetApprovalForAll(&_IoID.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IoID *IoIDTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _IoID.Contract.SetApprovalForAll(&_IoID.TransactOpts, operator, approved)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_IoID *IoIDTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _IoID.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_IoID *IoIDSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _IoID.Contract.SetMinter(&_IoID.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_IoID *IoIDTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _IoID.Contract.SetMinter(&_IoID.TransactOpts, _minter)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IoID *IoIDTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IoID.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IoID *IoIDSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IoID.Contract.TransferFrom(&_IoID.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IoID *IoIDTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IoID.Contract.TransferFrom(&_IoID.TransactOpts, from, to, tokenId)
}

// IoIDApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IoID contract.
type IoIDApprovalIterator struct {
	Event *IoIDApproval // Event containing the contract specifics and raw log

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
func (it *IoIDApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoIDApproval)
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
		it.Event = new(IoIDApproval)
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
func (it *IoIDApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoIDApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoIDApproval represents a Approval event raised by the IoID contract.
type IoIDApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IoID *IoIDFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IoIDApprovalIterator, error) {

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

	logs, sub, err := _IoID.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IoIDApprovalIterator{contract: _IoID.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IoID *IoIDFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IoIDApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IoID.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoIDApproval)
				if err := _IoID.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IoID *IoIDFilterer) ParseApproval(log types.Log) (*IoIDApproval, error) {
	event := new(IoIDApproval)
	if err := _IoID.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoIDApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IoID contract.
type IoIDApprovalForAllIterator struct {
	Event *IoIDApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IoIDApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoIDApprovalForAll)
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
		it.Event = new(IoIDApprovalForAll)
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
func (it *IoIDApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoIDApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoIDApprovalForAll represents a ApprovalForAll event raised by the IoID contract.
type IoIDApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IoID *IoIDFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IoIDApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IoID.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IoIDApprovalForAllIterator{contract: _IoID.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IoID *IoIDFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IoIDApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IoID.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoIDApprovalForAll)
				if err := _IoID.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_IoID *IoIDFilterer) ParseApprovalForAll(log types.Log) (*IoIDApprovalForAll, error) {
	event := new(IoIDApprovalForAll)
	if err := _IoID.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoIDCreateIoIDIterator is returned from FilterCreateIoID and is used to iterate over the raw logs and unpacked data for CreateIoID events raised by the IoID contract.
type IoIDCreateIoIDIterator struct {
	Event *IoIDCreateIoID // Event containing the contract specifics and raw log

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
func (it *IoIDCreateIoIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoIDCreateIoID)
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
		it.Event = new(IoIDCreateIoID)
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
func (it *IoIDCreateIoIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoIDCreateIoIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoIDCreateIoID represents a CreateIoID event raised by the IoID contract.
type IoIDCreateIoID struct {
	Owner  common.Address
	Id     *big.Int
	Wallet common.Address
	Did    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreateIoID is a free log retrieval operation binding the contract event 0x313a15bccdaa3cc35e31f4e2f6a0c398a1c735a9231dd4684399ea0307373062.
//
// Solidity: event CreateIoID(address indexed owner, uint256 id, address wallet, string did)
func (_IoID *IoIDFilterer) FilterCreateIoID(opts *bind.FilterOpts, owner []common.Address) (*IoIDCreateIoIDIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IoID.contract.FilterLogs(opts, "CreateIoID", ownerRule)
	if err != nil {
		return nil, err
	}
	return &IoIDCreateIoIDIterator{contract: _IoID.contract, event: "CreateIoID", logs: logs, sub: sub}, nil
}

// WatchCreateIoID is a free log subscription operation binding the contract event 0x313a15bccdaa3cc35e31f4e2f6a0c398a1c735a9231dd4684399ea0307373062.
//
// Solidity: event CreateIoID(address indexed owner, uint256 id, address wallet, string did)
func (_IoID *IoIDFilterer) WatchCreateIoID(opts *bind.WatchOpts, sink chan<- *IoIDCreateIoID, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IoID.contract.WatchLogs(opts, "CreateIoID", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoIDCreateIoID)
				if err := _IoID.contract.UnpackLog(event, "CreateIoID", log); err != nil {
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
func (_IoID *IoIDFilterer) ParseCreateIoID(log types.Log) (*IoIDCreateIoID, error) {
	event := new(IoIDCreateIoID)
	if err := _IoID.contract.UnpackLog(event, "CreateIoID", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoIDInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IoID contract.
type IoIDInitializedIterator struct {
	Event *IoIDInitialized // Event containing the contract specifics and raw log

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
func (it *IoIDInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoIDInitialized)
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
		it.Event = new(IoIDInitialized)
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
func (it *IoIDInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoIDInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoIDInitialized represents a Initialized event raised by the IoID contract.
type IoIDInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IoID *IoIDFilterer) FilterInitialized(opts *bind.FilterOpts) (*IoIDInitializedIterator, error) {

	logs, sub, err := _IoID.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IoIDInitializedIterator{contract: _IoID.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IoID *IoIDFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IoIDInitialized) (event.Subscription, error) {

	logs, sub, err := _IoID.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoIDInitialized)
				if err := _IoID.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_IoID *IoIDFilterer) ParseInitialized(log types.Log) (*IoIDInitialized, error) {
	event := new(IoIDInitialized)
	if err := _IoID.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoIDRemoveDIDWalletIterator is returned from FilterRemoveDIDWallet and is used to iterate over the raw logs and unpacked data for RemoveDIDWallet events raised by the IoID contract.
type IoIDRemoveDIDWalletIterator struct {
	Event *IoIDRemoveDIDWallet // Event containing the contract specifics and raw log

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
func (it *IoIDRemoveDIDWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoIDRemoveDIDWallet)
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
		it.Event = new(IoIDRemoveDIDWallet)
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
func (it *IoIDRemoveDIDWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoIDRemoveDIDWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoIDRemoveDIDWallet represents a RemoveDIDWallet event raised by the IoID contract.
type IoIDRemoveDIDWallet struct {
	Wallet common.Address
	Did    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemoveDIDWallet is a free log retrieval operation binding the contract event 0x5be405c75c3aee195a3c337b5ad1503937083fff6e6a6c29f9135252f379c64b.
//
// Solidity: event RemoveDIDWallet(address indexed wallet, string did)
func (_IoID *IoIDFilterer) FilterRemoveDIDWallet(opts *bind.FilterOpts, wallet []common.Address) (*IoIDRemoveDIDWalletIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _IoID.contract.FilterLogs(opts, "RemoveDIDWallet", walletRule)
	if err != nil {
		return nil, err
	}
	return &IoIDRemoveDIDWalletIterator{contract: _IoID.contract, event: "RemoveDIDWallet", logs: logs, sub: sub}, nil
}

// WatchRemoveDIDWallet is a free log subscription operation binding the contract event 0x5be405c75c3aee195a3c337b5ad1503937083fff6e6a6c29f9135252f379c64b.
//
// Solidity: event RemoveDIDWallet(address indexed wallet, string did)
func (_IoID *IoIDFilterer) WatchRemoveDIDWallet(opts *bind.WatchOpts, sink chan<- *IoIDRemoveDIDWallet, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _IoID.contract.WatchLogs(opts, "RemoveDIDWallet", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoIDRemoveDIDWallet)
				if err := _IoID.contract.UnpackLog(event, "RemoveDIDWallet", log); err != nil {
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
func (_IoID *IoIDFilterer) ParseRemoveDIDWallet(log types.Log) (*IoIDRemoveDIDWallet, error) {
	event := new(IoIDRemoveDIDWallet)
	if err := _IoID.contract.UnpackLog(event, "RemoveDIDWallet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoIDSetMinterIterator is returned from FilterSetMinter and is used to iterate over the raw logs and unpacked data for SetMinter events raised by the IoID contract.
type IoIDSetMinterIterator struct {
	Event *IoIDSetMinter // Event containing the contract specifics and raw log

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
func (it *IoIDSetMinterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoIDSetMinter)
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
		it.Event = new(IoIDSetMinter)
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
func (it *IoIDSetMinterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoIDSetMinterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoIDSetMinter represents a SetMinter event raised by the IoID contract.
type IoIDSetMinter struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMinter is a free log retrieval operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address indexed minter)
func (_IoID *IoIDFilterer) FilterSetMinter(opts *bind.FilterOpts, minter []common.Address) (*IoIDSetMinterIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _IoID.contract.FilterLogs(opts, "SetMinter", minterRule)
	if err != nil {
		return nil, err
	}
	return &IoIDSetMinterIterator{contract: _IoID.contract, event: "SetMinter", logs: logs, sub: sub}, nil
}

// WatchSetMinter is a free log subscription operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address indexed minter)
func (_IoID *IoIDFilterer) WatchSetMinter(opts *bind.WatchOpts, sink chan<- *IoIDSetMinter, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _IoID.contract.WatchLogs(opts, "SetMinter", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoIDSetMinter)
				if err := _IoID.contract.UnpackLog(event, "SetMinter", log); err != nil {
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
func (_IoID *IoIDFilterer) ParseSetMinter(log types.Log) (*IoIDSetMinter, error) {
	event := new(IoIDSetMinter)
	if err := _IoID.contract.UnpackLog(event, "SetMinter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoIDTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IoID contract.
type IoIDTransferIterator struct {
	Event *IoIDTransfer // Event containing the contract specifics and raw log

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
func (it *IoIDTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoIDTransfer)
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
		it.Event = new(IoIDTransfer)
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
func (it *IoIDTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoIDTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoIDTransfer represents a Transfer event raised by the IoID contract.
type IoIDTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IoID *IoIDFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IoIDTransferIterator, error) {

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

	logs, sub, err := _IoID.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IoIDTransferIterator{contract: _IoID.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IoID *IoIDFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IoIDTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IoID.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoIDTransfer)
				if err := _IoID.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IoID *IoIDFilterer) ParseTransfer(log types.Log) (*IoIDTransfer, error) {
	event := new(IoIDTransfer)
	if err := _IoID.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
