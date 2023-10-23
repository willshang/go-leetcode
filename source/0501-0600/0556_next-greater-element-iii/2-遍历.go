package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(nextGreaterElement(12))
}

// leetcode556_下一个更大元素III
func nextGreaterElement(n int) int {
	nums := []byte(strconv.Itoa(n))
	length := len(nums)
	left := length - 2
	// 以12385764为例，从后往前找到5<7 的升序情况，目标值为左边的数5
	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}
	if left >= 0 { // 存在升序的情况
		right := length - 1
		// 从后往前，找到第一个大于目标值的数，如6>5，然后交换
		for right >= 0 && nums[right] <= nums[left] {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	} else {
		return -1
	}
	left = left + 1
	right := length - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	res, _ := strconv.Atoi(string(nums))
	if res > math.MaxInt32 {
		return -1
	}
	return res
}
