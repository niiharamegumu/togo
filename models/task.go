package models

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title    string    `gorm:"size:255; not null"`
	Status   string    `gorm:"size:255; default:'Pending'; not null"`
	Priority int       `gorm:"size:100; default:0; not null; min:0; max:100"`
	ExpireAt time.Time `gorm:"type:datetime"`
}

const (
	StatusPending = "Pending"
	StatusDone    = "Done"
)

var TaskTableHeader = []string{"ID", "Title", "Status", "Priority", "Created", "Updated", "Expire"}

func (task *Task) RenderTaskTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(TaskTableHeader)
	var expireAtStr string
	if !task.ExpireAt.IsZero() {
		expireAtStr = task.ExpireAt.Format("2006/01/02")
	} else {
		expireAtStr = ""
	}
	table.Append([]string{
		fmt.Sprintf("%d", task.ID),
		task.Title,
		task.Status,
		fmt.Sprintf("%d", task.Priority),
		task.CreatedAt.Format("2006/01/02 15:04"),
		task.UpdatedAt.Format("2006/01/02 15:04"),
		expireAtStr,
	})
	table.Render()
}
