package twopointers

/**
 * Definition for a binary tree node.
 */

type BSTIteratorFWD struct {
	arr []int
}

func BSTIteratorFWDConstructor(root *TreeNode, forward bool) BSTIteratorFWD {
	it := BSTIteratorFWD{}
	it.inOrderTraversal(root, forward)
	return it
}

func (this *BSTIteratorFWD) inOrderTraversal(root *TreeNode, forward bool) {
	if root == nil {
		return
	}
	if forward {
		this.inOrderTraversal(root.Left, forward)
		this.arr = append(this.arr, root.Val)
		this.inOrderTraversal(root.Right, forward)
	} else {
		this.inOrderTraversal(root.Right, forward)
		this.arr = append(this.arr, root.Val)
		this.inOrderTraversal(root.Left, forward)
	}
}

func (this *BSTIteratorFWD) Next() int {
	val := this.arr[0]
	this.arr = this.arr[1:]
	return val
}

func (this *BSTIteratorFWD) HasNext() bool {
	return len(this.arr) > 0
}

func (this *BSTIteratorFWD) Peek() int {
	val := this.arr[0]
	return val
}
