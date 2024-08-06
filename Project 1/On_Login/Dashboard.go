package On_Login

import (
	"FileHandling/models"
	"fmt"
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
		fmt.Println()
		fmt.Println("\033[1;34mPress 2 to Manage To-Do List\033[0m")
		fmt.Println()
		fmt.Println("\033[1;34mPress 3 to Check Course Progress\033[0m")
		fmt.Println() //
		fmt.Println("\033[1;34mPress 4 to see Your Profile\033[0m")
		fmt.Println()
		fmt.Println("\033[1;34mPress 5 to Exit\033[0m")
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
			CourseProgress()
		case 5:
			fmt.Println("Logging out...")
			return
		case 4:
			UserProfile()

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
