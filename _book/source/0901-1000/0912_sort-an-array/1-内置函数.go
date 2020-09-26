package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{66, 33, 55, 22, 11, 99, 88, 77}
	sortArray(arr)
	fmt.Println(arr)
}

func sortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}
