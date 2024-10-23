package twopointers

import (
	"testing"
)

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	tests := []struct {
		s        string
		k        int
		expected int
	}{
		{"abcabc", 2, 2},    // Happy path
		{"aa", 1, 2},        // Happy path
		{"abc", 3, 3},       // Happy path
		{"abac", 2, 3},      // Happy path
		{"", 0, 0},          // Edge case: empty string
		{"aaaaaaa", 1, 7},   // Edge case: all same characters
		{"abcde", 0, 0},     // Edge case: k is 0
		{"abcde", 1, 1},     // Edge case: k is 1
		{"abcabcabc", 3, 9}, // Edge case: full string with all characters within k
		{"aabbcc", 2, 4},    // Edge case: repeating characters within limit
		{"abcabc", 4, 6},    // Edge case: k greater than number of unique chars
		{"a", 1, 1},         // Edge case: single character string
		{"ab", 1, 1},        // Edge case: two different characters with k=1
	}

	for _, test := range tests {
		result := lengthOfLongestSubstringKDistinct(test.s, test.k)
		if result != test.expected {
			t.Errorf("For input s: %s, k: %d, expected %d, got %d", test.s, test.k, test.expected, result)
		}
	}
}
