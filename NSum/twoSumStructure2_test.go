package nsum

import (
	"testing"
)

func TestTwoSum2(t *testing.T) {
	ts := TwoSum2Constructor()

	// Happy Path
	ts.AddTwoSum2(1)
	ts.AddTwoSum2(2)
	ts.AddTwoSum2(3)

	if !ts.FindTwoSum2(3) {
		t.Errorf("Expected true for target 3")
	}
	if !ts.FindTwoSum2(4) {
		t.Errorf("Expected true for target 4")
	}
	if !ts.FindTwoSum2(5) {
		t.Errorf("Expected true for target 5")
	}

	// Edge Cases
	ts.AddTwoSum2(0)
	if !ts.FindTwoSum2(1) {
		t.Errorf("Expected true for target 1 with numbers 1 and 0")
	}
	if !ts.FindTwoSum2(2) {
		t.Errorf("Expected true for target 2 with numbers 2 and 0 added")
	}

	ts = TwoSum2Constructor() // reset
	ts.AddTwoSum2(-1)
	ts.AddTwoSum2(-2)
	ts.AddTwoSum2(-3)

	if !ts.FindTwoSum2(-3) {
		t.Errorf("Expected true for target -3")
	}
	if ts.FindTwoSum2(-2) {
		t.Errorf("Expected false for target -2")
	}
	if ts.FindTwoSum2(-1) {
		t.Errorf("Expected false for target -1 with only -1 added")
	}

	// Test with duplicates
	ts.AddTwoSum2(1)
	ts.AddTwoSum2(1)

	if !ts.FindTwoSum2(2) {
		t.Errorf("Expected true for target 2 with duplicates")
	}
}
