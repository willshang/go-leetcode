package main

import "fmt"

func main() {
	fmt.Println(maxResult([]int{1, -1, -2, 4, -7, 3}, 2))
}


// leetcode1696_跳跃游戏VI
func maxResult(nums []int, k int) int {
	n := len(nums)
	if k > n {
		k = n
	}
	res := nums[0]
	stack := make([][2]int, 0)
	stack = append(stack, [2]int{0, nums[0]})
	for i := 1; i < len(nums); i++ {
		if stack[0][0] < i-k {
			stack = stack[1:]
		}
		res = stack[0][1] + nums[i]
		for len(stack) > 0 && stack[len(stack)-1][1] < res {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, [2]int{i, res})
	}
	return res
}
