package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{99, 1, 10, 8, 3, 49, 6}
	bucketSort(arr, 100)
	fmt.Println(arr)
}

func bucketSort(arr []int, bucketSize int) {
	maxValue := arr[0]
	minValue := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
		if arr[i] < minValue {
			minValue = arr[i]
		}
	}
	d := maxValue - minValue
	bucketList := make([][]int, bucketSize)
	for i := 0; i < bucketSize; i++ {
		bucketList[i] = make([]int, 0)
	}
	for i := 0; i < len(arr); i++ {
		index := (arr[i] - minValue) * (bucketSize - 1) / d
		bucketList[index] = append(bucketList[index], arr[i])
	}
	for i := 0; i < bucketSize; i++ {
		sort.Ints(bucketList[i])
	}
	res := make([]int, 0)
	for i := 0; i < bucketSize; i++ {
		for j := 0; j < len(bucketList[i]); j++ {
			res = append(res, bucketList[i][j])
		}
	}
	copy(arr, res)
}
