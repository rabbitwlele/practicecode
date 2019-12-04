package main

import "fmt"

func countSquares(matrix [][]int) int {
	for i := 0; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] += matrix[i-1][j]
		}
	}
	//for i := 0; i < len(matrix); i++ {
	//	fmt.Println(matrix[i])
	//}
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			for length := 1; length <= len(matrix); length++ {
				if ok(matrix, i, j, length) {
					//		fmt.Println("xxx", i, j, length)
					sum += 1
				} else {
					break
				}
			}
		}
	}
	return sum
}
func ok(matrix [][]int, i, j, length int) bool {
	i, j = i-1, j-1
	ii, jj := i+length, j+length
	//	fmt.Println("input", i, j, ii, jj)
	if ii >= len(matrix) {
		return false
	}
	if jj >= len(matrix[ii]) {
		return false
	}

	var m1, m2, m3, m4 int
	if i >= 0 && j >= 0 {
		m1 = matrix[i][j]
	}
	if i >= 0 && jj >= 0 {
		m2 = matrix[i][jj]
	}
	if ii >= 0 && j >= 0 {
		m3 = matrix[ii][j]
	}
	m4 = matrix[ii][jj]
	//fmt.Println("out", m1, m2, m3, m4)
	return (m4 - m2 - m3 + m1) == length*length
}

//[[0,1,1,1],[1,1,1,1],[0,1,1,1]]
func main() {
	fmt.Println(countSquares1([][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}))
}

func countSquares1(matrix [][]int) int {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] > 0 && i > 0 && j > 0 {
				matrix[i][j] += min(matrix[i-1][j-1], min(matrix[i-1][j], matrix[i][j-1]))
			}
			sum += matrix[i][j]
		}
		fmt.Println(matrix[i])
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
