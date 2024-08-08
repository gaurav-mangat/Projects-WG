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

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// UserTasks struct to hold tasks for each user

type UserTasks struct {
	Username string `json:"username"`
	Tasks    []Task `json:"tasks"`
}

// Defining the structure for Course list

//type Course1 struct {
//	Courses []struct {
//		ID      int    `json:"id"`
//		Title   string `json:"title"`
//		Lessons []struct {
//			ID    float32 `json:"id"`
//			Title string  `json:"title"`
//		} `json:"lessons"`
//	} `json:"courses"`
//}

type Lesson struct {
	ID    float64 `json:"id"`
	Title string  `json:"title"`
}

// Define the Course struct

type Course struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	Lessons []Lesson `json:"lessons"`
}

// Define the CoursesData struct to hold the array of courses

type CoursesData struct {
	Courses []Course `json:"courses"`
}

type Progress struct {
	Username          string            `json:"username"`
	CourseProgress    map[int][]float64 `json:"course_progress"`    // Course ID to list of completed Lesson IDs
	OverallCompletion float64           `json:"overall_completion"` // Overall completion percentage
	CourseCompletion  map[int]float64   `json:"course_completion"`  // Course ID to completion percentage
}
