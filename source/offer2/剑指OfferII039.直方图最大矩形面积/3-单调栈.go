package main

import "fmt"

func main() {
	fmt.Println(largestRectangleArea([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(largestRectangleArea([]int{1}))
}

// 剑指OfferII039.直方图最大矩形面积
func largestRectangleArea(heights []int) int {
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	n := len(heights)
	res := 0
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		// 递增栈
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
