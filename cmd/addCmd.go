package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new task",
	Aliases: []string{"a"},
	Example: "togo add [\"Task title\"]",
	Run:     addTask,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addTask(cmd *cobra.Command, args []string) {
	fmt.Printf("Enter the new task title : ")
	reader := bufio.NewReader(os.Stdin)
	taskTitle, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("🚨 Error reading input:", err)
		return
	}

	taskTitle = strings.TrimSpace(taskTitle)

	if taskTitle == "" {
		fmt.Println("👌 Exiting the process")
		return
	}

	task := models.Task{
		Title:  taskTitle,
		Status: models.StatusPending,
	}

	result := dbConn.Create(&task)
	if result.Error != nil {
		fmt.Println("🚨 Failed to add the task:", result.Error)
		return
	}

	fmt.Printf("👉 Add Task\n")
	task.RenderTaskTable()
}
