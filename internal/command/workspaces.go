package command

import (
	"github.com/rreubenreyes/toastman/internal/command/workspaces"
	"github.com/urfave/cli/v2"
)

func Workspaces() (cmd *cli.Command) {
	cmd = &cli.Command{
		Name:  "workspaces",
		Usage: "manage workspaces",
		Subcommands: []*cli.Command{
			workspaces.List(),
			workspaces.Create(),
			workspaces.Describe(),
			workspaces.Remove(),
			workspaces.Rename(),
		},
	}

	return
}
