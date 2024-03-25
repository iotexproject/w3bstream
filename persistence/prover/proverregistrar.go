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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"proverID\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"proverID\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"proverID\",\"type\":\"string\"}],\"name\":\"ProverPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"proverID\",\"type\":\"string\"}],\"name\":\"ProverUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"proverID\",\"type\":\"string\"}],\"name\":\"ProverUpserted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_proverID\",\"type\":\"string\"}],\"name\":\"canOperateProver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"createProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"pauseProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"proverIDs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"provers\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"unpauseProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// CanOperateProver is a free data retrieval call binding the contract method 0x822e883f.
//
// Solidity: function canOperateProver(address _operator, string _proverID) view returns(bool)
func (_Prover *ProverCaller) CanOperateProver(opts *bind.CallOpts, _operator common.Address, _proverID string) (bool, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "canOperateProver", _operator, _proverID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanOperateProver is a free data retrieval call binding the contract method 0x822e883f.
//
// Solidity: function canOperateProver(address _operator, string _proverID) view returns(bool)
func (_Prover *ProverSession) CanOperateProver(_operator common.Address, _proverID string) (bool, error) {
	return _Prover.Contract.CanOperateProver(&_Prover.CallOpts, _operator, _proverID)
}

// CanOperateProver is a free data retrieval call binding the contract method 0x822e883f.
//
// Solidity: function canOperateProver(address _operator, string _proverID) view returns(bool)
func (_Prover *ProverCallerSession) CanOperateProver(_operator common.Address, _proverID string) (bool, error) {
	return _Prover.Contract.CanOperateProver(&_Prover.CallOpts, _operator, _proverID)
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

// ProverIDs is a free data retrieval call binding the contract method 0x11584323.
//
// Solidity: function proverIDs(string ) view returns(uint64)
func (_Prover *ProverCaller) ProverIDs(opts *bind.CallOpts, arg0 string) (uint64, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "proverIDs", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ProverIDs is a free data retrieval call binding the contract method 0x11584323.
//
// Solidity: function proverIDs(string ) view returns(uint64)
func (_Prover *ProverSession) ProverIDs(arg0 string) (uint64, error) {
	return _Prover.Contract.ProverIDs(&_Prover.CallOpts, arg0)
}

// ProverIDs is a free data retrieval call binding the contract method 0x11584323.
//
// Solidity: function proverIDs(string ) view returns(uint64)
func (_Prover *ProverCallerSession) ProverIDs(arg0 string) (uint64, error) {
	return _Prover.Contract.ProverIDs(&_Prover.CallOpts, arg0)
}

// Provers is a free data retrieval call binding the contract method 0x128e8834.
//
// Solidity: function provers(uint64 ) view returns(string id, bool paused)
func (_Prover *ProverCaller) Provers(opts *bind.CallOpts, arg0 uint64) (struct {
	Id     string
	Paused bool
}, error) {
	var out []interface{}
	err := _Prover.contract.Call(opts, &out, "provers", arg0)

	outstruct := new(struct {
		Id     string
		Paused bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Paused = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Provers is a free data retrieval call binding the contract method 0x128e8834.
//
// Solidity: function provers(uint64 ) view returns(string id, bool paused)
func (_Prover *ProverSession) Provers(arg0 uint64) (struct {
	Id     string
	Paused bool
}, error) {
	return _Prover.Contract.Provers(&_Prover.CallOpts, arg0)
}

// Provers is a free data retrieval call binding the contract method 0x128e8834.
//
// Solidity: function provers(uint64 ) view returns(string id, bool paused)
func (_Prover *ProverCallerSession) Provers(arg0 uint64) (struct {
	Id     string
	Paused bool
}, error) {
	return _Prover.Contract.Provers(&_Prover.CallOpts, arg0)
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

// AddOperator is a paid mutator transaction binding the contract method 0x50545d2f.
//
// Solidity: function addOperator(string _id, address _operator) returns()
func (_Prover *ProverTransactor) AddOperator(opts *bind.TransactOpts, _id string, _operator common.Address) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "addOperator", _id, _operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x50545d2f.
//
// Solidity: function addOperator(string _id, address _operator) returns()
func (_Prover *ProverSession) AddOperator(_id string, _operator common.Address) (*types.Transaction, error) {
	return _Prover.Contract.AddOperator(&_Prover.TransactOpts, _id, _operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x50545d2f.
//
// Solidity: function addOperator(string _id, address _operator) returns()
func (_Prover *ProverTransactorSession) AddOperator(_id string, _operator common.Address) (*types.Transaction, error) {
	return _Prover.Contract.AddOperator(&_Prover.TransactOpts, _id, _operator)
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

// CreateProver is a paid mutator transaction binding the contract method 0x8d21b90d.
//
// Solidity: function createProver(string _id) returns()
func (_Prover *ProverTransactor) CreateProver(opts *bind.TransactOpts, _id string) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "createProver", _id)
}

// CreateProver is a paid mutator transaction binding the contract method 0x8d21b90d.
//
// Solidity: function createProver(string _id) returns()
func (_Prover *ProverSession) CreateProver(_id string) (*types.Transaction, error) {
	return _Prover.Contract.CreateProver(&_Prover.TransactOpts, _id)
}

// CreateProver is a paid mutator transaction binding the contract method 0x8d21b90d.
//
// Solidity: function createProver(string _id) returns()
func (_Prover *ProverTransactorSession) CreateProver(_id string) (*types.Transaction, error) {
	return _Prover.Contract.CreateProver(&_Prover.TransactOpts, _id)
}

// PauseProver is a paid mutator transaction binding the contract method 0x6bc85a7b.
//
// Solidity: function pauseProver(string _id) returns()
func (_Prover *ProverTransactor) PauseProver(opts *bind.TransactOpts, _id string) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "pauseProver", _id)
}

// PauseProver is a paid mutator transaction binding the contract method 0x6bc85a7b.
//
// Solidity: function pauseProver(string _id) returns()
func (_Prover *ProverSession) PauseProver(_id string) (*types.Transaction, error) {
	return _Prover.Contract.PauseProver(&_Prover.TransactOpts, _id)
}

// PauseProver is a paid mutator transaction binding the contract method 0x6bc85a7b.
//
// Solidity: function pauseProver(string _id) returns()
func (_Prover *ProverTransactorSession) PauseProver(_id string) (*types.Transaction, error) {
	return _Prover.Contract.PauseProver(&_Prover.TransactOpts, _id)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x335eb2cd.
//
// Solidity: function removeOperator(string _id, address _operator) returns()
func (_Prover *ProverTransactor) RemoveOperator(opts *bind.TransactOpts, _id string, _operator common.Address) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "removeOperator", _id, _operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x335eb2cd.
//
// Solidity: function removeOperator(string _id, address _operator) returns()
func (_Prover *ProverSession) RemoveOperator(_id string, _operator common.Address) (*types.Transaction, error) {
	return _Prover.Contract.RemoveOperator(&_Prover.TransactOpts, _id, _operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x335eb2cd.
//
// Solidity: function removeOperator(string _id, address _operator) returns()
func (_Prover *ProverTransactorSession) RemoveOperator(_id string, _operator common.Address) (*types.Transaction, error) {
	return _Prover.Contract.RemoveOperator(&_Prover.TransactOpts, _id, _operator)
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

// UnpauseProver is a paid mutator transaction binding the contract method 0xef83a25f.
//
// Solidity: function unpauseProver(string _id) returns()
func (_Prover *ProverTransactor) UnpauseProver(opts *bind.TransactOpts, _id string) (*types.Transaction, error) {
	return _Prover.contract.Transact(opts, "unpauseProver", _id)
}

// UnpauseProver is a paid mutator transaction binding the contract method 0xef83a25f.
//
// Solidity: function unpauseProver(string _id) returns()
func (_Prover *ProverSession) UnpauseProver(_id string) (*types.Transaction, error) {
	return _Prover.Contract.UnpauseProver(&_Prover.TransactOpts, _id)
}

// UnpauseProver is a paid mutator transaction binding the contract method 0xef83a25f.
//
// Solidity: function unpauseProver(string _id) returns()
func (_Prover *ProverTransactorSession) UnpauseProver(_id string) (*types.Transaction, error) {
	return _Prover.Contract.UnpauseProver(&_Prover.TransactOpts, _id)
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

// ProverOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the Prover contract.
type ProverOperatorAddedIterator struct {
	Event *ProverOperatorAdded // Event containing the contract specifics and raw log

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
func (it *ProverOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverOperatorAdded)
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
		it.Event = new(ProverOperatorAdded)
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
func (it *ProverOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverOperatorAdded represents a OperatorAdded event raised by the Prover contract.
type ProverOperatorAdded struct {
	ProverID common.Hash
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0x451b633daf7125a63499b6bfb37960137bf0ec934a77cde7c987cd73a763ffde.
//
// Solidity: event OperatorAdded(string indexed proverID, address indexed operator)
func (_Prover *ProverFilterer) FilterOperatorAdded(opts *bind.FilterOpts, proverID []string, operator []common.Address) (*ProverOperatorAddedIterator, error) {

	var proverIDRule []interface{}
	for _, proverIDItem := range proverID {
		proverIDRule = append(proverIDRule, proverIDItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "OperatorAdded", proverIDRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ProverOperatorAddedIterator{contract: _Prover.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0x451b633daf7125a63499b6bfb37960137bf0ec934a77cde7c987cd73a763ffde.
//
// Solidity: event OperatorAdded(string indexed proverID, address indexed operator)
func (_Prover *ProverFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *ProverOperatorAdded, proverID []string, operator []common.Address) (event.Subscription, error) {

	var proverIDRule []interface{}
	for _, proverIDItem := range proverID {
		proverIDRule = append(proverIDRule, proverIDItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "OperatorAdded", proverIDRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverOperatorAdded)
				if err := _Prover.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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
// Solidity: event OperatorAdded(string indexed proverID, address indexed operator)
func (_Prover *ProverFilterer) ParseOperatorAdded(log types.Log) (*ProverOperatorAdded, error) {
	event := new(ProverOperatorAdded)
	if err := _Prover.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the Prover contract.
type ProverOperatorRemovedIterator struct {
	Event *ProverOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *ProverOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverOperatorRemoved)
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
		it.Event = new(ProverOperatorRemoved)
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
func (it *ProverOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverOperatorRemoved represents a OperatorRemoved event raised by the Prover contract.
type ProverOperatorRemoved struct {
	ProverID common.Hash
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0xbe1acd4d169dac18fbcaede8b1a4f165d9765c8a41daca9d8df4fb1c64ce600a.
//
// Solidity: event OperatorRemoved(string indexed proverID, address indexed operator)
func (_Prover *ProverFilterer) FilterOperatorRemoved(opts *bind.FilterOpts, proverID []string, operator []common.Address) (*ProverOperatorRemovedIterator, error) {

	var proverIDRule []interface{}
	for _, proverIDItem := range proverID {
		proverIDRule = append(proverIDRule, proverIDItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "OperatorRemoved", proverIDRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ProverOperatorRemovedIterator{contract: _Prover.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0xbe1acd4d169dac18fbcaede8b1a4f165d9765c8a41daca9d8df4fb1c64ce600a.
//
// Solidity: event OperatorRemoved(string indexed proverID, address indexed operator)
func (_Prover *ProverFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *ProverOperatorRemoved, proverID []string, operator []common.Address) (event.Subscription, error) {

	var proverIDRule []interface{}
	for _, proverIDItem := range proverID {
		proverIDRule = append(proverIDRule, proverIDItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "OperatorRemoved", proverIDRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverOperatorRemoved)
				if err := _Prover.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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
// Solidity: event OperatorRemoved(string indexed proverID, address indexed operator)
func (_Prover *ProverFilterer) ParseOperatorRemoved(log types.Log) (*ProverOperatorRemoved, error) {
	event := new(ProverOperatorRemoved)
	if err := _Prover.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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
	ProverID common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProverPaused is a free log retrieval operation binding the contract event 0x118fec2c7e3b902c9952ae857ad0e3d298087c64abe593cdc0d015b06e1dd2ff.
//
// Solidity: event ProverPaused(string indexed proverID)
func (_Prover *ProverFilterer) FilterProverPaused(opts *bind.FilterOpts, proverID []string) (*ProverProverPausedIterator, error) {

	var proverIDRule []interface{}
	for _, proverIDItem := range proverID {
		proverIDRule = append(proverIDRule, proverIDItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "ProverPaused", proverIDRule)
	if err != nil {
		return nil, err
	}
	return &ProverProverPausedIterator{contract: _Prover.contract, event: "ProverPaused", logs: logs, sub: sub}, nil
}

// WatchProverPaused is a free log subscription operation binding the contract event 0x118fec2c7e3b902c9952ae857ad0e3d298087c64abe593cdc0d015b06e1dd2ff.
//
// Solidity: event ProverPaused(string indexed proverID)
func (_Prover *ProverFilterer) WatchProverPaused(opts *bind.WatchOpts, sink chan<- *ProverProverPaused, proverID []string) (event.Subscription, error) {

	var proverIDRule []interface{}
	for _, proverIDItem := range proverID {
		proverIDRule = append(proverIDRule, proverIDItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "ProverPaused", proverIDRule)
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

// ParseProverPaused is a log parse operation binding the contract event 0x118fec2c7e3b902c9952ae857ad0e3d298087c64abe593cdc0d015b06e1dd2ff.
//
// Solidity: event ProverPaused(string indexed proverID)
func (_Prover *ProverFilterer) ParseProverPaused(log types.Log) (*ProverProverPaused, error) {
	event := new(ProverProverPaused)
	if err := _Prover.contract.UnpackLog(event, "ProverPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverProverUnpausedIterator is returned from FilterProverUnpaused and is used to iterate over the raw logs and unpacked data for ProverUnpaused events raised by the Prover contract.
type ProverProverUnpausedIterator struct {
	Event *ProverProverUnpaused // Event containing the contract specifics and raw log

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
func (it *ProverProverUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverProverUnpaused)
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
		it.Event = new(ProverProverUnpaused)
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
func (it *ProverProverUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverProverUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverProverUnpaused represents a ProverUnpaused event raised by the Prover contract.
type ProverProverUnpaused struct {
	ProverID common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProverUnpaused is a free log retrieval operation binding the contract event 0xe4c7f87e678a08fac1b9ea1211f9b6a8e8c885440cbb0a9df454b0e04bcab146.
//
// Solidity: event ProverUnpaused(string indexed proverID)
func (_Prover *ProverFilterer) FilterProverUnpaused(opts *bind.FilterOpts, proverID []string) (*ProverProverUnpausedIterator, error) {

	var proverIDRule []interface{}
	for _, proverIDItem := range proverID {
		proverIDRule = append(proverIDRule, proverIDItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "ProverUnpaused", proverIDRule)
	if err != nil {
		return nil, err
	}
	return &ProverProverUnpausedIterator{contract: _Prover.contract, event: "ProverUnpaused", logs: logs, sub: sub}, nil
}

// WatchProverUnpaused is a free log subscription operation binding the contract event 0xe4c7f87e678a08fac1b9ea1211f9b6a8e8c885440cbb0a9df454b0e04bcab146.
//
// Solidity: event ProverUnpaused(string indexed proverID)
func (_Prover *ProverFilterer) WatchProverUnpaused(opts *bind.WatchOpts, sink chan<- *ProverProverUnpaused, proverID []string) (event.Subscription, error) {

	var proverIDRule []interface{}
	for _, proverIDItem := range proverID {
		proverIDRule = append(proverIDRule, proverIDItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "ProverUnpaused", proverIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverProverUnpaused)
				if err := _Prover.contract.UnpackLog(event, "ProverUnpaused", log); err != nil {
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

// ParseProverUnpaused is a log parse operation binding the contract event 0xe4c7f87e678a08fac1b9ea1211f9b6a8e8c885440cbb0a9df454b0e04bcab146.
//
// Solidity: event ProverUnpaused(string indexed proverID)
func (_Prover *ProverFilterer) ParseProverUnpaused(log types.Log) (*ProverProverUnpaused, error) {
	event := new(ProverProverUnpaused)
	if err := _Prover.contract.UnpackLog(event, "ProverUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverProverUpsertedIterator is returned from FilterProverUpserted and is used to iterate over the raw logs and unpacked data for ProverUpserted events raised by the Prover contract.
type ProverProverUpsertedIterator struct {
	Event *ProverProverUpserted // Event containing the contract specifics and raw log

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
func (it *ProverProverUpsertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverProverUpserted)
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
		it.Event = new(ProverProverUpserted)
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
func (it *ProverProverUpsertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverProverUpsertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverProverUpserted represents a ProverUpserted event raised by the Prover contract.
type ProverProverUpserted struct {
	ProverID common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProverUpserted is a free log retrieval operation binding the contract event 0x159da16ce23d3ebe16324b9add1c3a4d2262a28aecafabe4b60281590b7790f4.
//
// Solidity: event ProverUpserted(string indexed proverID)
func (_Prover *ProverFilterer) FilterProverUpserted(opts *bind.FilterOpts, proverID []string) (*ProverProverUpsertedIterator, error) {

	var proverIDRule []interface{}
	for _, proverIDItem := range proverID {
		proverIDRule = append(proverIDRule, proverIDItem)
	}

	logs, sub, err := _Prover.contract.FilterLogs(opts, "ProverUpserted", proverIDRule)
	if err != nil {
		return nil, err
	}
	return &ProverProverUpsertedIterator{contract: _Prover.contract, event: "ProverUpserted", logs: logs, sub: sub}, nil
}

// WatchProverUpserted is a free log subscription operation binding the contract event 0x159da16ce23d3ebe16324b9add1c3a4d2262a28aecafabe4b60281590b7790f4.
//
// Solidity: event ProverUpserted(string indexed proverID)
func (_Prover *ProverFilterer) WatchProverUpserted(opts *bind.WatchOpts, sink chan<- *ProverProverUpserted, proverID []string) (event.Subscription, error) {

	var proverIDRule []interface{}
	for _, proverIDItem := range proverID {
		proverIDRule = append(proverIDRule, proverIDItem)
	}

	logs, sub, err := _Prover.contract.WatchLogs(opts, "ProverUpserted", proverIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverProverUpserted)
				if err := _Prover.contract.UnpackLog(event, "ProverUpserted", log); err != nil {
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

// ParseProverUpserted is a log parse operation binding the contract event 0x159da16ce23d3ebe16324b9add1c3a4d2262a28aecafabe4b60281590b7790f4.
//
// Solidity: event ProverUpserted(string indexed proverID)
func (_Prover *ProverFilterer) ParseProverUpserted(log types.Log) (*ProverProverUpserted, error) {
	event := new(ProverProverUpserted)
	if err := _Prover.contract.UnpackLog(event, "ProverUpserted", log); err != nil {
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
