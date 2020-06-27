package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
}

// leetcode1051_高度检查器
func heightChecker(heights []int) int {
	temp := make([]int, len(heights))
	copy(temp, heights)
	sort.Ints(temp)
	res := 0
	for i := 0; i < len(temp); i++ {
		if temp[i] != heights[i] {
			res++
		}
	}
	return res
}
