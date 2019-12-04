type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	nums := make([]int, n+1)
	nodes := make([]*TreeNode, n+1)
	dfs(root, nums, nodes)
	//fmt.Println(nums)
	var c1, c2 int
	for _, node := range nodes {
		if node != nil && node.Val == x {
			if node.Left != nil {
				c1 = node.Left.Val
			}
			if node.Right != nil {
				c2 = node.Right.Val
			}
			break
		}
	}
	//fmt.Println(c1,c2)
	//fmt.Println(nums[c1],nums[c2],(n-nums[x]))
	if nums[c1]*2 > n {
		return true
	}
	if nums[c2]*2 > n {
		return true
	}
	if (n-nums[x])*2 > n {
		return true
	}
	return false
}

func dfs(root *TreeNode, ans []int, nodes []*TreeNode) int {
	if root == nil {
		return 0
	}
	nodes[root.Val] = root
	ans[root.Val] = dfs(root.Left, ans, nodes) + dfs(root.Right, ans, nodes) + 1
	return ans[root.Val]
}
