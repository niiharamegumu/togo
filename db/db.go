package db

import (
	"fmt"
	"os"
	"path/filepath"

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

	var dbPath string
	rootPath := os.Getenv("TOGO_PROJECT_ROOT_PATH")

	if rootPath != "" {
		dbPath = filepath.Join(rootPath, "tasks.db")
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("failed to get user home directory: %w", err)
		}
		togoDir := filepath.Join(home, ".togo")
		if _, err := os.Stat(togoDir); os.IsNotExist(err) {
			if err := os.MkdirAll(togoDir, 0755); err != nil {
				return nil, fmt.Errorf("failed to create directory %s: %w", togoDir, err)
			}
		}
		dbPath = filepath.Join(togoDir, "tasks.db")
	}

	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB.AutoMigrate(&models.Task{})

	return DB, nil
}
