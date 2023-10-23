package main

import "sort"

func main() {

}

// leetcode1913_两个数对之间的最大乘积差
func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return nums[n-1]*nums[n-2] - nums[0]*nums[1]
}
