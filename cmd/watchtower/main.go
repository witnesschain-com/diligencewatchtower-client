package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
	"github.com/witnesschain-com/diligencewatchtower-client/watcher"
	"github.com/witnesschain-com/diligencewatchtower-client/webserver"

	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
	coordinator "github.com/witnesschain-com/diligencewatchtower-client/coordinator"
)

func main() {
	app := &cli.App{
		Name:  "diligencewatchtower-client",
		Usage: "Watch optimismtic rollup networks",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config",
				Value: "config.json",
				Aliases: []string{"c"},
				Usage: "path of the config.json file that store rollup and contract details",
			},
			&cli.BoolFlag{
				Name: "version",
				Usage: "print the version of watchtower client",
			},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.Bool("version") {
				fmt.Print(VERSION)
				os.Exit(0)
			}
			Start(cCtx.String("config"))
			return nil
		},
	}

	app.Suggest = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var VERSION = "undefined"

const sleepTimeIfNoChainAssigned = 5 * time.Second

func Start(configPath string) {

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
	configData := wtCommon.LoadConfigFromJson(configPath)
	
	

	simplifiedConfig := wtCommon.LoadSimplifiedConfig(configData, nil)

	// validate config valus are set correctly and that watchtower address is valid
	if !wtCommon.PreStartupChecks(configData, simplifiedConfig) {
		wtCommon.Fatal("Pre-startup checks failed!")
	}

	config := simplifiedConfig
	vault, err := keystore.SetupVault(config.WatchtowerAddress, big.NewInt(config.ProofSubmissionChainID), config.PrivateKey, config.ExternalSignerEndpoint)
	if err != nil {
		wtCommon.Error(err)
	}

	if len(configData.WatchtowerAddress) == 0 {
		configData.WatchtowerAddress = vault.NewTransactOpts(nil).From.Hex()
	}

	wtCommon.Info("Starting Watchtower (" + VERSION + ") ...")


	go func() {
		for {
			coordinator.StartCoordinator(*simplifiedConfig)
			time.Sleep(10 * time.Second)
			wtCommon.Info("Restarting coordiantor client...")
		}
	}()

	for {
		Run(configData, simplifiedConfig)
		wtCommon.Error("Restarting watchtower...............")
	}

}

func Run(configData *wtCommon.WatchTowerConfig, simplifiedConfig *wtCommon.SimplifiedConfig) bool {
	watchingChain := ""

	// channel used by webserver to alert watchtower about changes to config.json
	configChan := make(chan bool, 1)

	// creates wait group
	var waitGroup sync.WaitGroup

	// waitgroup waits for 1 process
	waitGroup.Add(2)

	var server *http.Server

	// load the config params for the webserver
	webServerConfig := wtCommon.LoadWebServerConfig(configData)

	// run webserver in a go routine eldlessly
	go webserver.Start(
		server,
		webServerConfig,
		configChan,
	)

	// run watcher as goroutines and wait using the wait group
	go watcher.StartDiligenceWatcher(&waitGroup, configData, simplifiedConfig, configChan)
	go watcher.StartInclusionWatcher(&waitGroup, simplifiedConfig)
	wtCommon.Success("WitnessChain Watchtower started!")

	// wait for the 1 process to end (watchtower to stop)
	waitGroup.Wait()

	// watchtower stopped, irrespective of why, stop the webserver too
	webserver.Stop(server)

	if len(watchingChain) == 0 {
		wtCommon.Info("Sleeping for 5 seconds ...\n")
		time.Sleep(sleepTimeIfNoChainAssigned)
	}

	return true
}
