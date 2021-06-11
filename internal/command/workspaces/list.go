package workspaces

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func List() (cmd *cli.Command) {
	cmd = &cli.Command{
		Name: "list",
		Usage: "list workspaces",
		Action: list,
	}

	return
}

func list(c *cli.Context) error {
	fmt.Println("unimplemented")

	return nil
}
