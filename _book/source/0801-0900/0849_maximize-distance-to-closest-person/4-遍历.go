package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 0, 0, 0, 1, 0, 1}
	fmt.Println(maxDistToClosest(arr))
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0}))
}

// leetcode849_到最近的人的最大距离
func maxDistToClosest(seats []int) int {
	res := 0
	count := 0
	for i := 0; i < len(seats); i++ {
		if count == i {
			res = count
		} else {
			res = max(res, (count+count%2)/2)
		}
		if seats[i] == 1 {
			count = 0
		} else {
			count++
		}
	}
	return max(res, count)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
