package utils

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

type TodoListItem struct {
	Task        string `json:"task"`
	TaskId      int    `json:"task_id"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	Done        bool   `json:"done"`
}

var Todolist []TodoListItem

func init() {
	// Set up logging format
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func CreateTodoListFile() {
	// Create the todo list file if it doesn't exist
	file, err := os.OpenFile("todolist.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		logrus.Fatalf("Error creating todo list file: %v", err)
	}
	defer file.Close()
	logrus.Info("Todo list file created successfully.")
}

func ReadTodoListFile() {
	// Read the todo list file
	file, err := os.OpenFile("todolist.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		logrus.Fatalf("Error reading todo list file: %v", err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&Todolist)
	if err != nil {
		logrus.Fatalf("Error decoding todo list file: %v", err)
	}
	logrus.Info("Todo list file read successfully.")
}

func WriteTodoListFile() {
	// Write the todo list to the file
	file, err := os.OpenFile("todolist.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		logrus.Fatalf("Error writing todo list file: %v", err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(Todolist)
	if err != nil {
		logrus.Fatalf("Error encoding todo list file: %v", err)
	}
	logrus.Info("Todo list file written successfully.")
}
