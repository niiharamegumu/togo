package db

import (
	"fmt"
	"os"

	"github.com/niiharamegumu/togo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}

	rootPath := os.Getenv("TOGO_PROJECT_ROOT_PATH")
	if rootPath == "" {
		fmt.Println("Please set TOGO_PROJECT_ROOT_PATH in the .env file")
		return nil, fmt.Errorf("The environment variable TOGO_PROJECT_ROOT_PATH is not set")
	}

	dbPath := fmt.Sprintf("%s/%s", os.Getenv("TOGO_PROJECT_ROOT_PATH"), "tasks.db")
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB.AutoMigrate(&models.Task{})

	return DB, nil
}
