package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	res := make([]int, 0)
	for _, v := range nums1 {
		m1[v] = true
	}

	for _, v := range nums2 {
		if m1[v] != false {
			m2[v] = true
		}
	}

	for k := range m2 {
		res = append(res, k)
	}
	return res
}
