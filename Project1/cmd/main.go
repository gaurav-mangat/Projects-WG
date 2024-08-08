package main

import (
	li "FileHandling/Login_and_Signup"
	"FileHandling/utils"
	"fmt"
	"strconv"
)

func main() {
	for {
		// Welcome message with ASCII art
		fmt.Println("\033[1;36m") // Cyan bold
		fmt.Println("===================================")
		fmt.Println("     WELCOME TO OUR APPLICATION    ")
		fmt.Println("===================================")
		fmt.Println("\033[0m") // Reset color

		// Menu options with borders and colors

		fmt.Println("\033[1;32m") // Green bold
		fmt.Println("-----------------------------")
		fmt.Println("      Press 1 to Login")
		fmt.Println("-----------------------------")
		fmt.Println("      Press 2 to SignUp")
		fmt.Println("-----------------------------")
		fmt.Println("      Press 3 to Exit")
		fmt.Println("-----------------------------")
		fmt.Println("\033[0m") // Reset color

		var choice int

		// Prompt user for choice with color

		fmt.Print("\033[1;34m     Enter your choice: \033[0m") // Blue bold
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("\033[1;31m     Error reading input:", err, "\033[0m") // Red bold
			continue
		}

		switch choice {
		case 1:
			li.Login()
		case 2:
			var yearString string
			var yearInt int
			//fmt.Print("\033[1;34m\n     Enter your birth year: \033[0m") // Blue bold
			yearString = utils.ReadInput("\033[1;34m\n     Enter your birth year: \033[0m")
			yearInt, err = strconv.Atoi(yearString)

			if yearInt < 2004 && yearInt > 1910 {
				fmt.Println("\033[1;33m\n\n     Hurray! You are eligible to sign up 😃😃\n\033[0m") // Yellow bold
				li.SignUp()
			} else {
				fmt.Println("\033[1;31m\n\n  Invalid Birth Year!!! \033[0m")                  // Red bold
				fmt.Println("\033[1;31m\n\n  Sorry! You are not eligible for SignUp \033[0m") // Red bold
			}
		case 3:
			fmt.Println("\033[1;32m\n  Successfully exited the program....\033[0m") // Green bold
			return
		default:
			fmt.Println("\033[1;31m\n     Invalid choice\033[0m") // Red bold
		}
	}
}
