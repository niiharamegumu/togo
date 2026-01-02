package cmd

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/niiharamegumu/togo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB initializes an in-memory SQLite DB for testing
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&models.Task{})
	dbConn = db // Update global dbConn for cmd package
	return db
}

// captureOutput captures Stdout and returns it as a string
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	return <-outC
}

// mockStdin was removed because huh relies on tea.Program which is harder to mock via simple stdin redirection.
// Interactive commands are now manually verified.
