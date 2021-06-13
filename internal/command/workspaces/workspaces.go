package workspaces

import (
	"fmt"
	"os"
	"regexp"

	"github.com/rreubenreyes/toastman/internal/config"
)

func WorkspacesPath() string { return "workspaces" }

func DefaultWorkspaceName() string {
	return ".default"
}

func DefaultWorkspace() string {
	path := fmt.Sprintf("%s/%s/%s", config.ToastPath(), WorkspacesPath(), DefaultWorkspaceName())
	os.MkdirAll(path, os.ModePerm)

	return path
}

func IsValidWorkspaceName(name string) bool {
	expr, err := regexp.Compile(`^[\w\d\-\_\.]*$`)
	if err != nil {
		panic(err)
	}
	return expr.Match([]byte(name))
}
