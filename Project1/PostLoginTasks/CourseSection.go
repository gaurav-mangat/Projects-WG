package PostLoginTasks

import (
	"FileHandling/Config"
	"FileHandling/models"
	"FileHandling/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

//const progressfile = Config.ProgressFile

var course []models.Course

var userProgress *models.Progress

// Function to see course details
func SeeCourseDetails() {
	//const filename = "courses.json"

	// Read the JSON file
	data, err := os.ReadFile(Config.CourseFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Unmarshal the JSON data
	var coursesData models.CoursesData
	err = json.Unmarshal(data, &coursesData)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Print course details
	for _, course := range coursesData.Courses {
		fmt.Printf("\n\nCourse ID: %d\n", course.ID)
		fmt.Printf("Course Title: %s\n", course.Title)
		fmt.Println("Lessons:")
		for _, lesson := range course.Lessons {
			fmt.Printf("  Lesson ID: %.1f\n", lesson.ID)
			fmt.Printf("  Lesson Title: %s\n", lesson.Title)
		}
		fmt.Println()
	}
}

func UpdateProgress(username string) {
	userProgress = utils.InitializeUserProgress(username)
	if userProgress == nil {
		return
	}

	// Load course details to show lessons
	const filename = Config.CourseFile
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("\033[1;31mError reading file: %v\033[0m\n", err)
		return
	}

	var coursesData models.CoursesData
	err = json.Unmarshal(data, &coursesData)
	if err != nil {
		fmt.Printf("\033[1;31mError unmarshalling JSON: %v\033[0m\n", err)
		return
	}

	// Display all courses
	fmt.Println("\033[1;34m\n\nAvailable Courses:\033[0m")
	for _, course := range coursesData.Courses {
		fmt.Printf("  %d : %s\n", course.ID, course.Title)
	}

	// Ask user to select a course ID
	var courseID int
	fmt.Print("\033[1;34m\nEnter Course ID to update: \033[0m")
	fmt.Scan(&courseID)

	// Validate the selected course ID
	var courseLessons []models.Lesson
	var courseFound bool
	for _, course := range coursesData.Courses {
		if course.ID == courseID {
			courseLessons = course.Lessons
			courseFound = true
			break
		}
	}

	if !courseFound {
		fmt.Println("\033[1;31mInvalid Course ID. No lessons found for the given Course ID.\033[0m")
		return
	}

	// Display lessons for the chosen course
	fmt.Println("\033[1;34m\nLessons for Course ID", courseID, ":\033[0m")
	for _, lesson := range courseLessons {
		fmt.Printf("  %.1f : %s\n", lesson.ID, lesson.Title)
	}

	// Loop to allow multiple lesson ID inputs
	for {
		var lessonID float64
		fmt.Print("\033[1;34m\nEnter Lesson ID to mark as completed (or type 0 to finish): \033[0m")
		fmt.Scan(&lessonID)

		// Exit loop if user enters 0
		if lessonID == 0 {
			break
		}

		// Validate the selected lesson ID
		var lessonFound bool
		for _, lesson := range courseLessons {
			if lesson.ID == lessonID {
				lessonFound = true
				break
			}
		}

		if !lessonFound {
			fmt.Println("\033[1;31m\nInvalid Lesson ID. No such lesson found for the given Course ID.\033[0m")
			continue
		}

		// Check if the lesson ID is already completed
		completedLessons := userProgress.CourseProgress[courseID]
		lessonAlreadyCompleted := false
		for _, completedLessonID := range completedLessons {
			if completedLessonID == lessonID {
				lessonAlreadyCompleted = true
				break
			}
		}

		if lessonAlreadyCompleted {
			fmt.Println("\033[1;31m\nLesson ID already completed. Please enter a different Lesson ID.\033[0m")
			continue
		}

		// Update progress
		userProgress.CourseProgress[courseID] = append(userProgress.CourseProgress[courseID], lessonID)
		fmt.Println("\033[1;32m\nLesson ID added successfully!\033[0m")

		// Update progress percentages
		updateProgressPercentages(userProgress, coursesData)
	}

	if err := utils.SaveProgress(Config.ProgressFile); err != nil {
		fmt.Printf("\033[1;31mError saving progress: %v\033[0m\n", err)
		return
	}

	fmt.Println("\033[1;32m\nAll progress updates completed successfully!\033[0m")
}

func updateProgressPercentages(userProgress *models.Progress, coursesData models.CoursesData) {
	totalLessons := 0
	completedLessons := 0
	completedCourses := 0

	userProgress.CourseCompletion = make(map[int]float64)

	for _, course := range coursesData.Courses {
		totalLessons += len(course.Lessons)
		completedLessonsForCourse := len(userProgress.CourseProgress[course.ID])
		userProgress.CourseCompletion[course.ID] = float64(completedLessonsForCourse) / float64(len(course.Lessons)) * 100
		completedLessons += completedLessonsForCourse
		if completedLessonsForCourse == len(course.Lessons) {
			completedCourses++
		}
	}

	if totalLessons > 0 {
		userProgress.OverallCompletion = float64(completedCourses) / float64(len(coursesData.Courses)) * 100
	} else {
		userProgress.OverallCompletion = 0
	}
}

func DisplayProgress(username string) {
	userProgress := utils.InitializeUserProgress(username)
	if userProgress == nil {
		return
	}

	// Load course details to show lesson titles
	const filename = Config.CourseFile
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("\033[1;31mError reading file: %v\033[0m\n", err)
		return
	}

	var coursesData models.CoursesData
	err = json.Unmarshal(data, &coursesData)
	if err != nil {
		fmt.Printf("\033[1;31mError unmarshalling JSON: %v\033[0m\n", err)
		return
	}

	// Display overall and course-specific progress
	fmt.Println("\033[1;34m\n\nUser Progress:\033[0m")
	fmt.Printf("Overall Learning Path Completion: %.2f%%\n", userProgress.OverallCompletion)

	fmt.Println("\033[1;34m\nCompleted Courses:\033[0m")
	for courseID, completion := range userProgress.CourseCompletion {
		if completion == 100 {
			fmt.Printf("Course ID: %d\n", courseID)
		}
	}
	fmt.Println()

	for courseID, lessons := range userProgress.CourseProgress {
		fmt.Printf("Course ID %d : Completed Lessons %v\n", courseID, lessons)
		if completion, exists := userProgress.CourseCompletion[courseID]; exists {
			fmt.Printf("Completion Percentage for Course ID %d: %.2f%%\n", courseID, completion)
		}
		fmt.Println()
	}

	var courseID int
	fmt.Print("\033[1;34m\nEnter Course ID to see detailed progress: \033[0m")
	fmt.Scan(&courseID)

	if lessons, exists := userProgress.CourseProgress[courseID]; exists {
		fmt.Printf("\nDetailed Progress for Course ID %d:\n", courseID)
		for _, lessonID := range lessons {
			var lessonTitle string
			for _, course := range coursesData.Courses {
				if course.ID == courseID {
					for _, lesson := range course.Lessons {
						if lesson.ID == lessonID {
							lessonTitle = lesson.Title
							break
						}
					}
					break
				}
			}
			fmt.Printf("  %.1f : %s\n", lessonID, lessonTitle)
		}
		fmt.Printf("\nCompletion for Course ID %d: %.2f%%\n", courseID, userProgress.CourseCompletion[courseID])
	} else {
		fmt.Println("\033[1;31m\nNo progress found for this course\033[0m")
	}
}

