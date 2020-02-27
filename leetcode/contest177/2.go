package contest177

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	ru := make([]int, n)
	num := 0

	for i := 0; i < n; i++ {
		for _, now := range []int{leftChild[i], rightChild[i]} {
			if now < 0 {
				continue
			}
			ru[now]++
			if ru[now] > 1 {
				return false
			}
			num++
		}
	}
	return num == n-1
}
