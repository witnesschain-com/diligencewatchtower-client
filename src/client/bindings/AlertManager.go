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

// IAlertManagerAlert is an auto generated low-level Go binding around an user-defined struct.
type IAlertManagerAlert struct {
	ChainID           *big.Int
	L2BlockNumber     *big.Int
	OriginalStateRoot []byte
	ComputedStateRoot []byte
	ProofOfDiligence  []byte
	Sender            common.Address
}

// AlertManagerMetaData contains all meta data concerning the AlertManager contract.
var AlertManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"NewAlertRaised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"alertsByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"originalStateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"computedStateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofOfDiligence\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"alertsByChainIDBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"originalStateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"computedStateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofOfDiligence\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"getAlerts\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"originalStateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"computedStateRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofOfDiligence\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"internalType\":\"structIAlertManager.Alert[]\",\"name\":\"alerts\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOperatorRegistry\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"contractIL2ChainMapping\",\"name\":\"_l2ChainMapping\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ChainMapping\",\"outputs\":[{\"internalType\":\"contractIL2ChainMapping\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_originalOutputRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_computedOutputRoot\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_proofOfDiligence\",\"type\":\"bytes\"}],\"name\":\"raiseAlert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIOperatorRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOperatorRegistry\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setOperatorRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// AlertManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AlertManagerMetaData.ABI instead.
var AlertManagerABI = AlertManagerMetaData.ABI

// AlertManager is an auto generated Go binding around an Ethereum contract.
type AlertManager struct {
	AlertManagerCaller     // Read-only binding to the contract
	AlertManagerTransactor // Write-only binding to the contract
	AlertManagerFilterer   // Log filterer for contract events
}

// AlertManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AlertManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlertManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AlertManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlertManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AlertManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlertManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AlertManagerSession struct {
	Contract     *AlertManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AlertManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AlertManagerCallerSession struct {
	Contract *AlertManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AlertManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AlertManagerTransactorSession struct {
	Contract     *AlertManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AlertManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AlertManagerRaw struct {
	Contract *AlertManager // Generic contract binding to access the raw methods on
}

// AlertManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AlertManagerCallerRaw struct {
	Contract *AlertManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AlertManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AlertManagerTransactorRaw struct {
	Contract *AlertManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAlertManager creates a new instance of AlertManager, bound to a specific deployed contract.
func NewAlertManager(address common.Address, backend bind.ContractBackend) (*AlertManager, error) {
	contract, err := bindAlertManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AlertManager{AlertManagerCaller: AlertManagerCaller{contract: contract}, AlertManagerTransactor: AlertManagerTransactor{contract: contract}, AlertManagerFilterer: AlertManagerFilterer{contract: contract}}, nil
}

// NewAlertManagerCaller creates a new read-only instance of AlertManager, bound to a specific deployed contract.
func NewAlertManagerCaller(address common.Address, caller bind.ContractCaller) (*AlertManagerCaller, error) {
	contract, err := bindAlertManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AlertManagerCaller{contract: contract}, nil
}

// NewAlertManagerTransactor creates a new write-only instance of AlertManager, bound to a specific deployed contract.
func NewAlertManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AlertManagerTransactor, error) {
	contract, err := bindAlertManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AlertManagerTransactor{contract: contract}, nil
}

// NewAlertManagerFilterer creates a new log filterer instance of AlertManager, bound to a specific deployed contract.
func NewAlertManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AlertManagerFilterer, error) {
	contract, err := bindAlertManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AlertManagerFilterer{contract: contract}, nil
}

// bindAlertManager binds a generic wrapper to an already deployed contract.
func bindAlertManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AlertManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AlertManager *AlertManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AlertManager.Contract.AlertManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AlertManager *AlertManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlertManager.Contract.AlertManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AlertManager *AlertManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AlertManager.Contract.AlertManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AlertManager *AlertManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AlertManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AlertManager *AlertManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlertManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AlertManager *AlertManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AlertManager.Contract.contract.Transact(opts, method, params...)
}

