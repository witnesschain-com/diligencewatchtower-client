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

// WatchTowerManagerMetaData contains all meta data concerning the WatchTowerManager contract.
var WatchTowerManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractOperatorRegistry\",\"name\":\"_operatorRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"operatortaskNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"lastBlockToBeServed\",\"type\":\"uint32\"}],\"name\":\"TaskSubmited\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_SERVING_BLOCKS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestServeUntilBlock\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestServeUntilBlock\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorRegistry\",\"outputs\":[{\"internalType\":\"contractOperatorRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"publicKey\",\"type\":\"bytes32\"}],\"name\":\"registerOperatorasWatchTower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitTask\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskNumber\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WatchTowerManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use WatchTowerManagerMetaData.ABI instead.
var WatchTowerManagerABI = WatchTowerManagerMetaData.ABI

// WatchTowerManager is an auto generated Go binding around an Ethereum contract.
type WatchTowerManager struct {
	WatchTowerManagerCaller     // Read-only binding to the contract
	WatchTowerManagerTransactor // Write-only binding to the contract
	WatchTowerManagerFilterer   // Log filterer for contract events
}

// WatchTowerManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WatchTowerManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WatchTowerManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WatchTowerManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WatchTowerManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WatchTowerManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WatchTowerManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WatchTowerManagerSession struct {
	Contract     *WatchTowerManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WatchTowerManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WatchTowerManagerCallerSession struct {
	Contract *WatchTowerManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// WatchTowerManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WatchTowerManagerTransactorSession struct {
	Contract     *WatchTowerManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// WatchTowerManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WatchTowerManagerRaw struct {
	Contract *WatchTowerManager // Generic contract binding to access the raw methods on
}

// WatchTowerManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WatchTowerManagerCallerRaw struct {
	Contract *WatchTowerManagerCaller // Generic read-only contract binding to access the raw methods on
}

// WatchTowerManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WatchTowerManagerTransactorRaw struct {
	Contract *WatchTowerManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWatchTowerManager creates a new instance of WatchTowerManager, bound to a specific deployed contract.
func NewWatchTowerManager(address common.Address, backend bind.ContractBackend) (*WatchTowerManager, error) {
	contract, err := bindWatchTowerManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WatchTowerManager{WatchTowerManagerCaller: WatchTowerManagerCaller{contract: contract}, WatchTowerManagerTransactor: WatchTowerManagerTransactor{contract: contract}, WatchTowerManagerFilterer: WatchTowerManagerFilterer{contract: contract}}, nil
}

// NewWatchTowerManagerCaller creates a new read-only instance of WatchTowerManager, bound to a specific deployed contract.
func NewWatchTowerManagerCaller(address common.Address, caller bind.ContractCaller) (*WatchTowerManagerCaller, error) {
	contract, err := bindWatchTowerManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WatchTowerManagerCaller{contract: contract}, nil
}

// NewWatchTowerManagerTransactor creates a new write-only instance of WatchTowerManager, bound to a specific deployed contract.
func NewWatchTowerManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*WatchTowerManagerTransactor, error) {
	contract, err := bindWatchTowerManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WatchTowerManagerTransactor{contract: contract}, nil
}

// NewWatchTowerManagerFilterer creates a new log filterer instance of WatchTowerManager, bound to a specific deployed contract.
func NewWatchTowerManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*WatchTowerManagerFilterer, error) {
	contract, err := bindWatchTowerManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WatchTowerManagerFilterer{contract: contract}, nil
}

// bindWatchTowerManager binds a generic wrapper to an already deployed contract.
func bindWatchTowerManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WatchTowerManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WatchTowerManager *WatchTowerManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WatchTowerManager.Contract.WatchTowerManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WatchTowerManager *WatchTowerManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WatchTowerManager.Contract.WatchTowerManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WatchTowerManager *WatchTowerManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WatchTowerManager.Contract.WatchTowerManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WatchTowerManager *WatchTowerManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WatchTowerManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WatchTowerManager *WatchTowerManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WatchTowerManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WatchTowerManager *WatchTowerManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WatchTowerManager.Contract.contract.Transact(opts, method, params...)
}

