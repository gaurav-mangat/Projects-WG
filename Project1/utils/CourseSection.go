package utils

import (
	"FileHandling/Config"
	"FileHandling/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var UserProgress []models.Progress

// Load progress data from a file
func LoadProgress(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		// If the file doesn't exist, create an empty progress file
		if os.IsNotExist(err) {
			UserProgress = []models.Progress{}
			return SaveProgress(filename)
		}
		return err
	}

	err = json.Unmarshal(file, &UserProgress)
	if err != nil {
		return err
	}

	return nil
}

// Save progress data to a file
func SaveProgress(filename string) error {
	data, err := json.MarshalIndent(UserProgress, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func InitializeUserProgress(username string) *models.Progress {
	const progressFile = "progress.json"

	if err := LoadProgress(Config.ProgressFile); err != nil {
		fmt.Printf("\033[1;31mError loading progress: %v\033[0m\n", err)
		return nil
	}

	// Find the user's progress
	var userProgress *models.Progress
	for i, progress := range UserProgress {
		if progress.Username == username {
			userProgress = &UserProgress[i]
			return userProgress
		}
	}

	// Initialize progress if not found
	userProgress = &models.Progress{
		Username:       username,
		CourseProgress: make(map[int][]float64), // Initialize with an empty map
	}
	UserProgress = append(UserProgress, *userProgress)

	if err := SaveProgress(Config.ProgressFile); err != nil {
		fmt.Printf("\033[1;31mError saving progress: %v\033[0m\n", err)
		return nil
	}

	return userProgress
}
