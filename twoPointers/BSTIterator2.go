package twopointers

/**
 * Definition for a binary tree node.
 */

type BSTIterator2 struct {
	stack []*TreeNode
	cur   *TreeNode
}

func BSTIteratorConstructor2(root *TreeNode) BSTIterator2 {
	return BSTIterator2{stack: []*TreeNode{}, cur: root}
}

func (this *BSTIterator2) Next() int {
	for node := this.cur; node != nil; node = node.Left {
		this.stack = append(this.stack, node)
	}
	this.cur = this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	val := this.cur.Val
	this.cur = this.cur.Right
	return val
}

func (this *BSTIterator2) HasNext() bool {
	return this.cur != nil || len(this.stack) > 0
}
