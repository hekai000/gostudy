package twopointers

/**
 * Definition for a binary tree node.
 */

type BSTIterator struct {
	arr []int
}

func BSTIteratorConstructor(root *TreeNode) BSTIterator {
	it := BSTIterator{}
	it.inOrderTraversal(root)
	return it
}

func (this *BSTIterator) inOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	this.inOrderTraversal(root.Left)
	this.arr = append(this.arr, root.Val)
	this.inOrderTraversal(root.Right)
}

func (this *BSTIterator) Next() int {
	val := this.arr[0]
	this.arr = this.arr[1:]
	return val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.arr) > 0
}
