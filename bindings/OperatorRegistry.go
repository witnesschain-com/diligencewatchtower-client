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

// IOperatorRegistryOperator is an auto generated low-level Go binding around an user-defined struct.
type IOperatorRegistryOperator struct {
	OperatorAddress common.Address
	IsRegistered    bool
	WatchStatus     uint8
	OperatorType    uint8
	IsActive        bool
}

// OperatorRegistryMetaData contains all meta data concerning the OperatorRegistry contract.
var OperatorRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"enumIOperatorRegistry.WatchStatus\",\"name\":\"watchStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumIOperatorRegistry.OperatorType\",\"name\":\"operatorType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIOperatorRegistry.Operator\",\"name\":\"operator\",\"type\":\"tuple\"}],\"name\":\"OperatorDeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"enumIOperatorRegistry.WatchStatus\",\"name\":\"watchStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumIOperatorRegistry.OperatorType\",\"name\":\"operatorType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIOperatorRegistry.Operator\",\"name\":\"operator\",\"type\":\"tuple\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"enumIOperatorRegistry.WatchStatus\",\"name\":\"watchStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumIOperatorRegistry.OperatorType\",\"name\":\"operatorType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIOperatorRegistry.Operator\",\"name\":\"operator\",\"type\":\"tuple\"}],\"name\":\"OperatorSuspended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"operatorsList\",\"type\":\"address[]\"}],\"name\":\"addToOperatorWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"deRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllWatchtowerAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"watchtower\",\"type\":\"address\"}],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegationManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_slasherAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isActiveOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"watchtower\",\"type\":\"address\"}],\"name\":\"isValidWatchtower\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"enumIOperatorRegistry.WatchStatus\",\"name\":\"watchStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumIOperatorRegistry.OperatorType\",\"name\":\"operatorType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"watchtowerAddress\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"operatorsList\",\"type\":\"address[]\"}],\"name\":\"removeFromOperatorWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegationManagerAddress\",\"type\":\"address\"}],\"name\":\"setDelegationManagerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_slasherAddress\",\"type\":\"address\"}],\"name\":\"setSlasherAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slasherAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"suspend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"watchTowerAddressesList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OperatorRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use OperatorRegistryMetaData.ABI instead.
var OperatorRegistryABI = OperatorRegistryMetaData.ABI

// OperatorRegistry is an auto generated Go binding around an Ethereum contract.
type OperatorRegistry struct {
	OperatorRegistryCaller     // Read-only binding to the contract
	OperatorRegistryTransactor // Write-only binding to the contract
	OperatorRegistryFilterer   // Log filterer for contract events
}

// OperatorRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperatorRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperatorRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperatorRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperatorRegistrySession struct {
	Contract     *OperatorRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OperatorRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperatorRegistryCallerSession struct {
	Contract *OperatorRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// OperatorRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperatorRegistryTransactorSession struct {
	Contract     *OperatorRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OperatorRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperatorRegistryRaw struct {
	Contract *OperatorRegistry // Generic contract binding to access the raw methods on
}

// OperatorRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperatorRegistryCallerRaw struct {
	Contract *OperatorRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// OperatorRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperatorRegistryTransactorRaw struct {
	Contract *OperatorRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperatorRegistry creates a new instance of OperatorRegistry, bound to a specific deployed contract.
func NewOperatorRegistry(address common.Address, backend bind.ContractBackend) (*OperatorRegistry, error) {
	contract, err := bindOperatorRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistry{OperatorRegistryCaller: OperatorRegistryCaller{contract: contract}, OperatorRegistryTransactor: OperatorRegistryTransactor{contract: contract}, OperatorRegistryFilterer: OperatorRegistryFilterer{contract: contract}}, nil
}

// NewOperatorRegistryCaller creates a new read-only instance of OperatorRegistry, bound to a specific deployed contract.
func NewOperatorRegistryCaller(address common.Address, caller bind.ContractCaller) (*OperatorRegistryCaller, error) {
	contract, err := bindOperatorRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryCaller{contract: contract}, nil
}

// NewOperatorRegistryTransactor creates a new write-only instance of OperatorRegistry, bound to a specific deployed contract.
func NewOperatorRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*OperatorRegistryTransactor, error) {
	contract, err := bindOperatorRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryTransactor{contract: contract}, nil
}