func CourseProgress() {
	for {
		fmt.Println()
		fmt.Println()
		fmt.Println("\033[1;36m---------------------------------------------\033[0m") // Sky blue
		fmt.Println("\033[1;34m             COURSES SECTION             \033[0m")     // Blue
		fmt.Println("\033[1;36m---------------------------------------------\033[0m")
		fmt.Println()

		fmt.Println("\033[1;31m     Press 1 to See the Courses Assigned\033[0m")
		fmt.Println("\033[1;31m     Press 2 to Display Progress\033[0m")
		fmt.Println("\033[1;31m     Press 3 to Update the Progress\033[0m")
		fmt.Println("\033[1;31m     Press 4 to Exit to Dashboard\033[0m") // Option to exit

		subChoiceStr := utils.ReadInput("\n     Enter your choice: ")
		subChoice, err := strconv.Atoi(subChoiceStr)
		if err != nil {
			fmt.Println("Invalid choice. Please enter a number.")
			continue
		}

		switch subChoice {
		case 1:
			SeeCourseDetails()
		case 2:
			DisplayProgress(ActiveUser.Username)
		case 3:
			UpdateProgress(ActiveUser.Username)
		case 4:
			fmt.Println("Exiting to dashboard...")
			return // Exit the loop and return to the dashboard
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}
