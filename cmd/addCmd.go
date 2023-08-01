package cmd

import (
	"fmt"

	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new task",
	Example: "togo add [\"Task title\"]",
	Run:     addTask,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addTask(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("❌ タスクのタイトルを指定してください。")
		return
	}

	taskTitle := args[0]

	task := models.Task{
		Title:  taskTitle,
		Status: models.StatusPending,
	}

	result := dbConn.Create(&task)
	if result.Error != nil {
		fmt.Println("🚨 タスクの追加に失敗しました:", result.Error)
		return
	}

	fmt.Printf("👉 Add Task\n")
	task.RenderTaskTable()
}
