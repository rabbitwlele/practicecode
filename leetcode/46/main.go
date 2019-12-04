package main

import "fmt"

func main() {
	permute([]int{1, 2, 3, 4, 5})
}

func permute(nums []int) [][]int {
	var ret [][]int

	var dfs func(flag []int, arr []int)

	dfs = func(flag []int, arr []int) {
		if len(arr) == len(nums) {
			fmt.Println(arr)
			ret = append(ret, append([]int{}, arr...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if flag[i] == 0 {
				flag[i] = 1
				dfs(flag, append(arr, nums[i]))
				flag[i] = 0
			}
		}
	}

	dfs(make([]int, len(nums)), []int{})
	return ret
}
