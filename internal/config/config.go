package config

import (
	"log"
	"os"
)

var (
	Stderr    = log.New(os.Stderr, "", 0)
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
