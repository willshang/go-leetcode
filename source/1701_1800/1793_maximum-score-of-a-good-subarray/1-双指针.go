package main

import "fmt"

func main() {
	fmt.Println(maximumScore([]int{1, 4, 3, 7, 4, 5}, 3))
}

// leetcode1793_好子数组的最大分数
func maximumScore(nums []int, k int) int {
	left, right := k, k
	res := 0
	for minValue := nums[k]; minValue >= 1; minValue-- {
		for left >= 0 && nums[left] >= minValue {
			left--
		}
		for right < len(nums) && nums[right] >= minValue {
			right++
		}
		left++  // left注意+1
		right-- // right注意-1
		res = max(res, minValue*(right-left+1))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
