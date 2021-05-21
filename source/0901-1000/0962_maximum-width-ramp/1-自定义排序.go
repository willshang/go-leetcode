package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(maxWidthRamp([]int{7, 5, 1, 2, 3, 4, 5, 10, 8, 7, 6, 5}))
	fmt.Println(maxWidthRamp([]int{10, 10, 10, 7, 1, 6, 2, 1, 7}))
}

type Node struct {
	Value int
	Index int
}

func maxWidthRamp(A []int) int {
	res := 0
	n := len(A)
	arr := make([]Node, n)
	for i := 0; i < n; i++ {
		arr[i] = Node{
			Value: A[i],
			Index: i,
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].Value == arr[j].Value {
			return arr[i].Index < arr[j].Index
		}
		return arr[i].Value < arr[j].Value
	})
	minIndex := n // 保存当前遇到的最小下标
	for i := 0; i < n; i++ {
		res = max(res, arr[i].Index-minIndex)
		minIndex = min(minIndex, arr[i].Index)
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
