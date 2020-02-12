package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0{
		return [][]int{}
	}
	sort.Sort(Arr(intervals))
	//fmt.Println(intervals)
	var ans [][]int
	tmp := intervals[0]
	for _, interval := range intervals {
		if interval[0] > tmp[1] {
			ans = append(ans, tmp)
			tmp = interval
		} else {
			tmp[1] = max(tmp[1], interval[1])
		}
	}
	ans = append(ans,tmp)
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Arr [][]int

func (a Arr) Len() int      { return len(a) }
func (a Arr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Arr) Less(i, j int) bool {
	if a[i][0] < a[j][0] {
		return true
	}
	if a[i][0] == a[j][0] && a[i][1] < a[j][1] {
		return true
	}
	return false
}
func main() {

}
