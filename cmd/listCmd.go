package cmd

import (
	"fmt"
	"time"

	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:     "list [flags]",
		Short:   "List all tasks",
		Aliases: []string{"l"},
		Example: "togo list --sort [ i | t | p | c | u | d ] --sort-direction [asc|desc]",
		Run:     listTasks,
	}
	sortBy        string
	sortDirection string
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&sortBy, "sort", "s", "created_at", "\nSort tasks by column\n[options] :  id(i) | title(t) | priority(p) | created_at(c) | updated_at(u) | due_date(d)\n[default] : ")
	listCmd.Flags().StringVarP(&sortDirection, "sort-direction", "d", "asc", "\nSort direction\n[options] : asc | desc\n[default] : ")
}

func listTasks(cmd *cobra.Command, args []string) {
	if _, ok := columnsMapping[sortBy]; !ok {
		fmt.Println("❌ Invalid sort column", sortBy)
		return
	}

	var tasks []models.Task
	orderClause := fmt.Sprintf("%s %s", columnsMapping[sortBy], sortDirection)
	if columnsMapping[sortBy] == "due_date" {
		// Put zero/null dates at the end
		orderClause = fmt.Sprintf("CASE WHEN due_date IS NULL OR due_date = '0001-01-01 00:00:00+00:00' THEN 1 ELSE 0 END, due_date %s", sortDirection)
	}

	result := dbConn.Order(orderClause).Find(&tasks)
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

	models.RenderTasksTable(tasks)
}
