package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
}

// leetcode1464_数组中两元素的最大乘积
func maxProduct(nums []int) int {
	sort.Ints(nums)
	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}
