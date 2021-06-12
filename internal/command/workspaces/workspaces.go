package workspaces

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/rreubenreyes/toastman/internal/config"
)

func WorkspacesPath() string { return "workspaces" }

func DefaultWorkspace() string {
	path := fmt.Sprintf("%s/%s/default", config.ToastPath(), WorkspacesPath())
	os.MkdirAll(path, os.ModePerm)

	return path
}

func IsValidWorkspaceName(name []byte) bool {
	expr, err := regexp.Compile(`^[\w\d\-\_\.]*$`)
	if err != nil {
		log.Fatal("[workspaces] error parsing workspace name regexp")
	}
	return expr.Match(name)
}
