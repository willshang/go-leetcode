package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0, 1, 3}))
}

func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
