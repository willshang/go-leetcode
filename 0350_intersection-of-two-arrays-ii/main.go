package main

import "fmt"

func main() {
	nums1 := []int{4,9,5}
	nums2 := []int{9,4,9,8,4}

	fmt.Println(intersect(nums1,nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	res := []int{}
	for _, v := range nums1{
		m1[v] += 1
	}

	for _, v := range nums2{
		if m1[v] > 0{
			res = append(res,v)
			m1[v]--
		}
	}
	return res
}