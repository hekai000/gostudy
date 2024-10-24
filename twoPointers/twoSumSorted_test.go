package twopointers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSumSorted(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{[]int{1, 2, 3, 4, 6}, 6, []int{2, 4}},  // happy path
		{[]int{1, 2, 3, 4, 6}, 10, []int{4, 5}}, // happy path
		{[]int{1, 2, 3, 4, 6}, 5, []int{1, 4}},  // happy path
		{[]int{1, 2, 3, 4, 6}, 2, []int{}},      // edge case: smallest target
		{[]int{1, 2}, 3, []int{1, 2}},           // edge case: target equal to sum of two elements
		{[]int{1, 2}, 5, []int{}},               // edge case: target greater than sum
		{[]int{1, 2}, 0, []int{}},               // edge case: target less than smallest element
		{[]int{}, 1, []int{}},                   // edge case: empty array
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("numbers=%v,target=%d", tt.numbers, tt.target), func(t *testing.T) {
			got := twoSumSorted(tt.numbers, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
