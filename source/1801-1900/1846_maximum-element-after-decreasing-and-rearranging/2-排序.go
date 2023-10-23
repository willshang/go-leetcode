package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumElementAfterDecrementingAndRearranging([]int{0}))
}

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	res := 0
	for i := 0; i < len(arr); i++ {
		if res < arr[i] {
			res++
		}
	}
	return res
}
