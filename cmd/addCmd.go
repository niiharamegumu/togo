package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
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
	var (
		taskTitle   string
		priorityStr string
		dueDateStr  string
	)

	// If flags are set, use them as default or skip?
	// For simplicity and "optimizing for huh", we'll make it fully interactive but use flags as initial values if present.
	if dueDate != "" {
		dueDateStr = dueDate
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("Task Title").
				Value(&taskTitle).
				Validate(func(str string) error {
					if strings.TrimSpace(str) == "" {
						return fmt.Errorf("title cannot be empty")
					}
					return nil
				}),

			huh.NewInput().
				Title("Priority (0-100)").
				Value(&priorityStr).
				Validate(func(str string) error {
					if str == "" {
						return nil // Allow empty for default
					}
					p, err := strconv.Atoi(str)
					if err != nil {
						return fmt.Errorf("priority must be a number")
					}
					if p < 0 || p > 100 {
						return fmt.Errorf("priority must be between 0 and 100")
					}
					return nil
				}),

			huh.NewInput().
				Title("Due Date (YYYY-MM-DD)").
				Value(&dueDateStr).
				Validate(func(str string) error {
					if str == "" {
						return nil // Allow empty
					}
					_, err := time.Parse("2006-01-02", str)
					if err != nil {
						return fmt.Errorf("invalid date format")
					}
					return nil
				}),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("❌ Operation cancelled")
		return
	}

	priority := 0
	if priorityStr != "" {
		priority, _ = strconv.Atoi(priorityStr)
	}

	var dueDateTime time.Time
	if dueDateStr != "" {
		dueDateTime, _ = time.Parse("2006-01-02", dueDateStr)
	}

	var maxID *uint
	dbConn.Unscoped().Model(&models.Task{}).Select("MAX(id)").Scan(&maxID)

	var nextID uint
	if maxID == nil {
		nextID = 1
	} else {
		nextID = *maxID + 1
	}

	task := models.Task{
		Title:    taskTitle,
		Priority: priority,
		DueDate:  dueDateTime,
	}
	task.ID = nextID

	result := dbConn.Create(&task)
	if result.Error != nil {
		fmt.Println("🚨 Failed to add the task:", result.Error)
		return
	}

	fmt.Println("👉 Task Added")
	task.RenderTaskTable()
}
