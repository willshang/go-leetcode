package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

// leetcode11_盛最多水的容器
func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	res := 0
	for i < j {
		area := (j - i) * min(height[i], height[j])
		if area > res {
			res = area
		}
		// 移动较小的指针，尝试获取更大的面积
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
