package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 2, 3, 1}
	sortArray(arr)
	fmt.Println(arr)
}

func sortArray(nums []int) []int {
	buildHeap(nums, len(nums))
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapModify(nums, 0, i-1)
	}
	return nums
}

func buildHeap(arr []int, length int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapModify(arr, i, length-1)
	}
}

func heapModify(arr []int, start, end int) {
	temp := arr[start]
	for left := 2*start + 1; left <= end; left = 2*left + 1 {
		if left < end && arr[left] < arr[left+1] {
			left++
		}
		if arr[start] < arr[left] {
			arr[start] = arr[left]
			start = left
		}
		arr[start] = temp
	}
}
