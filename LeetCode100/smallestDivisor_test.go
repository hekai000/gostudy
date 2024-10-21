package leetcode100

import (
	"testing"
)

func TestSmallestDivisor(t *testing.T) {
	tests := []struct {
		nums      []int
		threshold int
		expected  int
	}{
		// Happy path
		{[]int{21212, 10101, 12121}, 1000000, 1},
		{[]int{19}, 5, 4},
		{[]int{1, 2, 5, 9}, 6, 5},
		{[]int{2, 3, 5, 7, 11}, 11, 3},

		// Edge cases
		{[]int{1}, 1, 1},       // Single element, divisor equals threshold
		{[]int{1, 2, 3}, 3, 3}, // All elements must be divided
		// {[]int{}, 10, 0},                 // Empty array
		{[]int{10, 20, 30, 40}, 4, 40},  // High divisor for a low threshold
		{[]int{10, 15, 25, 30}, 100, 1}, // Low divisor for a high threshold
		{[]int{1, 1000000}, 500, 2005},  // High range of values
	}

	for _, test := range tests {
		result := smallestDivisor(test.nums, test.threshold)
		if result != test.expected {
			t.Errorf("For nums: %v and threshold: %d, expected %d but got %d", test.nums, test.threshold, test.expected, result)
		}
	}
}
