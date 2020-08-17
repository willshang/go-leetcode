package main

import "fmt"

func main() {
	nums := []int{6, 0, 1, 0, 3, 12, 0, 7}
	moveZeroes(nums)
	fmt.Println(nums)
}

// leetcode283_移动零
func moveZeroes(nums []int) {
	length := 0
	for i := 0; i < len(nums); i++ {
		nums[i], nums[length] = nums[length], nums[i]
		if nums[length] != 0 {
			length++
		}
	}
}
