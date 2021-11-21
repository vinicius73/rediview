package cmd

import (
	"github.com/urfave/cli/v2"
	"github.com/vinicius73/rediview/pkg/config"
	"github.com/vinicius73/rediview/pkg/server"
	"net"
	"strconv"
	// conf "github.com/vinicius73/rediview/pkg/config"
)

func Server(_ config.Config) *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "start rediview server",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "port",
				Usage: "set server addr port",
				Value: 3073,
			},
		},
		Action: func(c *cli.Context) error {
			port := strconv.Itoa(c.Int("port"))

			return server.Start(c.Context, server.Config{Addr: net.JoinHostPort("", port)})
		},
	}
}
