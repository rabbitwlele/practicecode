package main

import "sort"

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	tmp := make(map[int][]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tmp[i-j] = append(tmp[i-j], mat[i][j])
		}
	}

	for k, _ := range tmp {
		sort.Ints(tmp[k])
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := tmp[i-j][0]
			tmp[i-j] = tmp[i-j][1:]
			mat[i][j] = k
		}
	}
	return mat
}

func main() {
}
