package main

import "fmt"

func main() {
	fmt.Println(threeSum([]int{1, 0, -1, -1}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0, 0, 0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	var ret [][]int
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	for a, _ := range m {
		m[a]--
		for b, _ := range m {
			if b >= a && m[b] > 0 {
				m[b]--
				c := 0 - a - b
				if c >= b && m[c] > 0 {
					ret = append(ret, []int{a, b, c})
				}
				m[b]++
			}
		}
		m[a]++
	}
	return ret
}
