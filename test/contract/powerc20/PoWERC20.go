// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package powerc20

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

// Powerc20MetaData contains all meta data concerning the Powerc20 contract.
var Powerc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_decimals_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_miningLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initialLimitPerMint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"difficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimitPerMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRemainingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitPerMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"mine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minedNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miningLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"miningTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupplyCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"uint256ToFr\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Powerc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Powerc20MetaData.ABI instead.
var Powerc20ABI = Powerc20MetaData.ABI

// Powerc20 is an auto generated Go binding around an Ethereum contract.
type Powerc20 struct {
	Powerc20Caller     // Read-only binding to the contract
	Powerc20Transactor // Write-only binding to the contract
	Powerc20Filterer   // Log filterer for contract events
}

// Powerc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Powerc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Powerc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Powerc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Powerc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Powerc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Powerc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Powerc20Session struct {
	Contract     *Powerc20         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Powerc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Powerc20CallerSession struct {
	Contract *Powerc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Powerc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Powerc20TransactorSession struct {
	Contract     *Powerc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Powerc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Powerc20Raw struct {
	Contract *Powerc20 // Generic contract binding to access the raw methods on
}

// Powerc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Powerc20CallerRaw struct {
	Contract *Powerc20Caller // Generic read-only contract binding to access the raw methods on
}

// Powerc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Powerc20TransactorRaw struct {
	Contract *Powerc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPowerc20 creates a new instance of Powerc20, bound to a specific deployed contract.
func NewPowerc20(address common.Address, backend bind.ContractBackend) (*Powerc20, error) {
	contract, err := bindPowerc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Powerc20{Powerc20Caller: Powerc20Caller{contract: contract}, Powerc20Transactor: Powerc20Transactor{contract: contract}, Powerc20Filterer: Powerc20Filterer{contract: contract}}, nil
}

// NewPowerc20Caller creates a new read-only instance of Powerc20, bound to a specific deployed contract.
func NewPowerc20Caller(address common.Address, caller bind.ContractCaller) (*Powerc20Caller, error) {
	contract, err := bindPowerc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Powerc20Caller{contract: contract}, nil
}

// NewPowerc20Transactor creates a new write-only instance of Powerc20, bound to a specific deployed contract.
func NewPowerc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Powerc20Transactor, error) {
	contract, err := bindPowerc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Powerc20Transactor{contract: contract}, nil
}

// NewPowerc20Filterer creates a new log filterer instance of Powerc20, bound to a specific deployed contract.
func NewPowerc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Powerc20Filterer, error) {
	contract, err := bindPowerc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Powerc20Filterer{contract: contract}, nil
}

