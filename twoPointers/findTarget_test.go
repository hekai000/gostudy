package twopointers

import (
	"testing"
)

func TestFindTarget(t *testing.T) {

	tests := []struct {
		name     string
		root     *TreeNode
		k        int
		expected bool
	}{
		{
			name:     "Happy Path: one",
			root:     &TreeNode{Val: 1, Left: nil, Right: nil},
			k:        2,
			expected: false,
		},
		{
			name:     "Happy Path: Target exists",
			root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}},
			k:        9,
			expected: true,
		},
		{
			name:     "Happy Path: Target exists with negative values",
			root:     &TreeNode{Val: -3, Left: &TreeNode{Val: -5}, Right: &TreeNode{Val: -2}},
			k:        -7,
			expected: true,
		},
		{
			name:     "Edge Case: Empty tree",
			root:     nil,
			k:        5,
			expected: false,
		},
		{
			name:     "Edge Case: Single node with value equals to k",
			root:     &TreeNode{Val: 5},
			k:        5,
			expected: false,
		},
		{
			name:     "Happy Path: Multiple nodes no two sum",
			root:     &TreeNode{Val: 10, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 15}},
			k:        20,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := findTarget(tt.root, tt.k)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
