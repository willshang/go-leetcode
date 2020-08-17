package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 2, 3, 1}
	fmt.Println(sortArray(arr))
	fmt.Println(arr)
}

func sortArray(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	return merge(sortArray(nums[:len(nums)/2]), sortArray(nums[len(nums)/2:]))
}

func merge(left, right []int) []int {
	res := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		res = append(res, left...)
	}
	if len(right) > 0 {
		res = append(res, right...)
	}
	return res
}
