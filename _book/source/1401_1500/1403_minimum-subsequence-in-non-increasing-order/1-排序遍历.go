package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minSubsequence([]int{4, 3, 10, 9, 8}))
	fmt.Println(minSubsequence([]int{4, 4, 7, 6, 7}))
}

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	target := sum / 2
	sum = 0
	res := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if sum <= target {
			res = append(res, nums[i])
			sum = sum + nums[i]
		}
	}
	return res
}
