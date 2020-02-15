package main

import "fmt"

type node struct {
	idx, height int
}

func push(que []node, x node) []node {
	for len(que) != 0 && que[len(que)-1].height >= x.height {
		que = que[:len(que)-1]
	}
	que = append(que, x)
	return que
}

func largestRectangleArea(heights []int) int {
	que := []node{}
	ll := make([]int, len(heights))
	rr := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		que = push(que, node{i, heights[i]})
		left := -1
		if len(que) >= 2 {
			left = que[len(que)-2].idx
		}
		ll[i] = i - left
	}
	que = []node{}
	for i := len(heights) - 1; i >= 0; i-- {
		que = push(que, node{i, heights[i]})
		right := len(heights)
		if len(que) >= 2 {
			right = que[len(que)-2].idx
		}
		rr[i] = right - i
	}
	//fmt.Println(ll)
	//fmt.Println(rr)
	ans := 0
	for i := 0; i < len(heights); i++ {
		if ans < (rr[i]+ll[i]-1)*heights[i] {
			ans = (rr[i] + ll[i] - 1) * heights[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{}))
	fmt.Println(largestRectangleArea([]int{1, 2, 1}))
	fmt.Println(largestRectangleArea([]int{1, 2, 1, 2, 1}))
	fmt.Println(largestRectangleArea([]int{1, 2, 1, 0, 1}))
	fmt.Println(largestRectangleArea([]int{1, 2, 1, 0, 10, 20}))
}
