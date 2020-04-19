package context182

func findLucky(arr []int) int {
	var f [501]int
	for _, x := range arr {
		f[x]++
	}
	for i := 500; i > 0; i-- {
		if f[i] == i {
			return i
		}
	}
	return -1
}
