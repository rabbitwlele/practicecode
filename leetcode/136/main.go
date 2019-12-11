package main

func singleNumber(nums []int) int {
	var sum int
	for _, num := range nums {
		sum ^= num
	}
	return sum
}
