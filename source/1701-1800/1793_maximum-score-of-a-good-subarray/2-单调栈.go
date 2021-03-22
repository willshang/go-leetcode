package main

import "fmt"

func main() {
	fmt.Println(maximumScore([]int{1, 4, 3, 7, 4, 5}, 3))
}

func maximumScore(nums []int, k int) int {
	res := 0
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < n; i++ {
		if left[i] < k && right[i] > k {
			res = max(res, nums[i]*(right[i]-left[i]-1))
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
