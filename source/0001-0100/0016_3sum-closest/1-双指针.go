package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return target
			}
			if abs(sum, target) < abs(res, target) {
				res = sum
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
