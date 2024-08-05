package Login_and_Signup

import "fmt"
import a "FileHandling/utils"

// Sign up a new user

func SignUp() {
	const filename = "users.json"

	// Load users from the file
	if err := a.LoadUsers(filename); err != nil {
		fmt.Printf("Error loading users: %v\n", err)
		return
	}

	var username string
	fmt.Print("\n\nEnter username (Username should only be a single word) : ")
	_, err := fmt.Scan(&username)
	if err != nil {
		fmt.Printf("Error reading username: %v\n", err)
	}
	//
	//password := a.ReadInput("Enter password (Password must have at least 12 characters and should contain lowercase, uppercase, numbers, and special characters): ")
	var password string
	fmt.Print("Enter password   : ")
	fmt.Scan(&password)
	if err != nil {
		fmt.Println("Password input failed: ", err)
		return
	}

	if !a.IsValidPassword(password) {
		fmt.Println("Password does not meet complexity requirements.")
		return
	}
	if !a.IsUsernameUnique(username) {
		fmt.Println("Username already exists... Please use another.")
		return
	}

	fullName := a.ReadInput("Enter full name: ")
	mobileNumber := a.ReadInput("Enter mobile number: ")
	gender := a.ReadInput("Enter gender (Male/Female/Other): ")

	hashedPassword, err := a.HashPassword(password)
	if err != nil {
		fmt.Printf("Error hashing password: %v\n", err)
		return
	}

	a.Users = append(a.Users, a.User{
		Username:     username,
		PasswordHash: hashedPassword,
		FullName:     fullName,
		MobileNumber: mobileNumber,
		Gender:       gender,
	})

	if err := a.SaveUsers(filename); err != nil {
		fmt.Printf("Error saving users: %v\n", err)
	} else {
		fmt.Println("User signed up successfully!")
		fmt.Println("\n\nPress 1 to Login \nPress 2 to exit")
		var ch int
		fmt.Println("Enter your choice :")
		fmt.Scanf("%d", &ch)

		switch ch {
		case 1:
			Login()

		case 2:
			return

		default:
			fmt.Println("Invalid choice")
			break
		}
	}
}
