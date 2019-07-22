package main

func shortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
	const max = 19999999
	const red, blue = 0, 1
	var g [2][101][101]int
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			g[red][i][j] = max
			g[blue][i][j] = max
		}
	}
	for _, path := range red_edges {
		g[red][path[0]][path[1]] = 1
	}
	for _, path := range blue_edges {
		g[blue][path[0]][path[1]] = 1
	}
	var dp [2][101]int
	for i := 0; i < n; i++ {
		dp[red][i] = max
		dp[blue][i] = max
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for mid := 0; mid < n; mid++ {
				dp[red][j] = min(dp[red][j], g[blue][0][mid]+g[red][mid][j])
				dp[blue][j] = min(dp[blue][j], g[red][0][mid]+g[blue][mid][j])
			}
		}
	}
	var ans = make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = min(dp[red][i], dp[blue][i])
		if ans[i] >= max {
			ans[i] = -1
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