// bindPowerc20 binds a generic wrapper to an already deployed contract.
func bindPowerc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Powerc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Powerc20 *Powerc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Powerc20.Contract.Powerc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Powerc20 *Powerc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Powerc20.Contract.Powerc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Powerc20 *Powerc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Powerc20.Contract.Powerc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Powerc20 *Powerc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Powerc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Powerc20 *Powerc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Powerc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Powerc20 *Powerc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Powerc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Powerc20 *Powerc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Powerc20 *Powerc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Powerc20.Contract.Allowance(&_Powerc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Powerc20 *Powerc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Powerc20.Contract.Allowance(&_Powerc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Powerc20 *Powerc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Powerc20 *Powerc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _Powerc20.Contract.BalanceOf(&_Powerc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Powerc20 *Powerc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Powerc20.Contract.BalanceOf(&_Powerc20.CallOpts, account)
}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(uint256)
func (_Powerc20 *Powerc20Caller) Challenge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "challenge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(uint256)
func (_Powerc20 *Powerc20Session) Challenge() (*big.Int, error) {
	return _Powerc20.Contract.Challenge(&_Powerc20.CallOpts)
}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(uint256)
func (_Powerc20 *Powerc20CallerSession) Challenge() (*big.Int, error) {
	return _Powerc20.Contract.Challenge(&_Powerc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Powerc20 *Powerc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Powerc20 *Powerc20Session) Decimals() (uint8, error) {
	return _Powerc20.Contract.Decimals(&_Powerc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Powerc20 *Powerc20CallerSession) Decimals() (uint8, error) {
	return _Powerc20.Contract.Decimals(&_Powerc20.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_Powerc20 *Powerc20Caller) Difficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "difficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_Powerc20 *Powerc20Session) Difficulty() (*big.Int, error) {
	return _Powerc20.Contract.Difficulty(&_Powerc20.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_Powerc20 *Powerc20CallerSession) Difficulty() (*big.Int, error) {
	return _Powerc20.Contract.Difficulty(&_Powerc20.CallOpts)
}

// GetLimitPerMint is a free data retrieval call binding the contract method 0xb32e82c0.
//
// Solidity: function getLimitPerMint() view returns(uint256)
func (_Powerc20 *Powerc20Caller) GetLimitPerMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "getLimitPerMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLimitPerMint is a free data retrieval call binding the contract method 0xb32e82c0.
//
// Solidity: function getLimitPerMint() view returns(uint256)
func (_Powerc20 *Powerc20Session) GetLimitPerMint() (*big.Int, error) {
	return _Powerc20.Contract.GetLimitPerMint(&_Powerc20.CallOpts)
}

// GetLimitPerMint is a free data retrieval call binding the contract method 0xb32e82c0.
//
// Solidity: function getLimitPerMint() view returns(uint256)
func (_Powerc20 *Powerc20CallerSession) GetLimitPerMint() (*big.Int, error) {
	return _Powerc20.Contract.GetLimitPerMint(&_Powerc20.CallOpts)
}

// GetRemainingSupply is a free data retrieval call binding the contract method 0xe4b7fb73.
//
// Solidity: function getRemainingSupply() view returns(uint256)
func (_Powerc20 *Powerc20Caller) GetRemainingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "getRemainingSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRemainingSupply is a free data retrieval call binding the contract method 0xe4b7fb73.
//
// Solidity: function getRemainingSupply() view returns(uint256)
func (_Powerc20 *Powerc20Session) GetRemainingSupply() (*big.Int, error) {
	return _Powerc20.Contract.GetRemainingSupply(&_Powerc20.CallOpts)
}

// GetRemainingSupply is a free data retrieval call binding the contract method 0xe4b7fb73.
//
// Solidity: function getRemainingSupply() view returns(uint256)
func (_Powerc20 *Powerc20CallerSession) GetRemainingSupply() (*big.Int, error) {
	return _Powerc20.Contract.GetRemainingSupply(&_Powerc20.CallOpts)
}

// LimitPerMint is a free data retrieval call binding the contract method 0xe2ce9f51.
//
// Solidity: function limitPerMint() view returns(uint256)
func (_Powerc20 *Powerc20Caller) LimitPerMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "limitPerMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitPerMint is a free data retrieval call binding the contract method 0xe2ce9f51.
//
// Solidity: function limitPerMint() view returns(uint256)
func (_Powerc20 *Powerc20Session) LimitPerMint() (*big.Int, error) {
	return _Powerc20.Contract.LimitPerMint(&_Powerc20.CallOpts)
}

// LimitPerMint is a free data retrieval call binding the contract method 0xe2ce9f51.
//
// Solidity: function limitPerMint() view returns(uint256)
func (_Powerc20 *Powerc20CallerSession) LimitPerMint() (*big.Int, error) {
	return _Powerc20.Contract.LimitPerMint(&_Powerc20.CallOpts)
}

// MinedNonces is a free data retrieval call binding the contract method 0x342a252a.
//
// Solidity: function minedNonces(address , uint256 ) view returns(bool)
func (_Powerc20 *Powerc20Caller) MinedNonces(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "minedNonces", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinedNonces is a free data retrieval call binding the contract method 0x342a252a.
//
// Solidity: function minedNonces(address , uint256 ) view returns(bool)
func (_Powerc20 *Powerc20Session) MinedNonces(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Powerc20.Contract.MinedNonces(&_Powerc20.CallOpts, arg0, arg1)
}

// MinedNonces is a free data retrieval call binding the contract method 0x342a252a.
//
// Solidity: function minedNonces(address , uint256 ) view returns(bool)
func (_Powerc20 *Powerc20CallerSession) MinedNonces(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Powerc20.Contract.MinedNonces(&_Powerc20.CallOpts, arg0, arg1)
}

// MiningLimit is a free data retrieval call binding the contract method 0xc2651503.
//
// Solidity: function miningLimit() view returns(uint256)
func (_Powerc20 *Powerc20Caller) MiningLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "miningLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningLimit is a free data retrieval call binding the contract method 0xc2651503.
//
// Solidity: function miningLimit() view returns(uint256)
func (_Powerc20 *Powerc20Session) MiningLimit() (*big.Int, error) {
	return _Powerc20.Contract.MiningLimit(&_Powerc20.CallOpts)
}

// MiningLimit is a free data retrieval call binding the contract method 0xc2651503.
//
// Solidity: function miningLimit() view returns(uint256)
func (_Powerc20 *Powerc20CallerSession) MiningLimit() (*big.Int, error) {
	return _Powerc20.Contract.MiningLimit(&_Powerc20.CallOpts)
}

// MiningTimes is a free data retrieval call binding the contract method 0x2719881e.
//
// Solidity: function miningTimes(address ) view returns(uint256)
func (_Powerc20 *Powerc20Caller) MiningTimes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "miningTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningTimes is a free data retrieval call binding the contract method 0x2719881e.
//
// Solidity: function miningTimes(address ) view returns(uint256)
func (_Powerc20 *Powerc20Session) MiningTimes(arg0 common.Address) (*big.Int, error) {
	return _Powerc20.Contract.MiningTimes(&_Powerc20.CallOpts, arg0)
}

// MiningTimes is a free data retrieval call binding the contract method 0x2719881e.
//
// Solidity: function miningTimes(address ) view returns(uint256)
func (_Powerc20 *Powerc20CallerSession) MiningTimes(arg0 common.Address) (*big.Int, error) {
	return _Powerc20.Contract.MiningTimes(&_Powerc20.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Powerc20 *Powerc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Powerc20 *Powerc20Session) Name() (string, error) {
	return _Powerc20.Contract.Name(&_Powerc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Powerc20 *Powerc20CallerSession) Name() (string, error) {
	return _Powerc20.Contract.Name(&_Powerc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Powerc20 *Powerc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Powerc20 *Powerc20Session) Symbol() (string, error) {
	return _Powerc20.Contract.Symbol(&_Powerc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Powerc20 *Powerc20CallerSession) Symbol() (string, error) {
	return _Powerc20.Contract.Symbol(&_Powerc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Powerc20 *Powerc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Powerc20 *Powerc20Session) TotalSupply() (*big.Int, error) {
	return _Powerc20.Contract.TotalSupply(&_Powerc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Powerc20 *Powerc20CallerSession) TotalSupply() (*big.Int, error) {
	return _Powerc20.Contract.TotalSupply(&_Powerc20.CallOpts)
}

// TotalSupplyCap is a free data retrieval call binding the contract method 0xbb102aea.
//
// Solidity: function totalSupplyCap() view returns(uint256)
func (_Powerc20 *Powerc20Caller) TotalSupplyCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "totalSupplyCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyCap is a free data retrieval call binding the contract method 0xbb102aea.
//
// Solidity: function totalSupplyCap() view returns(uint256)
func (_Powerc20 *Powerc20Session) TotalSupplyCap() (*big.Int, error) {
	return _Powerc20.Contract.TotalSupplyCap(&_Powerc20.CallOpts)
}

// TotalSupplyCap is a free data retrieval call binding the contract method 0xbb102aea.
//
// Solidity: function totalSupplyCap() view returns(uint256)
func (_Powerc20 *Powerc20CallerSession) TotalSupplyCap() (*big.Int, error) {
	return _Powerc20.Contract.TotalSupplyCap(&_Powerc20.CallOpts)
}

// Uint256ToFr is a free data retrieval call binding the contract method 0xbd228486.
//
// Solidity: function uint256ToFr(uint256 _value) pure returns(bytes32)
func (_Powerc20 *Powerc20Caller) Uint256ToFr(opts *bind.CallOpts, _value *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "uint256ToFr", _value)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Uint256ToFr is a free data retrieval call binding the contract method 0xbd228486.
//
// Solidity: function uint256ToFr(uint256 _value) pure returns(bytes32)
func (_Powerc20 *Powerc20Session) Uint256ToFr(_value *big.Int) ([32]byte, error) {
	return _Powerc20.Contract.Uint256ToFr(&_Powerc20.CallOpts, _value)
}

// Uint256ToFr is a free data retrieval call binding the contract method 0xbd228486.
//
// Solidity: function uint256ToFr(uint256 _value) pure returns(bytes32)
func (_Powerc20 *Powerc20CallerSession) Uint256ToFr(_value *big.Int) ([32]byte, error) {
	return _Powerc20.Contract.Uint256ToFr(&_Powerc20.CallOpts, _value)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Powerc20 *Powerc20Caller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Powerc20.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Powerc20 *Powerc20Session) Verifier() (common.Address, error) {
	return _Powerc20.Contract.Verifier(&_Powerc20.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Powerc20 *Powerc20CallerSession) Verifier() (common.Address, error) {
	return _Powerc20.Contract.Verifier(&_Powerc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Powerc20 *Powerc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Powerc20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Powerc20 *Powerc20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Powerc20.Contract.Approve(&_Powerc20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Powerc20 *Powerc20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Powerc20.Contract.Approve(&_Powerc20.TransactOpts, spender, value)
}

// Mine is a paid mutator transaction binding the contract method 0x050ce37a.
//
// Solidity: function mine(uint256 nonce, address sender, bytes proof) returns()
func (_Powerc20 *Powerc20Transactor) Mine(opts *bind.TransactOpts, nonce *big.Int, sender common.Address, proof []byte) (*types.Transaction, error) {
	return _Powerc20.contract.Transact(opts, "mine", nonce, sender, proof)
}

// Mine is a paid mutator transaction binding the contract method 0x050ce37a.
//
// Solidity: function mine(uint256 nonce, address sender, bytes proof) returns()
func (_Powerc20 *Powerc20Session) Mine(nonce *big.Int, sender common.Address, proof []byte) (*types.Transaction, error) {
	return _Powerc20.Contract.Mine(&_Powerc20.TransactOpts, nonce, sender, proof)
}

// Mine is a paid mutator transaction binding the contract method 0x050ce37a.
//
// Solidity: function mine(uint256 nonce, address sender, bytes proof) returns()
func (_Powerc20 *Powerc20TransactorSession) Mine(nonce *big.Int, sender common.Address, proof []byte) (*types.Transaction, error) {
	return _Powerc20.Contract.Mine(&_Powerc20.TransactOpts, nonce, sender, proof)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Powerc20 *Powerc20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Powerc20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Powerc20 *Powerc20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Powerc20.Contract.Transfer(&_Powerc20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Powerc20 *Powerc20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Powerc20.Contract.Transfer(&_Powerc20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Powerc20 *Powerc20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Powerc20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Powerc20 *Powerc20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Powerc20.Contract.TransferFrom(&_Powerc20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Powerc20 *Powerc20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Powerc20.Contract.TransferFrom(&_Powerc20.TransactOpts, from, to, value)
}

// Powerc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Powerc20 contract.
type Powerc20ApprovalIterator struct {
	Event *Powerc20Approval // Event containing the contract specifics and raw log

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
func (it *Powerc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Powerc20Approval)
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
		it.Event = new(Powerc20Approval)
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
func (it *Powerc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Powerc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Powerc20Approval represents a Approval event raised by the Powerc20 contract.
type Powerc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Powerc20 *Powerc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Powerc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Powerc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Powerc20ApprovalIterator{contract: _Powerc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Powerc20 *Powerc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Powerc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Powerc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Powerc20Approval)
				if err := _Powerc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Powerc20 *Powerc20Filterer) ParseApproval(log types.Log) (*Powerc20Approval, error) {
	event := new(Powerc20Approval)
	if err := _Powerc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Powerc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Powerc20 contract.
type Powerc20TransferIterator struct {
	Event *Powerc20Transfer // Event containing the contract specifics and raw log

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
func (it *Powerc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Powerc20Transfer)
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
		it.Event = new(Powerc20Transfer)
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
func (it *Powerc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Powerc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Powerc20Transfer represents a Transfer event raised by the Powerc20 contract.
type Powerc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Powerc20 *Powerc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Powerc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Powerc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Powerc20TransferIterator{contract: _Powerc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Powerc20 *Powerc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Powerc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Powerc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Powerc20Transfer)
				if err := _Powerc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Powerc20 *Powerc20Filterer) ParseTransfer(log types.Log) (*Powerc20Transfer, error) {
	event := new(Powerc20Transfer)
	if err := _Powerc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
