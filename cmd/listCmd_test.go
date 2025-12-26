package cmd

import (
	"testing"

	"github.com/niiharamegumu/togo/models"
	"github.com/stretchr/testify/assert"
)

func TestListTasks(t *testing.T) {
	db := setupTestDB(t)

	// Seed data
	db.Create(&models.Task{Title: "Pending Task", Status: models.StatusPending})
	db.Create(&models.Task{Title: "Done Task", Status: models.StatusDone})

	// Test default list (pending)
	output := captureOutput(func() {
		listTasks(nil, []string{})
	})
	assert.Contains(t, output, "Pending Task")
	assert.NotContains(t, output, "Done Task")

	// Test list done
	statusFlag = "done"
	output = captureOutput(func() {
		listTasks(nil, []string{"done"})
	})
	assert.Contains(t, output, "Done Task")
	assert.NotContains(t, output, "Pending Task")
	statusFlag = "pen" // reset

	// Test list all
	statusFlag = "all"
	output = captureOutput(func() {
		listTasks(nil, []string{"all"})
	})
	assert.Contains(t, output, "Pending Task")
	assert.Contains(t, output, "Done Task")
	statusFlag = "pen" // reset
}
