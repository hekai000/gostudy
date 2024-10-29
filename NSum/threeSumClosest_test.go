package nsum

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},    // Happy path
		{[]int{0, 0, 0}, 1, 0},         // Happy path with zeros
		{[]int{1, 1, 1, 1}, 2, 3},      // All elements are the same
		{[]int{1, 2, 3, 4, 5}, 10, 10}, // Multiple combinations
		{[]int{5, 5, 5, 5}, 15, 15},    // Edge case: target equals sum of all elements
		// {[]int{}, 0, 0},                  // Edge case: empty input
		// {[]int{1}, 1, 0},                 // Edge case: single element
		{[]int{-1, -1, -1, -1}, -3, -3},  // Edge case: all negative elements
		{[]int{100, 200, 300}, 500, 600}, // Edge case: large numbers
	}

	for _, tt := range tests {
		got := threeSumClosest(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("threeSumClosest(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}
