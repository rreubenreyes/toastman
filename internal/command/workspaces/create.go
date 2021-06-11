package workspaces

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Create() (cmd *cli.Command) {
	cmd = &cli.Command{
		Name: "create",
		Usage: "create a new workspace",
		Action: create,
	}

	return
}

func create(c *cli.Context) error {
	fmt.Println("unimplemented")

	return nil
}
