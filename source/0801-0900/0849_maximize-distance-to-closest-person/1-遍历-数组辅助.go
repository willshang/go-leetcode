package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 0, 0, 0, 1, 0, 1}
	fmt.Println(maxDistToClosest(arr))
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0}))
}

func maxDistToClosest(seats []int) int {
	n := len(seats)
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		left[i], right[i] = n, n
	}
	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			left[i] = 0
		} else if seats[i] != 1 && i > 0 {
			left[i] = left[i-1] + 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		if seats[i] == 1 {
			right[i] = 0
		} else if seats[i] != 1 && i < n-1 {
			right[i] = right[i+1] + 1
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if seats[i] == 0 {
			if min(left[i], right[i]) > res {
				res = min(left[i], right[i])
			}
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
