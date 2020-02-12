package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)-1
	now := nums[l]
	//fmt.Println(nums)
	for l < r {
		for r > l && nums[r] <= now {
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] >= now {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = now
	//fmt.Println(now, l, nums)
	l++
	if l == k {
		return now
	}
	if l < k {
		return findKthLargest(nums[l:], k-l)
	} else {
		return findKthLargest(nums[:l-1], k)
	}
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 1))
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 3))
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 4))
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 5))
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 6))
}
