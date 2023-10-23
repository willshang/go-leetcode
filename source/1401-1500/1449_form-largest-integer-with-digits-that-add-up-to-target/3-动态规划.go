package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestNumber([]int{4, 3, 2, 5, 6, 7, 2, 5, 5}, 8))
}

// leetcode1449_数位成本和为目标值的最大数字
func largestNumber(cost []int, target int) string {
	dp := make([]string, target+1)
	for i := 1; i <= target; i++ {
		dp[i] = "0"
		for j := 8; j >= 0; j-- {
			if cost[j] > i || (i-cost[j] != 0 && dp[i-cost[j]] == "0") {
				continue
			}
			a, b := dp[i], dp[i-cost[j]]+fmt.Sprintf("%d", j+1)
			if len(a) < len(b) {
				dp[i] = b
			} else if len(a) == len(b) && a < b {
				dp[i] = b
			}
		}
	}
	return dp[target]
}
