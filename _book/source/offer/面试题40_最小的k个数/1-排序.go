package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 1))
}

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	sort.Ints(arr)
	return arr[:k]
}
