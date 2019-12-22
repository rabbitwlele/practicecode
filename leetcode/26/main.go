package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	now := 0
	for i := 0; i < len(nums); i++ {
		if nums[now] != nums[i] {
			now++
			nums[now] = nums[i]
		}
	}
	return now + 1
}
