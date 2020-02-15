package main

import "fmt"

func minWindow(s string, t string) string {
	flag := make(map[rune]int)
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	str := []rune(s)
	for _, ch := range t {
		flag[ch]++
	}

	leftNum := len(t)
	l, r := 0, 0
	if _, ok := flag[str[l]]; ok {
		flag[rune(s[l])]--
		leftNum--
	}

	ans := "-1"
	for {
		if leftNum == 0 {
			if ans == "-1" || r-l+1 < len(ans) {
				ans = s[l : r+1]
			}
		}

		if leftNum > 0 {
			r++
			if r >= len(s) {
				break
			}
			if _, ok := flag[str[r]]; ok {
				flag[str[r]]--
				if flag[str[r]] >= 0 {
					leftNum--
				}

			}
		} else {
			if _, ok := flag[str[l]]; ok {
				flag[str[l]]++
				if flag[str[l]] > 0 {
					leftNum++
				}
			}
			l++
		}
	}
	if ans == "-1" {
		ans = ""
	}
	return ans
}
func main() {
	fmt.Println(minWindow("ACB", "ABC"))
	fmt.Println(minWindow("", "ABC"))
	fmt.Println(minWindow("ACB", ""))
	fmt.Println(minWindow("ADOBECODEBANC", "AABC"))
	fmt.Println(minWindow("A", "ABC"))
	fmt.Println(minWindow("aa", "aa"))
}
