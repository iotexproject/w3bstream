// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockerc20

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

// Mockerc20MetaData contains all meta data concerning the Mockerc20 contract.
var Mockerc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"initialAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60406080815234620003e65762000f60803803806200001e81620003eb565b9283398101608082820312620003e65781516001600160401b0390818111620003e657826200004f91850162000411565b906020928385015190828211620003e6576200006d91860162000411565b848601516001600160a01b0381169591929190869003620003e6576060015192805191808311620002e65760038054936001938486811c96168015620003db575b89871014620003c5578190601f968781116200036f575b5089908783116001146200030857600092620002fc575b505060001982841b1c191690841b1781555b8451918211620002e65760049485548481811c91168015620002db575b89821014620002c6578581116200027b575b508790858411600114620002105793839491849260009562000204575b50501b92600019911b1c19161782555b8415620001c2575060025490828201809211620001ad575060025560008381528083528481208054830190558451918252917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a351610adc9081620004848239f35b601190634e487b7160e01b6000525260246000fd5b855162461bcd60e51b815291820184905260248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260649150fd5b0151935038806200013a565b9190601f1984169287600052848a6000209460005b8c8983831062000263575050501062000248575b50505050811b0182556200014a565b01519060f884600019921b161c191690553880808062000239565b86860151895590970196948501948893500162000225565b86600052886000208680860160051c8201928b8710620002bc575b0160051c019085905b828110620002af5750506200011d565b600081550185906200029f565b9250819262000296565b602287634e487b7160e01b6000525260246000fd5b90607f16906200010b565b634e487b7160e01b600052604160045260246000fd5b015190503880620000dc565b90869350601f19831691856000528b6000209260005b8d8282106200035857505084116200033f575b505050811b018155620000ee565b015160001983861b60f8161c1916905538808062000331565b8385015186558a979095019493840193016200031e565b90915083600052896000208780850160051c8201928c8610620003bb575b918891869594930160051c01915b828110620003ab575050620000c5565b600081558594508891016200039b565b925081926200038d565b634e487b7160e01b600052602260045260246000fd5b95607f1695620000ae565b600080fd5b6040519190601f01601f191682016001600160401b03811183821017620002e657604052565b919080601f84011215620003e65782516001600160401b038111620002e65760209062000447601f8201601f19168301620003eb565b92818452828287010111620003e65760005b8181106200046f57508260009394955001015290565b85810183015184820184015282016200045956fe608060408181526004918236101561001657600080fd5b600092833560e01c91826306fdde03146105d057508163095ea7b3146105a657816318160ddd1461058757816323b872dd14610477578163313ce5671461045b57816339509351146103cf57816370a082311461038c57816395d89b411461020e578163a457c2d71461012657508063a9059cbb146100f65763dd62ed3e1461009e57600080fd5b346100f257806003193601126100f257806020926100ba610700565b6100c2610728565b73ffffffffffffffffffffffffffffffffffffffff91821683526001865283832091168252845220549051908152f35b5080fd5b50346100f257806003193601126100f25760209061011f610115610700565b602435903361074b565b5160018152f35b9050823461020b578260031936011261020b57610141610700565b918360243592338152600160205281812073ffffffffffffffffffffffffffffffffffffffff861682526020522054908282106101885760208561011f858503873361095a565b60849060208651917f08c379a0000000000000000000000000000000000000000000000000000000008352820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152fd5b80fd5b8383346100f257816003193601126100f257805190828454600181811c90808316928315610382575b602093848410811461035657838852879594939291811561031957506001146102bb575b50505003601f01601f191682019267ffffffffffffffff84118385101761028f575082918261028b9252826106b8565b0390f35b806041867f4e487b71000000000000000000000000000000000000000000000000000000006024945252fd5b8888529193925086917f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5b8284106103035750505090601f1992601f9282010191819361025b565b80548885018701528794509285019281016102e6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016848701525050151560051b830101905081601f601f1961025b565b60248960228c7f4e487b7100000000000000000000000000000000000000000000000000000000835252fd5b91607f1691610237565b5050346100f25760206003193601126100f2578060209273ffffffffffffffffffffffffffffffffffffffff6103c0610700565b16815280845220549051908152f35b82843461020b578160031936011261020b576103e9610700565b338252600160205282822073ffffffffffffffffffffffffffffffffffffffff821683526020528282205491602435830180931161042f5760208461011f85853361095a565b806011867f4e487b71000000000000000000000000000000000000000000000000000000006024945252fd5b5050346100f257816003193601126100f2576020905160128152f35b839150346100f25760606003193601126100f257610493610700565b61049b610728565b91846044359473ffffffffffffffffffffffffffffffffffffffff8416815260016020528181203382526020522054907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610501575b60208661011f87878761074b565b84821061052a575091839161051f6020969561011f9503338361095a565b9193948193506104f3565b60649060208751917f08c379a0000000000000000000000000000000000000000000000000000000008352820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152fd5b5050346100f257816003193601126100f2576020906002549051908152f35b5050346100f257806003193601126100f25760209061011f6105c6610700565b602435903361095a565b849084346106b457826003193601126106b45782600354600181811c908083169283156106aa575b6020938484108114610356578388528795949392918115610319575060011461064b5750505003601f01601f191682019267ffffffffffffffff84118385101761028f575082918261028b9252826106b8565b600388529193925086917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b8284106106945750505090601f1992601f9282010191819361025b565b8054888501870152879450928501928101610677565b91607f16916105f8565b8280fd5b60208082528251818301819052939260005b8581106106ec57505050601f19601f8460006040809697860101520116010190565b8181018301518482016040015282016106ca565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361072357565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361072357565b73ffffffffffffffffffffffffffffffffffffffff8091169182156108d65716918215610852576000828152806020526040812054918083106107ce57604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef958760209652828652038282205586815220818154019055604051908152a3565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152fd5b73ffffffffffffffffffffffffffffffffffffffff809116918215610a4c57169182156109c85760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925918360005260018252604060002085600052825280604060002055604051908152a3565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152fdfea164736f6c6343000813000a",
}