// AlertsByAddress is a free data retrieval call binding the contract method 0x6d8f4bc5.
//
// Solidity: function alertsByAddress(address ) view returns(uint256 chainID, uint256 l2BlockNumber, bytes originalStateRoot, bytes computedStateRoot, bytes proofOfDiligence, address sender)
func (_AlertManager *AlertManagerCaller) AlertsByAddress(opts *bind.CallOpts, arg0 common.Address) (struct {
	ChainID           *big.Int
	L2BlockNumber     *big.Int
	OriginalStateRoot []byte
	ComputedStateRoot []byte
	ProofOfDiligence  []byte
	Sender            common.Address
}, error) {
	var out []interface{}
	err := _AlertManager.contract.Call(opts, &out, "alertsByAddress", arg0)

	outstruct := new(struct {
		ChainID           *big.Int
		L2BlockNumber     *big.Int
		OriginalStateRoot []byte
		ComputedStateRoot []byte
		ProofOfDiligence  []byte
		Sender            common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.L2BlockNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OriginalStateRoot = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.ComputedStateRoot = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.ProofOfDiligence = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.Sender = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AlertsByAddress is a free data retrieval call binding the contract method 0x6d8f4bc5.
//
// Solidity: function alertsByAddress(address ) view returns(uint256 chainID, uint256 l2BlockNumber, bytes originalStateRoot, bytes computedStateRoot, bytes proofOfDiligence, address sender)
func (_AlertManager *AlertManagerSession) AlertsByAddress(arg0 common.Address) (struct {
	ChainID           *big.Int
	L2BlockNumber     *big.Int
	OriginalStateRoot []byte
	ComputedStateRoot []byte
	ProofOfDiligence  []byte
	Sender            common.Address
}, error) {
	return _AlertManager.Contract.AlertsByAddress(&_AlertManager.CallOpts, arg0)
}

// AlertsByAddress is a free data retrieval call binding the contract method 0x6d8f4bc5.
//
// Solidity: function alertsByAddress(address ) view returns(uint256 chainID, uint256 l2BlockNumber, bytes originalStateRoot, bytes computedStateRoot, bytes proofOfDiligence, address sender)
func (_AlertManager *AlertManagerCallerSession) AlertsByAddress(arg0 common.Address) (struct {
	ChainID           *big.Int
	L2BlockNumber     *big.Int
	OriginalStateRoot []byte
	ComputedStateRoot []byte
	ProofOfDiligence  []byte
	Sender            common.Address
}, error) {
	return _AlertManager.Contract.AlertsByAddress(&_AlertManager.CallOpts, arg0)
}

// AlertsByChainIDBlockNumber is a free data retrieval call binding the contract method 0x45b53f63.
//
// Solidity: function alertsByChainIDBlockNumber(uint256 , uint256 , uint256 ) view returns(uint256 chainID, uint256 l2BlockNumber, bytes originalStateRoot, bytes computedStateRoot, bytes proofOfDiligence, address sender)
func (_AlertManager *AlertManagerCaller) AlertsByChainIDBlockNumber(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (struct {
	ChainID           *big.Int
	L2BlockNumber     *big.Int
	OriginalStateRoot []byte
	ComputedStateRoot []byte
	ProofOfDiligence  []byte
	Sender            common.Address
}, error) {
	var out []interface{}
	err := _AlertManager.contract.Call(opts, &out, "alertsByChainIDBlockNumber", arg0, arg1, arg2)

	outstruct := new(struct {
		ChainID           *big.Int
		L2BlockNumber     *big.Int
		OriginalStateRoot []byte
		ComputedStateRoot []byte
		ProofOfDiligence  []byte
		Sender            common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.L2BlockNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OriginalStateRoot = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.ComputedStateRoot = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.ProofOfDiligence = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.Sender = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AlertsByChainIDBlockNumber is a free data retrieval call binding the contract method 0x45b53f63.
//
// Solidity: function alertsByChainIDBlockNumber(uint256 , uint256 , uint256 ) view returns(uint256 chainID, uint256 l2BlockNumber, bytes originalStateRoot, bytes computedStateRoot, bytes proofOfDiligence, address sender)
func (_AlertManager *AlertManagerSession) AlertsByChainIDBlockNumber(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (struct {
	ChainID           *big.Int
	L2BlockNumber     *big.Int
	OriginalStateRoot []byte
	ComputedStateRoot []byte
	ProofOfDiligence  []byte
	Sender            common.Address
}, error) {
	return _AlertManager.Contract.AlertsByChainIDBlockNumber(&_AlertManager.CallOpts, arg0, arg1, arg2)
}

// AlertsByChainIDBlockNumber is a free data retrieval call binding the contract method 0x45b53f63.
//
// Solidity: function alertsByChainIDBlockNumber(uint256 , uint256 , uint256 ) view returns(uint256 chainID, uint256 l2BlockNumber, bytes originalStateRoot, bytes computedStateRoot, bytes proofOfDiligence, address sender)
func (_AlertManager *AlertManagerCallerSession) AlertsByChainIDBlockNumber(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (struct {
	ChainID           *big.Int
	L2BlockNumber     *big.Int
	OriginalStateRoot []byte
	ComputedStateRoot []byte
	ProofOfDiligence  []byte
	Sender            common.Address
}, error) {
	return _AlertManager.Contract.AlertsByChainIDBlockNumber(&_AlertManager.CallOpts, arg0, arg1, arg2)
}

// GetAlerts is a free data retrieval call binding the contract method 0x9cd33a3f.
//
// Solidity: function getAlerts(uint256 _chainID, uint256 _l2BlockNumber) view returns((uint256,uint256,bytes,bytes,bytes,address)[] alerts)
func (_AlertManager *AlertManagerCaller) GetAlerts(opts *bind.CallOpts, _chainID *big.Int, _l2BlockNumber *big.Int) ([]IAlertManagerAlert, error) {
	var out []interface{}
	err := _AlertManager.contract.Call(opts, &out, "getAlerts", _chainID, _l2BlockNumber)

	if err != nil {
		return *new([]IAlertManagerAlert), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAlertManagerAlert)).(*[]IAlertManagerAlert)

	return out0, err

}

// GetAlerts is a free data retrieval call binding the contract method 0x9cd33a3f.
//
// Solidity: function getAlerts(uint256 _chainID, uint256 _l2BlockNumber) view returns((uint256,uint256,bytes,bytes,bytes,address)[] alerts)
func (_AlertManager *AlertManagerSession) GetAlerts(_chainID *big.Int, _l2BlockNumber *big.Int) ([]IAlertManagerAlert, error) {
	return _AlertManager.Contract.GetAlerts(&_AlertManager.CallOpts, _chainID, _l2BlockNumber)
}

// GetAlerts is a free data retrieval call binding the contract method 0x9cd33a3f.
//
// Solidity: function getAlerts(uint256 _chainID, uint256 _l2BlockNumber) view returns((uint256,uint256,bytes,bytes,bytes,address)[] alerts)
func (_AlertManager *AlertManagerCallerSession) GetAlerts(_chainID *big.Int, _l2BlockNumber *big.Int) ([]IAlertManagerAlert, error) {
	return _AlertManager.Contract.GetAlerts(&_AlertManager.CallOpts, _chainID, _l2BlockNumber)
}

// L2ChainMapping is a free data retrieval call binding the contract method 0x820b6f39.
//
// Solidity: function l2ChainMapping() view returns(address)
func (_AlertManager *AlertManagerCaller) L2ChainMapping(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlertManager.contract.Call(opts, &out, "l2ChainMapping")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ChainMapping is a free data retrieval call binding the contract method 0x820b6f39.
//
// Solidity: function l2ChainMapping() view returns(address)
func (_AlertManager *AlertManagerSession) L2ChainMapping() (common.Address, error) {
	return _AlertManager.Contract.L2ChainMapping(&_AlertManager.CallOpts)
}

// L2ChainMapping is a free data retrieval call binding the contract method 0x820b6f39.
//
// Solidity: function l2ChainMapping() view returns(address)
func (_AlertManager *AlertManagerCallerSession) L2ChainMapping() (common.Address, error) {
	return _AlertManager.Contract.L2ChainMapping(&_AlertManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AlertManager *AlertManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlertManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AlertManager *AlertManagerSession) Owner() (common.Address, error) {
	return _AlertManager.Contract.Owner(&_AlertManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AlertManager *AlertManagerCallerSession) Owner() (common.Address, error) {
	return _AlertManager.Contract.Owner(&_AlertManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AlertManager *AlertManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AlertManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AlertManager *AlertManagerSession) Paused() (bool, error) {
	return _AlertManager.Contract.Paused(&_AlertManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AlertManager *AlertManagerCallerSession) Paused() (bool, error) {
	return _AlertManager.Contract.Paused(&_AlertManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AlertManager *AlertManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AlertManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AlertManager *AlertManagerSession) ProxiableUUID() ([32]byte, error) {
	return _AlertManager.Contract.ProxiableUUID(&_AlertManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AlertManager *AlertManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AlertManager.Contract.ProxiableUUID(&_AlertManager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AlertManager *AlertManagerCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlertManager.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AlertManager *AlertManagerSession) Registry() (common.Address, error) {
	return _AlertManager.Contract.Registry(&_AlertManager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AlertManager *AlertManagerCallerSession) Registry() (common.Address, error) {
	return _AlertManager.Contract.Registry(&_AlertManager.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registry, address _l2ChainMapping) returns()
func (_AlertManager *AlertManagerTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _l2ChainMapping common.Address) (*types.Transaction, error) {
	return _AlertManager.contract.Transact(opts, "initialize", _registry, _l2ChainMapping)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registry, address _l2ChainMapping) returns()
func (_AlertManager *AlertManagerSession) Initialize(_registry common.Address, _l2ChainMapping common.Address) (*types.Transaction, error) {
	return _AlertManager.Contract.Initialize(&_AlertManager.TransactOpts, _registry, _l2ChainMapping)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registry, address _l2ChainMapping) returns()
func (_AlertManager *AlertManagerTransactorSession) Initialize(_registry common.Address, _l2ChainMapping common.Address) (*types.Transaction, error) {
	return _AlertManager.Contract.Initialize(&_AlertManager.TransactOpts, _registry, _l2ChainMapping)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AlertManager *AlertManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlertManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AlertManager *AlertManagerSession) Pause() (*types.Transaction, error) {
	return _AlertManager.Contract.Pause(&_AlertManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AlertManager *AlertManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _AlertManager.Contract.Pause(&_AlertManager.TransactOpts)
}

// RaiseAlert is a paid mutator transaction binding the contract method 0x0de4a0e7.
//
// Solidity: function raiseAlert(uint256 _chainID, uint256 _l2BlockNumber, bytes _originalOutputRoot, bytes _computedOutputRoot, bytes _proofOfDiligence) returns()
func (_AlertManager *AlertManagerTransactor) RaiseAlert(opts *bind.TransactOpts, _chainID *big.Int, _l2BlockNumber *big.Int, _originalOutputRoot []byte, _computedOutputRoot []byte, _proofOfDiligence []byte) (*types.Transaction, error) {
	return _AlertManager.contract.Transact(opts, "raiseAlert", _chainID, _l2BlockNumber, _originalOutputRoot, _computedOutputRoot, _proofOfDiligence)
}

// RaiseAlert is a paid mutator transaction binding the contract method 0x0de4a0e7.
//
// Solidity: function raiseAlert(uint256 _chainID, uint256 _l2BlockNumber, bytes _originalOutputRoot, bytes _computedOutputRoot, bytes _proofOfDiligence) returns()
func (_AlertManager *AlertManagerSession) RaiseAlert(_chainID *big.Int, _l2BlockNumber *big.Int, _originalOutputRoot []byte, _computedOutputRoot []byte, _proofOfDiligence []byte) (*types.Transaction, error) {
	return _AlertManager.Contract.RaiseAlert(&_AlertManager.TransactOpts, _chainID, _l2BlockNumber, _originalOutputRoot, _computedOutputRoot, _proofOfDiligence)
}

// RaiseAlert is a paid mutator transaction binding the contract method 0x0de4a0e7.
//
// Solidity: function raiseAlert(uint256 _chainID, uint256 _l2BlockNumber, bytes _originalOutputRoot, bytes _computedOutputRoot, bytes _proofOfDiligence) returns()
func (_AlertManager *AlertManagerTransactorSession) RaiseAlert(_chainID *big.Int, _l2BlockNumber *big.Int, _originalOutputRoot []byte, _computedOutputRoot []byte, _proofOfDiligence []byte) (*types.Transaction, error) {
	return _AlertManager.Contract.RaiseAlert(&_AlertManager.TransactOpts, _chainID, _l2BlockNumber, _originalOutputRoot, _computedOutputRoot, _proofOfDiligence)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AlertManager *AlertManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlertManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AlertManager *AlertManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _AlertManager.Contract.RenounceOwnership(&_AlertManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AlertManager *AlertManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AlertManager.Contract.RenounceOwnership(&_AlertManager.TransactOpts)
}

// SetOperatorRegistry is a paid mutator transaction binding the contract method 0x9d28fb86.
//
// Solidity: function setOperatorRegistry(address _registry) returns()
func (_AlertManager *AlertManagerTransactor) SetOperatorRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _AlertManager.contract.Transact(opts, "setOperatorRegistry", _registry)
}

// SetOperatorRegistry is a paid mutator transaction binding the contract method 0x9d28fb86.
//
// Solidity: function setOperatorRegistry(address _registry) returns()
func (_AlertManager *AlertManagerSession) SetOperatorRegistry(_registry common.Address) (*types.Transaction, error) {
	return _AlertManager.Contract.SetOperatorRegistry(&_AlertManager.TransactOpts, _registry)
}

// SetOperatorRegistry is a paid mutator transaction binding the contract method 0x9d28fb86.
//
// Solidity: function setOperatorRegistry(address _registry) returns()
func (_AlertManager *AlertManagerTransactorSession) SetOperatorRegistry(_registry common.Address) (*types.Transaction, error) {
	return _AlertManager.Contract.SetOperatorRegistry(&_AlertManager.TransactOpts, _registry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AlertManager *AlertManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AlertManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AlertManager *AlertManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AlertManager.Contract.TransferOwnership(&_AlertManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AlertManager *AlertManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AlertManager.Contract.TransferOwnership(&_AlertManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AlertManager *AlertManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlertManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AlertManager *AlertManagerSession) Unpause() (*types.Transaction, error) {
	return _AlertManager.Contract.Unpause(&_AlertManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AlertManager *AlertManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _AlertManager.Contract.Unpause(&_AlertManager.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AlertManager *AlertManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _AlertManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AlertManager *AlertManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AlertManager.Contract.UpgradeTo(&_AlertManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AlertManager *AlertManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AlertManager.Contract.UpgradeTo(&_AlertManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AlertManager *AlertManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AlertManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AlertManager *AlertManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AlertManager.Contract.UpgradeToAndCall(&_AlertManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AlertManager *AlertManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AlertManager.Contract.UpgradeToAndCall(&_AlertManager.TransactOpts, newImplementation, data)
}

// AlertManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the AlertManager contract.
type AlertManagerAdminChangedIterator struct {
	Event *AlertManagerAdminChanged // Event containing the contract specifics and raw log

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
func (it *AlertManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlertManagerAdminChanged)
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
		it.Event = new(AlertManagerAdminChanged)
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
func (it *AlertManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlertManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlertManagerAdminChanged represents a AdminChanged event raised by the AlertManager contract.
type AlertManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AlertManager *AlertManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AlertManagerAdminChangedIterator, error) {

	logs, sub, err := _AlertManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AlertManagerAdminChangedIterator{contract: _AlertManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AlertManager *AlertManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AlertManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _AlertManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlertManagerAdminChanged)
				if err := _AlertManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_AlertManager *AlertManagerFilterer) ParseAdminChanged(log types.Log) (*AlertManagerAdminChanged, error) {
	event := new(AlertManagerAdminChanged)
	if err := _AlertManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlertManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the AlertManager contract.
type AlertManagerBeaconUpgradedIterator struct {
	Event *AlertManagerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *AlertManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlertManagerBeaconUpgraded)
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
		it.Event = new(AlertManagerBeaconUpgraded)
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
func (it *AlertManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlertManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlertManagerBeaconUpgraded represents a BeaconUpgraded event raised by the AlertManager contract.
type AlertManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AlertManager *AlertManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AlertManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AlertManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AlertManagerBeaconUpgradedIterator{contract: _AlertManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AlertManager *AlertManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AlertManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AlertManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlertManagerBeaconUpgraded)
				if err := _AlertManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_AlertManager *AlertManagerFilterer) ParseBeaconUpgraded(log types.Log) (*AlertManagerBeaconUpgraded, error) {
	event := new(AlertManagerBeaconUpgraded)
	if err := _AlertManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlertManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AlertManager contract.
type AlertManagerInitializedIterator struct {
	Event *AlertManagerInitialized // Event containing the contract specifics and raw log

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
func (it *AlertManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlertManagerInitialized)
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
		it.Event = new(AlertManagerInitialized)
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
func (it *AlertManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlertManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlertManagerInitialized represents a Initialized event raised by the AlertManager contract.
type AlertManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AlertManager *AlertManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AlertManagerInitializedIterator, error) {

	logs, sub, err := _AlertManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AlertManagerInitializedIterator{contract: _AlertManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AlertManager *AlertManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AlertManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _AlertManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlertManagerInitialized)
				if err := _AlertManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AlertManager *AlertManagerFilterer) ParseInitialized(log types.Log) (*AlertManagerInitialized, error) {
	event := new(AlertManagerInitialized)
	if err := _AlertManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlertManagerNewAlertRaisedIterator is returned from FilterNewAlertRaised and is used to iterate over the raw logs and unpacked data for NewAlertRaised events raised by the AlertManager contract.
type AlertManagerNewAlertRaisedIterator struct {
	Event *AlertManagerNewAlertRaised // Event containing the contract specifics and raw log

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
func (it *AlertManagerNewAlertRaisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlertManagerNewAlertRaised)
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
		it.Event = new(AlertManagerNewAlertRaised)
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
func (it *AlertManagerNewAlertRaisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlertManagerNewAlertRaisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlertManagerNewAlertRaised represents a NewAlertRaised event raised by the AlertManager contract.
type AlertManagerNewAlertRaised struct {
	Sender        common.Address
	ChainID       *big.Int
	L2BlockNumber *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewAlertRaised is a free log retrieval operation binding the contract event 0x98dfa9179e1cf0fe5a0929afca50c816752df62173cdea2208730e805ca63fbc.
//
// Solidity: event NewAlertRaised(address sender, uint256 chainID, uint256 l2BlockNumber)
func (_AlertManager *AlertManagerFilterer) FilterNewAlertRaised(opts *bind.FilterOpts) (*AlertManagerNewAlertRaisedIterator, error) {

	logs, sub, err := _AlertManager.contract.FilterLogs(opts, "NewAlertRaised")
	if err != nil {
		return nil, err
	}
	return &AlertManagerNewAlertRaisedIterator{contract: _AlertManager.contract, event: "NewAlertRaised", logs: logs, sub: sub}, nil
}

// WatchNewAlertRaised is a free log subscription operation binding the contract event 0x98dfa9179e1cf0fe5a0929afca50c816752df62173cdea2208730e805ca63fbc.
//
// Solidity: event NewAlertRaised(address sender, uint256 chainID, uint256 l2BlockNumber)
func (_AlertManager *AlertManagerFilterer) WatchNewAlertRaised(opts *bind.WatchOpts, sink chan<- *AlertManagerNewAlertRaised) (event.Subscription, error) {

	logs, sub, err := _AlertManager.contract.WatchLogs(opts, "NewAlertRaised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlertManagerNewAlertRaised)
				if err := _AlertManager.contract.UnpackLog(event, "NewAlertRaised", log); err != nil {
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

// ParseNewAlertRaised is a log parse operation binding the contract event 0x98dfa9179e1cf0fe5a0929afca50c816752df62173cdea2208730e805ca63fbc.
//
// Solidity: event NewAlertRaised(address sender, uint256 chainID, uint256 l2BlockNumber)
func (_AlertManager *AlertManagerFilterer) ParseNewAlertRaised(log types.Log) (*AlertManagerNewAlertRaised, error) {
	event := new(AlertManagerNewAlertRaised)
	if err := _AlertManager.contract.UnpackLog(event, "NewAlertRaised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlertManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AlertManager contract.
type AlertManagerOwnershipTransferredIterator struct {
	Event *AlertManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AlertManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlertManagerOwnershipTransferred)
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
		it.Event = new(AlertManagerOwnershipTransferred)
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
func (it *AlertManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlertManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlertManagerOwnershipTransferred represents a OwnershipTransferred event raised by the AlertManager contract.
type AlertManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AlertManager *AlertManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AlertManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AlertManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AlertManagerOwnershipTransferredIterator{contract: _AlertManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AlertManager *AlertManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AlertManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AlertManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlertManagerOwnershipTransferred)
				if err := _AlertManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AlertManager *AlertManagerFilterer) ParseOwnershipTransferred(log types.Log) (*AlertManagerOwnershipTransferred, error) {
	event := new(AlertManagerOwnershipTransferred)
	if err := _AlertManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlertManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AlertManager contract.
type AlertManagerPausedIterator struct {
	Event *AlertManagerPaused // Event containing the contract specifics and raw log

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
func (it *AlertManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlertManagerPaused)
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
		it.Event = new(AlertManagerPaused)
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
func (it *AlertManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlertManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlertManagerPaused represents a Paused event raised by the AlertManager contract.
type AlertManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AlertManager *AlertManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*AlertManagerPausedIterator, error) {

	logs, sub, err := _AlertManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AlertManagerPausedIterator{contract: _AlertManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AlertManager *AlertManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AlertManagerPaused) (event.Subscription, error) {

	logs, sub, err := _AlertManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlertManagerPaused)
				if err := _AlertManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_AlertManager *AlertManagerFilterer) ParsePaused(log types.Log) (*AlertManagerPaused, error) {
	event := new(AlertManagerPaused)
	if err := _AlertManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlertManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AlertManager contract.
type AlertManagerUnpausedIterator struct {
	Event *AlertManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *AlertManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlertManagerUnpaused)
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
		it.Event = new(AlertManagerUnpaused)
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
func (it *AlertManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlertManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlertManagerUnpaused represents a Unpaused event raised by the AlertManager contract.
type AlertManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AlertManager *AlertManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AlertManagerUnpausedIterator, error) {

	logs, sub, err := _AlertManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AlertManagerUnpausedIterator{contract: _AlertManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AlertManager *AlertManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AlertManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _AlertManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlertManagerUnpaused)
				if err := _AlertManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_AlertManager *AlertManagerFilterer) ParseUnpaused(log types.Log) (*AlertManagerUnpaused, error) {
	event := new(AlertManagerUnpaused)
	if err := _AlertManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlertManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AlertManager contract.
type AlertManagerUpgradedIterator struct {
	Event *AlertManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *AlertManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlertManagerUpgraded)
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
		it.Event = new(AlertManagerUpgraded)
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
func (it *AlertManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlertManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlertManagerUpgraded represents a Upgraded event raised by the AlertManager contract.
type AlertManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AlertManager *AlertManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AlertManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AlertManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AlertManagerUpgradedIterator{contract: _AlertManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AlertManager *AlertManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AlertManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AlertManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlertManagerUpgraded)
				if err := _AlertManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AlertManager *AlertManagerFilterer) ParseUpgraded(log types.Log) (*AlertManagerUpgraded, error) {
	event := new(AlertManagerUpgraded)
	if err := _AlertManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
