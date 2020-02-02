package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	merge(a, 3, []int{}, 0)
	fmt.Println(a)
}

// leetcode 88 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		if nums1[m-1] < nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
	if m == 0 && n > 0 {
		for n > 0 {
			nums1[n-1] = nums2[n-1]
			n--
		}
	}
}
