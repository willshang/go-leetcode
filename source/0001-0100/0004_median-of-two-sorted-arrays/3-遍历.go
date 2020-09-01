package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

// leetcode4_寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	a, b := total/2, (total-1)/2
	count := 0
	res := 0
	for i, j := 0, 0; i < len(nums1) || j < len(nums2); count++ {
		if i < len(nums1) && (j == len(nums2) || nums1[i] < nums2[j]) {
			if count == a {
				res = res + nums1[i]
			}
			if count == b {
				res = res + nums1[i]
			}
			i++
		} else {
			if count == a {
				res = res + nums2[j]
			}
			if count == b {
				res = res + nums2[j]
			}
			j++
		}
	}
	return float64(res) / 2
}
