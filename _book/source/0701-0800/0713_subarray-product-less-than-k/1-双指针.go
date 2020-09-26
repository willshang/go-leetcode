package main

import "fmt"

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}

// leetcode713_乘积小于K的子数组
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	res := 0
	left := 0
	total := 1
	for right := 0; right < len(nums); right++ {
		total = total * nums[right]
		for k <= total {
			total = total / nums[left]
			left++
		}
		res = res + right - left + 1
	}
	return res
}
