package main

import "fmt"

func main() {
	fmt.Println(maxNumber([]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5))
}

// leetcode321_拼接最大数
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	res := make([]int, k)
	for i := 0; i <= k; i++ {
		if i <= len(nums1) && k-i <= len(nums2) {
			a := check(nums1, i)
			b := check(nums2, k-i)
			c := merge(a, b)
			if compare(res, c) == true {
				res = c
			}
		}
	}
	return res
}

func check(num []int, k int) []int {
	stack := make([]int, 0)
	count := len(num) - k
	for i := 0; i < len(num); i++ {
		for len(stack) > 0 && count > 0 && stack[len(stack)-1] < num[i] {
			stack = stack[:len(stack)-1]
			count--
		}
		stack = append(stack, num[i])
	}
	return stack[:k]
}

func merge(nums1, nums2 []int) []int {
	res := make([]int, len(nums1)+len(nums2))
	if len(nums1) == 0 {
		copy(res, nums2)
		return res
	}
	if len(nums2) == 0 {
		copy(res, nums1)
		return res
	}
	for i := 0; i < len(res); i++ {
		if compare(nums1, nums2) {
			res[i], nums2 = nums2[0], nums2[1:]
		} else {
			res[i], nums1 = nums1[0], nums1[1:]
		}
	}
	return res
}

func compare(nums1, nums2 []int) bool {
	for i := 0; i < len(nums1) && i < len(nums2); i++ {
		if nums1[i] != nums2[i] {
			return nums1[i] < nums2[i]
		}
	}
	return len(nums1) < len(nums2)
}
