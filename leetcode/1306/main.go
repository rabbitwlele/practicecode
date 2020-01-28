package main

func canReach(arr []int, start int) bool {
	flag := make([]bool, len(arr))
	var que []int
	que = append(que, start)
	flag[start] = true

	for len(que) != 0 {
		now := que[0]
		que = que[1:]
		if arr[now] == 0 {
			return true
		}
		for _, i := range []int{1, -1} {
			next := now + i*arr[now]
			if next >= 0 && next < len(arr) && !flag[next] {
				flag[next] = true
				que = append(que, next)
			}
		}
	}

	return false
}
