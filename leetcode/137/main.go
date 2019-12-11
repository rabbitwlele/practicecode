package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 1, 1, 123, 3, 3, 3}))
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
}

func singleNumber(nums []int) int {
	a, b := 0, 0
	for _, c := range nums {
		a, b = (b&c)|(a&^c), ((^a)&(^b)&c)|(b&^c)
		//fmt.Printf("%b\n%b\n%b\n-----------\n", c, a, b)
	}
	return b
}
