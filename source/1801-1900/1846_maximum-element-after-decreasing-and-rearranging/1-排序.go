package main

import "sort"

func main() {

}

// leetcode1846_减小和重新排列数组后的最大元素
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	n := len(arr)
	arr[0] = 1
	for i := 1; i < n; i++ {
		arr[i] = min(arr[i], arr[i-1]+1)
	}
	return arr[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
