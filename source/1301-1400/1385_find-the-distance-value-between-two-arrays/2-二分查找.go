package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2))
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	sort.Ints(arr2)
	for i := 0; i < len(arr1); i++ {
		if search(arr1[i], arr2, d) {
			res++
		}
	}
	return res
}

func search(target int, arr []int, d int) bool {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			if target-arr[mid] <= d {
				return false
			}
			left = mid + 1
		} else {
			if arr[mid]-target <= d {
				return false
			}
			right = mid - 1
		}
	}
	return true
}
