package db

import (
	"github.com/niiharamegumu/togo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

// ConnectDB データベースに接続する関数
func ConnectDB() (*gorm.DB, error) {
	if DB == nil {
		DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	}

	DB.AutoMigrate(&models.Task{})

	return DB, nil
}
