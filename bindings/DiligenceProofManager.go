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

// DiligenceProofManagerMetaData contains all meta data concerning the DiligenceProofManager contract.
var DiligenceProofManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"logMessage\",\"type\":\"string\"}],\"name\":\"LogDebugging\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_signatureProofOfDiligence\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalBounty\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"NewPODBountyClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalBounty\",\"type\":\"uint256\"}],\"name\":\"NewPODBountyInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_signatureProofOfDiligence\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalBounty\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"NewPOIBountyClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalBounty\",\"type\":\"uint256\"}],\"name\":\"NewPOIBountyInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"}],\"name\":\"getPODBountyAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"getPODMinerStateRoots\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"}],\"name\":\"getPOIBountyAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"getPOIMinerStateRoots\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"}],\"name\":\"getRangeForChainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOperatorRegistry\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"contractIL2ChainMapping\",\"name\":\"_l2ChainMapping\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ChainMapping\",\"outputs\":[{\"internalType\":\"contractIL2ChainMapping\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIOperatorRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIL2ChainMapping\",\"name\":\"_l2ChainMapping\",\"type\":\"address\"}],\"name\":\"setL2ChainMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOperatorRegistry\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setOperatorRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalBounty\",\"type\":\"uint256\"}],\"name\":\"setPODBounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalBounty\",\"type\":\"uint256\"}],\"name\":\"setPOIBounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proofOfDiligence\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signatureProofOfDiligence\",\"type\":\"bytes\"}],\"name\":\"submitPODProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proofOfInclusion\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signatureProofOfInclusion\",\"type\":\"bytes\"}],\"name\":\"submitPOIProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_range\",\"type\":\"uint256\"}],\"name\":\"updateRangeForChainID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// DiligenceProofManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DiligenceProofManagerMetaData.ABI instead.
var DiligenceProofManagerABI = DiligenceProofManagerMetaData.ABI

// DiligenceProofManager is an auto generated Go binding around an Ethereum contract.
type DiligenceProofManager struct {
	DiligenceProofManagerCaller     // Read-only binding to the contract
	DiligenceProofManagerTransactor // Write-only binding to the contract
	DiligenceProofManagerFilterer   // Log filterer for contract events
}

// DiligenceProofManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiligenceProofManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiligenceProofManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiligenceProofManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiligenceProofManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiligenceProofManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiligenceProofManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiligenceProofManagerSession struct {
	Contract     *DiligenceProofManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DiligenceProofManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiligenceProofManagerCallerSession struct {
	Contract *DiligenceProofManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// DiligenceProofManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiligenceProofManagerTransactorSession struct {
	Contract     *DiligenceProofManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// DiligenceProofManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiligenceProofManagerRaw struct {
	Contract *DiligenceProofManager // Generic contract binding to access the raw methods on
}

// DiligenceProofManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiligenceProofManagerCallerRaw struct {
	Contract *DiligenceProofManagerCaller // Generic read-only contract binding to access the raw methods on
}

// DiligenceProofManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiligenceProofManagerTransactorRaw struct {
	Contract *DiligenceProofManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiligenceProofManager creates a new instance of DiligenceProofManager, bound to a specific deployed contract.
func NewDiligenceProofManager(address common.Address, backend bind.ContractBackend) (*DiligenceProofManager, error) {
	contract, err := bindDiligenceProofManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManager{DiligenceProofManagerCaller: DiligenceProofManagerCaller{contract: contract}, DiligenceProofManagerTransactor: DiligenceProofManagerTransactor{contract: contract}, DiligenceProofManagerFilterer: DiligenceProofManagerFilterer{contract: contract}}, nil
}

// NewDiligenceProofManagerCaller creates a new read-only instance of DiligenceProofManager, bound to a specific deployed contract.
func NewDiligenceProofManagerCaller(address common.Address, caller bind.ContractCaller) (*DiligenceProofManagerCaller, error) {
	contract, err := bindDiligenceProofManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManagerCaller{contract: contract}, nil
}

// NewDiligenceProofManagerTransactor creates a new write-only instance of DiligenceProofManager, bound to a specific deployed contract.
func NewDiligenceProofManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DiligenceProofManagerTransactor, error) {
	contract, err := bindDiligenceProofManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManagerTransactor{contract: contract}, nil
}

// NewDiligenceProofManagerFilterer creates a new log filterer instance of DiligenceProofManager, bound to a specific deployed contract.
func NewDiligenceProofManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DiligenceProofManagerFilterer, error) {
	contract, err := bindDiligenceProofManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManagerFilterer{contract: contract}, nil
}

// bindDiligenceProofManager binds a generic wrapper to an already deployed contract.
func bindDiligenceProofManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DiligenceProofManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiligenceProofManager *DiligenceProofManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiligenceProofManager.Contract.DiligenceProofManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiligenceProofManager *DiligenceProofManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.DiligenceProofManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiligenceProofManager *DiligenceProofManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.DiligenceProofManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiligenceProofManager *DiligenceProofManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiligenceProofManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiligenceProofManager *DiligenceProofManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiligenceProofManager *DiligenceProofManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.contract.Transact(opts, method, params...)
}

