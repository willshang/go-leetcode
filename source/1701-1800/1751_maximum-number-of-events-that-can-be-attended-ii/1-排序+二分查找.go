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
		for j := 1; j <= k; j++ {
			dp[i][j] = dp[i-1][j]
			left, right := 0, i-1
			for left < right {
				mid := left + (right-left)/2
				if events[mid][1] < events[i-1][0] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			dp[i][j] = max(dp[i][j], dp[right][j-1]+events[i-1][2])
		}
	}
	res := 0
	for i := 1; i <= k; i++ {
		res = max(res, dp[n][i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
