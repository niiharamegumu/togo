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
}

var dbConn *gorm.DB

func init() {
	var err error
	dbConn, err = db.ConnectDB()
	if err != nil {
		fmt.Println("🚨 データベースに接続できませんでした:", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
