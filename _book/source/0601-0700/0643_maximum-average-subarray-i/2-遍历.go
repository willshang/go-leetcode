package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(arr, k))
}

func findMaxAverage(nums []int, k int) float64 {
	max := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if i+k > len(nums) {
			break
		}
		sum := 0
		for j := i; j < i+k; j++ {
			sum = sum + nums[j]
		}
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}
