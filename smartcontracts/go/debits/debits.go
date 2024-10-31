// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package debits

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

// DebitsMetaData contains all meta data concerning the Debits contract.
var DebitsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Distributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withheld\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withhold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withholdingOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DebitsABI is the input ABI used to generate the binding from.
// Deprecated: Use DebitsMetaData.ABI instead.
var DebitsABI = DebitsMetaData.ABI

// Debits is an auto generated Go binding around an Ethereum contract.
type Debits struct {
	DebitsCaller     // Read-only binding to the contract
	DebitsTransactor // Write-only binding to the contract
	DebitsFilterer   // Log filterer for contract events
}

// DebitsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DebitsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebitsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DebitsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebitsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DebitsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebitsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DebitsSession struct {
	Contract     *Debits           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DebitsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DebitsCallerSession struct {
	Contract *DebitsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DebitsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DebitsTransactorSession struct {
	Contract     *DebitsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DebitsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DebitsRaw struct {
	Contract *Debits // Generic contract binding to access the raw methods on
}

// DebitsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DebitsCallerRaw struct {
	Contract *DebitsCaller // Generic read-only contract binding to access the raw methods on
}

// DebitsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DebitsTransactorRaw struct {
	Contract *DebitsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDebits creates a new instance of Debits, bound to a specific deployed contract.
func NewDebits(address common.Address, backend bind.ContractBackend) (*Debits, error) {
	contract, err := bindDebits(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Debits{DebitsCaller: DebitsCaller{contract: contract}, DebitsTransactor: DebitsTransactor{contract: contract}, DebitsFilterer: DebitsFilterer{contract: contract}}, nil
}

// NewDebitsCaller creates a new read-only instance of Debits, bound to a specific deployed contract.
func NewDebitsCaller(address common.Address, caller bind.ContractCaller) (*DebitsCaller, error) {
	contract, err := bindDebits(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DebitsCaller{contract: contract}, nil
}

// NewDebitsTransactor creates a new write-only instance of Debits, bound to a specific deployed contract.
func NewDebitsTransactor(address common.Address, transactor bind.ContractTransactor) (*DebitsTransactor, error) {
	contract, err := bindDebits(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DebitsTransactor{contract: contract}, nil
}

// NewDebitsFilterer creates a new log filterer instance of Debits, bound to a specific deployed contract.
func NewDebitsFilterer(address common.Address, filterer bind.ContractFilterer) (*DebitsFilterer, error) {
	contract, err := bindDebits(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DebitsFilterer{contract: contract}, nil
}

// bindDebits binds a generic wrapper to an already deployed contract.
func bindDebits(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DebitsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Debits *DebitsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Debits.Contract.DebitsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Debits *DebitsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Debits.Contract.DebitsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Debits *DebitsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Debits.Contract.DebitsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Debits *DebitsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Debits.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Debits *DebitsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Debits.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Debits *DebitsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Debits.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address token, address owner) view returns(uint256)
func (_Debits *DebitsCaller) BalanceOf(opts *bind.CallOpts, token common.Address, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Debits.contract.Call(opts, &out, "balanceOf", token, owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address token, address owner) view returns(uint256)
func (_Debits *DebitsSession) BalanceOf(token common.Address, owner common.Address) (*big.Int, error) {
	return _Debits.Contract.BalanceOf(&_Debits.CallOpts, token, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address token, address owner) view returns(uint256)
func (_Debits *DebitsCallerSession) BalanceOf(token common.Address, owner common.Address) (*big.Int, error) {
	return _Debits.Contract.BalanceOf(&_Debits.CallOpts, token, owner)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Debits *DebitsCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Debits.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Debits *DebitsSession) Operator() (common.Address, error) {
	return _Debits.Contract.Operator(&_Debits.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Debits *DebitsCallerSession) Operator() (common.Address, error) {
	return _Debits.Contract.Operator(&_Debits.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Debits *DebitsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Debits.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Debits *DebitsSession) Owner() (common.Address, error) {
	return _Debits.Contract.Owner(&_Debits.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Debits *DebitsCallerSession) Owner() (common.Address, error) {
	return _Debits.Contract.Owner(&_Debits.CallOpts)
}

// WithholdingOf is a free data retrieval call binding the contract method 0x68e3e17d.
//
// Solidity: function withholdingOf(address token, address owner) view returns(uint256)
func (_Debits *DebitsCaller) WithholdingOf(opts *bind.CallOpts, token common.Address, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Debits.contract.Call(opts, &out, "withholdingOf", token, owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithholdingOf is a free data retrieval call binding the contract method 0x68e3e17d.
//
// Solidity: function withholdingOf(address token, address owner) view returns(uint256)
func (_Debits *DebitsSession) WithholdingOf(token common.Address, owner common.Address) (*big.Int, error) {
	return _Debits.Contract.WithholdingOf(&_Debits.CallOpts, token, owner)
}

// WithholdingOf is a free data retrieval call binding the contract method 0x68e3e17d.
//
// Solidity: function withholdingOf(address token, address owner) view returns(uint256)
func (_Debits *DebitsCallerSession) WithholdingOf(token common.Address, owner common.Address) (*big.Int, error) {
	return _Debits.Contract.WithholdingOf(&_Debits.CallOpts, token, owner)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Debits *DebitsTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Debits.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Debits *DebitsSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Debits.Contract.Deposit(&_Debits.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Debits *DebitsTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Debits.Contract.Deposit(&_Debits.TransactOpts, token, amount)
}

// Distribute is a paid mutator transaction binding the contract method 0x0e054260.
//
// Solidity: function distribute(address token, address _owner, address[] _recipients, uint256[] _amounts) returns()
func (_Debits *DebitsTransactor) Distribute(opts *bind.TransactOpts, token common.Address, _owner common.Address, _recipients []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Debits.contract.Transact(opts, "distribute", token, _owner, _recipients, _amounts)
}

// Distribute is a paid mutator transaction binding the contract method 0x0e054260.
//
// Solidity: function distribute(address token, address _owner, address[] _recipients, uint256[] _amounts) returns()
func (_Debits *DebitsSession) Distribute(token common.Address, _owner common.Address, _recipients []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Debits.Contract.Distribute(&_Debits.TransactOpts, token, _owner, _recipients, _amounts)
}

// Distribute is a paid mutator transaction binding the contract method 0x0e054260.
//
// Solidity: function distribute(address token, address _owner, address[] _recipients, uint256[] _amounts) returns()
func (_Debits *DebitsTransactorSession) Distribute(token common.Address, _owner common.Address, _recipients []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Debits.Contract.Distribute(&_Debits.TransactOpts, token, _owner, _recipients, _amounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Debits *DebitsTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Debits.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Debits *DebitsSession) Initialize() (*types.Transaction, error) {
	return _Debits.Contract.Initialize(&_Debits.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Debits *DebitsTransactorSession) Initialize() (*types.Transaction, error) {
	return _Debits.Contract.Initialize(&_Debits.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x0e6dfcd5.
//
// Solidity: function redeem(address token, address _owner, uint256 _amount) returns()
func (_Debits *DebitsTransactor) Redeem(opts *bind.TransactOpts, token common.Address, _owner common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Debits.contract.Transact(opts, "redeem", token, _owner, _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x0e6dfcd5.
//
// Solidity: function redeem(address token, address _owner, uint256 _amount) returns()
func (_Debits *DebitsSession) Redeem(token common.Address, _owner common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Debits.Contract.Redeem(&_Debits.TransactOpts, token, _owner, _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x0e6dfcd5.
//
// Solidity: function redeem(address token, address _owner, uint256 _amount) returns()
func (_Debits *DebitsTransactorSession) Redeem(token common.Address, _owner common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Debits.Contract.Redeem(&_Debits.TransactOpts, token, _owner, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Debits *DebitsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Debits.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Debits *DebitsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Debits.Contract.RenounceOwnership(&_Debits.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Debits *DebitsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Debits.Contract.RenounceOwnership(&_Debits.TransactOpts)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Debits *DebitsTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Debits.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Debits *DebitsSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Debits.Contract.SetOperator(&_Debits.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Debits *DebitsTransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Debits.Contract.SetOperator(&_Debits.TransactOpts, _operator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Debits *DebitsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Debits.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Debits *DebitsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Debits.Contract.TransferOwnership(&_Debits.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Debits *DebitsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Debits.Contract.TransferOwnership(&_Debits.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns()
func (_Debits *DebitsTransactor) Withdraw(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Debits.contract.Transact(opts, "withdraw", _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns()
func (_Debits *DebitsSession) Withdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Debits.Contract.Withdraw(&_Debits.TransactOpts, _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns()
func (_Debits *DebitsTransactorSession) Withdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Debits.Contract.Withdraw(&_Debits.TransactOpts, _token, _amount)
}

// Withhold is a paid mutator transaction binding the contract method 0x914e6c02.
//
// Solidity: function withhold(address token, address owner, uint256 amount) returns()
func (_Debits *DebitsTransactor) Withhold(opts *bind.TransactOpts, token common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Debits.contract.Transact(opts, "withhold", token, owner, amount)
}

// Withhold is a paid mutator transaction binding the contract method 0x914e6c02.
//
// Solidity: function withhold(address token, address owner, uint256 amount) returns()
func (_Debits *DebitsSession) Withhold(token common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Debits.Contract.Withhold(&_Debits.TransactOpts, token, owner, amount)
}

// Withhold is a paid mutator transaction binding the contract method 0x914e6c02.
//
// Solidity: function withhold(address token, address owner, uint256 amount) returns()
func (_Debits *DebitsTransactorSession) Withhold(token common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Debits.Contract.Withhold(&_Debits.TransactOpts, token, owner, amount)
}

// DebitsDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Debits contract.
type DebitsDepositedIterator struct {
	Event *DebitsDeposited // Event containing the contract specifics and raw log

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
func (it *DebitsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebitsDeposited)
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
		it.Event = new(DebitsDeposited)
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
func (it *DebitsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebitsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebitsDeposited represents a Deposited event raised by the Debits contract.
type DebitsDeposited struct {
	Token  common.Address
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed token, address indexed owner, uint256 amount)
func (_Debits *DebitsFilterer) FilterDeposited(opts *bind.FilterOpts, token []common.Address, owner []common.Address) (*DebitsDepositedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Debits.contract.FilterLogs(opts, "Deposited", tokenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &DebitsDepositedIterator{contract: _Debits.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed token, address indexed owner, uint256 amount)
func (_Debits *DebitsFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *DebitsDeposited, token []common.Address, owner []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Debits.contract.WatchLogs(opts, "Deposited", tokenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebitsDeposited)
				if err := _Debits.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed token, address indexed owner, uint256 amount)
func (_Debits *DebitsFilterer) ParseDeposited(log types.Log) (*DebitsDeposited, error) {
	event := new(DebitsDeposited)
	if err := _Debits.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebitsDistributedIterator is returned from FilterDistributed and is used to iterate over the raw logs and unpacked data for Distributed events raised by the Debits contract.
type DebitsDistributedIterator struct {
	Event *DebitsDistributed // Event containing the contract specifics and raw log

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
func (it *DebitsDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebitsDistributed)
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
		it.Event = new(DebitsDistributed)
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
func (it *DebitsDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebitsDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebitsDistributed represents a Distributed event raised by the Debits contract.
type DebitsDistributed struct {
	Token      common.Address
	Owner      common.Address
	Recipients []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDistributed is a free log retrieval operation binding the contract event 0x974cbf628e5e0af19bbf3b53fafc63ef2975a917dcf6ee7da0ac85bfda0cb3c0.
//
// Solidity: event Distributed(address indexed token, address indexed owner, address[] recipients, uint256[] amounts)
func (_Debits *DebitsFilterer) FilterDistributed(opts *bind.FilterOpts, token []common.Address, owner []common.Address) (*DebitsDistributedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Debits.contract.FilterLogs(opts, "Distributed", tokenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &DebitsDistributedIterator{contract: _Debits.contract, event: "Distributed", logs: logs, sub: sub}, nil
}

// WatchDistributed is a free log subscription operation binding the contract event 0x974cbf628e5e0af19bbf3b53fafc63ef2975a917dcf6ee7da0ac85bfda0cb3c0.
//
// Solidity: event Distributed(address indexed token, address indexed owner, address[] recipients, uint256[] amounts)
func (_Debits *DebitsFilterer) WatchDistributed(opts *bind.WatchOpts, sink chan<- *DebitsDistributed, token []common.Address, owner []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Debits.contract.WatchLogs(opts, "Distributed", tokenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebitsDistributed)
				if err := _Debits.contract.UnpackLog(event, "Distributed", log); err != nil {
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

// ParseDistributed is a log parse operation binding the contract event 0x974cbf628e5e0af19bbf3b53fafc63ef2975a917dcf6ee7da0ac85bfda0cb3c0.
//
// Solidity: event Distributed(address indexed token, address indexed owner, address[] recipients, uint256[] amounts)
func (_Debits *DebitsFilterer) ParseDistributed(log types.Log) (*DebitsDistributed, error) {
	event := new(DebitsDistributed)
	if err := _Debits.contract.UnpackLog(event, "Distributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebitsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Debits contract.
type DebitsInitializedIterator struct {
	Event *DebitsInitialized // Event containing the contract specifics and raw log

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
func (it *DebitsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebitsInitialized)
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
		it.Event = new(DebitsInitialized)
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
func (it *DebitsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebitsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebitsInitialized represents a Initialized event raised by the Debits contract.
type DebitsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Debits *DebitsFilterer) FilterInitialized(opts *bind.FilterOpts) (*DebitsInitializedIterator, error) {

	logs, sub, err := _Debits.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DebitsInitializedIterator{contract: _Debits.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Debits *DebitsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DebitsInitialized) (event.Subscription, error) {

	logs, sub, err := _Debits.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebitsInitialized)
				if err := _Debits.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Debits *DebitsFilterer) ParseInitialized(log types.Log) (*DebitsInitialized, error) {
	event := new(DebitsInitialized)
	if err := _Debits.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebitsOperatorSetIterator is returned from FilterOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorSet events raised by the Debits contract.
type DebitsOperatorSetIterator struct {
	Event *DebitsOperatorSet // Event containing the contract specifics and raw log

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
func (it *DebitsOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebitsOperatorSet)
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
		it.Event = new(DebitsOperatorSet)
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
func (it *DebitsOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebitsOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebitsOperatorSet represents a OperatorSet event raised by the Debits contract.
type DebitsOperatorSet struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSet is a free log retrieval operation binding the contract event 0x99d737e0adf2c449d71890b86772885ec7959b152ddb265f76325b6e68e105d3.
//
// Solidity: event OperatorSet(address indexed operator)
func (_Debits *DebitsFilterer) FilterOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*DebitsOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Debits.contract.FilterLogs(opts, "OperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DebitsOperatorSetIterator{contract: _Debits.contract, event: "OperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSet is a free log subscription operation binding the contract event 0x99d737e0adf2c449d71890b86772885ec7959b152ddb265f76325b6e68e105d3.
//
// Solidity: event OperatorSet(address indexed operator)
func (_Debits *DebitsFilterer) WatchOperatorSet(opts *bind.WatchOpts, sink chan<- *DebitsOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Debits.contract.WatchLogs(opts, "OperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebitsOperatorSet)
				if err := _Debits.contract.UnpackLog(event, "OperatorSet", log); err != nil {
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
// Solidity: event OperatorSet(address indexed operator)
func (_Debits *DebitsFilterer) ParseOperatorSet(log types.Log) (*DebitsOperatorSet, error) {
	event := new(DebitsOperatorSet)
	if err := _Debits.contract.UnpackLog(event, "OperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebitsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Debits contract.
type DebitsOwnershipTransferredIterator struct {
	Event *DebitsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DebitsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebitsOwnershipTransferred)
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
		it.Event = new(DebitsOwnershipTransferred)
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
func (it *DebitsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebitsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebitsOwnershipTransferred represents a OwnershipTransferred event raised by the Debits contract.
type DebitsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Debits *DebitsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DebitsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Debits.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DebitsOwnershipTransferredIterator{contract: _Debits.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Debits *DebitsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DebitsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Debits.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebitsOwnershipTransferred)
				if err := _Debits.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Debits *DebitsFilterer) ParseOwnershipTransferred(log types.Log) (*DebitsOwnershipTransferred, error) {
	event := new(DebitsOwnershipTransferred)
	if err := _Debits.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebitsRedeemedIterator is returned from FilterRedeemed and is used to iterate over the raw logs and unpacked data for Redeemed events raised by the Debits contract.
type DebitsRedeemedIterator struct {
	Event *DebitsRedeemed // Event containing the contract specifics and raw log

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
func (it *DebitsRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebitsRedeemed)
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
		it.Event = new(DebitsRedeemed)
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
func (it *DebitsRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebitsRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebitsRedeemed represents a Redeemed event raised by the Debits contract.
type DebitsRedeemed struct {
	Token  common.Address
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeemed is a free log retrieval operation binding the contract event 0x27d4634c833b7622a0acddbf7f746183625f105945e95c723ad1d5a9f2a0b6fc.
//
// Solidity: event Redeemed(address indexed token, address indexed owner, uint256 amount)
func (_Debits *DebitsFilterer) FilterRedeemed(opts *bind.FilterOpts, token []common.Address, owner []common.Address) (*DebitsRedeemedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Debits.contract.FilterLogs(opts, "Redeemed", tokenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &DebitsRedeemedIterator{contract: _Debits.contract, event: "Redeemed", logs: logs, sub: sub}, nil
}

// WatchRedeemed is a free log subscription operation binding the contract event 0x27d4634c833b7622a0acddbf7f746183625f105945e95c723ad1d5a9f2a0b6fc.
//
// Solidity: event Redeemed(address indexed token, address indexed owner, uint256 amount)
func (_Debits *DebitsFilterer) WatchRedeemed(opts *bind.WatchOpts, sink chan<- *DebitsRedeemed, token []common.Address, owner []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Debits.contract.WatchLogs(opts, "Redeemed", tokenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebitsRedeemed)
				if err := _Debits.contract.UnpackLog(event, "Redeemed", log); err != nil {
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

// ParseRedeemed is a log parse operation binding the contract event 0x27d4634c833b7622a0acddbf7f746183625f105945e95c723ad1d5a9f2a0b6fc.
//
// Solidity: event Redeemed(address indexed token, address indexed owner, uint256 amount)
func (_Debits *DebitsFilterer) ParseRedeemed(log types.Log) (*DebitsRedeemed, error) {
	event := new(DebitsRedeemed)
	if err := _Debits.contract.UnpackLog(event, "Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebitsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Debits contract.
type DebitsWithdrawnIterator struct {
	Event *DebitsWithdrawn // Event containing the contract specifics and raw log

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
func (it *DebitsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebitsWithdrawn)
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
		it.Event = new(DebitsWithdrawn)
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
func (it *DebitsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebitsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebitsWithdrawn represents a Withdrawn event raised by the Debits contract.
type DebitsWithdrawn struct {
	Token  common.Address
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed token, address indexed owner, uint256 amount)
func (_Debits *DebitsFilterer) FilterWithdrawn(opts *bind.FilterOpts, token []common.Address, owner []common.Address) (*DebitsWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Debits.contract.FilterLogs(opts, "Withdrawn", tokenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &DebitsWithdrawnIterator{contract: _Debits.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed token, address indexed owner, uint256 amount)
func (_Debits *DebitsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *DebitsWithdrawn, token []common.Address, owner []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Debits.contract.WatchLogs(opts, "Withdrawn", tokenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebitsWithdrawn)
				if err := _Debits.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed token, address indexed owner, uint256 amount)
func (_Debits *DebitsFilterer) ParseWithdrawn(log types.Log) (*DebitsWithdrawn, error) {
	event := new(DebitsWithdrawn)
	if err := _Debits.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebitsWithheldIterator is returned from FilterWithheld and is used to iterate over the raw logs and unpacked data for Withheld events raised by the Debits contract.
type DebitsWithheldIterator struct {
	Event *DebitsWithheld // Event containing the contract specifics and raw log

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
func (it *DebitsWithheldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebitsWithheld)
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
		it.Event = new(DebitsWithheld)
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
func (it *DebitsWithheldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebitsWithheldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebitsWithheld represents a Withheld event raised by the Debits contract.
type DebitsWithheld struct {
	Token  common.Address
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithheld is a free log retrieval operation binding the contract event 0x605593e759b1d56db020a6124fe8807317bdfe513fd666b5b0f4ec807559d9cd.
//
// Solidity: event Withheld(address indexed token, address indexed owner, uint256 amount)
func (_Debits *DebitsFilterer) FilterWithheld(opts *bind.FilterOpts, token []common.Address, owner []common.Address) (*DebitsWithheldIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Debits.contract.FilterLogs(opts, "Withheld", tokenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &DebitsWithheldIterator{contract: _Debits.contract, event: "Withheld", logs: logs, sub: sub}, nil
}

// WatchWithheld is a free log subscription operation binding the contract event 0x605593e759b1d56db020a6124fe8807317bdfe513fd666b5b0f4ec807559d9cd.
//
// Solidity: event Withheld(address indexed token, address indexed owner, uint256 amount)
func (_Debits *DebitsFilterer) WatchWithheld(opts *bind.WatchOpts, sink chan<- *DebitsWithheld, token []common.Address, owner []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Debits.contract.WatchLogs(opts, "Withheld", tokenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebitsWithheld)
				if err := _Debits.contract.UnpackLog(event, "Withheld", log); err != nil {
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

// ParseWithheld is a log parse operation binding the contract event 0x605593e759b1d56db020a6124fe8807317bdfe513fd666b5b0f4ec807559d9cd.
//
// Solidity: event Withheld(address indexed token, address indexed owner, uint256 amount)
func (_Debits *DebitsFilterer) ParseWithheld(log types.Log) (*DebitsWithheld, error) {
	event := new(DebitsWithheld)
	if err := _Debits.contract.UnpackLog(event, "Withheld", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
