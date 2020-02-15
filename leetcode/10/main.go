package main

import "fmt"

func isMatch(s string, p string) bool {
	f := make([][]bool, len(s)+1)
	for idx, _ := range f {
		f[idx] = make([]bool, len(p)+1)
	}
	s = " " + s
	p = " " + p
	f[0][0] = true
	for i := 1; i < len(p); i++ {
		if i+1 < len(p) && p[i+1] == '*' {
			continue
		}
		if p[i] == '*' {
			f[0][i] = f[0][i] || f[0][i-2]
		}
	}
	for i := 1; i < len(s); i++ {
		for j := 1; j < len(p); j++ {
			if j+1 < len(p) && p[j+1] == '*' {
				continue
			}
			switch p[j] {
			case '.':
				f[i][j] = f[i][j] || f[i-1][j-1]
			case '*':
				f[i][j] = f[i][j] || f[i][j-2]
				if p[j-1] == '.' || p[j-1] == s[i] {
					f[i][j] = f[i][j] || f[i-1][j-2] || f[i-1][j]
				}
			default:
				if s[i] == p[j] {
					f[i][j] = f[i][j] || f[i-1][j-1]
				}
			}
			//fmt.Println(i, j, f[i][j])
		}
	}
	return f[len(s)-1][len(p)-1]
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", ".*"))
	fmt.Println(isMatch("aab", "a*c*b*"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
	fmt.Println(isMatch("mississippi", "mis*is*ip*."))
	fmt.Println(isMatch("aab", "c*d*a*b"))
	fmt.Println(isMatch("aab", "ac*d*a*b"))
	fmt.Println(isMatch("a", "aa"))
}
