// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ioidregistry

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

// IoidregistryMetaData contains all meta data concerning the Ioidregistry contract.
var IoidregistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"NewDevice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"RemoveDevice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"store\",\"type\":\"address\"}],\"name\":\"SetIoIdStore\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"UpdateDevice\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"METHOD\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"deviceTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"documentHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"documentID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"documentURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ioIDStore\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ioID\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ioID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ioIDStore\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"}],\"name\":\"permitHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deviceContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deviceContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registeredNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ioIDStore\",\"type\":\"address\"}],\"name\":\"setIoIDStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"device\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IoidregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IoidregistryMetaData.ABI instead.
var IoidregistryABI = IoidregistryMetaData.ABI

// Ioidregistry is an auto generated Go binding around an Ethereum contract.
type Ioidregistry struct {
	IoidregistryCaller     // Read-only binding to the contract
	IoidregistryTransactor // Write-only binding to the contract
	IoidregistryFilterer   // Log filterer for contract events
}

// IoidregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IoidregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoidregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IoidregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoidregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IoidregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoidregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IoidregistrySession struct {
	Contract     *Ioidregistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IoidregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IoidregistryCallerSession struct {
	Contract *IoidregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IoidregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IoidregistryTransactorSession struct {
	Contract     *IoidregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IoidregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IoidregistryRaw struct {
	Contract *Ioidregistry // Generic contract binding to access the raw methods on
}

// IoidregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IoidregistryCallerRaw struct {
	Contract *IoidregistryCaller // Generic read-only contract binding to access the raw methods on
}

// IoidregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IoidregistryTransactorRaw struct {
	Contract *IoidregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIoidregistry creates a new instance of Ioidregistry, bound to a specific deployed contract.
func NewIoidregistry(address common.Address, backend bind.ContractBackend) (*Ioidregistry, error) {
	contract, err := bindIoidregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ioidregistry{IoidregistryCaller: IoidregistryCaller{contract: contract}, IoidregistryTransactor: IoidregistryTransactor{contract: contract}, IoidregistryFilterer: IoidregistryFilterer{contract: contract}}, nil
}

// NewIoidregistryCaller creates a new read-only instance of Ioidregistry, bound to a specific deployed contract.
func NewIoidregistryCaller(address common.Address, caller bind.ContractCaller) (*IoidregistryCaller, error) {
	contract, err := bindIoidregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IoidregistryCaller{contract: contract}, nil
}

// NewIoidregistryTransactor creates a new write-only instance of Ioidregistry, bound to a specific deployed contract.
func NewIoidregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IoidregistryTransactor, error) {
	contract, err := bindIoidregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IoidregistryTransactor{contract: contract}, nil
}

// NewIoidregistryFilterer creates a new log filterer instance of Ioidregistry, bound to a specific deployed contract.
func NewIoidregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IoidregistryFilterer, error) {
	contract, err := bindIoidregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IoidregistryFilterer{contract: contract}, nil
}

