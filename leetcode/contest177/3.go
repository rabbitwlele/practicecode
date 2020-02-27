package contest177

import "math"

func mySqrt(x int) int {
	f := float64(x)
	ff := math.Sqrt(f)
	return int(math.Floor(ff))
}

func a(x int) []int {
	for i := mySqrt(x); i >= 1; i-- {
		if x%i == 0 {
			return []int{i, x / i}
		}
	}
	return []int{1, x}
}
func closestDivisors(num int) []int {
	ans1 := a(num + 1)
	ans2 := a(num + 2)
	if ans1[1]-ans1[0] < ans2[1]-ans2[0] {
		return ans1
	}
	return ans2
}
