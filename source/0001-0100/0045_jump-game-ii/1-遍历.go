package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	last := len(nums) - 1
	res := 0
	for last > 0 {
		// 从前往后，找到第一个一步能走到终点的，更新终点的位置
		for i := 0; i < last; i++ {
			if i+nums[i] >= last {
				last = i
				res++
				break
			}
		}
	}
	return res
}
