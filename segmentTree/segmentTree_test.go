package segmentTree

import (
	"testing"
)

func TestSegmentTree(t *testing.T) {
	add := func(i, j int) int { return i + j }

	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "Happy Path - Non-empty",
			input:  []int{1, 3, 5, 7, 9, 11},
			expect: []int{36, 9, 27, 4, 5, 16, 11, 1, 3, 0, 0, 7, 9, 0, 0},
		},
		{
			name:   "Edge Case - Empty Array",
			input:  []int{},
			expect: []int{0},
		},
		{
			name:   "Edge Case - Single Element",
			input:  []int{42},
			expect: []int{42},
		},
		{
			name:   "Edge Case - Two Elements",
			input:  []int{1, 2},
			expect: []int{3, 1, 2, 0, 0, 0, 0},
		},
		{
			name:   "Edge Case - All Same Values",
			input:  []int{5, 5, 5, 5},
			expect: []int{20, 10, 10, 5, 5, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var st SegmentTree
			st.Init(tt.input, add)
			st.PrintSegmentTree()
			for i := range st.tree {
				if i < len(tt.expect) && st.tree[i] != tt.expect[i] {
					t.Errorf("unexpected tree value at index %d: got %d, want %d", i, st.tree[i], tt.expect[i])
				}
			}
		})
	}
}

func TestSegmentTree_Query(t *testing.T) {
	tests := []struct {
		name       string
		nums       []int
		queryLeft  int
		queryRight int
		expected   int
		oper       func(i, j int) int
	}{
		{
			name:       "Happy Path - Sum",
			nums:       []int{1, 2, 3, 4, 5},
			queryLeft:  1,
			queryRight: 3,
			expected:   9, // 2 + 3 + 4
			oper: func(i, j int) int {
				return i + j
			},
		},
		{
			name:       "Happy Path - Max",
			nums:       []int{1, 3, 2, 4, 5},
			queryLeft:  0,
			queryRight: 4,
			expected:   5, // max in full range
			oper: func(i, j int) int {
				if i > j {
					return i
				}
				return j
			},
		},
		{
			name:       "Edge Case - Empty Array",
			nums:       []int{},
			queryLeft:  0,
			queryRight: 0,
			expected:   0,
			oper: func(i, j int) int {
				return i + j
			},
		},
		{
			name:       "Edge Case - Single Element",
			nums:       []int{42},
			queryLeft:  0,
			queryRight: 0,
			expected:   42, // only one element in the range
			oper: func(i, j int) int {
				return i + j
			},
		},
		{
			name:       "Edge Case - Negative Numbers",
			nums:       []int{-1, -2, -3, -4, -5},
			queryLeft:  1,
			queryRight: 3,
			expected:   -9, // -2 + -3 + -4
			oper: func(i, j int) int {
				return i + j
			},
		},
		{
			name:       "Edge Case - Out of Bounds",
			nums:       []int{1, 2, 3},
			queryLeft:  -1,
			queryRight: 3,
			expected:   0, // should handle out of bounds gracefully
			oper: func(i, j int) int {
				return i + j
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var st SegmentTree
			st.Init(tt.nums, tt.oper)
			result := st.Query(tt.queryLeft, tt.queryRight)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
func TestSegmentTree_QueryLazy(t *testing.T) {
	// Test case with addition as the merge operation
	add := func(i, j int) int {
		return i + j
	}
	st := SegmentTree{}
	st.Init([]int{1, 2, 3, 4, 5}, add)

	// Happy path test
	result := st.QueryLazy(0, 4)
	if result != 15 {
		t.Errorf("Expected 15, got %d", result)
	}

	// Query a subarray
	result = st.QueryLazy(1, 3)
	if result != 9 {
		t.Errorf("Expected 9, got %d", result)
	}

	// Query with an empty range
	result = st.QueryLazy(2, 2)
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}

	// Edge case: invalid range (left > right)
	result = st.QueryLazy(4, 3)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}

	// Edge case: query out of bounds
	result = st.QueryLazy(-1, 6)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
func TestSegmentTree_Update(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		index    int
		value    int
		expected []int
	}{
		{
			name:     "Update middle element",
			nums:     []int{1, 3, 5, 7, 9, 11},
			index:    1,
			value:    10,
			expected: []int{43, 16, 27, 11, 5, 16, 11, 1, 10, 0, 0, 7, 9, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var st SegmentTree
			st.Init(tt.nums, func(i, j int) int {
				return i + j
			})

			st.Update(tt.index, tt.value)

			for i, val := range tt.expected {
				if st.tree[i] != val {
					t.Errorf("expected %d, got %d", val, st.tree[i])
				}
			}
		})
	}
}

func TestUpdateLazy(t *testing.T) {
	mergeFunc := func(i, j int) int {
		return i + j
	}
	var st SegmentTree
	nums := []int{1, 2, 3, 4, 5}
	st.Init(nums, mergeFunc)

	// Test case 1: Lazy update over the entire range
	st.UpdateLazy(0, 4, 10)
	if result := st.QueryLazy(0, 4); result != 65 {
		t.Errorf("Expected 65, got %d", result)
	}

	// Test case 2: Lazy update on sub-range
	st.UpdateLazy(1, 3, 5)
	if result := st.QueryLazy(0, 4); result != 80 {
		t.Errorf("Expected 80, got %d", result)
	}

	// Test case 3: Lazy update on a single element
	st.UpdateLazy(2, 2, 2)
	if result := st.QueryLazy(0, 4); result != 82 {
		t.Errorf("Expected 82, got %d", result)
	}

	// Test case 4: UpdateLazy outside the range
	st.UpdateLazy(5, 6, 3)
	if result := st.QueryLazy(0, 4); result != 82 {
		t.Errorf("Expected 82, got %d", result)
	}

	// Test case 5: Edge case with empty tree
	stEmpty := &SegmentTree{}
	stEmpty.UpdateLazy(0, 0, 5) // Should not panic or change anything
}

func TestLazyUpdateWithNoOverlap(t *testing.T) {
	mergeFunc := func(i, j int) int {
		return i + j
	}
	var st SegmentTree
	nums := []int{1, 2, 3, 4, 5}
	st.Init(nums, mergeFunc)

	// Initial update on one segment
	st.UpdateLazy(0, 1, 10)

	// Update with no overlap
	st.UpdateLazy(3, 4, 5)

	// Check the values to ensure they are correct
	if result := st.Query(0, 4); result != 45 {
		t.Errorf("Expected 45, got %d", result)
	}
}
