package main

import "fmt"

func movesToMakeZigzag(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var tmp = make([]int, len(nums)+2)
	for idx, num := range nums {
		tmp[idx+1] = num
	}
	tmp[len(nums)+1] = nums[len(nums)-1] + 1
	tmp[0] = nums[0] + 1
	var ans1 = 0
	for i := 1; i < len(tmp)-1; i += 2 {
		x := min(tmp[i-1], tmp[i+1])
		if x <= tmp[i] {
			ans1 += tmp[i] - x + 1
		}
	}
	var ans2 = 0
	for i := 2; i < len(tmp)-1; i += 2 {
		x := min(tmp[i-1], tmp[i+1])
		if x <= tmp[i] {
			ans2 += tmp[i] - x + 1
		}
	}
	return min(ans1, ans2)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	//	fmt.Println(movesToMakeZigzag([]int{1, 2, 3}))
	fmt.Println(movesToMakeZigzag([]int{2, 7, 10, 9, 8, 9}))
}
