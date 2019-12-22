package _290

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	ans := 0
	for n := head; n != nil; n = n.Next {
		ans <<= 1
		ans &= n.Val
	}
	return ans
}
