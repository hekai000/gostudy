package nsum

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   [][]int
	}{
		{
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			want:   [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			want:   [][]int{{2, 2, 2, 2}},
		},
		{
			nums:   []int{},
			target: 0,
			want:   [][]int{},
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			target: 100,
			want:   [][]int{},
		},
		{
			nums:   []int{0, 0, 0, 0},
			target: 0,
			want:   [][]int{{0, 0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums: %v, target: %d", tt.nums, tt.target), func(t *testing.T) {
			got := fourSum(tt.nums, tt.target)
			sort.Sort(sortable(got))
			sort.Sort(sortable(tt.want))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

type sortable [][]int

func (s sortable) Len() int {
	return len(s)
}

func (s sortable) Less(i, j int) bool {
	if len(s[i]) == 0 || len(s[j]) == 0 {
		return len(s[i]) < len(s[j])
	}
	for k := 0; k < len(s[i]); k++ {
		if s[i][k] != s[j][k] {
			return s[i][k] < s[j][k]
		}
	}
	return false
}

func (s sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
