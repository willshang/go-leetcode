package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{
		{1, 5},
		{1, 5},
		{1, 5}, {2, 3}, {2, 3}}
	fmt.Println(maxEvents(arr))
}

func maxEvents(events [][]int) int {
	n := len(events)
	minDay := events[0][0]
	sort.Slice(events, func(i int, j int) bool {
		if events[i][0] < minDay {
			minDay = events[i][0]
		}
		if events[j][0] < minDay {
			minDay = events[j][0]
		}
		return events[i][1] < events[j][1]
	})
	res := 0
	maxDay := events[n-1][1]
	dp := make([]bool, maxDay-minDay+1)
	for i := 0; i < n; i++ {
		for j := events[i][0]; j <= events[i][1]; j++ {
			if dp[j-minDay] == false {
				dp[j-minDay] = true
				res++
				break
			}
		}
	}
	return res
}
