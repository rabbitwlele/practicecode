package context178

import "container/heap"

// An Item is something we manage in a priority queue.
type Item struct {
	value    node // The value of the item; arbitrary.
	priority int  // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value node, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

type node struct {
	x, y int
}

/*
1 ，下一步往右走，也就是你会从 grid[i][j] 走到 grid[i][j + 1]
2 ，下一步往左走，也就是你会从 grid[i][j] 走到 grid[i][j - 1]
3 ，下一步往下走，也就是你会从 grid[i][j] 走到 grid[i + 1][j]
4 ，下一步往上走，也就是你会从 grid[i][j] 走到 grid[i - 1][j]
*/
func minCost(grid [][]int) int {
	xx := []int{0, 0, 0, 1, -1}
	yy := []int{0, 1, -1, 0, 0}
	que := &PriorityQueue{}
	var f [110][110]int
	mm, nnn := len(grid), len(grid[0])
	heap.Push(que, &Item{priority: 0, value: node{0, 0}})

	for que.Len() != 0 {
		n := heap.Pop(que).(*Item)
		nn := n.value
		if nn.x < 0 || nn.x >= mm || nn.y < 0 || nn.y >= nnn {
			continue
		}
		if n.priority >= f[nn.x][nn.y] {
			continue
		}
		f[nn.x][nn.y] = n.priority

		heap.Push(que, &Item{priority: n.priority,
			value: node{
				x: nn.x + xx[grid[nn.x][nn.y]],
				y: nn.y + yy[grid[nn.x][nn.y]],
			}})
		for i := 1; i < 5; i++ {
			if i == grid[nn.x][nn.y] {
				continue
			}
			heap.Push(que, &Item{priority: n.priority,
				value: node{
					x: nn.x + xx[i],
					y: nn.y + yy[i],
				}})
		}
	}
	return f[mm-1][nnn-1]
}
