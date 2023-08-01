package models

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"gorm.io/gorm"
)

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

func (task *Task) RenderTaskTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Status", "Date"})
	table.Append([]string{
		fmt.Sprintf("%d", task.ID),
		task.Title,
		task.Status,
		task.CreatedAt.Format("2006/01/02 15:04"),
	})
	table.Render()
}
