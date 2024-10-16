package leetcode100

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},   // happy path
		{"bbbbb", 1},      // all characters same
		{"pwwkew", 3},     // mix of repeated and unique characters
		{"", 0},           // empty string
		{"a", 1},          // single character
		{"abcdefg", 7},    // all unique characters
		{"aabbcc", 2},     // repeating pairs
		{"", 0},           // edge case empty string again
		{"abcdabcdea", 5}, // long string with repetition
	}

	for _, test := range tests {
		result := lengthOfLongestSubstring(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d but got %d", test.input, test.expected, result)
		}
	}
}
