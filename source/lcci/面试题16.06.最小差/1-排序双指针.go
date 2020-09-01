package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//fmt.Println(smallestDifference([]int{1, 3, 15, 11, 2}, []int{23, 127, 235, 19, 8}))
	fmt.Println(smallestDifference([]int{2147483647}, []int{0}))
}

// 程序员面试金典16.06_最小差
func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	i, j := 0, 0
	res := math.MaxInt32
	for i < len(a) && j < len(b) {
		res = min(res, abs(a[i], b[j]))
		if a[i] > b[j] {
			j++
		} else {
			i++
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
