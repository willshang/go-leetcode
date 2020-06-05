package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1}))
	//fmt.Println(maxDistToClosest([]int{1, 0, 0, 0}))
}

func maxDistToClosest(seats []int) int {
	n := len(seats)
	var arr []int
	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			arr = append(arr, i)
		}
	}
	if len(arr) == 0 {
		return 0
	}
	max := -1
	for i := 0; i < n-1; i++ {
		if arr[i+1]-arr[i] > max {
			max = arr[i+1] - arr[i]
		}
	}
	max = max / 2
	// 判断首尾
	if arr[0] > max {
		max = arr[0]
	}
	if n-arr[len(arr)-1]-1 > max {
		max = n - arr[len(arr)-1] - 1
	}
	return max
}
