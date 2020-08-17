package main

import "fmt"

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	res := make([]int, len(nums1))
	for i := 0; i < len(nums2); i++ {
		for j := i + 1; j < len(nums2); j++ {
			if nums2[j] > nums2[i] {
				m[nums2[i]] = nums2[j]
				break
			}
		}
	}
	for key, value := range nums1 {
		if _, ok := m[value]; ok {
			res[key] = m[value]
		} else {
			res[key] = -1
		}
	}
	return res
}
