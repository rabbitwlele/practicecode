package main

import "fmt"

func minTaps(n int, ranges []int) int {
	f := make([]int, n+1)
	for idx, tap := range ranges {
		start := max(0, idx-tap)
		end := idx + tap
		f[start] = end
	}

	end := 0
	ans := 0
	maxEnd := -1
	for i := 0; i <= n; i++ {
		if end < i {
			return -1
		}

		maxEnd = max(maxEnd, f[i])
		if i == end && i != n {
			end = maxEnd
			maxEnd = -1
			ans++
		}
	}
	return ans
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minTaps(5, []int{3, 4, 1, 1, 0, 0}))
	fmt.Println(minTaps(3, []int{0, 0, 0, 0}))
	fmt.Println(minTaps(7, []int{1, 2, 1, 0, 2, 1, 0, 1}))
	fmt.Println(minTaps(8, []int{4, 0, 0, 0, 0, 0, 0, 0, 4}))
	fmt.Println(minTaps(8, []int{4, 0, 0, 0, 0, 0, 0, 0, 3}))
	fmt.Println(minTaps(8, []int{4, 0, 0, 0, 4, 0, 0, 0, 4}))
}
