package cmd

import (
	"github.com/qcserestipy/gotodolist/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func DeleteItemCmd() *cobra.Command {
	// Define the add command
	var addCmd = &cobra.Command{
		Use:   "delete",
		Short: "Deletes a task from the todo list",
		Long:  `Deletes a task from the todo list by its index.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				logrus.Error("Please provide the id of a task to delete.")
				return
			}
			index := int(args[0][0] - '0')
			for i, task := range utils.Todolist {
				if task.TaskId == index {
					logrus.Infof("Deleting task: %s", task.Task)
					utils.Todolist = append(utils.Todolist[:i], utils.Todolist[i+1:]...)
					utils.WriteTodoListFile()
					break
				}
			}
		},
	}

	addCmd.Flags().IntP("index", "i", -1, "Index of the task to delete")
	return addCmd
}
