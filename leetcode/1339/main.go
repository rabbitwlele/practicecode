package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	ans := 0
	mod := 1000000000 + 7
	var dfs func(root *TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		return dfs(r.Left) + dfs(r.Right) + r.Val
	}
	sum := dfs(root)
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		ll := dfs(r.Left)
		rr := dfs(r.Right)
		ans = min(ans, min(abs(sum-2*ll), abs(sum-2*rr)))
		return ll + rr + r.Val
	}
	dfs(root)
	//ans = (sum-ans)/2
	return (((sum - ans) / 2) * ((sum + ans) / 2)) % mod
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
