package main

import "fmt"

func maxJumps(arr []int, d int) int {
	var f [1010][1010]int
	var i int
	for i := 0; i < len(arr); i++ {
		f[0][i] = 1
	}
	for i = 0; ; i++ {
		flag := false
		for k := 0; k < len(arr); k++ {
			if f[i][k] == 0 {
				continue
			}
			for j := 1; j <= d; j++ {
				if j+k >= len(arr) || arr[j+k] >= arr[k] {
					break
				}
				f[i+1][k+j] = max(f[i+1][k+j], f[i][k]+1)
				flag = true
			}
			for j := 1; j <= d; j++ {
				if k-j < 0 || arr[k-j] >= arr[k] {
					break
				}
				f[i+1][k-j] = max(f[i+1][k-j], f[i][k]+1)
				flag = true
			}
		}
		//fmt.Println(f[i+1][:len(arr)])
		if !flag {
			break
		}
	}
	ans := -1
	for j := 0; j < len(arr); j++ {
		ans = max(ans, f[i][j])
	}
	return ans
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maxJumps([]int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, 2))
	fmt.Println(maxJumps([]int{8,3, 6, 5, 4, 3, 2, 1}, 2))
}

func maxJumps1(arr []int, d int) int {
	var dp = make([]int, 1001)
	res := 0
	// 遍历每一个idx
	for i := 0; i < len(arr); i++ {
		res = max(res, dfs(arr, dp, i, d))
	}
	return res
}

func dfs(arr, dp []int, i, d int) int {
	if dp[i] != 0 {
		return dp[i]
	}
	res := 1
	// (i, i+d]
	for j := i + 1; j <= min(i+d, len(arr)-1) && arr[j] < arr[i]; j++ {
		res = max(res, 1+dfs(arr, dp, j, d))
	}
	// [i-d, i)
	for j := i - 1; j >= max(0, i-d) && arr[j] < arr[i]; j-- {
		res = max(res, 1+dfs(arr, dp, j, d))
	}
	dp[i] = res
	return dp[i]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
