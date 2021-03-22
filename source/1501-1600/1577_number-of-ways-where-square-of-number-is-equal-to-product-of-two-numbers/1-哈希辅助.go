package main

import "fmt"

func main() {
	fmt.Println(numTriplets([]int{7, 4}, []int{5, 2, 8, 9}))
}

// leetcode1577_数的平方等于两数乘积的方法数
func numTriplets(nums1 []int, nums2 []int) int {
	res := 0
	res = res + getCount(nums1, nums2)
	res = res + getCount(nums2, nums1)
	return res
}

func getCount(nums1 []int, nums2 []int) int {
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]*nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		for j := i + 1; j < len(nums2); j++ {
			res = res + m[nums2[i]*nums2[j]]
		}
	}
	return res
}
