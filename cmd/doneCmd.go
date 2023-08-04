package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:     "done",
	Short:   "Mark a task as done",
	Aliases: []string{"d"},
	Example: "togo done [id1] [id2] [id3] ...",
	Run:     markTaskAsDone,
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func markTaskAsDone(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("❌ Please specify the ID(s) of the task(s) to mark as completed")
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

		task.Status = models.StatusDone
		result = dbConn.Save(&task)
		if result.Error != nil {
			fmt.Printf("🚨 Failed to update the task with ID %s: %v\n", arg, result.Error)
			failedTasks = append(failedTasks, arg)
			continue
		}

		fmt.Printf("👉 Done Task ID %s\n", arg)
	}

	if len(failedTasks) > 0 {
		fmt.Printf("⚠️ Some tasks could not be marked as completed: %s\n", strings.Join(failedTasks, ", "))
	}
}
