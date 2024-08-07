package utils

import (
	"fmt"
	"regexp"
)

// IisValidInput checks if the input contains only numeric characters

func IisValidInput(input string) bool {
	match, _ := regexp.MatchString(`^\d+$`, input)
	if !match {
		fmt.Println("\033[1;31m\nInvalid Input. Input should contain only numeric characters.\033[0m")
		fmt.Println("\nTry again....")
		return false
	}
	return true
}
