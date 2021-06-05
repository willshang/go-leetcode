package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(nextGreaterElement(12))
}

func nextGreaterElement(n int) int {
	arr := make([]int, 0)
	for n > 0 {
		arr = append(arr, n%10)
		n = n / 10
	}
	reverse(arr, 0, len(arr)-1)
	arr = nextPermutation(arr)
	if arr == nil {
		return -1
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		res = res*10 + arr[i]
		if res > math.MaxInt32 {
			return -1
		}
	}
	return res
}

// leetcode31.下一个排列
func nextPermutation(nums []int) []int {
	n := len(nums)
	left := n - 2
	// 以12385764为例，从后往前找到5<7 的升序情况，目标值为左边的数5
	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}
	if left >= 0 { // 存在升序的情况
		right := n - 1
		// 从后往前，找到第一个大于目标值的数，如6>5，然后交换
		for right >= 0 && nums[right] <= nums[left] {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	} else {
		return nil
	}
	reverse(nums, left+1, n-1)
	return nums
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