// GetPODBountyAmount is a free data retrieval call binding the contract method 0xbfe3df45.
//
// Solidity: function getPODBountyAmount(uint256 _chainID) view returns(uint256)
func (_DiligenceProofManager *DiligenceProofManagerCaller) GetPODBountyAmount(opts *bind.CallOpts, _chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DiligenceProofManager.contract.Call(opts, &out, "getPODBountyAmount", _chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPODBountyAmount is a free data retrieval call binding the contract method 0xbfe3df45.
//
// Solidity: function getPODBountyAmount(uint256 _chainID) view returns(uint256)
func (_DiligenceProofManager *DiligenceProofManagerSession) GetPODBountyAmount(_chainID *big.Int) (*big.Int, error) {
	return _DiligenceProofManager.Contract.GetPODBountyAmount(&_DiligenceProofManager.CallOpts, _chainID)
}

// GetPODBountyAmount is a free data retrieval call binding the contract method 0xbfe3df45.
//
// Solidity: function getPODBountyAmount(uint256 _chainID) view returns(uint256)
func (_DiligenceProofManager *DiligenceProofManagerCallerSession) GetPODBountyAmount(_chainID *big.Int) (*big.Int, error) {
	return _DiligenceProofManager.Contract.GetPODBountyAmount(&_DiligenceProofManager.CallOpts, _chainID)
}

// GetPODMinerStateRoots is a free data retrieval call binding the contract method 0x14bb1f11.
//
// Solidity: function getPODMinerStateRoots(uint256 _chainID, uint256 _l2BlockNumber, address _miner) view returns(bytes)
func (_DiligenceProofManager *DiligenceProofManagerCaller) GetPODMinerStateRoots(opts *bind.CallOpts, _chainID *big.Int, _l2BlockNumber *big.Int, _miner common.Address) ([]byte, error) {
	var out []interface{}
	err := _DiligenceProofManager.contract.Call(opts, &out, "getPODMinerStateRoots", _chainID, _l2BlockNumber, _miner)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPODMinerStateRoots is a free data retrieval call binding the contract method 0x14bb1f11.
//
// Solidity: function getPODMinerStateRoots(uint256 _chainID, uint256 _l2BlockNumber, address _miner) view returns(bytes)
func (_DiligenceProofManager *DiligenceProofManagerSession) GetPODMinerStateRoots(_chainID *big.Int, _l2BlockNumber *big.Int, _miner common.Address) ([]byte, error) {
	return _DiligenceProofManager.Contract.GetPODMinerStateRoots(&_DiligenceProofManager.CallOpts, _chainID, _l2BlockNumber, _miner)
}

// GetPODMinerStateRoots is a free data retrieval call binding the contract method 0x14bb1f11.
//
// Solidity: function getPODMinerStateRoots(uint256 _chainID, uint256 _l2BlockNumber, address _miner) view returns(bytes)
func (_DiligenceProofManager *DiligenceProofManagerCallerSession) GetPODMinerStateRoots(_chainID *big.Int, _l2BlockNumber *big.Int, _miner common.Address) ([]byte, error) {
	return _DiligenceProofManager.Contract.GetPODMinerStateRoots(&_DiligenceProofManager.CallOpts, _chainID, _l2BlockNumber, _miner)
}

// GetPOIBountyAmount is a free data retrieval call binding the contract method 0xb0cdaa56.
//
// Solidity: function getPOIBountyAmount(uint256 _chainID) view returns(uint256)
func (_DiligenceProofManager *DiligenceProofManagerCaller) GetPOIBountyAmount(opts *bind.CallOpts, _chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DiligenceProofManager.contract.Call(opts, &out, "getPOIBountyAmount", _chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPOIBountyAmount is a free data retrieval call binding the contract method 0xb0cdaa56.
//
// Solidity: function getPOIBountyAmount(uint256 _chainID) view returns(uint256)
func (_DiligenceProofManager *DiligenceProofManagerSession) GetPOIBountyAmount(_chainID *big.Int) (*big.Int, error) {
	return _DiligenceProofManager.Contract.GetPOIBountyAmount(&_DiligenceProofManager.CallOpts, _chainID)
}

// GetPOIBountyAmount is a free data retrieval call binding the contract method 0xb0cdaa56.
//
// Solidity: function getPOIBountyAmount(uint256 _chainID) view returns(uint256)
func (_DiligenceProofManager *DiligenceProofManagerCallerSession) GetPOIBountyAmount(_chainID *big.Int) (*big.Int, error) {
	return _DiligenceProofManager.Contract.GetPOIBountyAmount(&_DiligenceProofManager.CallOpts, _chainID)
}

// GetPOIMinerStateRoots is a free data retrieval call binding the contract method 0xc285fcb2.
//
// Solidity: function getPOIMinerStateRoots(uint256 _chainID, uint256 _l2BlockNumber, address _miner) view returns(bytes)
func (_DiligenceProofManager *DiligenceProofManagerCaller) GetPOIMinerStateRoots(opts *bind.CallOpts, _chainID *big.Int, _l2BlockNumber *big.Int, _miner common.Address) ([]byte, error) {
	var out []interface{}
	err := _DiligenceProofManager.contract.Call(opts, &out, "getPOIMinerStateRoots", _chainID, _l2BlockNumber, _miner)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPOIMinerStateRoots is a free data retrieval call binding the contract method 0xc285fcb2.
//
// Solidity: function getPOIMinerStateRoots(uint256 _chainID, uint256 _l2BlockNumber, address _miner) view returns(bytes)
func (_DiligenceProofManager *DiligenceProofManagerSession) GetPOIMinerStateRoots(_chainID *big.Int, _l2BlockNumber *big.Int, _miner common.Address) ([]byte, error) {
	return _DiligenceProofManager.Contract.GetPOIMinerStateRoots(&_DiligenceProofManager.CallOpts, _chainID, _l2BlockNumber, _miner)
}

// GetPOIMinerStateRoots is a free data retrieval call binding the contract method 0xc285fcb2.
//
// Solidity: function getPOIMinerStateRoots(uint256 _chainID, uint256 _l2BlockNumber, address _miner) view returns(bytes)
func (_DiligenceProofManager *DiligenceProofManagerCallerSession) GetPOIMinerStateRoots(_chainID *big.Int, _l2BlockNumber *big.Int, _miner common.Address) ([]byte, error) {
	return _DiligenceProofManager.Contract.GetPOIMinerStateRoots(&_DiligenceProofManager.CallOpts, _chainID, _l2BlockNumber, _miner)
}

// GetRangeForChainID is a free data retrieval call binding the contract method 0xe61e1df7.
//
// Solidity: function getRangeForChainID(uint256 _chainID) view returns(uint256)
func (_DiligenceProofManager *DiligenceProofManagerCaller) GetRangeForChainID(opts *bind.CallOpts, _chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DiligenceProofManager.contract.Call(opts, &out, "getRangeForChainID", _chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRangeForChainID is a free data retrieval call binding the contract method 0xe61e1df7.
//
// Solidity: function getRangeForChainID(uint256 _chainID) view returns(uint256)
func (_DiligenceProofManager *DiligenceProofManagerSession) GetRangeForChainID(_chainID *big.Int) (*big.Int, error) {
	return _DiligenceProofManager.Contract.GetRangeForChainID(&_DiligenceProofManager.CallOpts, _chainID)
}

// GetRangeForChainID is a free data retrieval call binding the contract method 0xe61e1df7.
//
// Solidity: function getRangeForChainID(uint256 _chainID) view returns(uint256)
func (_DiligenceProofManager *DiligenceProofManagerCallerSession) GetRangeForChainID(_chainID *big.Int) (*big.Int, error) {
	return _DiligenceProofManager.Contract.GetRangeForChainID(&_DiligenceProofManager.CallOpts, _chainID)
}

// L2ChainMapping is a free data retrieval call binding the contract method 0x820b6f39.
//
// Solidity: function l2ChainMapping() view returns(address)
func (_DiligenceProofManager *DiligenceProofManagerCaller) L2ChainMapping(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DiligenceProofManager.contract.Call(opts, &out, "l2ChainMapping")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ChainMapping is a free data retrieval call binding the contract method 0x820b6f39.
//
// Solidity: function l2ChainMapping() view returns(address)
func (_DiligenceProofManager *DiligenceProofManagerSession) L2ChainMapping() (common.Address, error) {
	return _DiligenceProofManager.Contract.L2ChainMapping(&_DiligenceProofManager.CallOpts)
}

// L2ChainMapping is a free data retrieval call binding the contract method 0x820b6f39.
//
// Solidity: function l2ChainMapping() view returns(address)
func (_DiligenceProofManager *DiligenceProofManagerCallerSession) L2ChainMapping() (common.Address, error) {
	return _DiligenceProofManager.Contract.L2ChainMapping(&_DiligenceProofManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DiligenceProofManager *DiligenceProofManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DiligenceProofManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DiligenceProofManager *DiligenceProofManagerSession) Owner() (common.Address, error) {
	return _DiligenceProofManager.Contract.Owner(&_DiligenceProofManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DiligenceProofManager *DiligenceProofManagerCallerSession) Owner() (common.Address, error) {
	return _DiligenceProofManager.Contract.Owner(&_DiligenceProofManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DiligenceProofManager *DiligenceProofManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DiligenceProofManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DiligenceProofManager *DiligenceProofManagerSession) Paused() (bool, error) {
	return _DiligenceProofManager.Contract.Paused(&_DiligenceProofManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DiligenceProofManager *DiligenceProofManagerCallerSession) Paused() (bool, error) {
	return _DiligenceProofManager.Contract.Paused(&_DiligenceProofManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_DiligenceProofManager *DiligenceProofManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DiligenceProofManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_DiligenceProofManager *DiligenceProofManagerSession) ProxiableUUID() ([32]byte, error) {
	return _DiligenceProofManager.Contract.ProxiableUUID(&_DiligenceProofManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_DiligenceProofManager *DiligenceProofManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _DiligenceProofManager.Contract.ProxiableUUID(&_DiligenceProofManager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_DiligenceProofManager *DiligenceProofManagerCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DiligenceProofManager.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_DiligenceProofManager *DiligenceProofManagerSession) Registry() (common.Address, error) {
	return _DiligenceProofManager.Contract.Registry(&_DiligenceProofManager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_DiligenceProofManager *DiligenceProofManagerCallerSession) Registry() (common.Address, error) {
	return _DiligenceProofManager.Contract.Registry(&_DiligenceProofManager.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registry, address _l2ChainMapping) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _l2ChainMapping common.Address) (*types.Transaction, error) {
	return _DiligenceProofManager.contract.Transact(opts, "initialize", _registry, _l2ChainMapping)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registry, address _l2ChainMapping) returns()
func (_DiligenceProofManager *DiligenceProofManagerSession) Initialize(_registry common.Address, _l2ChainMapping common.Address) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.Initialize(&_DiligenceProofManager.TransactOpts, _registry, _l2ChainMapping)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registry, address _l2ChainMapping) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactorSession) Initialize(_registry common.Address, _l2ChainMapping common.Address) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.Initialize(&_DiligenceProofManager.TransactOpts, _registry, _l2ChainMapping)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiligenceProofManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DiligenceProofManager *DiligenceProofManagerSession) Pause() (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.Pause(&_DiligenceProofManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.Pause(&_DiligenceProofManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiligenceProofManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DiligenceProofManager *DiligenceProofManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.RenounceOwnership(&_DiligenceProofManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.RenounceOwnership(&_DiligenceProofManager.TransactOpts)
}

// SetL2ChainMapping is a paid mutator transaction binding the contract method 0x8d134f8d.
//
// Solidity: function setL2ChainMapping(address _l2ChainMapping) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactor) SetL2ChainMapping(opts *bind.TransactOpts, _l2ChainMapping common.Address) (*types.Transaction, error) {
	return _DiligenceProofManager.contract.Transact(opts, "setL2ChainMapping", _l2ChainMapping)
}

// SetL2ChainMapping is a paid mutator transaction binding the contract method 0x8d134f8d.
//
// Solidity: function setL2ChainMapping(address _l2ChainMapping) returns()
func (_DiligenceProofManager *DiligenceProofManagerSession) SetL2ChainMapping(_l2ChainMapping common.Address) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.SetL2ChainMapping(&_DiligenceProofManager.TransactOpts, _l2ChainMapping)
}

// SetL2ChainMapping is a paid mutator transaction binding the contract method 0x8d134f8d.
//
// Solidity: function setL2ChainMapping(address _l2ChainMapping) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactorSession) SetL2ChainMapping(_l2ChainMapping common.Address) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.SetL2ChainMapping(&_DiligenceProofManager.TransactOpts, _l2ChainMapping)
}

// SetOperatorRegistry is a paid mutator transaction binding the contract method 0x9d28fb86.
//
// Solidity: function setOperatorRegistry(address _registry) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactor) SetOperatorRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _DiligenceProofManager.contract.Transact(opts, "setOperatorRegistry", _registry)
}

// SetOperatorRegistry is a paid mutator transaction binding the contract method 0x9d28fb86.
//
// Solidity: function setOperatorRegistry(address _registry) returns()
func (_DiligenceProofManager *DiligenceProofManagerSession) SetOperatorRegistry(_registry common.Address) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.SetOperatorRegistry(&_DiligenceProofManager.TransactOpts, _registry)
}

// SetOperatorRegistry is a paid mutator transaction binding the contract method 0x9d28fb86.
//
// Solidity: function setOperatorRegistry(address _registry) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactorSession) SetOperatorRegistry(_registry common.Address) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.SetOperatorRegistry(&_DiligenceProofManager.TransactOpts, _registry)
}

// SetPODBounty is a paid mutator transaction binding the contract method 0x4e8d8889.
//
// Solidity: function setPODBounty(uint256 _chainID, uint256 _totalBounty) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactor) SetPODBounty(opts *bind.TransactOpts, _chainID *big.Int, _totalBounty *big.Int) (*types.Transaction, error) {
	return _DiligenceProofManager.contract.Transact(opts, "setPODBounty", _chainID, _totalBounty)
}

// SetPODBounty is a paid mutator transaction binding the contract method 0x4e8d8889.
//
// Solidity: function setPODBounty(uint256 _chainID, uint256 _totalBounty) returns()
func (_DiligenceProofManager *DiligenceProofManagerSession) SetPODBounty(_chainID *big.Int, _totalBounty *big.Int) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.SetPODBounty(&_DiligenceProofManager.TransactOpts, _chainID, _totalBounty)
}

// SetPODBounty is a paid mutator transaction binding the contract method 0x4e8d8889.
//
// Solidity: function setPODBounty(uint256 _chainID, uint256 _totalBounty) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactorSession) SetPODBounty(_chainID *big.Int, _totalBounty *big.Int) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.SetPODBounty(&_DiligenceProofManager.TransactOpts, _chainID, _totalBounty)
}

// SetPOIBounty is a paid mutator transaction binding the contract method 0xfbc1a55f.
//
// Solidity: function setPOIBounty(uint256 _chainID, uint256 _totalBounty) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactor) SetPOIBounty(opts *bind.TransactOpts, _chainID *big.Int, _totalBounty *big.Int) (*types.Transaction, error) {
	return _DiligenceProofManager.contract.Transact(opts, "setPOIBounty", _chainID, _totalBounty)
}

// SetPOIBounty is a paid mutator transaction binding the contract method 0xfbc1a55f.
//
// Solidity: function setPOIBounty(uint256 _chainID, uint256 _totalBounty) returns()
func (_DiligenceProofManager *DiligenceProofManagerSession) SetPOIBounty(_chainID *big.Int, _totalBounty *big.Int) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.SetPOIBounty(&_DiligenceProofManager.TransactOpts, _chainID, _totalBounty)
}

// SetPOIBounty is a paid mutator transaction binding the contract method 0xfbc1a55f.
//
// Solidity: function setPOIBounty(uint256 _chainID, uint256 _totalBounty) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactorSession) SetPOIBounty(_chainID *big.Int, _totalBounty *big.Int) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.SetPOIBounty(&_DiligenceProofManager.TransactOpts, _chainID, _totalBounty)
}

// SubmitPODProof is a paid mutator transaction binding the contract method 0xdabe0314.
//
// Solidity: function submitPODProof(uint256 _chainID, uint256 _l2BlockNumber, bytes _proofOfDiligence, bytes _signatureProofOfDiligence) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactor) SubmitPODProof(opts *bind.TransactOpts, _chainID *big.Int, _l2BlockNumber *big.Int, _proofOfDiligence []byte, _signatureProofOfDiligence []byte) (*types.Transaction, error) {
	return _DiligenceProofManager.contract.Transact(opts, "submitPODProof", _chainID, _l2BlockNumber, _proofOfDiligence, _signatureProofOfDiligence)
}

// SubmitPODProof is a paid mutator transaction binding the contract method 0xdabe0314.
//
// Solidity: function submitPODProof(uint256 _chainID, uint256 _l2BlockNumber, bytes _proofOfDiligence, bytes _signatureProofOfDiligence) returns()
func (_DiligenceProofManager *DiligenceProofManagerSession) SubmitPODProof(_chainID *big.Int, _l2BlockNumber *big.Int, _proofOfDiligence []byte, _signatureProofOfDiligence []byte) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.SubmitPODProof(&_DiligenceProofManager.TransactOpts, _chainID, _l2BlockNumber, _proofOfDiligence, _signatureProofOfDiligence)
}

// SubmitPODProof is a paid mutator transaction binding the contract method 0xdabe0314.
//
// Solidity: function submitPODProof(uint256 _chainID, uint256 _l2BlockNumber, bytes _proofOfDiligence, bytes _signatureProofOfDiligence) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactorSession) SubmitPODProof(_chainID *big.Int, _l2BlockNumber *big.Int, _proofOfDiligence []byte, _signatureProofOfDiligence []byte) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.SubmitPODProof(&_DiligenceProofManager.TransactOpts, _chainID, _l2BlockNumber, _proofOfDiligence, _signatureProofOfDiligence)
}

// SubmitPOIProof is a paid mutator transaction binding the contract method 0xa4c933c3.
//
// Solidity: function submitPOIProof(uint256 _chainID, uint256 _l2BlockNumber, bytes _proofOfInclusion, bytes _signatureProofOfInclusion) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactor) SubmitPOIProof(opts *bind.TransactOpts, _chainID *big.Int, _l2BlockNumber *big.Int, _proofOfInclusion []byte, _signatureProofOfInclusion []byte) (*types.Transaction, error) {
	return _DiligenceProofManager.contract.Transact(opts, "submitPOIProof", _chainID, _l2BlockNumber, _proofOfInclusion, _signatureProofOfInclusion)
}

// SubmitPOIProof is a paid mutator transaction binding the contract method 0xa4c933c3.
//
// Solidity: function submitPOIProof(uint256 _chainID, uint256 _l2BlockNumber, bytes _proofOfInclusion, bytes _signatureProofOfInclusion) returns()
func (_DiligenceProofManager *DiligenceProofManagerSession) SubmitPOIProof(_chainID *big.Int, _l2BlockNumber *big.Int, _proofOfInclusion []byte, _signatureProofOfInclusion []byte) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.SubmitPOIProof(&_DiligenceProofManager.TransactOpts, _chainID, _l2BlockNumber, _proofOfInclusion, _signatureProofOfInclusion)
}

