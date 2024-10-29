package nsum

import "testing"

func TestThreeSumMulti(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		result int
	}{
		// Happy path tests
		{arr: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, target: 8, result: 20},
		{arr: []int{0, 0, 0, 0}, target: 0, result: 4},
		{arr: []int{1, 2, 3, 4, 5}, target: 6, result: 1},

		// Edge case tests
		{arr: []int{}, target: 0, result: 0},                    // empty array
		{arr: []int{1}, target: 1, result: 0},                   // single element
		{arr: []int{1, 1, 1}, target: 3, result: 1},             // all elements the same
		{arr: []int{3, 3, 3, 3, 3}, target: 9, result: 10},      // all elements the same forming multiple combinations
		{arr: []int{-1, 0, 1, 2, -1, -4}, target: 0, result: 3}, // test with negative numbers
		//(0,1,2) (0,3,4) (1, 2,4)
	}

	for _, test := range tests {
		t.Run("Testing with array", func(t *testing.T) {
			if got := threeSumMulti(test.arr, test.target); got != test.result {
				t.Errorf("threeSumMulti(%v, %d) = %d; want %d", test.arr, test.target, got, test.result)
			}
		})
	}
}
