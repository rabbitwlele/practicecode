package _292

import "fmt"

func main() {
	fmt.Println(maxSideLength([][]int{
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 1, 1},
	}, 6))
}
func maxSideLength(mat [][]int, threshold int) int {
	for i := 0; i < len(mat); i++ {
		for j := 1; j < len(mat[i]); j++ {
			mat[i][j] += mat[i][j-1]
		}
	}

	for i := 1; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			mat[i][j] += mat[i-1][j]
		}
	}
	//for i := 0; i < len(mat); i++ {
	//	fmt.Println(mat[i])
	//}
	ans := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			for length := 1; length <= len(mat); length++ {
				s := size(mat, i, j, length)
				fmt.Println(i, j, length, s)
				if s > threshold {
					break
				}
				if s == threshold {
					if length > ans {
						ans = length
					}
				}
			}
		}
	}
	return ans
}

func size(matrix [][]int, i, j, length int) int {
	i, j = i-1, j-1
	ii, jj := i+length, j+length
	//	fmt.Println("input", i, j, ii, jj)
	if ii >= len(matrix) {
		return 100000 + 9
	}
	if jj >= len(matrix[ii]) {
		return 100000 + 9
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
	return m4 - m2 - m3 + m1
}
