package main

import "sort"

func main() {

}

func isPossibleDivide(nums []int, k int) bool {
	sort.Ints(nums)
	m := make(map[int]int)
	for idx, num := range nums {
		m[num]++
	}

	for _, num := range nums {
		if m[num] != 0 {
			for i := num; i < num+k; i++ {
				if m[i] < 0 {
					return false
				} else {
					m[i]--
				}
			}
		}
	}
	return true
}
