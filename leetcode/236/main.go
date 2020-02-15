package _36

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		tmp := dfs(root.Left) | dfs(root.Right)
		if root == p {
			tmp |= 1
		}
		if root == q {
			tmp |= 2
		}
		if ans == nil && tmp == 3 {
			ans = root
		}
		return tmp
	}
	dfs(root)
	return ans
}
