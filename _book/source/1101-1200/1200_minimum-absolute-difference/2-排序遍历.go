package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))
}

// leetcode1200_最小绝对差
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	result := make([][]int, 0)
	min := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if min > arr[i]-arr[i-1] {
			min = arr[i] - arr[i-1]
		}
	}
	for i := 1; i < len(arr); i++ {
		if min == arr[i]-arr[i-1] {
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}
	return result
}
