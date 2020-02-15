package _177

func canMakePaliQueries(s string, queries [][]int) []bool {
	var counts [26][]int
	for i := 0; i < 26; i++ {
		counts[i] = make([]int, len(s)+1)
	}
	for idx, ch := range s {
		for i := 0; i < 26; i++ {
			counts[i][idx+1] = counts[i][idx]
			if int(ch-'a') == i {
				counts[i][idx+1]++
			}
		}
	}
	ans := make([]bool, 0, len(queries))
	for _, query := range queries {
		k := ((query[1] - query[0] + 1) & 1) + query[2]*2
		for i := 0; i < 26; i++ {
			k -= (counts[i][query[1]+1] - counts[i][query[0]]) & 1
		}
		ans = append(ans, k >= 0)
	}
	return ans
}
