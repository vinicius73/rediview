package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
	"github.com/vinicius73/rediview/cli/cmd"
	conf "github.com/vinicius73/rediview/pkg/config"
	"github.com/vinicius73/rediview/pkg/support"
)

func main() {
	config := conf.Build()

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "show rediview version info",
	}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("Rediview %s (%s %s)\nVinicius Reis <vinicius73.dev>\ngithub.com/vinicius73/rediview\n", config.Version(), config.Commit(), config.BuildDate())
	}

	app := &cli.App{
		Name:    "redview",
		Usage:   "a interface for redis",
		Version: config.Version(),
		Before: func(c *cli.Context) error {
			logger := support.Logger(c.Args().First(), config.Tags())

			c.Context = config.WithContext(c.Context)
			c.Context = logger.WithContext(c.Context)

			return nil
		},
		Commands: cli.Commands{
			cmd.Server(config),
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
