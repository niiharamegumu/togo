package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update [id]",
	Short:   "Update a task's title",
	Aliases: []string{"u"},
	Example: "togo update [id]",
	Run:     updateTask,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateTask(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("❌ Please specify the task ID")
		return
	}

	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("❌ Invalid task ID. Please provide a valid integer ID.")
		return
	}

	var task models.Task
	result := dbConn.First(&task, taskID)
	if result.Error != nil {
		fmt.Println("🚨 Failed to retrieve the task:", result.Error)
		return
	}

	task.RenderTaskTable()

	fmt.Printf("Enter the new task title : ")
	reader := bufio.NewReader(os.Stdin)
	newTitle, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("🚨 Error reading input:", err)
		return
	}
	newTitle = strings.TrimSpace(newTitle)
	if newTitle == "" {
		newTitle = task.Title
	}
	task.Title = newTitle

	fmt.Printf("Enter the Priority (0-100) : ")
	reader = bufio.NewReader(os.Stdin)
	priorityStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("🚨 Error reading input:", err)
		return
	}
	priorityStr = strings.TrimSpace(priorityStr)
	priority, err := strconv.Atoi(priorityStr)
	if err != nil {
		priority = 0
	}
	if priority < 0 {
		priority = 0
	}
	if priority > 100 {
		priority = 100
	}
	task.Priority = priority

	result = dbConn.Save(&task)
	if result.Error != nil {
		fmt.Println("🚨 Failed to update the task:", result.Error)
		return
	}

	fmt.Printf("👉 Updated Task ID %v \n", task.ID)
	task.RenderTaskTable()
}
