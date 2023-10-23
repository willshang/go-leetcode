package main

import "math"

func main() {

}

// leetcode1283_使结果不超过阈值的最小除数
func smallestDivisor(nums []int, threshold int) int {
	left, right := 1, math.MaxInt32/10
	res := 0
	for left <= right {
		mid := left + (right-left)/2
		count := getValue(nums, mid)
		if count > threshold {
			left = mid + 1
		} else {
			right = mid - 1
			res = mid
		}
	}
	return res
}

func getValue(nums []int, target int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res = res + (nums[i]+target-1)/target // 向上取整
	}
	return res
}
