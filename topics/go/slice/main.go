package main

import "fmt"

func makeTest() {
	a1 := []int{}
	var a2 []int
	a3 := new([]int)
	a4 := make([]int, 0)
	fmt.Println(a1 == nil, a2 == nil, a3 == nil, *a3 == nil, a4 == nil)

	l := make([]int, 10, 20)
	fmt.Println(len(l), cap(l))
}

func main() {
	makeTest()
}
