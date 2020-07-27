package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	index1, index2 := n/2, (n-1)/2
	count := 0
	res := float64(0)
	for i, j := 0, 0; i < len(nums1) || j < len(nums2); count++ {
		fmt.Println(i, j, count)
		if i < len(nums1) && j < len(nums2) && (nums1[i] < nums2[j]) {
			i++
		} else {
			j++
		}
		if count == index1 || count == index2 {
			if i < len(nums1) {
				res = res + float64(nums1[i])
			} else {
				res = res + float64(nums2[j])
			}
		}
	}
	return res / 2
}
