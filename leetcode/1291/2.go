package _291

import "fmt"

func main() {
	fmt.Println(sequentialDigits(1000, 130000000))
}

const MAX = 1000000000 + 9

func sequentialDigits(low int, high int) []int {
	ans := make([]int, 0)
	for n := 2; n <= 9; n++ {
		for h := 1; h <= 9; h++ {
			x := a(h, n)
			//fmt.Println(h, n, x)
			if x < low {
				continue
			}
			if x > high {
				break
			}
			ans = append(ans, x)
		}
	}
	return ans
}

func a(h, n int) int {
	if h+n-1 > 9 {
		return MAX
	}
	ans := 0
	for n > 0 {
		ans *= 10
		ans += h
		h++
		n--
	}
	return ans
}
