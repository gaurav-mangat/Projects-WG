package Login_and_Signup

import (
	ol "FileHandling/On_Login"
	ab "FileHandling/utils"
	"fmt"
)

func Login() {
	const filename = "users.json"

	for {
		// Load users from the file
		if err := ab.LoadUsers(filename); err != nil {
			fmt.Printf("Error loading users: %v\n", err)
			return
		}

		var username string

		fmt.Print("\n\nEnter username  : ")
		fmt.Scan(&username)

		var password string
		fmt.Print("Enter password: ")
		fmt.Scan(&password)

		// Check credentials
		loginSuccessful := false
		for _, user := range ab.Users {
			if user.Username == username && ab.CheckPasswordHash(password, user.PasswordHash) {
				fmt.Println("Login successful!")
				loginSuccessful = true
				break
			}
		}

		if loginSuccessful {
			ol.LoggedInMenu() // Calling the function LoggedInMenu() after successful login
		}

		fmt.Println("Login failed. Please check your username and password.")

		// After unsuccesful login attempt

		fmt.Println("\nWhat would you like to do next?")
		fmt.Println("1. Retrying LogIn")
		fmt.Println("2. Sign up")
		fmt.Println("3. Exit")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		if choice == 1 {
			Login()
		} else if choice == 2 {
			SignUp() // Call the SignUp function
		} else if choice == 3 {
			fmt.Println("Exiting...")
			return // Exit the function to end the program
		} else {
			fmt.Println("Invalid choice. Exiting...")
			return
		}
	}
}
