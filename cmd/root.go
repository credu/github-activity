package cmd

import (
	"os"
)

func Execute() {
	args := os.Args

	if len(args) != 2 {
		ShowDefaultMessage()
		os.Exit(1)
	}

	username := args[1]
	ShowUserActivityByUsername(username)
}
