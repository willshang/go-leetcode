package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

// leetcode42_接雨水
func trap(height []int) int {
	res := 0
	if len(height) == 0 {
		return 0
	}
	left := 0
	right := len(height) - 1
	leftMax := 0  // 左边的最大值
	rightMax := 0 // 右边的最大值
	for left < right {
		// 当前坐标形成的面积=(min(左边最高，右边最高)-当前高度) * 宽度(1,可省略)
		// 选择高度低的一边处理并求最大值, 说明当前侧最大值小于另一侧
		if height[left] < height[right] {
			// 也可以写成这样
			// leftMax = max(leftMax, height[left])
			// res = res + leftMax - height[left]
			if height[left] >= leftMax { // 递增无法蓄水
				leftMax = height[left]
			} else {
				res = res + leftMax - height[left]
			}
			left++
		} else {
			// 也可以写成这样
			// rightMax = max(rightMax, height[right])
			// res = res + rightMax - height[right]
			if height[right] >= rightMax { // 递减无法蓄水
				rightMax = height[right]
			} else {
				res = res + rightMax - height[right]
			}
			right--
		}
	}
	return res
}
