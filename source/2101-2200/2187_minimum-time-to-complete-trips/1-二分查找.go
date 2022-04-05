package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumTime([]int{1, 2, 3}, 5))
}

// leetcode2187_完成旅途的最少时间
func minimumTime(time []int, totalTrips int) int64 {
	n := len(time)
	maxValue := 0
	for i := 0; i < n; i++ {
		maxValue = max(maxValue, time[i])
	}
	left := int64(1)
	right := int64(totalTrips) * int64(maxValue)
	for left < right {
		mid := left + (right-left)/2
		if check(time, mid) >= totalTrips {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(arr []int, per int64) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		count = count + int(per/int64(arr[i]))
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
