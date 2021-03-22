package main

import "fmt"

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
}

// leetcode1539_第k个缺失的正整数
func findKthPositive(arr []int, k int) int {
	left := 0
	right := len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid]-(mid+1) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return k + left
}
