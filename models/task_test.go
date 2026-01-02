package models

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRenderTasksTable(t *testing.T) {
	// Capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	now := time.Now()
	tasks := []Task{
		{
			ID:        1,
			CreatedAt: now,
			UpdatedAt: now,
			Title:     "Test Task",
			Priority:  10,
			DueDate:   now.Add(24 * time.Hour),
		},
	}

	RenderTasksTable(tasks)

	// Restore and read
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC

	assert.Contains(t, out, "Test Task")
	assert.Contains(t, out, "10")
	assert.Contains(t, out, "ID")
}

func TestTask_RenderTaskTable(t *testing.T) {
	// Simple wrapper test
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	task := &Task{
		Title: "Single Task",
	}
	task.RenderTaskTable()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	assert.Contains(t, buf.String(), "Single Task")
}

func TestDueDateStyling(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	pastDate := time.Now().Add(-24 * time.Hour)
	tasks := []Task{
		{
			Title:   "Past Task",
			DueDate: pastDate,
		},
	}

	RenderTasksTable(tasks)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	out := buf.String()

	// Check for red color code \x1b[31m
	assert.True(t, strings.Contains(out, "\x1b[31m"), "Should contain red color code for past due date")
}
