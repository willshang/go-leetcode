package main

import "fmt"

func main() {
	nums := []int{6, 0, 1, 0, 3, 12, 0, 7}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	length := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[length] = nums[i]
			length++
		}
	}

	for i := length; i < len(nums); i++ {
		nums[i] = 0
	}
}
