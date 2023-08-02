package cmd

import (
	"fmt"

	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update [id] [title]",
	Short:   "Update a task's title",
	Aliases: []string{"u"},
	Example: "togo update [id] [\"Update title\"]",
	Run:     updateTask,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateTask(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("❌ タスクIDとタイトルを指定してください。")
		return
	}

	taskID := args[0]
	taskTitle := args[1]

	var task models.Task
	result := dbConn.First(&task, taskID)
	if result.Error != nil {
		fmt.Println("🚨 タスクの取得に失敗しました:", result.Error)
		return
	}

	task.Title = taskTitle
	result = dbConn.Save(&task)
	if result.Error != nil {
		fmt.Println("🚨 タスクの更新に失敗しました:", result.Error)
		return
	}

	fmt.Printf("👉 Update Task ID %v \n", task.ID)
	task.RenderTaskTable()
}
