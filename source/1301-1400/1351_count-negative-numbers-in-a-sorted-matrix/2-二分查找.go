package main

import "fmt"

func main() {
	arr := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	//arr := [][]int{
	//	{3,2},
	//	{1,0},
	//}
	fmt.Println(countNegatives(arr))
}

func countNegatives(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		res = res + (len(grid[i]) - search(grid[i]))
	}
	return res
}

func search(arr []int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] >= 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
