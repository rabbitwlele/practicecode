package main

import "fmt"

func main() {
	//fmt.Println(minFlips([][]int{
	//	{0, 0},
	//	{0, 1},
	//}))
	fmt.Println(minFlips([][]int{
		{1, 1, 1},
		{1, 0, 1},
		{0, 0, 0},
	}))
}

func minFlips(mat [][]int) int {
	if toInt(mat) == 0 {
		return 0
	}
	n, m := len(mat), len(mat[0])

	flag := make(map[int]int)
	que := make([]data, 0)
	que = append(que, data{mat, 0})
	flag[toInt(mat)] = 0
	for len(que) != 0 {
		now := que[0]
		que = que[1:]
		//fmt.Println("000000000000000")
		//print(now.mat)
		//fmt.Printf("%d\n------------\n", now.cnt)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				new := filp(now.mat, i, j)
				tmp := toInt(new)
				//fmt.Println(i, j, now.cnt+1)
				//print(new)
				if tmp == 0 {
					return now.cnt + 1
				}
				if _, ok := flag[tmp]; !ok {
					flag[tmp] = now.cnt + 1
					que = append(que, data{new, now.cnt + 1})
				}
			}
		}
	}
	return -1
}

type data struct {
	mat [][]int
	cnt int
}

func filp(mat [][]int, ii, jj int) [][]int {
	ret := make([][]int, len(mat))

	for i := 0; i < len(mat); i++ {

		ret[i] = make([]int, len(mat[i]))

		for j := 0; j < len(mat[i]); j++ {
			if (i == ii && j == jj) ||
				(i == ii-1 && j == jj) || (i == ii && j == jj-1) ||
				(i == ii+1 && j == jj) || (i == ii && j == jj+1) {
				ret[i][j] = 1 ^ mat[i][j]
			} else {
				ret[i][j] = mat[i][j]
			}
		}
	}
	return ret
}

func toInt(mat [][]int) int {
	ret := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			ret <<= 1
			ret |= mat[i][j]
		}
	}
	return ret
}

func print(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		fmt.Println(mat[i])
	}
}
