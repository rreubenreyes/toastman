package workspaces

import (
	"fmt"
	"io/ioutil"

	"github.com/rreubenreyes/toastman/internal/config"
	"github.com/urfave/cli/v2"
)

func List() (cmd *cli.Command) {
	cmd = &cli.Command{
		Name:   "list",
		Usage:  "list workspaces",
		Action: list,
	}

	return
}

func list(c *cli.Context) error {
	DefaultWorkspace()

	workspaces, err := ioutil.ReadDir(config.ToastPath() + "/" + WorkspacesPath())
	if err != nil {
		panic(err)
	}
	for _, workspace := range workspaces {
		fmt.Println(workspace.Name())
	}

	return nil
}
