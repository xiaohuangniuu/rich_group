// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mycontract

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

// MycontractMetaData contains all meta data concerning the Mycontract contract.
var MycontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_PUBLIC_MINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_PER_TOKEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVENANCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAllowListActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numberOfTokens\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"numberOfTokens\",\"type\":\"uint8\"}],\"name\":\"mintAllowList\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"numAvailableToMint\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"reserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleIsActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"numAllowedToMint\",\"type\":\"uint8\"}],\"name\":\"setAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isAllowListActive\",\"type\":\"bool\"}],\"name\":\"setIsAllowListActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"provenance\",\"type\":\"string\"}],\"name\":\"setProvenance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"newState\",\"type\":\"bool\"}],\"name\":\"setSaleState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MycontractABI is the input ABI used to generate the binding from.
// Deprecated: Use MycontractMetaData.ABI instead.
var MycontractABI = MycontractMetaData.ABI

// Mycontract is an auto generated Go binding around an Ethereum contract.
type Mycontract struct {
	MycontractCaller     // Read-only binding to the contract
	MycontractTransactor // Write-only binding to the contract
	MycontractFilterer   // Log filterer for contract events
}

// MycontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MycontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MycontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MycontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MycontractSession struct {
	Contract     *Mycontract       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MycontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MycontractCallerSession struct {
	Contract *MycontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MycontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MycontractTransactorSession struct {
	Contract     *MycontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MycontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MycontractRaw struct {
	Contract *Mycontract // Generic contract binding to access the raw methods on
}

// MycontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MycontractCallerRaw struct {
	Contract *MycontractCaller // Generic read-only contract binding to access the raw methods on
}

// MycontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MycontractTransactorRaw struct {
	Contract *MycontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMycontract creates a new instance of Mycontract, bound to a specific deployed contract.
func NewMycontract(address common.Address, backend bind.ContractBackend) (*Mycontract, error) {
	contract, err := bindMycontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mycontract{MycontractCaller: MycontractCaller{contract: contract}, MycontractTransactor: MycontractTransactor{contract: contract}, MycontractFilterer: MycontractFilterer{contract: contract}}, nil
}

// NewMycontractCaller creates a new read-only instance of Mycontract, bound to a specific deployed contract.
func NewMycontractCaller(address common.Address, caller bind.ContractCaller) (*MycontractCaller, error) {
	contract, err := bindMycontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MycontractCaller{contract: contract}, nil
}

// NewMycontractTransactor creates a new write-only instance of Mycontract, bound to a specific deployed contract.
func NewMycontractTransactor(address common.Address, transactor bind.ContractTransactor) (*MycontractTransactor, error) {
	contract, err := bindMycontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MycontractTransactor{contract: contract}, nil
}

// NewMycontractFilterer creates a new log filterer instance of Mycontract, bound to a specific deployed contract.
func NewMycontractFilterer(address common.Address, filterer bind.ContractFilterer) (*MycontractFilterer, error) {
	contract, err := bindMycontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MycontractFilterer{contract: contract}, nil
}

// bindMycontract binds a generic wrapper to an already deployed contract.
func bindMycontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MycontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mycontract *MycontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mycontract.Contract.MycontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mycontract *MycontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.Contract.MycontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mycontract *MycontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mycontract.Contract.MycontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mycontract *MycontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mycontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mycontract *MycontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mycontract *MycontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mycontract.Contract.contract.Transact(opts, method, params...)
}

// MAXPUBLICMINT is a free data retrieval call binding the contract method 0x65f13097.
//
// Solidity: function MAX_PUBLIC_MINT() view returns(uint256)
func (_Mycontract *MycontractCaller) MAXPUBLICMINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "MAX_PUBLIC_MINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPUBLICMINT is a free data retrieval call binding the contract method 0x65f13097.
//
// Solidity: function MAX_PUBLIC_MINT() view returns(uint256)
func (_Mycontract *MycontractSession) MAXPUBLICMINT() (*big.Int, error) {
	return _Mycontract.Contract.MAXPUBLICMINT(&_Mycontract.CallOpts)
}

