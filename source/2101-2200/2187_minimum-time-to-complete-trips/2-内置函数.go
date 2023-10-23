package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumTime([]int{1, 2, 3}, 5))
}

func minimumTime(time []int, totalTrips int) int64 {
	n := len(time)
	maxValue := 0
	for i := 0; i < n; i++ {
		maxValue = max(maxValue, time[i])
	}
	right := totalTrips * maxValue
	return int64(sort.Search(right, func(per int) bool {
		count := 0
		for i := 0; i < len(time); i++ {
			count = count + int(per/time[i])
		}
		return count >= totalTrips
	}))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
