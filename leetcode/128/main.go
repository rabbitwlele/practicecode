package main

import "fmt"

func add(m map[int]int, a, b int) {
	m[a] = get(m, b)
}
func get(m map[int]int, x int) int {
	if _, ok := m[x]; !ok {
		return x
	}
	m[x] = get(m, m[x])
	return m[x]
}
func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		add(m, num, num-1)
	}
	ans := 0
	for _, num := range nums {
		f := get(m, num)
		//fmt.Println(num, f)
		if num-f > ans {
			ans = num - f
		}
	}
	return ans
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 5, 4, 200, 13, 3, 2, 6}))
	fmt.Println(longestConsecutive([]int{}))
}
