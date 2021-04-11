package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(breakfastNumber([]int{10, 20, 5}, []int{5, 5, 2}, 15))
	fmt.Println(breakfastNumber([]int{2, 1, 1}, []int{8, 9, 5, 1}, 9))
}

func breakfastNumber(staple []int, drinks []int, x int) int {
	sort.Ints(staple)
	sort.Ints(drinks)
	res := 0
	for i := 0; i < len(staple); i++ {
		target := x - staple[i]
		if target <= 0 {
			break
		}
		j := binarySearch(drinks, target)
		res = (res + j) % 1000000007
	}
	return res
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
