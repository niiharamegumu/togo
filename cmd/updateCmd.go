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

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update a task's title",
	Aliases: []string{"u"},
	Example: "togo update",
	Run:     updateTask,
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&dueDate, "due-date", "d", "", "\nSet the task's due date\n[format] : '2006-01-02'")
}

func updateTask(cmd *cobra.Command, args []string) {
	var tasks []models.Task
	result := dbConn.Find(&tasks)
	if result.Error != nil {
		fmt.Println("🚨 Failed to retrieve the task:", result.Error)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("👉 No Tasks to update")
		return
	}

	options := make([]huh.Option[string], len(tasks))
	for i, task := range tasks {
		cleanTitle := strings.ReplaceAll(task.Title, "\n", " ")
		label := fmt.Sprintf("[%d] %s", task.ID, cleanTitle)
		options[i] = huh.NewOption(label, strconv.Itoa(int(task.ID)))
	}

	var selectedIDStr string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select a task to update").
				Options(options...).
				Value(&selectedIDStr),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("❌ Selection cancelled or failed")
		return
	}

	taskID, _ := strconv.Atoi(selectedIDStr)
	var task models.Task
	result = dbConn.First(&task, taskID)
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

	// Pre-fill values
	newTitle := task.Title
	newPriorityStr := strconv.Itoa(task.Priority)
	newDueDateStr := ""
	if !task.DueDate.IsZero() {
		newDueDateStr = task.DueDate.Format("2006-01-02")
	}

	updateForm := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("Task Title").
				Value(&newTitle).
				Validate(func(str string) error {
					if strings.TrimSpace(str) == "" {
						return fmt.Errorf("title cannot be empty")
					}
					return nil
				}),

			huh.NewInput().
				Title("Priority (0-100)").
				Value(&newPriorityStr).
				Validate(func(str string) error {
					if str == "" {
						return nil
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
				Value(&newDueDateStr).
				Validate(func(str string) error {
					if str == "" {
						return nil
					}
					_, err := time.Parse("2006-01-02", str)
					if err != nil {
						return fmt.Errorf("invalid date format")
					}
					return nil
				}),
		),
	)

	err = updateForm.Run()
	if err != nil {
		fmt.Println("❌ Update cancelled")
		return
	}

	task.Title = newTitle

	if newPriorityStr != "" {
		p, _ := strconv.Atoi(newPriorityStr)
		task.Priority = p
	} else {
		task.Priority = 0
	}

	if newDueDateStr != "" {
		dd, _ := time.Parse("2006-01-02", newDueDateStr)
		task.DueDate = dd
	} else {
		task.DueDate = time.Time{}
	}

	result = dbConn.Save(&task)
	if result.Error != nil {
		fmt.Println("🚨 Failed to update the task:", result.Error)
		return
	}

	fmt.Printf("👉 Updated Task ID %v \n", task.ID)
	task.RenderTaskTable()
}
