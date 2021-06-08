package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(minFallingPathSum(arr))
}

func minFallingPathSum(arr [][]int) int {
	n := len(arr)
	for i := 1; i < n; i++ {
		temp := make([][2]int, n)
		for j := 0; j < n; j++ {
			temp[j] = [2]int{arr[i-1][j], j}
		}
		sort.Slice(temp, func(i, j int) bool {
			return temp[i][0] < temp[j][0]
		})
		for j := 0; j < n; j++ {
			if temp[0][1] != j {
				arr[i][j] = arr[i][j] + temp[0][0]
			} else {
				arr[i][j] = arr[i][j] + temp[1][0]
			}
		}
	}
	res := math.MaxInt32
	for j := 0; j < n; j++ {
		res = min(res, arr[n-1][j])
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
