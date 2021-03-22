package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
}

func maxProduct(nums []int) int {
	max := math.MinInt32
	next := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			next, max = max, nums[i]
		} else if nums[i] > next {
			next = nums[i]
		}
	}
	return (max - 1) * (next - 1)
}
