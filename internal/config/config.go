package config

import (
	"os"
)

var (
	toastpath = os.Getenv("TOASTMAN_PATH")
)

func ToastPath() string { return toastpath }

func init() {
	if toastpath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		toastpath = home + "/.toastman"
	}
}
