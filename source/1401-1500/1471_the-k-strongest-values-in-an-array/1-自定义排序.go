package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getStrongest([]int{1, 2, 3, 4, 5}, 2))
}

func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	mid := arr[(len(arr)-1)/2]
	sort.Slice(arr, func(i, j int) bool {
		if abs(arr[i]-mid) == abs(arr[j]-mid) {
			return arr[i] > arr[j]
		}
		return abs(arr[i]-mid) > abs(arr[j]-mid)
	})
	return arr[:k]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
