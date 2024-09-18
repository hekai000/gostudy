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
			expect: []int{36, 10, 28, 4, 16, 18, 0},
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