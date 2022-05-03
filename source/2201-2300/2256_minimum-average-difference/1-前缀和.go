package main

import "math"

func main() {

}

// leetcode2256_最小平均差
func minimumAverageDifference(nums []int) int {
	res := 0
	right := 0
	for _, v := range nums {
		right = right + v
	}
	left := 0
	minValue := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		left = left + nums[i]
		right = right - nums[i]
		var a, b int
		a = left / (i + 1)
		if i != len(nums)-1 {
			b = right / (len(nums) - i - 1)
		}
		if abs(a-b) < minValue { // 更新结果
			minValue = abs(a - b)
			res = i
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
