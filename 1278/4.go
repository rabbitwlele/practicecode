package main

func palindromePartition(s string, k int) int {
	var dp [110][110]int
	for i := 0; i < 110; i++ {
		for j := 0; j < 110; j++ {
			dp[i][j] = 100 + 10
		}
	}
	dp[0][0] = 0
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= k; j++ {
			for m := 0; m < i; m++ {
				dp[i][j] = min(dp[i][j], dp[m][j-1]+count(s[m:i]))
			}
		}
	}
	return dp[len(s)][k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func count(s string) int {
	cnt := 0
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			cnt++
		}
	}
	return cnt
}
