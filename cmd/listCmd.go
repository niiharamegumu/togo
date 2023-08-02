package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/niiharamegumu/togo/models"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var listCmd = &cobra.Command{
	Use:     "list [status]",
	Short:   "List tasks by status",
	Aliases: []string{"l"},
	Example: "togo list [ status: pen | done | all | '']",
	Run:     listTasks,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listTasks(cmd *cobra.Command, args []string) {
	var statusFilter string
	if len(args) > 0 {
		statusFilter = args[0]
	}

	var tasks []models.Task
	var result *gorm.DB

	switch statusFilter {
	case "pen", "":
		result = dbConn.Find(&tasks, "status = ?", models.StatusPending)
	case "done":
		result = dbConn.Find(&tasks, "status = ?", models.StatusDone)
	case "all":
		result = dbConn.Find(&tasks)
	default:
		fmt.Println("❌ Invalid status filter", statusFilter)
		return
	}

	if result.Error != nil {
		fmt.Println("🚨 Failed to retrieve the task:", result.Error)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("👉 No Tasks")
		return
	}

	today := time.Now().Format("TODAY:2006/01/02")
	fmt.Printf("%v\n", today)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(models.TaskTableHeader)
	n := len(tasks)
	for i, task := range tasks {
		table.Append([]string{
			fmt.Sprintf("%d", task.ID),
			task.Title,
			task.Status,
			fmt.Sprintf("%d", task.Priority),
			task.CreatedAt.Format("2006/01/02 15:04"),
		})

		if i != n-1 {
			table.Append([]string{"", "", "", ""})
		}
	}

	table.Render()
}
