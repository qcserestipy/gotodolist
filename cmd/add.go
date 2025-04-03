package cmd

import (
	"time"

	"github.com/qcserestipy/gotodolist/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func AddItemCmd() *cobra.Command {
	// Define the add command
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a new task to the todo list",
		Long:  `Add a new task to the todo list with an optional description.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				logrus.Error("Please provide a task to add.")
				return
			}
			description, err := cmd.Flags().GetString("description")
			if err != nil {
				logrus.Error("Error retrieving description flag:", err)
				return
			}
			if description == "" {
				logrus.Warn("No description provided for the task.")
				description = "No description"
			}
			utils.Todolist = append(utils.Todolist, utils.TodoListItem{
				Task:        args[0],
				TaskId:      len(utils.Todolist) + 1,
				Description: description,
				CreatedAt:   time.Now().Format(time.RFC3339),
				Done:        false,
			})
			utils.WriteTodoListFile()
			logrus.Infof("Added task: %s", args[0])
		},
	}

	addCmd.Flags().StringP("description", "d", "", "Description of the task")
	return addCmd
}