// NewOperatorRegistryFilterer creates a new log filterer instance of OperatorRegistry, bound to a specific deployed contract.
func NewOperatorRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*OperatorRegistryFilterer, error) {
	contract, err := bindOperatorRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryFilterer{contract: contract}, nil
}

// bindOperatorRegistry binds a generic wrapper to an already deployed contract.
func bindOperatorRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OperatorRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorRegistry *OperatorRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorRegistry.Contract.OperatorRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorRegistry *OperatorRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.OperatorRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorRegistry *OperatorRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.OperatorRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorRegistry *OperatorRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorRegistry *OperatorRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorRegistry *OperatorRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.contract.Transact(opts, method, params...)
}

// DelegationManagerAddress is a free data retrieval call binding the contract method 0xb15e6689.
//
// Solidity: function delegationManagerAddress() view returns(address)
func (_OperatorRegistry *OperatorRegistryCaller) DelegationManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "delegationManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManagerAddress is a free data retrieval call binding the contract method 0xb15e6689.
//
// Solidity: function delegationManagerAddress() view returns(address)
func (_OperatorRegistry *OperatorRegistrySession) DelegationManagerAddress() (common.Address, error) {
	return _OperatorRegistry.Contract.DelegationManagerAddress(&_OperatorRegistry.CallOpts)
}

// DelegationManagerAddress is a free data retrieval call binding the contract method 0xb15e6689.
//
// Solidity: function delegationManagerAddress() view returns(address)
func (_OperatorRegistry *OperatorRegistryCallerSession) DelegationManagerAddress() (common.Address, error) {
	return _OperatorRegistry.Contract.DelegationManagerAddress(&_OperatorRegistry.CallOpts)
}

// GetAllWatchtowerAddresses is a free data retrieval call binding the contract method 0xa7e59b65.
//
// Solidity: function getAllWatchtowerAddresses() view returns(address[])
func (_OperatorRegistry *OperatorRegistryCaller) GetAllWatchtowerAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "getAllWatchtowerAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllWatchtowerAddresses is a free data retrieval call binding the contract method 0xa7e59b65.
//
// Solidity: function getAllWatchtowerAddresses() view returns(address[])
func (_OperatorRegistry *OperatorRegistrySession) GetAllWatchtowerAddresses() ([]common.Address, error) {
	return _OperatorRegistry.Contract.GetAllWatchtowerAddresses(&_OperatorRegistry.CallOpts)
}

// GetAllWatchtowerAddresses is a free data retrieval call binding the contract method 0xa7e59b65.
//
// Solidity: function getAllWatchtowerAddresses() view returns(address[])
func (_OperatorRegistry *OperatorRegistryCallerSession) GetAllWatchtowerAddresses() ([]common.Address, error) {
	return _OperatorRegistry.Contract.GetAllWatchtowerAddresses(&_OperatorRegistry.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address watchtower) view returns(address operator)
func (_OperatorRegistry *OperatorRegistryCaller) GetOperator(opts *bind.CallOpts, watchtower common.Address) (common.Address, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "getOperator", watchtower)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address watchtower) view returns(address operator)
func (_OperatorRegistry *OperatorRegistrySession) GetOperator(watchtower common.Address) (common.Address, error) {
	return _OperatorRegistry.Contract.GetOperator(&_OperatorRegistry.CallOpts, watchtower)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address watchtower) view returns(address operator)
func (_OperatorRegistry *OperatorRegistryCallerSession) GetOperator(watchtower common.Address) (common.Address, error) {
	return _OperatorRegistry.Contract.GetOperator(&_OperatorRegistry.CallOpts, watchtower)
}

