package main

import "fmt"

type Node struct {
	data int
	next *Node
}

var ringLinkList *Node
var noRingLinkList *Node

func IsRingList(node *Node) bool {
	if node == nil {
		return false
	}
	slow := node
	fast := node

	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast.data == slow.data {
			return true
		}

	}
	return false
}

func main() {
	fmt.Println(IsRingList(ringLinkList))
	fmt.Println(IsRingList(noRingLinkList))
}

//初始化有环列表和无环列表
func init() {
	for i := 0; i < 20; i++ {
		ringLinkList = &Node{20 - i, ringLinkList}
		noRingLinkList = &Node{20 - i, noRingLinkList}
	}

	node1 := ringLinkList
	for i := 0; i < 5; i++ {
		node1 = node1.next
	}
	tail := ringLinkList
	for i := 0; i < 19; i++ {
		tail = tail.next
	}
	tail.next = node1
}
