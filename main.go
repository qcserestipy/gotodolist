package main

import (
	"os"

	"github.com/qcserestipy/gotodolist/cmd"
)

func main() {
	err := cmd.RootCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}
