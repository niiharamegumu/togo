package cmd

import (
	"fmt"
	"os"

	"github.com/niiharamegumu/togo/db"
	"github.com/niiharamegumu/togo/models"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update a task's title",
	Example: "togo update [id] [\"Update title\"]",
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]
		taskTitle := args[1]

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

		task.Title = taskTitle
		result = db.Save(&task)
		if result.Error != nil {
			fmt.Println("🚨 タスクの更新に失敗しました:", result.Error)
			return
		}

		fmt.Printf("👉 Update Task ID %v \n", task.ID)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Title", "Status", "Date"})
		table.Append([]string{
			fmt.Sprintf("%d", task.ID),
			task.Title,
			task.Status,
			task.CreatedAt.Format("2006/01/02 15:04"),
		})
		table.Render()
	},
}
