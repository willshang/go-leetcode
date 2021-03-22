package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if countBit(arr[i]) == countBit(arr[j]) {
			return arr[i] < arr[j]
		}
		return countBit(arr[i]) < countBit(arr[j])
	})
	return arr
}

func countBit(num int) int {
	res := 0
	for num > 0 {
		if num%2 == 1 {
			res++
		}
		num = num / 2
	}
	return res
}
