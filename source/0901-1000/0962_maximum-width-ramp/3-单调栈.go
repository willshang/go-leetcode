package main

import "fmt"

func main() {
	//fmt.Println(maxWidthRamp([]int{7, 5, 1, 2, 3, 4, 5, 10, 8, 7, 6, 5}))
	//fmt.Println(maxWidthRamp([]int{10, 10, 10, 7, 1, 6, 2, 1, 7}))
	fmt.Println(maxWidthRamp([]int{6, 0, 8, 2, 1, 5}))
}

// leetcode962_最大宽度坡
func maxWidthRamp(A []int) int {
	res := 0
	n := len(A)
	stack := make([]int, 0) // 递减栈
	for i := 0; i < n; i++ {
		if len(stack) == 0 || A[stack[len(stack)-1]] > A[i] {
			stack = append(stack, i)
		}
	}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && A[stack[len(stack)-1]] <= A[i] {
			index := stack[len(stack)-1] // 坡底index能形成的最大宽度
			stack = stack[:len(stack)-1]
			res = max(res, i-index)
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
