package main

import (
	li "FileHandling/Login_and_Signup"
	"fmt"
)

func main() {
	for {
		// Message to the user
		fmt.Println("\n\nPress 1 to Login")
		fmt.Println("----------------------")
		fmt.Println("Press 2 to SignUp")
		fmt.Println("----------------------")
		fmt.Println("Press 3 to exit")
		fmt.Println()

		var choice int

		// Using fmt.Scanln to read an integer input
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
			// Clear the buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		switch choice {
		case 1:
			li.Login()
		case 2:
			var year int
			fmt.Print("Enter your birth year: ")
			_, err := fmt.Scanln(&year)
			if err != nil {
				fmt.Println("Error reading input:", err)
				// Clear the buffer
				var discard string
				fmt.Scanln(&discard)
				continue
			}
			if year < 2004 {
				fmt.Println("You are eligible to sign up\n")
				li.SignUp()
			} else {
				fmt.Println("You are too young for this.....")
			}
		case 3:
			fmt.Println("Successfully exited the program....")
			return
		default:
			fmt.Println("\nInvalid choice")
		}
	}
}
