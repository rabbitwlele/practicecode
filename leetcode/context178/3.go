package context178

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	h := head
	l := 0
	for h != nil {
		l++
		h = h.Next
	}
	ans := false
	var dfs func(r *TreeNode, arr []int)
	dfs = func(r *TreeNode, arr []int) {
		if ans {
			return
		}
		if r == nil {
			return
		}
		arr = append(arr, r.Val)
		if same(l, head, arr) {
			ans = true
			return
		}
		dfs(r.Left, append([]int{}, arr...))
		dfs(r.Right, append([]int{}, arr...))
	}
	dfs(root, []int{})
	return ans
}
func same(l int, head *ListNode, arr []int) bool {
	if l > len(arr) {
		return false
	}
	h := head
	for i := len(arr) - l; i < len(arr); i++ {
		if arr[i] != h.Val {
			return false
		}
		h = h.Next
	}
	return true
}
