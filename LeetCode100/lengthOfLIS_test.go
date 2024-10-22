package leetcode100

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4}, // Happy path
		{[]int{1, 3, 6, 7, 9, 4, 2, 3}, 5},     // Happy path
		{[]int{}, 0},                           // Edge case: empty array
		{[]int{7}, 1},                          // Edge case: single element
		{[]int{10, 9, 8, 7, 6, 5}, 1},          // Edge case: decreasing sequence
		{[]int{1, 2, 3, 4, 5}, 5},              // Edge case: increasing sequence
		{[]int{1, 2, 1, 3, 1, 4}, 4},           // Edge case: duplicates with increasing subsequence
		{[]int{5, 6, 1, 2, 3, 4}, 4},           // Edge case: mix with increasing subsequence
	}

	for _, test := range tests {
		result := lengthOfLIS(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %d but got %d", test.input, test.expected, result)
		}
	}
}
