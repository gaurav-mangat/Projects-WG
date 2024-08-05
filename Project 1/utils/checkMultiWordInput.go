package utils

import (
	"fmt"
	"strings"
)

func CheckMultiWordInput(input string) bool {
	if strings.Contains(input, " ") {
		fmt.Println("\033[1;31mInvalid username: Username should only be a single word.\033[0m")
		fmt.Println("\nTry again....")
		return true
	}
	return false

}
