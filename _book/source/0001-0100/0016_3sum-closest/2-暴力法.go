package main

import (
	"fmt"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

// leetcode16_最接近的三数之和
func threeSumClosest(nums []int, target int) int {
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum, target) < abs(res, target) {
					res = sum
				}
			}
		}
	}
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
