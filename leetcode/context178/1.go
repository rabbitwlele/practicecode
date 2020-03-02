package context178

func smallerNumbersThanCurrent(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				ans[i]++
			}
		}
	}
	return ans
}
