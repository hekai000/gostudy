package twopointers

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		// Happy path
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		// Edge cases
		{[]int{1, 2, 1}, 2},
		{[]int{1, 0, 0, 0, 1}, 4},
		{[]int{}, 0},
		{[]int{0, 0, 0, 0}, 0},
		{[]int{5}, 0},
	}

	for _, tt := range tests {
		got := maxArea(tt.height)
		if got != tt.want {
			t.Errorf("maxArea(%v) = %v; want %v", tt.height, got, tt.want)
		}
	}
}
