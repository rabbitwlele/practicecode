package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numFactoredBinaryTrees([]int{2, 4}))
	fmt.Println(numFactoredBinaryTrees([]int{2, 4, 5, 10}))
	fmt.Println(numFactoredBinaryTrees([]int{45, 42, 2, 18, 23, 1170, 12, 41, 40, 9, 47, 24, 33, 28, 10, 32, 29, 17, 46, 11, 759, 37, 6, 26, 21, 49, 31, 14, 19, 8, 13, 7, 27, 22, 3, 36, 34, 38, 39, 30, 43, 15, 4, 16, 35, 25, 20, 44, 5, 48}))
}

const mod = 1000000007

func numFactoredBinaryTrees(A []int) int {
	sort.Ints(A)
	cnts := make(map[int]int, len(A))
	for i := 0; i < len(A); i++ {
		cnts[A[i]] = 1
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j <= i; j++ {
			to := A[i] * A[j]
			if cnt, ok := cnts[to]; ok {
				if i == j {
					cnts[to] = (cnt + (cnts[A[i]]*cnts[A[j]])%mod) % mod
				} else {
					cnts[to] = (cnt + (cnts[A[i]]*cnts[A[j]]<<1)%mod) % mod
				}
			}
		}
	}

	sums := 0
	for _, cnt := range cnts {
		//fmt.Println(num, cnt)
		sums = (sums + cnt) % mod
	}
	return sums
}
