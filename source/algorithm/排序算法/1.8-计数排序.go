package main

import "fmt"

func main() {
	arr := []int{66, 33, 55, 22, 11, 99, 88, 77}
	fmt.Println(countingSort(arr))
}

func countingSort(arr []int) []int {
	maxValue := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	bucket := make([]int, maxValue+1)
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]]++
	}
	index := 0
	for i := 0; i < len(bucket); i++ {
		for bucket[i] > 0 {
			arr[index] = i
			index++
			bucket[i]--
		}
	}
	return arr
}
