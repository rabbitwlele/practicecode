package main

import "fmt"

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 3}))
	fmt.Println(firstMissingPositive([]int{1, 1, -1, 2}))
	fmt.Println(firstMissingPositive([]int{9, 3, 1, 7, 2, 5}))
}

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] != i+1 {
			if nums[i] > len(nums) || nums[i] == nums[nums[i]-1] {
				nums[i] = 0
			} else {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			}
		}
	}
	//fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}