// bindIoidregistry binds a generic wrapper to an already deployed contract.
func bindIoidregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IoidregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ioidregistry *IoidregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ioidregistry.Contract.IoidregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ioidregistry *IoidregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ioidregistry.Contract.IoidregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ioidregistry *IoidregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ioidregistry.Contract.IoidregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ioidregistry *IoidregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ioidregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ioidregistry *IoidregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ioidregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ioidregistry *IoidregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ioidregistry.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Ioidregistry *IoidregistryCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ioidregistry.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Ioidregistry *IoidregistrySession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Ioidregistry.Contract.DOMAINSEPARATOR(&_Ioidregistry.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Ioidregistry *IoidregistryCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Ioidregistry.Contract.DOMAINSEPARATOR(&_Ioidregistry.CallOpts)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc49f91d3.
//
// Solidity: function EIP712DOMAIN_TYPEHASH() view returns(bytes32)
func (_Ioidregistry *IoidregistryCaller) EIP712DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ioidregistry.contract.Call(opts, &out, "EIP712DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc49f91d3.
//
// Solidity: function EIP712DOMAIN_TYPEHASH() view returns(bytes32)
func (_Ioidregistry *IoidregistrySession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _Ioidregistry.Contract.EIP712DOMAINTYPEHASH(&_Ioidregistry.CallOpts)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc49f91d3.
//
// Solidity: function EIP712DOMAIN_TYPEHASH() view returns(bytes32)
func (_Ioidregistry *IoidregistryCallerSession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _Ioidregistry.Contract.EIP712DOMAINTYPEHASH(&_Ioidregistry.CallOpts)
}

// METHOD is a free data retrieval call binding the contract method 0x67efe82e.
//
// Solidity: function METHOD() view returns(string)
func (_Ioidregistry *IoidregistryCaller) METHOD(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ioidregistry.contract.Call(opts, &out, "METHOD")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// METHOD is a free data retrieval call binding the contract method 0x67efe82e.
//
// Solidity: function METHOD() view returns(string)
func (_Ioidregistry *IoidregistrySession) METHOD() (string, error) {
	return _Ioidregistry.Contract.METHOD(&_Ioidregistry.CallOpts)
}

// METHOD is a free data retrieval call binding the contract method 0x67efe82e.
//
// Solidity: function METHOD() view returns(string)
func (_Ioidregistry *IoidregistryCallerSession) METHOD() (string, error) {
	return _Ioidregistry.Contract.METHOD(&_Ioidregistry.CallOpts)
}

// DeviceTokenId is a free data retrieval call binding the contract method 0x4965c7e1.
//
// Solidity: function deviceTokenId(address device) view returns(uint256)
func (_Ioidregistry *IoidregistryCaller) DeviceTokenId(opts *bind.CallOpts, device common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ioidregistry.contract.Call(opts, &out, "deviceTokenId", device)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeviceTokenId is a free data retrieval call binding the contract method 0x4965c7e1.
//
// Solidity: function deviceTokenId(address device) view returns(uint256)
func (_Ioidregistry *IoidregistrySession) DeviceTokenId(device common.Address) (*big.Int, error) {
	return _Ioidregistry.Contract.DeviceTokenId(&_Ioidregistry.CallOpts, device)
}

// DeviceTokenId is a free data retrieval call binding the contract method 0x4965c7e1.
//
// Solidity: function deviceTokenId(address device) view returns(uint256)
func (_Ioidregistry *IoidregistryCallerSession) DeviceTokenId(device common.Address) (*big.Int, error) {
	return _Ioidregistry.Contract.DeviceTokenId(&_Ioidregistry.CallOpts, device)
}

// DocumentHash is a free data retrieval call binding the contract method 0x35984678.
//
// Solidity: function documentHash(address device) view returns(bytes32)
func (_Ioidregistry *IoidregistryCaller) DocumentHash(opts *bind.CallOpts, device common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Ioidregistry.contract.Call(opts, &out, "documentHash", device)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DocumentHash is a free data retrieval call binding the contract method 0x35984678.
//
// Solidity: function documentHash(address device) view returns(bytes32)
func (_Ioidregistry *IoidregistrySession) DocumentHash(device common.Address) ([32]byte, error) {
	return _Ioidregistry.Contract.DocumentHash(&_Ioidregistry.CallOpts, device)
}

// DocumentHash is a free data retrieval call binding the contract method 0x35984678.
//
// Solidity: function documentHash(address device) view returns(bytes32)
func (_Ioidregistry *IoidregistryCallerSession) DocumentHash(device common.Address) ([32]byte, error) {
	return _Ioidregistry.Contract.DocumentHash(&_Ioidregistry.CallOpts, device)
}

// DocumentID is a free data retrieval call binding the contract method 0x4e93257e.
//
// Solidity: function documentID(address device) pure returns(string)
func (_Ioidregistry *IoidregistryCaller) DocumentID(opts *bind.CallOpts, device common.Address) (string, error) {
	var out []interface{}
	err := _Ioidregistry.contract.Call(opts, &out, "documentID", device)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DocumentID is a free data retrieval call binding the contract method 0x4e93257e.
//
// Solidity: function documentID(address device) pure returns(string)
func (_Ioidregistry *IoidregistrySession) DocumentID(device common.Address) (string, error) {
	return _Ioidregistry.Contract.DocumentID(&_Ioidregistry.CallOpts, device)
}

// DocumentID is a free data retrieval call binding the contract method 0x4e93257e.
//
// Solidity: function documentID(address device) pure returns(string)
func (_Ioidregistry *IoidregistryCallerSession) DocumentID(device common.Address) (string, error) {
	return _Ioidregistry.Contract.DocumentID(&_Ioidregistry.CallOpts, device)
}

// DocumentURI is a free data retrieval call binding the contract method 0x1d9dbae4.
//
// Solidity: function documentURI(address device) view returns(string)
func (_Ioidregistry *IoidregistryCaller) DocumentURI(opts *bind.CallOpts, device common.Address) (string, error) {
	var out []interface{}
	err := _Ioidregistry.contract.Call(opts, &out, "documentURI", device)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DocumentURI is a free data retrieval call binding the contract method 0x1d9dbae4.
//
// Solidity: function documentURI(address device) view returns(string)
func (_Ioidregistry *IoidregistrySession) DocumentURI(device common.Address) (string, error) {
	return _Ioidregistry.Contract.DocumentURI(&_Ioidregistry.CallOpts, device)
}

// DocumentURI is a free data retrieval call binding the contract method 0x1d9dbae4.
//
// Solidity: function documentURI(address device) view returns(string)
func (_Ioidregistry *IoidregistryCallerSession) DocumentURI(device common.Address) (string, error) {
	return _Ioidregistry.Contract.DocumentURI(&_Ioidregistry.CallOpts, device)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address device) view returns(bool)
func (_Ioidregistry *IoidregistryCaller) Exists(opts *bind.CallOpts, device common.Address) (bool, error) {
	var out []interface{}
	err := _Ioidregistry.contract.Call(opts, &out, "exists", device)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address device) view returns(bool)
func (_Ioidregistry *IoidregistrySession) Exists(device common.Address) (bool, error) {
	return _Ioidregistry.Contract.Exists(&_Ioidregistry.CallOpts, device)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address device) view returns(bool)
func (_Ioidregistry *IoidregistryCallerSession) Exists(device common.Address) (bool, error) {
	return _Ioidregistry.Contract.Exists(&_Ioidregistry.CallOpts, device)
}

// IoID is a free data retrieval call binding the contract method 0xc3b3135e.
//
// Solidity: function ioID() view returns(address)
func (_Ioidregistry *IoidregistryCaller) IoID(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ioidregistry.contract.Call(opts, &out, "ioID")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IoID is a free data retrieval call binding the contract method 0xc3b3135e.
//
// Solidity: function ioID() view returns(address)
func (_Ioidregistry *IoidregistrySession) IoID() (common.Address, error) {
	return _Ioidregistry.Contract.IoID(&_Ioidregistry.CallOpts)
}

// IoID is a free data retrieval call binding the contract method 0xc3b3135e.
//
// Solidity: function ioID() view returns(address)
func (_Ioidregistry *IoidregistryCallerSession) IoID() (common.Address, error) {
	return _Ioidregistry.Contract.IoID(&_Ioidregistry.CallOpts)
}

// IoIDStore is a free data retrieval call binding the contract method 0xb6d32d69.
//
// Solidity: function ioIDStore() view returns(address)
func (_Ioidregistry *IoidregistryCaller) IoIDStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ioidregistry.contract.Call(opts, &out, "ioIDStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IoIDStore is a free data retrieval call binding the contract method 0xb6d32d69.
//
// Solidity: function ioIDStore() view returns(address)
func (_Ioidregistry *IoidregistrySession) IoIDStore() (common.Address, error) {
	return _Ioidregistry.Contract.IoIDStore(&_Ioidregistry.CallOpts)
}

// IoIDStore is a free data retrieval call binding the contract method 0xb6d32d69.
//
// Solidity: function ioIDStore() view returns(address)
func (_Ioidregistry *IoidregistryCallerSession) IoIDStore() (common.Address, error) {
	return _Ioidregistry.Contract.IoIDStore(&_Ioidregistry.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address device) view returns(uint256)
func (_Ioidregistry *IoidregistryCaller) Nonces(opts *bind.CallOpts, device common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ioidregistry.contract.Call(opts, &out, "nonces", device)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address device) view returns(uint256)
func (_Ioidregistry *IoidregistrySession) Nonces(device common.Address) (*big.Int, error) {
	return _Ioidregistry.Contract.Nonces(&_Ioidregistry.CallOpts, device)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address device) view returns(uint256)
func (_Ioidregistry *IoidregistryCallerSession) Nonces(device common.Address) (*big.Int, error) {
	return _Ioidregistry.Contract.Nonces(&_Ioidregistry.CallOpts, device)
}

// PermitHash is a free data retrieval call binding the contract method 0xb274703d.
//
// Solidity: function permitHash(address owner, address device) view returns(bytes32)
func (_Ioidregistry *IoidregistryCaller) PermitHash(opts *bind.CallOpts, owner common.Address, device common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Ioidregistry.contract.Call(opts, &out, "permitHash", owner, device)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PermitHash is a free data retrieval call binding the contract method 0xb274703d.
//
// Solidity: function permitHash(address owner, address device) view returns(bytes32)
func (_Ioidregistry *IoidregistrySession) PermitHash(owner common.Address, device common.Address) ([32]byte, error) {
	return _Ioidregistry.Contract.PermitHash(&_Ioidregistry.CallOpts, owner, device)
}

// PermitHash is a free data retrieval call binding the contract method 0xb274703d.
//
// Solidity: function permitHash(address owner, address device) view returns(bytes32)
func (_Ioidregistry *IoidregistryCallerSession) PermitHash(owner common.Address, device common.Address) ([32]byte, error) {
	return _Ioidregistry.Contract.PermitHash(&_Ioidregistry.CallOpts, owner, device)
}

// RegisteredNFT is a free data retrieval call binding the contract method 0xb9d959fd.
//
// Solidity: function registeredNFT(address , uint256 ) view returns(bool)
func (_Ioidregistry *IoidregistryCaller) RegisteredNFT(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Ioidregistry.contract.Call(opts, &out, "registeredNFT", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredNFT is a free data retrieval call binding the contract method 0xb9d959fd.
//
// Solidity: function registeredNFT(address , uint256 ) view returns(bool)
func (_Ioidregistry *IoidregistrySession) RegisteredNFT(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Ioidregistry.Contract.RegisteredNFT(&_Ioidregistry.CallOpts, arg0, arg1)
}

// RegisteredNFT is a free data retrieval call binding the contract method 0xb9d959fd.
//
// Solidity: function registeredNFT(address , uint256 ) view returns(bool)
func (_Ioidregistry *IoidregistryCallerSession) RegisteredNFT(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Ioidregistry.Contract.RegisteredNFT(&_Ioidregistry.CallOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ioIDStore, address _ioID) returns()
func (_Ioidregistry *IoidregistryTransactor) Initialize(opts *bind.TransactOpts, _ioIDStore common.Address, _ioID common.Address) (*types.Transaction, error) {
	return _Ioidregistry.contract.Transact(opts, "initialize", _ioIDStore, _ioID)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ioIDStore, address _ioID) returns()
func (_Ioidregistry *IoidregistrySession) Initialize(_ioIDStore common.Address, _ioID common.Address) (*types.Transaction, error) {
	return _Ioidregistry.Contract.Initialize(&_Ioidregistry.TransactOpts, _ioIDStore, _ioID)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ioIDStore, address _ioID) returns()
func (_Ioidregistry *IoidregistryTransactorSession) Initialize(_ioIDStore common.Address, _ioID common.Address) (*types.Transaction, error) {
	return _Ioidregistry.Contract.Initialize(&_Ioidregistry.TransactOpts, _ioIDStore, _ioID)
}

// Register is a paid mutator transaction binding the contract method 0x39a4a241.
//
// Solidity: function register(address deviceContract, uint256 tokenId, address device, bytes32 hash, string uri, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Ioidregistry *IoidregistryTransactor) Register(opts *bind.TransactOpts, deviceContract common.Address, tokenId *big.Int, device common.Address, hash [32]byte, uri string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Ioidregistry.contract.Transact(opts, "register", deviceContract, tokenId, device, hash, uri, v, r, s)
}

// Register is a paid mutator transaction binding the contract method 0x39a4a241.
//
// Solidity: function register(address deviceContract, uint256 tokenId, address device, bytes32 hash, string uri, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Ioidregistry *IoidregistrySession) Register(deviceContract common.Address, tokenId *big.Int, device common.Address, hash [32]byte, uri string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Ioidregistry.Contract.Register(&_Ioidregistry.TransactOpts, deviceContract, tokenId, device, hash, uri, v, r, s)
}

// Register is a paid mutator transaction binding the contract method 0x39a4a241.
//
// Solidity: function register(address deviceContract, uint256 tokenId, address device, bytes32 hash, string uri, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Ioidregistry *IoidregistryTransactorSession) Register(deviceContract common.Address, tokenId *big.Int, device common.Address, hash [32]byte, uri string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Ioidregistry.Contract.Register(&_Ioidregistry.TransactOpts, deviceContract, tokenId, device, hash, uri, v, r, s)
}

// Register0 is a paid mutator transaction binding the contract method 0xb20187f1.
//
// Solidity: function register(address deviceContract, uint256 tokenId, address user, address device, bytes32 hash, string uri, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Ioidregistry *IoidregistryTransactor) Register0(opts *bind.TransactOpts, deviceContract common.Address, tokenId *big.Int, user common.Address, device common.Address, hash [32]byte, uri string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Ioidregistry.contract.Transact(opts, "register0", deviceContract, tokenId, user, device, hash, uri, v, r, s)
}

// Register0 is a paid mutator transaction binding the contract method 0xb20187f1.
//
// Solidity: function register(address deviceContract, uint256 tokenId, address user, address device, bytes32 hash, string uri, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Ioidregistry *IoidregistrySession) Register0(deviceContract common.Address, tokenId *big.Int, user common.Address, device common.Address, hash [32]byte, uri string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Ioidregistry.Contract.Register0(&_Ioidregistry.TransactOpts, deviceContract, tokenId, user, device, hash, uri, v, r, s)
}

// Register0 is a paid mutator transaction binding the contract method 0xb20187f1.
//
// Solidity: function register(address deviceContract, uint256 tokenId, address user, address device, bytes32 hash, string uri, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Ioidregistry *IoidregistryTransactorSession) Register0(deviceContract common.Address, tokenId *big.Int, user common.Address, device common.Address, hash [32]byte, uri string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Ioidregistry.Contract.Register0(&_Ioidregistry.TransactOpts, deviceContract, tokenId, user, device, hash, uri, v, r, s)
}

// Remove is a paid mutator transaction binding the contract method 0x937a7b2e.
//
// Solidity: function remove(address device, uint8 v, bytes32 r, bytes32 s) returns()
func (_Ioidregistry *IoidregistryTransactor) Remove(opts *bind.TransactOpts, device common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Ioidregistry.contract.Transact(opts, "remove", device, v, r, s)
}

// Remove is a paid mutator transaction binding the contract method 0x937a7b2e.
//
// Solidity: function remove(address device, uint8 v, bytes32 r, bytes32 s) returns()
func (_Ioidregistry *IoidregistrySession) Remove(device common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Ioidregistry.Contract.Remove(&_Ioidregistry.TransactOpts, device, v, r, s)
}

// Remove is a paid mutator transaction binding the contract method 0x937a7b2e.
//
// Solidity: function remove(address device, uint8 v, bytes32 r, bytes32 s) returns()
func (_Ioidregistry *IoidregistryTransactorSession) Remove(device common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Ioidregistry.Contract.Remove(&_Ioidregistry.TransactOpts, device, v, r, s)
}

// SetIoIDStore is a paid mutator transaction binding the contract method 0x4b38b385.
//
// Solidity: function setIoIDStore(address _ioIDStore) returns()
func (_Ioidregistry *IoidregistryTransactor) SetIoIDStore(opts *bind.TransactOpts, _ioIDStore common.Address) (*types.Transaction, error) {
	return _Ioidregistry.contract.Transact(opts, "setIoIDStore", _ioIDStore)
}

// SetIoIDStore is a paid mutator transaction binding the contract method 0x4b38b385.
//
// Solidity: function setIoIDStore(address _ioIDStore) returns()
func (_Ioidregistry *IoidregistrySession) SetIoIDStore(_ioIDStore common.Address) (*types.Transaction, error) {
	return _Ioidregistry.Contract.SetIoIDStore(&_Ioidregistry.TransactOpts, _ioIDStore)
}

// SetIoIDStore is a paid mutator transaction binding the contract method 0x4b38b385.
//
// Solidity: function setIoIDStore(address _ioIDStore) returns()
func (_Ioidregistry *IoidregistryTransactorSession) SetIoIDStore(_ioIDStore common.Address) (*types.Transaction, error) {
	return _Ioidregistry.Contract.SetIoIDStore(&_Ioidregistry.TransactOpts, _ioIDStore)
}

// Update is a paid mutator transaction binding the contract method 0xfd8fa0c1.
//
// Solidity: function update(address device, bytes32 hash, string uri, uint8 v, bytes32 r, bytes32 s) returns()
func (_Ioidregistry *IoidregistryTransactor) Update(opts *bind.TransactOpts, device common.Address, hash [32]byte, uri string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Ioidregistry.contract.Transact(opts, "update", device, hash, uri, v, r, s)
}

// Update is a paid mutator transaction binding the contract method 0xfd8fa0c1.
//
// Solidity: function update(address device, bytes32 hash, string uri, uint8 v, bytes32 r, bytes32 s) returns()
func (_Ioidregistry *IoidregistrySession) Update(device common.Address, hash [32]byte, uri string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Ioidregistry.Contract.Update(&_Ioidregistry.TransactOpts, device, hash, uri, v, r, s)
}

// Update is a paid mutator transaction binding the contract method 0xfd8fa0c1.
//
// Solidity: function update(address device, bytes32 hash, string uri, uint8 v, bytes32 r, bytes32 s) returns()
func (_Ioidregistry *IoidregistryTransactorSession) Update(device common.Address, hash [32]byte, uri string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Ioidregistry.Contract.Update(&_Ioidregistry.TransactOpts, device, hash, uri, v, r, s)
}

// IoidregistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Ioidregistry contract.
type IoidregistryInitializedIterator struct {
	Event *IoidregistryInitialized // Event containing the contract specifics and raw log

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
func (it *IoidregistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoidregistryInitialized)
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
		it.Event = new(IoidregistryInitialized)
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
func (it *IoidregistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoidregistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoidregistryInitialized represents a Initialized event raised by the Ioidregistry contract.
type IoidregistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ioidregistry *IoidregistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*IoidregistryInitializedIterator, error) {

	logs, sub, err := _Ioidregistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IoidregistryInitializedIterator{contract: _Ioidregistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ioidregistry *IoidregistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IoidregistryInitialized) (event.Subscription, error) {

	logs, sub, err := _Ioidregistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoidregistryInitialized)
				if err := _Ioidregistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Ioidregistry *IoidregistryFilterer) ParseInitialized(log types.Log) (*IoidregistryInitialized, error) {
	event := new(IoidregistryInitialized)
	if err := _Ioidregistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoidregistryNewDeviceIterator is returned from FilterNewDevice and is used to iterate over the raw logs and unpacked data for NewDevice events raised by the Ioidregistry contract.
type IoidregistryNewDeviceIterator struct {
	Event *IoidregistryNewDevice // Event containing the contract specifics and raw log

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
func (it *IoidregistryNewDeviceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoidregistryNewDevice)
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
		it.Event = new(IoidregistryNewDevice)
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
func (it *IoidregistryNewDeviceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoidregistryNewDeviceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoidregistryNewDevice represents a NewDevice event raised by the Ioidregistry contract.
type IoidregistryNewDevice struct {
	Device common.Address
	Owner  common.Address
	Hash   [32]byte
	Uri    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewDevice is a free log retrieval operation binding the contract event 0x9b3397f29f76b91430e368033ed86df2fc794753ab8239898fb655ab0213f888.
//
// Solidity: event NewDevice(address indexed device, address owner, bytes32 hash, string uri)
func (_Ioidregistry *IoidregistryFilterer) FilterNewDevice(opts *bind.FilterOpts, device []common.Address) (*IoidregistryNewDeviceIterator, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _Ioidregistry.contract.FilterLogs(opts, "NewDevice", deviceRule)
	if err != nil {
		return nil, err
	}
	return &IoidregistryNewDeviceIterator{contract: _Ioidregistry.contract, event: "NewDevice", logs: logs, sub: sub}, nil
}

// WatchNewDevice is a free log subscription operation binding the contract event 0x9b3397f29f76b91430e368033ed86df2fc794753ab8239898fb655ab0213f888.
//
// Solidity: event NewDevice(address indexed device, address owner, bytes32 hash, string uri)
func (_Ioidregistry *IoidregistryFilterer) WatchNewDevice(opts *bind.WatchOpts, sink chan<- *IoidregistryNewDevice, device []common.Address) (event.Subscription, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _Ioidregistry.contract.WatchLogs(opts, "NewDevice", deviceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoidregistryNewDevice)
				if err := _Ioidregistry.contract.UnpackLog(event, "NewDevice", log); err != nil {
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

// ParseNewDevice is a log parse operation binding the contract event 0x9b3397f29f76b91430e368033ed86df2fc794753ab8239898fb655ab0213f888.
//
// Solidity: event NewDevice(address indexed device, address owner, bytes32 hash, string uri)
func (_Ioidregistry *IoidregistryFilterer) ParseNewDevice(log types.Log) (*IoidregistryNewDevice, error) {
	event := new(IoidregistryNewDevice)
	if err := _Ioidregistry.contract.UnpackLog(event, "NewDevice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoidregistryRemoveDeviceIterator is returned from FilterRemoveDevice and is used to iterate over the raw logs and unpacked data for RemoveDevice events raised by the Ioidregistry contract.
type IoidregistryRemoveDeviceIterator struct {
	Event *IoidregistryRemoveDevice // Event containing the contract specifics and raw log

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
func (it *IoidregistryRemoveDeviceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoidregistryRemoveDevice)
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
		it.Event = new(IoidregistryRemoveDevice)
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
func (it *IoidregistryRemoveDeviceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoidregistryRemoveDeviceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoidregistryRemoveDevice represents a RemoveDevice event raised by the Ioidregistry contract.
type IoidregistryRemoveDevice struct {
	Device common.Address
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemoveDevice is a free log retrieval operation binding the contract event 0x7fb1f1a379ee4b2c5e787bdcba983dff2cb148ae93c6341beafaca37b8ce8abe.
//
// Solidity: event RemoveDevice(address indexed device, address owner)
func (_Ioidregistry *IoidregistryFilterer) FilterRemoveDevice(opts *bind.FilterOpts, device []common.Address) (*IoidregistryRemoveDeviceIterator, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _Ioidregistry.contract.FilterLogs(opts, "RemoveDevice", deviceRule)
	if err != nil {
		return nil, err
	}
	return &IoidregistryRemoveDeviceIterator{contract: _Ioidregistry.contract, event: "RemoveDevice", logs: logs, sub: sub}, nil
}

// WatchRemoveDevice is a free log subscription operation binding the contract event 0x7fb1f1a379ee4b2c5e787bdcba983dff2cb148ae93c6341beafaca37b8ce8abe.
//
// Solidity: event RemoveDevice(address indexed device, address owner)
func (_Ioidregistry *IoidregistryFilterer) WatchRemoveDevice(opts *bind.WatchOpts, sink chan<- *IoidregistryRemoveDevice, device []common.Address) (event.Subscription, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _Ioidregistry.contract.WatchLogs(opts, "RemoveDevice", deviceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoidregistryRemoveDevice)
				if err := _Ioidregistry.contract.UnpackLog(event, "RemoveDevice", log); err != nil {
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
func (_Ioidregistry *IoidregistryFilterer) ParseRemoveDevice(log types.Log) (*IoidregistryRemoveDevice, error) {
	event := new(IoidregistryRemoveDevice)
	if err := _Ioidregistry.contract.UnpackLog(event, "RemoveDevice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoidregistrySetIoIdStoreIterator is returned from FilterSetIoIdStore and is used to iterate over the raw logs and unpacked data for SetIoIdStore events raised by the Ioidregistry contract.
type IoidregistrySetIoIdStoreIterator struct {
	Event *IoidregistrySetIoIdStore // Event containing the contract specifics and raw log

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
func (it *IoidregistrySetIoIdStoreIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoidregistrySetIoIdStore)
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
		it.Event = new(IoidregistrySetIoIdStore)
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
func (it *IoidregistrySetIoIdStoreIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoidregistrySetIoIdStoreIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoidregistrySetIoIdStore represents a SetIoIdStore event raised by the Ioidregistry contract.
type IoidregistrySetIoIdStore struct {
	Store common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetIoIdStore is a free log retrieval operation binding the contract event 0xc64dd2865233a317a221a0952325ee30744b237e5b6e4367fff3aeea454dfee9.
//
// Solidity: event SetIoIdStore(address indexed store)
func (_Ioidregistry *IoidregistryFilterer) FilterSetIoIdStore(opts *bind.FilterOpts, store []common.Address) (*IoidregistrySetIoIdStoreIterator, error) {

	var storeRule []interface{}
	for _, storeItem := range store {
		storeRule = append(storeRule, storeItem)
	}

	logs, sub, err := _Ioidregistry.contract.FilterLogs(opts, "SetIoIdStore", storeRule)
	if err != nil {
		return nil, err
	}
	return &IoidregistrySetIoIdStoreIterator{contract: _Ioidregistry.contract, event: "SetIoIdStore", logs: logs, sub: sub}, nil
}

// WatchSetIoIdStore is a free log subscription operation binding the contract event 0xc64dd2865233a317a221a0952325ee30744b237e5b6e4367fff3aeea454dfee9.
//
// Solidity: event SetIoIdStore(address indexed store)
func (_Ioidregistry *IoidregistryFilterer) WatchSetIoIdStore(opts *bind.WatchOpts, sink chan<- *IoidregistrySetIoIdStore, store []common.Address) (event.Subscription, error) {

	var storeRule []interface{}
	for _, storeItem := range store {
		storeRule = append(storeRule, storeItem)
	}

	logs, sub, err := _Ioidregistry.contract.WatchLogs(opts, "SetIoIdStore", storeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoidregistrySetIoIdStore)
				if err := _Ioidregistry.contract.UnpackLog(event, "SetIoIdStore", log); err != nil {
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
func (_Ioidregistry *IoidregistryFilterer) ParseSetIoIdStore(log types.Log) (*IoidregistrySetIoIdStore, error) {
	event := new(IoidregistrySetIoIdStore)
	if err := _Ioidregistry.contract.UnpackLog(event, "SetIoIdStore", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoidregistryUpdateDeviceIterator is returned from FilterUpdateDevice and is used to iterate over the raw logs and unpacked data for UpdateDevice events raised by the Ioidregistry contract.
type IoidregistryUpdateDeviceIterator struct {
	Event *IoidregistryUpdateDevice // Event containing the contract specifics and raw log

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
func (it *IoidregistryUpdateDeviceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoidregistryUpdateDevice)
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
		it.Event = new(IoidregistryUpdateDevice)
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
func (it *IoidregistryUpdateDeviceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoidregistryUpdateDeviceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoidregistryUpdateDevice represents a UpdateDevice event raised by the Ioidregistry contract.
type IoidregistryUpdateDevice struct {
	Device common.Address
	Owner  common.Address
	Hash   [32]byte
	Uri    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateDevice is a free log retrieval operation binding the contract event 0x56c4adbfac097b5bec49fa8c702ae682a2d4c10bee1ac0aad3af7f3768d0fc0f.
//
// Solidity: event UpdateDevice(address indexed device, address owner, bytes32 hash, string uri)
func (_Ioidregistry *IoidregistryFilterer) FilterUpdateDevice(opts *bind.FilterOpts, device []common.Address) (*IoidregistryUpdateDeviceIterator, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _Ioidregistry.contract.FilterLogs(opts, "UpdateDevice", deviceRule)
	if err != nil {
		return nil, err
	}
	return &IoidregistryUpdateDeviceIterator{contract: _Ioidregistry.contract, event: "UpdateDevice", logs: logs, sub: sub}, nil
}

// WatchUpdateDevice is a free log subscription operation binding the contract event 0x56c4adbfac097b5bec49fa8c702ae682a2d4c10bee1ac0aad3af7f3768d0fc0f.
//
// Solidity: event UpdateDevice(address indexed device, address owner, bytes32 hash, string uri)
func (_Ioidregistry *IoidregistryFilterer) WatchUpdateDevice(opts *bind.WatchOpts, sink chan<- *IoidregistryUpdateDevice, device []common.Address) (event.Subscription, error) {

	var deviceRule []interface{}
	for _, deviceItem := range device {
		deviceRule = append(deviceRule, deviceItem)
	}

	logs, sub, err := _Ioidregistry.contract.WatchLogs(opts, "UpdateDevice", deviceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoidregistryUpdateDevice)
				if err := _Ioidregistry.contract.UnpackLog(event, "UpdateDevice", log); err != nil {
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

// ParseUpdateDevice is a log parse operation binding the contract event 0x56c4adbfac097b5bec49fa8c702ae682a2d4c10bee1ac0aad3af7f3768d0fc0f.
//
// Solidity: event UpdateDevice(address indexed device, address owner, bytes32 hash, string uri)
func (_Ioidregistry *IoidregistryFilterer) ParseUpdateDevice(log types.Log) (*IoidregistryUpdateDevice, error) {
	event := new(IoidregistryUpdateDevice)
	if err := _Ioidregistry.contract.UnpackLog(event, "UpdateDevice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
