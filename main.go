package main

import (
	"log"
	"os"

	"github.com/rreubenreyes/toastman/internal/command"
	"github.com/rreubenreyes/toastman/internal/config"
	"github.com/urfave/cli/v2"
)

var stderr = log.New(os.Stderr, "", 0)

func main() {
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
		Before: func(_ *cli.Context) error {
			// Create workspaces directory if it doesn't already exist
			os.MkdirAll(config.ToastPath()+"/workspaces", os.ModePerm)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		stderr.Fatal(err)
	}
}
