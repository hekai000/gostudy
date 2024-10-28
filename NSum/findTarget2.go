package nsum

func findTarget2(root *TreeNode, k int) bool {
	left := BSTIteratorFWDConstructor(root, true)
	right := BSTIteratorFWDConstructor(root, false)
	for left.HasNext() && right.HasNext() {
		l, r := left.Peek(), right.Peek()
		if l >= r {
			return false
		}
		if l+r == k {
			return true
		} else if l+r < k {
			left.Next()
		} else {
			right.Next()
		}
	}

	return false
}
