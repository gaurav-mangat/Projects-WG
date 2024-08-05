package models

// Define the struct for user registration

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"` // Store hashed password
	FullName     string `json:"full_name"`
	MobileNumber string `json:"mobile_number"`
	Gender       string `json:"gender"`
}

// Define the structure for daily status

type DailyStatus struct {
	Username string `json:"username"`
	Status   []struct {
		Date        string `json:"date"`
		DailyStatus string `json:"dailyStatus"`
	} `json:"status"`
}
