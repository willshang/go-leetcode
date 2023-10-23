package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestNumber([]int{4, 3, 2, 5, 6, 7, 2, 5, 5}, 8))
}

func largestNumber(cost []int, target int) string {
	dp := make([]int, target+1)
	for i := 0; i <= target; i++ {
		dp[i] = math.MinInt32
	}
	dp[0] = 0 // dp[i]表示花费总代价恰好为i时的最大长度
	for i := 0; i < len(cost); i++ {
		for j := cost[i]; j <= target; j++ {
			dp[j] = max(dp[j], dp[j-cost[i]]+1)
		}
	}
	if dp[target] < 0 {
		return "0"
	}
	res := ""
	j := target
	for i := 8; i >= 0; i-- {
		for c := cost[i]; j >= c && dp[j] == dp[j-c]+1; j = j - c {
			res = res + string('1'+i)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
