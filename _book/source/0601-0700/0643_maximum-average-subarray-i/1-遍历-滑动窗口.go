package main

import "fmt"

func main() {
	arr := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(arr, k))
}

// leetcode643_子数组最大平均数 I
func findMaxAverage(nums []int, k int) float64 {
	temp := 0
	for i := 0; i < k; i++ {
		temp = temp + nums[i]
	}
	max := temp
	for i := k; i < len(nums); i++ {
		temp = temp + nums[i] - nums[i-k]
		if max < temp {
			max = temp
		}
	}
	return float64(max) / float64(k)
}
