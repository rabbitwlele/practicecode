func numEquivDominoPairs(dominoes [][]int) int {
	m := make(map[int]int)
	ans := 0
	for _, dominoe := range dominoes {
		a, b := dominoe[0], dominoe[1]
		if a > b {
			a, b = b, a
		}
		mask := a<<6 + b
		m[mask]++
	}
	for _, value := range m {
		if value >= 2 {
			ans += value * (value - 1) / 2
		}
	}
	return ans
}
