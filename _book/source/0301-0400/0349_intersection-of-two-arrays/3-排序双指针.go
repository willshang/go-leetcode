package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums1 := []int{4, 9, 5}
	//nums2 := []int{9, 4, 9, 8, 4}
	//fmt.Println(intersection(nums1, nums2))

	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := make([]int, 0)
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			if len(res) == 0 || res[len(res)-1] != nums1[i] {
				res = append(res, nums1[i])
			}
			i++
			j++
		}
	}
	return res
}
