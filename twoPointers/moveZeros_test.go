package twopointers

import (
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{0, 1, 0, 3, 12}, expected: []int{1, 3, 12, 0, 0}}, // happy path
		{input: []int{0, 0, 0, 0}, expected: []int{0, 0, 0, 0}},         // all zeros
		{input: []int{1, 2, 3, 4}, expected: []int{1, 2, 3, 4}},         // no zeros
		{input: []int{}, expected: []int{}},                             // empty array
		{input: []int{1}, expected: []int{1}},                           // single non-zero
		{input: []int{0}, expected: []int{0}},                           // single zero
		{input: []int{1, 0, 2, 0, 3}, expected: []int{1, 2, 3, 0, 0}},   // mixed
	}

	for _, test := range tests {
		moveZeroes(test.input)
		for i, v := range test.expected {
			if test.input[i] != v {
				t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, test.input)
			}
		}
	}
}
