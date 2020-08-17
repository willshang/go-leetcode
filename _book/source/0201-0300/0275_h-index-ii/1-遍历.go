package main

import (
	"fmt"
)

func main() {
	fmt.Println(hIndex([]int{0, 1, 3, 5, 6}))
}

// leetcode275_H指数II
func hIndex(citations []int) int {
	for i := 0; i < len(citations); i++ {
		if citations[i] >= len(citations)-i {
			return len(citations) - i
		}
	}
	return 0
}
