package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))
}

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	result := make([][]int, 0)
	min := arr[1] - arr[0]
	result = append(result, []int{arr[0], arr[1]})
	for i := 2; i < len(arr); i++ {
		value := arr[i] - arr[i-1]
		if value < min {
			min = value
			result = make([][]int, 0)
			result = append(result, []int{arr[i-1], arr[i]})
		} else if value == min {
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}
	return result
}
