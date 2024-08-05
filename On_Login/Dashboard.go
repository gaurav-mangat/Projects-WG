package On_Login

import (
	a "FileHandling/utils"
	"fmt"
	"strconv"
)

func LoggedInMenu() {
	for {
		fmt.Println("\n\nPress 1 to Add Daily Status")
		fmt.Println("----------------------")
		fmt.Println("Press 2 to Manage To-Do List")
		fmt.Println("----------------------")
		fmt.Println("Press 3 to Check Course Progress")
		fmt.Println("----------------------")
		fmt.Println("Press 4 to Logout")
		fmt.Println()

		choiceStr := a.ReadInput("Enter your choice: ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid choice. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			AddDailyStatus()
		case 2:
			fmt.Println("Press 1 to Add a Task")
			fmt.Println("Press 2 to Mark Task as Complete")
			subChoiceStr := a.ReadInput("Enter your choice: ")
			subChoice, err := strconv.Atoi(subChoiceStr)
			if err != nil {
				fmt.Println("Invalid choice. Please enter a number.")
				continue
			}
			if subChoice == 1 {
				AddTodoItem()
			} else if subChoice == 2 {
				MarkTodoComplete()
			} else {
				fmt.Println("Invalid choice. Returning to main menu.")
			}
		case 3:
			fmt.Println("Press 1 to Mark Section Complete")
			fmt.Println("Press 2 to Display Progress")
			subChoiceStr := a.ReadInput("Enter your choice: ")
			subChoice, err := strconv.Atoi(subChoiceStr)
			if err != nil {
				fmt.Println("Invalid choice. Please enter a number.")
				continue
			}
			if subChoice == 1 {
				MarkCourseSectionComplete()
			} else if subChoice == 2 {
				DisplayCourseProgress()
			} else {
				fmt.Println("Invalid choice. Returning to main menu.")
			}
		case 4:
			fmt.Println("Logging out...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
