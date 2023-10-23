package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1)%2 == 1 {
		return float64(nums1[len(nums1)/2])
	}
	return float64(nums1[len(nums1)/2]+nums1[len(nums1)/2-1]) / 2
}
