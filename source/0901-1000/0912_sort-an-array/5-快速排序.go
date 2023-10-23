package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 2, 3, 1}
	sortArray(arr)
	fmt.Println(arr)
}

// leetcode912_排序数组
func sortArray(nums []int) []int {
	quick(nums, 0, len(nums)-1)
	return nums
}

func quick(arr []int, left, right int) {
	if left >= right {
		return
	}
	index := partition(arr, left, right)
	quick(arr, left, index-1)
	quick(arr, index+1, right)
}

func partition(arr []int, left, right int) int {
	baseValue := arr[left] // 基准值
	for left < right {
		for baseValue <= arr[right] && left < right {
			right-- // 依次查找大于基准值的位置
		}
		arr[left] = arr[right]
		for arr[left] <= baseValue && left < right {
			left++ // 依次查找小于基准值的位置
		}
		arr[right] = arr[left]
	}
	arr[right] = baseValue
	return right
}
