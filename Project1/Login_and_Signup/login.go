package Login_and_Signup

import (
	"FileHandling/Config"
	ol "FileHandling/PostLoginTasks"
	"FileHandling/utils"
	"fmt"
)

func Login() {
	const filename = Config.UserFile

	// Load users from the file once
	if err := utils.LoadUsers(Config.UserFile); err != nil {
		fmt.Printf("Error loading users: %v\n", err)
		return
	}

	attemptsLeft := 3

	for attemptsLeft > 0 {
		var username, password string
		fmt.Println()
		fmt.Println()
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m") // Sky blue
		fmt.Println("\033[1;31m                          LOG IN                                \033[0m") // Red bold
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m")

		username = utils.ReadInput("\n             Enter username: ")
		if !utils.IsValidInput2(username) {
			return
		}

		password = utils.ReadInput("             Enter password: ")

		if !utils.IsValidInput(password) {
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
				return // Exit after successful login
			}
		}

		if loginSuccessful {
			// Successful login, exit the loop
			break
		} else {
			// Failed login, decrement the attempts left
			attemptsLeft--
			if attemptsLeft == 0 {
				fmt.Println("\033[1;31mLogin failed. You have exhausted all attempts.\033[0m") // Red bold
				return
			}
			fmt.Printf("Login failed. Please check your username and password. You have %d attempt(s) left.\n", attemptsLeft)

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
