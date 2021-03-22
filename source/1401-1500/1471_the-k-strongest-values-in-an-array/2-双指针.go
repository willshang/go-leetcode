package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getStrongest([]int{1, 2, 3, 4, 5}, 2))
}

// leetcode1471_数组中的k个最强值
func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	mid := arr[(len(arr)-1)/2]
	res := make([]int, 0)
	left, right := 0, len(arr)-1
	for k > 0 {
		if arr[right]-mid >= mid-arr[left] {
			res = append(res, arr[right])
			right--
		} else {
			res = append(res, arr[left])
			left++
		}
		k--
	}
	return res
}