// IsActiveOperator is a free data retrieval call binding the contract method 0x3367cca5.
//
// Solidity: function isActiveOperator(address operator) view returns(bool)
func (_OperatorRegistry *OperatorRegistryCaller) IsActiveOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "isActiveOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveOperator is a free data retrieval call binding the contract method 0x3367cca5.
//
// Solidity: function isActiveOperator(address operator) view returns(bool)
func (_OperatorRegistry *OperatorRegistrySession) IsActiveOperator(operator common.Address) (bool, error) {
	return _OperatorRegistry.Contract.IsActiveOperator(&_OperatorRegistry.CallOpts, operator)
}

// IsActiveOperator is a free data retrieval call binding the contract method 0x3367cca5.
//
// Solidity: function isActiveOperator(address operator) view returns(bool)
func (_OperatorRegistry *OperatorRegistryCallerSession) IsActiveOperator(operator common.Address) (bool, error) {
	return _OperatorRegistry.Contract.IsActiveOperator(&_OperatorRegistry.CallOpts, operator)
}

// IsValidWatchtower is a free data retrieval call binding the contract method 0x9a521382.
//
// Solidity: function isValidWatchtower(address watchtower) view returns(bool)
func (_OperatorRegistry *OperatorRegistryCaller) IsValidWatchtower(opts *bind.CallOpts, watchtower common.Address) (bool, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "isValidWatchtower", watchtower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidWatchtower is a free data retrieval call binding the contract method 0x9a521382.
//
// Solidity: function isValidWatchtower(address watchtower) view returns(bool)
func (_OperatorRegistry *OperatorRegistrySession) IsValidWatchtower(watchtower common.Address) (bool, error) {
	return _OperatorRegistry.Contract.IsValidWatchtower(&_OperatorRegistry.CallOpts, watchtower)
}

// IsValidWatchtower is a free data retrieval call binding the contract method 0x9a521382.
//
// Solidity: function isValidWatchtower(address watchtower) view returns(bool)
func (_OperatorRegistry *OperatorRegistryCallerSession) IsValidWatchtower(watchtower common.Address) (bool, error) {
	return _OperatorRegistry.Contract.IsValidWatchtower(&_OperatorRegistry.CallOpts, watchtower)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(address operatorAddress, bool isRegistered, uint8 watchStatus, uint8 operatorType, bool isActive)
func (_OperatorRegistry *OperatorRegistryCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (struct {
	OperatorAddress common.Address
	IsRegistered    bool
	WatchStatus     uint8
	OperatorType    uint8
	IsActive        bool
}, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "operators", arg0)

	outstruct := new(struct {
		OperatorAddress common.Address
		IsRegistered    bool
		WatchStatus     uint8
		OperatorType    uint8
		IsActive        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OperatorAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsRegistered = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.WatchStatus = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.OperatorType = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.IsActive = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(address operatorAddress, bool isRegistered, uint8 watchStatus, uint8 operatorType, bool isActive)
func (_OperatorRegistry *OperatorRegistrySession) Operators(arg0 common.Address) (struct {
	OperatorAddress common.Address
	IsRegistered    bool
	WatchStatus     uint8
	OperatorType    uint8
	IsActive        bool
}, error) {
	return _OperatorRegistry.Contract.Operators(&_OperatorRegistry.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(address operatorAddress, bool isRegistered, uint8 watchStatus, uint8 operatorType, bool isActive)
func (_OperatorRegistry *OperatorRegistryCallerSession) Operators(arg0 common.Address) (struct {
	OperatorAddress common.Address
	IsRegistered    bool
	WatchStatus     uint8
	OperatorType    uint8
	IsActive        bool
}, error) {
	return _OperatorRegistry.Contract.Operators(&_OperatorRegistry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OperatorRegistry *OperatorRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OperatorRegistry *OperatorRegistrySession) Owner() (common.Address, error) {
	return _OperatorRegistry.Contract.Owner(&_OperatorRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OperatorRegistry *OperatorRegistryCallerSession) Owner() (common.Address, error) {
	return _OperatorRegistry.Contract.Owner(&_OperatorRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_OperatorRegistry *OperatorRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_OperatorRegistry *OperatorRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _OperatorRegistry.Contract.ProxiableUUID(&_OperatorRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_OperatorRegistry *OperatorRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _OperatorRegistry.Contract.ProxiableUUID(&_OperatorRegistry.CallOpts)
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() view returns(address)
func (_OperatorRegistry *OperatorRegistryCaller) SlasherAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "slasherAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() view returns(address)
func (_OperatorRegistry *OperatorRegistrySession) SlasherAddress() (common.Address, error) {
	return _OperatorRegistry.Contract.SlasherAddress(&_OperatorRegistry.CallOpts)
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() view returns(address)
func (_OperatorRegistry *OperatorRegistryCallerSession) SlasherAddress() (common.Address, error) {
	return _OperatorRegistry.Contract.SlasherAddress(&_OperatorRegistry.CallOpts)
}

// WatchTowerAddressesList is a free data retrieval call binding the contract method 0xf598cd12.
//
// Solidity: function watchTowerAddressesList(uint256 ) view returns(address)
func (_OperatorRegistry *OperatorRegistryCaller) WatchTowerAddressesList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "watchTowerAddressesList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WatchTowerAddressesList is a free data retrieval call binding the contract method 0xf598cd12.
//
// Solidity: function watchTowerAddressesList(uint256 ) view returns(address)
func (_OperatorRegistry *OperatorRegistrySession) WatchTowerAddressesList(arg0 *big.Int) (common.Address, error) {
	return _OperatorRegistry.Contract.WatchTowerAddressesList(&_OperatorRegistry.CallOpts, arg0)
}

// WatchTowerAddressesList is a free data retrieval call binding the contract method 0xf598cd12.
//
// Solidity: function watchTowerAddressesList(uint256 ) view returns(address)
func (_OperatorRegistry *OperatorRegistryCallerSession) WatchTowerAddressesList(arg0 *big.Int) (common.Address, error) {
	return _OperatorRegistry.Contract.WatchTowerAddressesList(&_OperatorRegistry.CallOpts, arg0)
}

// AddToOperatorWhitelist is a paid mutator transaction binding the contract method 0xc8525c3e.
//
// Solidity: function addToOperatorWhitelist(address[] operatorsList) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) AddToOperatorWhitelist(opts *bind.TransactOpts, operatorsList []common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "addToOperatorWhitelist", operatorsList)
}

// AddToOperatorWhitelist is a paid mutator transaction binding the contract method 0xc8525c3e.
//
// Solidity: function addToOperatorWhitelist(address[] operatorsList) returns()
func (_OperatorRegistry *OperatorRegistrySession) AddToOperatorWhitelist(operatorsList []common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.AddToOperatorWhitelist(&_OperatorRegistry.TransactOpts, operatorsList)
}

// AddToOperatorWhitelist is a paid mutator transaction binding the contract method 0xc8525c3e.
//
// Solidity: function addToOperatorWhitelist(address[] operatorsList) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) AddToOperatorWhitelist(operatorsList []common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.AddToOperatorWhitelist(&_OperatorRegistry.TransactOpts, operatorsList)
}

// DeRegister is a paid mutator transaction binding the contract method 0x5b114af6.
//
// Solidity: function deRegister(address operatorAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) DeRegister(opts *bind.TransactOpts, operatorAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "deRegister", operatorAddress)
}

// DeRegister is a paid mutator transaction binding the contract method 0x5b114af6.
//
// Solidity: function deRegister(address operatorAddress) returns()
func (_OperatorRegistry *OperatorRegistrySession) DeRegister(operatorAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.DeRegister(&_OperatorRegistry.TransactOpts, operatorAddress)
}

// DeRegister is a paid mutator transaction binding the contract method 0x5b114af6.
//
// Solidity: function deRegister(address operatorAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) DeRegister(operatorAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.DeRegister(&_OperatorRegistry.TransactOpts, operatorAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _delegationManagerAddress, address _slasherAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) Initialize(opts *bind.TransactOpts, _delegationManagerAddress common.Address, _slasherAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "initialize", _delegationManagerAddress, _slasherAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _delegationManagerAddress, address _slasherAddress) returns()
func (_OperatorRegistry *OperatorRegistrySession) Initialize(_delegationManagerAddress common.Address, _slasherAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.Initialize(&_OperatorRegistry.TransactOpts, _delegationManagerAddress, _slasherAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _delegationManagerAddress, address _slasherAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) Initialize(_delegationManagerAddress common.Address, _slasherAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.Initialize(&_OperatorRegistry.TransactOpts, _delegationManagerAddress, _slasherAddress)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address operatorAddress, address watchtowerAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) Register(opts *bind.TransactOpts, operatorAddress common.Address, watchtowerAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "register", operatorAddress, watchtowerAddress)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address operatorAddress, address watchtowerAddress) returns()
func (_OperatorRegistry *OperatorRegistrySession) Register(operatorAddress common.Address, watchtowerAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.Register(&_OperatorRegistry.TransactOpts, operatorAddress, watchtowerAddress)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address operatorAddress, address watchtowerAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) Register(operatorAddress common.Address, watchtowerAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.Register(&_OperatorRegistry.TransactOpts, operatorAddress, watchtowerAddress)
}

// RemoveFromOperatorWhitelist is a paid mutator transaction binding the contract method 0x9efbea23.
//
// Solidity: function removeFromOperatorWhitelist(address[] operatorsList) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) RemoveFromOperatorWhitelist(opts *bind.TransactOpts, operatorsList []common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "removeFromOperatorWhitelist", operatorsList)
}

// RemoveFromOperatorWhitelist is a paid mutator transaction binding the contract method 0x9efbea23.
//
// Solidity: function removeFromOperatorWhitelist(address[] operatorsList) returns()
func (_OperatorRegistry *OperatorRegistrySession) RemoveFromOperatorWhitelist(operatorsList []common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.RemoveFromOperatorWhitelist(&_OperatorRegistry.TransactOpts, operatorsList)
}

// RemoveFromOperatorWhitelist is a paid mutator transaction binding the contract method 0x9efbea23.
//
// Solidity: function removeFromOperatorWhitelist(address[] operatorsList) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) RemoveFromOperatorWhitelist(operatorsList []common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.RemoveFromOperatorWhitelist(&_OperatorRegistry.TransactOpts, operatorsList)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OperatorRegistry *OperatorRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OperatorRegistry *OperatorRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _OperatorRegistry.Contract.RenounceOwnership(&_OperatorRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OperatorRegistry.Contract.RenounceOwnership(&_OperatorRegistry.TransactOpts)
}

// SetDelegationManagerAddress is a paid mutator transaction binding the contract method 0x7831ad20.
//
// Solidity: function setDelegationManagerAddress(address _delegationManagerAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) SetDelegationManagerAddress(opts *bind.TransactOpts, _delegationManagerAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "setDelegationManagerAddress", _delegationManagerAddress)
}

// SetDelegationManagerAddress is a paid mutator transaction binding the contract method 0x7831ad20.
//
// Solidity: function setDelegationManagerAddress(address _delegationManagerAddress) returns()
func (_OperatorRegistry *OperatorRegistrySession) SetDelegationManagerAddress(_delegationManagerAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.SetDelegationManagerAddress(&_OperatorRegistry.TransactOpts, _delegationManagerAddress)
}

// SetDelegationManagerAddress is a paid mutator transaction binding the contract method 0x7831ad20.
//
// Solidity: function setDelegationManagerAddress(address _delegationManagerAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) SetDelegationManagerAddress(_delegationManagerAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.SetDelegationManagerAddress(&_OperatorRegistry.TransactOpts, _delegationManagerAddress)
}

// SetSlasherAddress is a paid mutator transaction binding the contract method 0xaddd9cf5.
//
// Solidity: function setSlasherAddress(address _slasherAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) SetSlasherAddress(opts *bind.TransactOpts, _slasherAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "setSlasherAddress", _slasherAddress)
}

// SetSlasherAddress is a paid mutator transaction binding the contract method 0xaddd9cf5.
//
// Solidity: function setSlasherAddress(address _slasherAddress) returns()
func (_OperatorRegistry *OperatorRegistrySession) SetSlasherAddress(_slasherAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.SetSlasherAddress(&_OperatorRegistry.TransactOpts, _slasherAddress)
}

// SetSlasherAddress is a paid mutator transaction binding the contract method 0xaddd9cf5.
//
// Solidity: function setSlasherAddress(address _slasherAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) SetSlasherAddress(_slasherAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.SetSlasherAddress(&_OperatorRegistry.TransactOpts, _slasherAddress)
}

// Suspend is a paid mutator transaction binding the contract method 0x286781c7.
//
// Solidity: function suspend(address operatorAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) Suspend(opts *bind.TransactOpts, operatorAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "suspend", operatorAddress)
}

// Suspend is a paid mutator transaction binding the contract method 0x286781c7.
//
// Solidity: function suspend(address operatorAddress) returns()
func (_OperatorRegistry *OperatorRegistrySession) Suspend(operatorAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.Suspend(&_OperatorRegistry.TransactOpts, operatorAddress)
}

// Suspend is a paid mutator transaction binding the contract method 0x286781c7.
//
// Solidity: function suspend(address operatorAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) Suspend(operatorAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.Suspend(&_OperatorRegistry.TransactOpts, operatorAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OperatorRegistry *OperatorRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.TransferOwnership(&_OperatorRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.TransferOwnership(&_OperatorRegistry.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_OperatorRegistry *OperatorRegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.UpgradeTo(&_OperatorRegistry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.UpgradeTo(&_OperatorRegistry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_OperatorRegistry *OperatorRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_OperatorRegistry *OperatorRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.UpgradeToAndCall(&_OperatorRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.UpgradeToAndCall(&_OperatorRegistry.TransactOpts, newImplementation, data)
}

// OperatorRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the OperatorRegistry contract.
type OperatorRegistryAdminChangedIterator struct {
	Event *OperatorRegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryAdminChanged)
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
		it.Event = new(OperatorRegistryAdminChanged)
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
func (it *OperatorRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryAdminChanged represents a AdminChanged event raised by the OperatorRegistry contract.
type OperatorRegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*OperatorRegistryAdminChangedIterator, error) {

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryAdminChangedIterator{contract: _OperatorRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *OperatorRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryAdminChanged)
				if err := _OperatorRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_OperatorRegistry *OperatorRegistryFilterer) ParseAdminChanged(log types.Log) (*OperatorRegistryAdminChanged, error) {
	event := new(OperatorRegistryAdminChanged)
	if err := _OperatorRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the OperatorRegistry contract.
type OperatorRegistryBeaconUpgradedIterator struct {
	Event *OperatorRegistryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryBeaconUpgraded)
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
		it.Event = new(OperatorRegistryBeaconUpgraded)
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
func (it *OperatorRegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryBeaconUpgraded represents a BeaconUpgraded event raised by the OperatorRegistry contract.
type OperatorRegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*OperatorRegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryBeaconUpgradedIterator{contract: _OperatorRegistry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *OperatorRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryBeaconUpgraded)
				if err := _OperatorRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_OperatorRegistry *OperatorRegistryFilterer) ParseBeaconUpgraded(log types.Log) (*OperatorRegistryBeaconUpgraded, error) {
	event := new(OperatorRegistryBeaconUpgraded)
	if err := _OperatorRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OperatorRegistry contract.
type OperatorRegistryInitializedIterator struct {
	Event *OperatorRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryInitialized)
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
		it.Event = new(OperatorRegistryInitialized)
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
func (it *OperatorRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryInitialized represents a Initialized event raised by the OperatorRegistry contract.
type OperatorRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*OperatorRegistryInitializedIterator, error) {

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryInitializedIterator{contract: _OperatorRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OperatorRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryInitialized)
				if err := _OperatorRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OperatorRegistry *OperatorRegistryFilterer) ParseInitialized(log types.Log) (*OperatorRegistryInitialized, error) {
	event := new(OperatorRegistryInitialized)
	if err := _OperatorRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryOperatorDeRegisteredIterator is returned from FilterOperatorDeRegistered and is used to iterate over the raw logs and unpacked data for OperatorDeRegistered events raised by the OperatorRegistry contract.
type OperatorRegistryOperatorDeRegisteredIterator struct {
	Event *OperatorRegistryOperatorDeRegistered // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryOperatorDeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryOperatorDeRegistered)
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
		it.Event = new(OperatorRegistryOperatorDeRegistered)
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
func (it *OperatorRegistryOperatorDeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryOperatorDeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryOperatorDeRegistered represents a OperatorDeRegistered event raised by the OperatorRegistry contract.
type OperatorRegistryOperatorDeRegistered struct {
	Operator IOperatorRegistryOperator
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeRegistered is a free log retrieval operation binding the contract event 0xf17b1a9cbf131df0d08ef9ae44c569f1f58d03d6da953bc1db70aaaea0ef145f.
//
// Solidity: event OperatorDeRegistered((address,bool,uint8,uint8,bool) operator)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterOperatorDeRegistered(opts *bind.FilterOpts) (*OperatorRegistryOperatorDeRegisteredIterator, error) {

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "OperatorDeRegistered")
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryOperatorDeRegisteredIterator{contract: _OperatorRegistry.contract, event: "OperatorDeRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeRegistered is a free log subscription operation binding the contract event 0xf17b1a9cbf131df0d08ef9ae44c569f1f58d03d6da953bc1db70aaaea0ef145f.
//
// Solidity: event OperatorDeRegistered((address,bool,uint8,uint8,bool) operator)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchOperatorDeRegistered(opts *bind.WatchOpts, sink chan<- *OperatorRegistryOperatorDeRegistered) (event.Subscription, error) {

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "OperatorDeRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryOperatorDeRegistered)
				if err := _OperatorRegistry.contract.UnpackLog(event, "OperatorDeRegistered", log); err != nil {
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

// ParseOperatorDeRegistered is a log parse operation binding the contract event 0xf17b1a9cbf131df0d08ef9ae44c569f1f58d03d6da953bc1db70aaaea0ef145f.
//
// Solidity: event OperatorDeRegistered((address,bool,uint8,uint8,bool) operator)
func (_OperatorRegistry *OperatorRegistryFilterer) ParseOperatorDeRegistered(log types.Log) (*OperatorRegistryOperatorDeRegistered, error) {
	event := new(OperatorRegistryOperatorDeRegistered)
	if err := _OperatorRegistry.contract.UnpackLog(event, "OperatorDeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the OperatorRegistry contract.
type OperatorRegistryOperatorRegisteredIterator struct {
	Event *OperatorRegistryOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryOperatorRegistered)
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
		it.Event = new(OperatorRegistryOperatorRegistered)
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
func (it *OperatorRegistryOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryOperatorRegistered represents a OperatorRegistered event raised by the OperatorRegistry contract.
type OperatorRegistryOperatorRegistered struct {
	Operator IOperatorRegistryOperator
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x0a4785e15bc3cda51fd88af2b1b98a6a739074b9960dd51528f299a29d29eb4f.
//
// Solidity: event OperatorRegistered((address,bool,uint8,uint8,bool) operator)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterOperatorRegistered(opts *bind.FilterOpts) (*OperatorRegistryOperatorRegisteredIterator, error) {

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "OperatorRegistered")
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryOperatorRegisteredIterator{contract: _OperatorRegistry.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x0a4785e15bc3cda51fd88af2b1b98a6a739074b9960dd51528f299a29d29eb4f.
//
// Solidity: event OperatorRegistered((address,bool,uint8,uint8,bool) operator)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *OperatorRegistryOperatorRegistered) (event.Subscription, error) {

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "OperatorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryOperatorRegistered)
				if err := _OperatorRegistry.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x0a4785e15bc3cda51fd88af2b1b98a6a739074b9960dd51528f299a29d29eb4f.
//
// Solidity: event OperatorRegistered((address,bool,uint8,uint8,bool) operator)
func (_OperatorRegistry *OperatorRegistryFilterer) ParseOperatorRegistered(log types.Log) (*OperatorRegistryOperatorRegistered, error) {
	event := new(OperatorRegistryOperatorRegistered)
	if err := _OperatorRegistry.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryOperatorSuspendedIterator is returned from FilterOperatorSuspended and is used to iterate over the raw logs and unpacked data for OperatorSuspended events raised by the OperatorRegistry contract.
type OperatorRegistryOperatorSuspendedIterator struct {
	Event *OperatorRegistryOperatorSuspended // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryOperatorSuspendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryOperatorSuspended)
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
		it.Event = new(OperatorRegistryOperatorSuspended)
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
func (it *OperatorRegistryOperatorSuspendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryOperatorSuspendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryOperatorSuspended represents a OperatorSuspended event raised by the OperatorRegistry contract.
type OperatorRegistryOperatorSuspended struct {
	Operator IOperatorRegistryOperator
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSuspended is a free log retrieval operation binding the contract event 0x30168a62be230d0fd5f630a818bf1f2ad87578747c505e13cf1a9382eaa71494.
//
// Solidity: event OperatorSuspended((address,bool,uint8,uint8,bool) operator)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterOperatorSuspended(opts *bind.FilterOpts) (*OperatorRegistryOperatorSuspendedIterator, error) {

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "OperatorSuspended")
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryOperatorSuspendedIterator{contract: _OperatorRegistry.contract, event: "OperatorSuspended", logs: logs, sub: sub}, nil
}

// WatchOperatorSuspended is a free log subscription operation binding the contract event 0x30168a62be230d0fd5f630a818bf1f2ad87578747c505e13cf1a9382eaa71494.
//
// Solidity: event OperatorSuspended((address,bool,uint8,uint8,bool) operator)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchOperatorSuspended(opts *bind.WatchOpts, sink chan<- *OperatorRegistryOperatorSuspended) (event.Subscription, error) {

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "OperatorSuspended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryOperatorSuspended)
				if err := _OperatorRegistry.contract.UnpackLog(event, "OperatorSuspended", log); err != nil {
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

// ParseOperatorSuspended is a log parse operation binding the contract event 0x30168a62be230d0fd5f630a818bf1f2ad87578747c505e13cf1a9382eaa71494.
//
// Solidity: event OperatorSuspended((address,bool,uint8,uint8,bool) operator)
func (_OperatorRegistry *OperatorRegistryFilterer) ParseOperatorSuspended(log types.Log) (*OperatorRegistryOperatorSuspended, error) {
	event := new(OperatorRegistryOperatorSuspended)
	if err := _OperatorRegistry.contract.UnpackLog(event, "OperatorSuspended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OperatorRegistry contract.
type OperatorRegistryOwnershipTransferredIterator struct {
	Event *OperatorRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryOwnershipTransferred)
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
		it.Event = new(OperatorRegistryOwnershipTransferred)
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
func (it *OperatorRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the OperatorRegistry contract.
type OperatorRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OperatorRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryOwnershipTransferredIterator{contract: _OperatorRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OperatorRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryOwnershipTransferred)
				if err := _OperatorRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OperatorRegistry *OperatorRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*OperatorRegistryOwnershipTransferred, error) {
	event := new(OperatorRegistryOwnershipTransferred)
	if err := _OperatorRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the OperatorRegistry contract.
type OperatorRegistryUpgradedIterator struct {
	Event *OperatorRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryUpgraded)
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
		it.Event = new(OperatorRegistryUpgraded)
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
func (it *OperatorRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryUpgraded represents a Upgraded event raised by the OperatorRegistry contract.
type OperatorRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*OperatorRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryUpgradedIterator{contract: _OperatorRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *OperatorRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryUpgraded)
				if err := _OperatorRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_OperatorRegistry *OperatorRegistryFilterer) ParseUpgraded(log types.Log) (*OperatorRegistryUpgraded, error) {
	event := new(OperatorRegistryUpgraded)
	if err := _OperatorRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
