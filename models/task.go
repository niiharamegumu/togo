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
	DueDate  time.Time `gorm:"type:datetime"`
}

const (
	StatusPending = "Pending"
	StatusDone    = "Done"
)

var TaskTableHeader = []string{"ID", "Title", "Status", "Priority", "Created", "Updated", "DueDate"}

func (task *Task) RenderTaskTable() {
	RenderTasksTable([]Task{*task})
}

func RenderTasksTable(tasks []Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(TaskTableHeader)

	n := len(tasks)
	for i, task := range tasks {
		var dueDateStr string
		if !task.DueDate.IsZero() && task.DueDate.Before(time.Now()) {
			dueDateStr = fmt.Sprintf("\x1b[31m%s\x1b[0m", task.DueDate.Format("2006/01/02"))
		} else if !task.DueDate.IsZero() {
			dueDateStr = task.DueDate.Format("2006/01/02")
		} else {
			dueDateStr = ""
		}

		table.Append([]string{
			fmt.Sprintf("%d", task.ID),
			task.Title,
			task.Status,
			fmt.Sprintf("%d", task.Priority),
			task.CreatedAt.Format("2006/01/02 15:04"),
			task.UpdatedAt.Format("2006/01/02 15:04"),
			dueDateStr,
		})

		// Add empty row between tasks for better readability if multiple
		if n > 1 && i != n-1 {
			emptyRow := make([]string, len(TaskTableHeader))
			table.Append(emptyRow)
		}
	}

	table.Render()
}
