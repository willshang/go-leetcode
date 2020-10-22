package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

func canJump(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if i <= max {
			if i+nums[i] > max {
				max = i + nums[i]
			}
			if max >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}
