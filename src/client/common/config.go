/*
	Copyright (c) 2024 Kaleidoscope Blockchain Inc.

	Unless specified otherwise, this work is licensed under the Creative Commons
	Attribution-NonCommercial 4.0 International License.

	To view a copy of this license, visit:
		http://creativecommons.org/licenses/by-nc/4.0/
*/

package common

import (
	"crypto/ecdsa"
	"encoding/json"
	"io"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

// WatchTowerConfig is used to store all the configurable parameters
// of the watchtower
type WatchTowerConfig struct {
	// L1 - ethereum RPC urls
	PrivateKey                  string `json:"private_key"`
	Vault                       string `json:"encrypted_vault_directory"`
	EthWebsocketURL             string `json:"eth_websocket_url"`
	EthTestnetWebsocketURL      string `json:"eth_testnet_websocket_url"`
	ProofSubmissionWebsocketURL string `json:"proof_submission_chain_url"`
	CurrentlyWatchingL1         string `json:"currently_watching_l1"`
	EthTestnetChainID           int    `json:"eth_testnet_chain_id"`
	ProofSubmissionChainID      int    `json:"proof_submission_chain_id"`

	// L2 - config params
	L2                  L2     `json:"l2"`
	CurrentlyWatchingL2 string `json:"currently_watching_l2"`

	// keys and tuning
	WatchtowerAddress          string `json:"watchtower_address"`
	Retries        int   `json:"watchtower_retries"`
	ReceiptTimeout int   `json:"receipt_timeout"`
	GasPrice       int64 `json:"gas_price"`

	// Webserver/API params
	HostName                   string `json:"host_name"`
	Port                       string `json:"port"`
	AlertURL                   string `json:"watchtower_failure_alert_url"`
	WitnesschainCoordinatorUrl string `json:"witnesschain_coordinator_url"`

	// Contract addresses
	AlertManagerAddress          string `json:"alert_manager_address"`
	DiligenceProofManagerAddress string `json:"diligence_proof_manager_address"`
	OperatorRegistry             string `json:"operator_registry"`
	RollupRegistry               string `json:"rollup_registry"`
	ExternalSignerEndpoint       string `json:"external_signer_endpoint"`
}

func ValidateConfig(config *WatchTowerConfig) bool {
	Info("Validating config values ...")
	var isValid bool = true

	isSupportedChain := map[string]bool{
		"optimism": true,
		"base":     true,
		"mode":     true,
		"zora":     true,
	}

	// validate L1 chain and it's RPC set correctly
	if config.CurrentlyWatchingL1 == "" {
		Error("Invalid config: please set 'currently_watching_l1' and corresponding websocket_url")
		isValid = false
	}
	if config.CurrentlyWatchingL1 == "eth_testnet" {
		if config.EthTestnetWebsocketURL == "" {
			Error("Invalid config: please set 'eth_testnet_websocket_url'")
			isValid = false
		}
		if config.EthTestnetChainID == 0 {
			Error("Invalid config: please set 'eth_testnet_chain_id'")
			isValid = false
		}
	}
	if config.CurrentlyWatchingL1 == "eth" && config.EthWebsocketURL == "" {
		Error("Invalid config: please set 'eth_websocket_url'")
		isValid = false
	}

	if config.ProofSubmissionWebsocketURL == "" {
		Error("Invalid config: please set 'proof_submission_chain_url'")
		isValid = false
	}
	if config.ProofSubmissionChainID == 0 {
		Error("Invalid config: please set 'proof_submission_chain_id'")
		isValid = false
	}

	// validate L2 chain and it's RPC is set correctly
	if config.CurrentlyWatchingL2 == "" {
		Error("Invalid config: please set 'currently_watching_l2'")
		isValid = false
	}

	if !isSupportedChain[config.CurrentlyWatchingL2] {
		Error("Invalid config: update 'currently_watching_l2', as " + config.CurrentlyWatchingL2 + " is not unsupported")
		isValid = false
	}

	if len(config.L2.OPChains) == 0 {
		Error("Invalid config: please set network params in op_super_chains")
		isValid = false
	} else {
		for _, chain := range config.L2.OPChains {
			if chain.Name == "" {
				Error("Invalid config: please set network param 'chain_name' for in op_super_chains")
				isValid = false
			}
			if !isSupportedChain[chain.Name] {
				Error("Invalid config: unsupported chain " + chain.Name)
				isValid = false
				continue
			}
			if chain.ChainID == 0 {
				Error("Invalid config: please set network param 'chain_id' for " + chain.Name)
				isValid = false
			}
			if chain.L2OOAddress == "" {
				Error("Invalid config: please set network param 'l2oo_address' for " + chain.Name)
				isValid = false
			}
			if chain.NodeRPCURL == "" {
				Error("Invalid config: please set network param 'op_node_rpc_url' for " + chain.Name)
				isValid = false
			}
			if chain.RPCURL == "" {
				Error("Invalid config: please set network param 'op_geth_rpc_url' for " + chain.Name)
				isValid = false
			}
			if chain.WebsocketURL == "" {
				Warning("Network param 'websocket_url' not set for " + chain.Name)
			}
		}
	}

	// validate contract addresses are set
	if config.AlertManagerAddress == "" {
		Error("Invalid config: please set 'alert_manager_address'")
		isValid = false
	}
	if config.DiligenceProofManagerAddress == "" {
		Error("Invalid config: please set 'diligence_proof_manager_address'")
		isValid = false
	}
	if config.OperatorRegistry == "" {
		Error("Invalid config: please set 'operator_registry'")
		isValid = false
	}
	if config.RollupRegistry == "" {
		Warning("'rollup_registry' not set")
	}

	// validate optional configs
	if config.HostName == "" {
		Warning("'host_name' not set")
	}

	if config.AlertURL == "" {
		Warning("'watchtower_failure_alert_url' not set")
	}

	if config.ReceiptTimeout == 0 {
		Warning("'receipt_timeout' not set, will default to 180")
		config.ReceiptTimeout = 180
	}
	if config.Retries == 0 {
		Warning("'watchtower_retries' not set, will default to 5")
		config.Retries = 5
	}
	if config.GasPrice == 0 {
		Warning("'gas_price' not set, will use network estimates")
	}

	if isValid {
		Success("Config is Valid")
	} else {
		Error("Config validation failed! please fix above issues")
	}

	if config.ExternalSignerEndpoint == "" && config.PrivateKey == "" && config.Vault == "" {
		Error("Incorrect config, please set at least one of the following: external_signer_endpoint, encrypted_vault_directory, or private_key")
		isValid = false
	}

	return isValid

}

// `Optimism` is used to store all the configurable
// parameters of op-stack based chains
type Optimism struct {
	WebsocketURL string `json:"op_geth_websocket_url"`
	Name         string `json:"chain_name"`
	ChainID      int    `json:"chain_id"`
	RPCURL       string `json:"op_geth_rpc_url"`
	NodeRPCURL   string `json:"op_node_rpc_url"`
	L2OOAddress  string `json:"l2oo_address"`
}

// `Arbitrum` is used to store all the configurable
// parameters of arbitrum rollup
type Arbitrum struct {
	// not yet supported
}
type L2 struct {
	OPChains []Optimism `json:"op_super_chains"`
	Arbitrum Arbitrum   `json:"arbitrum"`
}

// `WebServerConfig` stores the configurable params of the webserver
// these are the only modifiable aspects of the webserver
type WebServerConfig struct {
	HostName            string
	Port                string
	PublicKeyAddressHex string
	CurrentlyWatchingL2 string
}

// `SimplifiedConfig` is used to store chain agnostic config
// for watchtower to make it easier to manage and use the values
type SimplifiedConfig struct {
	L1WebsocketURL               string
	L2WebsocketURL               string
	ProofSubmissionWebsocketURL  string
	ProofSubmissionChainID       int64
	AlertManagerAddress          ethCommon.Address
	DiligenceProofManagerAddress ethCommon.Address
	L2ExecRPCURL                 string
	L2NodeRPCURL                 string
	StateCommitmentAddress       ethCommon.Address
	CurrentL2ChainID             int64
	CurrentL1ChainID             int64
	Retries                      int
	ReceiptTimeout               int
	GasPrice                     int64
	WatchtowerAddress            common.Address
	ExternalSignerEndpoint       string
	Vault                        string
	PrivateKey                   *ecdsa.PrivateKey
}

// `LoadConfigFromJson` returns a config object of type `WatchTowerConfig` by loading
// the values from the locally present `config.json` file
func LoadConfigFromJson() *WatchTowerConfig {

	Info("Reading the config from `config.json` ...")
	configFile, err := os.Open("config.json")
	if err != nil {
		Error(err)
		os.Exit(123)
	}
	defer configFile.Close()

	configData, err := io.ReadAll(configFile)
	if err != nil {
		Error(err)
		os.Exit(123)
	}

	var Config WatchTowerConfig
	err = json.Unmarshal(configData, &Config)
	if err != nil {
		Error("Can not unmarshal JSON")
		Error(err)
		os.Exit(123)
	}

	Success("Read `config.json`")

	return &Config

}

// `UpdateConfigToJson` writes the passed config object
// of type `WatchTowerConfig` to the “config.json“ file,
// it return a boolean value, `true` for success and `false` otherwise
func UpdateConfigToJson(config *WatchTowerConfig) bool {

	configJson, err := json.Marshal(*config)
	if err != nil {
		Fatal("Cannot marshal JSON")
		Error(err)
		return false
	}

	err = os.WriteFile("config.json", configJson, 0600)
	if err != nil {
		Fatal(err)
		return false
	}

	return true

}

func LoadSimplifiedConfig(config *WatchTowerConfig, simpleConfig *SimplifiedConfig) *SimplifiedConfig {

	if simpleConfig == nil {
		simpleConfig = new(SimplifiedConfig)
	}

	simpleConfig.WatchtowerAddress = common.HexToAddress(config.WatchtowerAddress)
	simpleConfig.AlertManagerAddress = ethCommon.HexToAddress(config.AlertManagerAddress)
	simpleConfig.DiligenceProofManagerAddress = ethCommon.HexToAddress(config.DiligenceProofManagerAddress)
	simpleConfig.Retries = config.Retries
	simpleConfig.ReceiptTimeout = config.ReceiptTimeout
	simpleConfig.GasPrice = config.GasPrice
	simpleConfig.ProofSubmissionWebsocketURL = config.ProofSubmissionWebsocketURL
	simpleConfig.ProofSubmissionChainID = int64(config.ProofSubmissionChainID)
	simpleConfig.ExternalSignerEndpoint = config.ExternalSignerEndpoint
	simpleConfig.Vault = config.Vault

	if len(config.PrivateKey) > 0 {
		key := config.PrivateKey
		if config.PrivateKey[0:2] == "0x"{
			key = config.PrivateKey[2:]
		}
		private_key, err := crypto.HexToECDSA(key)
		if err != nil {
			Fatal(err)
		}
		simpleConfig.PrivateKey = private_key
	}

	switch config.CurrentlyWatchingL1 {
	case "eth":
		simpleConfig.L1WebsocketURL = config.EthWebsocketURL
		simpleConfig.CurrentL1ChainID = 1
	case "eth_testnet":
		simpleConfig.L1WebsocketURL = config.EthTestnetWebsocketURL
		simpleConfig.CurrentL1ChainID = int64(config.EthTestnetChainID)
	default:
		simpleConfig.L1WebsocketURL = config.EthTestnetWebsocketURL
		simpleConfig.CurrentL1ChainID = 11155111 //sepolia
	}

	var chainParamsSet bool = false
	for _, chain := range config.L2.OPChains {
		if chain.Name == config.CurrentlyWatchingL2 {
			simpleConfig.L2WebsocketURL = chain.WebsocketURL
			simpleConfig.L2ExecRPCURL = chain.RPCURL
			simpleConfig.L2NodeRPCURL = chain.NodeRPCURL
			simpleConfig.StateCommitmentAddress = ethCommon.HexToAddress(chain.L2OOAddress)
			simpleConfig.CurrentL2ChainID = int64(chain.ChainID)
			chainParamsSet = true
		}
	}

	if !chainParamsSet {
		Error("Config not found for " + config.CurrentlyWatchingL2)
		Warning("Program will continue with old config if set earlier")
	}

	var isValidConfig bool = true

	if simpleConfig.L1WebsocketURL == "" {
		Warning("No websocket url found for " + config.CurrentlyWatchingL1)
		isValidConfig = false
	}
	if simpleConfig.L2WebsocketURL == "" {
		Warning("No websocket url found for " + config.CurrentlyWatchingL2)
		isValidConfig = false
	}
	if simpleConfig.L2ExecRPCURL == "" {
		Warning("No RPCURL found for " + config.CurrentlyWatchingL2)
		isValidConfig = false
	}
	if simpleConfig.L2NodeRPCURL == "" {
		Warning("No NodeRPCURL found for " + config.CurrentlyWatchingL2)
		isValidConfig = false
	}
	if simpleConfig.CurrentL2ChainID == 0 {
		Warning("No L2ChainID found for " + config.CurrentlyWatchingL2)
		isValidConfig = false
	}
	if simpleConfig.StateCommitmentAddress == ethCommon.HexToAddress("") {
		Warning("No State commitment address (L2OO) found for " + config.CurrentlyWatchingL2)
		isValidConfig = false
	}
	if !isValidConfig {
		Fatal("Invalid/Missing configuration for " + config.CurrentlyWatchingL2 + " please fix your config file")
		return nil
	}

	return simpleConfig
}

func LoadWebServerConfig(config *WatchTowerConfig) *WebServerConfig {
	var webServerConfig WebServerConfig

	webServerConfig.HostName = config.HostName
	webServerConfig.Port = config.Port
	webServerConfig.CurrentlyWatchingL2 = config.CurrentlyWatchingL2
	webServerConfig.PublicKeyAddressHex = config.WatchtowerAddress

	return &webServerConfig

}
