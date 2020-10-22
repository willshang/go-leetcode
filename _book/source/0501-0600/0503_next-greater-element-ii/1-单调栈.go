package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}

// leetcode503_下一个更大元素II
func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	if len(nums) == 0 {
		return res
	}
	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}
	stack := make([]int, 0)
	for i := 0; i < len(nums)*2; i++ {
		index := i % len(nums)
		for len(stack) > 0 && nums[index] > nums[stack[len(stack)-1]] {
			if res[stack[len(stack)-1]] == -1 {
				res[stack[len(stack)-1]] = nums[index]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, index)
	}
	return res
}
