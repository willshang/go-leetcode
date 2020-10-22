package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallestK([]int{1, 3, 5, 7, 2, 4, 6, 8}, 4))
}

func smallestK(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
