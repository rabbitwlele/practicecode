package context179

func numTimesAllBlue(light []int) int {
	max := 0
	ans := 0
	for i := 0; i < len(light); i++ {
		if light[i] > max {
			max = light[i]
		}
		if max == i+1 {
			ans++
		}
	}
	return ans
}
