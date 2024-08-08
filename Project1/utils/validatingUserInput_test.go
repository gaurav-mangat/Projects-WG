package utils

import (
	"testing"
)

// TestIsValidInput tests the IsValidInput function
func TestIsValidInput(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"validinput", true},
		{"invalid input", false},
		{"another_valid_input", true},
		{"  leading_space", false},
		{"trailing_space ", false},
		{" multiple  spaces ", false},
	}

	for _, test := range tests {
		result := IsValidInput(test.input)
		if result != test.expected {
			t.Errorf("IsValidInput(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// TestIsValidInput2 tests the IsValidInput2 function
func TestIsValidInput2(t *testing.T) {
	tests := []struct {
		input         string
		expectedValid bool
	}{
		{"validinput", true},
		{"another_valid_input", true},
		{"trailing_space ", false},
		{" multiple  spaces ", false},
	}

	for _, test := range tests {
		result := IsValidInput2(test.input)
		if result != test.expectedValid {
			t.Errorf("IsValidInput2(%q) = %v; want %v", test.input, result, test.expectedValid)
		}
	}

}
