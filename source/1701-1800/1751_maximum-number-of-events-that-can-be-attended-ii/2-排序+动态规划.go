package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{
		{1, 2, 4},
		{3, 4, 3},
		{2, 3, 1},
	}
	fmt.Println(maxValue(arr, 2))
}

// leetcode1751_最多可以参加的会议数目II
func maxValue(events [][]int, k int) int {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		if events[i][1] == events[j][1] {
			return events[i][2] < events[j][2]
		}
		return events[i][1] < events[j][1]
	})
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	for i := 1; i <= n; i++ {
		index := 0
		for j := i - 1; j >= 1; j-- {
			if events[j-1][1] < events[i-1][0] {
				index = j
				break
			}
		}
		for j := 1; j <= k; j++ {
			dp[i][j] = max(dp[i-1][j], dp[index][j-1]+events[i-1][2])
		}
	}
	return dp[n][k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
