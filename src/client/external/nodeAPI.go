package external

import (
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	wtCommon "github.com/diligencewatchtower-client/common"
)

func initialiseNodeAPI(logger logging.Logger, PublicKeyAddressHex string) *nodeapi.NodeApi {

	nodeApi := nodeapi.NewNodeApi("watchtowerAVS", "v0.0.1", ":8080", logger)
	nodeApi.RegisterNewService(PublicKeyAddressHex, "witnesschain-watchtower", "witnesschain watchtower for roll-ups", nodeapi.ServiceStatusInitializing)
	nodeApi.Start()
	return nodeApi
}

func InitialiseNodeAPIWithLogger(
	configData *wtCommon.WatchTowerConfig,

) *nodeapi.NodeApi {
	logger := GetELLogger()
	webServerConfig := wtCommon.LoadWebServerConfig(configData)
	nodeApi := initialiseNodeAPI(logger, webServerConfig.PublicKeyAddressHex)
	nodeApi.UpdateHealth(nodeapi.Healthy) // Eg. use case
	return nodeApi
}
