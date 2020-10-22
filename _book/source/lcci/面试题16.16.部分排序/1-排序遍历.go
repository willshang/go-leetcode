package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subSort([]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}))
}

func subSort(array []int) []int {
	temp := make([]int, len(array))
	copy(temp, array)
	sort.Ints(temp)
	left, right := -1, -1
	for i := 0; i < len(array); i++ {
		if temp[i] != array[i] {
			left = i
			break
		}
	}
	for i := len(array) - 1; i >= 0; i-- {
		if temp[i] != array[i] {
			right = i
			break
		}
	}
	return []int{left, right}
}