// MAXSERVINGBLOCKS is a free data retrieval call binding the contract method 0xd2887e9f.
//
// Solidity: function MAX_SERVING_BLOCKS() view returns(uint32)
func (_WatchTowerManager *WatchTowerManagerCaller) MAXSERVINGBLOCKS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _WatchTowerManager.contract.Call(opts, &out, "MAX_SERVING_BLOCKS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXSERVINGBLOCKS is a free data retrieval call binding the contract method 0xd2887e9f.
//
// Solidity: function MAX_SERVING_BLOCKS() view returns(uint32)
func (_WatchTowerManager *WatchTowerManagerSession) MAXSERVINGBLOCKS() (uint32, error) {
	return _WatchTowerManager.Contract.MAXSERVINGBLOCKS(&_WatchTowerManager.CallOpts)
}

// MAXSERVINGBLOCKS is a free data retrieval call binding the contract method 0xd2887e9f.
//
// Solidity: function MAX_SERVING_BLOCKS() view returns(uint32)
func (_WatchTowerManager *WatchTowerManagerCallerSession) MAXSERVINGBLOCKS() (uint32, error) {
	return _WatchTowerManager.Contract.MAXSERVINGBLOCKS(&_WatchTowerManager.CallOpts)
}

// GetLatestServeUntilBlock is a free data retrieval call binding the contract method 0x7e3c3349.
//
// Solidity: function getLatestServeUntilBlock() view returns(uint32)
func (_WatchTowerManager *WatchTowerManagerCaller) GetLatestServeUntilBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _WatchTowerManager.contract.Call(opts, &out, "getLatestServeUntilBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetLatestServeUntilBlock is a free data retrieval call binding the contract method 0x7e3c3349.
//
// Solidity: function getLatestServeUntilBlock() view returns(uint32)
func (_WatchTowerManager *WatchTowerManagerSession) GetLatestServeUntilBlock() (uint32, error) {
	return _WatchTowerManager.Contract.GetLatestServeUntilBlock(&_WatchTowerManager.CallOpts)
}

// GetLatestServeUntilBlock is a free data retrieval call binding the contract method 0x7e3c3349.
//
// Solidity: function getLatestServeUntilBlock() view returns(uint32)
func (_WatchTowerManager *WatchTowerManagerCallerSession) GetLatestServeUntilBlock() (uint32, error) {
	return _WatchTowerManager.Contract.GetLatestServeUntilBlock(&_WatchTowerManager.CallOpts)
}

// LatestServeUntilBlock is a free data retrieval call binding the contract method 0x758f8dba.
//
// Solidity: function latestServeUntilBlock() view returns(uint32)
func (_WatchTowerManager *WatchTowerManagerCaller) LatestServeUntilBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _WatchTowerManager.contract.Call(opts, &out, "latestServeUntilBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestServeUntilBlock is a free data retrieval call binding the contract method 0x758f8dba.
//
// Solidity: function latestServeUntilBlock() view returns(uint32)
func (_WatchTowerManager *WatchTowerManagerSession) LatestServeUntilBlock() (uint32, error) {
	return _WatchTowerManager.Contract.LatestServeUntilBlock(&_WatchTowerManager.CallOpts)
}

// LatestServeUntilBlock is a free data retrieval call binding the contract method 0x758f8dba.
//
// Solidity: function latestServeUntilBlock() view returns(uint32)
func (_WatchTowerManager *WatchTowerManagerCallerSession) LatestServeUntilBlock() (uint32, error) {
	return _WatchTowerManager.Contract.LatestServeUntilBlock(&_WatchTowerManager.CallOpts)
}

// OperatorRegistry is a free data retrieval call binding the contract method 0x58c2225b.
//
// Solidity: function operatorRegistry() view returns(address)
func (_WatchTowerManager *WatchTowerManagerCaller) OperatorRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WatchTowerManager.contract.Call(opts, &out, "operatorRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorRegistry is a free data retrieval call binding the contract method 0x58c2225b.
//
// Solidity: function operatorRegistry() view returns(address)
func (_WatchTowerManager *WatchTowerManagerSession) OperatorRegistry() (common.Address, error) {
	return _WatchTowerManager.Contract.OperatorRegistry(&_WatchTowerManager.CallOpts)
}

// OperatorRegistry is a free data retrieval call binding the contract method 0x58c2225b.
//
// Solidity: function operatorRegistry() view returns(address)
func (_WatchTowerManager *WatchTowerManagerCallerSession) OperatorRegistry() (common.Address, error) {
	return _WatchTowerManager.Contract.OperatorRegistry(&_WatchTowerManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WatchTowerManager *WatchTowerManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WatchTowerManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WatchTowerManager *WatchTowerManagerSession) Owner() (common.Address, error) {
	return _WatchTowerManager.Contract.Owner(&_WatchTowerManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WatchTowerManager *WatchTowerManagerCallerSession) Owner() (common.Address, error) {
	return _WatchTowerManager.Contract.Owner(&_WatchTowerManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_WatchTowerManager *WatchTowerManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _WatchTowerManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_WatchTowerManager *WatchTowerManagerSession) TaskNumber() (uint32, error) {
	return _WatchTowerManager.Contract.TaskNumber(&_WatchTowerManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_WatchTowerManager *WatchTowerManagerCallerSession) TaskNumber() (uint32, error) {
	return _WatchTowerManager.Contract.TaskNumber(&_WatchTowerManager.CallOpts)
}

// RegisterOperatorasWatchTower is a paid mutator transaction binding the contract method 0x8fa22ad0.
//
// Solidity: function registerOperatorasWatchTower(address operator, bytes32 publicKey) returns()
func (_WatchTowerManager *WatchTowerManagerTransactor) RegisterOperatorasWatchTower(opts *bind.TransactOpts, operator common.Address, publicKey [32]byte) (*types.Transaction, error) {
	return _WatchTowerManager.contract.Transact(opts, "registerOperatorasWatchTower", operator, publicKey)
}

// RegisterOperatorasWatchTower is a paid mutator transaction binding the contract method 0x8fa22ad0.
//
// Solidity: function registerOperatorasWatchTower(address operator, bytes32 publicKey) returns()
func (_WatchTowerManager *WatchTowerManagerSession) RegisterOperatorasWatchTower(operator common.Address, publicKey [32]byte) (*types.Transaction, error) {
	return _WatchTowerManager.Contract.RegisterOperatorasWatchTower(&_WatchTowerManager.TransactOpts, operator, publicKey)
}

// RegisterOperatorasWatchTower is a paid mutator transaction binding the contract method 0x8fa22ad0.
//
// Solidity: function registerOperatorasWatchTower(address operator, bytes32 publicKey) returns()
func (_WatchTowerManager *WatchTowerManagerTransactorSession) RegisterOperatorasWatchTower(operator common.Address, publicKey [32]byte) (*types.Transaction, error) {
	return _WatchTowerManager.Contract.RegisterOperatorasWatchTower(&_WatchTowerManager.TransactOpts, operator, publicKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WatchTowerManager *WatchTowerManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WatchTowerManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WatchTowerManager *WatchTowerManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _WatchTowerManager.Contract.RenounceOwnership(&_WatchTowerManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WatchTowerManager *WatchTowerManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WatchTowerManager.Contract.RenounceOwnership(&_WatchTowerManager.TransactOpts)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x4abbef9b.
//
// Solidity: function submitTask() returns(uint32)
func (_WatchTowerManager *WatchTowerManagerTransactor) SubmitTask(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WatchTowerManager.contract.Transact(opts, "submitTask")
}

// SubmitTask is a paid mutator transaction binding the contract method 0x4abbef9b.
//
// Solidity: function submitTask() returns(uint32)
func (_WatchTowerManager *WatchTowerManagerSession) SubmitTask() (*types.Transaction, error) {
	return _WatchTowerManager.Contract.SubmitTask(&_WatchTowerManager.TransactOpts)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x4abbef9b.
//
// Solidity: function submitTask() returns(uint32)
func (_WatchTowerManager *WatchTowerManagerTransactorSession) SubmitTask() (*types.Transaction, error) {
	return _WatchTowerManager.Contract.SubmitTask(&_WatchTowerManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WatchTowerManager *WatchTowerManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WatchTowerManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WatchTowerManager *WatchTowerManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WatchTowerManager.Contract.TransferOwnership(&_WatchTowerManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WatchTowerManager *WatchTowerManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WatchTowerManager.Contract.TransferOwnership(&_WatchTowerManager.TransactOpts, newOwner)
}

// WatchTowerManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WatchTowerManager contract.
type WatchTowerManagerOwnershipTransferredIterator struct {
	Event *WatchTowerManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WatchTowerManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WatchTowerManagerOwnershipTransferred)
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
		it.Event = new(WatchTowerManagerOwnershipTransferred)
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
func (it *WatchTowerManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WatchTowerManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WatchTowerManagerOwnershipTransferred represents a OwnershipTransferred event raised by the WatchTowerManager contract.
type WatchTowerManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WatchTowerManager *WatchTowerManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WatchTowerManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WatchTowerManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WatchTowerManagerOwnershipTransferredIterator{contract: _WatchTowerManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WatchTowerManager *WatchTowerManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WatchTowerManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WatchTowerManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WatchTowerManagerOwnershipTransferred)
				if err := _WatchTowerManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WatchTowerManager *WatchTowerManagerFilterer) ParseOwnershipTransferred(log types.Log) (*WatchTowerManagerOwnershipTransferred, error) {
	event := new(WatchTowerManagerOwnershipTransferred)
	if err := _WatchTowerManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WatchTowerManagerTaskSubmitedIterator is returned from FilterTaskSubmited and is used to iterate over the raw logs and unpacked data for TaskSubmited events raised by the WatchTowerManager contract.
type WatchTowerManagerTaskSubmitedIterator struct {
	Event *WatchTowerManagerTaskSubmited // Event containing the contract specifics and raw log

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
func (it *WatchTowerManagerTaskSubmitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WatchTowerManagerTaskSubmited)
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
		it.Event = new(WatchTowerManagerTaskSubmited)
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
func (it *WatchTowerManagerTaskSubmitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WatchTowerManagerTaskSubmitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WatchTowerManagerTaskSubmited represents a TaskSubmited event raised by the WatchTowerManager contract.
type WatchTowerManagerTaskSubmited struct {
	Operator            common.Address
	OperatortaskNumber  uint32
	LastBlockToBeServed uint32
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTaskSubmited is a free log retrieval operation binding the contract event 0xa12a0d6d059e62f1f1ff5e187dc823d3b39af2c2cd7bc56dfbbb720c283fa1b8.
//
// Solidity: event TaskSubmited(address operator, uint32 operatortaskNumber, uint32 lastBlockToBeServed)
func (_WatchTowerManager *WatchTowerManagerFilterer) FilterTaskSubmited(opts *bind.FilterOpts) (*WatchTowerManagerTaskSubmitedIterator, error) {

	logs, sub, err := _WatchTowerManager.contract.FilterLogs(opts, "TaskSubmited")
	if err != nil {
		return nil, err
	}
	return &WatchTowerManagerTaskSubmitedIterator{contract: _WatchTowerManager.contract, event: "TaskSubmited", logs: logs, sub: sub}, nil
}

// WatchTaskSubmited is a free log subscription operation binding the contract event 0xa12a0d6d059e62f1f1ff5e187dc823d3b39af2c2cd7bc56dfbbb720c283fa1b8.
//
// Solidity: event TaskSubmited(address operator, uint32 operatortaskNumber, uint32 lastBlockToBeServed)
func (_WatchTowerManager *WatchTowerManagerFilterer) WatchTaskSubmited(opts *bind.WatchOpts, sink chan<- *WatchTowerManagerTaskSubmited) (event.Subscription, error) {

	logs, sub, err := _WatchTowerManager.contract.WatchLogs(opts, "TaskSubmited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WatchTowerManagerTaskSubmited)
				if err := _WatchTowerManager.contract.UnpackLog(event, "TaskSubmited", log); err != nil {
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

// ParseTaskSubmited is a log parse operation binding the contract event 0xa12a0d6d059e62f1f1ff5e187dc823d3b39af2c2cd7bc56dfbbb720c283fa1b8.
//
// Solidity: event TaskSubmited(address operator, uint32 operatortaskNumber, uint32 lastBlockToBeServed)
func (_WatchTowerManager *WatchTowerManagerFilterer) ParseTaskSubmited(log types.Log) (*WatchTowerManagerTaskSubmited, error) {
	event := new(WatchTowerManagerTaskSubmited)
	if err := _WatchTowerManager.contract.UnpackLog(event, "TaskSubmited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
