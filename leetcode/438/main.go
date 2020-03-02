package main

import "fmt"

func findAnagrams(s string, p string) []int {
	num := len(p)
	m := make(map[byte]int)
	for _, ch := range []byte(p) {
		m[ch]++
	}
	l, r := 0, -1
	ans := []int{}
	for {
		if num > 0 {
			r++
			if r >= len(s) {
				break
			}
			if cnt, ok := m[s[r]]; ok {
				m[s[r]]--
				if cnt > 0 {
					num--
				}
			}
		} else {
			if cnt, ok := m[s[l]]; ok {
				m[s[l]]++
				if cnt >= 0 {
					num++
				}
			}
			l++
		}
		//fmt.Println(m, num, l, r)
		if num == 0 && r-l+1 == len(p) {
			ans = append(ans, l)
		}
	}
	return ans
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}
