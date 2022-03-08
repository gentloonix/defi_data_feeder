// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PairReader

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

// PairReaderMetaData contains all meta data concerning the PairReader contract.
var PairReaderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PairReaderABI is the input ABI used to generate the binding from.
// Deprecated: Use PairReaderMetaData.ABI instead.
var PairReaderABI = PairReaderMetaData.ABI

// PairReader is an auto generated Go binding around an Ethereum contract.
type PairReader struct {
	PairReaderCaller     // Read-only binding to the contract
	PairReaderTransactor // Write-only binding to the contract
	PairReaderFilterer   // Log filterer for contract events
}

// PairReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairReaderSession struct {
	Contract     *PairReader       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairReaderCallerSession struct {
	Contract *PairReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PairReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairReaderTransactorSession struct {
	Contract     *PairReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PairReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairReaderRaw struct {
	Contract *PairReader // Generic contract binding to access the raw methods on
}

// PairReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairReaderCallerRaw struct {
	Contract *PairReaderCaller // Generic read-only contract binding to access the raw methods on
}

// PairReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairReaderTransactorRaw struct {
	Contract *PairReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairReader creates a new instance of PairReader, bound to a specific deployed contract.
func NewPairReader(address common.Address, backend bind.ContractBackend) (*PairReader, error) {
	contract, err := bindPairReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PairReader{PairReaderCaller: PairReaderCaller{contract: contract}, PairReaderTransactor: PairReaderTransactor{contract: contract}, PairReaderFilterer: PairReaderFilterer{contract: contract}}, nil
}

// NewPairReaderCaller creates a new read-only instance of PairReader, bound to a specific deployed contract.
func NewPairReaderCaller(address common.Address, caller bind.ContractCaller) (*PairReaderCaller, error) {
	contract, err := bindPairReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairReaderCaller{contract: contract}, nil
}

// NewPairReaderTransactor creates a new write-only instance of PairReader, bound to a specific deployed contract.
func NewPairReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*PairReaderTransactor, error) {
	contract, err := bindPairReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairReaderTransactor{contract: contract}, nil
}

// NewPairReaderFilterer creates a new log filterer instance of PairReader, bound to a specific deployed contract.
func NewPairReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*PairReaderFilterer, error) {
	contract, err := bindPairReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairReaderFilterer{contract: contract}, nil
}

// bindPairReader binds a generic wrapper to an already deployed contract.
func bindPairReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairReaderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairReader *PairReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PairReader.Contract.PairReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairReader *PairReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairReader.Contract.PairReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairReader *PairReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairReader.Contract.PairReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairReader *PairReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PairReader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairReader *PairReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairReader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairReader *PairReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairReader.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PairReader *PairReaderCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairReader.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PairReader *PairReaderSession) Factory() (common.Address, error) {
	return _PairReader.Contract.Factory(&_PairReader.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PairReader *PairReaderCallerSession) Factory() (common.Address, error) {
	return _PairReader.Contract.Factory(&_PairReader.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_PairReader *PairReaderCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _PairReader.contract.Call(opts, &out, "getReserves")

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
func (_PairReader *PairReaderSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _PairReader.Contract.GetReserves(&_PairReader.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_PairReader *PairReaderCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _PairReader.Contract.GetReserves(&_PairReader.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_PairReader *PairReaderCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairReader.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_PairReader *PairReaderSession) KLast() (*big.Int, error) {
	return _PairReader.Contract.KLast(&_PairReader.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_PairReader *PairReaderCallerSession) KLast() (*big.Int, error) {
	return _PairReader.Contract.KLast(&_PairReader.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_PairReader *PairReaderCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairReader.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_PairReader *PairReaderSession) Price0CumulativeLast() (*big.Int, error) {
	return _PairReader.Contract.Price0CumulativeLast(&_PairReader.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_PairReader *PairReaderCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _PairReader.Contract.Price0CumulativeLast(&_PairReader.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_PairReader *PairReaderCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PairReader.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_PairReader *PairReaderSession) Price1CumulativeLast() (*big.Int, error) {
	return _PairReader.Contract.Price1CumulativeLast(&_PairReader.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_PairReader *PairReaderCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _PairReader.Contract.Price1CumulativeLast(&_PairReader.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_PairReader *PairReaderCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairReader.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_PairReader *PairReaderSession) Token0() (common.Address, error) {
	return _PairReader.Contract.Token0(&_PairReader.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_PairReader *PairReaderCallerSession) Token0() (common.Address, error) {
	return _PairReader.Contract.Token0(&_PairReader.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_PairReader *PairReaderCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PairReader.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_PairReader *PairReaderSession) Token1() (common.Address, error) {
	return _PairReader.Contract.Token1(&_PairReader.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_PairReader *PairReaderCallerSession) Token1() (common.Address, error) {
	return _PairReader.Contract.Token1(&_PairReader.CallOpts)
}
