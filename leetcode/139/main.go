package main

import "fmt"

func main() {
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

func wordBreak(s string, wordDict []string) bool {
	f := make([]int, len(s)+1)
	f[0] = 1
	for i := 0; i <= len(s); i++ {
		if f[i] == 1 {
			for _, word := range wordDict {
				if len(word)+i > len(s) {
					continue
				}
				tmp := s[i : i+len(word)]
				flag := true
				for j := 0; j < len(word); j++ {
					if word[j] != tmp[j] {
						flag = false
						break
					}
				}
				if flag {
					f[i+len(word)] = 1
				}
			}
		}
	}
	return f[len(s)] == 1
}
