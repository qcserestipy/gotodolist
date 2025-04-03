package cmd

import (
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
			if index < 0 || index >= len(Todolist) {
				logrus.Error("Invalid index provided.")
				return
			}
			// Find the task by index and delete it
			for _, task := range Todolist {
				if task.TaskId == index {
					logrus.Infof("Deleting task: %s", task.Task)
					// Remove the task from the list
					Todolist = append(Todolist[:index], Todolist[index+1:]...)
					logrus.Infof("Deleted task at index: %d", index)
					break
				}
			}
			// Todolist = append(Todolist[:index], Todolist[index+1:]...)
			// logrus.Infof("Deleted task at index: %d", index)
		},
	}

	addCmd.Flags().IntP("index", "i", -1, "Index of the task to delete")
	return addCmd
}
