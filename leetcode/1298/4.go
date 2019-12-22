package main

func main() {
	maxCandies([]int{1, 1, 0, 0, 0, 0})
}
func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	var ans = 0
	var hasKeys [1010]bool
	flag := false
	for idx, x := range status {
		flag = true
		hasKeys[idx] = (x == 1)
	}
	boxs := make(map[int]int)
	for _, box := range initialBoxes {
		flag = true
		boxs[box] = 1
	}

	for flag {
		flag = false
		newboxs := make(map[int]int)
		for now, _ := range boxs {
			if hasKeys[now] {
				ans += candies[now]
				for _, key := range keys[now] {
					hasKeys[key] = true
					flag = true
				}
				for _, box := range containedBoxes[now] {
					newboxs[box] = 1
					flag = true
				}
			}
			newboxs[now] = 1
		}
		boxs = newboxs
	}
	return ans
}
