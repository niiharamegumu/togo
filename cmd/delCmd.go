package cmd

import (
	"fmt"

	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:     "del",
	Short:   "delete task",
	Example: "togo del [id]",
	Run:     deleteTask,
}

func init() {
	rootCmd.AddCommand(delCmd)
}

func deleteTask(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("❌ 削除するタスクのIDを指定してください。")
		return
	}

	taskID := args[0]

	var task models.Task
	result := dbConn.First(&task, taskID)
	if result.Error != nil {
		fmt.Println("🚨 タスクの取得に失敗しました:", result.Error)
		return
	}

	result = dbConn.Delete(&task)
	if result.Error != nil {
		fmt.Println("🚨 タスクの削除に失敗しました:", result.Error)
		return
	}

	fmt.Print("👉 Delete Task")
}