// SubmitPOIProof is a paid mutator transaction binding the contract method 0xa4c933c3.
//
// Solidity: function submitPOIProof(uint256 _chainID, uint256 _l2BlockNumber, bytes _proofOfInclusion, bytes _signatureProofOfInclusion) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactorSession) SubmitPOIProof(_chainID *big.Int, _l2BlockNumber *big.Int, _proofOfInclusion []byte, _signatureProofOfInclusion []byte) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.SubmitPOIProof(&_DiligenceProofManager.TransactOpts, _chainID, _l2BlockNumber, _proofOfInclusion, _signatureProofOfInclusion)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DiligenceProofManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DiligenceProofManager *DiligenceProofManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.TransferOwnership(&_DiligenceProofManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.TransferOwnership(&_DiligenceProofManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiligenceProofManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DiligenceProofManager *DiligenceProofManagerSession) Unpause() (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.Unpause(&_DiligenceProofManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.Unpause(&_DiligenceProofManager.TransactOpts)
}

// UpdateRangeForChainID is a paid mutator transaction binding the contract method 0xe9fb9254.
//
// Solidity: function updateRangeForChainID(uint256 _chainID, uint256 _range) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactor) UpdateRangeForChainID(opts *bind.TransactOpts, _chainID *big.Int, _range *big.Int) (*types.Transaction, error) {
	return _DiligenceProofManager.contract.Transact(opts, "updateRangeForChainID", _chainID, _range)
}

