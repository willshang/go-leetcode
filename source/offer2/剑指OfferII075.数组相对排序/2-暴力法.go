package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
}

// 剑指OfferII075.数组相对排序
func relativeSortArray(arr1 []int, arr2 []int) []int {
	count := 0
	for i := 0; i < len(arr2); i++ {
		for j := count; j < len(arr1); j++ {
			if arr2[i] == arr1[j] {
				arr1[count], arr1[j] = arr1[j], arr1[count]
				count++
			}
		}
	}
	sort.Ints(arr1[count:])
	return arr1
}
