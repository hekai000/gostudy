package unionfind

import (
	"testing"
)

func TestUnionFindCount(t *testing.T) {
	tests := []struct {
		n                int
		unions           [][2]int
		expectedCount    []int
		expectedMaxCount int
	}{
		{
			n:                5,
			unions:           [][2]int{{0, 1}, {1, 2}},
			expectedCount:    []int{3, 1, 1, 1, 1},
			expectedMaxCount: 3,
		},
		{
			n:                3,
			unions:           [][2]int{{0, 1}, {0, 2}},
			expectedCount:    []int{3, 1, 1},
			expectedMaxCount: 3,
		},
		{
			n:                4,
			unions:           [][2]int{{0, 1}, {1, 2}, {2, 3}},
			expectedCount:    []int{4, 1, 1, 1},
			expectedMaxCount: 4,
		},
		{
			n:                0,
			unions:           [][2]int{},
			expectedCount:    []int{},
			expectedMaxCount: 0,
		},
		{
			n:                1,
			unions:           [][2]int{},
			expectedCount:    []int{1},
			expectedMaxCount: 0,
		},
	}

	for _, test := range tests {
		uf := &UnionFindCount{}
		uf.Init(test.n)

		for _, union := range test.unions {
			uf.Union(union[0], union[1])
		}

		if got := uf.Count(); !equal(got, test.expectedCount) {
			t.Errorf("Count() = %v, want %v", got, test.expectedCount)
		}

		if got := uf.MaxUnionCount(); got != test.expectedMaxCount {
			t.Errorf("MaxUnionCount() = %v, want %v", got, test.expectedMaxCount)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
