package binarysearch

import "testing"

func TestSplitArray(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		// Happy path
		{[]int{7, 2, 5, 10, 8}, 2, 18},
		{[]int{1, 2, 3, 4, 5}, 2, 9},
		{[]int{1, 4, 4}, 3, 4},

		// Edge cases
		{[]int{1}, 1, 1},                  // Single element array
		{[]int{1, 2}, 1, 3},               // Two elements, k=1
		{[]int{1, 2, 3, 4, 5, 6}, 6, 6},   // k equal to length of nums
		{[]int{10, 10, 10}, 3, 10},        // All elements the same
		{[]int{5, 6, 7, 8, 9, 10}, 1, 45}, // Large k
		{[]int{100, 200, 300}, 2, 300},    // Large values
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := splitArray(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("splitArray(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
