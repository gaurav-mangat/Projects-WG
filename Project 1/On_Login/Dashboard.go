package On_Login

import (
	"FileHandling/models"
	a "FileHandling/utils"
	"fmt"
	"strconv"
)

var ActiveUser models.User

func Dashboard(activeUser models.User) {
	for {
		// User Dashboard
		fmt.Println()
		fmt.Println("\033[1;36m---------------------------------------------\033[0m")  // Sky blue
		fmt.Println("\033[1;32m                DASHBOARD                     \033[0m") // Green
		fmt.Println("\033[1;36m---------------------------------------------\033[0m")  // Sky blue
		fmt.Println("\n\033[1;34mPress 1 for Daily Status section\033[0m")
		fmt.Println("\033[1;36m----------------\033[0m") // Sky blue
		fmt.Println("\033[1;34mPress 2 to Manage To-Do List\033[0m")
		fmt.Println("\033[1;36m----------------\033[0m") // Sky blue
		fmt.Println("\033[1;34mPress 3 to Check Course Progress\033[0m")
		fmt.Println("\033[1;36m----------------\033[0m") // Sky blue
		fmt.Println("\033[1;34mPress 4 to Logout\033[0m")
		fmt.Println()

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}

		switch choice {
		case 1:
			DailyStatusSection()
		case 2:
			TaskManagementSection()
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
