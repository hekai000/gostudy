package leetcode100

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			l1:   &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			l2:   &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			want: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			l1:   &ListNode{Val: 0},
			l2:   &ListNode{Val: 0},
			want: &ListNode{Val: 0},
		},
		{
			l1:   &ListNode{Val: 9, Next: &ListNode{Val: 9}},
			l2:   &ListNode{Val: 1},
			want: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
		{
			l1:   &ListNode{Val: 1},
			l2:   &ListNode{Val: 9, Next: &ListNode{Val: 9}},
			want: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
		{
			l1:   &ListNode{Val: 5},
			l2:   nil,
			want: &ListNode{Val: 5},
		},
		{
			l1:   nil,
			l2:   &ListNode{Val: 5},
			want: &ListNode{Val: 5},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := addTwoNumbers(tt.l1, tt.l2); !equal(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
