package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(increasingTriplet([]int{10, 7, 11, 8, 9, 10}))
	fmt.Println(increasingTriplet([]int{3, 4, 2, 6}))
}

// leetcode334_递增的三元子序列
func increasingTriplet(nums []int) bool {
	a, b := math.MaxInt32, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if a >= nums[i] {
			a = nums[i]
		} else if b >= nums[i] {
			b = nums[i]
		} else {
			return true
		}
	}
	return false
}
