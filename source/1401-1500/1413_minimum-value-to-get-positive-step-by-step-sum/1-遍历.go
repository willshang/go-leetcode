package main

import "fmt"

func main() {
	fmt.Println(minStartValue([]int{-3, 2, -3, 4, 2}))
	fmt.Println(minStartValue([]int{1, 2}))
}

// leetcode1413_逐步求和得到正数的最小值
func minStartValue(nums []int) int {
	min := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum = sum + nums[i]
		if sum < min {
			min = sum
		}
	}
	if min >= 0 {
		return 1
	}
	return 1 - min
}
