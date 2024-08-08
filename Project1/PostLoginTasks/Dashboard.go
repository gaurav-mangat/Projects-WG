package PostLoginTasks

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
		fmt.Println("\n\033[1;34m     1. Daily Status Section\033[0m")
		fmt.Println()
		fmt.Println("\033[1;34m     2. Manage To-Do List\033[0m")
		fmt.Println()
		fmt.Println("\033[1;34m     3. Courses Section\033[0m")
		fmt.Println() //
		fmt.Println("\033[1;34m     4. Your Profile\033[0m")
		fmt.Println()
		fmt.Println("\033[1;34m     5. Exit\033[0m")
		fmt.Println()

		var choice int
		fmt.Print("     Enter your choice: ")
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
