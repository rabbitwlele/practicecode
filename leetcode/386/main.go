package main

import (
	"fmt"
)

func lexicalOrder(n int) []int {
	var ans []int
	var dfs func(x int)
	dfs = func(x int) {
		if x > n {
			return
		}
		ans = append(ans, x)
		x *= 10
		for i := 0; i <= 9; i++ {
			dfs(x + i)
		}
	}
	for i := 1; i <= 9; i++ {
		dfs(i)
	}
	return ans
}
func main() {
	fmt.Println(lexicalOrder(10100))
}
