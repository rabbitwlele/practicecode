package main

import "fmt"

func grayCode(n int) []int {

	ans := make([]int, 1<<uint(n))
	pow := 1

	for i := 1; i < (1 << uint(n)); i++ {
		if pow <= i {
			pow <<= 1
		}
		ans[i] = ans[pow-1-i] | (pow >> 1)
	}
	return ans
}

func main() {
	fmt.Println(grayCode(3))
}
