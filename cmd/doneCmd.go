package cmd

import (
	"fmt"

	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:     "done",
	Short:   "Mark a task as done",
	Aliases: []string{"d"},
	Example: "togo done [id]",
	Run:     markTaskAsDone,
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func markTaskAsDone(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("❌ Please specify the ID of the task to mark as completed")
		return
	}

	taskID := args[0]

	var task models.Task
	result := dbConn.First(&task, taskID)
	if result.Error != nil {
		fmt.Println("🚨 Failed to retrieve the task:", result.Error)
		return
	}

	task.Status = models.StatusDone
	result = dbConn.Save(&task)
	if result.Error != nil {
		fmt.Println("🚨 Failed to update the task:", result.Error)
		return
	}

	fmt.Println("👉 Done Task")
	task.RenderTaskTable()
}
