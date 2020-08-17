package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	res := 0
	stack := make([]int, 0)
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			bottom := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				prev := stack[len(stack)-1]
				// 横着的面积=长(min(height[i], height[prev])-bottom)*宽(i-prev-1)
				h := min(height[i], height[prev]) - bottom
				w := i - prev - 1
				area := h * w
				res = res + area
			}
		}
		stack = append(stack, i)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
