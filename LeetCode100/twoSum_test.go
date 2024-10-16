package leetcode100

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "happy path",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "negative numbers",
			nums:   []int{-3, 4, 3, 90},
			target: 0,
			want:   []int{0, 2},
		},
		{
			name:   "target is the sum of the same number",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "no solution",
			nums:   []int{1, 2, 3},
			target: 7,
			want:   []int{},
		},
		{
			name:   "empty array",
			nums:   []int{},
			target: 1,
			want:   []int{},
		},
		{
			name:   "one element (no solution)",
			nums:   []int{5},
			target: 5,
			want:   []int{},
		},
		{
			name:   "two elements (solution)",
			nums:   []int{1, 4},
			target: 5,
			want:   []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
