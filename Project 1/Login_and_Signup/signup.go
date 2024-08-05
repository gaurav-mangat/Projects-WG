package Login_and_Signup

import (
	"FileHandling/models"
	"FileHandling/utils"
	"fmt"
	"regexp"
	"strings"
)

// Sign up a new user

func SignUp() {
	const filename = "users.json"

	// Load users from the file
	if err := utils.LoadUsers(filename); err != nil {
		fmt.Printf("\033[1;31mError loading users: %v\033[0m\n", err) // Red bold
		return
	}

	// Signup form

	fmt.Println()
	fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m")    // Sky blue
	fmt.Println("\033[1;31m                       SIGN UP FORM                                \033[0m") // Red bold
	fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m")
	fmt.Println()
	username := utils.ReadInput("\033[1;34mEnter username (Username should only be a single word): \033[0m")
	if strings.Contains(username, " ") {
		fmt.Println("\033[1;31m\nInvalid username: Username should only be a single word.\033[0m")
		fmt.Println("\nTry again....")
		return
	}

	if !utils.IsUsernameUnique(username) {
		fmt.Println("\033[1;31mUsername already exists... Please use another.\033[0m") // Red bold
		return
	}

	password := utils.ReadInput("\033[1;34mEnter password (min 9 chars, include lowercase, uppercase, numbers, special): \033[0m")
	if strings.Contains(password, " ") {
		fmt.Println("\033[1;31m\nYou can't use space in password.\033[0m")
		fmt.Println("\nTry again....")
		return
	}

	if !utils.IsValidPassword(password) {
		fmt.Println("\033[1;31mPassword does not meet complexity requirements.\033[0m") // Red bold
		return
	}

	fullName := utils.ReadInput("\033[1;34mEnter full name: \033[0m")

	var mobileNumber string
	valid := false

	for !valid {
		mobileNumber = utils.ReadInput("\033[1;34mEnter mobile number: \033[0m")
		if isValidMobileNumber(mobileNumber) {
			valid = true

		} else {
			fmt.Println("\033[1;31mInvalid mobile number. Please enter a 10-digit number starting with 6, 7, 8, or 9.\033[0m")
		}
	}

	gender := utils.ReadInput("\033[1;34mEnter gender (Male/Female/Other): \033[0m")

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Printf("\033[1;31mError hashing password: %v\033[0m\n", err) // Red bold
		return
	}

	utils.Users = append(utils.Users, models.User{
		Username:     username,
		PasswordHash: hashedPassword,
		FullName:     fullName,
		MobileNumber: mobileNumber,
		Gender:       gender,
	})

	if err := utils.SaveUsers(filename); err != nil {
		fmt.Printf("\033[1;31mError saving users: %v\033[0m\n", err) // Red bold
	} else {
		fmt.Println("\033[1;32m\n\nUser signed up successfully!\033[0m") // Green bold
		fmt.Println("\n\nPress 1 to Login \nPress 2 to Exit")
		var choice int
		fmt.Print("\033[1;34m\nEnter your choice: \033[0m") // Blue bold
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Printf("\033[1;31mError reading choice: %v\033[0m\n", err) // Red bold
			return
		}

		switch choice {
		case 1:
			Login()
		case 2:
			return
		default:
			fmt.Println("\033[1;31mInvalid choice\033[0m") // Red bold
		}
	}
}

// Function to validate mobile number
func isValidMobileNumber(number string) bool {
	match, _ := regexp.MatchString(`^[6-9]\d{9}$`, number)
	return match
}
