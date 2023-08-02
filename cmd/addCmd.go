package cmd

import (
	"fmt"

	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new task",
	Aliases: []string{"a"},
	Example: "togo add [\"Task title\"]",
	Run:     addTask,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addTask(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("❌ Please specify the task title")
		return
	}

	taskTitle := args[0]

	task := models.Task{
		Title:  taskTitle,
		Status: models.StatusPending,
	}

	result := dbConn.Create(&task)
	if result.Error != nil {
		fmt.Println("🚨 Failed to add the task:", result.Error)
		return
	}

	fmt.Printf("👉 Add Task\n")
	task.RenderTaskTable()
}
