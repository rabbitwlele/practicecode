package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := root.Val
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		ll := dfs(root.Left)
		rr := dfs(root.Right)
		depth := max(root.Val, root.Val+max(ll, rr))
		ans = max(ans, max(ll+rr+root.Val, depth))
		return depth
	}
	dfs(root)
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
