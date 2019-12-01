package main

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	if tomatoSlices%2 == 1 {
		return []int{}
	}
	tomatoSlices /= 2
	if tomatoSlices < cheeseSlices || tomatoSlices > 2*cheeseSlices {
		return []int{}
	}
	return []int{tomatoSlices - cheeseSlices, 2*cheeseSlices - tomatoSlices}
}

func main() {

}
