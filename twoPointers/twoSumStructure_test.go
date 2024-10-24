package twopointers

import "testing"

func TestTwoSumStructure(t *testing.T) {
	ts := Constructor()

	// Happy path
	ts.Add(1)
	ts.Add(3)
	ts.Add(5)

	if !ts.Find(4) {
		t.Errorf("Expected to find 4, but got false")
	}

	if ts.Find(7) {
		t.Errorf("Expected not to find 7, but got true")
	}

	// Edge cases
	ts.Add(2)
	ts.Add(2)

	if !ts.Find(4) {
		t.Errorf("Expected to find 4, but got false")
	}

	if !ts.Find(5) {
		t.Errorf("Expected to find 5, but got false")
	}

	if !ts.Find(6) {
		t.Errorf("Expected not to find 6, but got true")
	}

	if ts.Find(0) {
		t.Errorf("Expected not to find 0, but got true")
	}

	// Testing with duplicate values
	ts.Add(3)

	if !ts.Find(6) {
		t.Errorf("Expected to find 6, but got false")
	}
}
