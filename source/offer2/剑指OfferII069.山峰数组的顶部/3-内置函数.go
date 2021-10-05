package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{0, 2, 1, 0}
	fmt.Println(peakIndexInMountainArray(arr))
}

func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	return sort.Search(n-1, func(i int) bool {
		return arr[i] > arr[i+1]
	})
}
