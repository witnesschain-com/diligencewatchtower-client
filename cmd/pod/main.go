package main

import (
	"log"
	"os"
	"sync"

	"github.com/urfave/cli/v2"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
	"github.com/witnesschain-com/diligencewatchtower-client/watcher"
)

func main() {
	var configPath string
	app := &cli.App{
		Name:  "diligencewatchtower-client Proof of Diligence",
		Usage: "Proof of Diligence for rollup networks",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config",
				Value: "config.json",
				Aliases: []string{"c"},
				Destination: &configPath,
				Usage: "path of the config.json file that store rollup and contract details",
			},
		},
		Action: func(cCtx *cli.Context) error {

			configData := wtCommon.LoadConfigFromJson(configPath)
			simplifiedConfig := wtCommon.LoadSimplifiedConfig(configData, nil)

			wtCommon.Info("Starting Proof of Diligence Watchtower ...")

			var wg sync.WaitGroup
			wg.Add(1)

			configChan := make(chan bool, 1)
			watcher.StartDiligenceWatcher(&wg, configData, simplifiedConfig, configChan)

			wg.Wait()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

