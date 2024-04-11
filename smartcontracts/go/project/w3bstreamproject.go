// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package project

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

// W3bstreamProjectProjectConfig is an auto generated low-level Go binding around an user-defined struct.
type W3bstreamProjectProjectConfig struct {
	Uri  string
	Hash [32]byte
}

// ProjectMetaData contains all meta data concerning the Project contract.
var ProjectMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"AttributeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MinterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ProjectConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"}],\"name\":\"ProjectPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"}],\"name\":\"ProjectResumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"attribute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"attributes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_keys\",\"type\":\"bytes32[]\"}],\"name\":\"attributesOf\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"values_\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"config\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"internalType\":\"structW3bstreamProject.ProjectConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"projectId_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_projectId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"_keys\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_values\",\"type\":\"bytes[]\"}],\"name\":\"setAttributes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"updateConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProjectABI is the input ABI used to generate the binding from.
// Deprecated: Use ProjectMetaData.ABI instead.
var ProjectABI = ProjectMetaData.ABI

// Project is an auto generated Go binding around an Ethereum contract.
type Project struct {
	ProjectCaller     // Read-only binding to the contract
	ProjectTransactor // Write-only binding to the contract
	ProjectFilterer   // Log filterer for contract events
}

// ProjectCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProjectCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProjectTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProjectFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProjectSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProjectSession struct {
	Contract     *Project          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProjectCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProjectCallerSession struct {
	Contract *ProjectCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ProjectTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProjectTransactorSession struct {
	Contract     *ProjectTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ProjectRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProjectRaw struct {
	Contract *Project // Generic contract binding to access the raw methods on
}

// ProjectCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProjectCallerRaw struct {
	Contract *ProjectCaller // Generic read-only contract binding to access the raw methods on
}

// ProjectTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProjectTransactorRaw struct {
	Contract *ProjectTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProject creates a new instance of Project, bound to a specific deployed contract.
func NewProject(address common.Address, backend bind.ContractBackend) (*Project, error) {
	contract, err := bindProject(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Project{ProjectCaller: ProjectCaller{contract: contract}, ProjectTransactor: ProjectTransactor{contract: contract}, ProjectFilterer: ProjectFilterer{contract: contract}}, nil
}

// NewProjectCaller creates a new read-only instance of Project, bound to a specific deployed contract.
func NewProjectCaller(address common.Address, caller bind.ContractCaller) (*ProjectCaller, error) {
	contract, err := bindProject(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectCaller{contract: contract}, nil
}

// NewProjectTransactor creates a new write-only instance of Project, bound to a specific deployed contract.
func NewProjectTransactor(address common.Address, transactor bind.ContractTransactor) (*ProjectTransactor, error) {
	contract, err := bindProject(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProjectTransactor{contract: contract}, nil
}

// NewProjectFilterer creates a new log filterer instance of Project, bound to a specific deployed contract.
func NewProjectFilterer(address common.Address, filterer bind.ContractFilterer) (*ProjectFilterer, error) {
	contract, err := bindProject(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProjectFilterer{contract: contract}, nil
}

// bindProject binds a generic wrapper to an already deployed contract.
func bindProject(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProjectMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Project *ProjectRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Project.Contract.ProjectCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Project *ProjectRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Project.Contract.ProjectTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Project *ProjectRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Project.Contract.ProjectTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Project *ProjectCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Project.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Project *ProjectTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Project.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Project *ProjectTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Project.Contract.contract.Transact(opts, method, params...)
}

// Attribute is a free data retrieval call binding the contract method 0x40341e4b.
//
// Solidity: function attribute(uint256 _projectId, bytes32 _name) view returns(bytes)
func (_Project *ProjectCaller) Attribute(opts *bind.CallOpts, _projectId *big.Int, _name [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "attribute", _projectId, _name)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Attribute is a free data retrieval call binding the contract method 0x40341e4b.
//
// Solidity: function attribute(uint256 _projectId, bytes32 _name) view returns(bytes)
func (_Project *ProjectSession) Attribute(_projectId *big.Int, _name [32]byte) ([]byte, error) {
	return _Project.Contract.Attribute(&_Project.CallOpts, _projectId, _name)
}

// Attribute is a free data retrieval call binding the contract method 0x40341e4b.
//
// Solidity: function attribute(uint256 _projectId, bytes32 _name) view returns(bytes)
func (_Project *ProjectCallerSession) Attribute(_projectId *big.Int, _name [32]byte) ([]byte, error) {
	return _Project.Contract.Attribute(&_Project.CallOpts, _projectId, _name)
}

// Attributes is a free data retrieval call binding the contract method 0x73c773db.
//
// Solidity: function attributes(uint256 , bytes32 ) view returns(bytes)
func (_Project *ProjectCaller) Attributes(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "attributes", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Attributes is a free data retrieval call binding the contract method 0x73c773db.
//
// Solidity: function attributes(uint256 , bytes32 ) view returns(bytes)
func (_Project *ProjectSession) Attributes(arg0 *big.Int, arg1 [32]byte) ([]byte, error) {
	return _Project.Contract.Attributes(&_Project.CallOpts, arg0, arg1)
}

// Attributes is a free data retrieval call binding the contract method 0x73c773db.
//
// Solidity: function attributes(uint256 , bytes32 ) view returns(bytes)
func (_Project *ProjectCallerSession) Attributes(arg0 *big.Int, arg1 [32]byte) ([]byte, error) {
	return _Project.Contract.Attributes(&_Project.CallOpts, arg0, arg1)
}

// AttributesOf is a free data retrieval call binding the contract method 0xe76d621c.
//
// Solidity: function attributesOf(uint256 _projectId, bytes32[] _keys) view returns(bytes[] values_)
func (_Project *ProjectCaller) AttributesOf(opts *bind.CallOpts, _projectId *big.Int, _keys [][32]byte) ([][]byte, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "attributesOf", _projectId, _keys)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// AttributesOf is a free data retrieval call binding the contract method 0xe76d621c.
//
// Solidity: function attributesOf(uint256 _projectId, bytes32[] _keys) view returns(bytes[] values_)
func (_Project *ProjectSession) AttributesOf(_projectId *big.Int, _keys [][32]byte) ([][]byte, error) {
	return _Project.Contract.AttributesOf(&_Project.CallOpts, _projectId, _keys)
}

// AttributesOf is a free data retrieval call binding the contract method 0xe76d621c.
//
// Solidity: function attributesOf(uint256 _projectId, bytes32[] _keys) view returns(bytes[] values_)
func (_Project *ProjectCallerSession) AttributesOf(_projectId *big.Int, _keys [][32]byte) ([][]byte, error) {
	return _Project.Contract.AttributesOf(&_Project.CallOpts, _projectId, _keys)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Project *ProjectCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Project *ProjectSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Project.Contract.BalanceOf(&_Project.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Project *ProjectCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Project.Contract.BalanceOf(&_Project.CallOpts, owner)
}

// Config is a free data retrieval call binding the contract method 0x84691767.
//
// Solidity: function config(uint256 _projectId) view returns((string,bytes32))
func (_Project *ProjectCaller) Config(opts *bind.CallOpts, _projectId *big.Int) (W3bstreamProjectProjectConfig, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "config", _projectId)

	if err != nil {
		return *new(W3bstreamProjectProjectConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(W3bstreamProjectProjectConfig)).(*W3bstreamProjectProjectConfig)

	return out0, err

}

// Config is a free data retrieval call binding the contract method 0x84691767.
//
// Solidity: function config(uint256 _projectId) view returns((string,bytes32))
func (_Project *ProjectSession) Config(_projectId *big.Int) (W3bstreamProjectProjectConfig, error) {
	return _Project.Contract.Config(&_Project.CallOpts, _projectId)
}

// Config is a free data retrieval call binding the contract method 0x84691767.
//
// Solidity: function config(uint256 _projectId) view returns((string,bytes32))
func (_Project *ProjectCallerSession) Config(_projectId *big.Int) (W3bstreamProjectProjectConfig, error) {
	return _Project.Contract.Config(&_Project.CallOpts, _projectId)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Project *ProjectCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Project *ProjectSession) Count() (*big.Int, error) {
	return _Project.Contract.Count(&_Project.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Project *ProjectCallerSession) Count() (*big.Int, error) {
	return _Project.Contract.Count(&_Project.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Project *ProjectCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Project *ProjectSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Project.Contract.GetApproved(&_Project.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Project *ProjectCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Project.Contract.GetApproved(&_Project.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Project *ProjectCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Project *ProjectSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Project.Contract.IsApprovedForAll(&_Project.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Project *ProjectCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Project.Contract.IsApprovedForAll(&_Project.CallOpts, owner, operator)
}

// IsPaused is a free data retrieval call binding the contract method 0xbdf2a43c.
//
// Solidity: function isPaused(uint256 _projectId) view returns(bool)
func (_Project *ProjectCaller) IsPaused(opts *bind.CallOpts, _projectId *big.Int) (bool, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "isPaused", _projectId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xbdf2a43c.
//
// Solidity: function isPaused(uint256 _projectId) view returns(bool)
func (_Project *ProjectSession) IsPaused(_projectId *big.Int) (bool, error) {
	return _Project.Contract.IsPaused(&_Project.CallOpts, _projectId)
}

// IsPaused is a free data retrieval call binding the contract method 0xbdf2a43c.
//
// Solidity: function isPaused(uint256 _projectId) view returns(bool)
func (_Project *ProjectCallerSession) IsPaused(_projectId *big.Int) (bool, error) {
	return _Project.Contract.IsPaused(&_Project.CallOpts, _projectId)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Project *ProjectCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Project *ProjectSession) Minter() (common.Address, error) {
	return _Project.Contract.Minter(&_Project.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Project *ProjectCallerSession) Minter() (common.Address, error) {
	return _Project.Contract.Minter(&_Project.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Project *ProjectCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Project *ProjectSession) Name() (string, error) {
	return _Project.Contract.Name(&_Project.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Project *ProjectCallerSession) Name() (string, error) {
	return _Project.Contract.Name(&_Project.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Project *ProjectCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Project *ProjectSession) Owner() (common.Address, error) {
	return _Project.Contract.Owner(&_Project.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Project *ProjectCallerSession) Owner() (common.Address, error) {
	return _Project.Contract.Owner(&_Project.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Project *ProjectCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Project *ProjectSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Project.Contract.OwnerOf(&_Project.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Project *ProjectCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Project.Contract.OwnerOf(&_Project.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Project *ProjectCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Project *ProjectSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Project.Contract.SupportsInterface(&_Project.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Project *ProjectCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Project.Contract.SupportsInterface(&_Project.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Project *ProjectCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Project *ProjectSession) Symbol() (string, error) {
	return _Project.Contract.Symbol(&_Project.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Project *ProjectCallerSession) Symbol() (string, error) {
	return _Project.Contract.Symbol(&_Project.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Project *ProjectCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Project.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Project *ProjectSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Project.Contract.TokenURI(&_Project.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Project *ProjectCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Project.Contract.TokenURI(&_Project.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Project *ProjectTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Project *ProjectSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Project.Contract.Approve(&_Project.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Project *ProjectTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Project.Contract.Approve(&_Project.TransactOpts, to, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(address _minter, string _name, string _symbol) returns()
func (_Project *ProjectTransactor) Initialize(opts *bind.TransactOpts, _minter common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "initialize", _minter, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(address _minter, string _name, string _symbol) returns()
func (_Project *ProjectSession) Initialize(_minter common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _Project.Contract.Initialize(&_Project.TransactOpts, _minter, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(address _minter, string _name, string _symbol) returns()
func (_Project *ProjectTransactorSession) Initialize(_minter common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _Project.Contract.Initialize(&_Project.TransactOpts, _minter, _name, _symbol)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _owner) returns(uint256 projectId_)
func (_Project *ProjectTransactor) Mint(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "mint", _owner)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _owner) returns(uint256 projectId_)
func (_Project *ProjectSession) Mint(_owner common.Address) (*types.Transaction, error) {
	return _Project.Contract.Mint(&_Project.TransactOpts, _owner)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _owner) returns(uint256 projectId_)
func (_Project *ProjectTransactorSession) Mint(_owner common.Address) (*types.Transaction, error) {
	return _Project.Contract.Mint(&_Project.TransactOpts, _owner)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 _projectId) returns()
func (_Project *ProjectTransactor) Pause(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "pause", _projectId)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 _projectId) returns()
func (_Project *ProjectSession) Pause(_projectId *big.Int) (*types.Transaction, error) {
	return _Project.Contract.Pause(&_Project.TransactOpts, _projectId)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 _projectId) returns()
func (_Project *ProjectTransactorSession) Pause(_projectId *big.Int) (*types.Transaction, error) {
	return _Project.Contract.Pause(&_Project.TransactOpts, _projectId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Project *ProjectTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Project *ProjectSession) RenounceOwnership() (*types.Transaction, error) {
	return _Project.Contract.RenounceOwnership(&_Project.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Project *ProjectTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Project.Contract.RenounceOwnership(&_Project.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x414000b5.
//
// Solidity: function resume(uint256 _projectId) returns()
func (_Project *ProjectTransactor) Resume(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "resume", _projectId)
}

// Resume is a paid mutator transaction binding the contract method 0x414000b5.
//
// Solidity: function resume(uint256 _projectId) returns()
func (_Project *ProjectSession) Resume(_projectId *big.Int) (*types.Transaction, error) {
	return _Project.Contract.Resume(&_Project.TransactOpts, _projectId)
}

// Resume is a paid mutator transaction binding the contract method 0x414000b5.
//
// Solidity: function resume(uint256 _projectId) returns()
func (_Project *ProjectTransactorSession) Resume(_projectId *big.Int) (*types.Transaction, error) {
	return _Project.Contract.Resume(&_Project.TransactOpts, _projectId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Project *ProjectTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Project *ProjectSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Project.Contract.SafeTransferFrom(&_Project.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Project *ProjectTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Project.Contract.SafeTransferFrom(&_Project.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Project *ProjectTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Project *ProjectSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Project.Contract.SafeTransferFrom0(&_Project.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Project *ProjectTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Project.Contract.SafeTransferFrom0(&_Project.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Project *ProjectTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Project *ProjectSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Project.Contract.SetApprovalForAll(&_Project.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Project *ProjectTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Project.Contract.SetApprovalForAll(&_Project.TransactOpts, operator, approved)
}

// SetAttributes is a paid mutator transaction binding the contract method 0xe24ce5b4.
//
// Solidity: function setAttributes(uint64 _projectId, bytes32[] _keys, bytes[] _values) returns()
func (_Project *ProjectTransactor) SetAttributes(opts *bind.TransactOpts, _projectId uint64, _keys [][32]byte, _values [][]byte) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "setAttributes", _projectId, _keys, _values)
}

// SetAttributes is a paid mutator transaction binding the contract method 0xe24ce5b4.
//
// Solidity: function setAttributes(uint64 _projectId, bytes32[] _keys, bytes[] _values) returns()
func (_Project *ProjectSession) SetAttributes(_projectId uint64, _keys [][32]byte, _values [][]byte) (*types.Transaction, error) {
	return _Project.Contract.SetAttributes(&_Project.TransactOpts, _projectId, _keys, _values)
}

// SetAttributes is a paid mutator transaction binding the contract method 0xe24ce5b4.
//
// Solidity: function setAttributes(uint64 _projectId, bytes32[] _keys, bytes[] _values) returns()
func (_Project *ProjectTransactorSession) SetAttributes(_projectId uint64, _keys [][32]byte, _values [][]byte) (*types.Transaction, error) {
	return _Project.Contract.SetAttributes(&_Project.TransactOpts, _projectId, _keys, _values)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Project *ProjectTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Project *ProjectSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Project.Contract.SetMinter(&_Project.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Project *ProjectTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Project.Contract.SetMinter(&_Project.TransactOpts, _minter)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Project *ProjectTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Project *ProjectSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Project.Contract.TransferFrom(&_Project.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Project *ProjectTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Project.Contract.TransferFrom(&_Project.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Project *ProjectTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Project *ProjectSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Project.Contract.TransferOwnership(&_Project.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Project *ProjectTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Project.Contract.TransferOwnership(&_Project.TransactOpts, newOwner)
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x4b597aa3.
//
// Solidity: function updateConfig(uint256 _projectId, string _uri, bytes32 _hash) returns()
func (_Project *ProjectTransactor) UpdateConfig(opts *bind.TransactOpts, _projectId *big.Int, _uri string, _hash [32]byte) (*types.Transaction, error) {
	return _Project.contract.Transact(opts, "updateConfig", _projectId, _uri, _hash)
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x4b597aa3.
//
// Solidity: function updateConfig(uint256 _projectId, string _uri, bytes32 _hash) returns()
func (_Project *ProjectSession) UpdateConfig(_projectId *big.Int, _uri string, _hash [32]byte) (*types.Transaction, error) {
	return _Project.Contract.UpdateConfig(&_Project.TransactOpts, _projectId, _uri, _hash)
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x4b597aa3.
//
// Solidity: function updateConfig(uint256 _projectId, string _uri, bytes32 _hash) returns()
func (_Project *ProjectTransactorSession) UpdateConfig(_projectId *big.Int, _uri string, _hash [32]byte) (*types.Transaction, error) {
	return _Project.Contract.UpdateConfig(&_Project.TransactOpts, _projectId, _uri, _hash)
}

// ProjectApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Project contract.
type ProjectApprovalIterator struct {
	Event *ProjectApproval // Event containing the contract specifics and raw log

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
func (it *ProjectApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectApproval)
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
		it.Event = new(ProjectApproval)
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
func (it *ProjectApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectApproval represents a Approval event raised by the Project contract.
type ProjectApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Project *ProjectFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ProjectApprovalIterator, error) {

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

	logs, sub, err := _Project.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ProjectApprovalIterator{contract: _Project.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Project *ProjectFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ProjectApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Project.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectApproval)
				if err := _Project.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Project *ProjectFilterer) ParseApproval(log types.Log) (*ProjectApproval, error) {
	event := new(ProjectApproval)
	if err := _Project.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Project contract.
type ProjectApprovalForAllIterator struct {
	Event *ProjectApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ProjectApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectApprovalForAll)
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
		it.Event = new(ProjectApprovalForAll)
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
func (it *ProjectApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectApprovalForAll represents a ApprovalForAll event raised by the Project contract.
type ProjectApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Project *ProjectFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ProjectApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Project.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ProjectApprovalForAllIterator{contract: _Project.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Project *ProjectFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ProjectApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Project.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectApprovalForAll)
				if err := _Project.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Project *ProjectFilterer) ParseApprovalForAll(log types.Log) (*ProjectApprovalForAll, error) {
	event := new(ProjectApprovalForAll)
	if err := _Project.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectAttributeSetIterator is returned from FilterAttributeSet and is used to iterate over the raw logs and unpacked data for AttributeSet events raised by the Project contract.
type ProjectAttributeSetIterator struct {
	Event *ProjectAttributeSet // Event containing the contract specifics and raw log

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
func (it *ProjectAttributeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectAttributeSet)
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
		it.Event = new(ProjectAttributeSet)
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
func (it *ProjectAttributeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectAttributeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectAttributeSet represents a AttributeSet event raised by the Project contract.
type ProjectAttributeSet struct {
	ProjectId *big.Int
	Key       [32]byte
	Value     []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAttributeSet is a free log retrieval operation binding the contract event 0x840db4c564ec8ec61fd9377b125346993b20659d558d2e066e33c588b60f9fc3.
//
// Solidity: event AttributeSet(uint256 indexed projectId, bytes32 indexed key, bytes value)
func (_Project *ProjectFilterer) FilterAttributeSet(opts *bind.FilterOpts, projectId []*big.Int, key [][32]byte) (*ProjectAttributeSetIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Project.contract.FilterLogs(opts, "AttributeSet", projectIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &ProjectAttributeSetIterator{contract: _Project.contract, event: "AttributeSet", logs: logs, sub: sub}, nil
}

// WatchAttributeSet is a free log subscription operation binding the contract event 0x840db4c564ec8ec61fd9377b125346993b20659d558d2e066e33c588b60f9fc3.
//
// Solidity: event AttributeSet(uint256 indexed projectId, bytes32 indexed key, bytes value)
func (_Project *ProjectFilterer) WatchAttributeSet(opts *bind.WatchOpts, sink chan<- *ProjectAttributeSet, projectId []*big.Int, key [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Project.contract.WatchLogs(opts, "AttributeSet", projectIdRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectAttributeSet)
				if err := _Project.contract.UnpackLog(event, "AttributeSet", log); err != nil {
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

// ParseAttributeSet is a log parse operation binding the contract event 0x840db4c564ec8ec61fd9377b125346993b20659d558d2e066e33c588b60f9fc3.
//
// Solidity: event AttributeSet(uint256 indexed projectId, bytes32 indexed key, bytes value)
func (_Project *ProjectFilterer) ParseAttributeSet(log types.Log) (*ProjectAttributeSet, error) {
	event := new(ProjectAttributeSet)
	if err := _Project.contract.UnpackLog(event, "AttributeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Project contract.
type ProjectInitializedIterator struct {
	Event *ProjectInitialized // Event containing the contract specifics and raw log

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
func (it *ProjectInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectInitialized)
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
		it.Event = new(ProjectInitialized)
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
func (it *ProjectInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectInitialized represents a Initialized event raised by the Project contract.
type ProjectInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Project *ProjectFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProjectInitializedIterator, error) {

	logs, sub, err := _Project.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProjectInitializedIterator{contract: _Project.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Project *ProjectFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProjectInitialized) (event.Subscription, error) {

	logs, sub, err := _Project.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectInitialized)
				if err := _Project.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Project *ProjectFilterer) ParseInitialized(log types.Log) (*ProjectInitialized, error) {
	event := new(ProjectInitialized)
	if err := _Project.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectMinterSetIterator is returned from FilterMinterSet and is used to iterate over the raw logs and unpacked data for MinterSet events raised by the Project contract.
type ProjectMinterSetIterator struct {
	Event *ProjectMinterSet // Event containing the contract specifics and raw log

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
func (it *ProjectMinterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectMinterSet)
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
		it.Event = new(ProjectMinterSet)
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
func (it *ProjectMinterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectMinterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectMinterSet represents a MinterSet event raised by the Project contract.
type ProjectMinterSet struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterSet is a free log retrieval operation binding the contract event 0x726b590ef91a8c76ad05bbe91a57ef84605276528f49cd47d787f558a4e755b6.
//
// Solidity: event MinterSet(address indexed minter)
func (_Project *ProjectFilterer) FilterMinterSet(opts *bind.FilterOpts, minter []common.Address) (*ProjectMinterSetIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Project.contract.FilterLogs(opts, "MinterSet", minterRule)
	if err != nil {
		return nil, err
	}
	return &ProjectMinterSetIterator{contract: _Project.contract, event: "MinterSet", logs: logs, sub: sub}, nil
}

// WatchMinterSet is a free log subscription operation binding the contract event 0x726b590ef91a8c76ad05bbe91a57ef84605276528f49cd47d787f558a4e755b6.
//
// Solidity: event MinterSet(address indexed minter)
func (_Project *ProjectFilterer) WatchMinterSet(opts *bind.WatchOpts, sink chan<- *ProjectMinterSet, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Project.contract.WatchLogs(opts, "MinterSet", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectMinterSet)
				if err := _Project.contract.UnpackLog(event, "MinterSet", log); err != nil {
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
// Solidity: event MinterSet(address indexed minter)
func (_Project *ProjectFilterer) ParseMinterSet(log types.Log) (*ProjectMinterSet, error) {
	event := new(ProjectMinterSet)
	if err := _Project.contract.UnpackLog(event, "MinterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Project contract.
type ProjectOwnershipTransferredIterator struct {
	Event *ProjectOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProjectOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectOwnershipTransferred)
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
		it.Event = new(ProjectOwnershipTransferred)
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
func (it *ProjectOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectOwnershipTransferred represents a OwnershipTransferred event raised by the Project contract.
type ProjectOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Project *ProjectFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProjectOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Project.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProjectOwnershipTransferredIterator{contract: _Project.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Project *ProjectFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProjectOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Project.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectOwnershipTransferred)
				if err := _Project.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Project *ProjectFilterer) ParseOwnershipTransferred(log types.Log) (*ProjectOwnershipTransferred, error) {
	event := new(ProjectOwnershipTransferred)
	if err := _Project.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectProjectConfigUpdatedIterator is returned from FilterProjectConfigUpdated and is used to iterate over the raw logs and unpacked data for ProjectConfigUpdated events raised by the Project contract.
type ProjectProjectConfigUpdatedIterator struct {
	Event *ProjectProjectConfigUpdated // Event containing the contract specifics and raw log

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
func (it *ProjectProjectConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectProjectConfigUpdated)
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
		it.Event = new(ProjectProjectConfigUpdated)
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
func (it *ProjectProjectConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectProjectConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectProjectConfigUpdated represents a ProjectConfigUpdated event raised by the Project contract.
type ProjectProjectConfigUpdated struct {
	ProjectId *big.Int
	Uri       string
	Hash      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectConfigUpdated is a free log retrieval operation binding the contract event 0xa9ee0c223bc138bec6ebb21e09d00d5423fc3bbc210bdb6aef9d190b0641aecb.
//
// Solidity: event ProjectConfigUpdated(uint256 indexed projectId, string uri, bytes32 hash)
func (_Project *ProjectFilterer) FilterProjectConfigUpdated(opts *bind.FilterOpts, projectId []*big.Int) (*ProjectProjectConfigUpdatedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Project.contract.FilterLogs(opts, "ProjectConfigUpdated", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ProjectProjectConfigUpdatedIterator{contract: _Project.contract, event: "ProjectConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchProjectConfigUpdated is a free log subscription operation binding the contract event 0xa9ee0c223bc138bec6ebb21e09d00d5423fc3bbc210bdb6aef9d190b0641aecb.
//
// Solidity: event ProjectConfigUpdated(uint256 indexed projectId, string uri, bytes32 hash)
func (_Project *ProjectFilterer) WatchProjectConfigUpdated(opts *bind.WatchOpts, sink chan<- *ProjectProjectConfigUpdated, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Project.contract.WatchLogs(opts, "ProjectConfigUpdated", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectProjectConfigUpdated)
				if err := _Project.contract.UnpackLog(event, "ProjectConfigUpdated", log); err != nil {
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

// ParseProjectConfigUpdated is a log parse operation binding the contract event 0xa9ee0c223bc138bec6ebb21e09d00d5423fc3bbc210bdb6aef9d190b0641aecb.
//
// Solidity: event ProjectConfigUpdated(uint256 indexed projectId, string uri, bytes32 hash)
func (_Project *ProjectFilterer) ParseProjectConfigUpdated(log types.Log) (*ProjectProjectConfigUpdated, error) {
	event := new(ProjectProjectConfigUpdated)
	if err := _Project.contract.UnpackLog(event, "ProjectConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectProjectPausedIterator is returned from FilterProjectPaused and is used to iterate over the raw logs and unpacked data for ProjectPaused events raised by the Project contract.
type ProjectProjectPausedIterator struct {
	Event *ProjectProjectPaused // Event containing the contract specifics and raw log

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
func (it *ProjectProjectPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectProjectPaused)
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
		it.Event = new(ProjectProjectPaused)
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
func (it *ProjectProjectPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectProjectPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectProjectPaused represents a ProjectPaused event raised by the Project contract.
type ProjectProjectPaused struct {
	ProjectId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectPaused is a free log retrieval operation binding the contract event 0x9f505f325627bdd7f5a6dd8bcceecdc48a989f647561427d61d35b7a50703f79.
//
// Solidity: event ProjectPaused(uint256 indexed projectId)
func (_Project *ProjectFilterer) FilterProjectPaused(opts *bind.FilterOpts, projectId []*big.Int) (*ProjectProjectPausedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Project.contract.FilterLogs(opts, "ProjectPaused", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ProjectProjectPausedIterator{contract: _Project.contract, event: "ProjectPaused", logs: logs, sub: sub}, nil
}

// WatchProjectPaused is a free log subscription operation binding the contract event 0x9f505f325627bdd7f5a6dd8bcceecdc48a989f647561427d61d35b7a50703f79.
//
// Solidity: event ProjectPaused(uint256 indexed projectId)
func (_Project *ProjectFilterer) WatchProjectPaused(opts *bind.WatchOpts, sink chan<- *ProjectProjectPaused, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Project.contract.WatchLogs(opts, "ProjectPaused", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectProjectPaused)
				if err := _Project.contract.UnpackLog(event, "ProjectPaused", log); err != nil {
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

// ParseProjectPaused is a log parse operation binding the contract event 0x9f505f325627bdd7f5a6dd8bcceecdc48a989f647561427d61d35b7a50703f79.
//
// Solidity: event ProjectPaused(uint256 indexed projectId)
func (_Project *ProjectFilterer) ParseProjectPaused(log types.Log) (*ProjectProjectPaused, error) {
	event := new(ProjectProjectPaused)
	if err := _Project.contract.UnpackLog(event, "ProjectPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectProjectResumedIterator is returned from FilterProjectResumed and is used to iterate over the raw logs and unpacked data for ProjectResumed events raised by the Project contract.
type ProjectProjectResumedIterator struct {
	Event *ProjectProjectResumed // Event containing the contract specifics and raw log

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
func (it *ProjectProjectResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectProjectResumed)
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
		it.Event = new(ProjectProjectResumed)
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
func (it *ProjectProjectResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectProjectResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectProjectResumed represents a ProjectResumed event raised by the Project contract.
type ProjectProjectResumed struct {
	ProjectId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectResumed is a free log retrieval operation binding the contract event 0x8c936416fd11c0291d9c7f69aad7e8847c5228b15a3969a1ac6a1c7bf394cd75.
//
// Solidity: event ProjectResumed(uint256 indexed projectId)
func (_Project *ProjectFilterer) FilterProjectResumed(opts *bind.FilterOpts, projectId []*big.Int) (*ProjectProjectResumedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Project.contract.FilterLogs(opts, "ProjectResumed", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ProjectProjectResumedIterator{contract: _Project.contract, event: "ProjectResumed", logs: logs, sub: sub}, nil
}

// WatchProjectResumed is a free log subscription operation binding the contract event 0x8c936416fd11c0291d9c7f69aad7e8847c5228b15a3969a1ac6a1c7bf394cd75.
//
// Solidity: event ProjectResumed(uint256 indexed projectId)
func (_Project *ProjectFilterer) WatchProjectResumed(opts *bind.WatchOpts, sink chan<- *ProjectProjectResumed, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Project.contract.WatchLogs(opts, "ProjectResumed", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectProjectResumed)
				if err := _Project.contract.UnpackLog(event, "ProjectResumed", log); err != nil {
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

// ParseProjectResumed is a log parse operation binding the contract event 0x8c936416fd11c0291d9c7f69aad7e8847c5228b15a3969a1ac6a1c7bf394cd75.
//
// Solidity: event ProjectResumed(uint256 indexed projectId)
func (_Project *ProjectFilterer) ParseProjectResumed(log types.Log) (*ProjectProjectResumed, error) {
	event := new(ProjectProjectResumed)
	if err := _Project.contract.UnpackLog(event, "ProjectResumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProjectTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Project contract.
type ProjectTransferIterator struct {
	Event *ProjectTransfer // Event containing the contract specifics and raw log

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
func (it *ProjectTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProjectTransfer)
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
		it.Event = new(ProjectTransfer)
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
func (it *ProjectTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProjectTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProjectTransfer represents a Transfer event raised by the Project contract.
type ProjectTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Project *ProjectFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ProjectTransferIterator, error) {

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

	logs, sub, err := _Project.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ProjectTransferIterator{contract: _Project.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Project *ProjectFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ProjectTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Project.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProjectTransfer)
				if err := _Project.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Project *ProjectFilterer) ParseTransfer(log types.Log) (*ProjectTransfer, error) {
	event := new(ProjectTransfer)
	if err := _Project.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
