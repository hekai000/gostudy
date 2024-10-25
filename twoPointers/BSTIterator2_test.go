package twopointers

import (
	"testing"
)

func TestBSTIterator2(t *testing.T) {
	// Helper function to create a new TreeNode
	newNode := func(val int) *TreeNode {
		return &TreeNode{Val: val, Left: nil, Right: nil}
	}

	// Test case for a basic tree
	root := newNode(7)
	root.Left = newNode(3)
	root.Right = newNode(15)
	root.Right.Left = newNode(9)
	root.Right.Right = newNode(20)

	iterator := BSTIteratorConstructor2(root)

	// Happy Path
	if !iterator.HasNext() {
		t.Fatalf("Expected HasNext() to return true")
	}
	if val := iterator.Next(); val != 3 {
		t.Fatalf("Expected Next() to return 3, got %d", val)
	}
	if val := iterator.Next(); val != 7 {
		t.Fatalf("Expected Next() to return 7, got %d", val)
	}
	if val := iterator.Next(); val != 9 {
		t.Fatalf("Expected Next() to return 9, got %d", val)
	}
	if val := iterator.Next(); val != 15 {
		t.Fatalf("Expected Next() to return 15, got %d", val)
	}
	if val := iterator.Next(); val != 20 {
		t.Fatalf("Expected Next() to return 20, got %d", val)
	}
	if iterator.HasNext() {
		t.Fatalf("Expected HasNext() to return false")
	}

	// Edge case: Empty tree
	emptyIterator := BSTIteratorConstructor2(nil)
	if emptyIterator.HasNext() {
		t.Fatalf("Expected HasNext() to return false for an empty tree")
	}
}
