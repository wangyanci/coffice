package app

import (
	"fmt"

	"vueApp/app/cmd"

	"gopkg.in/urfave/cli.v2"
)

func NewApp() *cli.App {
	app := &cli.App{
		Name: "greet",
		Usage: "welcome to huangsewangzhan!",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},

		Flags: []cli.Flag{

		},

		Commands: []*cli.Command{
			cmd.WebCmd,
			cmd.VersionCmd,
		},
	}

	return app
}
