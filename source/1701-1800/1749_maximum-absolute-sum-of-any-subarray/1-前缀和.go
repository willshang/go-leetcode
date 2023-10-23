package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxAbsoluteSum([]int{1, -3, 2, 3, -4}))
}

// leetcode1749_任意子数组和的绝对值的最大值
func maxAbsoluteSum(nums []int) int {
	arr := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		arr[i] = arr[i-1] + nums[i-1]
	}
	sort.Ints(arr)
	return arr[len(nums)] - arr[0]
}
