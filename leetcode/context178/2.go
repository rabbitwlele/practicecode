package main

import (
	"fmt"
	"sort"
)

type Node struct {
	n    byte
	f    [26]int
	isEx bool
}

func rankTeams(votes []string) string {
	f := make([]Node, 26)
	for i := 0; i < 26; i++ {
		f[i].n = byte('A' + i)
	}

	for i := 0; i < len(votes[0]); i++ {
		for j := 0; j < len(votes); j++ {
			f[votes[j][i]-'A'].f[i]++
			f[votes[j][i]-'A'].isEx = true
		}
	}

	sort.Slice(f, func(i, j int) bool {
		for k := 0; k < 26; k++ {
			if f[i].f[k] > f[j].f[k] {
				return true
			} else if f[i].f[k] < f[j].f[k] {
				return false
			}
		}
		return f[i].n < f[j].n
	})
	ans := []byte{}
	for i := 0; i < 26; i++ {
		if f[i].isEx {
			ans = append(ans, f[i].n)
			//fmt.Println(f[i])
		}
	}
	return string(ans)
}
func main() {
	fmt.Println(rankTeams([]string{"ABC", "ACB", "ABC", "ACB", "ACB"}))
}
