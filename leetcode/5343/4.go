package _343

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Value    int // 值
	Priority int // 优先级
	Index    int // 元素在堆中的索引
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

// 注意这里是优先级大的在前
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	n := len(*pq)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]

	x.Index = -1
	*pq = old[0 : n-1]
	return x
}

func (pq *PriorityQueue) Update(item *Item, value int, priority int) {
	item.Value = value
	item.Priority = priority

	// Fix操作比 先Remove再Push要快
	heap.Fix(pq, item.Index)
}

func isPossible(target []int) bool {
	var q = &PriorityQueue{}
	sum := 0

	for idx, a := range target {
		if a != 1 {
			heap.Push(q, &Item{
				Value:    idx,
				Priority: a,
			})
		}
		sum += a
	}

	for q.Len() != 0 {
		x := heap.Pop(q).(*Item)
		to := x.Priority - sum + x.Priority
		sum -= x.Priority - to
		//fmt.Println(x, sum, to)
		if to <= 0 {
			return false
		}
		if to != 1 {
			heap.Push(q, &Item{
				Value:    x.Value,
				Priority: to,
			})
		}
	}
	return true
}

func main() {
	fmt.Println(isPossible([]int{9, 3, 5}))
	fmt.Println(isPossible([]int{1, 1, 1, 1}))
}
