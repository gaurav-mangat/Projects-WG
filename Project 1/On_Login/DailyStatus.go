package On_Login

import (
	d "FileHandling/models"
	a "FileHandling/utils"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

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

	// Load daily statuses from the file
	if err := loadDailyStatuses(filename); err != nil {
		fmt.Printf("Error loading daily statuses: %v\n", err)
		return
	}
	date := time.Now().Format("2006-01-02")
	//date := a.ReadInput("Enter date (YYYY-MM-DD): ")
	status := a.ReadInput("Enter daily status: ")

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

	if err := saveDailyStatuses(filename); err != nil {
		fmt.Printf("Error saving daily statuses: %v\n", err)
	} else {
		fmt.Println("Daily status added successfully!")
	}
}
