package main

import "fmt"

func main() {
	fmt.Println(smallestDivisor([]int{1, 2, 5, 9}, 6))
	fmt.Println(smallestDivisor([]int{1, 2, 3}, 6))
	fmt.Println(smallestDivisor([]int{2, 3, 5, 7, 11}, 11))
}

func smallestDivisor(nums []int, threshold int) int {
	ans := func(d int) int {
		ret := 0
		for _, num := range nums {
			ret += num / d
			if num%d != 0 {
				ret++
			}
		}
		return ret
	}

	l, r := 1, 1000000+10
	for l <= r {
		mid := (l + r) >> 1
		cnt := ans(mid)
		if cnt > threshold {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
