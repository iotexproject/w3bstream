// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package znode

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

// ZnodeMetaData contains all meta data concerning the Znode contract.
var ZnodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_did\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_did\",\"type\":\"string\"}],\"name\":\"createZNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"znodeDID\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"znodeDID\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_did\",\"type\":\"string\"}],\"name\":\"pauseZNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_did\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_did\",\"type\":\"string\"}],\"name\":\"unpauseZNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"znodeDID\",\"type\":\"string\"}],\"name\":\"ZNodePaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"znodeDID\",\"type\":\"string\"}],\"name\":\"ZNodeUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"znodeDID\",\"type\":\"string\"}],\"name\":\"ZNodeUpserted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_znodeDID\",\"type\":\"string\"}],\"name\":\"canOperateZNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"znodeDIDs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"znodes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ZnodeABI is the input ABI used to generate the binding from.
// Deprecated: Use ZnodeMetaData.ABI instead.
var ZnodeABI = ZnodeMetaData.ABI

// Znode is an auto generated Go binding around an Ethereum contract.
type Znode struct {
	ZnodeCaller     // Read-only binding to the contract
	ZnodeTransactor // Write-only binding to the contract
	ZnodeFilterer   // Log filterer for contract events
}

// ZnodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZnodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZnodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZnodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZnodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZnodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZnodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZnodeSession struct {
	Contract     *Znode            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZnodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZnodeCallerSession struct {
	Contract *ZnodeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZnodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZnodeTransactorSession struct {
	Contract     *ZnodeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZnodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZnodeRaw struct {
	Contract *Znode // Generic contract binding to access the raw methods on
}

// ZnodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZnodeCallerRaw struct {
	Contract *ZnodeCaller // Generic read-only contract binding to access the raw methods on
}

// ZnodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZnodeTransactorRaw struct {
	Contract *ZnodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZnode creates a new instance of Znode, bound to a specific deployed contract.
func NewZnode(address common.Address, backend bind.ContractBackend) (*Znode, error) {
	contract, err := bindZnode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Znode{ZnodeCaller: ZnodeCaller{contract: contract}, ZnodeTransactor: ZnodeTransactor{contract: contract}, ZnodeFilterer: ZnodeFilterer{contract: contract}}, nil
}

// NewZnodeCaller creates a new read-only instance of Znode, bound to a specific deployed contract.
func NewZnodeCaller(address common.Address, caller bind.ContractCaller) (*ZnodeCaller, error) {
	contract, err := bindZnode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZnodeCaller{contract: contract}, nil
}

// NewZnodeTransactor creates a new write-only instance of Znode, bound to a specific deployed contract.
func NewZnodeTransactor(address common.Address, transactor bind.ContractTransactor) (*ZnodeTransactor, error) {
	contract, err := bindZnode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZnodeTransactor{contract: contract}, nil
}

// NewZnodeFilterer creates a new log filterer instance of Znode, bound to a specific deployed contract.
func NewZnodeFilterer(address common.Address, filterer bind.ContractFilterer) (*ZnodeFilterer, error) {
	contract, err := bindZnode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZnodeFilterer{contract: contract}, nil
}

// bindZnode binds a generic wrapper to an already deployed contract.
func bindZnode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZnodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Znode *ZnodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Znode.Contract.ZnodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Znode *ZnodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Znode.Contract.ZnodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Znode *ZnodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Znode.Contract.ZnodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Znode *ZnodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Znode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Znode *ZnodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Znode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Znode *ZnodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Znode.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Znode *ZnodeCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Znode.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Znode *ZnodeSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Znode.Contract.BalanceOf(&_Znode.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Znode *ZnodeCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Znode.Contract.BalanceOf(&_Znode.CallOpts, owner)
}

// CanOperateZNode is a free data retrieval call binding the contract method 0xb04e93c7.
//
// Solidity: function canOperateZNode(address _operator, string _znodeDID) view returns(bool)
func (_Znode *ZnodeCaller) CanOperateZNode(opts *bind.CallOpts, _operator common.Address, _znodeDID string) (bool, error) {
	var out []interface{}
	err := _Znode.contract.Call(opts, &out, "canOperateZNode", _operator, _znodeDID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanOperateZNode is a free data retrieval call binding the contract method 0xb04e93c7.
//
// Solidity: function canOperateZNode(address _operator, string _znodeDID) view returns(bool)
func (_Znode *ZnodeSession) CanOperateZNode(_operator common.Address, _znodeDID string) (bool, error) {
	return _Znode.Contract.CanOperateZNode(&_Znode.CallOpts, _operator, _znodeDID)
}

// CanOperateZNode is a free data retrieval call binding the contract method 0xb04e93c7.
//
// Solidity: function canOperateZNode(address _operator, string _znodeDID) view returns(bool)
func (_Znode *ZnodeCallerSession) CanOperateZNode(_operator common.Address, _znodeDID string) (bool, error) {
	return _Znode.Contract.CanOperateZNode(&_Znode.CallOpts, _operator, _znodeDID)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Znode *ZnodeCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Znode.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Znode *ZnodeSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Znode.Contract.GetApproved(&_Znode.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Znode *ZnodeCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Znode.Contract.GetApproved(&_Znode.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Znode *ZnodeCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Znode.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Znode *ZnodeSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Znode.Contract.IsApprovedForAll(&_Znode.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Znode *ZnodeCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Znode.Contract.IsApprovedForAll(&_Znode.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Znode *ZnodeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Znode.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Znode *ZnodeSession) Name() (string, error) {
	return _Znode.Contract.Name(&_Znode.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Znode *ZnodeCallerSession) Name() (string, error) {
	return _Znode.Contract.Name(&_Znode.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Znode *ZnodeCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Znode.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Znode *ZnodeSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Znode.Contract.OwnerOf(&_Znode.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Znode *ZnodeCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Znode.Contract.OwnerOf(&_Znode.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Znode *ZnodeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Znode.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Znode *ZnodeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Znode.Contract.SupportsInterface(&_Znode.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Znode *ZnodeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Znode.Contract.SupportsInterface(&_Znode.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Znode *ZnodeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Znode.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Znode *ZnodeSession) Symbol() (string, error) {
	return _Znode.Contract.Symbol(&_Znode.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Znode *ZnodeCallerSession) Symbol() (string, error) {
	return _Znode.Contract.Symbol(&_Znode.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Znode *ZnodeCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Znode.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Znode *ZnodeSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Znode.Contract.TokenURI(&_Znode.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Znode *ZnodeCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Znode.Contract.TokenURI(&_Znode.CallOpts, tokenId)
}

// ZnodeDIDs is a free data retrieval call binding the contract method 0xee2ea723.
//
// Solidity: function znodeDIDs(string ) view returns(uint64)
func (_Znode *ZnodeCaller) ZnodeDIDs(opts *bind.CallOpts, arg0 string) (uint64, error) {
	var out []interface{}
	err := _Znode.contract.Call(opts, &out, "znodeDIDs", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ZnodeDIDs is a free data retrieval call binding the contract method 0xee2ea723.
//
// Solidity: function znodeDIDs(string ) view returns(uint64)
func (_Znode *ZnodeSession) ZnodeDIDs(arg0 string) (uint64, error) {
	return _Znode.Contract.ZnodeDIDs(&_Znode.CallOpts, arg0)
}

// ZnodeDIDs is a free data retrieval call binding the contract method 0xee2ea723.
//
// Solidity: function znodeDIDs(string ) view returns(uint64)
func (_Znode *ZnodeCallerSession) ZnodeDIDs(arg0 string) (uint64, error) {
	return _Znode.Contract.ZnodeDIDs(&_Znode.CallOpts, arg0)
}

// Znodes is a free data retrieval call binding the contract method 0x5b85f5ec.
//
// Solidity: function znodes(uint64 ) view returns(string did, bool paused)
func (_Znode *ZnodeCaller) Znodes(opts *bind.CallOpts, arg0 uint64) (struct {
	Did    string
	Paused bool
}, error) {
	var out []interface{}
	err := _Znode.contract.Call(opts, &out, "znodes", arg0)

	outstruct := new(struct {
		Did    string
		Paused bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Did = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Paused = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Znodes is a free data retrieval call binding the contract method 0x5b85f5ec.
//
// Solidity: function znodes(uint64 ) view returns(string did, bool paused)
func (_Znode *ZnodeSession) Znodes(arg0 uint64) (struct {
	Did    string
	Paused bool
}, error) {
	return _Znode.Contract.Znodes(&_Znode.CallOpts, arg0)
}

// Znodes is a free data retrieval call binding the contract method 0x5b85f5ec.
//
// Solidity: function znodes(uint64 ) view returns(string did, bool paused)
func (_Znode *ZnodeCallerSession) Znodes(arg0 uint64) (struct {
	Did    string
	Paused bool
}, error) {
	return _Znode.Contract.Znodes(&_Znode.CallOpts, arg0)
}

// AddOperator is a paid mutator transaction binding the contract method 0x50545d2f.
//
// Solidity: function addOperator(string _did, address _operator) returns()
func (_Znode *ZnodeTransactor) AddOperator(opts *bind.TransactOpts, _did string, _operator common.Address) (*types.Transaction, error) {
	return _Znode.contract.Transact(opts, "addOperator", _did, _operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x50545d2f.
//
// Solidity: function addOperator(string _did, address _operator) returns()
func (_Znode *ZnodeSession) AddOperator(_did string, _operator common.Address) (*types.Transaction, error) {
	return _Znode.Contract.AddOperator(&_Znode.TransactOpts, _did, _operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x50545d2f.
//
// Solidity: function addOperator(string _did, address _operator) returns()
func (_Znode *ZnodeTransactorSession) AddOperator(_did string, _operator common.Address) (*types.Transaction, error) {
	return _Znode.Contract.AddOperator(&_Znode.TransactOpts, _did, _operator)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Znode *ZnodeTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Znode.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Znode *ZnodeSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Znode.Contract.Approve(&_Znode.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Znode *ZnodeTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Znode.Contract.Approve(&_Znode.TransactOpts, to, tokenId)
}

// CreateZNode is a paid mutator transaction binding the contract method 0xe522dc51.
//
// Solidity: function createZNode(string _did) returns()
func (_Znode *ZnodeTransactor) CreateZNode(opts *bind.TransactOpts, _did string) (*types.Transaction, error) {
	return _Znode.contract.Transact(opts, "createZNode", _did)
}

// CreateZNode is a paid mutator transaction binding the contract method 0xe522dc51.
//
// Solidity: function createZNode(string _did) returns()
func (_Znode *ZnodeSession) CreateZNode(_did string) (*types.Transaction, error) {
	return _Znode.Contract.CreateZNode(&_Znode.TransactOpts, _did)
}

// CreateZNode is a paid mutator transaction binding the contract method 0xe522dc51.
//
// Solidity: function createZNode(string _did) returns()
func (_Znode *ZnodeTransactorSession) CreateZNode(_did string) (*types.Transaction, error) {
	return _Znode.Contract.CreateZNode(&_Znode.TransactOpts, _did)
}

// PauseZNode is a paid mutator transaction binding the contract method 0xb0c9e9d8.
//
// Solidity: function pauseZNode(string _did) returns()
func (_Znode *ZnodeTransactor) PauseZNode(opts *bind.TransactOpts, _did string) (*types.Transaction, error) {
	return _Znode.contract.Transact(opts, "pauseZNode", _did)
}

// PauseZNode is a paid mutator transaction binding the contract method 0xb0c9e9d8.
//
// Solidity: function pauseZNode(string _did) returns()
func (_Znode *ZnodeSession) PauseZNode(_did string) (*types.Transaction, error) {
	return _Znode.Contract.PauseZNode(&_Znode.TransactOpts, _did)
}

// PauseZNode is a paid mutator transaction binding the contract method 0xb0c9e9d8.
//
// Solidity: function pauseZNode(string _did) returns()
func (_Znode *ZnodeTransactorSession) PauseZNode(_did string) (*types.Transaction, error) {
	return _Znode.Contract.PauseZNode(&_Znode.TransactOpts, _did)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x335eb2cd.
//
// Solidity: function removeOperator(string _did, address _operator) returns()
func (_Znode *ZnodeTransactor) RemoveOperator(opts *bind.TransactOpts, _did string, _operator common.Address) (*types.Transaction, error) {
	return _Znode.contract.Transact(opts, "removeOperator", _did, _operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x335eb2cd.
//
// Solidity: function removeOperator(string _did, address _operator) returns()
func (_Znode *ZnodeSession) RemoveOperator(_did string, _operator common.Address) (*types.Transaction, error) {
	return _Znode.Contract.RemoveOperator(&_Znode.TransactOpts, _did, _operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x335eb2cd.
//
// Solidity: function removeOperator(string _did, address _operator) returns()
func (_Znode *ZnodeTransactorSession) RemoveOperator(_did string, _operator common.Address) (*types.Transaction, error) {
	return _Znode.Contract.RemoveOperator(&_Znode.TransactOpts, _did, _operator)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Znode *ZnodeTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Znode.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Znode *ZnodeSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Znode.Contract.SafeTransferFrom(&_Znode.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Znode *ZnodeTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Znode.Contract.SafeTransferFrom(&_Znode.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Znode *ZnodeTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Znode.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Znode *ZnodeSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Znode.Contract.SafeTransferFrom0(&_Znode.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Znode *ZnodeTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Znode.Contract.SafeTransferFrom0(&_Znode.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Znode *ZnodeTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Znode.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Znode *ZnodeSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Znode.Contract.SetApprovalForAll(&_Znode.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Znode *ZnodeTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Znode.Contract.SetApprovalForAll(&_Znode.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Znode *ZnodeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Znode.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Znode *ZnodeSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Znode.Contract.TransferFrom(&_Znode.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Znode *ZnodeTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Znode.Contract.TransferFrom(&_Znode.TransactOpts, from, to, tokenId)
}

// UnpauseZNode is a paid mutator transaction binding the contract method 0x612aa00e.
//
// Solidity: function unpauseZNode(string _did) returns()
func (_Znode *ZnodeTransactor) UnpauseZNode(opts *bind.TransactOpts, _did string) (*types.Transaction, error) {
	return _Znode.contract.Transact(opts, "unpauseZNode", _did)
}

// UnpauseZNode is a paid mutator transaction binding the contract method 0x612aa00e.
//
// Solidity: function unpauseZNode(string _did) returns()
func (_Znode *ZnodeSession) UnpauseZNode(_did string) (*types.Transaction, error) {
	return _Znode.Contract.UnpauseZNode(&_Znode.TransactOpts, _did)
}

// UnpauseZNode is a paid mutator transaction binding the contract method 0x612aa00e.
//
// Solidity: function unpauseZNode(string _did) returns()
func (_Znode *ZnodeTransactorSession) UnpauseZNode(_did string) (*types.Transaction, error) {
	return _Znode.Contract.UnpauseZNode(&_Znode.TransactOpts, _did)
}

// ZnodeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Znode contract.
type ZnodeApprovalIterator struct {
	Event *ZnodeApproval // Event containing the contract specifics and raw log

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
func (it *ZnodeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZnodeApproval)
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
		it.Event = new(ZnodeApproval)
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
func (it *ZnodeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZnodeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZnodeApproval represents a Approval event raised by the Znode contract.
type ZnodeApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Znode *ZnodeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ZnodeApprovalIterator, error) {

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

	logs, sub, err := _Znode.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZnodeApprovalIterator{contract: _Znode.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Znode *ZnodeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZnodeApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Znode.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZnodeApproval)
				if err := _Znode.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Znode *ZnodeFilterer) ParseApproval(log types.Log) (*ZnodeApproval, error) {
	event := new(ZnodeApproval)
	if err := _Znode.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZnodeApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Znode contract.
type ZnodeApprovalForAllIterator struct {
	Event *ZnodeApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ZnodeApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZnodeApprovalForAll)
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
		it.Event = new(ZnodeApprovalForAll)
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
func (it *ZnodeApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZnodeApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZnodeApprovalForAll represents a ApprovalForAll event raised by the Znode contract.
type ZnodeApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Znode *ZnodeFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ZnodeApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Znode.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ZnodeApprovalForAllIterator{contract: _Znode.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Znode *ZnodeFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ZnodeApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Znode.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZnodeApprovalForAll)
				if err := _Znode.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Znode *ZnodeFilterer) ParseApprovalForAll(log types.Log) (*ZnodeApprovalForAll, error) {
	event := new(ZnodeApprovalForAll)
	if err := _Znode.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZnodeOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the Znode contract.
type ZnodeOperatorAddedIterator struct {
	Event *ZnodeOperatorAdded // Event containing the contract specifics and raw log

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
func (it *ZnodeOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZnodeOperatorAdded)
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
		it.Event = new(ZnodeOperatorAdded)
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
func (it *ZnodeOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZnodeOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZnodeOperatorAdded represents a OperatorAdded event raised by the Znode contract.
type ZnodeOperatorAdded struct {
	ZnodeDID common.Hash
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0x451b633daf7125a63499b6bfb37960137bf0ec934a77cde7c987cd73a763ffde.
//
// Solidity: event OperatorAdded(string indexed znodeDID, address indexed operator)
func (_Znode *ZnodeFilterer) FilterOperatorAdded(opts *bind.FilterOpts, znodeDID []string, operator []common.Address) (*ZnodeOperatorAddedIterator, error) {

	var znodeDIDRule []interface{}
	for _, znodeDIDItem := range znodeDID {
		znodeDIDRule = append(znodeDIDRule, znodeDIDItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Znode.contract.FilterLogs(opts, "OperatorAdded", znodeDIDRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ZnodeOperatorAddedIterator{contract: _Znode.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0x451b633daf7125a63499b6bfb37960137bf0ec934a77cde7c987cd73a763ffde.
//
// Solidity: event OperatorAdded(string indexed znodeDID, address indexed operator)
func (_Znode *ZnodeFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *ZnodeOperatorAdded, znodeDID []string, operator []common.Address) (event.Subscription, error) {

	var znodeDIDRule []interface{}
	for _, znodeDIDItem := range znodeDID {
		znodeDIDRule = append(znodeDIDRule, znodeDIDItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Znode.contract.WatchLogs(opts, "OperatorAdded", znodeDIDRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZnodeOperatorAdded)
				if err := _Znode.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0x451b633daf7125a63499b6bfb37960137bf0ec934a77cde7c987cd73a763ffde.
//
// Solidity: event OperatorAdded(string indexed znodeDID, address indexed operator)
func (_Znode *ZnodeFilterer) ParseOperatorAdded(log types.Log) (*ZnodeOperatorAdded, error) {
	event := new(ZnodeOperatorAdded)
	if err := _Znode.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZnodeOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the Znode contract.
type ZnodeOperatorRemovedIterator struct {
	Event *ZnodeOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *ZnodeOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZnodeOperatorRemoved)
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
		it.Event = new(ZnodeOperatorRemoved)
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
func (it *ZnodeOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZnodeOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZnodeOperatorRemoved represents a OperatorRemoved event raised by the Znode contract.
type ZnodeOperatorRemoved struct {
	ZnodeDID common.Hash
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0xbe1acd4d169dac18fbcaede8b1a4f165d9765c8a41daca9d8df4fb1c64ce600a.
//
// Solidity: event OperatorRemoved(string indexed znodeDID, address indexed operator)
func (_Znode *ZnodeFilterer) FilterOperatorRemoved(opts *bind.FilterOpts, znodeDID []string, operator []common.Address) (*ZnodeOperatorRemovedIterator, error) {

	var znodeDIDRule []interface{}
	for _, znodeDIDItem := range znodeDID {
		znodeDIDRule = append(znodeDIDRule, znodeDIDItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Znode.contract.FilterLogs(opts, "OperatorRemoved", znodeDIDRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ZnodeOperatorRemovedIterator{contract: _Znode.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0xbe1acd4d169dac18fbcaede8b1a4f165d9765c8a41daca9d8df4fb1c64ce600a.
//
// Solidity: event OperatorRemoved(string indexed znodeDID, address indexed operator)
func (_Znode *ZnodeFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *ZnodeOperatorRemoved, znodeDID []string, operator []common.Address) (event.Subscription, error) {

	var znodeDIDRule []interface{}
	for _, znodeDIDItem := range znodeDID {
		znodeDIDRule = append(znodeDIDRule, znodeDIDItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Znode.contract.WatchLogs(opts, "OperatorRemoved", znodeDIDRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZnodeOperatorRemoved)
				if err := _Znode.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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

// ParseOperatorRemoved is a log parse operation binding the contract event 0xbe1acd4d169dac18fbcaede8b1a4f165d9765c8a41daca9d8df4fb1c64ce600a.
//
// Solidity: event OperatorRemoved(string indexed znodeDID, address indexed operator)
func (_Znode *ZnodeFilterer) ParseOperatorRemoved(log types.Log) (*ZnodeOperatorRemoved, error) {
	event := new(ZnodeOperatorRemoved)
	if err := _Znode.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZnodeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Znode contract.
type ZnodeTransferIterator struct {
	Event *ZnodeTransfer // Event containing the contract specifics and raw log

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
func (it *ZnodeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZnodeTransfer)
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
		it.Event = new(ZnodeTransfer)
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
func (it *ZnodeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZnodeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZnodeTransfer represents a Transfer event raised by the Znode contract.
type ZnodeTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Znode *ZnodeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ZnodeTransferIterator, error) {

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

	logs, sub, err := _Znode.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZnodeTransferIterator{contract: _Znode.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Znode *ZnodeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZnodeTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Znode.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZnodeTransfer)
				if err := _Znode.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Znode *ZnodeFilterer) ParseTransfer(log types.Log) (*ZnodeTransfer, error) {
	event := new(ZnodeTransfer)
	if err := _Znode.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZnodeZNodePausedIterator is returned from FilterZNodePaused and is used to iterate over the raw logs and unpacked data for ZNodePaused events raised by the Znode contract.
type ZnodeZNodePausedIterator struct {
	Event *ZnodeZNodePaused // Event containing the contract specifics and raw log

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
func (it *ZnodeZNodePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZnodeZNodePaused)
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
		it.Event = new(ZnodeZNodePaused)
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
func (it *ZnodeZNodePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZnodeZNodePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZnodeZNodePaused represents a ZNodePaused event raised by the Znode contract.
type ZnodeZNodePaused struct {
	ZnodeDID common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterZNodePaused is a free log retrieval operation binding the contract event 0x3b3591532ac11876cdfe50eeaa1e4a653370539aa7b2061a222517f12d5d610e.
//
// Solidity: event ZNodePaused(string indexed znodeDID)
func (_Znode *ZnodeFilterer) FilterZNodePaused(opts *bind.FilterOpts, znodeDID []string) (*ZnodeZNodePausedIterator, error) {

	var znodeDIDRule []interface{}
	for _, znodeDIDItem := range znodeDID {
		znodeDIDRule = append(znodeDIDRule, znodeDIDItem)
	}

	logs, sub, err := _Znode.contract.FilterLogs(opts, "ZNodePaused", znodeDIDRule)
	if err != nil {
		return nil, err
	}
	return &ZnodeZNodePausedIterator{contract: _Znode.contract, event: "ZNodePaused", logs: logs, sub: sub}, nil
}

// WatchZNodePaused is a free log subscription operation binding the contract event 0x3b3591532ac11876cdfe50eeaa1e4a653370539aa7b2061a222517f12d5d610e.
//
// Solidity: event ZNodePaused(string indexed znodeDID)
func (_Znode *ZnodeFilterer) WatchZNodePaused(opts *bind.WatchOpts, sink chan<- *ZnodeZNodePaused, znodeDID []string) (event.Subscription, error) {

	var znodeDIDRule []interface{}
	for _, znodeDIDItem := range znodeDID {
		znodeDIDRule = append(znodeDIDRule, znodeDIDItem)
	}

	logs, sub, err := _Znode.contract.WatchLogs(opts, "ZNodePaused", znodeDIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZnodeZNodePaused)
				if err := _Znode.contract.UnpackLog(event, "ZNodePaused", log); err != nil {
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

// ParseZNodePaused is a log parse operation binding the contract event 0x3b3591532ac11876cdfe50eeaa1e4a653370539aa7b2061a222517f12d5d610e.
//
// Solidity: event ZNodePaused(string indexed znodeDID)
func (_Znode *ZnodeFilterer) ParseZNodePaused(log types.Log) (*ZnodeZNodePaused, error) {
	event := new(ZnodeZNodePaused)
	if err := _Znode.contract.UnpackLog(event, "ZNodePaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZnodeZNodeUnpausedIterator is returned from FilterZNodeUnpaused and is used to iterate over the raw logs and unpacked data for ZNodeUnpaused events raised by the Znode contract.
type ZnodeZNodeUnpausedIterator struct {
	Event *ZnodeZNodeUnpaused // Event containing the contract specifics and raw log

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
func (it *ZnodeZNodeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZnodeZNodeUnpaused)
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
		it.Event = new(ZnodeZNodeUnpaused)
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
func (it *ZnodeZNodeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZnodeZNodeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZnodeZNodeUnpaused represents a ZNodeUnpaused event raised by the Znode contract.
type ZnodeZNodeUnpaused struct {
	ZnodeDID common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterZNodeUnpaused is a free log retrieval operation binding the contract event 0x7940133dd0d674bd27bae54738b9b150d03482f500fe908e943309807e6a8a62.
//
// Solidity: event ZNodeUnpaused(string indexed znodeDID)
func (_Znode *ZnodeFilterer) FilterZNodeUnpaused(opts *bind.FilterOpts, znodeDID []string) (*ZnodeZNodeUnpausedIterator, error) {

	var znodeDIDRule []interface{}
	for _, znodeDIDItem := range znodeDID {
		znodeDIDRule = append(znodeDIDRule, znodeDIDItem)
	}

	logs, sub, err := _Znode.contract.FilterLogs(opts, "ZNodeUnpaused", znodeDIDRule)
	if err != nil {
		return nil, err
	}
	return &ZnodeZNodeUnpausedIterator{contract: _Znode.contract, event: "ZNodeUnpaused", logs: logs, sub: sub}, nil
}

// WatchZNodeUnpaused is a free log subscription operation binding the contract event 0x7940133dd0d674bd27bae54738b9b150d03482f500fe908e943309807e6a8a62.
//
// Solidity: event ZNodeUnpaused(string indexed znodeDID)
func (_Znode *ZnodeFilterer) WatchZNodeUnpaused(opts *bind.WatchOpts, sink chan<- *ZnodeZNodeUnpaused, znodeDID []string) (event.Subscription, error) {

	var znodeDIDRule []interface{}
	for _, znodeDIDItem := range znodeDID {
		znodeDIDRule = append(znodeDIDRule, znodeDIDItem)
	}

	logs, sub, err := _Znode.contract.WatchLogs(opts, "ZNodeUnpaused", znodeDIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZnodeZNodeUnpaused)
				if err := _Znode.contract.UnpackLog(event, "ZNodeUnpaused", log); err != nil {
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

// ParseZNodeUnpaused is a log parse operation binding the contract event 0x7940133dd0d674bd27bae54738b9b150d03482f500fe908e943309807e6a8a62.
//
// Solidity: event ZNodeUnpaused(string indexed znodeDID)
func (_Znode *ZnodeFilterer) ParseZNodeUnpaused(log types.Log) (*ZnodeZNodeUnpaused, error) {
	event := new(ZnodeZNodeUnpaused)
	if err := _Znode.contract.UnpackLog(event, "ZNodeUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZnodeZNodeUpsertedIterator is returned from FilterZNodeUpserted and is used to iterate over the raw logs and unpacked data for ZNodeUpserted events raised by the Znode contract.
type ZnodeZNodeUpsertedIterator struct {
	Event *ZnodeZNodeUpserted // Event containing the contract specifics and raw log

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
func (it *ZnodeZNodeUpsertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZnodeZNodeUpserted)
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
		it.Event = new(ZnodeZNodeUpserted)
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
func (it *ZnodeZNodeUpsertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZnodeZNodeUpsertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZnodeZNodeUpserted represents a ZNodeUpserted event raised by the Znode contract.
type ZnodeZNodeUpserted struct {
	ZnodeDID common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterZNodeUpserted is a free log retrieval operation binding the contract event 0x5d67f654709f9147ec5d8c04b24191c2b359fce2fa0a04ab282ad4132cd6cddb.
//
// Solidity: event ZNodeUpserted(string indexed znodeDID)
func (_Znode *ZnodeFilterer) FilterZNodeUpserted(opts *bind.FilterOpts, znodeDID []string) (*ZnodeZNodeUpsertedIterator, error) {

	var znodeDIDRule []interface{}
	for _, znodeDIDItem := range znodeDID {
		znodeDIDRule = append(znodeDIDRule, znodeDIDItem)
	}

	logs, sub, err := _Znode.contract.FilterLogs(opts, "ZNodeUpserted", znodeDIDRule)
	if err != nil {
		return nil, err
	}
	return &ZnodeZNodeUpsertedIterator{contract: _Znode.contract, event: "ZNodeUpserted", logs: logs, sub: sub}, nil
}

// WatchZNodeUpserted is a free log subscription operation binding the contract event 0x5d67f654709f9147ec5d8c04b24191c2b359fce2fa0a04ab282ad4132cd6cddb.
//
// Solidity: event ZNodeUpserted(string indexed znodeDID)
func (_Znode *ZnodeFilterer) WatchZNodeUpserted(opts *bind.WatchOpts, sink chan<- *ZnodeZNodeUpserted, znodeDID []string) (event.Subscription, error) {

	var znodeDIDRule []interface{}
	for _, znodeDIDItem := range znodeDID {
		znodeDIDRule = append(znodeDIDRule, znodeDIDItem)
	}

	logs, sub, err := _Znode.contract.WatchLogs(opts, "ZNodeUpserted", znodeDIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZnodeZNodeUpserted)
				if err := _Znode.contract.UnpackLog(event, "ZNodeUpserted", log); err != nil {
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

// ParseZNodeUpserted is a log parse operation binding the contract event 0x5d67f654709f9147ec5d8c04b24191c2b359fce2fa0a04ab282ad4132cd6cddb.
//
// Solidity: event ZNodeUpserted(string indexed znodeDID)
func (_Znode *ZnodeFilterer) ParseZNodeUpserted(log types.Log) (*ZnodeZNodeUpserted, error) {
	event := new(ZnodeZNodeUpserted)
	if err := _Znode.contract.UnpackLog(event, "ZNodeUpserted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
