package main

import "fmt"

func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	const mod = 1000000000 + 7
	var f [510][30][2]int
	lo := toInts(s1)
	hi := toInts(s2)
	e := toInts(evil)

	for i := lo[0]; i <= hi[0]; i++ {
		if i != e[0] {
			f[0][i][0] = 1
		} else {
			f[0][i][1] = 1
		}
	}

	fmt.Println(f[0])
	for i := 1; i < n; i++ {
		for x := 0; x < 26; x++ {
			for k := 0; k < len(evil); k++ {
				if f[i-1][x][k] == 0 {
					continue
				}
				l, r := 0, 25
				if x == lo[i-1] {
					l = lo[i]
				}
				if x == hi[i-1] {
					r = hi[i]
				}
				fmt.Println(string(rune(x+'a')), x, l, r)

				for to := l; to <= r; to++ {
					if to == e[k] {
						f[i][to][k+1] = (f[i][to][k+1] + f[i-1][x][k]) % mod
					} else {
						f[i][to][k] = (f[i][to][k] + f[i-1][x][k]) % mod
					}
				}
			}
		}
		fmt.Println(f[i])
	}

	ans := 0
	for i := 0; i < 26; i++ {
		for k := 0; k < len(e); k++ {
			ans = (ans + f[n-1][i][k]) % mod
		}
	}
	return ans
}

func toInts(s string) []int {
	ret := make([]int, len(s))
	for idx, x := range s {
		ret[idx] = int(x - 'a')
	}
	return ret
}

/*
3
"szc"
"zyi"
"p"
*/
func main() {
	fmt.Println(findGoodStrings(3, "szc", "zyi", "p"))
	//fmt.Println(findGoodStrings(4, "leea", "leet", "z"))
}
