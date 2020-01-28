package main

import "fmt"

func minimumDistance(word string) int {
	var dp [310][26][26]int
	for i := 1; i < 310; i++ {
		for j := 0; j < 26; j++ {
			for z := 0; z < 26; z++ {
				dp[i][j][z] = 1999999999
			}
		}
	}
	for idx, ch := range []rune(word) {
		z := ascii(ch)
		//fmt.Println(string(ch), z)
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				a := rune(i + 'A')
				b := rune(j + 'A')
				//fmt.Println(z, i, j)
				dp[idx+1][i][z] = min(dp[idx+1][i][z], dp[idx][i][j]+dis(b, ch))
				dp[idx+1][z][j] = min(dp[idx+1][z][j], dp[idx][j][i]+dis(a, ch))
			}
		}
	}
	ans := 1999999999
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			ans = min(ans, dp[len(word)][i][j])
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func dis(a, b rune) int {
	tmp := ascii(a)
	ai := tmp / 6
	aj := tmp % 6
	tmp = ascii(b)
	bi := tmp / 6
	bj := tmp % 6
	if tmp < 0 {
		tmp = -tmp
	}
	return abs(ai-bi) + abs(aj-bj)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func ascii(ch rune) int {
	return int(ch - 'A')
}

func main() {
	fmt.Println(minimumDistance("CAKE"))
	fmt.Println(minimumDistance("HAPPY"))
	fmt.Println(minimumDistance("NEW"))
	fmt.Println(minimumDistance("YEAR"))
}
