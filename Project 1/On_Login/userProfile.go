package On_Login

import (
	"FileHandling/models"
	"encoding/json"
	"fmt"
	"os"
)

var users []models.User

func loadUserDetails() error {
	const userFilename = "users.json"
	file, err := os.ReadFile(userFilename)
	if err != nil {
		if os.IsNotExist(err) {
			users = []models.User{}
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &users)
}

// Display user profile
func UserProfile() {
	if err := loadUserDetails(); err != nil {
		fmt.Printf("Error loading user details: %v\n", err)
		return
	}

	var foundUser *models.User
	for _, user := range users {
		if user.Username == ActiveUser.Username {
			foundUser = &user
			break
		}
	}

	if foundUser == nil {
		fmt.Println("User not found.")
		return
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("\033[1;36m---------------------------------------------\033[0m") // Sky blue
	fmt.Println("\033[1;34m               USER PROFILE                 \033[0m")  // Blue
	fmt.Println("\033[1;36m---------------------------------------------\033[0m")
	fmt.Printf("\nUsername     : %s\n", foundUser.Username)
	fmt.Printf("Full Name      : %s\n", foundUser.FullName)
	fmt.Printf("Mobile Number  : %s\n", foundUser.MobileNumber)
	fmt.Printf("Gender         : %s\n", foundUser.Gender)
	fmt.Println("\033[1;36m---------------------------------------------\033[0m")
	fmt.Println()
	fmt.Println()
}
