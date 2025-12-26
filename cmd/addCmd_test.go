package cmd

import (
	"testing"

	"github.com/niiharamegumu/togo/models"
	"github.com/stretchr/testify/assert"
)

func TestAddTask(t *testing.T) {
	db := setupTestDB(t)

	// Mock input for title and priority
	// Title: "New Task\n\n" (the second empty line to finish)
	// Priority: "50"
	input := "New Task\n\n50\n"

	output := captureOutput(func() {
		mockStdin(input, func() {
			addTask(nil, []string{})
		})
	})

	assert.Contains(t, output, "Task Added")
	assert.Contains(t, output, "New Task")
	assert.Contains(t, output, "50")

	var task models.Task
	db.First(&task)
	assert.Equal(t, "New Task", task.Title)
	assert.Equal(t, 50, task.Priority)
}

func TestAddTask_EmptyTitle(t *testing.T) {
	setupTestDB(t)

	// Empty title (just newline)
	input := "\n"

	output := captureOutput(func() {
		mockStdin(input, func() {
			addTask(nil, []string{})
		})
	})

	assert.Contains(t, output, "Exiting the process")
}

func TestAddTask_InvalidPriority(t *testing.T) {
	db := setupTestDB(t)

	// Invalid priority "abc" should fallback to 0
	input := "Invalid Priority Task\n\nabc\n"

	captureOutput(func() {
		mockStdin(input, func() {
			addTask(nil, []string{})
		})
	})

	var task models.Task
	db.First(&task)
	assert.Equal(t, 0, task.Priority)
}
