package unionfind

import "testing"

func TestUnionFind(t *testing.T) {
	uf := UnionFind{}
	uf.Init(10)

	// Happy Path
	if uf.TotalCount() != 10 {
		t.Errorf("Expected count to be 10, got %d", uf.TotalCount())
	}

	uf.Union(1, 2)
	if uf.Find(1) != uf.Find(2) {
		t.Errorf("Expected 1 and 2 to be connected")
	}

	uf.Union(2, 3)
	if uf.Find(1) != uf.Find(3) {
		t.Errorf("Expected 1 and 3 to be connected")
	}

	if uf.TotalCount() != 8 {
		t.Errorf("Expected count to be 8 after unions, got %d", uf.TotalCount())
	}

	// Edge cases
	uf.Union(1, 2) // Union repeat
	if uf.TotalCount() != 8 {
		t.Errorf("Expected count to be 8 after repeating union, got %d", uf.TotalCount())
	}

	uf.Union(3, 4)
	if uf.Find(2) != uf.Find(4) {
		t.Errorf("Expected 2 and 4 to not be connected")
	}

	uf.Union(5, 6)
	uf.Union(6, 7)
	uf.Union(7, 8)
	if uf.TotalCount() != 4 {
		t.Errorf("Expected count to be 5 after additional unions, got %d", uf.TotalCount())
	}

	// Test with invalid input
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on invalid union")
		}
	}()

	uf.Union(10, 11) // This should panic as the indices are out of bounds
}
