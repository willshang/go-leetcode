package main

import (
	"fmt"
)

func main() {
	fmt.Println(smallestK([]int{1, 3, 5, 7, 2, 4, 6, 8}, 4))
}

// 程序员面试金典17.14_最小K个数
func smallestK(arr []int, k int) []int {
	return quickSort(arr, 0, len(arr)-1, k)
}

func quickSort(arr []int, left, right, k int) []int {
	if left > right {
		return nil
	}
	index := partition(arr, left, right)
	if index == k {
		return arr[:k]
	} else if index < k {
		return quickSort(arr, index+1, right, k)
	}
	return quickSort(arr, left, index-1, k)
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
