package cmd

import (
	"fmt"

	"github.com/niiharamegumu/togo/db"
	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:     "done",
	Short:   "Mark a task as done",
	Example: "togo update [id]",
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]

		db, err := db.ConnectDB()
		if err != nil {
			fmt.Println("🚨 データベースに接続できませんでした:", err)
			return
		}

		var task models.Task
		result := db.First(&task, taskID)
		if result.Error != nil {
			fmt.Println("🚨 タスクの取得に失敗しました:", result.Error)
			return
		}

		task.Status = models.StatusDone
		result = db.Save(&task)
		if result.Error != nil {
			fmt.Println("🚨 タスクの更新に失敗しました:", result.Error)
			return
		}

		fmt.Println("👉 Done Task")
	},
}
