package main

import "fmt"

type union struct {
	f []int
}

func NewUnion(n int) *union {
	f := make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = i
	}
	return &union{f}
}
func (list *union) GetF(x int) int {
	if list.f[x] == x {
		return x
	}
	list.f[x] = list.GetF(list.f[x])
	return list.f[x]
}

func (list *union) Add(a, b int) {
	list.f[list.GetF(a)] = list.GetF(b)
}

func makeConnected(n int, connections [][]int) int {
	if n > len(connections)+1 {
		return -1
	}

	un := NewUnion(n)
	for _, connection := range connections {
		un.Add(connection[0], connection[1])
		//fmt.Println(un.f)
	}
	ans := -1
	for i := 0; i < n; i++ {
		//fmt.Println(i,un.GetF(i),un.f[i])
		if i == un.GetF(i) {
			ans++
		}
	}
	return ans
}
func main() {
	un := NewUnion(5)
	fmt.Println(1)
	un.Add(1, 2)
	fmt.Println(1)
	un.Add(2, 3)
	fmt.Println(1)

	for i := 0; i < 5; i++ {
		fmt.Println(i, un.GetF(i))
	}
}
