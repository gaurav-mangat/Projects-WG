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

// Defining the structure for Course list

type Course struct {
	Courses []struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Lessons []struct {
			ID    float32 `json:"id"`
			Title string  `json:"title"`
		} `json:"lessons"`
	} `json:"courses"`
}

type Progress struct {
	Username string           `json:"username"`
	Courses  []CourseProgress `json:"courses"`
}

type CourseProgress struct {
	CourseID         int       `json:"course_id"`
	CompletedLessons []float32 `json:"completed_lessons"`
}
