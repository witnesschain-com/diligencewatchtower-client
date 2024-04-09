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

// TypesRewardPoints is an auto generated low-level Go binding around an user-defined struct.
type TypesRewardPoints struct {
	InclusionProofPoints *big.Int
	DiligenceProofPoints *big.Int
}

// EigenTowerMetaData contains all meta data concerning the EigenTower contract.
var EigenTowerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggregator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggregator\",\"type\":\"address\"}],\"name\":\"AggregatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"InvalidOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldL2ChainMapping\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newL2ChainMapping\",\"type\":\"address\"}],\"name\":\"L2ChainMappingUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2BlockNumberEnd\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rewardHash\",\"type\":\"bytes32\"}],\"name\":\"NewRewardsUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRegistry\",\"type\":\"address\"}],\"name\":\"RegistryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getOperatorRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"inclusionProofPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"diligenceProofPoints\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.RewardPoints\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOperatorRegistry\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"contractIL2ChainMapping\",\"name\":\"_l2ChainMapping\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_agg\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ChainMapping\",\"outputs\":[{\"internalType\":\"contractIL2ChainMapping\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"inclusionProofPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"diligenceProofPoints\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIOperatorRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"setAggregatorAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIL2ChainMapping\",\"name\":\"_l2chainmapping\",\"type\":\"address\"}],\"name\":\"setL2ChainMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOperatorRegistry\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumBegin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumEnd\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_operatorsList\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"inclusionProofPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"diligenceProofPoints\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.RewardPoints[]\",\"name\":\"_proofRewards\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"_rewardHash\",\"type\":\"bytes32\"}],\"name\":\"updateReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// EigenTowerABI is the input ABI used to generate the binding from.
// Deprecated: Use EigenTowerMetaData.ABI instead.
var EigenTowerABI = EigenTowerMetaData.ABI

// EigenTower is an auto generated Go binding around an Ethereum contract.
type EigenTower struct {
	EigenTowerCaller     // Read-only binding to the contract
	EigenTowerTransactor // Write-only binding to the contract
	EigenTowerFilterer   // Log filterer for contract events
}

// EigenTowerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EigenTowerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenTowerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EigenTowerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenTowerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EigenTowerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenTowerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EigenTowerSession struct {
	Contract     *EigenTower       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EigenTowerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EigenTowerCallerSession struct {
	Contract *EigenTowerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EigenTowerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EigenTowerTransactorSession struct {
	Contract     *EigenTowerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EigenTowerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EigenTowerRaw struct {
	Contract *EigenTower // Generic contract binding to access the raw methods on
}

// EigenTowerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EigenTowerCallerRaw struct {
	Contract *EigenTowerCaller // Generic read-only contract binding to access the raw methods on
}

// EigenTowerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EigenTowerTransactorRaw struct {
	Contract *EigenTowerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEigenTower creates a new instance of EigenTower, bound to a specific deployed contract.
func NewEigenTower(address common.Address, backend bind.ContractBackend) (*EigenTower, error) {
	contract, err := bindEigenTower(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EigenTower{EigenTowerCaller: EigenTowerCaller{contract: contract}, EigenTowerTransactor: EigenTowerTransactor{contract: contract}, EigenTowerFilterer: EigenTowerFilterer{contract: contract}}, nil
}

// NewEigenTowerCaller creates a new read-only instance of EigenTower, bound to a specific deployed contract.
func NewEigenTowerCaller(address common.Address, caller bind.ContractCaller) (*EigenTowerCaller, error) {
	contract, err := bindEigenTower(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EigenTowerCaller{contract: contract}, nil
}

// NewEigenTowerTransactor creates a new write-only instance of EigenTower, bound to a specific deployed contract.
func NewEigenTowerTransactor(address common.Address, transactor bind.ContractTransactor) (*EigenTowerTransactor, error) {
	contract, err := bindEigenTower(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EigenTowerTransactor{contract: contract}, nil
}

// NewEigenTowerFilterer creates a new log filterer instance of EigenTower, bound to a specific deployed contract.
func NewEigenTowerFilterer(address common.Address, filterer bind.ContractFilterer) (*EigenTowerFilterer, error) {
	contract, err := bindEigenTower(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EigenTowerFilterer{contract: contract}, nil
}

// bindEigenTower binds a generic wrapper to an already deployed contract.
func bindEigenTower(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EigenTowerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EigenTower *EigenTowerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EigenTower.Contract.EigenTowerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EigenTower *EigenTowerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenTower.Contract.EigenTowerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EigenTower *EigenTowerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EigenTower.Contract.EigenTowerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EigenTower *EigenTowerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EigenTower.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EigenTower *EigenTowerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenTower.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EigenTower *EigenTowerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EigenTower.Contract.contract.Transact(opts, method, params...)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_EigenTower *EigenTowerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenTower.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_EigenTower *EigenTowerSession) Aggregator() (common.Address, error) {
	return _EigenTower.Contract.Aggregator(&_EigenTower.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_EigenTower *EigenTowerCallerSession) Aggregator() (common.Address, error) {
	return _EigenTower.Contract.Aggregator(&_EigenTower.CallOpts)
}

// GetOperatorRewards is a free data retrieval call binding the contract method 0xc9045a30.
//
// Solidity: function getOperatorRewards(address _operator) view returns((uint256,uint256))
func (_EigenTower *EigenTowerCaller) GetOperatorRewards(opts *bind.CallOpts, _operator common.Address) (TypesRewardPoints, error) {
	var out []interface{}
	err := _EigenTower.contract.Call(opts, &out, "getOperatorRewards", _operator)

	if err != nil {
		return *new(TypesRewardPoints), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesRewardPoints)).(*TypesRewardPoints)

	return out0, err

}

// GetOperatorRewards is a free data retrieval call binding the contract method 0xc9045a30.
//
// Solidity: function getOperatorRewards(address _operator) view returns((uint256,uint256))
func (_EigenTower *EigenTowerSession) GetOperatorRewards(_operator common.Address) (TypesRewardPoints, error) {
	return _EigenTower.Contract.GetOperatorRewards(&_EigenTower.CallOpts, _operator)
}

// GetOperatorRewards is a free data retrieval call binding the contract method 0xc9045a30.
//
// Solidity: function getOperatorRewards(address _operator) view returns((uint256,uint256))
func (_EigenTower *EigenTowerCallerSession) GetOperatorRewards(_operator common.Address) (TypesRewardPoints, error) {
	return _EigenTower.Contract.GetOperatorRewards(&_EigenTower.CallOpts, _operator)
}

// L2ChainMapping is a free data retrieval call binding the contract method 0x820b6f39.
//
// Solidity: function l2ChainMapping() view returns(address)
func (_EigenTower *EigenTowerCaller) L2ChainMapping(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenTower.contract.Call(opts, &out, "l2ChainMapping")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ChainMapping is a free data retrieval call binding the contract method 0x820b6f39.
//
// Solidity: function l2ChainMapping() view returns(address)
func (_EigenTower *EigenTowerSession) L2ChainMapping() (common.Address, error) {
	return _EigenTower.Contract.L2ChainMapping(&_EigenTower.CallOpts)
}

// L2ChainMapping is a free data retrieval call binding the contract method 0x820b6f39.
//
// Solidity: function l2ChainMapping() view returns(address)
func (_EigenTower *EigenTowerCallerSession) L2ChainMapping() (common.Address, error) {
	return _EigenTower.Contract.L2ChainMapping(&_EigenTower.CallOpts)
}

// OperatorRewards is a free data retrieval call binding the contract method 0x41a2b8d6.
//
// Solidity: function operatorRewards(address ) view returns(uint256 inclusionProofPoints, uint256 diligenceProofPoints)
func (_EigenTower *EigenTowerCaller) OperatorRewards(opts *bind.CallOpts, arg0 common.Address) (struct {
	InclusionProofPoints *big.Int
	DiligenceProofPoints *big.Int
}, error) {
	var out []interface{}
	err := _EigenTower.contract.Call(opts, &out, "operatorRewards", arg0)

	outstruct := new(struct {
		InclusionProofPoints *big.Int
		DiligenceProofPoints *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InclusionProofPoints = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DiligenceProofPoints = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OperatorRewards is a free data retrieval call binding the contract method 0x41a2b8d6.
//
// Solidity: function operatorRewards(address ) view returns(uint256 inclusionProofPoints, uint256 diligenceProofPoints)
func (_EigenTower *EigenTowerSession) OperatorRewards(arg0 common.Address) (struct {
	InclusionProofPoints *big.Int
	DiligenceProofPoints *big.Int
}, error) {
	return _EigenTower.Contract.OperatorRewards(&_EigenTower.CallOpts, arg0)
}

// OperatorRewards is a free data retrieval call binding the contract method 0x41a2b8d6.
//
// Solidity: function operatorRewards(address ) view returns(uint256 inclusionProofPoints, uint256 diligenceProofPoints)
func (_EigenTower *EigenTowerCallerSession) OperatorRewards(arg0 common.Address) (struct {
	InclusionProofPoints *big.Int
	DiligenceProofPoints *big.Int
}, error) {
	return _EigenTower.Contract.OperatorRewards(&_EigenTower.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EigenTower *EigenTowerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenTower.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EigenTower *EigenTowerSession) Owner() (common.Address, error) {
	return _EigenTower.Contract.Owner(&_EigenTower.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EigenTower *EigenTowerCallerSession) Owner() (common.Address, error) {
	return _EigenTower.Contract.Owner(&_EigenTower.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EigenTower *EigenTowerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EigenTower.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EigenTower *EigenTowerSession) Paused() (bool, error) {
	return _EigenTower.Contract.Paused(&_EigenTower.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EigenTower *EigenTowerCallerSession) Paused() (bool, error) {
	return _EigenTower.Contract.Paused(&_EigenTower.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EigenTower *EigenTowerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EigenTower.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EigenTower *EigenTowerSession) ProxiableUUID() ([32]byte, error) {
	return _EigenTower.Contract.ProxiableUUID(&_EigenTower.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EigenTower *EigenTowerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _EigenTower.Contract.ProxiableUUID(&_EigenTower.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_EigenTower *EigenTowerCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenTower.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_EigenTower *EigenTowerSession) Registry() (common.Address, error) {
	return _EigenTower.Contract.Registry(&_EigenTower.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_EigenTower *EigenTowerCallerSession) Registry() (common.Address, error) {
	return _EigenTower.Contract.Registry(&_EigenTower.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _registry, address _l2ChainMapping, address _agg) returns()
func (_EigenTower *EigenTowerTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _l2ChainMapping common.Address, _agg common.Address) (*types.Transaction, error) {
	return _EigenTower.contract.Transact(opts, "initialize", _registry, _l2ChainMapping, _agg)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _registry, address _l2ChainMapping, address _agg) returns()
func (_EigenTower *EigenTowerSession) Initialize(_registry common.Address, _l2ChainMapping common.Address, _agg common.Address) (*types.Transaction, error) {
	return _EigenTower.Contract.Initialize(&_EigenTower.TransactOpts, _registry, _l2ChainMapping, _agg)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _registry, address _l2ChainMapping, address _agg) returns()
func (_EigenTower *EigenTowerTransactorSession) Initialize(_registry common.Address, _l2ChainMapping common.Address, _agg common.Address) (*types.Transaction, error) {
	return _EigenTower.Contract.Initialize(&_EigenTower.TransactOpts, _registry, _l2ChainMapping, _agg)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EigenTower *EigenTowerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenTower.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EigenTower *EigenTowerSession) Pause() (*types.Transaction, error) {
	return _EigenTower.Contract.Pause(&_EigenTower.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EigenTower *EigenTowerTransactorSession) Pause() (*types.Transaction, error) {
	return _EigenTower.Contract.Pause(&_EigenTower.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EigenTower *EigenTowerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenTower.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EigenTower *EigenTowerSession) RenounceOwnership() (*types.Transaction, error) {
	return _EigenTower.Contract.RenounceOwnership(&_EigenTower.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EigenTower *EigenTowerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EigenTower.Contract.RenounceOwnership(&_EigenTower.TransactOpts)
}

// SetAggregatorAddress is a paid mutator transaction binding the contract method 0x47b32448.
//
// Solidity: function setAggregatorAddress(address _aggregator) returns()
func (_EigenTower *EigenTowerTransactor) SetAggregatorAddress(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _EigenTower.contract.Transact(opts, "setAggregatorAddress", _aggregator)
}

// SetAggregatorAddress is a paid mutator transaction binding the contract method 0x47b32448.
//
// Solidity: function setAggregatorAddress(address _aggregator) returns()
func (_EigenTower *EigenTowerSession) SetAggregatorAddress(_aggregator common.Address) (*types.Transaction, error) {
	return _EigenTower.Contract.SetAggregatorAddress(&_EigenTower.TransactOpts, _aggregator)
}

// SetAggregatorAddress is a paid mutator transaction binding the contract method 0x47b32448.
//
// Solidity: function setAggregatorAddress(address _aggregator) returns()
func (_EigenTower *EigenTowerTransactorSession) SetAggregatorAddress(_aggregator common.Address) (*types.Transaction, error) {
	return _EigenTower.Contract.SetAggregatorAddress(&_EigenTower.TransactOpts, _aggregator)
}

// SetL2ChainMapping is a paid mutator transaction binding the contract method 0x8d134f8d.
//
// Solidity: function setL2ChainMapping(address _l2chainmapping) returns()
func (_EigenTower *EigenTowerTransactor) SetL2ChainMapping(opts *bind.TransactOpts, _l2chainmapping common.Address) (*types.Transaction, error) {
	return _EigenTower.contract.Transact(opts, "setL2ChainMapping", _l2chainmapping)
}

// SetL2ChainMapping is a paid mutator transaction binding the contract method 0x8d134f8d.
//
// Solidity: function setL2ChainMapping(address _l2chainmapping) returns()
func (_EigenTower *EigenTowerSession) SetL2ChainMapping(_l2chainmapping common.Address) (*types.Transaction, error) {
	return _EigenTower.Contract.SetL2ChainMapping(&_EigenTower.TransactOpts, _l2chainmapping)
}

// SetL2ChainMapping is a paid mutator transaction binding the contract method 0x8d134f8d.
//
// Solidity: function setL2ChainMapping(address _l2chainmapping) returns()
func (_EigenTower *EigenTowerTransactorSession) SetL2ChainMapping(_l2chainmapping common.Address) (*types.Transaction, error) {
	return _EigenTower.Contract.SetL2ChainMapping(&_EigenTower.TransactOpts, _l2chainmapping)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_EigenTower *EigenTowerTransactor) SetRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _EigenTower.contract.Transact(opts, "setRegistry", _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_EigenTower *EigenTowerSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _EigenTower.Contract.SetRegistry(&_EigenTower.TransactOpts, _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_EigenTower *EigenTowerTransactorSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _EigenTower.Contract.SetRegistry(&_EigenTower.TransactOpts, _registry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EigenTower *EigenTowerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EigenTower.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EigenTower *EigenTowerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EigenTower.Contract.TransferOwnership(&_EigenTower.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EigenTower *EigenTowerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EigenTower.Contract.TransferOwnership(&_EigenTower.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EigenTower *EigenTowerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenTower.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EigenTower *EigenTowerSession) Unpause() (*types.Transaction, error) {
	return _EigenTower.Contract.Unpause(&_EigenTower.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EigenTower *EigenTowerTransactorSession) Unpause() (*types.Transaction, error) {
	return _EigenTower.Contract.Unpause(&_EigenTower.TransactOpts)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x395373b0.
//
// Solidity: function updateReward(uint256 _chainID, uint256 _blockNumBegin, uint256 _blockNumEnd, address[] _operatorsList, (uint256,uint256)[] _proofRewards, bytes32 _rewardHash) returns()
func (_EigenTower *EigenTowerTransactor) UpdateReward(opts *bind.TransactOpts, _chainID *big.Int, _blockNumBegin *big.Int, _blockNumEnd *big.Int, _operatorsList []common.Address, _proofRewards []TypesRewardPoints, _rewardHash [32]byte) (*types.Transaction, error) {
	return _EigenTower.contract.Transact(opts, "updateReward", _chainID, _blockNumBegin, _blockNumEnd, _operatorsList, _proofRewards, _rewardHash)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x395373b0.
//
// Solidity: function updateReward(uint256 _chainID, uint256 _blockNumBegin, uint256 _blockNumEnd, address[] _operatorsList, (uint256,uint256)[] _proofRewards, bytes32 _rewardHash) returns()
func (_EigenTower *EigenTowerSession) UpdateReward(_chainID *big.Int, _blockNumBegin *big.Int, _blockNumEnd *big.Int, _operatorsList []common.Address, _proofRewards []TypesRewardPoints, _rewardHash [32]byte) (*types.Transaction, error) {
	return _EigenTower.Contract.UpdateReward(&_EigenTower.TransactOpts, _chainID, _blockNumBegin, _blockNumEnd, _operatorsList, _proofRewards, _rewardHash)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x395373b0.
//
// Solidity: function updateReward(uint256 _chainID, uint256 _blockNumBegin, uint256 _blockNumEnd, address[] _operatorsList, (uint256,uint256)[] _proofRewards, bytes32 _rewardHash) returns()
func (_EigenTower *EigenTowerTransactorSession) UpdateReward(_chainID *big.Int, _blockNumBegin *big.Int, _blockNumEnd *big.Int, _operatorsList []common.Address, _proofRewards []TypesRewardPoints, _rewardHash [32]byte) (*types.Transaction, error) {
	return _EigenTower.Contract.UpdateReward(&_EigenTower.TransactOpts, _chainID, _blockNumBegin, _blockNumEnd, _operatorsList, _proofRewards, _rewardHash)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EigenTower *EigenTowerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _EigenTower.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EigenTower *EigenTowerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EigenTower.Contract.UpgradeTo(&_EigenTower.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EigenTower *EigenTowerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EigenTower.Contract.UpgradeTo(&_EigenTower.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EigenTower *EigenTowerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EigenTower.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EigenTower *EigenTowerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EigenTower.Contract.UpgradeToAndCall(&_EigenTower.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EigenTower *EigenTowerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EigenTower.Contract.UpgradeToAndCall(&_EigenTower.TransactOpts, newImplementation, data)
}

// EigenTowerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the EigenTower contract.
type EigenTowerAdminChangedIterator struct {
	Event *EigenTowerAdminChanged // Event containing the contract specifics and raw log

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
func (it *EigenTowerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenTowerAdminChanged)
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
		it.Event = new(EigenTowerAdminChanged)
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
func (it *EigenTowerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenTowerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenTowerAdminChanged represents a AdminChanged event raised by the EigenTower contract.
type EigenTowerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EigenTower *EigenTowerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*EigenTowerAdminChangedIterator, error) {

	logs, sub, err := _EigenTower.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &EigenTowerAdminChangedIterator{contract: _EigenTower.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EigenTower *EigenTowerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *EigenTowerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _EigenTower.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenTowerAdminChanged)
				if err := _EigenTower.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_EigenTower *EigenTowerFilterer) ParseAdminChanged(log types.Log) (*EigenTowerAdminChanged, error) {
	event := new(EigenTowerAdminChanged)
	if err := _EigenTower.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenTowerAggregatorUpdatedIterator is returned from FilterAggregatorUpdated and is used to iterate over the raw logs and unpacked data for AggregatorUpdated events raised by the EigenTower contract.
type EigenTowerAggregatorUpdatedIterator struct {
	Event *EigenTowerAggregatorUpdated // Event containing the contract specifics and raw log

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
func (it *EigenTowerAggregatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenTowerAggregatorUpdated)
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
		it.Event = new(EigenTowerAggregatorUpdated)
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
func (it *EigenTowerAggregatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenTowerAggregatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenTowerAggregatorUpdated represents a AggregatorUpdated event raised by the EigenTower contract.
type EigenTowerAggregatorUpdated struct {
	OldAggregator common.Address
	NewAggregator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAggregatorUpdated is a free log retrieval operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address oldAggregator, address newAggregator)
func (_EigenTower *EigenTowerFilterer) FilterAggregatorUpdated(opts *bind.FilterOpts) (*EigenTowerAggregatorUpdatedIterator, error) {

	logs, sub, err := _EigenTower.contract.FilterLogs(opts, "AggregatorUpdated")
	if err != nil {
		return nil, err
	}
	return &EigenTowerAggregatorUpdatedIterator{contract: _EigenTower.contract, event: "AggregatorUpdated", logs: logs, sub: sub}, nil
}

// WatchAggregatorUpdated is a free log subscription operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address oldAggregator, address newAggregator)
func (_EigenTower *EigenTowerFilterer) WatchAggregatorUpdated(opts *bind.WatchOpts, sink chan<- *EigenTowerAggregatorUpdated) (event.Subscription, error) {

	logs, sub, err := _EigenTower.contract.WatchLogs(opts, "AggregatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenTowerAggregatorUpdated)
				if err := _EigenTower.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
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

// ParseAggregatorUpdated is a log parse operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address oldAggregator, address newAggregator)
func (_EigenTower *EigenTowerFilterer) ParseAggregatorUpdated(log types.Log) (*EigenTowerAggregatorUpdated, error) {
	event := new(EigenTowerAggregatorUpdated)
	if err := _EigenTower.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenTowerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the EigenTower contract.
type EigenTowerBeaconUpgradedIterator struct {
	Event *EigenTowerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *EigenTowerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenTowerBeaconUpgraded)
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
		it.Event = new(EigenTowerBeaconUpgraded)
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
func (it *EigenTowerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenTowerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenTowerBeaconUpgraded represents a BeaconUpgraded event raised by the EigenTower contract.
type EigenTowerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EigenTower *EigenTowerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*EigenTowerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EigenTower.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &EigenTowerBeaconUpgradedIterator{contract: _EigenTower.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EigenTower *EigenTowerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *EigenTowerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EigenTower.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenTowerBeaconUpgraded)
				if err := _EigenTower.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_EigenTower *EigenTowerFilterer) ParseBeaconUpgraded(log types.Log) (*EigenTowerBeaconUpgraded, error) {
	event := new(EigenTowerBeaconUpgraded)
	if err := _EigenTower.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenTowerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EigenTower contract.
type EigenTowerInitializedIterator struct {
	Event *EigenTowerInitialized // Event containing the contract specifics and raw log

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
func (it *EigenTowerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenTowerInitialized)
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
		it.Event = new(EigenTowerInitialized)
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
func (it *EigenTowerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenTowerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenTowerInitialized represents a Initialized event raised by the EigenTower contract.
type EigenTowerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EigenTower *EigenTowerFilterer) FilterInitialized(opts *bind.FilterOpts) (*EigenTowerInitializedIterator, error) {

	logs, sub, err := _EigenTower.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EigenTowerInitializedIterator{contract: _EigenTower.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EigenTower *EigenTowerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EigenTowerInitialized) (event.Subscription, error) {

	logs, sub, err := _EigenTower.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenTowerInitialized)
				if err := _EigenTower.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EigenTower *EigenTowerFilterer) ParseInitialized(log types.Log) (*EigenTowerInitialized, error) {
	event := new(EigenTowerInitialized)
	if err := _EigenTower.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenTowerInvalidOperatorIterator is returned from FilterInvalidOperator and is used to iterate over the raw logs and unpacked data for InvalidOperator events raised by the EigenTower contract.
type EigenTowerInvalidOperatorIterator struct {
	Event *EigenTowerInvalidOperator // Event containing the contract specifics and raw log

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
func (it *EigenTowerInvalidOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenTowerInvalidOperator)
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
		it.Event = new(EigenTowerInvalidOperator)
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
func (it *EigenTowerInvalidOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenTowerInvalidOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenTowerInvalidOperator represents a InvalidOperator event raised by the EigenTower contract.
type EigenTowerInvalidOperator struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvalidOperator is a free log retrieval operation binding the contract event 0x3eaa03e38745f9ef282a3d3ae730e428a187a9075ccc2a3b7a1a318fe83d25be.
//
// Solidity: event InvalidOperator(address operator)
func (_EigenTower *EigenTowerFilterer) FilterInvalidOperator(opts *bind.FilterOpts) (*EigenTowerInvalidOperatorIterator, error) {

	logs, sub, err := _EigenTower.contract.FilterLogs(opts, "InvalidOperator")
	if err != nil {
		return nil, err
	}
	return &EigenTowerInvalidOperatorIterator{contract: _EigenTower.contract, event: "InvalidOperator", logs: logs, sub: sub}, nil
}

// WatchInvalidOperator is a free log subscription operation binding the contract event 0x3eaa03e38745f9ef282a3d3ae730e428a187a9075ccc2a3b7a1a318fe83d25be.
//
// Solidity: event InvalidOperator(address operator)
func (_EigenTower *EigenTowerFilterer) WatchInvalidOperator(opts *bind.WatchOpts, sink chan<- *EigenTowerInvalidOperator) (event.Subscription, error) {

	logs, sub, err := _EigenTower.contract.WatchLogs(opts, "InvalidOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenTowerInvalidOperator)
				if err := _EigenTower.contract.UnpackLog(event, "InvalidOperator", log); err != nil {
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

// ParseInvalidOperator is a log parse operation binding the contract event 0x3eaa03e38745f9ef282a3d3ae730e428a187a9075ccc2a3b7a1a318fe83d25be.
//
// Solidity: event InvalidOperator(address operator)
func (_EigenTower *EigenTowerFilterer) ParseInvalidOperator(log types.Log) (*EigenTowerInvalidOperator, error) {
	event := new(EigenTowerInvalidOperator)
	if err := _EigenTower.contract.UnpackLog(event, "InvalidOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenTowerL2ChainMappingUpdatedIterator is returned from FilterL2ChainMappingUpdated and is used to iterate over the raw logs and unpacked data for L2ChainMappingUpdated events raised by the EigenTower contract.
type EigenTowerL2ChainMappingUpdatedIterator struct {
	Event *EigenTowerL2ChainMappingUpdated // Event containing the contract specifics and raw log

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
func (it *EigenTowerL2ChainMappingUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenTowerL2ChainMappingUpdated)
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
		it.Event = new(EigenTowerL2ChainMappingUpdated)
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
func (it *EigenTowerL2ChainMappingUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenTowerL2ChainMappingUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenTowerL2ChainMappingUpdated represents a L2ChainMappingUpdated event raised by the EigenTower contract.
type EigenTowerL2ChainMappingUpdated struct {
	OldL2ChainMapping common.Address
	NewL2ChainMapping common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterL2ChainMappingUpdated is a free log retrieval operation binding the contract event 0xafcfe8816011d6bd68dd7f59cbd0865f721cc58085f72778b69065c75a19b5ca.
//
// Solidity: event L2ChainMappingUpdated(address oldL2ChainMapping, address newL2ChainMapping)
func (_EigenTower *EigenTowerFilterer) FilterL2ChainMappingUpdated(opts *bind.FilterOpts) (*EigenTowerL2ChainMappingUpdatedIterator, error) {

	logs, sub, err := _EigenTower.contract.FilterLogs(opts, "L2ChainMappingUpdated")
	if err != nil {
		return nil, err
	}
	return &EigenTowerL2ChainMappingUpdatedIterator{contract: _EigenTower.contract, event: "L2ChainMappingUpdated", logs: logs, sub: sub}, nil
}

// WatchL2ChainMappingUpdated is a free log subscription operation binding the contract event 0xafcfe8816011d6bd68dd7f59cbd0865f721cc58085f72778b69065c75a19b5ca.
//
// Solidity: event L2ChainMappingUpdated(address oldL2ChainMapping, address newL2ChainMapping)
func (_EigenTower *EigenTowerFilterer) WatchL2ChainMappingUpdated(opts *bind.WatchOpts, sink chan<- *EigenTowerL2ChainMappingUpdated) (event.Subscription, error) {

	logs, sub, err := _EigenTower.contract.WatchLogs(opts, "L2ChainMappingUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenTowerL2ChainMappingUpdated)
				if err := _EigenTower.contract.UnpackLog(event, "L2ChainMappingUpdated", log); err != nil {
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

// ParseL2ChainMappingUpdated is a log parse operation binding the contract event 0xafcfe8816011d6bd68dd7f59cbd0865f721cc58085f72778b69065c75a19b5ca.
//
// Solidity: event L2ChainMappingUpdated(address oldL2ChainMapping, address newL2ChainMapping)
func (_EigenTower *EigenTowerFilterer) ParseL2ChainMappingUpdated(log types.Log) (*EigenTowerL2ChainMappingUpdated, error) {
	event := new(EigenTowerL2ChainMappingUpdated)
	if err := _EigenTower.contract.UnpackLog(event, "L2ChainMappingUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenTowerNewRewardsUpdateIterator is returned from FilterNewRewardsUpdate and is used to iterate over the raw logs and unpacked data for NewRewardsUpdate events raised by the EigenTower contract.
type EigenTowerNewRewardsUpdateIterator struct {
	Event *EigenTowerNewRewardsUpdate // Event containing the contract specifics and raw log

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
func (it *EigenTowerNewRewardsUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenTowerNewRewardsUpdate)
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
		it.Event = new(EigenTowerNewRewardsUpdate)
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
func (it *EigenTowerNewRewardsUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenTowerNewRewardsUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenTowerNewRewardsUpdate represents a NewRewardsUpdate event raised by the EigenTower contract.
type EigenTowerNewRewardsUpdate struct {
	ChainId          *big.Int
	L2BlockNumberEnd *big.Int
	RewardHash       [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewRewardsUpdate is a free log retrieval operation binding the contract event 0x36b8234ba3e9f8897b5448ed687509b776a4ad94a88ffd28d6aef8bee37ee610.
//
// Solidity: event NewRewardsUpdate(uint256 indexed chainId, uint256 indexed l2BlockNumberEnd, bytes32 indexed rewardHash)
func (_EigenTower *EigenTowerFilterer) FilterNewRewardsUpdate(opts *bind.FilterOpts, chainId []*big.Int, l2BlockNumberEnd []*big.Int, rewardHash [][32]byte) (*EigenTowerNewRewardsUpdateIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var l2BlockNumberEndRule []interface{}
	for _, l2BlockNumberEndItem := range l2BlockNumberEnd {
		l2BlockNumberEndRule = append(l2BlockNumberEndRule, l2BlockNumberEndItem)
	}
	var rewardHashRule []interface{}
	for _, rewardHashItem := range rewardHash {
		rewardHashRule = append(rewardHashRule, rewardHashItem)
	}

	logs, sub, err := _EigenTower.contract.FilterLogs(opts, "NewRewardsUpdate", chainIdRule, l2BlockNumberEndRule, rewardHashRule)
	if err != nil {
		return nil, err
	}
	return &EigenTowerNewRewardsUpdateIterator{contract: _EigenTower.contract, event: "NewRewardsUpdate", logs: logs, sub: sub}, nil
}

// WatchNewRewardsUpdate is a free log subscription operation binding the contract event 0x36b8234ba3e9f8897b5448ed687509b776a4ad94a88ffd28d6aef8bee37ee610.
//
// Solidity: event NewRewardsUpdate(uint256 indexed chainId, uint256 indexed l2BlockNumberEnd, bytes32 indexed rewardHash)
func (_EigenTower *EigenTowerFilterer) WatchNewRewardsUpdate(opts *bind.WatchOpts, sink chan<- *EigenTowerNewRewardsUpdate, chainId []*big.Int, l2BlockNumberEnd []*big.Int, rewardHash [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var l2BlockNumberEndRule []interface{}
	for _, l2BlockNumberEndItem := range l2BlockNumberEnd {
		l2BlockNumberEndRule = append(l2BlockNumberEndRule, l2BlockNumberEndItem)
	}
	var rewardHashRule []interface{}
	for _, rewardHashItem := range rewardHash {
		rewardHashRule = append(rewardHashRule, rewardHashItem)
	}

	logs, sub, err := _EigenTower.contract.WatchLogs(opts, "NewRewardsUpdate", chainIdRule, l2BlockNumberEndRule, rewardHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenTowerNewRewardsUpdate)
				if err := _EigenTower.contract.UnpackLog(event, "NewRewardsUpdate", log); err != nil {
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

// ParseNewRewardsUpdate is a log parse operation binding the contract event 0x36b8234ba3e9f8897b5448ed687509b776a4ad94a88ffd28d6aef8bee37ee610.
//
// Solidity: event NewRewardsUpdate(uint256 indexed chainId, uint256 indexed l2BlockNumberEnd, bytes32 indexed rewardHash)
func (_EigenTower *EigenTowerFilterer) ParseNewRewardsUpdate(log types.Log) (*EigenTowerNewRewardsUpdate, error) {
	event := new(EigenTowerNewRewardsUpdate)
	if err := _EigenTower.contract.UnpackLog(event, "NewRewardsUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenTowerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EigenTower contract.
type EigenTowerOwnershipTransferredIterator struct {
	Event *EigenTowerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EigenTowerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenTowerOwnershipTransferred)
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
		it.Event = new(EigenTowerOwnershipTransferred)
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
func (it *EigenTowerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenTowerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenTowerOwnershipTransferred represents a OwnershipTransferred event raised by the EigenTower contract.
type EigenTowerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EigenTower *EigenTowerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EigenTowerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EigenTower.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EigenTowerOwnershipTransferredIterator{contract: _EigenTower.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EigenTower *EigenTowerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EigenTowerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EigenTower.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenTowerOwnershipTransferred)
				if err := _EigenTower.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EigenTower *EigenTowerFilterer) ParseOwnershipTransferred(log types.Log) (*EigenTowerOwnershipTransferred, error) {
	event := new(EigenTowerOwnershipTransferred)
	if err := _EigenTower.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenTowerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EigenTower contract.
type EigenTowerPausedIterator struct {
	Event *EigenTowerPaused // Event containing the contract specifics and raw log

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
func (it *EigenTowerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenTowerPaused)
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
		it.Event = new(EigenTowerPaused)
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
func (it *EigenTowerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenTowerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenTowerPaused represents a Paused event raised by the EigenTower contract.
type EigenTowerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EigenTower *EigenTowerFilterer) FilterPaused(opts *bind.FilterOpts) (*EigenTowerPausedIterator, error) {

	logs, sub, err := _EigenTower.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EigenTowerPausedIterator{contract: _EigenTower.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EigenTower *EigenTowerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EigenTowerPaused) (event.Subscription, error) {

	logs, sub, err := _EigenTower.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenTowerPaused)
				if err := _EigenTower.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_EigenTower *EigenTowerFilterer) ParsePaused(log types.Log) (*EigenTowerPaused, error) {
	event := new(EigenTowerPaused)
	if err := _EigenTower.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenTowerRegistryUpdatedIterator is returned from FilterRegistryUpdated and is used to iterate over the raw logs and unpacked data for RegistryUpdated events raised by the EigenTower contract.
type EigenTowerRegistryUpdatedIterator struct {
	Event *EigenTowerRegistryUpdated // Event containing the contract specifics and raw log

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
func (it *EigenTowerRegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenTowerRegistryUpdated)
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
		it.Event = new(EigenTowerRegistryUpdated)
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
func (it *EigenTowerRegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenTowerRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenTowerRegistryUpdated represents a RegistryUpdated event raised by the EigenTower contract.
type EigenTowerRegistryUpdated struct {
	OldRegistry common.Address
	NewRegistry common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistryUpdated is a free log retrieval operation binding the contract event 0x482b97c53e48ffa324a976e2738053e9aff6eee04d8aac63b10e19411d869b82.
//
// Solidity: event RegistryUpdated(address oldRegistry, address newRegistry)
func (_EigenTower *EigenTowerFilterer) FilterRegistryUpdated(opts *bind.FilterOpts) (*EigenTowerRegistryUpdatedIterator, error) {

	logs, sub, err := _EigenTower.contract.FilterLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &EigenTowerRegistryUpdatedIterator{contract: _EigenTower.contract, event: "RegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchRegistryUpdated is a free log subscription operation binding the contract event 0x482b97c53e48ffa324a976e2738053e9aff6eee04d8aac63b10e19411d869b82.
//
// Solidity: event RegistryUpdated(address oldRegistry, address newRegistry)
func (_EigenTower *EigenTowerFilterer) WatchRegistryUpdated(opts *bind.WatchOpts, sink chan<- *EigenTowerRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _EigenTower.contract.WatchLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenTowerRegistryUpdated)
				if err := _EigenTower.contract.UnpackLog(event, "RegistryUpdated", log); err != nil {
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

// ParseRegistryUpdated is a log parse operation binding the contract event 0x482b97c53e48ffa324a976e2738053e9aff6eee04d8aac63b10e19411d869b82.
//
// Solidity: event RegistryUpdated(address oldRegistry, address newRegistry)
func (_EigenTower *EigenTowerFilterer) ParseRegistryUpdated(log types.Log) (*EigenTowerRegistryUpdated, error) {
	event := new(EigenTowerRegistryUpdated)
	if err := _EigenTower.contract.UnpackLog(event, "RegistryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenTowerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EigenTower contract.
type EigenTowerUnpausedIterator struct {
	Event *EigenTowerUnpaused // Event containing the contract specifics and raw log

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
func (it *EigenTowerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenTowerUnpaused)
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
		it.Event = new(EigenTowerUnpaused)
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
func (it *EigenTowerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenTowerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenTowerUnpaused represents a Unpaused event raised by the EigenTower contract.
type EigenTowerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EigenTower *EigenTowerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EigenTowerUnpausedIterator, error) {

	logs, sub, err := _EigenTower.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EigenTowerUnpausedIterator{contract: _EigenTower.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EigenTower *EigenTowerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EigenTowerUnpaused) (event.Subscription, error) {

	logs, sub, err := _EigenTower.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenTowerUnpaused)
				if err := _EigenTower.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_EigenTower *EigenTowerFilterer) ParseUnpaused(log types.Log) (*EigenTowerUnpaused, error) {
	event := new(EigenTowerUnpaused)
	if err := _EigenTower.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenTowerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EigenTower contract.
type EigenTowerUpgradedIterator struct {
	Event *EigenTowerUpgraded // Event containing the contract specifics and raw log

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
func (it *EigenTowerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenTowerUpgraded)
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
		it.Event = new(EigenTowerUpgraded)
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
func (it *EigenTowerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenTowerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenTowerUpgraded represents a Upgraded event raised by the EigenTower contract.
type EigenTowerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EigenTower *EigenTowerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EigenTowerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EigenTower.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EigenTowerUpgradedIterator{contract: _EigenTower.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EigenTower *EigenTowerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EigenTowerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EigenTower.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenTowerUpgraded)
				if err := _EigenTower.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_EigenTower *EigenTowerFilterer) ParseUpgraded(log types.Log) (*EigenTowerUpgraded, error) {
	event := new(EigenTowerUpgraded)
	if err := _EigenTower.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
