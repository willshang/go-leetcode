package main

import "fmt"

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
}

// leetcode983_最低票价
func mincostTickets(days []int, costs []int) int {
	n := days[len(days)-1] + 1
	dp := make([]int, n)
	for i := 0; i < len(days); i++ {
		dp[days[i]] = 1 // 出行日
	}
	for i := 1; i < n; i++ {
		if dp[i] > 0 {
			dp[i] = min(dp[i-1]+costs[0],
				min(dp[max(i-7, 0)]+costs[1], dp[max(i-30, 0)]+costs[2]))
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
