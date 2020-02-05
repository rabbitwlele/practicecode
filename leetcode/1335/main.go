package main

import "fmt"

func minDifficulty(jobDifficulty []int, d int) int {
	var f [12][310]int
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 300; j++ {
			f[i][j] = 1000000009
		}
	}
	f[0][0] = 0

	for i := 0; i < d; i++ {
		for j := i; j < len(jobDifficulty); j++ {
			tmp := -1
			for k := j; k < len(jobDifficulty); k++ {
				tmp = max(tmp, jobDifficulty[k])
				//fmt.Println(tmp)
				f[i+1][k+1] = min(f[i][j]+tmp, f[i+1][k+1])
			}
		}
		//fmt.Println(f[i+1][:len(jobDifficulty)+1])
	}
	if f[d][len(jobDifficulty)] == 1000000009 {
		return -1
	}
	return f[d][len(jobDifficulty)]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minDifficulty([]int{6, 5, 4, 3, 2, 1}, 2))
	fmt.Println(minDifficulty([]int{9, 9, 9}, 4))
	fmt.Println(minDifficulty([]int{1, 1, 1}, 3))
	fmt.Println(minDifficulty([]int{11, 111, 22, 222, 33, 333, 44, 444}, 6))
	fmt.Println(minDifficulty([]int{7, 1, 7, 1, 7, 1}, 3))
}
