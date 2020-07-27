package main

import "fmt"

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}

// leetcode496_下一个更大元素I
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	res := make([]int, len(nums1))
	stack := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if len(stack) > 0 {
			for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
				top := stack[len(stack)-1]
				m[top] = nums2[i]
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, nums2[i])
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
