package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	leftNum := (len(nums1)+len(nums2))/2 + ((len(nums1) + len(nums2)) & 1)
	l, r, mid := leftNum-len(nums2)-1, min(leftNum, len(nums1))-1, 0
	l1, r1, l2, r2 := 0, 0, 0, 0
	for l <= r {
		mid = (l + r) >> 1
		//fmt.Println(l, r, mid+1)
		l1, r1 = get(nums1, mid+1)
		l2, r2 = get(nums2, leftNum-mid-1)
		if max(l1, l2) <= min(r1, r2) {
			break
		}
		if l1 > r2 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	fmt.Println(l1, l2, r1, r2)
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(max(l1, l2))
	}
	return (float64(max(l1, l2)) + float64(min(r1, r2))) / 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func get(arr []int, num int) (int, int) {
	l, r := 0, 0
	if num-1 < 0 {
		l = -999999999
	} else {
		l = arr[num-1]
	}
	if num >= len(arr) {
		r = 999999999
	} else {
		fmt.Println(num)
		r = arr[num]
	}
	return l, r

}
func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{1, 1}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 5, 6, 7, 8}, []int{3, 4}))
}
