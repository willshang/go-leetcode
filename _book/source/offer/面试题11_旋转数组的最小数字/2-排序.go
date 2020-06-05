package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
}

func minArray(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}
