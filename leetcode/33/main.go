package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 4))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 5))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 6))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 7))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 1))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 2))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 123))

	fmt.Println(search([]int{4, 5, 6, 7, 0, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 1, 2}, 1))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	head := nums[0]
	for l < r {
		mid := (l + r) >> 1
		x := nums[mid]
		//		fmt.Println("b:", l, r, mid, x)
		if (target >= head && (x < head || x >= target)) ||
			(target < head && (x < head && x >= target)) {
			r = mid
		} else {
			l = mid + 1
		}
		//		fmt.Println("e:", l, r, mid, x)
	}
	if nums[r] == target {
		return r
	}
	return -1
}
