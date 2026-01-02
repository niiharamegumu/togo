package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/niiharamegumu/togo/models"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:     "del",
	Short:   "delete task",
	Aliases: []string{"de"},
	Example: "togo del",
	Run:     deleteTask,
}

func init() {
	rootCmd.AddCommand(delCmd)
}

func deleteTask(cmd *cobra.Command, args []string) {
	var tasks []models.Task
	result := dbConn.Find(&tasks)
	if result.Error != nil {
		fmt.Println("🚨 Failed to retrieve tasks:", result.Error)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("👉 No Tasks to delete")
		return
	}

	options := make([]huh.Option[string], len(tasks))
	for i, task := range tasks {
		cleanTitle := strings.ReplaceAll(task.Title, "\n", " ")
		label := fmt.Sprintf("[%d] %s", task.ID, cleanTitle)
		options[i] = huh.NewOption(label, strconv.Itoa(int(task.ID)))
	}

	var selectedIDs []string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Select tasks to delete").
				Options(options...).
				Value(&selectedIDs),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("❌ Selection cancelled or failed")
		return
	}

	if len(selectedIDs) == 0 {
		fmt.Println("👉 No tasks selected")
		return
	}

	for _, idStr := range selectedIDs {
		id, _ := strconv.Atoi(idStr)
		res := dbConn.Delete(&models.Task{}, id)
		if res.Error != nil {
			fmt.Printf("🚨 Failed to delete task ID %d: %v\n", id, res.Error)
		} else {
			fmt.Printf("👉 Deleted Task ID %d\n", id)
		}
	}
}
