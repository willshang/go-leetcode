package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total%2 == 1 {
		mid := total / 2
		return float64(getKth(nums1, nums2, mid+1))
	}
	mid1, mid2 := total/2-1, total/2
	return float64(getKth(nums1, nums2, mid1+1)+getKth(nums1, nums2, mid2+1)) / 2.0
}

func getKth(nums1 []int, nums2 []int, k int) int {
	a, b := 0, 0
	for {
		if a == len(nums1) {
			return nums2[b+k-1]
		}
		if b == len(nums2) {
			return nums1[a+k-1]
		}
		if k == 1 {
			return min(nums1[a], nums2[b])
		}
		mid := k / 2
		newA := min(a+mid, len(nums1)) - 1
		newB := min(b+mid, len(nums2)) - 1
		valueA, valueB := nums1[newA], nums2[newB]
		if valueA < valueB {
			k = k - (newA - a + 1)
			a = newA + 1
		} else {
			k = k - (newB - b + 1)
			b = newB + 1
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
