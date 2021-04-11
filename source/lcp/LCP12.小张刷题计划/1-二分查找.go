package main

import "fmt"

func main() {
	fmt.Println(minTime([]int{1, 2, 3, 3}, 2))
}

// leetcode_lcp12_小张刷题计划
func minTime(time []int, m int) int {
	left, right, mid := 0, 0, 0
	for i := 0; i < len(time); i++ {
		right = right + time[i]
	}
	// 二分查找一个数mid，使time数组能分割成m个和不小于mid的子数组
	for left <= right {
		mid = left + (right-left)/2
		if check(time, mid, m) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(arr []int, mid, m int) bool {
	maxValue := 0
	sum := 0
	count := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
		if sum-maxValue > mid {
			count++
			if count >= m {
				return false
			}
			sum = arr[i]
			maxValue = arr[i]
		}
	}
	return true
}
