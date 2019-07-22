package main

import "fmt"

func mctFromLeafValues(arr []int) int {
	var dp [101][101]int
	for wide := 2; wide <= len(arr); wide++ {
		for i := 0; i+wide <= len(arr); i++ {
			j := wide + i

			dp[i][j] = 1999999999
			for mid := i + 1; mid < j; mid++ {
				tmp := getMaxNum(arr[i:mid])*getMaxNum(arr[mid:j]) + dp[i][mid] + dp[mid][j]
				if tmp < dp[i][j] {
					dp[i][j] = tmp
				}
			}
		}
	}
	return dp[0][len(arr)]
}

func getMaxNum(arr []int) int {
	ans := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > ans {
			ans = arr[i]
		}
	}
	return ans
}

func mctFromLeafValues1(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	ans := -1
	for i := 1; i < len(arr); i++ {
		l := arr[0:i]
		r := arr[i:]
		temp := getMaxNum(l)*getMaxNum(r) + mctFromLeafValues(l) + mctFromLeafValues(r)
		if ans == -1 {
			ans = temp
		} else if ans > temp {
			ans = temp
		}
	}
	return ans
}

func main() {
	fmt.Println(mctFromLeafValues([]int{6, 2, 4}))
	//fmt.Println(mctFromLeafValues([]int{7, 5, 12, 2, 2, 3, 13, 8, 4, 9, 12, 9, 3, 10, 4, 13, 7, 5, 15}))
}
