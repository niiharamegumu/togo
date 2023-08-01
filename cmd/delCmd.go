package cmd

import (
	"fmt"

	"github.com/niiharamegumu/togo/db"
	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:     "del",
	Short:   "delete task",
	Example: "togo delete [id]",
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

		result = db.Delete(&task)
		if result.Error != nil {
			fmt.Println("🚨 タスクの削除に失敗しました:", result.Error)
			return
		}

		fmt.Println("👉 Delete Task")
	},
}
