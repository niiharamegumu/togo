package models

import "gorm.io/gorm"

// Task データモデル
type Task struct {
	gorm.Model
	Title  string
	Status string
}

const (
	StatusPending = "Pending"
	StatusDone    = "Done"
)
