package _353

import (
	"container/heap"
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
	return pq[i].Priority < pq[j].Priority
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
	heap.Fix(pq, item.Index)
}

func maxEvents(events [][]int) int {
	var evts [10010][]int
	for _, event := range events {
		evts[event[0]] = append(evts[event[0]], event[1])
	}
	que := &PriorityQueue{}
	ans := 0
	for i := 1; i <= 10000; i++ {
		for _, end := range evts[i] {
			heap.Push(que, &Item{Priority: end})
		}
		for que.Len() != 0 {
			tmp := heap.Pop(que).(*Item).Priority
			if tmp >= i {
				ans++
				break
			}
		}
	}
	return ans
}
