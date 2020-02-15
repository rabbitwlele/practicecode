package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	bytes := make([]byte, len(s))
	for idx, x := range []byte(s) {
		bytes[len(s)-idx-1] = x
	}
	ss := string(bytes)
	var f [1010][1010]int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(ss); j++ {
			if s[i] == ss[j] {
				f[i+1][j+1] = max(f[i+1][j+1], f[i][j]+1)
			}
		}
	}
	ans := 0
	a, b := 0, 0
	for i := 1; i <= len(s); i++ {
		//fmt.Println(i, len(s)-i+1, f[i][len(s)-i+1]*2-1)
		if ans < 2*f[i][len(s)-i+1]-1 {

			ans = 2*f[i][len(s)-i+1] - 1
			a, b = i-f[i][len(s)-i+1], i+f[i][len(s)-i+1]-1
		}
		//fmt.Println(i, len(s)-i, f[i][len(s)-i]*2)
		if ans < f[i][len(s)-i]*2 {
			ans = f[i][len(s)-i] * 2
			a, b = i-f[i][len(s)-i], i+f[i][len(s)-i]
		}
	}
	//fmt.Println(a, b)
	return s[a:b]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(longestPalindrome("aacdefcaa"))
	fmt.Println(longestPalindrome("aabaacasdfwefq"))
	fmt.Println(longestPalindrome("aabbaacasdfwefqcaabbaa"))
	fmt.Println(longestPalindrome("aaaa"))
}
