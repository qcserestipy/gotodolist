package cmd

import (
	"os"

	"github.com/qcserestipy/gotodolist/pkg/utils"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	// Define the root command
	var rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "A simple CLI todo list application",
		Long:  `A simple CLI todo list application that allows you to add, remove, and view tasks.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if _, err := os.Stat("todolist.json"); os.IsNotExist(err) {
				utils.CreateTodoListFile()
			} else {
				utils.ReadTodoListFile()
			}
			return nil
		},
	}

	// Add subcommands
	rootCmd.AddCommand(
		AddItemCmd(),
		DeleteItemCmd(),
	)

	return rootCmd
}
