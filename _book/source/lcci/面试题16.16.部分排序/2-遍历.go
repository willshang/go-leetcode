package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(subSort([]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}))
}

// 程序员面试金典16.16_部分排序
func subSort(array []int) []int {
	left, right := -1, -1
	maxValue := math.MinInt32
	minValue := math.MaxInt32
	for i := 0; i < len(array); i++ {
		if array[i] >= maxValue {
			maxValue = array[i]
		} else {
			right = i
		}
	}
	for i := len(array) - 1; i >= 0; i-- {
		if minValue >= array[i] {
			minValue = array[i]
		} else {
			left = i
		}
	}
	return []int{left, right}
}
