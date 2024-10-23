package binarysearch

import "testing"

func TestMinDays(t *testing.T) {
	tests := []struct {
		bloomDay []int
		m        int
		k        int
		expected int
	}{
		// Happy path
		{[]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 4, 2, 9},
		{[]int{1, 10, 3, 10, 2}, 3, 1, 3}, // Minimum days to bloom 3 flowers each requiring 1 day
		{[]int{1, 2, 3, 4, 5}, 1, 2, 2},   // Minimum days to bloom 1 flower with 2 length
		{[]int{7, 7, 7, 7, 7}, 1, 1, 7},   // All flowers bloom the same day

		// Edge cases
		{[]int{1, 2, 3}, 5, 1, -1},         // Not enough flowers
		{[]int{1, 2, 3}, 1, 5, -1},         // Not enough flowers in a group
		{[]int{1, 2, 3, 4, 5}, 0, 2, 0},    // If m is 0, should return 0
		{[]int{}, 1, 1, -1},                // Empty bloomDay
		{[]int{1}, 1, 1, 1},                // Single flower
		{[]int{1, 10, 3, 10, 2}, 2, 2, 10}, // Minimum days for 2 flowers each requiring 2 days
	}

	for _, test := range tests {
		result := minDays(test.bloomDay, test.m, test.k)
		if result != test.expected {
			t.Errorf("For bloomDay %v, m = %d, k = %d; expected %d, got %d", test.bloomDay, test.m, test.k, test.expected, result)
		}
	}
}
