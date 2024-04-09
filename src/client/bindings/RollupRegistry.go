// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// RollupRegistryMetaData contains all meta data concerning the RollupRegistry contract.
var RollupRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"RollUpDeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"RollUpRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"RollUpSuspended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"deRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollups\",\"outputs\":[{\"internalType\":\"uint256[2][]\",\"name\":\"\",\"type\":\"uint256[2][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"layerNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numOfWatchtowers\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rollups\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"layerNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numOfWatchtowers\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"suspend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// RollupRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupRegistryMetaData.ABI instead.
var RollupRegistryABI = RollupRegistryMetaData.ABI

// RollupRegistry is an auto generated Go binding around an Ethereum contract.
type RollupRegistry struct {
	RollupRegistryCaller     // Read-only binding to the contract
	RollupRegistryTransactor // Write-only binding to the contract
	RollupRegistryFilterer   // Log filterer for contract events
}

// RollupRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupRegistrySession struct {
	Contract     *RollupRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupRegistryCallerSession struct {
	Contract *RollupRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RollupRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupRegistryTransactorSession struct {
	Contract     *RollupRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RollupRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupRegistryRaw struct {
	Contract *RollupRegistry // Generic contract binding to access the raw methods on
}

// RollupRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupRegistryCallerRaw struct {
	Contract *RollupRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RollupRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupRegistryTransactorRaw struct {
	Contract *RollupRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupRegistry creates a new instance of RollupRegistry, bound to a specific deployed contract.
func NewRollupRegistry(address common.Address, backend bind.ContractBackend) (*RollupRegistry, error) {
	contract, err := bindRollupRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupRegistry{RollupRegistryCaller: RollupRegistryCaller{contract: contract}, RollupRegistryTransactor: RollupRegistryTransactor{contract: contract}, RollupRegistryFilterer: RollupRegistryFilterer{contract: contract}}, nil
}

// NewRollupRegistryCaller creates a new read-only instance of RollupRegistry, bound to a specific deployed contract.
func NewRollupRegistryCaller(address common.Address, caller bind.ContractCaller) (*RollupRegistryCaller, error) {
	contract, err := bindRollupRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupRegistryCaller{contract: contract}, nil
}

// NewRollupRegistryTransactor creates a new write-only instance of RollupRegistry, bound to a specific deployed contract.
func NewRollupRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupRegistryTransactor, error) {
	contract, err := bindRollupRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupRegistryTransactor{contract: contract}, nil
}

// NewRollupRegistryFilterer creates a new log filterer instance of RollupRegistry, bound to a specific deployed contract.
func NewRollupRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupRegistryFilterer, error) {
	contract, err := bindRollupRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupRegistryFilterer{contract: contract}, nil
}

// bindRollupRegistry binds a generic wrapper to an already deployed contract.
func bindRollupRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RollupRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupRegistry *RollupRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupRegistry.Contract.RollupRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupRegistry *RollupRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupRegistry.Contract.RollupRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupRegistry *RollupRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupRegistry.Contract.RollupRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupRegistry *RollupRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupRegistry *RollupRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupRegistry *RollupRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetRollups is a free data retrieval call binding the contract method 0x45d06c66.
//
// Solidity: function getRollups() view returns(uint256[2][])
func (_RollupRegistry *RollupRegistryCaller) GetRollups(opts *bind.CallOpts) ([][2]*big.Int, error) {
	var out []interface{}
	err := _RollupRegistry.contract.Call(opts, &out, "getRollups")

	if err != nil {
		return *new([][2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][2]*big.Int)).(*[][2]*big.Int)

	return out0, err

}

// GetRollups is a free data retrieval call binding the contract method 0x45d06c66.
//
// Solidity: function getRollups() view returns(uint256[2][])
func (_RollupRegistry *RollupRegistrySession) GetRollups() ([][2]*big.Int, error) {
	return _RollupRegistry.Contract.GetRollups(&_RollupRegistry.CallOpts)
}

// GetRollups is a free data retrieval call binding the contract method 0x45d06c66.
//
// Solidity: function getRollups() view returns(uint256[2][])
func (_RollupRegistry *RollupRegistryCallerSession) GetRollups() ([][2]*big.Int, error) {
	return _RollupRegistry.Contract.GetRollups(&_RollupRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupRegistry *RollupRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupRegistry *RollupRegistrySession) Owner() (common.Address, error) {
	return _RollupRegistry.Contract.Owner(&_RollupRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupRegistry *RollupRegistryCallerSession) Owner() (common.Address, error) {
	return _RollupRegistry.Contract.Owner(&_RollupRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RollupRegistry *RollupRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RollupRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RollupRegistry *RollupRegistrySession) Paused() (bool, error) {
	return _RollupRegistry.Contract.Paused(&_RollupRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RollupRegistry *RollupRegistryCallerSession) Paused() (bool, error) {
	return _RollupRegistry.Contract.Paused(&_RollupRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RollupRegistry *RollupRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RollupRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RollupRegistry *RollupRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _RollupRegistry.Contract.ProxiableUUID(&_RollupRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RollupRegistry *RollupRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _RollupRegistry.Contract.ProxiableUUID(&_RollupRegistry.CallOpts)
}

// Rollups is a free data retrieval call binding the contract method 0xb794e5a3.
//
// Solidity: function rollups(uint256 ) view returns(bool isRegistered, bool isActive, uint256 layerNumber, uint256 numOfWatchtowers)
func (_RollupRegistry *RollupRegistryCaller) Rollups(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IsRegistered     bool
	IsActive         bool
	LayerNumber      *big.Int
	NumOfWatchtowers *big.Int
}, error) {
	var out []interface{}
	err := _RollupRegistry.contract.Call(opts, &out, "rollups", arg0)

	outstruct := new(struct {
		IsRegistered     bool
		IsActive         bool
		LayerNumber      *big.Int
		NumOfWatchtowers *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsRegistered = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsActive = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.LayerNumber = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.NumOfWatchtowers = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Rollups is a free data retrieval call binding the contract method 0xb794e5a3.
//
// Solidity: function rollups(uint256 ) view returns(bool isRegistered, bool isActive, uint256 layerNumber, uint256 numOfWatchtowers)
func (_RollupRegistry *RollupRegistrySession) Rollups(arg0 *big.Int) (struct {
	IsRegistered     bool
	IsActive         bool
	LayerNumber      *big.Int
	NumOfWatchtowers *big.Int
}, error) {
	return _RollupRegistry.Contract.Rollups(&_RollupRegistry.CallOpts, arg0)
}

// Rollups is a free data retrieval call binding the contract method 0xb794e5a3.
//
// Solidity: function rollups(uint256 ) view returns(bool isRegistered, bool isActive, uint256 layerNumber, uint256 numOfWatchtowers)
func (_RollupRegistry *RollupRegistryCallerSession) Rollups(arg0 *big.Int) (struct {
	IsRegistered     bool
	IsActive         bool
	LayerNumber      *big.Int
	NumOfWatchtowers *big.Int
}, error) {
	return _RollupRegistry.Contract.Rollups(&_RollupRegistry.CallOpts, arg0)
}

// DeRegister is a paid mutator transaction binding the contract method 0x480db2bb.
//
// Solidity: function deRegister(uint256 chainID) returns()
func (_RollupRegistry *RollupRegistryTransactor) DeRegister(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _RollupRegistry.contract.Transact(opts, "deRegister", chainID)
}

// DeRegister is a paid mutator transaction binding the contract method 0x480db2bb.
//
// Solidity: function deRegister(uint256 chainID) returns()
func (_RollupRegistry *RollupRegistrySession) DeRegister(chainID *big.Int) (*types.Transaction, error) {
	return _RollupRegistry.Contract.DeRegister(&_RollupRegistry.TransactOpts, chainID)
}

// DeRegister is a paid mutator transaction binding the contract method 0x480db2bb.
//
// Solidity: function deRegister(uint256 chainID) returns()
func (_RollupRegistry *RollupRegistryTransactorSession) DeRegister(chainID *big.Int) (*types.Transaction, error) {
	return _RollupRegistry.Contract.DeRegister(&_RollupRegistry.TransactOpts, chainID)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RollupRegistry *RollupRegistryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupRegistry.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RollupRegistry *RollupRegistrySession) Initialize() (*types.Transaction, error) {
	return _RollupRegistry.Contract.Initialize(&_RollupRegistry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RollupRegistry *RollupRegistryTransactorSession) Initialize() (*types.Transaction, error) {
	return _RollupRegistry.Contract.Initialize(&_RollupRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RollupRegistry *RollupRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RollupRegistry *RollupRegistrySession) Pause() (*types.Transaction, error) {
	return _RollupRegistry.Contract.Pause(&_RollupRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RollupRegistry *RollupRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _RollupRegistry.Contract.Pause(&_RollupRegistry.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0xfaa5c564.
//
// Solidity: function register(uint256 chainID, uint256 layerNumber, uint256 numOfWatchtowers) returns()
func (_RollupRegistry *RollupRegistryTransactor) Register(opts *bind.TransactOpts, chainID *big.Int, layerNumber *big.Int, numOfWatchtowers *big.Int) (*types.Transaction, error) {
	return _RollupRegistry.contract.Transact(opts, "register", chainID, layerNumber, numOfWatchtowers)
}

// Register is a paid mutator transaction binding the contract method 0xfaa5c564.
//
// Solidity: function register(uint256 chainID, uint256 layerNumber, uint256 numOfWatchtowers) returns()
func (_RollupRegistry *RollupRegistrySession) Register(chainID *big.Int, layerNumber *big.Int, numOfWatchtowers *big.Int) (*types.Transaction, error) {
	return _RollupRegistry.Contract.Register(&_RollupRegistry.TransactOpts, chainID, layerNumber, numOfWatchtowers)
}

// Register is a paid mutator transaction binding the contract method 0xfaa5c564.
//
// Solidity: function register(uint256 chainID, uint256 layerNumber, uint256 numOfWatchtowers) returns()
func (_RollupRegistry *RollupRegistryTransactorSession) Register(chainID *big.Int, layerNumber *big.Int, numOfWatchtowers *big.Int) (*types.Transaction, error) {
	return _RollupRegistry.Contract.Register(&_RollupRegistry.TransactOpts, chainID, layerNumber, numOfWatchtowers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupRegistry *RollupRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupRegistry *RollupRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _RollupRegistry.Contract.RenounceOwnership(&_RollupRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupRegistry *RollupRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RollupRegistry.Contract.RenounceOwnership(&_RollupRegistry.TransactOpts)
}

// Suspend is a paid mutator transaction binding the contract method 0x4b865846.
//
// Solidity: function suspend(uint256 chainID) returns()
func (_RollupRegistry *RollupRegistryTransactor) Suspend(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _RollupRegistry.contract.Transact(opts, "suspend", chainID)
}

// Suspend is a paid mutator transaction binding the contract method 0x4b865846.
//
// Solidity: function suspend(uint256 chainID) returns()
func (_RollupRegistry *RollupRegistrySession) Suspend(chainID *big.Int) (*types.Transaction, error) {
	return _RollupRegistry.Contract.Suspend(&_RollupRegistry.TransactOpts, chainID)
}

// Suspend is a paid mutator transaction binding the contract method 0x4b865846.
//
// Solidity: function suspend(uint256 chainID) returns()
func (_RollupRegistry *RollupRegistryTransactorSession) Suspend(chainID *big.Int) (*types.Transaction, error) {
	return _RollupRegistry.Contract.Suspend(&_RollupRegistry.TransactOpts, chainID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupRegistry *RollupRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RollupRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupRegistry *RollupRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RollupRegistry.Contract.TransferOwnership(&_RollupRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupRegistry *RollupRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RollupRegistry.Contract.TransferOwnership(&_RollupRegistry.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RollupRegistry *RollupRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RollupRegistry *RollupRegistrySession) Unpause() (*types.Transaction, error) {
	return _RollupRegistry.Contract.Unpause(&_RollupRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RollupRegistry *RollupRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _RollupRegistry.Contract.Unpause(&_RollupRegistry.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RollupRegistry *RollupRegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _RollupRegistry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RollupRegistry *RollupRegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _RollupRegistry.Contract.UpgradeTo(&_RollupRegistry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RollupRegistry *RollupRegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _RollupRegistry.Contract.UpgradeTo(&_RollupRegistry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RollupRegistry *RollupRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RollupRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RollupRegistry *RollupRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RollupRegistry.Contract.UpgradeToAndCall(&_RollupRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RollupRegistry *RollupRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RollupRegistry.Contract.UpgradeToAndCall(&_RollupRegistry.TransactOpts, newImplementation, data)
}

// RollupRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the RollupRegistry contract.
type RollupRegistryAdminChangedIterator struct {
	Event *RollupRegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *RollupRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRegistryAdminChanged)
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
		it.Event = new(RollupRegistryAdminChanged)
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
func (it *RollupRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRegistryAdminChanged represents a AdminChanged event raised by the RollupRegistry contract.
type RollupRegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RollupRegistry *RollupRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*RollupRegistryAdminChangedIterator, error) {

	logs, sub, err := _RollupRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &RollupRegistryAdminChangedIterator{contract: _RollupRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RollupRegistry *RollupRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *RollupRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _RollupRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRegistryAdminChanged)
				if err := _RollupRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RollupRegistry *RollupRegistryFilterer) ParseAdminChanged(log types.Log) (*RollupRegistryAdminChanged, error) {
	event := new(RollupRegistryAdminChanged)
	if err := _RollupRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the RollupRegistry contract.
type RollupRegistryBeaconUpgradedIterator struct {
	Event *RollupRegistryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *RollupRegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRegistryBeaconUpgraded)
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
		it.Event = new(RollupRegistryBeaconUpgraded)
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
func (it *RollupRegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRegistryBeaconUpgraded represents a BeaconUpgraded event raised by the RollupRegistry contract.
type RollupRegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RollupRegistry *RollupRegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*RollupRegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _RollupRegistry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &RollupRegistryBeaconUpgradedIterator{contract: _RollupRegistry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RollupRegistry *RollupRegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *RollupRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _RollupRegistry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRegistryBeaconUpgraded)
				if err := _RollupRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RollupRegistry *RollupRegistryFilterer) ParseBeaconUpgraded(log types.Log) (*RollupRegistryBeaconUpgraded, error) {
	event := new(RollupRegistryBeaconUpgraded)
	if err := _RollupRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RollupRegistry contract.
type RollupRegistryInitializedIterator struct {
	Event *RollupRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *RollupRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRegistryInitialized)
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
		it.Event = new(RollupRegistryInitialized)
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
func (it *RollupRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRegistryInitialized represents a Initialized event raised by the RollupRegistry contract.
type RollupRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RollupRegistry *RollupRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*RollupRegistryInitializedIterator, error) {

	logs, sub, err := _RollupRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RollupRegistryInitializedIterator{contract: _RollupRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RollupRegistry *RollupRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RollupRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _RollupRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRegistryInitialized)
				if err := _RollupRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RollupRegistry *RollupRegistryFilterer) ParseInitialized(log types.Log) (*RollupRegistryInitialized, error) {
	event := new(RollupRegistryInitialized)
	if err := _RollupRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RollupRegistry contract.
type RollupRegistryOwnershipTransferredIterator struct {
	Event *RollupRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RollupRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRegistryOwnershipTransferred)
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
		it.Event = new(RollupRegistryOwnershipTransferred)
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
func (it *RollupRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the RollupRegistry contract.
type RollupRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupRegistry *RollupRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RollupRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RollupRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RollupRegistryOwnershipTransferredIterator{contract: _RollupRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupRegistry *RollupRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RollupRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RollupRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRegistryOwnershipTransferred)
				if err := _RollupRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RollupRegistry *RollupRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*RollupRegistryOwnershipTransferred, error) {
	event := new(RollupRegistryOwnershipTransferred)
	if err := _RollupRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the RollupRegistry contract.
type RollupRegistryPausedIterator struct {
	Event *RollupRegistryPaused // Event containing the contract specifics and raw log

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
func (it *RollupRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRegistryPaused)
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
		it.Event = new(RollupRegistryPaused)
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
func (it *RollupRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRegistryPaused represents a Paused event raised by the RollupRegistry contract.
type RollupRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RollupRegistry *RollupRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*RollupRegistryPausedIterator, error) {

	logs, sub, err := _RollupRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RollupRegistryPausedIterator{contract: _RollupRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RollupRegistry *RollupRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RollupRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _RollupRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRegistryPaused)
				if err := _RollupRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RollupRegistry *RollupRegistryFilterer) ParsePaused(log types.Log) (*RollupRegistryPaused, error) {
	event := new(RollupRegistryPaused)
	if err := _RollupRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRegistryRollUpDeRegisteredIterator is returned from FilterRollUpDeRegistered and is used to iterate over the raw logs and unpacked data for RollUpDeRegistered events raised by the RollupRegistry contract.
type RollupRegistryRollUpDeRegisteredIterator struct {
	Event *RollupRegistryRollUpDeRegistered // Event containing the contract specifics and raw log

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
func (it *RollupRegistryRollUpDeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRegistryRollUpDeRegistered)
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
		it.Event = new(RollupRegistryRollUpDeRegistered)
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
func (it *RollupRegistryRollUpDeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRegistryRollUpDeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRegistryRollUpDeRegistered represents a RollUpDeRegistered event raised by the RollupRegistry contract.
type RollupRegistryRollUpDeRegistered struct {
	ChainID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRollUpDeRegistered is a free log retrieval operation binding the contract event 0x67c3283580b3cdb043e7fcaf3df415e863fbc94b0a745c7e001428f02689838c.
//
// Solidity: event RollUpDeRegistered(uint256 chainID)
func (_RollupRegistry *RollupRegistryFilterer) FilterRollUpDeRegistered(opts *bind.FilterOpts) (*RollupRegistryRollUpDeRegisteredIterator, error) {

	logs, sub, err := _RollupRegistry.contract.FilterLogs(opts, "RollUpDeRegistered")
	if err != nil {
		return nil, err
	}
	return &RollupRegistryRollUpDeRegisteredIterator{contract: _RollupRegistry.contract, event: "RollUpDeRegistered", logs: logs, sub: sub}, nil
}

// WatchRollUpDeRegistered is a free log subscription operation binding the contract event 0x67c3283580b3cdb043e7fcaf3df415e863fbc94b0a745c7e001428f02689838c.
//
// Solidity: event RollUpDeRegistered(uint256 chainID)
func (_RollupRegistry *RollupRegistryFilterer) WatchRollUpDeRegistered(opts *bind.WatchOpts, sink chan<- *RollupRegistryRollUpDeRegistered) (event.Subscription, error) {

	logs, sub, err := _RollupRegistry.contract.WatchLogs(opts, "RollUpDeRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRegistryRollUpDeRegistered)
				if err := _RollupRegistry.contract.UnpackLog(event, "RollUpDeRegistered", log); err != nil {
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

// ParseRollUpDeRegistered is a log parse operation binding the contract event 0x67c3283580b3cdb043e7fcaf3df415e863fbc94b0a745c7e001428f02689838c.
//
// Solidity: event RollUpDeRegistered(uint256 chainID)
func (_RollupRegistry *RollupRegistryFilterer) ParseRollUpDeRegistered(log types.Log) (*RollupRegistryRollUpDeRegistered, error) {
	event := new(RollupRegistryRollUpDeRegistered)
	if err := _RollupRegistry.contract.UnpackLog(event, "RollUpDeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRegistryRollUpRegisteredIterator is returned from FilterRollUpRegistered and is used to iterate over the raw logs and unpacked data for RollUpRegistered events raised by the RollupRegistry contract.
type RollupRegistryRollUpRegisteredIterator struct {
	Event *RollupRegistryRollUpRegistered // Event containing the contract specifics and raw log

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
func (it *RollupRegistryRollUpRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRegistryRollUpRegistered)
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
		it.Event = new(RollupRegistryRollUpRegistered)
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
func (it *RollupRegistryRollUpRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRegistryRollUpRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRegistryRollUpRegistered represents a RollUpRegistered event raised by the RollupRegistry contract.
type RollupRegistryRollUpRegistered struct {
	ChainID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRollUpRegistered is a free log retrieval operation binding the contract event 0x5cfa33b2a612d25e41e15ef94201e6dca255fd885ad00a96dbb4802d3f447e9e.
//
// Solidity: event RollUpRegistered(uint256 chainID)
func (_RollupRegistry *RollupRegistryFilterer) FilterRollUpRegistered(opts *bind.FilterOpts) (*RollupRegistryRollUpRegisteredIterator, error) {

	logs, sub, err := _RollupRegistry.contract.FilterLogs(opts, "RollUpRegistered")
	if err != nil {
		return nil, err
	}
	return &RollupRegistryRollUpRegisteredIterator{contract: _RollupRegistry.contract, event: "RollUpRegistered", logs: logs, sub: sub}, nil
}

// WatchRollUpRegistered is a free log subscription operation binding the contract event 0x5cfa33b2a612d25e41e15ef94201e6dca255fd885ad00a96dbb4802d3f447e9e.
//
// Solidity: event RollUpRegistered(uint256 chainID)
func (_RollupRegistry *RollupRegistryFilterer) WatchRollUpRegistered(opts *bind.WatchOpts, sink chan<- *RollupRegistryRollUpRegistered) (event.Subscription, error) {

	logs, sub, err := _RollupRegistry.contract.WatchLogs(opts, "RollUpRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRegistryRollUpRegistered)
				if err := _RollupRegistry.contract.UnpackLog(event, "RollUpRegistered", log); err != nil {
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

// ParseRollUpRegistered is a log parse operation binding the contract event 0x5cfa33b2a612d25e41e15ef94201e6dca255fd885ad00a96dbb4802d3f447e9e.
//
// Solidity: event RollUpRegistered(uint256 chainID)
func (_RollupRegistry *RollupRegistryFilterer) ParseRollUpRegistered(log types.Log) (*RollupRegistryRollUpRegistered, error) {
	event := new(RollupRegistryRollUpRegistered)
	if err := _RollupRegistry.contract.UnpackLog(event, "RollUpRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRegistryRollUpSuspendedIterator is returned from FilterRollUpSuspended and is used to iterate over the raw logs and unpacked data for RollUpSuspended events raised by the RollupRegistry contract.
type RollupRegistryRollUpSuspendedIterator struct {
	Event *RollupRegistryRollUpSuspended // Event containing the contract specifics and raw log

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
func (it *RollupRegistryRollUpSuspendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRegistryRollUpSuspended)
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
		it.Event = new(RollupRegistryRollUpSuspended)
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
func (it *RollupRegistryRollUpSuspendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRegistryRollUpSuspendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRegistryRollUpSuspended represents a RollUpSuspended event raised by the RollupRegistry contract.
type RollupRegistryRollUpSuspended struct {
	ChainID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRollUpSuspended is a free log retrieval operation binding the contract event 0xa1ac7f558d2b2897a2364d8ef7692e559af85eff21fc58ca8ee852f03ff69c01.
//
// Solidity: event RollUpSuspended(uint256 chainID)
func (_RollupRegistry *RollupRegistryFilterer) FilterRollUpSuspended(opts *bind.FilterOpts) (*RollupRegistryRollUpSuspendedIterator, error) {

	logs, sub, err := _RollupRegistry.contract.FilterLogs(opts, "RollUpSuspended")
	if err != nil {
		return nil, err
	}
	return &RollupRegistryRollUpSuspendedIterator{contract: _RollupRegistry.contract, event: "RollUpSuspended", logs: logs, sub: sub}, nil
}

// WatchRollUpSuspended is a free log subscription operation binding the contract event 0xa1ac7f558d2b2897a2364d8ef7692e559af85eff21fc58ca8ee852f03ff69c01.
//
// Solidity: event RollUpSuspended(uint256 chainID)
func (_RollupRegistry *RollupRegistryFilterer) WatchRollUpSuspended(opts *bind.WatchOpts, sink chan<- *RollupRegistryRollUpSuspended) (event.Subscription, error) {

	logs, sub, err := _RollupRegistry.contract.WatchLogs(opts, "RollUpSuspended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRegistryRollUpSuspended)
				if err := _RollupRegistry.contract.UnpackLog(event, "RollUpSuspended", log); err != nil {
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

// ParseRollUpSuspended is a log parse operation binding the contract event 0xa1ac7f558d2b2897a2364d8ef7692e559af85eff21fc58ca8ee852f03ff69c01.
//
// Solidity: event RollUpSuspended(uint256 chainID)
func (_RollupRegistry *RollupRegistryFilterer) ParseRollUpSuspended(log types.Log) (*RollupRegistryRollUpSuspended, error) {
	event := new(RollupRegistryRollUpSuspended)
	if err := _RollupRegistry.contract.UnpackLog(event, "RollUpSuspended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the RollupRegistry contract.
type RollupRegistryUnpausedIterator struct {
	Event *RollupRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *RollupRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRegistryUnpaused)
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
		it.Event = new(RollupRegistryUnpaused)
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
func (it *RollupRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRegistryUnpaused represents a Unpaused event raised by the RollupRegistry contract.
type RollupRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RollupRegistry *RollupRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RollupRegistryUnpausedIterator, error) {

	logs, sub, err := _RollupRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RollupRegistryUnpausedIterator{contract: _RollupRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RollupRegistry *RollupRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RollupRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _RollupRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRegistryUnpaused)
				if err := _RollupRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RollupRegistry *RollupRegistryFilterer) ParseUnpaused(log types.Log) (*RollupRegistryUnpaused, error) {
	event := new(RollupRegistryUnpaused)
	if err := _RollupRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the RollupRegistry contract.
type RollupRegistryUpgradedIterator struct {
	Event *RollupRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *RollupRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRegistryUpgraded)
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
		it.Event = new(RollupRegistryUpgraded)
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
func (it *RollupRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRegistryUpgraded represents a Upgraded event raised by the RollupRegistry contract.
type RollupRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RollupRegistry *RollupRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RollupRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RollupRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RollupRegistryUpgradedIterator{contract: _RollupRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RollupRegistry *RollupRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RollupRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RollupRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRegistryUpgraded)
				if err := _RollupRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RollupRegistry *RollupRegistryFilterer) ParseUpgraded(log types.Log) (*RollupRegistryUpgraded, error) {
	event := new(RollupRegistryUpgraded)
	if err := _RollupRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
