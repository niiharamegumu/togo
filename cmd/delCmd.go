package cmd

import (
	"fmt"

	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:     "del",
	Short:   "delete task",
	Aliases: []string{"de"},
	Example: "togo del [id]",
	Run:     deleteTask,
}

func init() {
	rootCmd.AddCommand(delCmd)
}

func deleteTask(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("❌ Please specify the ID of the task to be deleted")
		return
	}

	taskID := args[0]

	var task models.Task
	result := dbConn.First(&task, taskID)
	if result.Error != nil {
		fmt.Println("🚨 Failed to retrieve the task:", result.Error)
		return
	}

	result = dbConn.Delete(&task)
	if result.Error != nil {
		fmt.Println("🚨 Failed to delete the task:", result.Error)
		return
	}

	fmt.Println("👉 Delete Task")
}
