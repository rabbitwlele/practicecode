package context182

func numTeams(rating []int) int {
	ans := 0
	for i := 0; i < len(rating); i++ {
		for j := i + 1; j < len(rating); j++ {
			for k := j + 1; k < len(rating); k++ {
				if (rating[i] < rating[j] && rating[j] < rating[k]) ||
					(rating[i] > rating[j] && rating[j] > rating[k]) {
					ans++
				}
			}
		}
	}
	return ans
}
