package main
import (
	"fmt"
)
func minJumps(arr []int) int {
	m := make(map[int][]int)
	f := make([]int, len(arr))
	for idx, x := range arr {
		f[idx] = 1000000009
		m[x] = append(m[x], idx)
	}
	que := make([]int, 0)
	que = append(que, 0)
	f[0] = 0
	for len(que) != 0 {
		now := que[0]
		que = que[1:]
		if now == len(arr)+1 {
			return f[now]
		}

		for _, x := range []int{1, -1} {
			if 0 <= now+x && now+x < len(arr) && f[now]+1 < f[now+x] {
				f[now+x] = f[now] + 1
				que = append(que, now+x)
			}
		}

		for _, x := range m[arr[now]] {
			if f[now]+1 < f[x] {
				f[x] = f[now] + 1
				que = append(que, x)
			}
		}
	//	fmt.Println(f)
	}
	return f[len(arr)-1]
}

func main() {
	fmt.Println(minJumps([]int{100,-23,-23,404,100,23,23,23,3,404}))
	fmt.Println(minJumps([]int{11,22,7,7,7,7,7,7,7,22,13}))
}
