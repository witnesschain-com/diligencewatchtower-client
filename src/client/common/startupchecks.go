package common

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/witnesschain-com/diligencewatchtower-client/bindings"
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

	Success(config.WatchtowerAddress)

	OperatorRegistryAddress := common.HexToAddress(config.OperatorRegistry)

	operatorRegistryContract, err := bindings.NewOperatorRegistry(OperatorRegistryAddress, L1Client)
	if err != nil {
		Fatal(err)
	}
	simplifiedConfig := LoadSimplifiedConfig(config, nil)
	isValid, err := operatorRegistryContract.IsValidWatchtower(&callOpts, simplifiedConfig.WatchtowerAddress)
	if err != nil {
		Fatal(err)
	} else if isValid {
		Success("Watchtower" + simplifiedConfig.WatchtowerAddress.Hex() + " is not registered with WitnessChain, please register ")
	} else {
		Error("Watchtower" + simplifiedConfig.WatchtowerAddress.Hex() +  "address is invalid, please ensure that your watchtower's eth address is registered with WitnessChain")
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
