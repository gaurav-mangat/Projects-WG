package utils

import (
	"FileHandling/models"
	"encoding/json"
	"os"
	"testing"
)

func TestLoadUsers(t *testing.T) {
	// Create a test file with user data
	testFile := "test_users.json"
	defer os.Remove(testFile)

	users := []models.User{
		{Username: "user1", PasswordHash: "hash1", FullName: "User One", MobileNumber: "9876543210", Gender: "M"},
		{Username: "user2", PasswordHash: "hash2", FullName: "User Two", MobileNumber: "8765432109", Gender: "F"},
	}
	fileContent, err := json.Marshal(users)
	if err != nil {
		t.Fatalf("Error marshaling users: %v", err)
	}
	err = os.WriteFile(testFile, fileContent, 0644)
	if err != nil {
		t.Fatalf("Error writing test file: %v", err)
	}

	// Test LoadUsers
	err = LoadUsers(testFile)
	if err != nil {
		t.Errorf("LoadUsers returned an error: %v", err)
	}

	// Verify loaded users
	if len(Users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(Users))
	}
	if Users[0].Username != "user1" || Users[1].Username != "user2" {
		t.Errorf("Users not loaded correctly: %+v", Users)
	}
}

func TestSaveUsers(t *testing.T) {
	// Set up users to save
	Users = []models.User{
		{Username: "saveUser1", PasswordHash: "saveHash1", FullName: "Save User One", MobileNumber: "1111111111", Gender: "M"},
		{Username: "saveUser2", PasswordHash: "saveHash2", FullName: "Save User Two", MobileNumber: "2222222222", Gender: "F"},
	}

	// Create a test file to save users
	testFile := "test_save_users.json"
	defer os.Remove(testFile)

	// Test SaveUsers
	err := SaveUsers(testFile)
	if err != nil {
		t.Errorf("SaveUsers returned an error: %v", err)
	}

	// Verify saved users
	fileContent, err := os.ReadFile(testFile)
	if err != nil {
		t.Errorf("Error reading test file: %v", err)
	}
	var savedUsers []models.User
	err = json.Unmarshal(fileContent, &savedUsers)
	if err != nil {
		t.Errorf("Error unmarshaling saved users: %v", err)
	}

	if len(savedUsers) != 2 || savedUsers[0].Username != "saveUser1" || savedUsers[1].Username != "saveUser2" {
		t.Errorf("Users not saved correctly: %+v", savedUsers)
	}
}

func TestIsUsernameUnique(t *testing.T) {
	Users = []models.User{
		{Username: "existingUser"},
	}

	tests := []struct {
		username string
		expected bool
	}{
		{"newUser", true},
		{"existingUser", false},
	}

	for _, test := range tests {
		result := IsUsernameUnique(test.username)
		if result != test.expected {
			t.Errorf("IsUsernameUnique(%q) = %v; want %v", test.username, result, test.expected)
		}
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"Password123!", true},
		{"password", false},
		{"PASSWORD123!", false},
		{"Password", false},
		{"Password123", false},
		{"Passw@1", false},
	}

	for _, test := range tests {
		result := IsValidPassword(test.password)
		if result != test.expected {
			t.Errorf("IsValidPassword(%q) = %v; want %v", test.password, result, test.expected)
		}
	}
}

func TestHashPassword(t *testing.T) {
	password := "Password123!"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned an error: %v", err)
	}

	if !CheckPasswordHash(password, hash) {
		t.Errorf("HashPassword and CheckPasswordHash did not match for password: %s", password)
	}
}

func TestIsValidMobileNumber(t *testing.T) {
	tests := []struct {
		number   string
		expected bool
	}{
		{"9876543210", true},
		{"1234567890", false},
		{"987654321", false},
		{"98765432100", false},
		{"abcdefghij", false},
	}

	for _, test := range tests {
		result := IsValidMobileNumber(test.number)
		if result != test.expected {
			t.Errorf("IsValidMobileNumber(%q) = %v; want %v", test.number, result, test.expected)
		}
	}
}
