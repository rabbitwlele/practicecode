package main

import "fmt"

func main() {
	fmt.Println(subtractProductAndSum(123))
	fmt.Println(subtractProductAndSum())
	fmt.Println(subtractProductAndSum(11))
}

func subtractProductAndSum(n int) int {
	a1, a2 := 1, 0
	for n != 0 {
		tmp := n % 10
		n /= 10
		a1 *= tmp
		a2 += tmp
	}
	if a1 < a2 {
		return a2 - a1
	}
	return a1 - a2
}
