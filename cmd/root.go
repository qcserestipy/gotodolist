package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type TodoListItem struct {
	Task        string `json:"task"`
	TaskId      int    `json:"task_id"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	Done        bool   `json:"done"`
}

var Todolist []TodoListItem

func RootCmd() *cobra.Command {
	// Define the root command
	var rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "A simple CLI todo list application",
		Long:  `A simple CLI todo list application that allows you to add, remove, and view tasks.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("Welcome to the todo list application!")
		},
	}

	// Add subcommands
	rootCmd.AddCommand(
		AddItemCmd(),
		DeleteItemCmd(),
	)

	return rootCmd
}
