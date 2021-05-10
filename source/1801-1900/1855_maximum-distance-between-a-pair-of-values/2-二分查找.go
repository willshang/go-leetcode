package main

import "sort"

func main() {

}

func maxDistance(nums1 []int, nums2 []int) int {
	res := 0
	for j := 0; j < len(nums2); j++ {
		n := min(j, len(nums1))
		i := sort.Search(n, func(i int) bool {
			return nums1[i] <= nums2[j]
		})
		if i < n {
			if j-i > res {
				res = j - i
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
