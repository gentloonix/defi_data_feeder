// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PairBinder

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
)

// PairBinderMetaData contains all meta data concerning the PairBinder contract.
var PairBinderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PairBinderABI is the input ABI used to generate the binding from.
// Deprecated: Use PairBinderMetaData.ABI instead.
var PairBinderABI = PairBinderMetaData.ABI

// PairBinder is an auto generated Go binding around an Ethereum contract.
type PairBinder struct {
	PairBinderCaller     // Read-only binding to the contract
	PairBinderTransactor // Write-only binding to the contract
	PairBinderFilterer   // Log filterer for contract events
}

// PairBinderCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairBinderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairBinderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairBinderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairBinderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairBinderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairBinderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairBinderSession struct {
	Contract     *PairBinder       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairBinderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairBinderCallerSession struct {
	Contract *PairBinderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PairBinderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairBinderTransactorSession struct {
	Contract     *PairBinderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PairBinderRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairBinderRaw struct {
	Contract *PairBinder // Generic contract binding to access the raw methods on
}

// PairBinderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairBinderCallerRaw struct {
	Contract *PairBinderCaller // Generic read-only contract binding to access the raw methods on
}

// PairBinderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairBinderTransactorRaw struct {
	Contract *PairBinderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairBinder creates a new instance of PairBinder, bound to a specific deployed contract.
func NewPairBinder(address common.Address, backend bind.ContractBackend) (*PairBinder, error) {
	contract, err := bindPairBinder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PairBinder{PairBinderCaller: PairBinderCaller{contract: contract}, PairBinderTransactor: PairBinderTransactor{contract: contract}, PairBinderFilterer: PairBinderFilterer{contract: contract}}, nil
}

// NewPairBinderCaller creates a new read-only instance of PairBinder, bound to a specific deployed contract.
func NewPairBinderCaller(address common.Address, caller bind.ContractCaller) (*PairBinderCaller, error) {
	contract, err := bindPairBinder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairBinderCaller{contract: contract}, nil
}

// NewPairBinderTransactor creates a new write-only instance of PairBinder, bound to a specific deployed contract.
func NewPairBinderTransactor(address common.Address, transactor bind.ContractTransactor) (*PairBinderTransactor, error) {
	contract, err := bindPairBinder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairBinderTransactor{contract: contract}, nil
}

// NewPairBinderFilterer creates a new log filterer instance of PairBinder, bound to a specific deployed contract.
func NewPairBinderFilterer(address common.Address, filterer bind.ContractFilterer) (*PairBinderFilterer, error) {
	contract, err := bindPairBinder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairBinderFilterer{contract: contract}, nil
}

// bindPairBinder binds a generic wrapper to an already deployed contract.
func bindPairBinder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairBinderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairBinder *PairBinderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PairBinder.Contract.PairBinderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairBinder *PairBinderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairBinder.Contract.PairBinderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairBinder *PairBinderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairBinder.Contract.PairBinderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairBinder *PairBinderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PairBinder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairBinder *PairBinderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairBinder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairBinder *PairBinderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairBinder.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PairBinder *PairBinderCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairBinder.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PairBinder *PairBinderSession) Factory() (common.Address, error) {
	return _PairBinder.Contract.Factory(&_PairBinder.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PairBinder *PairBinderCallerSession) Factory() (common.Address, error) {
	return _PairBinder.Contract.Factory(&_PairBinder.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_PairBinder *PairBinderCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _PairBinder.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_PairBinder *PairBinderSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _PairBinder.Contract.GetReserves(&_PairBinder.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_PairBinder *PairBinderCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _PairBinder.Contract.GetReserves(&_PairBinder.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_PairBinder *PairBinderCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairBinder.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_PairBinder *PairBinderSession) KLast() (*big.Int, error) {
	return _PairBinder.Contract.KLast(&_PairBinder.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_PairBinder *PairBinderCallerSession) KLast() (*big.Int, error) {
	return _PairBinder.Contract.KLast(&_PairBinder.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_PairBinder *PairBinderCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairBinder.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_PairBinder *PairBinderSession) Price0CumulativeLast() (*big.Int, error) {
	return _PairBinder.Contract.Price0CumulativeLast(&_PairBinder.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_PairBinder *PairBinderCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _PairBinder.Contract.Price0CumulativeLast(&_PairBinder.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_PairBinder *PairBinderCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairBinder.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_PairBinder *PairBinderSession) Price1CumulativeLast() (*big.Int, error) {
	return _PairBinder.Contract.Price1CumulativeLast(&_PairBinder.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_PairBinder *PairBinderCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _PairBinder.Contract.Price1CumulativeLast(&_PairBinder.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_PairBinder *PairBinderCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairBinder.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_PairBinder *PairBinderSession) Token0() (common.Address, error) {
	return _PairBinder.Contract.Token0(&_PairBinder.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_PairBinder *PairBinderCallerSession) Token0() (common.Address, error) {
	return _PairBinder.Contract.Token0(&_PairBinder.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_PairBinder *PairBinderCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairBinder.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_PairBinder *PairBinderSession) Token1() (common.Address, error) {
	return _PairBinder.Contract.Token1(&_PairBinder.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_PairBinder *PairBinderCallerSession) Token1() (common.Address, error) {
	return _PairBinder.Contract.Token1(&_PairBinder.CallOpts)
}
