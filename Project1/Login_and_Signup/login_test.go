package Login_and_Signup_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"FileHandling/Login_and_Signup"
	"FileHandling/models"
	"FileHandling/utils"
)

// TestLogin attempts to indirectly test the Login function by setting up conditions
// and observing outcomes. Due to the nature of the Login function, direct testing
// is challenging without significant refactoring of the original code.
func TestLogin(t *testing.T) {
	// Setup: Create a temporary file with user data
	tmpfile, err := ioutil.TempFile("", "test_users_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up

	users := []models.User{
		{Username: "testuser", PasswordHash: "$2a$10$N9qoE5iKtPdVZy3lLckOPuXwZqvxpMQVvJt5VO/zZb5kvR78kFjTC"}, // bcrypt hash of "testpassword"
	}
	fileContent, err := json.Marshal(users)
	if err != nil {
		t.Fatalf("Failed to marshal users: %v", err)
	}
	err = ioutil.WriteFile(tmpfile.Name(), fileContent, 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Setup: Override Config.UserFile to point to our temp file
	oldUserFile := utils.Config.UserFile
	utils.Config.UserFile = tmpfile.Name()
	defer func() { utils.Config.UserFile = oldUserFile }()

	// Test: Attempt a successful login
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = ioutil.NopCloser(bytes.NewBufferString("testuser\ntestpassword\n"))

	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	var buf bytes.Buffer
	os.Stdout = &buf

	Login_and_Signup.Login()

	output := buf.String()
	if !strings.Contains(output, "Login successful!") {
		t.Errorf("Expected login to succeed, got:\n%s", output)
	}
}
