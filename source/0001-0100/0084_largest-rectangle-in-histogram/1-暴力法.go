package main

import "fmt"

func main() {
	// fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{1}))
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	res := 0
	for i := 0; i < n; i++ {
		height := heights[i]
		for j := i; j < n; j++ {
			width := j - i + 1
			height = min(height, heights[j])
			res = max(res, width*height)
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
