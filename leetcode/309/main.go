package main

import "fmt"

func maxProfit(prices []int) int {
	var f [2][3]int
	now, pre := 0, 1
	for idx, price := range prices {
		f[now][0] = f[pre][2]
		if idx > 0 {
			f[now][0] = max(f[now][0], f[pre][0]+price-prices[idx-1])
		}
		f[now][2] = max(f[pre][0], f[pre][2])
		//fmt.Println(price, f[now])
		now, pre = pre, now
	}

	return max(f[pre][0], f[pre][2])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1, 2, 3, 19, 2}))
	fmt.Println(maxProfit([]int{6, 1, 6, 4, 3, 0, 2}))
}
