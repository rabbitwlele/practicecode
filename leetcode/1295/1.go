package main

func main() {

}

func findNumbers(nums []int) int {
	ret := 0
	for _, num := range nums {
		if x(num)%2 == 0 {
			ret++
		}
	}
	return ret
}

func x(x int) int {
	ret := 0
	for x != 0 {
		x /= 10
		ret++
	}
	return ret
}
