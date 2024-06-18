package common

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/witnesschain-com/diligencewatchtower-client/bindings"
)

// Check if we are eigen-layer operator
// Check if we are whitelisted by witnesschain
func ValidateWatchtowerAddress(simplifiedConfig *SimplifiedConfig) bool {
	callOpts := bind.CallOpts{}

	L1Client, err := GetConnection(
		simplifiedConfig.ProofSubmissionWebsocketURL,
		simplifiedConfig.Retries,
	)
	if err != nil {
		Fatal(err)
	}
	defer L1Client.Close()

	Info("Validating watchtower address ...")
	Info("watchtower address: " + simplifiedConfig.WatchtowerAddress.Hex())

	operatorRegistryContract, err := bindings.NewOperatorRegistry(simplifiedConfig.OperatorRegistryAddress, L1Client)
	if err != nil {
		Fatal(err)
	}

	isValid, err := operatorRegistryContract.IsValidWatchtower(&callOpts, simplifiedConfig.WatchtowerAddress)
	if err != nil {
		Fatal(err)
	} else if isValid {
		Success("Watchtower: " + simplifiedConfig.WatchtowerAddress.Hex() + " is registered with WitnessChain")
	} else {
		Error("Watchtower" + simplifiedConfig.WatchtowerAddress.Hex() + "address is invalid, please ensure that your watchtower's eth address is registered with WitnessChain")
		return false
	}

	return true

}

func PreStartupChecks(config *WatchTowerConfig, simplifiedConfig *SimplifiedConfig) bool {

	// validate necessary config is set correctly in config file
	configIsValid := ValidateConfig(config)

	// validate the watchtower address is registered and belongs to active operator
	watchtowerAddressIsValid := ValidateWatchtowerAddress(simplifiedConfig)

	return (configIsValid && watchtowerAddressIsValid)

}
