package common

import (
	"github.com/diligencewatchtower-client/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Check if we are eigen-layer operator
// Check if we are whitelisted by witnesschain
func ValidateWatchtowerAddress(config *WatchTowerConfig) bool {
	Info("Validating watchtower address ...")
	callOpts := bind.CallOpts{}

	L1Client, err := GetConnection(
		config.ProofSubmissionWebsocketURL,
		config.Retries,
	)
	if err != nil {
		Fatal(err)
	}
	defer L1Client.Close()

	publicKeyAddress, err := GetPublicKeyAddressFromPrivateKey(config.PrivateKey)
	if err != nil {
		Error(err)
	}

	publicKeyAddressHex := publicKeyAddress.Hex()
	Info("Your `publicKeyAddressHex` is")
	Success(publicKeyAddressHex)

	OperatorRegistryAddress := common.HexToAddress(config.OperatorRegistry)

	operatorRegistryContract, err := bindings.NewOperatorRegistry(OperatorRegistryAddress, L1Client)
	if err != nil {
		Fatal(err)
	}
	isValid, err := operatorRegistryContract.IsValidWatchtower(&callOpts, publicKeyAddress)
	if err != nil {
		Fatal(err)
	} else if isValid {
		Success("Watchtower address is valid")
	} else {
		Error("Watchtower address is invalid, please ensure that your watchtower's eth address is registered with WitnessChain")
		return false
	}

	return true

}

func PreStartupChecks(config *WatchTowerConfig) bool {

	// validate necessary config is set correctly in config file
	configIsValid := ValidateConfig(config)

	// validate the watchtower address is registered and belongs to active operator
	watchtowerAddressIsValid := ValidateWatchtowerAddress(config)

	return (configIsValid && watchtowerAddressIsValid)

}
