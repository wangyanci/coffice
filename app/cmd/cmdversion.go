package cmd

import (
	"gopkg.in/urfave/cli.v2"
)

var VersionCmd *cli.Command

func init() {
	VersionCmd = &cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "show the app version!",
		Action:  ShowVersion,
	}
}

func ShowVersion(c *cli.Context) error {
	return nil
}
