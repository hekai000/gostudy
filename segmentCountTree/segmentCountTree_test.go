package segmentCountTree

import (
	"testing"
)

func TestSegmentCountTree(t *testing.T) {
	mergeFunc := func(i, j int) int { return i + j }
	var st SegmentCountTree
	st.Init([]int{1, 3, 5, 7, 9, 11}, mergeFunc)
	st.buildSegmentCountTree(0, 0, len(st.data)-1)

	tests := []struct {
		queryLeft  int
		queryRight int
		expected   int
	}{
		{1, 5, 3},   // Query in the middle
		{3, 9, 3},   // Query including more elements
		{0, 11, 6},  // Query all elements
		{1, 1, 0},   // Query single element not present
		{12, 15, 0}, // Query out of range
	}

	for _, test := range tests {
		result := st.Query(test.queryLeft, test.queryRight)
		if result != test.expected {
			t.Errorf("Query(%d, %d) = %d; expected %d", test.queryLeft, test.queryRight, result, test.expected)
		}
	}

	st.Update(5) // Update an existing value
	result := st.Query(3, 7)
	if result != 4 {
		t.Errorf("Query(3, 7) after Update(5) = %d; expected 4", result)
	}

	st.Update(10) // Update a non-existing value
	result = st.Query(9, 11)
	if result != 3 {
		t.Errorf("Query(9, 11) after Update(10) = %d; expected 3", result)
	}
}
