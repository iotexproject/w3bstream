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

// IoIDRegistryMetaData contains all meta data concerning the IoIDRegistry contract.
var IoIDRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"NewDevice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"RemoveDevice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"store\",\"type\":\"address\"}],\"name\":\"SetIoIdStore\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"UpdateDevice\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"METHOD\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"deviceTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"documentHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"documentID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"documentURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ioIDStore\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ioID\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ioID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ioIDStore\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"permitHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deviceContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registeredNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ioIDStore\",\"type\":\"address\"}],\"name\":\"setIoIDStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IoIDRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IoIDRegistryMetaData.ABI instead.
var IoIDRegistryABI = IoIDRegistryMetaData.ABI

// IoIDRegistry is an auto generated Go binding around an Ethereum contract.
type IoIDRegistry struct {
	IoIDRegistryCaller     // Read-only binding to the contract
	IoIDRegistryTransactor // Write-only binding to the contract
	IoIDRegistryFilterer   // Log filterer for contract events
}

// IoIDRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IoIDRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoIDRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IoIDRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoIDRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IoIDRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoIDRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IoIDRegistrySession struct {
	Contract     *IoIDRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IoIDRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IoIDRegistryCallerSession struct {
	Contract *IoIDRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IoIDRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IoIDRegistryTransactorSession struct {
	Contract     *IoIDRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IoIDRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IoIDRegistryRaw struct {
	Contract *IoIDRegistry // Generic contract binding to access the raw methods on
}

// IoIDRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IoIDRegistryCallerRaw struct {
	Contract *IoIDRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IoIDRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IoIDRegistryTransactorRaw struct {
	Contract *IoIDRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIoIDRegistry creates a new instance of IoIDRegistry, bound to a specific deployed contract.
func NewIoIDRegistry(address common.Address, backend bind.ContractBackend) (*IoIDRegistry, error) {
	contract, err := bindIoIDRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IoIDRegistry{IoIDRegistryCaller: IoIDRegistryCaller{contract: contract}, IoIDRegistryTransactor: IoIDRegistryTransactor{contract: contract}, IoIDRegistryFilterer: IoIDRegistryFilterer{contract: contract}}, nil
}

// NewIoIDRegistryCaller creates a new read-only instance of IoIDRegistry, bound to a specific deployed contract.
func NewIoIDRegistryCaller(address common.Address, caller bind.ContractCaller) (*IoIDRegistryCaller, error) {
	contract, err := bindIoIDRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IoIDRegistryCaller{contract: contract}, nil
}

// NewIoIDRegistryTransactor creates a new write-only instance of IoIDRegistry, bound to a specific deployed contract.
func NewIoIDRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IoIDRegistryTransactor, error) {
	contract, err := bindIoIDRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IoIDRegistryTransactor{contract: contract}, nil
}

// NewIoIDRegistryFilterer creates a new log filterer instance of IoIDRegistry, bound to a specific deployed contract.
func NewIoIDRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IoIDRegistryFilterer, error) {
	contract, err := bindIoIDRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IoIDRegistryFilterer{contract: contract}, nil
}

// bindIoIDRegistry binds a generic wrapper to an already deployed contract.
func bindIoIDRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IoIDRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoIDRegistry *IoIDRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IoIDRegistry.Contract.IoIDRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoIDRegistry *IoIDRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoIDRegistry.Contract.IoIDRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoIDRegistry *IoIDRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoIDRegistry.Contract.IoIDRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoIDRegistry *IoIDRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IoIDRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoIDRegistry *IoIDRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoIDRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoIDRegistry *IoIDRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoIDRegistry.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IoIDRegistry *IoIDRegistryCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IoIDRegistry.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IoIDRegistry *IoIDRegistrySession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IoIDRegistry.Contract.DOMAINSEPARATOR(&_IoIDRegistry.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IoIDRegistry *IoIDRegistryCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IoIDRegistry.Contract.DOMAINSEPARATOR(&_IoIDRegistry.CallOpts)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc49f91d3.
//
// Solidity: function EIP712DOMAIN_TYPEHASH() view returns(bytes32)
func (_IoIDRegistry *IoIDRegistryCaller) EIP712DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IoIDRegistry.contract.Call(opts, &out, "EIP712DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc49f91d3.
//
// Solidity: function EIP712DOMAIN_TYPEHASH() view returns(bytes32)
func (_IoIDRegistry *IoIDRegistrySession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _IoIDRegistry.Contract.EIP712DOMAINTYPEHASH(&_IoIDRegistry.CallOpts)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc49f91d3.
//
// Solidity: function EIP712DOMAIN_TYPEHASH() view returns(bytes32)
func (_IoIDRegistry *IoIDRegistryCallerSession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _IoIDRegistry.Contract.EIP712DOMAINTYPEHASH(&_IoIDRegistry.CallOpts)
}

// METHOD is a free data retrieval call binding the contract method 0x67efe82e.
//
// Solidity: function METHOD() view returns(string)
func (_IoIDRegistry *IoIDRegistryCaller) METHOD(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IoIDRegistry.contract.Call(opts, &out, "METHOD")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// METHOD is a free data retrieval call binding the contract method 0x67efe82e.
//
// Solidity: function METHOD() view returns(string)
func (_IoIDRegistry *IoIDRegistrySession) METHOD() (string, error) {
	return _IoIDRegistry.Contract.METHOD(&_IoIDRegistry.CallOpts)
}

// METHOD is a free data retrieval call binding the contract method 0x67efe82e.
//
// Solidity: function METHOD() view returns(string)
func (_IoIDRegistry *IoIDRegistryCallerSession) METHOD() (string, error) {
	return _IoIDRegistry.Contract.METHOD(&_IoIDRegistry.CallOpts)
}

// DeviceTokenId is a free data retrieval call binding the contract method 0x4965c7e1.
//
// Solidity: function deviceTokenId(address device) view returns(uint256)
func (_IoIDRegistry *IoIDRegistryCaller) DeviceTokenId(opts *bind.CallOpts, device common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IoIDRegistry.contract.Call(opts, &out, "deviceTokenId", device)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeviceTokenId is a free data retrieval call binding the contract method 0x4965c7e1.
//
// Solidity: function deviceTokenId(address device) view returns(uint256)
func (_IoIDRegistry *IoIDRegistrySession) DeviceTokenId(device common.Address) (*big.Int, error) {
	return _IoIDRegistry.Contract.DeviceTokenId(&_IoIDRegistry.CallOpts, device)
}

// DeviceTokenId is a free data retrieval call binding the contract method 0x4965c7e1.
//
// Solidity: function deviceTokenId(address device) view returns(uint256)
func (_IoIDRegistry *IoIDRegistryCallerSession) DeviceTokenId(device common.Address) (*big.Int, error) {
	return _IoIDRegistry.Contract.DeviceTokenId(&_IoIDRegistry.CallOpts, device)
}

// DocumentHash is a free data retrieval call binding the contract method 0x35984678.
//
// Solidity: function documentHash(address device) view returns(bytes32)
func (_IoIDRegistry *IoIDRegistryCaller) DocumentHash(opts *bind.CallOpts, device common.Address) ([32]byte, error) {
	var out []interface{}
	err := _IoIDRegistry.contract.Call(opts, &out, "documentHash", device)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DocumentHash is a free data retrieval call binding the contract method 0x35984678.
//
// Solidity: function documentHash(address device) view returns(bytes32)
func (_IoIDRegistry *IoIDRegistrySession) DocumentHash(device common.Address) ([32]byte, error) {
	return _IoIDRegistry.Contract.DocumentHash(&_IoIDRegistry.CallOpts, device)
}

// DocumentHash is a free data retrieval call binding the contract method 0x35984678.
//
// Solidity: function documentHash(address device) view returns(bytes32)
func (_IoIDRegistry *IoIDRegistryCallerSession) DocumentHash(device common.Address) ([32]byte, error) {
	return _IoIDRegistry.Contract.DocumentHash(&_IoIDRegistry.CallOpts, device)
}

// DocumentID is a free data retrieval call binding the contract method 0x4e93257e.
//
// Solidity: function documentID(address device) pure returns(string)
func (_IoIDRegistry *IoIDRegistryCaller) DocumentID(opts *bind.CallOpts, device common.Address) (string, error) {
	var out []interface{}
	err := _IoIDRegistry.contract.Call(opts, &out, "documentID", device)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DocumentID is a free data retrieval call binding the contract method 0x4e93257e.
//
// Solidity: function documentID(address device) pure returns(string)
func (_IoIDRegistry *IoIDRegistrySession) DocumentID(device common.Address) (string, error) {
	return _IoIDRegistry.Contract.DocumentID(&_IoIDRegistry.CallOpts, device)
}

// DocumentID is a free data retrieval call binding the contract method 0x4e93257e.
//
// Solidity: function documentID(address device) pure returns(string)
func (_IoIDRegistry *IoIDRegistryCallerSession) DocumentID(device common.Address) (string, error) {
	return _IoIDRegistry.Contract.DocumentID(&_IoIDRegistry.CallOpts, device)
}

// DocumentURI is a free data retrieval call binding the contract method 0x1d9dbae4.
//
// Solidity: function documentURI(address device) view returns(string)
func (_IoIDRegistry *IoIDRegistryCaller) DocumentURI(opts *bind.CallOpts, device common.Address) (string, error) {
	var out []interface{}
	err := _IoIDRegistry.contract.Call(opts, &out, "documentURI", device)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DocumentURI is a free data retrieval call binding the contract method 0x1d9dbae4.
//
// Solidity: function documentURI(address device) view returns(string)
func (_IoIDRegistry *IoIDRegistrySession) DocumentURI(device common.Address) (string, error) {
	return _IoIDRegistry.Contract.DocumentURI(&_IoIDRegistry.CallOpts, device)
}

// DocumentURI is a free data retrieval call binding the contract method 0x1d9dbae4.
//
// Solidity: function documentURI(address device) view returns(string)
func (_IoIDRegistry *IoIDRegistryCallerSession) DocumentURI(device common.Address) (string, error) {
	return _IoIDRegistry.Contract.DocumentURI(&_IoIDRegistry.CallOpts, device)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address device) view returns(bool)
func (_IoIDRegistry *IoIDRegistryCaller) Exists(opts *bind.CallOpts, device common.Address) (bool, error) {
	var out []interface{}
	err := _IoIDRegistry.contract.Call(opts, &out, "exists", device)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address device) view returns(bool)
func (_IoIDRegistry *IoIDRegistrySession) Exists(device common.Address) (bool, error) {
	return _IoIDRegistry.Contract.Exists(&_IoIDRegistry.CallOpts, device)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address device) view returns(bool)
func (_IoIDRegistry *IoIDRegistryCallerSession) Exists(device common.Address) (bool, error) {
	return _IoIDRegistry.Contract.Exists(&_IoIDRegistry.CallOpts, device)
}

// IoID is a free data retrieval call binding the contract method 0xc3b3135e.
//
// Solidity: function ioID() view returns(address)
func (_IoIDRegistry *IoIDRegistryCaller) IoID(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IoIDRegistry.contract.Call(opts, &out, "ioID")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IoID is a free data retrieval call binding the contract method 0xc3b3135e.
//
// Solidity: function ioID() view returns(address)
func (_IoIDRegistry *IoIDRegistrySession) IoID() (common.Address, error) {
	return _IoIDRegistry.Contract.IoID(&_IoIDRegistry.CallOpts)
}

// IoID is a free data retrieval call binding the contract method 0xc3b3135e.
//
// Solidity: function ioID() view returns(address)
func (_IoIDRegistry *IoIDRegistryCallerSession) IoID() (common.Address, error) {
	return _IoIDRegistry.Contract.IoID(&_IoIDRegistry.CallOpts)
}

// IoIDStore is a free data retrieval call binding the contract method 0xb6d32d69.
//
// Solidity: function ioIDStore() view returns(address)
func (_IoIDRegistry *IoIDRegistryCaller) IoIDStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IoIDRegistry.contract.Call(opts, &out, "ioIDStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IoIDStore is a free data retrieval call binding the contract method 0xb6d32d69.
//
// Solidity: function ioIDStore() view returns(address)
func (_IoIDRegistry *IoIDRegistrySession) IoIDStore() (common.Address, error) {
	return _IoIDRegistry.Contract.IoIDStore(&_IoIDRegistry.CallOpts)
}

// IoIDStore is a free data retrieval call binding the contract method 0xb6d32d69.
//
// Solidity: function ioIDStore() view returns(address)
func (_IoIDRegistry *IoIDRegistryCallerSession) IoIDStore() (common.Address, error) {
	return _IoIDRegistry.Contract.IoIDStore(&_IoIDRegistry.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address device) view returns(uint256)
func (_IoIDRegistry *IoIDRegistryCaller) Nonces(opts *bind.CallOpts, device common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IoIDRegistry.contract.Call(opts, &out, "nonces", device)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address device) view returns(uint256)
func (_IoIDRegistry *IoIDRegistrySession) Nonces(device common.Address) (*big.Int, error) {
	return _IoIDRegistry.Contract.Nonces(&_IoIDRegistry.CallOpts, device)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address device) view returns(uint256)
func (_IoIDRegistry *IoIDRegistryCallerSession) Nonces(device common.Address) (*big.Int, error) {
	return _IoIDRegistry.Contract.Nonces(&_IoIDRegistry.CallOpts, device)
}

// PermitHash is a free data retrieval call binding the contract method 0xb274703d.
//
// Solidity: function permitHash(address owner, address device) view returns(bytes32)
func (_IoIDRegistry *IoIDRegistryCaller) PermitHash(opts *bind.CallOpts, owner common.Address, device common.Address) ([32]byte, error) {
	var out []interface{}
	err := _IoIDRegistry.contract.Call(opts, &out, "permitHash", owner, device)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PermitHash is a free data retrieval call binding the contract method 0xb274703d.
//
// Solidity: function permitHash(address owner, address device) view returns(bytes32)
func (_IoIDRegistry *IoIDRegistrySession) PermitHash(owner common.Address, device common.Address) ([32]byte, error) {
	return _IoIDRegistry.Contract.PermitHash(&_IoIDRegistry.CallOpts, owner, device)
}

// PermitHash is a free data retrieval call binding the contract method 0xb274703d.
//
// Solidity: function permitHash(address owner, address device) view returns(bytes32)
func (_IoIDRegistry *IoIDRegistryCallerSession) PermitHash(owner common.Address, device common.Address) ([32]byte, error) {
	return _IoIDRegistry.Contract.PermitHash(&_IoIDRegistry.CallOpts, owner, device)
}

// RegisteredNFT is a free data retrieval call binding the contract method 0xb9d959fd.
//
// Solidity: function registeredNFT(address , uint256 ) view returns(bool)
func (_IoIDRegistry *IoIDRegistryCaller) RegisteredNFT(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _IoIDRegistry.contract.Call(opts, &out, "registeredNFT", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredNFT is a free data retrieval call binding the contract method 0xb9d959fd.
//
// Solidity: function registeredNFT(address , uint256 ) view returns(bool)
func (_IoIDRegistry *IoIDRegistrySession) RegisteredNFT(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _IoIDRegistry.Contract.RegisteredNFT(&_IoIDRegistry.CallOpts, arg0, arg1)
}

// RegisteredNFT is a free data retrieval call binding the contract method 0xb9d959fd.
//
// Solidity: function registeredNFT(address , uint256 ) view returns(bool)
func (_IoIDRegistry *IoIDRegistryCallerSession) RegisteredNFT(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _IoIDRegistry.Contract.RegisteredNFT(&_IoIDRegistry.CallOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ioIDStore, address _ioID) returns()
func (_IoIDRegistry *IoIDRegistryTransactor) Initialize(opts *bind.TransactOpts, _ioIDStore common.Address, _ioID common.Address) (*types.Transaction, error) {
	return _IoIDRegistry.contract.Transact(opts, "initialize", _ioIDStore, _ioID)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ioIDStore, address _ioID) returns()
func (_IoIDRegistry *IoIDRegistrySession) Initialize(_ioIDStore common.Address, _ioID common.Address) (*types.Transaction, error) {
	return _IoIDRegistry.Contract.Initialize(&_IoIDRegistry.TransactOpts, _ioIDStore, _ioID)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ioIDStore, address _ioID) returns()
func (_IoIDRegistry *IoIDRegistryTransactorSession) Initialize(_ioIDStore common.Address, _ioID common.Address) (*types.Transaction, error) {
	return _IoIDRegistry.Contract.Initialize(&_IoIDRegistry.TransactOpts, _ioIDStore, _ioID)
}

// Register is a paid mutator transaction binding the contract method 0x39a4a241.
//
// Solidity: function register(address deviceContract, uint256 tokenId, address device, bytes32 hash, string uri, uint8 v, bytes32 r, bytes32 s) returns()
func (_IoIDRegistry *IoIDRegistryTransactor) Register(opts *bind.TransactOpts, deviceContract common.Address, tokenId *big.Int, device common.Address, hash [32]byte, uri string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IoIDRegistry.contract.Transact(opts, "register", deviceContract, tokenId, device, hash, uri, v, r, s)
}

// Register is a paid mutator transaction binding the contract method 0x39a4a241.
//
// Solidity: function register(address deviceContract, uint256 tokenId, address device, bytes32 hash, string uri, uint8 v, bytes32 r, bytes32 s) returns()
func (_IoIDRegistry *IoIDRegistrySession) Register(deviceContract common.Address, tokenId *big.Int, device common.Address, hash [32]byte, uri string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IoIDRegistry.Contract.Register(&_IoIDRegistry.TransactOpts, deviceContract, tokenId, device, hash, uri, v, r, s)
}

// Register is a paid mutator transaction binding the contract method 0x39a4a241.
//
// Solidity: function register(address deviceContract, uint256 tokenId, address device, bytes32 hash, string uri, uint8 v, bytes32 r, bytes32 s) returns()
func (_IoIDRegistry *IoIDRegistryTransactorSession) Register(deviceContract common.Address, tokenId *big.Int, device common.Address, hash [32]byte, uri string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IoIDRegistry.Contract.Register(&_IoIDRegistry.TransactOpts, deviceContract, tokenId, device, hash, uri, v, r, s)
}

// Remove is a paid mutator transaction binding the contract method 0x937a7b2e.
//
// Solidity: function remove(address device, uint8 v, bytes32 r, bytes32 s) returns()
func (_IoIDRegistry *IoIDRegistryTransactor) Remove(opts *bind.TransactOpts, device common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IoIDRegistry.contract.Transact(opts, "remove", device, v, r, s)
}

// Remove is a paid mutator transaction binding the contract method 0x937a7b2e.
//
// Solidity: function remove(address device, uint8 v, bytes32 r, bytes32 s) returns()
func (_IoIDRegistry *IoIDRegistrySession) Remove(device common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IoIDRegistry.Contract.Remove(&_IoIDRegistry.TransactOpts, device, v, r, s)
}

// Remove is a paid mutator transaction binding the contract method 0x937a7b2e.
//
// Solidity: function remove(address device, uint8 v, bytes32 r, bytes32 s) returns()
func (_IoIDRegistry *IoIDRegistryTransactorSession) Remove(device common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IoIDRegistry.Contract.Remove(&_IoIDRegistry.TransactOpts, device, v, r, s)
}

// SetIoIDStore is a paid mutator transaction binding the contract method 0x4b38b385.
//
// Solidity: function setIoIDStore(address _ioIDStore) returns()
func (_IoIDRegistry *IoIDRegistryTransactor) SetIoIDStore(opts *bind.TransactOpts, _ioIDStore common.Address) (*types.Transaction, error) {
	return _IoIDRegistry.contract.Transact(opts, "setIoIDStore", _ioIDStore)
}

// SetIoIDStore is a paid mutator transaction binding the contract method 0x4b38b385.
//
// Solidity: function setIoIDStore(address _ioIDStore) returns()
func (_IoIDRegistry *IoIDRegistrySession) SetIoIDStore(_ioIDStore common.Address) (*types.Transaction, error) {
	return _IoIDRegistry.Contract.SetIoIDStore(&_IoIDRegistry.TransactOpts, _ioIDStore)
}

// SetIoIDStore is a paid mutator transaction binding the contract method 0x4b38b385.
//
// Solidity: function setIoIDStore(address _ioIDStore) returns()
func (_IoIDRegistry *IoIDRegistryTransactorSession) SetIoIDStore(_ioIDStore common.Address) (*types.Transaction, error) {
	return _IoIDRegistry.Contract.SetIoIDStore(&_IoIDRegistry.TransactOpts, _ioIDStore)
}

// Update is a paid mutator transaction binding the contract method 0xfd8fa0c1.
//
// Solidity: function update(address device, bytes32 hash, string uri, uint8 v, bytes32 r, bytes32 s) returns()
func (_IoIDRegistry *IoIDRegistryTransactor) Update(opts *bind.TransactOpts, device common.Address, hash [32]byte, uri string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IoIDRegistry.contract.Transact(opts, "update", device, hash, uri, v, r, s)
}

// Update is a paid mutator transaction binding the contract method 0xfd8fa0c1.
//
// Solidity: function update(address device, bytes32 hash, string uri, uint8 v, bytes32 r, bytes32 s) returns()
func (_IoIDRegistry *IoIDRegistrySession) Update(device common.Address, hash [32]byte, uri string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IoIDRegistry.Contract.Update(&_IoIDRegistry.TransactOpts, device, hash, uri, v, r, s)
}

// Update is a paid mutator transaction binding the contract method 0xfd8fa0c1.
//
// Solidity: function update(address device, bytes32 hash, string uri, uint8 v, bytes32 r, bytes32 s) returns()
func (_IoIDRegistry *IoIDRegistryTransactorSession) Update(device common.Address, hash [32]byte, uri string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IoIDRegistry.Contract.Update(&_IoIDRegistry.TransactOpts, device, hash, uri, v, r, s)
}

// IoIDRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IoIDRegistry contract.
type IoIDRegistryInitializedIterator struct {
	Event *IoIDRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *IoIDRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoIDRegistryInitialized)
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
		it.Event = new(IoIDRegistryInitialized)
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
func (it *IoIDRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoIDRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoIDRegistryInitialized represents a Initialized event raised by the IoIDRegistry contract.
type IoIDRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IoIDRegistry *IoIDRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*IoIDRegistryInitializedIterator, error) {

	logs, sub, err := _IoIDRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IoIDRegistryInitializedIterator{contract: _IoIDRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IoIDRegistry *IoIDRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IoIDRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _IoIDRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoIDRegistryInitialized)
				if err := _IoIDRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_IoIDRegistry *IoIDRegistryFilterer) ParseInitialized(log types.Log) (*IoIDRegistryInitialized, error) {
	event := new(IoIDRegistryInitialized)
	if err := _IoIDRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoIDRegistryNewDeviceIterator is returned from FilterNewDevice and is used to iterate over the raw logs and unpacked data for NewDevice events raised by the IoIDRegistry contract.
type IoIDRegistryNewDeviceIterator struct {
	Event *IoIDRegistryNewDevice // Event containing the contract specifics and raw log

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
func (it *IoIDRegistryNewDeviceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoIDRegistryNewDevice)
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
		it.Event = new(IoIDRegistryNewDevice)
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
func (it *IoIDRegistryNewDeviceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoIDRegistryNewDeviceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoIDRegistryNewDevice represents a NewDevice event raised by the IoIDRegistry contract.
type IoIDRegistryNewDevice struct {
	Device common.Address
	Owner  common.Address
	Hash   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewDevice is a free log retrieval operation binding the contract event 0x64693517e7e211d1bddfb7b25365841ef0903138547bda2758c5c12ebb47221e.
//
// Solidity: event NewDevice(address indexed device, address owner, bytes32 hash)
func (_IoIDRegistry *IoIDRegistryFilterer) FilterNewDevice(opts *bind.FilterOpts, device []common.Address) (*IoIDRegistryNewDeviceIterator, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _IoIDRegistry.contract.FilterLogs(opts, "NewDevice", deviceRule)
	if err != nil {
		return nil, err
	}
	return &IoIDRegistryNewDeviceIterator{contract: _IoIDRegistry.contract, event: "NewDevice", logs: logs, sub: sub}, nil
}

// WatchNewDevice is a free log subscription operation binding the contract event 0x64693517e7e211d1bddfb7b25365841ef0903138547bda2758c5c12ebb47221e.
//
// Solidity: event NewDevice(address indexed device, address owner, bytes32 hash)
func (_IoIDRegistry *IoIDRegistryFilterer) WatchNewDevice(opts *bind.WatchOpts, sink chan<- *IoIDRegistryNewDevice, device []common.Address) (event.Subscription, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _IoIDRegistry.contract.WatchLogs(opts, "NewDevice", deviceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoIDRegistryNewDevice)
				if err := _IoIDRegistry.contract.UnpackLog(event, "NewDevice", log); err != nil {
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

// ParseNewDevice is a log parse operation binding the contract event 0x64693517e7e211d1bddfb7b25365841ef0903138547bda2758c5c12ebb47221e.
//
// Solidity: event NewDevice(address indexed device, address owner, bytes32 hash)
func (_IoIDRegistry *IoIDRegistryFilterer) ParseNewDevice(log types.Log) (*IoIDRegistryNewDevice, error) {
	event := new(IoIDRegistryNewDevice)
	if err := _IoIDRegistry.contract.UnpackLog(event, "NewDevice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoIDRegistryRemoveDeviceIterator is returned from FilterRemoveDevice and is used to iterate over the raw logs and unpacked data for RemoveDevice events raised by the IoIDRegistry contract.
type IoIDRegistryRemoveDeviceIterator struct {
	Event *IoIDRegistryRemoveDevice // Event containing the contract specifics and raw log

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
func (it *IoIDRegistryRemoveDeviceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoIDRegistryRemoveDevice)
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
		it.Event = new(IoIDRegistryRemoveDevice)
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
func (it *IoIDRegistryRemoveDeviceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoIDRegistryRemoveDeviceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoIDRegistryRemoveDevice represents a RemoveDevice event raised by the IoIDRegistry contract.
type IoIDRegistryRemoveDevice struct {
	Device common.Address
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemoveDevice is a free log retrieval operation binding the contract event 0x7fb1f1a379ee4b2c5e787bdcba983dff2cb148ae93c6341beafaca37b8ce8abe.
//
// Solidity: event RemoveDevice(address indexed device, address owner)
func (_IoIDRegistry *IoIDRegistryFilterer) FilterRemoveDevice(opts *bind.FilterOpts, device []common.Address) (*IoIDRegistryRemoveDeviceIterator, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _IoIDRegistry.contract.FilterLogs(opts, "RemoveDevice", deviceRule)
	if err != nil {
		return nil, err
	}
	return &IoIDRegistryRemoveDeviceIterator{contract: _IoIDRegistry.contract, event: "RemoveDevice", logs: logs, sub: sub}, nil
}

// WatchRemoveDevice is a free log subscription operation binding the contract event 0x7fb1f1a379ee4b2c5e787bdcba983dff2cb148ae93c6341beafaca37b8ce8abe.
//
// Solidity: event RemoveDevice(address indexed device, address owner)
func (_IoIDRegistry *IoIDRegistryFilterer) WatchRemoveDevice(opts *bind.WatchOpts, sink chan<- *IoIDRegistryRemoveDevice, device []common.Address) (event.Subscription, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _IoIDRegistry.contract.WatchLogs(opts, "RemoveDevice", deviceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoIDRegistryRemoveDevice)
				if err := _IoIDRegistry.contract.UnpackLog(event, "RemoveDevice", log); err != nil {
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

// ParseRemoveDevice is a log parse operation binding the contract event 0x7fb1f1a379ee4b2c5e787bdcba983dff2cb148ae93c6341beafaca37b8ce8abe.
//
// Solidity: event RemoveDevice(address indexed device, address owner)
func (_IoIDRegistry *IoIDRegistryFilterer) ParseRemoveDevice(log types.Log) (*IoIDRegistryRemoveDevice, error) {
	event := new(IoIDRegistryRemoveDevice)
	if err := _IoIDRegistry.contract.UnpackLog(event, "RemoveDevice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoIDRegistrySetIoIdStoreIterator is returned from FilterSetIoIdStore and is used to iterate over the raw logs and unpacked data for SetIoIdStore events raised by the IoIDRegistry contract.
type IoIDRegistrySetIoIdStoreIterator struct {
	Event *IoIDRegistrySetIoIdStore // Event containing the contract specifics and raw log

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
func (it *IoIDRegistrySetIoIdStoreIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoIDRegistrySetIoIdStore)
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
		it.Event = new(IoIDRegistrySetIoIdStore)
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
func (it *IoIDRegistrySetIoIdStoreIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoIDRegistrySetIoIdStoreIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoIDRegistrySetIoIdStore represents a SetIoIdStore event raised by the IoIDRegistry contract.
type IoIDRegistrySetIoIdStore struct {
	Store common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetIoIdStore is a free log retrieval operation binding the contract event 0xc64dd2865233a317a221a0952325ee30744b237e5b6e4367fff3aeea454dfee9.
//
// Solidity: event SetIoIdStore(address indexed store)
func (_IoIDRegistry *IoIDRegistryFilterer) FilterSetIoIdStore(opts *bind.FilterOpts, store []common.Address) (*IoIDRegistrySetIoIdStoreIterator, error) {

	var storeRule []interface{}
	for _, storeItem := range store {
		storeRule = append(storeRule, storeItem)
	}

	logs, sub, err := _IoIDRegistry.contract.FilterLogs(opts, "SetIoIdStore", storeRule)
	if err != nil {
		return nil, err
	}
	return &IoIDRegistrySetIoIdStoreIterator{contract: _IoIDRegistry.contract, event: "SetIoIdStore", logs: logs, sub: sub}, nil
}

// WatchSetIoIdStore is a free log subscription operation binding the contract event 0xc64dd2865233a317a221a0952325ee30744b237e5b6e4367fff3aeea454dfee9.
//
// Solidity: event SetIoIdStore(address indexed store)
func (_IoIDRegistry *IoIDRegistryFilterer) WatchSetIoIdStore(opts *bind.WatchOpts, sink chan<- *IoIDRegistrySetIoIdStore, store []common.Address) (event.Subscription, error) {

	var storeRule []interface{}
	for _, storeItem := range store {
		storeRule = append(storeRule, storeItem)
	}

	logs, sub, err := _IoIDRegistry.contract.WatchLogs(opts, "SetIoIdStore", storeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoIDRegistrySetIoIdStore)
				if err := _IoIDRegistry.contract.UnpackLog(event, "SetIoIdStore", log); err != nil {
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

// ParseSetIoIdStore is a log parse operation binding the contract event 0xc64dd2865233a317a221a0952325ee30744b237e5b6e4367fff3aeea454dfee9.
//
// Solidity: event SetIoIdStore(address indexed store)
func (_IoIDRegistry *IoIDRegistryFilterer) ParseSetIoIdStore(log types.Log) (*IoIDRegistrySetIoIdStore, error) {
	event := new(IoIDRegistrySetIoIdStore)
	if err := _IoIDRegistry.contract.UnpackLog(event, "SetIoIdStore", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoIDRegistryUpdateDeviceIterator is returned from FilterUpdateDevice and is used to iterate over the raw logs and unpacked data for UpdateDevice events raised by the IoIDRegistry contract.
type IoIDRegistryUpdateDeviceIterator struct {
	Event *IoIDRegistryUpdateDevice // Event containing the contract specifics and raw log

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
func (it *IoIDRegistryUpdateDeviceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoIDRegistryUpdateDevice)
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
		it.Event = new(IoIDRegistryUpdateDevice)
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
func (it *IoIDRegistryUpdateDeviceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoIDRegistryUpdateDeviceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoIDRegistryUpdateDevice represents a UpdateDevice event raised by the IoIDRegistry contract.
type IoIDRegistryUpdateDevice struct {
	Device common.Address
	Owner  common.Address
	Hash   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateDevice is a free log retrieval operation binding the contract event 0x50c559b3a189fe72395f086f1ed1afe46c593ec513dde8c31b8f9b6f14b8560c.
//
// Solidity: event UpdateDevice(address indexed device, address owner, bytes32 hash)
func (_IoIDRegistry *IoIDRegistryFilterer) FilterUpdateDevice(opts *bind.FilterOpts, device []common.Address) (*IoIDRegistryUpdateDeviceIterator, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _IoIDRegistry.contract.FilterLogs(opts, "UpdateDevice", deviceRule)
	if err != nil {
		return nil, err
	}
	return &IoIDRegistryUpdateDeviceIterator{contract: _IoIDRegistry.contract, event: "UpdateDevice", logs: logs, sub: sub}, nil
}

// WatchUpdateDevice is a free log subscription operation binding the contract event 0x50c559b3a189fe72395f086f1ed1afe46c593ec513dde8c31b8f9b6f14b8560c.
//
// Solidity: event UpdateDevice(address indexed device, address owner, bytes32 hash)
func (_IoIDRegistry *IoIDRegistryFilterer) WatchUpdateDevice(opts *bind.WatchOpts, sink chan<- *IoIDRegistryUpdateDevice, device []common.Address) (event.Subscription, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _IoIDRegistry.contract.WatchLogs(opts, "UpdateDevice", deviceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoIDRegistryUpdateDevice)
				if err := _IoIDRegistry.contract.UnpackLog(event, "UpdateDevice", log); err != nil {
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

// ParseUpdateDevice is a log parse operation binding the contract event 0x50c559b3a189fe72395f086f1ed1afe46c593ec513dde8c31b8f9b6f14b8560c.
//
// Solidity: event UpdateDevice(address indexed device, address owner, bytes32 hash)
func (_IoIDRegistry *IoIDRegistryFilterer) ParseUpdateDevice(log types.Log) (*IoIDRegistryUpdateDevice, error) {
	event := new(IoIDRegistryUpdateDevice)
	if err := _IoIDRegistry.contract.UnpackLog(event, "UpdateDevice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
