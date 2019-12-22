package _293

import "fmt"

func main() {
	fmt.Println(shortestPath([][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 0},
	}, 1))
}

var step = [][]int{
	{1, 0},
	{0, 1},
	{0, -1},
	{-1, 0},
}

type node struct {
	i, j, k, step int
}

func shortestPath(grid [][]int, k int) int {
	n, m := len(grid), len(grid[0])
	var flag [49 * 49][49][49]int
	que := make([]node, 0)
	que = append(que, node{0, 0, 0, 0})
	flag[0][0][0] = 1

	for len(que) != 0 {
		head := que[0]
		que = que[1:]
		if head.i == n-1 && head.j == m-1 {
			return head.step
		}
		for _, s := range step {
			i, j := head.i+s[0], head.j+s[1]
			if 0 <= i && i < n && 0 <= j && j < m {
				now := node{i, j, head.k + grid[i][j], head.step + 1}
				if now.k <= k && flag[now.k][now.i][now.j] == 0 {
					flag[now.k][now.i][now.j] = 1
					que = append(que, now)
				}
			}
		}
	}
	return -1
}
