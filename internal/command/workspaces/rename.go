package workspaces

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Rename() (cmd *cli.Command) {
	cmd = &cli.Command{
		Name: "rename",
		Usage: "rename workspaces",
		Action: rename,
	}

	return
}

func rename(c *cli.Context) error {
	fmt.Println("unimplemented")

	return nil
}
