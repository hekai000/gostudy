package nsum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{
			input:    []int{7, -1, -1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}},
		},
		{
			input:    []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, 0, 1}, {-1, -1, 2}},
		},
		{
			input:    []int{},
			expected: [][]int{},
		},
		{
			input:    []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			input:    []int{1, 2, -2, -1},
			expected: [][]int{},
		},
		{
			input:    []int{-1, 0, 1, 1, 2},
			expected: [][]int{{-1, 0, 1}},
		},
	}

	for _, test := range tests {
		result := threeSum(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
		}
	}
}
