package nsum

import (
	"reflect"
	"testing"
)

func TestKSum(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		target int
		start  int
		result [][]int
	}{
		// Happy path
		{[]int{1, 2, 3, 4, 5}, 2, 5, 0, [][]int{{1, 4}, {2, 3}}},
		{[]int{1, 0, -1, 0, -2, 2}, 4, 0, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{1, 2, 1, 0, -1}, 3, 2, 0, [][]int{{-1, 1, 2}, {0, 1, 1}}},

		// Edge cases
		{[]int{}, 2, 0, 0, [][]int{}},                 // empty array
		{[]int{1}, 2, 2, 0, [][]int{}},                // single element
		{[]int{2, 2, 2, 2}, 2, 4, 0, [][]int{{2, 2}}}, // duplicates with valid pairs
		{[]int{1, 2, 3}, 3, 6, 0, [][]int{{1, 2, 3}}}, // k equals length of nums
		{[]int{-1, 0, 1}, 2, 0, 0, [][]int{{-1, 1}}},  // negative and positive numbers
	}

	for _, test := range tests {
		got := kSum(test.nums, test.k, test.target, test.start)
		if !reflect.DeepEqual(got, test.result) {
			t.Errorf("kSum(%v, %d, %d, %d) = %v; want %v", test.nums, test.k, test.target, test.start, got, test.result)
		}
	}
}
