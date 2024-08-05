package function

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
	"strings"
)

// Define the user structure with additional fields

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"` // Store hashed password
	FullName     string `json:"full_name"`
	MobileNumber string `json:"mobile_number"`
	Gender       string `json:"gender"`
}

// Define a global variable for user storage

var Users []User

// Create a buffered reader

var Reader = bufio.NewReader(os.Stdin)

// Read input from user with prompt

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	input, err := Reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	return strings.TrimSpace(input)
}

// Load users from JSON file

func LoadUsers(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist, no users to load
			Users = []User{}
			return nil
		}
		return err
	}

	// Check if the file is empty
	if len(file) == 0 {
		Users = []User{}
		return nil
	}

	// Try to unmarshal the file content into users slice

	if err := json.Unmarshal(file, &Users); err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		Users = []User{} // Reset Users to empty slice
		return err
	}
	return nil
}

// Save Users to JSON file

func SaveUsers(filename string) error {
	file, err := json.MarshalIndent(Users, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, file, 0644)
}

// Checking if the Username is unique or not

func IsUsernameUnique(username string) bool {
	for _, user := range Users {
		if user.Username == username {
			return false
		}
	}
	return true
}

//  Validating the password for given specification

func IsValidPassword(password string) bool {
	var (
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString
		hasLower   = regexp.MustCompile(`[a-z]`).MatchString
		hasNumber  = regexp.MustCompile(`[0-9]`).MatchString
		hasSpecial = regexp.MustCompile(`[!@#\$%\^&\*\(\)_+\-=\[\]\;:'",.<>?/|\\]`).MatchString
	)

	if len(password) > 12 && hasUpper(password) && hasLower(password) && hasNumber(password) && hasSpecial(password) {
		return true
	}
	return false
}

// Hash a password using bcrypt

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Compare a hashed password with a plaintext password

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return true
	}
	return false
}
