package main

import "fmt"

func largest1BorderedSquare(grid [][]int) int {
	var ans = 0
	var w, h = len(grid[0]), len(grid)
	var dp [2][110][110]int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == 1 {
				dp[0][i][j], dp[1][i][j] = 1, 1
				if j-1 >= 0 {
					dp[0][i][j] += dp[0][i][j-1]
				}
				if i-1 >= 0 {
					dp[1][i][j] += dp[1][i-1][j]
				}
			}

			for l := 0; l < 100; l++ {
				si, sj := i-l, j-l
				if si < 0 || sj < 0 {
					break
				}
				if dp[1][i][sj] > l && dp[0][si][j] > l &&
					dp[1][i][j] > l && dp[0][i][j] > l {
					if ans <= l {
						ans = l + 1
					}
				}
			}
		}
	}
	return ans * ans
}

func main() {
	fmt.Println(largest1BorderedSquare([][]int{
		{1, 1, 1}, {1, 0, 1}, {1, 1, 1},
	}))
}
