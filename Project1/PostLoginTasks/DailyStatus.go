package PostLoginTasks

import (
	"FileHandling/Config"
	d "FileHandling/models"
	a "FileHandling/utils"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func init() {
	// Load daily statuses from the file
	if err := loadDailyStatuses("daily_status.json"); err != nil {
		fmt.Printf("Error loading daily statuses: %v\n", err)
		return
	}
}

// Define a global variable for daily statuses
var dailyStatuses []d.DailyStatus

// Load daily statuses from JSON file
func loadDailyStatuses(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist, no statuses to load
			dailyStatuses = []d.DailyStatus{}
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &dailyStatuses)
}

// Save daily statuses to JSON file
func saveDailyStatuses(filename string) error {
	file, err := json.MarshalIndent(dailyStatuses, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, file, 0644)
}

// Add daily status

func AddDailyStatus() {
	const filename = "daily_status.json"

	date := time.Now().Format("2006-01-02")
	status := a.ReadInput("\nEnter daily status: ")

	fmt.Println("Active User: ", ActiveUser.Username)

	// Append new status
	newEntry := struct {
		Date        string `json:"date"`
		DailyStatus string `json:"dailyStatus"`
	}{
		Date:        date,
		DailyStatus: status,
	}

	var found bool
	for i, users := range dailyStatuses {
		if users.Username == ActiveUser.Username {
			// User found, update the status
			dailyStatuses[i].Status = append(dailyStatuses[i].Status, newEntry)
			found = true
			break
		}
	}

	if !found {
		// No user found, create a new entry
		data := d.DailyStatus{
			Username: ActiveUser.Username,
			Status: []struct {
				Date        string `json:"date"`
				DailyStatus string `json:"dailyStatus"`
			}{
				{Date: date, DailyStatus: status},
			},
		}
		dailyStatuses = append(dailyStatuses, data)
	}

	if err := saveDailyStatuses(Config.DailyStatusFile); err != nil {
		fmt.Printf("Error saving daily statuses: %v\n", err)
	} else {
		fmt.Println("Daily status added successfully!")
	}
}

// View daily status
func viewDailyStatus() {
	date := a.ReadInput("\nEnter the date (YYYY-MM-DD) to view daily status: ")
	var count int
	var found bool
	for _, users := range dailyStatuses {
		if users.Username == ActiveUser.Username {
			for _, statusEntry := range users.Status {
				if statusEntry.Date == date {
					count++
					if count == 1 {
						fmt.Printf("\033[1;34m\nDate: %s\033[0m\n", statusEntry.Date)
					}
					fmt.Printf("\033[1;32mStatus %d: %s\033[0m\n", count, statusEntry.DailyStatus)
					found = true

				}
			}
			break
		}
	}

	if !found {
		fmt.Println("\033[1;31mNo status found for the specified date.\033[0m") // Red
	}
}

// Daily Status Section Menu

func DailyStatusSection() {
	for {
		fmt.Println()
		fmt.Println()
		fmt.Println("\033[1;36m---------------------------------------------\033[0m") // Sky blue
		fmt.Println("\033[1;34m          DAILY STATUS SECTION             \033[0m")   // Blue
		fmt.Println("\033[1;36m---------------------------------------------\033[0m")
		fmt.Println()
		fmt.Println("1. Add Daily Status")
		fmt.Println("2. View Daily Status")
		fmt.Println("3. Back to Dashboard")
		fmt.Println()

		var choice int
		fmt.Print("\033[1;34mEnter your choice: \033[0m") // Blue

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("\033[1;31mError reading input. Please enter a valid number.\033[0m") // Red
			continue
		}

		switch choice {
		case 1:
			AddDailyStatus()
		case 2:
			viewDailyStatus()
		case 3:
			return // Exit the function
		default:
			fmt.Println("\033[1;31mInvalid choice. Please select a valid option.\033[0m") // Red
		}
	}
}
