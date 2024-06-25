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

// AddressesMetaData contains all meta data concerning the Addresses contract.
var AddressesMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"delegationManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegationManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_slasherAddress\",\"type\":\"address\"}],\"name\":\"initializeAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slasherAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AddressesABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressesMetaData.ABI instead.
var AddressesABI = AddressesMetaData.ABI

// Addresses is an auto generated Go binding around an Ethereum contract.
type Addresses struct {
	AddressesCaller     // Read-only binding to the contract
	AddressesTransactor // Write-only binding to the contract
	AddressesFilterer   // Log filterer for contract events
}

// AddressesCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressesSession struct {
	Contract     *Addresses        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressesCallerSession struct {
	Contract *AddressesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AddressesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressesTransactorSession struct {
	Contract     *AddressesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AddressesRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressesRaw struct {
	Contract *Addresses // Generic contract binding to access the raw methods on
}

// AddressesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressesCallerRaw struct {
	Contract *AddressesCaller // Generic read-only contract binding to access the raw methods on
}

// AddressesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressesTransactorRaw struct {
	Contract *AddressesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddresses creates a new instance of Addresses, bound to a specific deployed contract.
func NewAddresses(address common.Address, backend bind.ContractBackend) (*Addresses, error) {
	contract, err := bindAddresses(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Addresses{AddressesCaller: AddressesCaller{contract: contract}, AddressesTransactor: AddressesTransactor{contract: contract}, AddressesFilterer: AddressesFilterer{contract: contract}}, nil
}

// NewAddressesCaller creates a new read-only instance of Addresses, bound to a specific deployed contract.
func NewAddressesCaller(address common.Address, caller bind.ContractCaller) (*AddressesCaller, error) {
	contract, err := bindAddresses(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressesCaller{contract: contract}, nil
}

// NewAddressesTransactor creates a new write-only instance of Addresses, bound to a specific deployed contract.
func NewAddressesTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressesTransactor, error) {
	contract, err := bindAddresses(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressesTransactor{contract: contract}, nil
}

// NewAddressesFilterer creates a new log filterer instance of Addresses, bound to a specific deployed contract.
func NewAddressesFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressesFilterer, error) {
	contract, err := bindAddresses(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressesFilterer{contract: contract}, nil
}

// bindAddresses binds a generic wrapper to an already deployed contract.
func bindAddresses(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Addresses *AddressesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Addresses.Contract.AddressesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Addresses *AddressesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Addresses.Contract.AddressesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Addresses *AddressesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Addresses.Contract.AddressesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Addresses *AddressesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Addresses.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Addresses *AddressesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Addresses.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Addresses *AddressesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Addresses.Contract.contract.Transact(opts, method, params...)
}

// DelegationManagerAddress is a free data retrieval call binding the contract method 0xb15e6689.
//
// Solidity: function delegationManagerAddress() view returns(address)
func (_Addresses *AddressesCaller) DelegationManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Addresses.contract.Call(opts, &out, "delegationManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManagerAddress is a free data retrieval call binding the contract method 0xb15e6689.
//
// Solidity: function delegationManagerAddress() view returns(address)
func (_Addresses *AddressesSession) DelegationManagerAddress() (common.Address, error) {
	return _Addresses.Contract.DelegationManagerAddress(&_Addresses.CallOpts)
}

// DelegationManagerAddress is a free data retrieval call binding the contract method 0xb15e6689.
//
// Solidity: function delegationManagerAddress() view returns(address)
func (_Addresses *AddressesCallerSession) DelegationManagerAddress() (common.Address, error) {
	return _Addresses.Contract.DelegationManagerAddress(&_Addresses.CallOpts)
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() view returns(address)
func (_Addresses *AddressesCaller) SlasherAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Addresses.contract.Call(opts, &out, "slasherAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() view returns(address)
func (_Addresses *AddressesSession) SlasherAddress() (common.Address, error) {
	return _Addresses.Contract.SlasherAddress(&_Addresses.CallOpts)
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() view returns(address)
func (_Addresses *AddressesCallerSession) SlasherAddress() (common.Address, error) {
	return _Addresses.Contract.SlasherAddress(&_Addresses.CallOpts)
}

// InitializeAddresses is a paid mutator transaction binding the contract method 0x30b7bdad.
//
// Solidity: function initializeAddresses(address _delegationManagerAddress, address _slasherAddress) returns()
func (_Addresses *AddressesTransactor) InitializeAddresses(opts *bind.TransactOpts, _delegationManagerAddress common.Address, _slasherAddress common.Address) (*types.Transaction, error) {
	return _Addresses.contract.Transact(opts, "initializeAddresses", _delegationManagerAddress, _slasherAddress)
}

// InitializeAddresses is a paid mutator transaction binding the contract method 0x30b7bdad.
//
// Solidity: function initializeAddresses(address _delegationManagerAddress, address _slasherAddress) returns()
func (_Addresses *AddressesSession) InitializeAddresses(_delegationManagerAddress common.Address, _slasherAddress common.Address) (*types.Transaction, error) {
	return _Addresses.Contract.InitializeAddresses(&_Addresses.TransactOpts, _delegationManagerAddress, _slasherAddress)
}

// InitializeAddresses is a paid mutator transaction binding the contract method 0x30b7bdad.
//
// Solidity: function initializeAddresses(address _delegationManagerAddress, address _slasherAddress) returns()
func (_Addresses *AddressesTransactorSession) InitializeAddresses(_delegationManagerAddress common.Address, _slasherAddress common.Address) (*types.Transaction, error) {
	return _Addresses.Contract.InitializeAddresses(&_Addresses.TransactOpts, _delegationManagerAddress, _slasherAddress)
}

// AddressesInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Addresses contract.
type AddressesInitializedIterator struct {
	Event *AddressesInitialized // Event containing the contract specifics and raw log

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
func (it *AddressesInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressesInitialized)
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
		it.Event = new(AddressesInitialized)
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
func (it *AddressesInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressesInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressesInitialized represents a Initialized event raised by the Addresses contract.
type AddressesInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Addresses *AddressesFilterer) FilterInitialized(opts *bind.FilterOpts) (*AddressesInitializedIterator, error) {

	logs, sub, err := _Addresses.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AddressesInitializedIterator{contract: _Addresses.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Addresses *AddressesFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AddressesInitialized) (event.Subscription, error) {

	logs, sub, err := _Addresses.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressesInitialized)
				if err := _Addresses.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Addresses *AddressesFilterer) ParseInitialized(log types.Log) (*AddressesInitialized, error) {
	event := new(AddressesInitialized)
	if err := _Addresses.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
