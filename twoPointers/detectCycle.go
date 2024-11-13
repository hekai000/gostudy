package twopointers

type ListNode struct {
	Val  int
	Next *ListNode
}

/***
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

不允许修改 链表。

解法：

我们可以用快慢指针的方法，慢指针每次移动一步，快指针每次移动两步。如果链表中存在环，则快指针会在环中相遇。

如果链表中不存在环，则快指针会在链表末尾相遇，此时返回 null。

如果链表中存在环，则我们可以让慢指针从头开始，快指针从相遇点开始，每次移动一步，直到快慢指针相遇。此时慢指针所指向的节点就是链表的入环点。

时间复杂度：O(n)，其中 n 是链表的长度。
***/

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
