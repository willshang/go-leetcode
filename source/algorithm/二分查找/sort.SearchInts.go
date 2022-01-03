package main

import (
	"fmt"
	"sort"
)

func main() {
	var index int
	arr := []int{1, 1, 10, 10, 50, 51, 60, 90, 100, 100, 100, 101}
	index = sort.SearchInts(arr, 100) // 8
	fmt.Println(index)
	index = sort.SearchInts(arr, 101) // 11
	fmt.Println(index)
	index = sort.SearchInts(arr, 1) // 0
	fmt.Println(index)
	index = sort.SearchInts(arr, 0) // 0
	fmt.Println(index)
	index = sort.SearchInts(arr, -1) //0
	fmt.Println(index)
	index = sort.SearchInts(arr, 95)
	if index < len(arr) && arr[index] == 95 {
		fmt.Println("95在数组中:\t", arr)
	} else {
		fmt.Println("95不在数组中") // 输出
	}
}
