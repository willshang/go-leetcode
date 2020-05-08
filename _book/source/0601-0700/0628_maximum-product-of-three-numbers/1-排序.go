package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-1, 2, 3, 4, 5, -1}
	fmt.Println(maximumProduct(arr))

	fmt.Println(maximumProduct([]int{-1, -1, 4, -4, -1}))
	fmt.Println(maximumProduct([]int{-1, 1, 2}))
}

// leetcode628_三个数的最大乘积
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	return max(nums[0]*nums[1]*nums[len(nums)-1],
		nums[len(nums)-3]*nums[len(nums)-2]*nums[len(nums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
