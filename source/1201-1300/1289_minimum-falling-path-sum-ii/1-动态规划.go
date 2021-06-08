package main

import (
	"fmt"
	"math"
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
	firstMin, secondMin := 0, 0
	firstIndex := -1
	for i := 0; i < n; i++ {
		fMin, sMin := math.MaxInt32, math.MaxInt32
		fIndex := -1
		for j := 0; j < n; j++ {
			sum := 0
			if firstIndex != j { // 不等于最小值所在行，就+最小值
				sum = firstMin + arr[i][j]
			} else { // 等于最小值所在行，就+次小值
				sum = secondMin + arr[i][j]
			}
			if sum < fMin {
				sMin = fMin
				fMin = sum
				fIndex = j
			} else if sum < sMin {
				sMin = sum
			}
		}
		firstMin = fMin
		secondMin = sMin
		firstIndex = fIndex
	}
	return firstMin
}
