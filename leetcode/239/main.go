package main

import "fmt"

func push(que []int, x int) []int {
	for len(que) != 0 && que[len(que)-1] < x {
		que = que[:len(que)-1]
	}
	que = append(que, x)
	return que
}

func pop(que []int, x int) []int {
	if len(que) != 0 && que[0] == x {
		return que[1:]
	}
	return que
}

func maxSlidingWindow(nums []int, k int) []int {
	que := []int{}
	ans := []int{}
	if len(nums) == 0 {
		return ans
	}
	for i := 0; i < k-1; i++ {
		que = push(que, nums[i])
	}
	for i := k - 1; i < len(nums); i++ {
		que = push(que, nums[i])
		ans = append(ans, que[0])
		que = pop(que, nums[i-k+1])
	}
	return ans
}
func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
