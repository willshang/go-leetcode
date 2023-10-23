package main

import "sort"

func main() {

}

// leetcode462_最少移动次数使数组元素相等II
func minMoves2(nums []int) int {
	sort.Ints(nums)
	res := 0
	left, right := 0, len(nums)-1
	for left < right {
		// 无论选哪个数作为最终值，前后第n个值需要移动之和不变
		// nums[right]-target+target-nums[left] = nums[right] - nums[left]
		res = res + nums[right] - nums[left]
		left++
		right--
	}
	return res
}
