package main

import (
	"fmt"
	"sort"
)

func main() {
	x := permuteUnique1([]int{-1, -1, 1, 1, 2, 2, -1})

	for _, arr := range x {
		cnt := 0
		for _, arr1 := range x {
			if equal(arr, arr1) {
				cnt++
			}
		}
		if cnt > 1 {
			fmt.Println(arr)
		}
	}
	//	fmt.Println("000000000000")

}

func equal(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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
func permuteUnique1(nums []int) [][]int {
	var ret [][]int
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	var dfs func(int, []int)
	dfs = func(level int, out []int) {
		if level == len(nums) {
			ret = append(ret, append([]int{}, out...))
			return
		}
		for n, cnt := range m {
			if cnt != 0 {
				m[n]--
				dfs(level+1, append(out, n))
				m[n]++
			}
		}
	}
	dfs(0, []int{})
	return ret
}
