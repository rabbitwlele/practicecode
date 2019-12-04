package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canReorderDoubled([]int{3, 1, 3, 6}))
	fmt.Println(canReorderDoubled([]int{2, 1, 2, 6}))
	fmt.Println(canReorderDoubled([]int{4, -2, 2, -4}))
	fmt.Println(canReorderDoubled([]int{1, 2, 4, 16, 8, 4}))
}

type Arr []int

func (arr Arr) Len() int           { return len(arr) }
func (arr Arr) Swap(i, j int)      { arr[i], arr[j] = arr[j], arr[i] }
func (arr Arr) Less(i, j int) bool { return abs(arr[i]) < abs(arr[j]) }
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func canReorderDoubled(A []int) bool {
	sort.Sort(Arr(A))
	m := make(map[int]int)

	for i := 0; i < len(A); i++ {
		m[A[i]]++
	}

	for i := 0; i < len(A); i++ {
		if m[A[i]] > 0 {
			if m[A[i]] > m[A[i]*2] {
				return false
			} else {
				m[A[i]*2] -= m[A[i]]
				m[A[i]] = 0
			}
		}
	}
	return true
}
