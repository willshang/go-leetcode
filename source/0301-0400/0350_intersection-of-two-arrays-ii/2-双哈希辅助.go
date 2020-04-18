package main

import (
	"fmt"
)

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	res := make([]int, 0)
	for _, v := range nums1 {
		m1[v]++
	}

	for _, v := range nums2 {
		if m1[v] != 0 && m1[v] > m2[v] {
			m2[v]++
		}
	}

	for k := range m2 {
		for i := 0; i < m2[k]; i++ {
			res = append(res, k)
		}
	}
	return res
}
