package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[rune]int)
	l, r := 0, 0
	str := []rune(s)
	m[str[r]] = 1
	ans := 1
	for {
		//fmt.Println(l, r)
		if r-l+1 == len(m) {
			ans = max(ans, r-l+1)
			r++
			if r >= len(str) {
				break
			}
			m[str[r]]++
		} else {
			m[str[l]]--
			if m[str[l]] == 0 {
				delete(m, str[l])
			}
			l++
		}
	}
	return ans
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcda"))
}
