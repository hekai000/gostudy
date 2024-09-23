package lrucache

import (
	"testing"
)

func TestNode(t *testing.T) {
	// Happy path
	node1 := Node{Key: 1, Value: 100}
	if node1.Key != 1 || node1.Value != 100 {
		t.Errorf("Expected Key: 1, Value: 100, Got Key: %d, Value: %d", node1.Key, node1.Value)
	}

	node2 := Node{Key: 2, Value: 200}
	node1.Next = &node2
	node2.Prev = &node1

	// Check Next and Prev pointers
	if node1.Next != &node2 {
		t.Errorf("Expected node1 Next to be node2")
	}
	if node2.Prev != &node1 {
		t.Errorf("Expected node2 Prev to be node1")
	}

	// Edge Cases
	node3 := Node{Key: 3, Value: 300}
	node2.Next = &node3
	node3.Prev = &node2

	if node2.Next != &node3 {
		t.Errorf("Expected node2 Next to be node3")
	}
	if node3.Prev != &node2 {
		t.Errorf("Expected node3 Prev to be node2")
	}

	// Test uninitialized Node
	var node4 Node
	if node4.Key != 0 || node4.Value != 0 {
		t.Errorf("Expected uninitialized Node Key: 0, Value: 0, Got Key: %d, Value: %d", node4.Key, node4.Value)
	}
}
