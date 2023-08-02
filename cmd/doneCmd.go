package cmd

import (
	"fmt"

	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:     "done",
	Short:   "Mark a task as done",
	Aliases: []string{"d"},
	Example: "togo done [id]",
	Run:     markTaskAsDone,
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func markTaskAsDone(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("❌ 完了済みにするタスクのIDを指定してください。")
		return
	}

	taskID := args[0]

	var task models.Task
	result := dbConn.First(&task, taskID)
	if result.Error != nil {
		fmt.Println("🚨 タスクの取得に失敗しました:", result.Error)
		return
	}

	task.Status = models.StatusDone
	result = dbConn.Save(&task)
	if result.Error != nil {
		fmt.Println("🚨 タスクの更新に失敗しました:", result.Error)
		return
	}

	fmt.Printf("👉 Done Task\n")
	task.RenderTaskTable()
}
