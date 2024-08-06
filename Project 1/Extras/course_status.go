package On_Login

import (
	a "FileHandling/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Define the structure for course section
type CourseSection struct {
	SectionName string `json:"section_name"`
	Completed   bool   `json:"completed"`
}

// Define a global variable for course sections
var courseSections []CourseSection

// Load course sections from JSON file
func loadCourseSections(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist, no sections to load
			courseSections = []CourseSection{}
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &courseSections)
}

// Save course sections to JSON file
func saveCourseSections(filename string) error {
	file, err := json.MarshalIndent(courseSections, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, file, 0644)
}

// Mark a course section as complete

func MarkCourseSectionComplete() {
	const filename = "course_sections.json"

	// Load course sections from the file
	if err := loadCourseSections(filename); err != nil {
		fmt.Printf("Error loading course sections: %v\n", err)
		return
	}

	for i, section := range courseSections {
		fmt.Printf("%d: %s (Completed: %v)\n", i+1, section.SectionName, section.Completed)
	}

	indexStr := a.ReadInput("Enter the number of the section to mark as complete: ")
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 1 || index > len(courseSections) {
		fmt.Println("Invalid section number.")
		return
	}

	courseSections[index-1].Completed = true

	if err := saveCourseSections(filename); err != nil {
		fmt.Printf("Error saving course sections: %v\n", err)
	} else {
		fmt.Println("Course section marked as complete!")
	}
}

// Display course progress
func DisplayCourseProgress() {
	const filename = "course_sections.json"

	// Load course sections from the file
	if err := loadCourseSections(filename); err != nil {
		fmt.Printf("Error loading course sections: %v\n", err)
		return
	}

	totalSections := len(courseSections)
	completedSections := 0

	for _, section := range courseSections {
		if section.Completed {
			completedSections++
		}
	}

	progress := float64(completedSections) / float64(totalSections) * 100
	fmt.Printf("Course Progress: %.2f%%\n", progress)
}
