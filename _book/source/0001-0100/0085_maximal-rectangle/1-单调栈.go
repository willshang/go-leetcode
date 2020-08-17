package main

import (
	"fmt"
)

func main() {
	arr := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalRectangle(arr))
}

// leetcode85_最大矩形
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	res := 0
	n, m := len(matrix), len(matrix[0])
	height := make([]int, m) // 高度
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' {
				height[j] = 0
			} else {
				height[j] = height[j] + 1
			}
		}
		res = max(res, getMaxArea(height))
	}
	return res
}

func getMaxArea(heights []int) int {
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	n := len(heights)
	res := 0
	stack := make([]int, 0) // 递增栈
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i - stack[len(stack)-1] - 1
			res = max(res, height*width)
		}
		stack = append(stack, i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
