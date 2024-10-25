package twopointers

import (
	"fmt"
	"testing"
)

func TestFindTarget2(t *testing.T) {
	type testCase struct {
		root *TreeNode
		k    int
		want bool
	}

	tests := []testCase{
		// Happy path
		{newTree([]int{1, 2, 3}), 5, true},
		{newTree([]int{5, 3, 6, 2, 4, -200, 7}), 9, true},   // 4 + 5
		{newTree([]int{5, 3, 6, 2, 4, -200, 7}), 28, false}, // No pairs

		// Edge cases
		{nil, 9, false},                     // Empty tree
		{newTree([]int{2}), 3, false},       // Single node, no pair
		{newTree([]int{2, 1, 3}), 4, true},  // Multiple nodes with one valid pair
		{newTree([]int{5, 2, 8}), 10, true}, // Valid pair; 2 + 8
		{newTree([]int{5, 2, 8}), 5, false}, // Only one node value
		{newTree([]int{1, 2, 3}), 5, true},  // Valid pair; 2 + 3
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("findTarget2(%v, %v)", tt.root, tt.k), func(t *testing.T) {
			got := findTarget2(tt.root, tt.k)
			if got != tt.want {
				t.Errorf("findTarget2() = %v; want %v", got, tt.want)
			}
		})
	}
}

func newTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	root := &TreeNode{Val: values[0]}
	for _, v := range values[1:] {
		insert(root, v)
	}
	return root
}

func insert(node *TreeNode, val int) {
	if val == -200 {
		return
	}
	if val < node.Val {
		if node.Left == nil {
			node.Left = &TreeNode{Val: val}
		} else {
			insert(node.Left, val)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Val: val}
		} else {
			insert(node.Right, val)
		}
	}
}
