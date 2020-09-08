package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumGap([]int{3, 6, 9, 1}))
}

func maximumGap(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	res := 0
	minValue, maxValue := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		minValue = min(minValue, nums[i])
		maxValue = max(maxValue, nums[i])
	}
	bucketSize := (maxValue-minValue)/len(nums) + 1
	bucketNum := (maxValue-minValue)/bucketSize + 1
	arr := make([][]int, bucketNum)
	for i := 0; i < len(nums); i++ {
		index := (nums[i] - minValue) / bucketSize
		if len(arr[index]) == 0 {
			arr[index] = make([]int, 2)
			arr[index][0], arr[index][1] = nums[i], nums[i]
		} else {
			arr[index][0] = min(arr[index][0], nums[i])
			arr[index][1] = max(arr[index][1], nums[i])
		}
	}
	prev := 0
	for i := 0; i < bucketNum; i++ {
		if len(arr[i]) == 0 {
			continue
		}
		res = max(res, arr[i][0]-arr[prev][1])
		prev = i
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
