package main

func findUnsortedSubarray(nums []int) int {
	ll := make([]int, len(nums))
	rr := make([]int, len(nums))
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] >= max {
			ll[i] = 1
			max = ll[i]
		}
	}
	min := nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= min {
			min = nums[i]
			rr[i] = 1
		}
	}
	l, r := len(nums), -1
	for i := 0; i < len(nums); i++ {
		if ll[i] != 1 || rr[i] != 1 {
			l = min(l, i)
			r = max(r, i)
		}
	}
	if l >= r {
		return 0
	}
	return r - l + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
