package cmd

import (
	"fmt"
	"testing"

	"github.com/niiharamegumu/togo/models"
	"github.com/stretchr/testify/assert"
)

func TestMarkTaskAsDone(t *testing.T) {
	db := setupTestDB(t)

	db.Create(&models.Task{Title: "Task to complete", Status: models.StatusPending})
	var task models.Task
	db.First(&task)

	output := captureOutput(func() {
		markTaskAsDone(nil, []string{fmt.Sprintf("%d", task.ID)})
	})

	assert.Contains(t, output, "Done Task ID")

	db.First(&task)
	assert.Equal(t, models.StatusDone, task.Status)
}

func TestDeleteTask(t *testing.T) {
	db := setupTestDB(t)

	db.Create(&models.Task{Title: "Task to delete"})
	var task models.Task
	db.First(&task)
	taskID := task.ID

	output := captureOutput(func() {
		deleteTask(nil, []string{fmt.Sprintf("%d", taskID)})
	})

	assert.Contains(t, output, "Deleted Task ID")

	result := db.First(&task, taskID)
	assert.Error(t, result.Error, "Task should be deleted")
}
