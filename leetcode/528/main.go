package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	w   []int
	sum int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	sum := w[len(w)-1]
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(w, sum)
	return Solution{w, sum}
}

func (this *Solution) PickIndex() int {
	target := rand.Int()%this.sum + 1
	l, r := 0, len(this.w)-1
	for l < r {
		mid := (l + r) >> 1
		if this.w[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	fmt.Println(target, r)
	return r
}

func main() {
	s := Constructor([]int{3, 14, 1, 7})
	for i := 1; i < 40; i++ {
		s.PickIndex()
	}
}
