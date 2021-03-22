package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(arrayRankTransform([]int{40, 10, 20, 30}))
	fmt.Println(arrayRankTransform([]int{-9999999, 9999999}))
}

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	min := math.MaxInt32
	max := math.MinInt32
	for i := 0; i < len(arr); i++ {
		if arr[i] <= min {
			min = arr[i]
		}
		if arr[i] >= max {
			max = arr[i]
		}
	}
	length := max - min + 1
	temp := make([]int, length)
	for i := 0; i < length; i++ {
		temp[i] = math.MinInt32
	}
	for i := 0; i < len(arr); i++ {
		temp[arr[i]-min] = -1
	}
	count := 0
	for i := 0; i < length; i++ {
		if temp[i] == -1 {
			temp[i] = count
			count++
		}
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = temp[arr[i]-min] + 1
	}
	return arr
}
