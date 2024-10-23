package binarysearch

import "testing"

func TestMaxDistance(t *testing.T) {
	tests := []struct {
		position []int
		m        int
		expected int
	}{
		// Happy path
		{[]int{1, 2, 3, 4, 7}, 3, 3},
		{[]int{1, 3, 6, 7, 10}, 3, 4},

		// Edge cases
		{[]int{1, 2}, 2, 1},                  // Minimum distance for 2 positions
		{[]int{1, 1000000000}, 2, 999999999}, // Maximum distance in large numbers
		{[]int{1}, 1, 0},                     // Only one position
		// {[]int{}, 1, 0},                      // No positions
		{[]int{1, 2, 3, 4, 5}, 0, 0}, // m = 0
	}

	for _, test := range tests {
		result := maxDistance(test.position, test.m)
		if result != test.expected {
			t.Errorf("For position %v and m = %d, expected %d but got %d", test.position, test.m, test.expected, result)
		}
	}
}
