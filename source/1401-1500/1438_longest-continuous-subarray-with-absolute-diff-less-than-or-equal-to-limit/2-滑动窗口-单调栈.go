package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestSubarray([]int{8, 2, 9, 7}, 4))
}

func longestSubarray(nums []int, limit int) int {
	res := 0
	maxStack, minStack := make([]int, 0), make([]int, 0)
	var i int
	for j := 0; j < len(nums); j++ {
		for len(minStack) > 0 && minStack[len(minStack)-1] > nums[j] {
			minStack = minStack[:len(minStack)-1]
		}
		minStack = append(minStack, nums[j])
		for len(maxStack) > 0 && maxStack[len(maxStack)-1] < nums[j] {
			maxStack = maxStack[:len(maxStack)-1]
		}
		maxStack = append(maxStack, nums[j])
		for maxStack[0]-minStack[0] > limit { // 移除左边
			if nums[i] == maxStack[0] {
				maxStack = maxStack[1:]
			}
			if nums[i] == minStack[0] {
				minStack = minStack[1:]
			}
			i++
		}
		res = max(res, j-i+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
