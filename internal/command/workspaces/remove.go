package workspaces

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Remove() (cmd *cli.Command) {
	cmd = &cli.Command{
		Name: "remove",
		Usage: "remove workspaces",
		Action: remove,
	}

	return
}

func remove(c *cli.Context) error {
	fmt.Println("unimplemented")

	return nil
}
