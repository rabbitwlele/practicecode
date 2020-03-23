package context179

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	relations := make([][]int, len(manager))
	for i, m := range manager {
		if m == -1 {
			continue
		}
		relations[m] = append(relations[m], i)
	}

	var dfs func(r int) int
	dfs = func(r int) int {
		if len(relations[r]) == 0 {
			return 0
		}
		tmp := 0
		for i := 0; i < len(relations[r]); i++ {
			tmp = max(tmp, dfs(relations[r][i]))
		}
		return tmp + informTime[r]
	}
	return dfs(headID)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
