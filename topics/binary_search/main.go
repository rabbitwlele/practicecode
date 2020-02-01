package main

import "fmt"

func searchOnlyOne(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = r - 1
		} else {
			l = l + 1
		}
	}
	return -1
}

func searchTheFirst(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		//	fmt.Println(l, r, mid)
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[r] == target {
		return r
	}
	return -1
}

func searchTheLast(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func main() {
	fmt.Println(searchTheFirst([]int{1, 1, 2, 2, 3, 3}, 1))
	fmt.Println(searchTheFirst([]int{1, 1, 2, 2, 2, 3, 3}, 2))
	fmt.Println(searchTheFirst([]int{1, 1, 2, 2, 3, 3}, 3))
	fmt.Println(searchTheFirst([]int{1, 1, 2, 2, 3, 3}, 4))

	fmt.Println(searchTheLast([]int{1, 1, 2, 2, 3, 3}, 1))
	fmt.Println(searchTheLast([]int{1, 1, 2, 2, 3, 3}, 2))
	fmt.Println(searchTheLast([]int{1, 1, 2, 2, 3, 3}, 3))
	fmt.Println(searchTheLast([]int{1, 1, 2, 2, 3, 3}, 5))
}
