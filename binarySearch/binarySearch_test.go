package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "Happy path - target is in the middle",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "Happy path - target is the first element",
			arr:      []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "Happy path - target is the last element",
			arr:      []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},
		{
			name:     "Edge case - empty array",
			arr:      []int{},
			target:   1,
			expected: -1,
		},
		{
			name:     "Edge case - target not in array",
			arr:      []int{1, 2, 3, 4, 5},
			target:   6,
			expected: -1,
		},
		{
			name:     "Edge case - target less than min",
			arr:      []int{1, 2, 3, 4, 5},
			target:   0,
			expected: -1,
		},
		{
			name:     "Edge case - target greater than max",
			arr:      []int{1, 2, 3, 4, 5},
			target:   10,
			expected: -1,
		},
		{
			name:     "Edge case - single element array, target present",
			arr:      []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "Edge case - single element array, target absent",
			arr:      []int{1},
			target:   2,
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := binarysearch(tc.arr, tc.target)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}
