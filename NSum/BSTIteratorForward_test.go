package nsum

import (
	"testing"
)

func TestBSTIteratorFWD(t *testing.T) {
	// Happy Path
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}

	iterator := BSTIteratorFWDConstructor(root, true)

	if !iterator.HasNext() {
		t.Errorf("Expected HasNext() to be true")
	}
	if val := iterator.Next(); val != 1 {
		t.Errorf("Expected Next() to return 1, got %d", val)
	}
	if val := iterator.Next(); val != 3 {
		t.Errorf("Expected Next() to return 3, got %d", val)
	}
	if val := iterator.Next(); val != 4 {
		t.Errorf("Expected Next() to return 4, got %d", val)
	}
	if iterator.HasNext() {
		t.Errorf("Expected HasNext() to be false")
	}

	// Edge Case: Empty Tree
	var emptyRoot *TreeNode
	emptyIterator := BSTIteratorFWDConstructor(emptyRoot, true)
	if emptyIterator.HasNext() {
		t.Errorf("Expected HasNext() to be false for empty tree")
	}

	// Edge Case: Single Node Tree
	oneNodeRoot := &TreeNode{Val: 5}
	oneNodeIterator := BSTIteratorFWDConstructor(oneNodeRoot, true)

	if !oneNodeIterator.HasNext() {
		t.Errorf("Expected HasNext() to be true for single node tree")
	}
	if val := oneNodeIterator.Next(); val != 5 {
		t.Errorf("Expected Next() to return 5, got %d", val)
	}
	if oneNodeIterator.HasNext() {
		t.Errorf("Expected HasNext() to be false after single node accessed")
	}
}
