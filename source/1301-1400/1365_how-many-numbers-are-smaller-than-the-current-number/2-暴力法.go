package main

import "fmt"

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
}

// leetcode1365_有多少小于当前数字的数字
func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				res[i]++
			}
		}
	}
	return res
}
