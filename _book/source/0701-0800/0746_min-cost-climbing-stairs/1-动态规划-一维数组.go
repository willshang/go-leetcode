package main

import "fmt"

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}

/*
用dp[i]表示爬i个台阶所需要的成本，所以dp[0]=0，dp[1]=0
每次爬i个楼梯，计算的都是从倒数第一个结束，还是从倒数第二个结束
动态转移方程为:
dp[i] = min{dp[i-2]+cost[i-2] , dp[i-1]+cost[i-1]};
*/
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