// MAXPUBLICMINT is a free data retrieval call binding the contract method 0x65f13097.
//
// Solidity: function MAX_PUBLIC_MINT() view returns(uint256)
func (_Mycontract *MycontractCallerSession) MAXPUBLICMINT() (*big.Int, error) {
	return _Mycontract.Contract.MAXPUBLICMINT(&_Mycontract.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Mycontract *MycontractCaller) MAXSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "MAX_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Mycontract *MycontractSession) MAXSUPPLY() (*big.Int, error) {
	return _Mycontract.Contract.MAXSUPPLY(&_Mycontract.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Mycontract *MycontractCallerSession) MAXSUPPLY() (*big.Int, error) {
	return _Mycontract.Contract.MAXSUPPLY(&_Mycontract.CallOpts)
}

// PRICEPERTOKEN is a free data retrieval call binding the contract method 0x833b9499.
//
// Solidity: function PRICE_PER_TOKEN() view returns(uint256)
func (_Mycontract *MycontractCaller) PRICEPERTOKEN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "PRICE_PER_TOKEN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPERTOKEN is a free data retrieval call binding the contract method 0x833b9499.
//
// Solidity: function PRICE_PER_TOKEN() view returns(uint256)
func (_Mycontract *MycontractSession) PRICEPERTOKEN() (*big.Int, error) {
	return _Mycontract.Contract.PRICEPERTOKEN(&_Mycontract.CallOpts)
}

// PRICEPERTOKEN is a free data retrieval call binding the contract method 0x833b9499.
//
// Solidity: function PRICE_PER_TOKEN() view returns(uint256)
func (_Mycontract *MycontractCallerSession) PRICEPERTOKEN() (*big.Int, error) {
	return _Mycontract.Contract.PRICEPERTOKEN(&_Mycontract.CallOpts)
}

// PROVENANCE is a free data retrieval call binding the contract method 0x6373a6b1.
//
// Solidity: function PROVENANCE() view returns(string)
func (_Mycontract *MycontractCaller) PROVENANCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "PROVENANCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PROVENANCE is a free data retrieval call binding the contract method 0x6373a6b1.
//
// Solidity: function PROVENANCE() view returns(string)
func (_Mycontract *MycontractSession) PROVENANCE() (string, error) {
	return _Mycontract.Contract.PROVENANCE(&_Mycontract.CallOpts)
}

// PROVENANCE is a free data retrieval call binding the contract method 0x6373a6b1.
//
// Solidity: function PROVENANCE() view returns(string)
func (_Mycontract *MycontractCallerSession) PROVENANCE() (string, error) {
	return _Mycontract.Contract.PROVENANCE(&_Mycontract.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Mycontract *MycontractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Mycontract *MycontractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Mycontract.Contract.BalanceOf(&_Mycontract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Mycontract *MycontractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Mycontract.Contract.BalanceOf(&_Mycontract.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Mycontract *MycontractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Mycontract *MycontractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Mycontract.Contract.GetApproved(&_Mycontract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Mycontract *MycontractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Mycontract.Contract.GetApproved(&_Mycontract.CallOpts, tokenId)
}

// IsAllowListActive is a free data retrieval call binding the contract method 0x29fc6bae.
//
// Solidity: function isAllowListActive() view returns(bool)
func (_Mycontract *MycontractCaller) IsAllowListActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "isAllowListActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowListActive is a free data retrieval call binding the contract method 0x29fc6bae.
//
// Solidity: function isAllowListActive() view returns(bool)
func (_Mycontract *MycontractSession) IsAllowListActive() (bool, error) {
	return _Mycontract.Contract.IsAllowListActive(&_Mycontract.CallOpts)
}

// IsAllowListActive is a free data retrieval call binding the contract method 0x29fc6bae.
//
// Solidity: function isAllowListActive() view returns(bool)
func (_Mycontract *MycontractCallerSession) IsAllowListActive() (bool, error) {
	return _Mycontract.Contract.IsAllowListActive(&_Mycontract.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Mycontract *MycontractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Mycontract *MycontractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Mycontract.Contract.IsApprovedForAll(&_Mycontract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Mycontract *MycontractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Mycontract.Contract.IsApprovedForAll(&_Mycontract.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mycontract *MycontractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mycontract *MycontractSession) Name() (string, error) {
	return _Mycontract.Contract.Name(&_Mycontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mycontract *MycontractCallerSession) Name() (string, error) {
	return _Mycontract.Contract.Name(&_Mycontract.CallOpts)
}

// NumAvailableToMint is a free data retrieval call binding the contract method 0xc04a2836.
//
// Solidity: function numAvailableToMint(address addr) view returns(uint8)
func (_Mycontract *MycontractCaller) NumAvailableToMint(opts *bind.CallOpts, addr common.Address) (uint8, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "numAvailableToMint", addr)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NumAvailableToMint is a free data retrieval call binding the contract method 0xc04a2836.
//
// Solidity: function numAvailableToMint(address addr) view returns(uint8)
func (_Mycontract *MycontractSession) NumAvailableToMint(addr common.Address) (uint8, error) {
	return _Mycontract.Contract.NumAvailableToMint(&_Mycontract.CallOpts, addr)
}

// NumAvailableToMint is a free data retrieval call binding the contract method 0xc04a2836.
//
// Solidity: function numAvailableToMint(address addr) view returns(uint8)
func (_Mycontract *MycontractCallerSession) NumAvailableToMint(addr common.Address) (uint8, error) {
	return _Mycontract.Contract.NumAvailableToMint(&_Mycontract.CallOpts, addr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mycontract *MycontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mycontract *MycontractSession) Owner() (common.Address, error) {
	return _Mycontract.Contract.Owner(&_Mycontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mycontract *MycontractCallerSession) Owner() (common.Address, error) {
	return _Mycontract.Contract.Owner(&_Mycontract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Mycontract *MycontractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Mycontract *MycontractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Mycontract.Contract.OwnerOf(&_Mycontract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Mycontract *MycontractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Mycontract.Contract.OwnerOf(&_Mycontract.CallOpts, tokenId)
}

// SaleIsActive is a free data retrieval call binding the contract method 0xeb8d2444.
//
// Solidity: function saleIsActive() view returns(bool)
func (_Mycontract *MycontractCaller) SaleIsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "saleIsActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SaleIsActive is a free data retrieval call binding the contract method 0xeb8d2444.
//
// Solidity: function saleIsActive() view returns(bool)
func (_Mycontract *MycontractSession) SaleIsActive() (bool, error) {
	return _Mycontract.Contract.SaleIsActive(&_Mycontract.CallOpts)
}

// SaleIsActive is a free data retrieval call binding the contract method 0xeb8d2444.
//
// Solidity: function saleIsActive() view returns(bool)
func (_Mycontract *MycontractCallerSession) SaleIsActive() (bool, error) {
	return _Mycontract.Contract.SaleIsActive(&_Mycontract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mycontract *MycontractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mycontract *MycontractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Mycontract.Contract.SupportsInterface(&_Mycontract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mycontract *MycontractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Mycontract.Contract.SupportsInterface(&_Mycontract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mycontract *MycontractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mycontract *MycontractSession) Symbol() (string, error) {
	return _Mycontract.Contract.Symbol(&_Mycontract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mycontract *MycontractCallerSession) Symbol() (string, error) {
	return _Mycontract.Contract.Symbol(&_Mycontract.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Mycontract *MycontractCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Mycontract *MycontractSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Mycontract.Contract.TokenByIndex(&_Mycontract.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Mycontract *MycontractCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Mycontract.Contract.TokenByIndex(&_Mycontract.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Mycontract *MycontractCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Mycontract *MycontractSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Mycontract.Contract.TokenOfOwnerByIndex(&_Mycontract.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Mycontract *MycontractCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Mycontract.Contract.TokenOfOwnerByIndex(&_Mycontract.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Mycontract *MycontractCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Mycontract *MycontractSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Mycontract.Contract.TokenURI(&_Mycontract.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Mycontract *MycontractCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Mycontract.Contract.TokenURI(&_Mycontract.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mycontract *MycontractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mycontract *MycontractSession) TotalSupply() (*big.Int, error) {
	return _Mycontract.Contract.TotalSupply(&_Mycontract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mycontract *MycontractCallerSession) TotalSupply() (*big.Int, error) {
	return _Mycontract.Contract.TotalSupply(&_Mycontract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Mycontract *MycontractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Mycontract *MycontractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.Approve(&_Mycontract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Mycontract *MycontractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.Approve(&_Mycontract.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 numberOfTokens) payable returns()
func (_Mycontract *MycontractTransactor) Mint(opts *bind.TransactOpts, numberOfTokens *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "mint", numberOfTokens)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 numberOfTokens) payable returns()
func (_Mycontract *MycontractSession) Mint(numberOfTokens *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.Mint(&_Mycontract.TransactOpts, numberOfTokens)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 numberOfTokens) payable returns()
func (_Mycontract *MycontractTransactorSession) Mint(numberOfTokens *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.Mint(&_Mycontract.TransactOpts, numberOfTokens)
}

// MintAllowList is a paid mutator transaction binding the contract method 0xddff5b1c.
//
// Solidity: function mintAllowList(uint8 numberOfTokens) payable returns()
func (_Mycontract *MycontractTransactor) MintAllowList(opts *bind.TransactOpts, numberOfTokens uint8) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "mintAllowList", numberOfTokens)
}

// MintAllowList is a paid mutator transaction binding the contract method 0xddff5b1c.
//
// Solidity: function mintAllowList(uint8 numberOfTokens) payable returns()
func (_Mycontract *MycontractSession) MintAllowList(numberOfTokens uint8) (*types.Transaction, error) {
	return _Mycontract.Contract.MintAllowList(&_Mycontract.TransactOpts, numberOfTokens)
}

// MintAllowList is a paid mutator transaction binding the contract method 0xddff5b1c.
//
// Solidity: function mintAllowList(uint8 numberOfTokens) payable returns()
func (_Mycontract *MycontractTransactorSession) MintAllowList(numberOfTokens uint8) (*types.Transaction, error) {
	return _Mycontract.Contract.MintAllowList(&_Mycontract.TransactOpts, numberOfTokens)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mycontract *MycontractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mycontract *MycontractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mycontract.Contract.RenounceOwnership(&_Mycontract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mycontract *MycontractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mycontract.Contract.RenounceOwnership(&_Mycontract.TransactOpts)
}

// Reserve is a paid mutator transaction binding the contract method 0x819b25ba.
//
// Solidity: function reserve(uint256 n) returns()
func (_Mycontract *MycontractTransactor) Reserve(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "reserve", n)
}

// Reserve is a paid mutator transaction binding the contract method 0x819b25ba.
//
// Solidity: function reserve(uint256 n) returns()
func (_Mycontract *MycontractSession) Reserve(n *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.Reserve(&_Mycontract.TransactOpts, n)
}

// Reserve is a paid mutator transaction binding the contract method 0x819b25ba.
//
// Solidity: function reserve(uint256 n) returns()
func (_Mycontract *MycontractTransactorSession) Reserve(n *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.Reserve(&_Mycontract.TransactOpts, n)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Mycontract *MycontractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Mycontract *MycontractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SafeTransferFrom(&_Mycontract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Mycontract *MycontractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SafeTransferFrom(&_Mycontract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Mycontract *MycontractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Mycontract *MycontractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Mycontract.Contract.SafeTransferFrom0(&_Mycontract.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Mycontract *MycontractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Mycontract.Contract.SafeTransferFrom0(&_Mycontract.TransactOpts, from, to, tokenId, _data)
}

// SetAllowList is a paid mutator transaction binding the contract method 0x8295784d.
//
// Solidity: function setAllowList(address[] addresses, uint8 numAllowedToMint) returns()
func (_Mycontract *MycontractTransactor) SetAllowList(opts *bind.TransactOpts, addresses []common.Address, numAllowedToMint uint8) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "setAllowList", addresses, numAllowedToMint)
}

// SetAllowList is a paid mutator transaction binding the contract method 0x8295784d.
//
// Solidity: function setAllowList(address[] addresses, uint8 numAllowedToMint) returns()
func (_Mycontract *MycontractSession) SetAllowList(addresses []common.Address, numAllowedToMint uint8) (*types.Transaction, error) {
	return _Mycontract.Contract.SetAllowList(&_Mycontract.TransactOpts, addresses, numAllowedToMint)
}

// SetAllowList is a paid mutator transaction binding the contract method 0x8295784d.
//
// Solidity: function setAllowList(address[] addresses, uint8 numAllowedToMint) returns()
func (_Mycontract *MycontractTransactorSession) SetAllowList(addresses []common.Address, numAllowedToMint uint8) (*types.Transaction, error) {
	return _Mycontract.Contract.SetAllowList(&_Mycontract.TransactOpts, addresses, numAllowedToMint)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mycontract *MycontractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mycontract *MycontractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mycontract.Contract.SetApprovalForAll(&_Mycontract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mycontract *MycontractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mycontract.Contract.SetApprovalForAll(&_Mycontract.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Mycontract *MycontractTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI_ string) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "setBaseURI", baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Mycontract *MycontractSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _Mycontract.Contract.SetBaseURI(&_Mycontract.TransactOpts, baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Mycontract *MycontractTransactorSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _Mycontract.Contract.SetBaseURI(&_Mycontract.TransactOpts, baseURI_)
}

// SetIsAllowListActive is a paid mutator transaction binding the contract method 0x718bc4af.
//
// Solidity: function setIsAllowListActive(bool _isAllowListActive) returns()
func (_Mycontract *MycontractTransactor) SetIsAllowListActive(opts *bind.TransactOpts, _isAllowListActive bool) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "setIsAllowListActive", _isAllowListActive)
}

// SetIsAllowListActive is a paid mutator transaction binding the contract method 0x718bc4af.
//
// Solidity: function setIsAllowListActive(bool _isAllowListActive) returns()
func (_Mycontract *MycontractSession) SetIsAllowListActive(_isAllowListActive bool) (*types.Transaction, error) {
	return _Mycontract.Contract.SetIsAllowListActive(&_Mycontract.TransactOpts, _isAllowListActive)
}

// SetIsAllowListActive is a paid mutator transaction binding the contract method 0x718bc4af.
//
// Solidity: function setIsAllowListActive(bool _isAllowListActive) returns()
func (_Mycontract *MycontractTransactorSession) SetIsAllowListActive(_isAllowListActive bool) (*types.Transaction, error) {
	return _Mycontract.Contract.SetIsAllowListActive(&_Mycontract.TransactOpts, _isAllowListActive)
}

// SetProvenance is a paid mutator transaction binding the contract method 0xffe630b5.
//
// Solidity: function setProvenance(string provenance) returns()
func (_Mycontract *MycontractTransactor) SetProvenance(opts *bind.TransactOpts, provenance string) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "setProvenance", provenance)
}

// SetProvenance is a paid mutator transaction binding the contract method 0xffe630b5.
//
// Solidity: function setProvenance(string provenance) returns()
func (_Mycontract *MycontractSession) SetProvenance(provenance string) (*types.Transaction, error) {
	return _Mycontract.Contract.SetProvenance(&_Mycontract.TransactOpts, provenance)
}

// SetProvenance is a paid mutator transaction binding the contract method 0xffe630b5.
//
// Solidity: function setProvenance(string provenance) returns()
func (_Mycontract *MycontractTransactorSession) SetProvenance(provenance string) (*types.Transaction, error) {
	return _Mycontract.Contract.SetProvenance(&_Mycontract.TransactOpts, provenance)
}

// SetSaleState is a paid mutator transaction binding the contract method 0xc4e37095.
//
// Solidity: function setSaleState(bool newState) returns()
func (_Mycontract *MycontractTransactor) SetSaleState(opts *bind.TransactOpts, newState bool) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "setSaleState", newState)
}

// SetSaleState is a paid mutator transaction binding the contract method 0xc4e37095.
//
// Solidity: function setSaleState(bool newState) returns()
func (_Mycontract *MycontractSession) SetSaleState(newState bool) (*types.Transaction, error) {
	return _Mycontract.Contract.SetSaleState(&_Mycontract.TransactOpts, newState)
}

// SetSaleState is a paid mutator transaction binding the contract method 0xc4e37095.
//
// Solidity: function setSaleState(bool newState) returns()
func (_Mycontract *MycontractTransactorSession) SetSaleState(newState bool) (*types.Transaction, error) {
	return _Mycontract.Contract.SetSaleState(&_Mycontract.TransactOpts, newState)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Mycontract *MycontractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Mycontract *MycontractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.TransferFrom(&_Mycontract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Mycontract *MycontractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.TransferFrom(&_Mycontract.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mycontract *MycontractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mycontract *MycontractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.TransferOwnership(&_Mycontract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mycontract *MycontractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mycontract.Contract.TransferOwnership(&_Mycontract.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Mycontract *MycontractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Mycontract *MycontractSession) Withdraw() (*types.Transaction, error) {
	return _Mycontract.Contract.Withdraw(&_Mycontract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Mycontract *MycontractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Mycontract.Contract.Withdraw(&_Mycontract.TransactOpts)
}

// MycontractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Mycontract contract.
type MycontractApprovalIterator struct {
	Event *MycontractApproval // Event containing the contract specifics and raw log

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
func (it *MycontractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractApproval)
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
		it.Event = new(MycontractApproval)
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
func (it *MycontractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractApproval represents a Approval event raised by the Mycontract contract.
type MycontractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Mycontract *MycontractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MycontractApprovalIterator, error) {

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

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MycontractApprovalIterator{contract: _Mycontract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Mycontract *MycontractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MycontractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractApproval)
				if err := _Mycontract.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Mycontract *MycontractFilterer) ParseApproval(log types.Log) (*MycontractApproval, error) {
	event := new(MycontractApproval)
	if err := _Mycontract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Mycontract contract.
type MycontractApprovalForAllIterator struct {
	Event *MycontractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MycontractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractApprovalForAll)
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
		it.Event = new(MycontractApprovalForAll)
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
func (it *MycontractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractApprovalForAll represents a ApprovalForAll event raised by the Mycontract contract.
type MycontractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Mycontract *MycontractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MycontractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MycontractApprovalForAllIterator{contract: _Mycontract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Mycontract *MycontractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MycontractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractApprovalForAll)
				if err := _Mycontract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Mycontract *MycontractFilterer) ParseApprovalForAll(log types.Log) (*MycontractApprovalForAll, error) {
	event := new(MycontractApprovalForAll)
	if err := _Mycontract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mycontract contract.
type MycontractOwnershipTransferredIterator struct {
	Event *MycontractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MycontractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractOwnershipTransferred)
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
		it.Event = new(MycontractOwnershipTransferred)
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
func (it *MycontractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractOwnershipTransferred represents a OwnershipTransferred event raised by the Mycontract contract.
type MycontractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mycontract *MycontractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MycontractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MycontractOwnershipTransferredIterator{contract: _Mycontract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mycontract *MycontractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MycontractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractOwnershipTransferred)
				if err := _Mycontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mycontract *MycontractFilterer) ParseOwnershipTransferred(log types.Log) (*MycontractOwnershipTransferred, error) {
	event := new(MycontractOwnershipTransferred)
	if err := _Mycontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Mycontract contract.
type MycontractTransferIterator struct {
	Event *MycontractTransfer // Event containing the contract specifics and raw log

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
func (it *MycontractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractTransfer)
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
		it.Event = new(MycontractTransfer)
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
func (it *MycontractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractTransfer represents a Transfer event raised by the Mycontract contract.
type MycontractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Mycontract *MycontractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MycontractTransferIterator, error) {

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

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MycontractTransferIterator{contract: _Mycontract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Mycontract *MycontractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MycontractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractTransfer)
				if err := _Mycontract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Mycontract *MycontractFilterer) ParseTransfer(log types.Log) (*MycontractTransfer, error) {
	event := new(MycontractTransfer)
	if err := _Mycontract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