// Mockerc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Mockerc20MetaData.ABI instead.
var Mockerc20ABI = Mockerc20MetaData.ABI

// Mockerc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Mockerc20MetaData.Bin instead.
var Mockerc20Bin = Mockerc20MetaData.Bin

// DeployMockerc20 deploys a new Ethereum contract, binding an instance of Mockerc20 to it.
func DeployMockerc20(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, initialAccount common.Address, initialBalance *big.Int) (common.Address, *types.Transaction, *Mockerc20, error) {
	parsed, err := Mockerc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Mockerc20Bin), backend, name, symbol, initialAccount, initialBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mockerc20{Mockerc20Caller: Mockerc20Caller{contract: contract}, Mockerc20Transactor: Mockerc20Transactor{contract: contract}, Mockerc20Filterer: Mockerc20Filterer{contract: contract}}, nil
}

// Mockerc20 is an auto generated Go binding around an Ethereum contract.
type Mockerc20 struct {
	Mockerc20Caller     // Read-only binding to the contract
	Mockerc20Transactor // Write-only binding to the contract
	Mockerc20Filterer   // Log filterer for contract events
}

// Mockerc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Mockerc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mockerc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Mockerc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mockerc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Mockerc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mockerc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Mockerc20Session struct {
	Contract     *Mockerc20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Mockerc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Mockerc20CallerSession struct {
	Contract *Mockerc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Mockerc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Mockerc20TransactorSession struct {
	Contract     *Mockerc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Mockerc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Mockerc20Raw struct {
	Contract *Mockerc20 // Generic contract binding to access the raw methods on
}

// Mockerc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Mockerc20CallerRaw struct {
	Contract *Mockerc20Caller // Generic read-only contract binding to access the raw methods on
}

// Mockerc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Mockerc20TransactorRaw struct {
	Contract *Mockerc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMockerc20 creates a new instance of Mockerc20, bound to a specific deployed contract.
func NewMockerc20(address common.Address, backend bind.ContractBackend) (*Mockerc20, error) {
	contract, err := bindMockerc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mockerc20{Mockerc20Caller: Mockerc20Caller{contract: contract}, Mockerc20Transactor: Mockerc20Transactor{contract: contract}, Mockerc20Filterer: Mockerc20Filterer{contract: contract}}, nil
}

// NewMockerc20Caller creates a new read-only instance of Mockerc20, bound to a specific deployed contract.
func NewMockerc20Caller(address common.Address, caller bind.ContractCaller) (*Mockerc20Caller, error) {
	contract, err := bindMockerc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Mockerc20Caller{contract: contract}, nil
}

// NewMockerc20Transactor creates a new write-only instance of Mockerc20, bound to a specific deployed contract.
func NewMockerc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Mockerc20Transactor, error) {
	contract, err := bindMockerc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Mockerc20Transactor{contract: contract}, nil
}

// NewMockerc20Filterer creates a new log filterer instance of Mockerc20, bound to a specific deployed contract.
func NewMockerc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Mockerc20Filterer, error) {
	contract, err := bindMockerc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Mockerc20Filterer{contract: contract}, nil
}

