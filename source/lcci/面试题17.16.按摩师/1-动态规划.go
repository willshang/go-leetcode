package main

import "fmt"

func main() {
	fmt.Println(massage([]int{1, 2, 3, 1}))
}

// 程序员面试金典17.16_按摩师
func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	a := nums[0]
	b := max(a, nums[1])

	for i := 2; i < len(nums); i++ {
		a, b = b, max(a+nums[i], b)
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
