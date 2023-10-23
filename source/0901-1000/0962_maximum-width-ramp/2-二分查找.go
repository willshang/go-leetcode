package main

import (
	"fmt"
)

func main() {
	//fmt.Println(maxWidthRamp([]int{7, 5, 1, 2, 3, 4, 5, 10, 8, 7, 6, 5}))
	//fmt.Println(maxWidthRamp([]int{10, 10, 10, 7, 1, 6, 2, 1, 7}))
	fmt.Println(maxWidthRamp([]int{6, 0, 8, 2, 1, 5}))
}

type Node struct {
	Value int
	Index int
}

func maxWidthRamp(A []int) int {
	res := 0
	n := len(A)
	arr := make([]Node, 0) // Value递增数组
	arr = append(arr, Node{
		Value: A[n-1],
		Index: n - 1,
	})
	for i := n - 2; i >= 0; i-- {
		left, right := 0, len(arr)
		for left < right {
			mid := left + (right-left)/2
			if arr[mid].Value < A[i] { // 不满足条件
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left < len(arr) {
			index := arr[left].Index
			res = max(res, index-i)
		} else {
			arr = append(arr, Node{
				Value: A[i],
				Index: i,
			})
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
