package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
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

