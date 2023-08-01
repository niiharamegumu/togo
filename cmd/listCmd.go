package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/niiharamegumu/togo/db"
	"github.com/niiharamegumu/togo/models"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var listCmd = &cobra.Command{
	Use:   "list [status]",
	Short: "List tasks by status",
	Run: func(cmd *cobra.Command, args []string) {
		today := time.Now().Format("[TODAY:2006/01/02]")
		var statusFilter string
		if len(args) > 0 {
			statusFilter = args[0]
		}

		db, err := db.ConnectDB()
		if err != nil {
			fmt.Println("データベースに接続できませんでした:", err)
			return
		}

		var tasks []models.Task
		var result *gorm.DB

		if statusFilter == "pen" || statusFilter ==  "" {
			result = db.Find(&tasks, "status = ?", models.StatusPending)
		}else if statusFilter == "done" {
			result = db.Find(&tasks, "status = ?", models.StatusDone)
		} else {
			result = db.Find(&tasks)
		}

		if result.Error != nil {
			fmt.Println("タスクの取得に失敗しました:", result.Error)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("=== No Tasks ===")
			return
		}

		fmt.Printf("%v TOGO LIST : \n", today)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Title", "Status", "Date"})

		n := len(tasks)
		for i, task := range tasks {
			table.Append([]string{
				fmt.Sprintf("%d", task.ID),
				task.Title,
				task.Status,
				task.CreatedAt.Format("2006/01/02 15:04"),
			})

			if i != n-1 {
				table.Append([]string{"", "", "", ""})
			}
		}

		table.Render()
	},
}