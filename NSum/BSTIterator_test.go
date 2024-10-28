package nsum

import (
	"testing"
)

func TestBSTIterator(t *testing.T) {
	var insertIntoBST func(*TreeNode, int) *TreeNode
	// Helper function to create a binary tree
	createTree := func(vals []int) *TreeNode {
		if len(vals) == 0 {
			return nil
		}
		root := &TreeNode{Val: vals[0]}
		for _, v := range vals[1:] {
			insertIntoBST(root, v)
		}
		return root
	}

	// Helper function to insert value into binary search tree
	insertIntoBST = func(node *TreeNode, val int) *TreeNode {
		if node == nil {
			return &TreeNode{Val: val}
		}
		if val < node.Val {
			node.Left = insertIntoBST(node.Left, val)
		} else {
			node.Right = insertIntoBST(node.Right, val)
		}
		return node
	}

	t.Run("Happy Path", func(t *testing.T) {
		root := createTree([]int{7, 3, 15, 9, 20})
		iter := BSTIteratorConstructor(root)

		expected := []int{3, 7, 9, 15, 20}
		for _, val := range expected {
			if !iter.HasNext() {
				t.Error("Expected HasNext to be true")
			}
			nextVal := iter.Next()
			if nextVal != val {
				t.Errorf("Expected next value to be %d, got %d", val, nextVal)
			}
		}
		if iter.HasNext() {
			t.Error("Expected HasNext to be false")
		}
	})

	t.Run("Empty Tree", func(t *testing.T) {
		var root *TreeNode
		iter := BSTIteratorConstructor(root)

		if iter.HasNext() {
			t.Error("Expected HasNext to be false on empty tree")
		}
	})

	t.Run("Single Node Tree", func(t *testing.T) {
		root := &TreeNode{Val: 5}
		iter := BSTIteratorConstructor(root)

		if !iter.HasNext() {
			t.Error("Expected HasNext to be true for single node tree")
		}
		if val := iter.Next(); val != 5 {
			t.Errorf("Expected next value to be 5, got %d", val)
		}
		if iter.HasNext() {
			t.Error("Expected HasNext to be false after accessing the single node")
		}
	})
}
