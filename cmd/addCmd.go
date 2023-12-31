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

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new task",
	Aliases: []string{"a"},
	Example: "togo add [\"Task title\"]",
	Run:     addTask,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&dueDate, "due-date", "d", "", "\nSet the task's due date\n[format] : '2006-01-02'")
}

func addTask(cmd *cobra.Command, args []string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input the new task title")
	fmt.Println("Press Enter twice to finish, you can enter multiple lines")
	var taskTitleBuilder strings.Builder
	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}
		taskTitleBuilder.WriteString(line)
		taskTitleBuilder.WriteString("\n")
	}

	taskTitle := strings.TrimSpace(taskTitleBuilder.String())
	if taskTitle == "" {
		fmt.Println("👌 Exiting the process")
		return
	}

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

	var dueDateTime time.Time
	if dueDate != "" {
		dueDateTime, err = time.Parse("2006-01-02", dueDate)
		if err != nil {
			fmt.Println("🚨 Invalid date format for --due-date flag. Please use 'YYYY-MM-DD'.")
			return
		}
	}

	task := models.Task{
		Title:    taskTitle,
		Status:   models.StatusPending,
		Priority: priority,
		DueDate:  dueDateTime,
	}

	result := dbConn.Create(&task)
	if result.Error != nil {
		fmt.Println("🚨 Failed to add the task:", result.Error)
		return
	}

	fmt.Println("👉 Task Added")
	task.RenderTaskTable()
}
