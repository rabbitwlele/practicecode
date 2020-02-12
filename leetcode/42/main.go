package main

import (
	"fmt"
)

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	ll := make([]int, len(height))
	rr := make([]int, len(height))
	ll[0] = height[0]
	for i := 1; i < len(height); i++ {
		ll[i] = max(height[i], ll[i-1])
	}
	rr[len(rr)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rr[i] = max(height[i], rr[i+1])
	}
	ans := 0
	for i := 1; i < len(height)-1; i++ {
		tmp := min(rr[i+1], ll[i-1])
		//fmt.Println(i, ll[i-1], rr[i+1], height[i])
		if tmp > height[i] {
			ans += tmp - height[i]
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
