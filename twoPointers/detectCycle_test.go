package twopointers

import (
	"testing"
)

func TestDetectCycle(t *testing.T) {
	// Happy path
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node2 // creates a cycle

	cycleNode := detectCycle(node1)
	if cycleNode != node2 {
		t.Errorf("Expected cycle node to be %v, got %v", node2.Val, cycleNode.Val)
	}

	// No cycle
	node1 = &ListNode{Val: 1}
	node2 = &ListNode{Val: 2}
	node1.Next = node2

	cycleNode = detectCycle(node1)
	if cycleNode != nil {
		t.Errorf("Expected no cycle, got node %v", cycleNode.Val)
	}

	// Single node without cycle
	node1 = &ListNode{Val: 1}
	cycleNode = detectCycle(node1)
	if cycleNode != nil {
		t.Errorf("Expected no cycle, got node %v", cycleNode)
	}

	// Single node with cycle
	node1.Next = node1 // creates a cycle
	cycleNode = detectCycle(node1)
	if cycleNode != node1 {
		t.Errorf("Expected cycle node to be %v, got %v", node1.Val, cycleNode.Val)
	}
}
