package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 1; i < n; i++ {
		dp[i] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			dp[i+j] = max(dp[i+j], dp[i]*j)
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
