package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minTime([]int{1, 2, 3, 3}, 2))
}

func minTime(time []int, m int) int {
	left, right, mid := 0, 0, 0
	for i := 0; i < len(time); i++ {
		right = right + time[i]
	}
	// 二分查找一个数mid，使time数组能分割成m个和不小于mid的子数组
	res := math.MaxInt32
	for left <= right {
		mid = left + (right-left)/2
		if check(time, mid) <= m {
			if mid < res {
				res = mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func check(arr []int, mid int) int {
	res := 1
	maxValue := 0
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
		if sum-maxValue > mid {
			sum = arr[i]
			maxValue = arr[i]
			res++
		}
	}
	return res
}
