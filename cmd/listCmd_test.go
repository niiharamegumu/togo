package cmd

import (
	"testing"

	"github.com/niiharamegumu/togo/models"
	"github.com/stretchr/testify/assert"
)

func TestListTasks(t *testing.T) {
	db := setupTestDB(t)

	// Seed data
	db.Create(&models.Task{Title: "Pending Task"})
	db.Create(&models.Task{Title: "Done Task"})

	// Test default list (all)
	output := captureOutput(func() {
		listTasks(nil, []string{})
	})
	assert.Contains(t, output, "Pending Task")
	assert.Contains(t, output, "Done Task")
}
