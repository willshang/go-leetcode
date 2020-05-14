package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(arr, k))
}

func findMaxAverage(nums []int, k int) float64 {
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	max := sum[k-1]
	for i := k; i < len(nums); i++ {
		if sum[i]-sum[i-k] > max {
			max = sum[i] - sum[i-k]
		}
	}
	return float64(max) / float64(k)
}
