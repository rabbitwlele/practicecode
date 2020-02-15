package _69

func majorityElement(nums []int) int {
	l, r := 0, 0

	for {
		if nums[l] == -99999999 {
			l++
		} else {
			if nums[l] == nums[r] {
				r++
			} else {
				nums[l], nums[r] = -99999999, -99999999
				l++
				r++
			}
		}
		if r >= len(nums) {
			break
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != -99999999 {
			return nums[i]
		}
	}
	return 0
}
