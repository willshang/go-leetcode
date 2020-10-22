package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	if len(nums) == 0 {
		return res
	}
	stack := make([]int, 0)
	for i := 2*len(nums) - 1; i >= 0; i-- {
		index := i % len(nums)
		for len(stack) > 0 && nums[index] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[index] = -1
		} else {
			res[index] = stack[len(stack)-1]
		}
		stack = append(stack, nums[index])
	}
	return res
}
