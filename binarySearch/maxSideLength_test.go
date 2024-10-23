package binarysearch

import "testing"

func TestMaxSideLength(t *testing.T) {
	tests := []struct {
		mat       [][]int
		threshold int
		expected  int
	}{
		{
			mat: [][]int{
				{18, 70},
				{61, 1},
				{25, 85},
				{14, 40},
				{11, 96},
				{97, 96},
				{63, 45},
			},
			threshold: 40184,
			expected:  2,
		},
		{
			mat: [][]int{
				{1},
			},
			threshold: 1,
			expected:  1,
		},
		{
			mat: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			threshold: 20,
			expected:  2,
		},
		{
			mat: [][]int{
				{7, 6, 8},
				{1, 4, 2},
				{2, 3, 5},
			},
			threshold: 15,
			expected:  2,
		},
		{
			mat: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			threshold: 5,
			expected:  2,
		},
		{
			mat: [][]int{
				{1, 2},
				{3, 4},
			},
			threshold: 10,
			expected:  2,
		},
		{
			mat: [][]int{
				{1},
			},
			threshold: 1,
			expected:  1,
		},
		{
			mat: [][]int{
				{1},
			},
			threshold: 0,
			expected:  0,
		},
		{
			mat: [][]int{
				{1, 2},
				{3, 4},
			},
			threshold: 5,
			expected:  1,
		},
	}

	for _, test := range tests {
		result := maxSideLength(test.mat, test.threshold)
		if result != test.expected {
			t.Errorf("For matrix: %v and threshold: %d, expected %d but got %d", test.mat, test.threshold, test.expected, result)
		}
	}
}
