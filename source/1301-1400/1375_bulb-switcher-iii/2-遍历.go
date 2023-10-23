package main

import "fmt"

func main() {
	fmt.Println(numTimesAllBlue([]int{4, 1, 2, 3}))
}

// leetcode1375_灯泡开关III
func numTimesAllBlue(light []int) int {
	res := 0
	maxValue := 0
	for i := 0; i < len(light); i++ {
		// 最大亮起来的灯等于前面灯的数量，那么说明前面灯都亮了
		maxValue = max(maxValue, light[i])
		if maxValue == i+1 {
			res++
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
