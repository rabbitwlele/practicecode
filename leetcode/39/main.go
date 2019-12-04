package main

func main() {
	combinationSum([]int{2, 3, 6, 7}, 7)
	combinationSum([]int{2, 3, 5}, 8)
}

func combinationSum(candidates []int, target int) [][]int {
	var ret [][]int

	var dfs func(i int, sum int, arr []int)
	dfs = func(i int, sum int, arr []int) {
		if sum == target {
			//fmt.Println(arr)
			ret = append(ret, append([]int{}, arr...))
			return
		}
		if sum > target {
			return
		}

		for ; i < len(candidates); i++ {
			dfs(i, sum+candidates[i], append(arr, candidates[i]))
		}
	}
	dfs(0, 0, []int{})
	return ret
}
