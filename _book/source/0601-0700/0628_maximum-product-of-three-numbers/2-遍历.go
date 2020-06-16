package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-1, 2, 3, 4, 5, -1}
	fmt.Println(maximumProduct(arr))
	fmt.Println(maximumProduct([]int{-1, -1, 4, -4, -1}))
	fmt.Println(maximumProduct([]int{-1, 1, 2}))
}

func maximumProduct(nums []int) int {
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	min1, min2 := math.MaxInt32, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] <= min1 {
			min2 = min1
			min1 = nums[i]
		} else if nums[i] <= min2 {
			min2 = nums[i]
		}
		if nums[i] >= max1 {
			max3 = max2
			max2 = max1
			max1 = nums[i]
		} else if nums[i] >= max2 {
			max3 = max2
			max2 = nums[i]
		} else if nums[i] >= max3 {
			max3 = nums[i]
		}
	}
	return max(min1*min2*max1, max1*max2*max3)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
