package main

import (
	"fmt"
)

func main() {
	//	fmt.Println(maxFreq("aababcaab", 2, 3, 4))
	//	fmt.Println(maxFreq("aabcabcab", 2, 2, 4))
	fmt.Println(maxFreq("babcbceccaaacddbdaedbadcddcbdbcbaaddbcabcccbacebda", 1, 1, 1))
}
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	ret := make(map[string]int)
	for l := minSize; l <= minSize; l++ {
		flag := make(map[byte]int)

		for i := 0; i < l; i++ {
			flag[s[i]]++
		}
		if len(flag) <= maxLetters {
			ret[s[0:l]]++
		}

		for i := l; i < len(s); i++ {
			flag[s[i-l]]--
			flag[s[i]]++
			if flag[s[i-l]] == 0 {
				delete(flag, s[i-l])
			}
			//fmt.Println(flag)
			if len(flag) <= maxLetters {
				ret[s[i-l+1:i+1]]++
			}
		}
	}
	x := 0
	for _, cnt := range ret {
		if cnt > x {
			x = cnt
		}
	}
	//fmt.Println(ret)

	return x
}
