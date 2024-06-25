package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
	"github.com/witnesschain-com/diligencewatchtower-client/coordinator"
)

func main() {
	var configPath string
	app := &cli.App{
		Name:  "diligencewatchtower-client Proof of Inclusion",
		Usage: "Proof of Inclusion for rollup networks",
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

			wtCommon.Info("Starting Coordinator ...")

			coordinator.StartCoordinator(*simplifiedConfig)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
