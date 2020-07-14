package main

import "fmt"

func main() {
	//fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{2, 0, 1, 0, 1}))
}

func canJump(nums []int) bool {
	zero := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if zero > 0 {
			if i+nums[i] > zero {
				zero = -1
			}
			continue
		}
		if nums[i] == 0 {
			zero = i
			continue
		}
	}
	return zero < 0
}
