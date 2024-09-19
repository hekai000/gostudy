package main

import (
	"testing"
)

func TestSegmentTree(t *testing.T) {
	add := func(i, j int) int { return i + j }

	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "Happy Path - Non-empty",
			input:  []int{1, 3, 5, 7, 9, 11},
			expect: []int{36, 9, 27, 4, 5, 16, 11, 1, 3, 0, 0, 7, 9, 0, 0},
		},
		{
			name:   "Edge Case - Empty Array",
			input:  []int{},
			expect: []int{0},
		},
		{
			name:   "Edge Case - Single Element",
			input:  []int{42},
			expect: []int{42},
		},
		{
			name:   "Edge Case - Two Elements",
			input:  []int{1, 2},
			expect: []int{3, 1, 2, 0, 0, 0, 0},
		},
		{
			name:   "Edge Case - All Same Values",
			input:  []int{5, 5, 5, 5},
			expect: []int{20, 10, 10, 5, 5, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var st SegmentTree
			st.Init(tt.input, add)

			for i := range st.tree {
				if i < len(tt.expect) && st.tree[i] != tt.expect[i] {
					t.Errorf("unexpected tree value at index %d: got %d, want %d", i, st.tree[i], tt.expect[i])
				}
			}
		})
	}
}

func TestSegmentTree_Query(t *testing.T) {
	tests := []struct {
		name       string
		nums       []int
		queryLeft  int
		queryRight int
		expected   int
		oper       func(i, j int) int
	}{
		{
			name:       "Happy Path - Sum",
			nums:       []int{1, 2, 3, 4, 5},
			queryLeft:  1,
			queryRight: 3,
			expected:   9, // 2 + 3 + 4
			oper: func(i, j int) int {
				return i + j
			},
		},
		{
			name:       "Happy Path - Max",
			nums:       []int{1, 3, 2, 4, 5},
			queryLeft:  0,
			queryRight: 4,
			expected:   5, // max in full range
			oper: func(i, j int) int {
				if i > j {
					return i
				}
				return j
			},
		},
		{
			name:       "Edge Case - Empty Array",
			nums:       []int{},
			queryLeft:  0,
			queryRight: 0,
			expected:   0,
			oper: func(i, j int) int {
				return i + j
			},
		},
		{
			name:       "Edge Case - Single Element",
			nums:       []int{42},
			queryLeft:  0,
			queryRight: 0,
			expected:   42, // only one element in the range
			oper: func(i, j int) int {
				return i + j
			},
		},
		{
			name:       "Edge Case - Negative Numbers",
			nums:       []int{-1, -2, -3, -4, -5},
			queryLeft:  1,
			queryRight: 3,
			expected:   -9, // -2 + -3 + -4
			oper: func(i, j int) int {
				return i + j
			},
		},
		{
			name:       "Edge Case - Out of Bounds",
			nums:       []int{1, 2, 3},
			queryLeft:  -1,
			queryRight: 3,
			expected:   0, // should handle out of bounds gracefully
			oper: func(i, j int) int {
				return i + j
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var st SegmentTree
			st.Init(tt.nums, tt.oper)
			result := st.Query(tt.queryLeft, tt.queryRight)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
