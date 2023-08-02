package models

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title    string `gorm:"size:255; not null"`
	Status   string `gorm:"size:255; default:'Pending'; not null"`
	Priority int    `gorm:"size:100; default:0; not null; min:0; max:100"`
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
