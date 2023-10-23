package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
}

// leetcode1480_一维数组的动态和
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	return nums
}
