package nsum

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	bstMap := make(map[int]int)
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if _, ok := bstMap[k-node.Val]; ok {
			return true
		} else {
			bstMap[node.Val] += 1
		}
		if dfs(node.Left) || dfs(node.Right) {
			return true
		}
		return false
	}
	return dfs(root)

}
