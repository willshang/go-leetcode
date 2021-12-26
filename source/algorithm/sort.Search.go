package main

import (
	"fmt"
	"sort"
)

func main() {
	var index int
	arr := []int{1, 1, 10, 10, 50, 51, 60, 90, 100, 100, 100, 101}
	index = sort.Search(len(arr), func(i int) bool {
		return arr[i] >= 100 // 8
	})
	fmt.Println(index)
	index = sort.Search(len(arr), func(i int) bool {
		return arr[i] >= 101 // 11
	})
	fmt.Println(index)
	index = sort.Search(len(arr), func(i int) bool {
		return arr[i] >= 1 // 0
	})
	fmt.Println(index)
	index = sort.Search(len(arr), func(i int) bool {
		return arr[i] >= 0 // 0
	})
	fmt.Println(index)
	index = sort.Search(len(arr), func(i int) bool {
		return arr[i] >= -1 // 0
	})
	fmt.Println(index)
	index = sort.Search(len(arr), func(i int) bool {
		return arr[i] >= 95 // 8
	})
	if index < len(arr) && arr[index] == 95 {
		fmt.Println("95在数组中:\t", arr)
	} else {
		fmt.Println("95不在数组中") // 输出
	}
}
