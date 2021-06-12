package workspaces

import (
	"fmt"
	"os"

	"github.com/rreubenreyes/toastman/internal/config"
	"github.com/urfave/cli/v2"
)

func Create() (cmd *cli.Command) {
	cmd = &cli.Command{
		Name:   "create",
		Usage:  "create a new workspace",
		Action: create,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Usage: "name of the new workspace",
			},
		},
	}

	return
}

func create(c *cli.Context) error {
	name := c.String("name")

	if name == "" {
		cli.ShowSubcommandHelpAndExit(c, 0)
		return nil
	}

	if !IsValidWorkspaceName([]byte(name)) {
		panic(fmt.Sprintf("invalid workspace name \"%s\"", name))
	}

	newWorkspacePath := fmt.Sprintf("%s/%s/%s", config.ToastPath(), WorkspacesPath(), name)
	os.MkdirAll(newWorkspacePath, os.ModePerm)

	return nil
}
