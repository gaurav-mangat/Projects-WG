package utils

import (
	"FileHandling/models"
	"os"
	"testing"
)

const testProgressFile = "test_progress.json"

// Helper function to delete the test file
func cleanup() {
	os.Remove(testProgressFile)
}

func TestLoadProgress(t *testing.T) {
	// Create a valid JSON file for testing
	validJSON := `[{"username":"testuser","course_progress":{},"overall_completion":0.0,"course_completion":{}}]`
	err := os.WriteFile(testProgressFile, []byte(validJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write test JSON file: %v", err)
	}
	defer cleanup()

	err = LoadProgress(testProgressFile)
	if err != nil {
		t.Fatalf("Failed to load progress: %v", err)
	}
}

func TestSaveProgress(t *testing.T) {
	// Save a sample progress
	UserProgress = []models.Progress{
		{
			Username: "testuser",
			CourseProgress: map[int][]float64{
				1: {101.1, 101.2},
			},
			OverallCompletion: 50.0,
			CourseCompletion: map[int]float64{
				1: 50.0,
			},
		},
	}
	defer cleanup()

	err := SaveProgress(testProgressFile)
	if err != nil {
		t.Fatalf("Failed to save progress: %v", err)
	}
}

func TestInitializeUserProgress(t *testing.T) {
	defer cleanup()

	// Initialize user progress and check results
	username := "newuser"
	progress := InitializeUserProgress(username)
	if progress == nil {
		t.Fatalf("Expected non-nil user progress")
	}
	if progress.Username != username {
		t.Fatalf("Expected username %s, got %s", username, progress.Username)
	}
}
