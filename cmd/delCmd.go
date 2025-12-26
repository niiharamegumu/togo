package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:     "del",
	Short:   "delete task",
	Aliases: []string{"de"},
	Example: "togo del [id1] [id2] [id3] ...",
	Run:     deleteTask,
}

func init() {
	rootCmd.AddCommand(delCmd)
}

func deleteTask(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("❌ Please specify the ID(s) of the task(s) to be deleted")
		return
	}

	var failedTasks []string
	for _, arg := range args {
		taskID, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("❌ Invalid task ID: %s\n", arg)
			failedTasks = append(failedTasks, arg)
			continue
		}

		var task models.Task
		result := dbConn.First(&task, taskID)
		if result.Error != nil {
			fmt.Printf("🚨 Failed to retrieve the task with ID %s: %v\n", arg, result.Error)
			failedTasks = append(failedTasks, arg)
			continue
		}

		result = dbConn.Unscoped().Delete(&task)
		if result.Error != nil {
			fmt.Printf("🚨 Failed to delete the task with ID %s: %v\n", arg, result.Error)
			failedTasks = append(failedTasks, arg)
			continue
		}

		fmt.Printf("👉 Deleted Task ID %s\n", arg)
	}

	if len(failedTasks) > 0 {
		fmt.Printf("⚠️ Some tasks could not be deleted: %s\n", strings.Join(failedTasks, ", "))
	}
}
