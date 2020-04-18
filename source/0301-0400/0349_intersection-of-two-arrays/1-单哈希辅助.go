package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}

// leetcode349_两个数组的交集
func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v] = 1
	}
	for _, v := range nums2 {
		if m[v] == 1 {
			res = append(res, v)
			m[v] += 1
		}
	}
	return res
}
