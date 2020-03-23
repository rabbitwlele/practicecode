package main

import "fmt"

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	e := make([][]int, n+1)
	for _, edge := range edges {
		e[edge[0]] = append(e[edge[0]], edge[1])
		e[edge[1]] = append(e[edge[1]], edge[0])
	}
	e[1] = append(e[1], 0)

	f := make([][]int, t+1)
	flag := make([]bool, n+1)

	for i := 0; i <= t; i++ {
		f[i] = make([]int, n+1)
	}
	flag[1] = true
	f[0][1] = 1

	for tt := 1; tt <= t; tt++ {
		for i := 1; i <= n; i++ {
			if f[tt-1][i] == 0 {
				continue
			}

			sum := len(e[i]) - 1
			if sum == 0 {
				f[tt][i] = f[tt-1][i]
			} else {
				for _, k := range e[i] {
					if !flag[k] {
						f[tt][k] = f[tt-1][i] * sum
						flag[k] = true
					}
				}
			}
		}
	}

	if f[t][target] == 0 {
		return 0
	}
	return 1 / float64(f[t][target])
}

func main() {
	fmt.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 2, 4))
	fmt.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 1, 7))
	fmt.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 20, 6))
}
