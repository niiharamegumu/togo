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

var (
	listCmd = &cobra.Command{
		Use:     "list [flags]",
		Short:   "List tasks by status",
		Aliases: []string{"l"},
		Example: "togo list --status [status: pen | done | all] --sort [ i | t | s | p | c | u | d ] --sort-direction [asc|desc]",
		Run:     listTasks,
	}
	sortBy        string
	sortDirection string
	statusFlag    string
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&sortBy, "sort", "s", "created_at", "\nSort tasks by column\n[options] :  id(i) | title(t) | status(s) | priority(p) | created_at(c) | updated_at(u) | due_date(d)\n[default] : ")
	listCmd.Flags().StringVarP(&sortDirection, "sort-direction", "d", "asc", "\nSort direction\n[options] : asc | desc\n[default] : ")
	listCmd.Flags().StringVarP(&statusFlag, "status", "S", "pen", "\nFilter tasks by status\n[options] : pen | done | all\n[default] : ")
}

func listTasks(cmd *cobra.Command, args []string) {
	var statusFilter string
	if statusFlag != "" {
		statusFilter = statusFlag
	}

	if _, ok := columnsMapping[sortBy]; !ok {
		fmt.Println("❌ Invalid sort column", sortBy)
		return
	}

	var tasks []models.Task
	var result *gorm.DB

	switch statusFilter {
	case "pen", "":
		result = dbConn.Order(fmt.Sprintf("%s %s", columnsMapping[sortBy], sortDirection)).Find(&tasks, "status = ?", models.StatusPending)
		if result.Error != nil {
			fmt.Println("🚨 Failed to retrieve the task:", result.Error)
			return
		}
	case "done":
		result = dbConn.Order(fmt.Sprintf("%s %s", columnsMapping[sortBy], sortDirection)).Find(&tasks, "status = ?", models.StatusDone)
		if result.Error != nil {
			fmt.Println("🚨 Failed to retrieve the task:", result.Error)
			return
		}
	case "all":
		result = dbConn.Find(&tasks)
		if result.Error != nil {
			fmt.Println("🚨 Failed to retrieve the task:", result.Error)
			return
		}
	default:
		fmt.Println("❌ Invalid status filter", statusFilter)
		return
	}
	if len(tasks) == 0 {
		fmt.Println("👉 No Tasks")
		return
	}

	today := time.Now().Format("TODAY:2006/01/02")
	fmt.Printf("%v\n", today)
	n := len(tasks)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(models.TaskTableHeader)
	for i, task := range tasks {
		var dueDateStr string
		if !task.DueDate.IsZero() && task.DueDate.Before(time.Now()) {
			dueDateStr = fmt.Sprintf("\x1b[31m%s\x1b[0m", task.DueDate.Format("2006/01/02"))
		} else if !task.DueDate.IsZero() {
			dueDateStr = task.DueDate.Format("2006/01/02")
		} else {
			dueDateStr = ""
		}
		table.Append([]string{
			fmt.Sprintf("%d", task.ID),
			task.Title,
			task.Status,
			fmt.Sprintf("%d", task.Priority),
			task.CreatedAt.Format("2006/01/02 15:04"),
			task.UpdatedAt.Format("2006/01/02 15:04"),
			dueDateStr,
		})
		if i != n-1 {
			emptyRow := make([]string, len(models.TaskTableHeader))
			table.Append(emptyRow)
		}
	}

	table.Render()
}
