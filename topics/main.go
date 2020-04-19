package main

import (
	"fmt"
)

func combinationSum1(candidates []int, target int) [][]int {
	dp := make([][][][]int, len(candidates)+1)
	for i := 0; i <= len(candidates); i++ {
		dp[i] = make([][][]int, target+1)
	}
	for i := 0; i <= len(candidates); i++ {
		dp[i][0] = [][]int{[]int{}}
	}
	for i := 1; i <= len(candidates); i++ {
		for j := candidates[i-1]; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			for _, arr := range dp[i-1][j-candidates[i-1]] {
				dp[i][j] = append(dp[i][j], append(arr, candidates[i-1]))
			}
			//	fmt.Println(i, j, dp[i][j])
		}
	}
	return dp[len(candidates)][target]
}

func combinationSum2(candidates []int, target int) [][]int {
	dp := make(map[int][][]int)
	dp[0] = [][]int{[]int{}}
	for _, num := range candidates {
		for sum, arrs := range clone(dp) {
			if sum+num > target {
				continue
			}
			for _, arr := range arrs {
				dp[sum+num] = append(dp[sum+num], append(arr, num))
			}
		}
	}
	return dp[target]
}

func clone(m map[int][][]int) map[int][][]int {
	tmp := make(map[int][][]int)
	for k, v := range m {
		tmp[k] = v
	}
	return tmp
}

func main() {
	fmt.Println(combinationSum2([]int{1, 3, 3, 7}, 7))
}
