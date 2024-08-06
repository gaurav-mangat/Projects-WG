package Login_and_Signup

import (
	ol "FileHandling/On_Login"
	"FileHandling/utils"
	"fmt"
	"strings"
)

func Login() {
	const filename = "users.json"

	// Load users from the file once
	if err := utils.LoadUsers(filename); err != nil {
		fmt.Printf("Error loading users: %v\n", err)
		return
	}

	for {
		var username, password string
		fmt.Println()
		fmt.Println()
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m") // Sky blue
		fmt.Println("\033[1;31m                          LOG IN                                \033[0m") // Red bold
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m")

		username = utils.ReadInput("\n             Enter username: ")
		if strings.Contains(username, " ") {
			fmt.Println("\033[1;31m            User name is only of one word.\033[0m")
			fmt.Println("\nPlease try again....")
			return
		}

		password = utils.ReadInput("             Enter password: ")

		if strings.Contains(password, " ") {
			fmt.Println("\033[1;31mP\nassword doesn't contain any spaces.\033[0m")
			fmt.Println("\nPlease try again....")
			return
		}
		fmt.Println()

		// Check credentials
		loginSuccessful := false
		for _, user := range utils.Users {
			if user.Username == username && utils.CheckPasswordHash(password, user.PasswordHash) {
				fmt.Println("\033[1;31m              Login successful!\033[0m") // Red bold
				loginSuccessful = true
				ol.ActiveUser = user
				ol.Dashboard(user)
				fmt.Println()
				break
			}
		}

		if loginSuccessful {
			// Successful login, exit the loop
			break
		} else {
			// Failed login, prompt the user
			fmt.Println("Login failed. Please check your username and password.")

			fmt.Println("\nWhat would you like to do next?")
			fmt.Println("1. Retry Login")
			fmt.Println("2. Sign up")
			fmt.Println("3. Exit")
			var choice int
			fmt.Print("Enter your choice: ")
			fmt.Scan(&choice)

			switch choice {
			case 1:
				// Retry login
				continue
			case 2:
				// Call the SignUp function
				SignUp()
				return // Return to avoid retrying after sign up
			case 3:
				// Exit
				fmt.Println("Exiting...")
				return
			default:
				// Invalid choice
				fmt.Println("Invalid choice. Exiting...")
				return
			}
		}
	}
}
