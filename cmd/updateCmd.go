package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

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

	updateCmd.Flags().StringVarP(&dueDate, "due-date", "d", "", "\nSet the task's due date\n[format] : '2006-01-02'")
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

	var dueDateTime time.Time
	if dueDate != "" {
		dueDateTime, err = time.Parse("2006-01-02", dueDate)
		if err != nil {
			fmt.Println("🚨 Invalid date format for --due-date flag. Please use 'YYYY-MM-DD'.")
			return
		}
		task.DueDate = dueDateTime
	}

	task.RenderTaskTable()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input the new task title")
	fmt.Println("Press Enter twice to finish, you can enter multiple lines")
	var newTitleBuilder strings.Builder
	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}
		newTitleBuilder.WriteString(line)
		newTitleBuilder.WriteString("\n")
	}

	newTitle := strings.TrimSpace(newTitleBuilder.String())
	if newTitle == "" {
		newTitle = task.Title
	}
	task.Title = newTitle

	fmt.Print("Enter the Priority (0-100): ")
	scanner.Scan()
	priorityStr := scanner.Text()

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
