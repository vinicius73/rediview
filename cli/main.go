package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	conf "github.com/vinicius73/rediview/pkg/config"
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
			c.Context = config.WithContext(c.Context)
			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
