package nsum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},        // Happy path
		{[]int{3, 2, 4}, 6, []int{1, 2}},             // Happy path
		{[]int{3, 3}, 6, []int{1, 0}},                // Happy path
		{[]int{1, 5, 7, 3, 2, 4}, 10, []int{3, 2}},   // Edge case with larger array
		{[]int{0, 2, 4, 3, 3}, 6, []int{1, 2}},       // Edge case with duplicates
		{[]int{-1, -2, -3, -4, -5}, -8, []int{4, 2}}, // Edge case with negative numbers
		{[]int{}, 0, []int{}},                        // Edge case with empty array
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			if got := twoSum(test.nums, test.target); !reflect.DeepEqual(got, test.result) {
				t.Errorf("twoSum(%v, %d) = %v; want %v", test.nums, test.target, got, test.result)
			}
		})
	}
}
