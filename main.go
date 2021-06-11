package main

import (
	"log"
	"os"

	"github.com/rreubenreyes/toastman/internal/command"
	"github.com/urfave/cli/v2"
)

func main() {
	stderr := log.New(os.Stderr, "", 0)

	app := &cli.App{
		Name:  "toastman",
		Usage: "friendly http client",
		Commands: []*cli.Command{
			command.Workspaces(),
		},
		Action: func(c *cli.Context) error {
			cli.ShowAppHelpAndExit(c, 0)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		stderr.Fatal(err)
	}
}
