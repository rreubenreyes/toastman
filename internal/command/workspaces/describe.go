package workspaces

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Describe() (cmd *cli.Command) {
	cmd = &cli.Command{
		Name: "describe",
		Usage: "show details of a workspace",
		Action: describe,
	}

	return
}

func describe(c *cli.Context) error {
	fmt.Println("unimplemented")

	return nil
}
