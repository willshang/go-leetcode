package main

import "fmt"

func main() {
	nums1 := []int{4, 1, 3, 6}
	nums2 := []int{1, 2, 7, 4}
	fmt.Println(nextGreaterElement(nums1, nums2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for i, n := range nums2 {
		m[n] = i
	}
	res := make([]int, len(nums1))
	for i, n := range nums1 {
		res[i] = -1
		for j := m[n] + 1; j < len(nums2); j++ {
			if n < nums2[j] {
				res[i] = nums2[j]
				break
			}
		}
	}
	return res
}
