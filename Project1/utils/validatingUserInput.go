package utils

import (
	"fmt"
	"strings"
)

func IsValidInput(input string) bool {
	if strings.Contains(input, " ") {

		return false
	}
	return true
}

func IsValidInput2(input string) bool {
	if strings.Contains(input, " ") {
		fmt.Println("\033[1;31m\nInvalid Input\033[0m")
		fmt.Println("\nTry again....")
		return false
	}
	return true
}
