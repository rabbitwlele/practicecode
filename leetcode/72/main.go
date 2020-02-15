package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	word1 = " " + word1
	word2 = " " + word2
	f := make([][]int, len(word1))
	for idx, _ := range f {
		f[idx] = make([]int, len(word2))
	}
	for i := 0; i < len(word2); i++ {
		f[0][i] = i
	}
	for i := 0; i < len(word1); i++ {
		f[i][0] = i
	}
	for i := 1; i < len(word1); i++ {
		for j := 1; j < len(word2); j++ {
			if word1[i] == word2[j] {
				f[i][j] = f[i-1][j-1]
			} else {
				f[i][j] = f[i-1][j-1] + 1
			}
			f[i][j] = min(f[i][j], min(f[i-1][j], f[i][j-1])+1)
			//fmt.Println(i, j, f[i][j])
		}
	}
	return f[len(word1)-1][len(word2)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minDistance("intention", "execution"))
	fmt.Println(minDistance("", "a"))
	fmt.Println(minDistance("a", ""))
}
