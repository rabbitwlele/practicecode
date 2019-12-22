package main

import "fmt"

func main() {
	fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	fmt.Println(canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
}

func canCross(stones []int) bool {
	f := make([]map[int]bool, len(stones))
	idxStone := make(map[int]int)
	for idx := 0; idx < len(stones); idx++ {
		f[idx] = make(map[int]bool)
		idxStone[stones[idx]] = idx
	}
	f[0][0] = true

	for i := 0; i < len(stones)-1; i++ {
		for k, _ := range f[i] {
			for _, d := range []int{1, -1} {
				idx := idxStone[stones[i]+k+d]
				if idx != 0 {
					f[idx][k+d] = true
				}
			}
		}
	}
	return len(f[len(f)-1]) != 0
}
