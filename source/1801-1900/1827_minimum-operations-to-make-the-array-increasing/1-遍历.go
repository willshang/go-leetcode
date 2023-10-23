package main

import "fmt"

func main() {
	fmt.Println(minOperations([]int{1, 1, 1}))
}

// leetcode1827_最少操作使数组递增
func minOperations(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	res := 0
	for i := 1; i < n; i++ {
		if nums[i-1] >= nums[i] {
			res = res + nums[i-1] + 1 - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}
	return res
}
