package main

import "sort"

func main() {

}

func maxSatisfaction(satisfaction []int) int {
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] > satisfaction[j]
	})
	if satisfaction[0] <= 0 {
		return 0
	}
	n := len(satisfaction)
	dp := make([]int, n)
	dp[0] = satisfaction[0]
	sum := satisfaction[0]
	for i := 1; i < n; i++ {
		sum = sum + satisfaction[i]
		dp[i] = max(dp[i-1], dp[i-1]+sum)
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
