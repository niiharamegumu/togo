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

// mockStdin mocks Stdin with the provided input string
func mockStdin(input string, f func()) {
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	stdinScanner = nil // Reset scanner singleton for test

	w.Write([]byte(input))
	w.Close()

	f()
	os.Stdin = oldStdin
	stdinScanner = nil // Reset back
}
