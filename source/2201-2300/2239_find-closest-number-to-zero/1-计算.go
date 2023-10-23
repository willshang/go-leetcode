package main

import "math"

func main() {

}

// leetcode2239_找到最接近0的数字
func findClosestNumber(nums []int) int {
	res := math.MaxInt32 / 10
	for i := 0; i < len(nums); i++ {
		if abs(nums[i]) < abs(res) {
			res = nums[i]
		} else if abs(nums[i]) == abs(res) {
			res = max(res, nums[i])
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
