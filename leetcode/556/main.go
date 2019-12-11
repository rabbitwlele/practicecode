package main

import (
	"fmt"
)

func main() {
	fmt.Println(nextGreaterElement(12))
	fmt.Println(nextGreaterElement(21))
	fmt.Println(nextGreaterElement(1234))
	fmt.Println(nextGreaterElement(230241))
	fmt.Println(nextGreaterElement(2147483647))
}
func nextGreaterElement(n int) int {
	var ns []int
	tmp := n
	for tmp != 0 {
		ns = append(ns, tmp%10)
		tmp /= 10
	}

	for i := 1; i < len(ns); i++ {
		if ns[i] >= ns[i-1] {
			continue
		}

		var k int
		for k = 0; k < i; k++ {
			if ns[k] > ns[i] {
				break
			}
		}
		ns[k], ns[i] = ns[i], ns[k]
		for f, e := 0, i-1; f < e; f, e = f+1, e-1 {
			ns[f], ns[e] = ns[e], ns[f]
		}
		break
	}

	ret := 0
	for i := len(ns) - 1; i >= 0; i-- {
		ret *= 10
		ret += ns[i]
	}
	if ret == n || ret > ((1<<31)-1) {
		ret = -1
	}
	return ret
}

/*
type arr []int

func (a arr) Len() int           { return len(a) }
func (a arr) Less(i, j int) bool { return a[i] > a[j] }
func (a arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func nextGreaterElement(n int) int {
	var max = []int{4, 2, 9, 4, 9, 6, 7, 2, 9, 5}
	var ns []int
	tmp := n
	for tmp != 0 {
		ns = append(ns, tmp%10)
		tmp /= 10
	}
	sort.Sort(arr(ns))

	if len(ns) == len(max) {
		for idx, x := range max {
			for i := idx; i < len(ns); i++ {
				if ns[i] > x {
					continue
				}
				tmp := ns[i]
				for j := i; j > idx; j-- {
					ns[j] = ns[j-1]
				}
				ns[idx] = tmp
				if tmp < x {
					goto next
				}
				break
			}
		}
	}

next:
	ret := 0
	for _, x := range ns {
		ret *= 10
		ret += x
	}
	if ret == n {
		ret = -1
	}
	return ret
}
*/
