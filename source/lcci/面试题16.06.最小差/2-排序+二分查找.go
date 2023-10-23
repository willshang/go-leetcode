package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//fmt.Println(smallestDifference([]int{1, 3, 15, 11, 2}, []int{23, 127, 235, 19, 8}))
	fmt.Println(smallestDifference([]int{-2147483648, 1}, []int{2147483647, 0}))
}

func smallestDifference(a []int, b []int) int {
	sort.Ints(b)
	res := math.MaxInt32
	for i := 0; i < len(a); i++ {
		left, right := 0, len(b)-1
		for left <= right {
			mid := left + (right-left)/2
			if b[mid] == a[i] {
				return 0
			} else if b[mid] > a[i] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if left < len(b) {
			res = min(res, abs(a[i], b[left]))
		}
		if left > 0 {
			res = min(res, abs(a[i], b[left-1]))
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
