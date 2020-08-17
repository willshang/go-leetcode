package main

import (
	"fmt"
)

func main() {
	arr := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	fmt.Println(kthSmallest(arr, 3))
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)/2
		if check(matrix, mid, k, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 左下角查找
func check(matrix [][]int, mid, k, n int) bool {
	i := n - 1
	j := 0
	num := 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			// 往右,说明左边一列都小于mid
			num = num + i + 1
			j++
		} else {
			// 往上
			i--
		}
	}
	return num >= k
}
