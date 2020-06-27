package main

import (
	"fmt"
)

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	temp := make([]int, 1001)
	for i := range arr1 {
		temp[arr1[i]]++
	}
	count := 0
	for i := range arr2 {
		for temp[arr2[i]] > 0 {
			arr1[count] = arr2[i]
			temp[arr2[i]]--
			count++
		}
	}
	for i := 0; i < len(temp); i++ {
		for temp[i] > 0 {
			arr1[count] = i
			temp[i]--
			count++
		}
	}
	return arr1
}
