package twopointers

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var bstMap = make(map[int]int)

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	cur := root
	if _, ok := bstMap[k-cur.Val]; ok {
		return true
	} else {
		bstMap[cur.Val] += 1
	}
	if findTarget(cur.Left, k) || findTarget(cur.Right, k) {
		return true
	}
	return false

}
