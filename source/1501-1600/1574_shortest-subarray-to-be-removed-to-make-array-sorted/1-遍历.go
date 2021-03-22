package main

import (
	"fmt"
)

func main() {
	fmt.Println(findLengthOfShortestSubarray([]int{1}))
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3}))
	fmt.Println(findLengthOfShortestSubarray([]int{5, 4, 3, 2, 1}))
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5}))
	fmt.Println(findLengthOfShortestSubarray([]int{2, 2, 2, 1, 1, 1}))
	fmt.Println(findLengthOfShortestSubarray([]int{13, 0, 14, 7, 18, 18, 18, 16, 8, 15, 20}))
}

// leetcode1574_删除最短的子数组使剩余数组有序
func findLengthOfShortestSubarray(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	flag := true
	left := 0
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			flag = false
			break
		} else {
			left = i
		}
	}
	if flag == true {
		return 0
	}
	right := len(arr) - 1
	for i := len(arr) - 1; i >= 1; i-- {
		if arr[i-1] <= arr[i] {
			right = i - 1
		} else {
			break
		}
	}
	leftC, rightC := 0, 0
	for i := left; i >= 0 && arr[i] > arr[right]; i-- {
		leftC++
	}
	for i := right; i < len(arr) && arr[i] < arr[left]; i++ {
		rightC++
	}
	res := 0
	res = max(res, (left+1)+(len(arr)-right)-leftC)
	res = max(res, (left+1)+(len(arr)-right)-rightC)
	return len(arr) - res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
