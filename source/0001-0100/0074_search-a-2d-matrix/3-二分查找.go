package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix(arr, 5))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] <= target && matrix[i][len(matrix[i])-1] >= target {
			res := binarySearch(matrix[i], target)
			if res == true {
				return true
			}
		}
	}
	return false
}

func binarySearch(arr []int, target int) bool {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
