package models

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Task struct {
	ID        uint      `gorm:"primarykey"`
	Title     string    `gorm:"size:255; not null"`
	Priority  int       `gorm:"size:100; default:0; not null; min:0; max:100"`
	DueDate   time.Time `gorm:"type:datetime"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var TaskTableHeader = []string{"ID", "TITLE", "CREATED", "DUEDATE", "UPDATED", "PRIORITY"}

func (task *Task) RenderTaskTable() {
	RenderTasksTable([]Task{*task})
}

func RenderTasksTable(tasks []Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(TaskTableHeader)

	// Set fixed width for Title column (index 1)
	table.SetColMinWidth(1, 50)

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
			task.CreatedAt.Format("2006/01/02 15:04"),
			dueDateStr,
			task.UpdatedAt.Format("2006/01/02 15:04"),
			fmt.Sprintf("%d", task.Priority),
		})

		// Add empty row between tasks for better readability if multiple
		if n > 1 && i != n-1 {
			emptyRow := make([]string, len(TaskTableHeader))
			table.Append(emptyRow)
		}
	}

	table.Render()
}
