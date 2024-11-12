package nsum

import (
	"fmt"
	"testing"
)

func TestKSum2(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		k      int
		want   int
	}{
		// Happy path test cases

		{[]int{1, 2, 3, 4, 5}, 5, 2, 2}, // 1+4,2+3
		{[]int{1, 2, 3, 4, 5}, 5, 1, 1}, // 5
		{[]int{1, 1, 1, 1}, 2, 2, 6},    // 1+1
		{[]int{1, 2, 3, 4, 5}, 9, 3, 2}, // 2+3+4
		{[]int{2, 4, 6, 8}, 12, 2, 1},   // 4+8

		// Edge cases
		{[]int{}, 0, 0, 1},        // Empty input
		{[]int{1}, 1, 1, 1},       // Single element matches
		{[]int{1}, 1, 2, 0},       // Single element, k > 1
		{[]int{1, 2, 3}, 7, 3, 0}, // No combination to reach target

	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("kSum2(%v, %d, %d)", tt.nums, tt.target, tt.k), func(t *testing.T) {
			got := kSum2(tt.nums, tt.target, tt.k)
			if got != tt.want {
				t.Errorf("kSum2(%v, %d, %d) = %d; want %d", tt.nums, tt.target, tt.k, got, tt.want)
			}
		})
	}
}
