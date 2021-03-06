package main

import (
	"fmt"
	"sort"
)

func main() {
	//combinationSum21([]int{10, 1, 2, 7, 6, 1, 5, 1, 1, 1, 1, 1, 1}, 8)
	combinationSum21([]int{1, 1, 1, 1, 2, 2}, 5)
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var ret [][]int
	var nums []int
	var cnts []int

	for i := 0; i < len(candidates); i++ {
		if len(nums) != 0 && candidates[i] == nums[len(nums)-1] {
			cnts[len(cnts)-1]++
		} else {
			nums = append(nums, candidates[i])
			cnts = append(cnts, 1)
		}
	}

	var dfs func(i, sum int, arr []int)
	dfs = func(i, sum int, arr []int) {
		if sum == target {
			//fmt.Println(arr)
			ret = append(ret, append([]int{}, arr...))
			return
		}
		if sum > target {
			return
		}

		for ; i < len(nums); i++ {
			var add []int
			for j := 0; j < cnts[i]; j++ {
				add = append(add, nums[i])
				dfs(i+1, sum+nums[i]*(j+1), append(arr, add...))
			}
		}
	}
	dfs(0, 0, []int{})
	return ret
}
func combinationSum21(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var ret [][]int
	nums := candidates

	var dfs func(i, sum int, arr []int)
	dfs = func(i, sum int, arr []int) {
		if sum == target {
			ret = append(ret, append([]int{}, arr...))
			fmt.Println(arr)
			return
		}
		if sum > target {
			return
		}
		for j := i; j < len(nums); j++ {
			if j > i && nums[j] == nums[j-1] {
				continue
			}
			fmt.Println(j, nums[j])
			dfs(j+1, sum+nums[j], append(arr, nums[j]))
		}
	}
	dfs(0, 0, []int{})
	return ret
}