// UpdateRangeForChainID is a paid mutator transaction binding the contract method 0xe9fb9254.
//
// Solidity: function updateRangeForChainID(uint256 _chainID, uint256 _range) returns()
func (_DiligenceProofManager *DiligenceProofManagerSession) UpdateRangeForChainID(_chainID *big.Int, _range *big.Int) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.UpdateRangeForChainID(&_DiligenceProofManager.TransactOpts, _chainID, _range)
}

// UpdateRangeForChainID is a paid mutator transaction binding the contract method 0xe9fb9254.
//
// Solidity: function updateRangeForChainID(uint256 _chainID, uint256 _range) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactorSession) UpdateRangeForChainID(_chainID *big.Int, _range *big.Int) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.UpdateRangeForChainID(&_DiligenceProofManager.TransactOpts, _chainID, _range)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _DiligenceProofManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_DiligenceProofManager *DiligenceProofManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.UpgradeTo(&_DiligenceProofManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.UpgradeTo(&_DiligenceProofManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _DiligenceProofManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_DiligenceProofManager *DiligenceProofManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.UpgradeToAndCall(&_DiligenceProofManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_DiligenceProofManager *DiligenceProofManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _DiligenceProofManager.Contract.UpgradeToAndCall(&_DiligenceProofManager.TransactOpts, newImplementation, data)
}

// DiligenceProofManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the DiligenceProofManager contract.
type DiligenceProofManagerAdminChangedIterator struct {
	Event *DiligenceProofManagerAdminChanged // Event containing the contract specifics and raw log

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
func (it *DiligenceProofManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiligenceProofManagerAdminChanged)
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
		it.Event = new(DiligenceProofManagerAdminChanged)
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
func (it *DiligenceProofManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiligenceProofManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiligenceProofManagerAdminChanged represents a AdminChanged event raised by the DiligenceProofManager contract.
type DiligenceProofManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*DiligenceProofManagerAdminChangedIterator, error) {

	logs, sub, err := _DiligenceProofManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManagerAdminChangedIterator{contract: _DiligenceProofManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *DiligenceProofManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _DiligenceProofManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiligenceProofManagerAdminChanged)
				if err := _DiligenceProofManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_DiligenceProofManager *DiligenceProofManagerFilterer) ParseAdminChanged(log types.Log) (*DiligenceProofManagerAdminChanged, error) {
	event := new(DiligenceProofManagerAdminChanged)
	if err := _DiligenceProofManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiligenceProofManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the DiligenceProofManager contract.
type DiligenceProofManagerBeaconUpgradedIterator struct {
	Event *DiligenceProofManagerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *DiligenceProofManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiligenceProofManagerBeaconUpgraded)
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
		it.Event = new(DiligenceProofManagerBeaconUpgraded)
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
func (it *DiligenceProofManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiligenceProofManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiligenceProofManagerBeaconUpgraded represents a BeaconUpgraded event raised by the DiligenceProofManager contract.
type DiligenceProofManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*DiligenceProofManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _DiligenceProofManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManagerBeaconUpgradedIterator{contract: _DiligenceProofManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *DiligenceProofManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _DiligenceProofManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiligenceProofManagerBeaconUpgraded)
				if err := _DiligenceProofManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_DiligenceProofManager *DiligenceProofManagerFilterer) ParseBeaconUpgraded(log types.Log) (*DiligenceProofManagerBeaconUpgraded, error) {
	event := new(DiligenceProofManagerBeaconUpgraded)
	if err := _DiligenceProofManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiligenceProofManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DiligenceProofManager contract.
type DiligenceProofManagerInitializedIterator struct {
	Event *DiligenceProofManagerInitialized // Event containing the contract specifics and raw log

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
func (it *DiligenceProofManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiligenceProofManagerInitialized)
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
		it.Event = new(DiligenceProofManagerInitialized)
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
func (it *DiligenceProofManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiligenceProofManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiligenceProofManagerInitialized represents a Initialized event raised by the DiligenceProofManager contract.
type DiligenceProofManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*DiligenceProofManagerInitializedIterator, error) {

	logs, sub, err := _DiligenceProofManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManagerInitializedIterator{contract: _DiligenceProofManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DiligenceProofManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _DiligenceProofManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiligenceProofManagerInitialized)
				if err := _DiligenceProofManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DiligenceProofManager *DiligenceProofManagerFilterer) ParseInitialized(log types.Log) (*DiligenceProofManagerInitialized, error) {
	event := new(DiligenceProofManagerInitialized)
	if err := _DiligenceProofManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiligenceProofManagerLogDebuggingIterator is returned from FilterLogDebugging and is used to iterate over the raw logs and unpacked data for LogDebugging events raised by the DiligenceProofManager contract.
type DiligenceProofManagerLogDebuggingIterator struct {
	Event *DiligenceProofManagerLogDebugging // Event containing the contract specifics and raw log

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
func (it *DiligenceProofManagerLogDebuggingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiligenceProofManagerLogDebugging)
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
		it.Event = new(DiligenceProofManagerLogDebugging)
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
func (it *DiligenceProofManagerLogDebuggingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiligenceProofManagerLogDebuggingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiligenceProofManagerLogDebugging represents a LogDebugging event raised by the DiligenceProofManager contract.
type DiligenceProofManagerLogDebugging struct {
	LogMessage string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogDebugging is a free log retrieval operation binding the contract event 0x98eb5a2c29958dab5d1549c7d2ac634fc85a96d2467361831b8d4e379f7ff9ec.
//
// Solidity: event LogDebugging(string logMessage)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) FilterLogDebugging(opts *bind.FilterOpts) (*DiligenceProofManagerLogDebuggingIterator, error) {

	logs, sub, err := _DiligenceProofManager.contract.FilterLogs(opts, "LogDebugging")
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManagerLogDebuggingIterator{contract: _DiligenceProofManager.contract, event: "LogDebugging", logs: logs, sub: sub}, nil
}

// WatchLogDebugging is a free log subscription operation binding the contract event 0x98eb5a2c29958dab5d1549c7d2ac634fc85a96d2467361831b8d4e379f7ff9ec.
//
// Solidity: event LogDebugging(string logMessage)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) WatchLogDebugging(opts *bind.WatchOpts, sink chan<- *DiligenceProofManagerLogDebugging) (event.Subscription, error) {

	logs, sub, err := _DiligenceProofManager.contract.WatchLogs(opts, "LogDebugging")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiligenceProofManagerLogDebugging)
				if err := _DiligenceProofManager.contract.UnpackLog(event, "LogDebugging", log); err != nil {
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

// ParseLogDebugging is a log parse operation binding the contract event 0x98eb5a2c29958dab5d1549c7d2ac634fc85a96d2467361831b8d4e379f7ff9ec.
//
// Solidity: event LogDebugging(string logMessage)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) ParseLogDebugging(log types.Log) (*DiligenceProofManagerLogDebugging, error) {
	event := new(DiligenceProofManagerLogDebugging)
	if err := _DiligenceProofManager.contract.UnpackLog(event, "LogDebugging", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiligenceProofManagerNewPODBountyClaimedIterator is returned from FilterNewPODBountyClaimed and is used to iterate over the raw logs and unpacked data for NewPODBountyClaimed events raised by the DiligenceProofManager contract.
type DiligenceProofManagerNewPODBountyClaimedIterator struct {
	Event *DiligenceProofManagerNewPODBountyClaimed // Event containing the contract specifics and raw log

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
func (it *DiligenceProofManagerNewPODBountyClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiligenceProofManagerNewPODBountyClaimed)
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
		it.Event = new(DiligenceProofManagerNewPODBountyClaimed)
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
func (it *DiligenceProofManagerNewPODBountyClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiligenceProofManagerNewPODBountyClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiligenceProofManagerNewPODBountyClaimed represents a NewPODBountyClaimed event raised by the DiligenceProofManager contract.
type DiligenceProofManagerNewPODBountyClaimed struct {
	ChainID                   *big.Int
	L2BlockNumber             *big.Int
	SignatureProofOfDiligence []byte
	TotalBounty               *big.Int
	Miner                     common.Address
	Timestamp                 *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterNewPODBountyClaimed is a free log retrieval operation binding the contract event 0x4843b6b4ef4b45c687f6bb5269ab2c04eee6c7ec950cc7db9e8f51cbc8974c1d.
//
// Solidity: event NewPODBountyClaimed(uint256 indexed _chainID, uint256 indexed _l2BlockNumber, bytes _signatureProofOfDiligence, uint256 _totalBounty, address indexed _miner, uint256 _timestamp)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) FilterNewPODBountyClaimed(opts *bind.FilterOpts, _chainID []*big.Int, _l2BlockNumber []*big.Int, _miner []common.Address) (*DiligenceProofManagerNewPODBountyClaimedIterator, error) {

	var _chainIDRule []interface{}
	for _, _chainIDItem := range _chainID {
		_chainIDRule = append(_chainIDRule, _chainIDItem)
	}
	var _l2BlockNumberRule []interface{}
	for _, _l2BlockNumberItem := range _l2BlockNumber {
		_l2BlockNumberRule = append(_l2BlockNumberRule, _l2BlockNumberItem)
	}

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	logs, sub, err := _DiligenceProofManager.contract.FilterLogs(opts, "NewPODBountyClaimed", _chainIDRule, _l2BlockNumberRule, _minerRule)
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManagerNewPODBountyClaimedIterator{contract: _DiligenceProofManager.contract, event: "NewPODBountyClaimed", logs: logs, sub: sub}, nil
}

// WatchNewPODBountyClaimed is a free log subscription operation binding the contract event 0x4843b6b4ef4b45c687f6bb5269ab2c04eee6c7ec950cc7db9e8f51cbc8974c1d.
//
// Solidity: event NewPODBountyClaimed(uint256 indexed _chainID, uint256 indexed _l2BlockNumber, bytes _signatureProofOfDiligence, uint256 _totalBounty, address indexed _miner, uint256 _timestamp)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) WatchNewPODBountyClaimed(opts *bind.WatchOpts, sink chan<- *DiligenceProofManagerNewPODBountyClaimed, _chainID []*big.Int, _l2BlockNumber []*big.Int, _miner []common.Address) (event.Subscription, error) {

	var _chainIDRule []interface{}
	for _, _chainIDItem := range _chainID {
		_chainIDRule = append(_chainIDRule, _chainIDItem)
	}
	var _l2BlockNumberRule []interface{}
	for _, _l2BlockNumberItem := range _l2BlockNumber {
		_l2BlockNumberRule = append(_l2BlockNumberRule, _l2BlockNumberItem)
	}

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	logs, sub, err := _DiligenceProofManager.contract.WatchLogs(opts, "NewPODBountyClaimed", _chainIDRule, _l2BlockNumberRule, _minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiligenceProofManagerNewPODBountyClaimed)
				if err := _DiligenceProofManager.contract.UnpackLog(event, "NewPODBountyClaimed", log); err != nil {
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

// ParseNewPODBountyClaimed is a log parse operation binding the contract event 0x4843b6b4ef4b45c687f6bb5269ab2c04eee6c7ec950cc7db9e8f51cbc8974c1d.
//
// Solidity: event NewPODBountyClaimed(uint256 indexed _chainID, uint256 indexed _l2BlockNumber, bytes _signatureProofOfDiligence, uint256 _totalBounty, address indexed _miner, uint256 _timestamp)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) ParseNewPODBountyClaimed(log types.Log) (*DiligenceProofManagerNewPODBountyClaimed, error) {
	event := new(DiligenceProofManagerNewPODBountyClaimed)
	if err := _DiligenceProofManager.contract.UnpackLog(event, "NewPODBountyClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiligenceProofManagerNewPODBountyInitializedIterator is returned from FilterNewPODBountyInitialized and is used to iterate over the raw logs and unpacked data for NewPODBountyInitialized events raised by the DiligenceProofManager contract.
type DiligenceProofManagerNewPODBountyInitializedIterator struct {
	Event *DiligenceProofManagerNewPODBountyInitialized // Event containing the contract specifics and raw log

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
func (it *DiligenceProofManagerNewPODBountyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiligenceProofManagerNewPODBountyInitialized)
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
		it.Event = new(DiligenceProofManagerNewPODBountyInitialized)
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
func (it *DiligenceProofManagerNewPODBountyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiligenceProofManagerNewPODBountyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiligenceProofManagerNewPODBountyInitialized represents a NewPODBountyInitialized event raised by the DiligenceProofManager contract.
type DiligenceProofManagerNewPODBountyInitialized struct {
	ChainID     *big.Int
	TotalBounty *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewPODBountyInitialized is a free log retrieval operation binding the contract event 0x0f0c773c4ddf39017c86ae4f66b1085a36b6a2c36a821c57670d732b00a9eadd.
//
// Solidity: event NewPODBountyInitialized(uint256 _chainID, uint256 _totalBounty)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) FilterNewPODBountyInitialized(opts *bind.FilterOpts) (*DiligenceProofManagerNewPODBountyInitializedIterator, error) {

	logs, sub, err := _DiligenceProofManager.contract.FilterLogs(opts, "NewPODBountyInitialized")
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManagerNewPODBountyInitializedIterator{contract: _DiligenceProofManager.contract, event: "NewPODBountyInitialized", logs: logs, sub: sub}, nil
}

// WatchNewPODBountyInitialized is a free log subscription operation binding the contract event 0x0f0c773c4ddf39017c86ae4f66b1085a36b6a2c36a821c57670d732b00a9eadd.
//
// Solidity: event NewPODBountyInitialized(uint256 _chainID, uint256 _totalBounty)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) WatchNewPODBountyInitialized(opts *bind.WatchOpts, sink chan<- *DiligenceProofManagerNewPODBountyInitialized) (event.Subscription, error) {

	logs, sub, err := _DiligenceProofManager.contract.WatchLogs(opts, "NewPODBountyInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiligenceProofManagerNewPODBountyInitialized)
				if err := _DiligenceProofManager.contract.UnpackLog(event, "NewPODBountyInitialized", log); err != nil {
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

// ParseNewPODBountyInitialized is a log parse operation binding the contract event 0x0f0c773c4ddf39017c86ae4f66b1085a36b6a2c36a821c57670d732b00a9eadd.
//
// Solidity: event NewPODBountyInitialized(uint256 _chainID, uint256 _totalBounty)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) ParseNewPODBountyInitialized(log types.Log) (*DiligenceProofManagerNewPODBountyInitialized, error) {
	event := new(DiligenceProofManagerNewPODBountyInitialized)
	if err := _DiligenceProofManager.contract.UnpackLog(event, "NewPODBountyInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiligenceProofManagerNewPOIBountyClaimedIterator is returned from FilterNewPOIBountyClaimed and is used to iterate over the raw logs and unpacked data for NewPOIBountyClaimed events raised by the DiligenceProofManager contract.
type DiligenceProofManagerNewPOIBountyClaimedIterator struct {
	Event *DiligenceProofManagerNewPOIBountyClaimed // Event containing the contract specifics and raw log

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
func (it *DiligenceProofManagerNewPOIBountyClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiligenceProofManagerNewPOIBountyClaimed)
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
		it.Event = new(DiligenceProofManagerNewPOIBountyClaimed)
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
func (it *DiligenceProofManagerNewPOIBountyClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiligenceProofManagerNewPOIBountyClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiligenceProofManagerNewPOIBountyClaimed represents a NewPOIBountyClaimed event raised by the DiligenceProofManager contract.
type DiligenceProofManagerNewPOIBountyClaimed struct {
	ChainID                   *big.Int
	L2BlockNumber             *big.Int
	SignatureProofOfDiligence []byte
	TotalBounty               *big.Int
	Miner                     common.Address
	Timestamp                 *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterNewPOIBountyClaimed is a free log retrieval operation binding the contract event 0xf12a89358d46d24eecf056793dac006c5642ddb743abc3ebb0a9a7afa6fcacf9.
//
// Solidity: event NewPOIBountyClaimed(uint256 indexed _chainID, uint256 indexed _l2BlockNumber, bytes _signatureProofOfDiligence, uint256 _totalBounty, address indexed _miner, uint256 _timestamp)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) FilterNewPOIBountyClaimed(opts *bind.FilterOpts, _chainID []*big.Int, _l2BlockNumber []*big.Int, _miner []common.Address) (*DiligenceProofManagerNewPOIBountyClaimedIterator, error) {

	var _chainIDRule []interface{}
	for _, _chainIDItem := range _chainID {
		_chainIDRule = append(_chainIDRule, _chainIDItem)
	}
	var _l2BlockNumberRule []interface{}
	for _, _l2BlockNumberItem := range _l2BlockNumber {
		_l2BlockNumberRule = append(_l2BlockNumberRule, _l2BlockNumberItem)
	}

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	logs, sub, err := _DiligenceProofManager.contract.FilterLogs(opts, "NewPOIBountyClaimed", _chainIDRule, _l2BlockNumberRule, _minerRule)
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManagerNewPOIBountyClaimedIterator{contract: _DiligenceProofManager.contract, event: "NewPOIBountyClaimed", logs: logs, sub: sub}, nil
}

// WatchNewPOIBountyClaimed is a free log subscription operation binding the contract event 0xf12a89358d46d24eecf056793dac006c5642ddb743abc3ebb0a9a7afa6fcacf9.
//
// Solidity: event NewPOIBountyClaimed(uint256 indexed _chainID, uint256 indexed _l2BlockNumber, bytes _signatureProofOfDiligence, uint256 _totalBounty, address indexed _miner, uint256 _timestamp)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) WatchNewPOIBountyClaimed(opts *bind.WatchOpts, sink chan<- *DiligenceProofManagerNewPOIBountyClaimed, _chainID []*big.Int, _l2BlockNumber []*big.Int, _miner []common.Address) (event.Subscription, error) {

	var _chainIDRule []interface{}
	for _, _chainIDItem := range _chainID {
		_chainIDRule = append(_chainIDRule, _chainIDItem)
	}
	var _l2BlockNumberRule []interface{}
	for _, _l2BlockNumberItem := range _l2BlockNumber {
		_l2BlockNumberRule = append(_l2BlockNumberRule, _l2BlockNumberItem)
	}

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	logs, sub, err := _DiligenceProofManager.contract.WatchLogs(opts, "NewPOIBountyClaimed", _chainIDRule, _l2BlockNumberRule, _minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiligenceProofManagerNewPOIBountyClaimed)
				if err := _DiligenceProofManager.contract.UnpackLog(event, "NewPOIBountyClaimed", log); err != nil {
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

// ParseNewPOIBountyClaimed is a log parse operation binding the contract event 0xf12a89358d46d24eecf056793dac006c5642ddb743abc3ebb0a9a7afa6fcacf9.
//
// Solidity: event NewPOIBountyClaimed(uint256 indexed _chainID, uint256 indexed _l2BlockNumber, bytes _signatureProofOfDiligence, uint256 _totalBounty, address indexed _miner, uint256 _timestamp)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) ParseNewPOIBountyClaimed(log types.Log) (*DiligenceProofManagerNewPOIBountyClaimed, error) {
	event := new(DiligenceProofManagerNewPOIBountyClaimed)
	if err := _DiligenceProofManager.contract.UnpackLog(event, "NewPOIBountyClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiligenceProofManagerNewPOIBountyInitializedIterator is returned from FilterNewPOIBountyInitialized and is used to iterate over the raw logs and unpacked data for NewPOIBountyInitialized events raised by the DiligenceProofManager contract.
type DiligenceProofManagerNewPOIBountyInitializedIterator struct {
	Event *DiligenceProofManagerNewPOIBountyInitialized // Event containing the contract specifics and raw log

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
func (it *DiligenceProofManagerNewPOIBountyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiligenceProofManagerNewPOIBountyInitialized)
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
		it.Event = new(DiligenceProofManagerNewPOIBountyInitialized)
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
func (it *DiligenceProofManagerNewPOIBountyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiligenceProofManagerNewPOIBountyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiligenceProofManagerNewPOIBountyInitialized represents a NewPOIBountyInitialized event raised by the DiligenceProofManager contract.
type DiligenceProofManagerNewPOIBountyInitialized struct {
	ChainID     *big.Int
	TotalBounty *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewPOIBountyInitialized is a free log retrieval operation binding the contract event 0x4ef58e80ae4475d8e97a37750e07ab9415ee5f2f166ab0e9436aecf8093adca9.
//
// Solidity: event NewPOIBountyInitialized(uint256 _chainID, uint256 _totalBounty)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) FilterNewPOIBountyInitialized(opts *bind.FilterOpts) (*DiligenceProofManagerNewPOIBountyInitializedIterator, error) {

	logs, sub, err := _DiligenceProofManager.contract.FilterLogs(opts, "NewPOIBountyInitialized")
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManagerNewPOIBountyInitializedIterator{contract: _DiligenceProofManager.contract, event: "NewPOIBountyInitialized", logs: logs, sub: sub}, nil
}

// WatchNewPOIBountyInitialized is a free log subscription operation binding the contract event 0x4ef58e80ae4475d8e97a37750e07ab9415ee5f2f166ab0e9436aecf8093adca9.
//
// Solidity: event NewPOIBountyInitialized(uint256 _chainID, uint256 _totalBounty)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) WatchNewPOIBountyInitialized(opts *bind.WatchOpts, sink chan<- *DiligenceProofManagerNewPOIBountyInitialized) (event.Subscription, error) {

	logs, sub, err := _DiligenceProofManager.contract.WatchLogs(opts, "NewPOIBountyInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiligenceProofManagerNewPOIBountyInitialized)
				if err := _DiligenceProofManager.contract.UnpackLog(event, "NewPOIBountyInitialized", log); err != nil {
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

// ParseNewPOIBountyInitialized is a log parse operation binding the contract event 0x4ef58e80ae4475d8e97a37750e07ab9415ee5f2f166ab0e9436aecf8093adca9.
//
// Solidity: event NewPOIBountyInitialized(uint256 _chainID, uint256 _totalBounty)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) ParseNewPOIBountyInitialized(log types.Log) (*DiligenceProofManagerNewPOIBountyInitialized, error) {
	event := new(DiligenceProofManagerNewPOIBountyInitialized)
	if err := _DiligenceProofManager.contract.UnpackLog(event, "NewPOIBountyInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiligenceProofManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DiligenceProofManager contract.
type DiligenceProofManagerOwnershipTransferredIterator struct {
	Event *DiligenceProofManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DiligenceProofManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiligenceProofManagerOwnershipTransferred)
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
		it.Event = new(DiligenceProofManagerOwnershipTransferred)
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
func (it *DiligenceProofManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiligenceProofManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiligenceProofManagerOwnershipTransferred represents a OwnershipTransferred event raised by the DiligenceProofManager contract.
type DiligenceProofManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DiligenceProofManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DiligenceProofManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManagerOwnershipTransferredIterator{contract: _DiligenceProofManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DiligenceProofManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DiligenceProofManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiligenceProofManagerOwnershipTransferred)
				if err := _DiligenceProofManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DiligenceProofManager *DiligenceProofManagerFilterer) ParseOwnershipTransferred(log types.Log) (*DiligenceProofManagerOwnershipTransferred, error) {
	event := new(DiligenceProofManagerOwnershipTransferred)
	if err := _DiligenceProofManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiligenceProofManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the DiligenceProofManager contract.
type DiligenceProofManagerPausedIterator struct {
	Event *DiligenceProofManagerPaused // Event containing the contract specifics and raw log

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
func (it *DiligenceProofManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiligenceProofManagerPaused)
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
		it.Event = new(DiligenceProofManagerPaused)
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
func (it *DiligenceProofManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiligenceProofManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiligenceProofManagerPaused represents a Paused event raised by the DiligenceProofManager contract.
type DiligenceProofManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*DiligenceProofManagerPausedIterator, error) {

	logs, sub, err := _DiligenceProofManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManagerPausedIterator{contract: _DiligenceProofManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *DiligenceProofManagerPaused) (event.Subscription, error) {

	logs, sub, err := _DiligenceProofManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiligenceProofManagerPaused)
				if err := _DiligenceProofManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_DiligenceProofManager *DiligenceProofManagerFilterer) ParsePaused(log types.Log) (*DiligenceProofManagerPaused, error) {
	event := new(DiligenceProofManagerPaused)
	if err := _DiligenceProofManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiligenceProofManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the DiligenceProofManager contract.
type DiligenceProofManagerUnpausedIterator struct {
	Event *DiligenceProofManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *DiligenceProofManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiligenceProofManagerUnpaused)
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
		it.Event = new(DiligenceProofManagerUnpaused)
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
func (it *DiligenceProofManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiligenceProofManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiligenceProofManagerUnpaused represents a Unpaused event raised by the DiligenceProofManager contract.
type DiligenceProofManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*DiligenceProofManagerUnpausedIterator, error) {

	logs, sub, err := _DiligenceProofManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManagerUnpausedIterator{contract: _DiligenceProofManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *DiligenceProofManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _DiligenceProofManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiligenceProofManagerUnpaused)
				if err := _DiligenceProofManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_DiligenceProofManager *DiligenceProofManagerFilterer) ParseUnpaused(log types.Log) (*DiligenceProofManagerUnpaused, error) {
	event := new(DiligenceProofManagerUnpaused)
	if err := _DiligenceProofManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiligenceProofManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the DiligenceProofManager contract.
type DiligenceProofManagerUpgradedIterator struct {
	Event *DiligenceProofManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *DiligenceProofManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiligenceProofManagerUpgraded)
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
		it.Event = new(DiligenceProofManagerUpgraded)
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
func (it *DiligenceProofManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiligenceProofManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiligenceProofManagerUpgraded represents a Upgraded event raised by the DiligenceProofManager contract.
type DiligenceProofManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*DiligenceProofManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _DiligenceProofManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &DiligenceProofManagerUpgradedIterator{contract: _DiligenceProofManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_DiligenceProofManager *DiligenceProofManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *DiligenceProofManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _DiligenceProofManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiligenceProofManagerUpgraded)
				if err := _DiligenceProofManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_DiligenceProofManager *DiligenceProofManagerFilterer) ParseUpgraded(log types.Log) (*DiligenceProofManagerUpgraded, error) {
	event := new(DiligenceProofManagerUpgraded)
	if err := _DiligenceProofManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
