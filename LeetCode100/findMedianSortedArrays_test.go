package leetcode100

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		// Happy path
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{}, []int{1}, 1.0},
		{[]int{2}, []int{}, 2.0},
		{[]int{100001}, []int{100000}, 100000.5},

		// Edge cases
		{[]int{}, []int{}, 0.0},                        // edge case: both arrays empty
		{[]int{1}, []int{2}, 1.5},                      // edge case: one element in each array
		{[]int{1, 2, 3}, []int{4, 5}, 3.0},             // edge case: different lengths
		{[]int{1, 3, 8}, []int{7, 9, 10, 11}, 8.0},     // edge case: various sizes with overlap
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9}, 5.0}, // edge case: longer arrays
	}

	for _, tt := range tests {
		t.Run("Testing findMedianSortedArrays", func(t *testing.T) {
			got := findMedianSortedArrays(tt.nums1, tt.nums2)
			if got != tt.want {
				t.Errorf("findMedianSortedArrays(%v, %v) = %v; want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
