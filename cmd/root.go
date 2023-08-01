package cmd

import (
	"fmt"
	"os"

	"github.com/niiharamegumu/togo/db"
	"github.com/niiharamegumu/togo/models"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Display available commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available commands:")
		for _, c := range cmd.Root().Commands() {
			if c.Name() != "help" {
				fmt.Printf("- %s: %s\n", c.Name(), c.Short)
			}
		}
	},
}

var rootCmd = &cobra.Command{
	Use:   "togo",
	Short: "Task Management CLI",
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Example: "togo add [\"Task title\"]",
	Run: func(cmd *cobra.Command, args []string) {
		taskTitle := args[0]

		task := models.Task{
			Title:  taskTitle,
			Status: models.StatusPending,
		}

		db, err := db.ConnectDB()
		if err != nil {
			fmt.Println("データベースに接続できませんでした:", err)
			return
		}

		result := db.Create(&task)
		if result.Error != nil {
			fmt.Println("タスクの追加に失敗しました:", result.Error)
			return
		}

		fmt.Println("=== Add Task ===")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := db.ConnectDB()
		if err != nil {
			fmt.Println("データベースに接続できませんでした:", err)
			return
		}

		var tasks []models.Task
		result := db.Find(&tasks)
		if result.Error != nil {
			fmt.Println("タスクの取得に失敗しました:", result.Error)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("=== Not Tasks ===")
			return
		}

		fmt.Println("TOGO LIST:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Title", "Status", "Date"})

		for _, task := range tasks {
			table.Append([]string{
				fmt.Sprintf("%d", task.ID),
				task.Title,
				task.Status,
				task.CreatedAt.Format("2006/01/02 15:04"),
			})
		}

		table.Render()
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task's title",
	Example: "togo update [id] [\"Update title\"]",
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]
		taskTitle := args[1]

		db, err := db.ConnectDB()
		if err != nil {
			fmt.Println("データベースに接続できませんでした:", err)
			return
		}

		var task models.Task
		result := db.First(&task, taskID)
		if result.Error != nil {
			fmt.Println("タスクの取得に失敗しました:", result.Error)
			return
		}

		task.Title = taskTitle
		result = db.Save(&task)
		if result.Error != nil {
			fmt.Println("タスクの更新に失敗しました:", result.Error)
			return
		}

		fmt.Printf("=== Update Task ID %d ===", task.ID)
	},
}

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task as done",
	Example: "togo update [id]",
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]

		db, err := db.ConnectDB()
		if err != nil {
			fmt.Println("データベースに接続できませんでした:", err)
			return
		}

		var task models.Task
		result := db.First(&task, taskID)
		if result.Error != nil {
			fmt.Println("タスクの取得に失敗しました:", result.Error)
			return
		}

		task.Status = models.StatusDone
		result = db.Save(&task)
		if result.Error != nil {
			fmt.Println("タスクの更新に失敗しました:", result.Error)
			return
		}

		fmt.Println("=== Done Task ===")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(doneCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