// bindMockerc20 binds a generic wrapper to an already deployed contract.
func bindMockerc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Mockerc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mockerc20 *Mockerc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mockerc20.Contract.Mockerc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mockerc20 *Mockerc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mockerc20.Contract.Mockerc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mockerc20 *Mockerc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mockerc20.Contract.Mockerc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mockerc20 *Mockerc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mockerc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mockerc20 *Mockerc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mockerc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mockerc20 *Mockerc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mockerc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mockerc20 *Mockerc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mockerc20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mockerc20 *Mockerc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mockerc20.Contract.Allowance(&_Mockerc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mockerc20 *Mockerc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mockerc20.Contract.Allowance(&_Mockerc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mockerc20 *Mockerc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mockerc20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mockerc20 *Mockerc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mockerc20.Contract.BalanceOf(&_Mockerc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mockerc20 *Mockerc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mockerc20.Contract.BalanceOf(&_Mockerc20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mockerc20 *Mockerc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Mockerc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mockerc20 *Mockerc20Session) Decimals() (uint8, error) {
	return _Mockerc20.Contract.Decimals(&_Mockerc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mockerc20 *Mockerc20CallerSession) Decimals() (uint8, error) {
	return _Mockerc20.Contract.Decimals(&_Mockerc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mockerc20 *Mockerc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mockerc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mockerc20 *Mockerc20Session) Name() (string, error) {
	return _Mockerc20.Contract.Name(&_Mockerc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mockerc20 *Mockerc20CallerSession) Name() (string, error) {
	return _Mockerc20.Contract.Name(&_Mockerc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mockerc20 *Mockerc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mockerc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mockerc20 *Mockerc20Session) Symbol() (string, error) {
	return _Mockerc20.Contract.Symbol(&_Mockerc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mockerc20 *Mockerc20CallerSession) Symbol() (string, error) {
	return _Mockerc20.Contract.Symbol(&_Mockerc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mockerc20 *Mockerc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mockerc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mockerc20 *Mockerc20Session) TotalSupply() (*big.Int, error) {
	return _Mockerc20.Contract.TotalSupply(&_Mockerc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mockerc20 *Mockerc20CallerSession) TotalSupply() (*big.Int, error) {
	return _Mockerc20.Contract.TotalSupply(&_Mockerc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mockerc20 *Mockerc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mockerc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mockerc20 *Mockerc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mockerc20.Contract.Approve(&_Mockerc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mockerc20 *Mockerc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mockerc20.Contract.Approve(&_Mockerc20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mockerc20 *Mockerc20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mockerc20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mockerc20 *Mockerc20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mockerc20.Contract.DecreaseAllowance(&_Mockerc20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mockerc20 *Mockerc20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mockerc20.Contract.DecreaseAllowance(&_Mockerc20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mockerc20 *Mockerc20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mockerc20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mockerc20 *Mockerc20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mockerc20.Contract.IncreaseAllowance(&_Mockerc20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mockerc20 *Mockerc20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mockerc20.Contract.IncreaseAllowance(&_Mockerc20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Mockerc20 *Mockerc20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mockerc20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Mockerc20 *Mockerc20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mockerc20.Contract.Transfer(&_Mockerc20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Mockerc20 *Mockerc20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mockerc20.Contract.Transfer(&_Mockerc20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Mockerc20 *Mockerc20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mockerc20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Mockerc20 *Mockerc20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mockerc20.Contract.TransferFrom(&_Mockerc20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Mockerc20 *Mockerc20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mockerc20.Contract.TransferFrom(&_Mockerc20.TransactOpts, from, to, amount)
}

// Mockerc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Mockerc20 contract.
type Mockerc20ApprovalIterator struct {
	Event *Mockerc20Approval // Event containing the contract specifics and raw log

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
func (it *Mockerc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Mockerc20Approval)
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
		it.Event = new(Mockerc20Approval)
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
func (it *Mockerc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Mockerc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Mockerc20Approval represents a Approval event raised by the Mockerc20 contract.
type Mockerc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mockerc20 *Mockerc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Mockerc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mockerc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Mockerc20ApprovalIterator{contract: _Mockerc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mockerc20 *Mockerc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Mockerc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mockerc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Mockerc20Approval)
				if err := _Mockerc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Mockerc20 *Mockerc20Filterer) ParseApproval(log types.Log) (*Mockerc20Approval, error) {
	event := new(Mockerc20Approval)
	if err := _Mockerc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Mockerc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Mockerc20 contract.
type Mockerc20TransferIterator struct {
	Event *Mockerc20Transfer // Event containing the contract specifics and raw log

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
func (it *Mockerc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Mockerc20Transfer)
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
		it.Event = new(Mockerc20Transfer)
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
func (it *Mockerc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Mockerc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Mockerc20Transfer represents a Transfer event raised by the Mockerc20 contract.
type Mockerc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mockerc20 *Mockerc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Mockerc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mockerc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Mockerc20TransferIterator{contract: _Mockerc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mockerc20 *Mockerc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Mockerc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mockerc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Mockerc20Transfer)
				if err := _Mockerc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Mockerc20 *Mockerc20Filterer) ParseTransfer(log types.Log) (*Mockerc20Transfer, error) {
	event := new(Mockerc20Transfer)
	if err := _Mockerc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
