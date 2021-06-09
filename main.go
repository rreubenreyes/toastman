package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type Args struct {
	subcommand string
}

func main() {
	app := &cli.App{
		Name: "toastman",
		Usage: "CLI HTTP client with Postman-like features",
		Action: func (c *cli.Context) error {
			fmt.Println("the main command performs an HTTP request")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
