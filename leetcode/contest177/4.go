package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestMultipleOfThree(digits []int) string {
	var f [3][]int
	for _, digit := range digits {
		f[digit%3] = append(f[digit%3], digit)
	}
	sort.Slice(f[2], func(i, j int) bool {
		return f[2][i] > f[2][j]
	})
	sort.Slice(f[1], func(i, j int) bool {
		return f[1][i] > f[1][j]
	})
	if len(f[1]) < len(f[2]) {
		f[1], f[2] = f[2], f[1]
	}

	diff := (len(f[2]) - len(f[1])) % 3

	f[0] = append(f[0], f[1][:]...)
	f[0] = append(f[0], f[2][:]...)

	sort.Slice(f[0], func(i, j int) bool {
		return f[0][i] > f[0][j]
	})

	//fmt.Println(f[0])
	ans := ""
	for _, digit := range f[0] {
		if len(ans) == 0 && digit == 0 {
			ans = "0"
			break
		}
		ans += strconv.Itoa(digit)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(largestMultipleOfThree([]int{8, 6, 7, 1, 0}))
	fmt.Println(largestMultipleOfThree([]int{8, 1, 9}))
	fmt.Println(largestMultipleOfThree([]int{1}))
	fmt.Println(largestMultipleOfThree([]int{0, 0, 0, 0}))
	fmt.Println(largestMultipleOfThree([]int{8, 6, 7, 1, 0, 0, 0}))
	fmt.Println(largestMultipleOfThree([]int{2, 2, 2, 2, 2, 2, 2, 2, 5, 5, 5, 0}))
}
