package main

import "fmt"

func main() {
	fmt.Println(search1([]int{1, 3, 1, 1, 1}, 3))
	fmt.Println(search1([]int{1, 3, 0, 1, 1}, 0))
	fmt.Println(search1([]int{1, 3, 0, 1, 1}, 123))
}

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	return rotateSearch(nums, 0, len(nums)-1, target)
}

func rotateSearch(nums []int, l, r, target int) bool {
	if l > r {
		return false
	}
	mid := (l + r) >> 1
	head, end := nums[l], nums[r]
	x := nums[mid]
	//fmt.Println(l, r, mid, x, head)
	if x == target {
		return true
	}

	if x == head && x == end {
		return rotateSearch(nums, l, mid-1, target) || rotateSearch(nums, mid+1, r, target)
	}

	if (target >= head && (x < head || x > target)) ||
		(target < head && (x < head && x > target)) {
		return rotateSearch(nums, l, mid-1, target)
	} else {
		return rotateSearch(nums, mid+1, r, target)
	}
}

func search1(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		head, end := nums[l], nums[r]
		x := nums[mid]

		//		fmt.Println(l, r, mid, x, head)
		if x == target {
			return true
		}

		if x == head && x == end {
			l++
			continue
		}

		if (target >= head && (x < head || x > target)) ||
			(target < head && (x < head && x > target)) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
