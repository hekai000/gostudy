package leetcode100

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"babad", "bab"},               // Happy path
		{"cbbd", "bb"},                 // Happy path
		{"a", "a"},                     // Single character
		{"", ""},                       // Empty string
		{"abcdefg", "a"},               // No palindrome longer than 1
		{"aibohphobia", "aibohphobia"}, // Full palindrome
		{"racecar", "racecar"},         // Odd length palindrome
		// {"A man, a plan, a canal: Panama", "ana"}, // Non-alphanumeric characters
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := longestPalindrome(test.input)
			if result != test.expected {
				t.Errorf("expected %s, got %s", test.expected, result)
			}
		})
	}
}
