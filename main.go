package main

import (
	"fmt"
	"os"

	"github.com/qcserestipy/gotodolist/cmd"
	"github.com/sirupsen/logrus"
)

func init() {
	// Initialize logrus timeformat
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func main() {
	err := cmd.RootCmd().Execute()
	fmt.Println(cmd.Todolist)
	if err != nil {
		os.Exit(1)
	}
}
