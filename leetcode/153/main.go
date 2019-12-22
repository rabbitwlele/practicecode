package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{4, 5, 6, 7, 2}))
	fmt.Println(findMin([]int{1, 2, 3, 4}))
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	head := nums[l]
	var mid int
	for l < r {
		mid = (l + r) >> 1
		if nums[mid] < head {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return min(nums[r], head)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
