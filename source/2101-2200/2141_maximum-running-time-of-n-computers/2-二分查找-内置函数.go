package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxRunTime(2, []int{3, 3, 3}))
}

func maxRunTime(n int, batteries []int) int64 {
	sum := 0
	for i := 0; i < len(batteries); i++ {
		sum = sum + batteries[i]
	}
	return int64(sort.Search(sum/n, func(target int) bool {
		target++
		total := 0
		for i := 0; i < len(batteries); i++ {
			total = total + min(batteries[i], target)
		}
		return target*n > total
	}))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
