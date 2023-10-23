package main

import "fmt"

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		// 非正数处理
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		value := abs(nums[i])
		// 标负
		if value <= n {
			nums[value-1] = -abs(abs(nums[value-1]))
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
