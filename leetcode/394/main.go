package main

import (
	"fmt"
)

//3[a2[c]]
func decodeString1(s string) string {
	ans, _ := dfs(s, 0)
	return ans
}

func decodeString(s string) string {
	ans, _ := dfs(s, 0)
	return ans
}
func dfs(s string, idx int) (string, int) {
	res := ""
	multi := 0
	for idx < len(s) {
		ch := s[idx]
		if ch == '[' {
			tmp, tidx := dfs(s, idx+1)
			res += mutliStr(multi, tmp)
			multi = 0
			idx = tidx
		} else if ch == ']' {
			return res, idx
		} else if '0' <= ch && ch <= '9' {
			multi = multi*10 + int(ch-'0')
		} else {
			res += string(ch)
		}
		idx++
	}
	return res, idx + 1
}

func mutliStr(mutli int, str string) string {
	ans := ""
	for i := 0; i < mutli; i++ {
		ans += str
	}
	return ans
}
func main() {
	//fmt.Println(decodeString("2[aaa]"))
	//fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("3[a]2[b2[c]e]"))
}
