/*
	Copyright (c) 2024 Kaleidoscope Blockchain Inc.

	Unless specified otherwise, this work is licensed under the Creative Commons
	Attribution-NonCommercial 4.0 International License.

	To view a copy of this license, visit:
		http://creativecommons.org/licenses/by-nc/4.0/
*/

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/nodeapi"

	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
	coordinator "github.com/witnesschain-com/diligencewatchtower-client/coordinator"
	"github.com/witnesschain-com/diligencewatchtower-client/external"
	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
	"github.com/witnesschain-com/diligencewatchtower-client/watcher"
	"github.com/witnesschain-com/diligencewatchtower-client/webserver"
)

const version = "v1.0"
const sleepTimeIfNoChainAssigned = 5 * time.Second

func main() {
	fmt.Print(`
                           ___
                    o___.-'  /
                    │      __\
                    │___.-'
                    │
                    │
            _   _   │   _   _
           [_]_[_]_[_]_[_]_[_]
           [__│__│__│__│__│__]
             [_│__│__│__│__]
             [__│__│^_│__│_]
             [_│__│/^\_│__│]
             [__│_// \\__│_]
             [_│__│__|__│__]
             [__│_│_|_│__│_]
             [_│__|__│__│__]
             [__│__│__│__│_]
  _   _   _  [_│__│__│__│__]  _   _   _
_[_]_[_]_[_]_[__│__│__│__│_]_[_]_[_]_[_]_
  _│__│__│__│[_│__│__│__│__]│__│__│__│__│_
     │  │  │ [  │  │  │  │ ] │  │  │  │ `)

	fmt.Println("")
	fmt.Println("")

	// read config
	configData := wtCommon.LoadConfigFromJson()

	simplifiedConfig := wtCommon.LoadSimplifiedConfig(configData, nil)

	// validate config valus are set correctly and that watchtower address is valid
	if !wtCommon.PreStartupChecks(configData, simplifiedConfig) {
		wtCommon.Fatal("Pre-startup checks failed!")
	}

	vault, err := keystore.SetupVault(simplifiedConfig)
	if err != nil {
		wtCommon.Error(err)
	}

	if len(configData.WatchtowerAddress) == 0 {
		configData.WatchtowerAddress = vault.NewTransactOpts(nil).From.Hex()
	}

	wtCommon.Info("Starting Watchtower (" + version + ") ...")

	// setup EL monitoring and nodeAPI
	external.InitialiseELMonitoring()
	nodeApi := external.InitialiseNodeAPIWithLogger(configData)

	go func() {
		for {
			coordinator.StartCoordinator(*simplifiedConfig)
			time.Sleep(10 * time.Second)
			wtCommon.Info("Restarting coordiantor client...")
		}
	}()

	for {
		Run(configData, simplifiedConfig, nodeApi)
	}

}

func Run(configData *wtCommon.WatchTowerConfig, simplifiedConfig *wtCommon.SimplifiedConfig, nodeApi *nodeapi.NodeApi) bool {
	watchingChain := ""

	// channel used by webserver to alert watchtower about changes to config.json
	configChan := make(chan bool, 1)

	// creates wait group
	var waitGroup sync.WaitGroup

	// waitgroup waits for 1 process
	waitGroup.Add(1)

	var server *http.Server

	// load the config params for the webserver
	webServerConfig := wtCommon.LoadWebServerConfig(configData)

	// run webserver in a go routine eldlessly
	go webserver.Start(
		server,
		webServerConfig,
		configChan,
	)

	err := nodeApi.UpdateServiceStatus(webServerConfig.PublicKeyAddressHex, nodeapi.ServiceStatusUp)
	if err != nil {
		wtCommon.Error(err)
	}

	// run watcher as goroutines and wait using the wait group
	go watcher.StartWatcher(&waitGroup, configData, simplifiedConfig, configChan)
	wtCommon.Success("WitnessChain Watchtower started!")

	// wait for the 1 process to end (watchtower to stop)
	waitGroup.Wait()

	// watchtower stopped, irrespective of why, stop the webserver too
	webserver.Stop(server)

	err = nodeApi.UpdateServiceStatus(webServerConfig.PublicKeyAddressHex, nodeapi.ServiceStatusDown)
	if err != nil {
		wtCommon.Error(err)
	}

	if len(watchingChain) == 0 {
		wtCommon.Info("Sleeping for 5 seconds ...\n")
		time.Sleep(sleepTimeIfNoChainAssigned)
	}

	return true
}
