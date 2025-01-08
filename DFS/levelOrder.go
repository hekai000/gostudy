/*
 * [102] 二叉树的层序遍历
 */
package dfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans = [][]int{}

func levelOrder(root *TreeNode) [][]int {

	ans = [][]int{}
	dfs(root, 0)
	return ans

}

func dfs(root *TreeNode, h int) {
	if root == nil {
		return
	}
	if root != nil {
		if len(ans) <= h {
			ans = append(ans, []int{})

		}
		ans[h] = append(ans[h], root.Val)
	}
	if root.Left != nil {
		dfs(root.Left, h+1)
	}
	if root.Right != nil {
		dfs(root.Right, h+1)
	}
}
