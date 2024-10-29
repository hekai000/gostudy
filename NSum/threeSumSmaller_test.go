package nsum

import "testing"

func TestThreeSumSmaller(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		// Happy path test cases
		{[]int{-2, 0, 1, 3}, 2, 2},     // combinations: (-2, 0, 1), (-2, 0, 3)
		{[]int{0, 0, 3, 4}, 5, 2},      // combinations: (0, 0, 3)
		{[]int{-1, 1, 2, -2, 0}, 0, 4}, // combinations: (-2, -1, 0), (-2, -1, 1), (-2, -1, 2),(-2, 0, 1)

		// Edge cases
		{[]int{1, 1, 1}, 3, 0},       // no valid combinations
		{[]int{1, 2}, 3, 0},          // less than 3 elements
		{[]int{}, 0, 0},              // empty input
		{[]int{1, 1, 1, 1, 1}, 4, 6}, // all elements the same
		//i, j, k (0, 1, 2) (0,1,3) (0,1,4) (1, 2, 3) (1, 2, 4) (2, 3, 4)
		{[]int{-1, -1, -1, 1, 1}, 0, 5}, // combinations for negative and positive
		//i, j, k (0,1,2) (0,1,3) (0,1,4) (1,2,3) (1,2,4)

	}

	for _, tt := range tests {
		got := threeSumSmaller(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("threeSumSmaller(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}
