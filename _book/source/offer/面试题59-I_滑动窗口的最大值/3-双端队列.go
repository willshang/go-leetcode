package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{}, 0))
}

// 剑指offer_面试题59-I_滑动窗口的最大值
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	deque := make([]int, 0)
	for i := 0; i < k; i++ {
		for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	for i := k; i < len(nums); i++ {
		res = append(res, nums[deque[0]])
		for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		deque = append(deque, i)
	}
	res = append(res, nums[deque[0]])
	return res
}
