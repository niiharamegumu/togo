package cmd

import (
	"fmt"
	"os"

	"github.com/niiharamegumu/togo/db"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var rootCmd = &cobra.Command{
	Use:   "togo",
	Short: "Task Management CLI",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		dbConn, err = db.ConnectDB()
		if err != nil {
			return fmt.Errorf("could not connect to the database: %w", err)
		}
		return nil
	},
}

var dbConn *gorm.DB
var dueDate string
var columnsMapping = map[string]string{
	"id":         "id",         // ID
	"i":          "id",         // ID shorthand
	"title":      "title",      // Title
	"t":          "title",      // Title shorthand
	"priority":   "priority",   // Priority
	"p":          "priority",   // Priority shorthand
	"created_at": "created_at", // CreatedAt
	"c":          "created_at", // CreatedAt shorthand
	"updated_at": "updated_at", // UpdatedAt
	"u":          "updated_at", // UpdatedAt shorthand
	"due_date":   "due_date",   // DueDate
	"d":          "due_date",   // DueDate shorthand
}

func init() {
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
