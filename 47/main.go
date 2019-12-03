package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 2, 1, 1, 1, 1, 1, 2}))
}

func permuteUnique(nums []int) [][]int {
	var ret [][]int
	idxs := make([]int, len(nums))
	out := make([]int, len(nums))
	for idx := 0; idx < len(nums); idx++ {
		out[idx] = -999999999999
	}
	sort.Ints(nums)
	dfs(nums, 0, idxs, out, &ret)
	return ret
}

func dfs(input []int, idx int, idxs []int, output []int, ret *[][]int) {
	//	fmt.Println("xx", output, idxs)
	if idx == len(input) {
		*ret = append(*ret, append([]int{}, output...))
		return
	}
	i := 0
	if idx != 0 && input[idx] == input[idx-1] {
		i = idxs[idx-1] + 1
	}
	for ; i < len(idxs); i++ {
		if output[i] == -999999999999 {
			output[i] = input[idx]
			idxs[idx] = i
			dfs(input, idx+1, idxs, output, ret)
			output[i] = -999999999999
			idxs[idx] = 0
		}
	}
}
