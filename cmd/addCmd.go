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

	task := models.Task{
		Title:    taskTitle,
		Status:   models.StatusPending,
		Priority: priority,
	}

	result := dbConn.Create(&task)
	if result.Error != nil {
		fmt.Println("🚨 Failed to add the task:", result.Error)
		return
	}

	fmt.Printf("👉 Add Task\n")
	task.RenderTaskTable()
}
