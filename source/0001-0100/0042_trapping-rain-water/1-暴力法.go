package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	res := 0
	for i := 0; i < len(height); i++ {
		left, right := 0, 0
		for j := i; j >= 0; j-- {
			left = max(left, height[j])
		}
		for j := i; j < len(height); j++ {
			right = max(right, height[j])
		}
		// 当前坐标形成的面积=(min(左边最高，右边最高)-当前高度) * 宽度(1,可省略)
		area := min(left, right) - height[i]
		res = res + area
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
