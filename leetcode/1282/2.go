package main

func main() {

}

func groupThePeople(groupSizes []int) [][]int {
	var m = make(map[int][]int)
	var ret [][]int
	for id, size := range groupSizes {
		m[size] = append(m[size], id)
		if len(m[size]) == size {
			ret = append(ret, m[size])
			m[size] = []int{}
		}
	}
	return ret
}
