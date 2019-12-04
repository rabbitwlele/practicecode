package main

func tribonacci(n int) int {
	var ans [38]int
	ans[0] = 0
	ans[1] = 1
	ans[2] = 1
	for i := 3; i <= n; i++ {
		ans[i] = ans[i-1] + ans[i-2] + ans[i-3]
	}
	return ans[n]
}
