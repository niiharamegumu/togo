package cmd

import (
	"fmt"

	"github.com/niiharamegumu/togo/db"
	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new task",
	Example: "togo add [\"Task title\"]",
	Run: func(cmd *cobra.Command, args []string) {
		taskTitle := args[0]

		task := models.Task{
			Title:  taskTitle,
			Status: models.StatusPending,
		}

		db, err := db.ConnectDB()
		if err != nil {
			fmt.Println("🚨 データベースに接続できませんでした:", err)
			return
		}

		result := db.Create(&task)
		if result.Error != nil {
			fmt.Println("🚨 タスクの追加に失敗しました:", result.Error)
			return
		}

		fmt.Println("👉 Add Task")
	},
}
