package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{1, 4, 4, 4, 4}))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	n := len(nums)
	length := 1 << n
	for i := 0; i < length; i++ {

	}
	return res
}
