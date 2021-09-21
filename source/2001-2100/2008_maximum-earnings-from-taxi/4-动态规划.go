package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxTaxiEarnings(5, [][]int{
		{2, 5, 4},
		{1, 5, 1},
	}))
}

func maxTaxiEarnings(n int, rides [][]int) int64 {
	dp := make([]int, n+1) // dp[i]到达i位置的最大盈利
	sort.Slice(rides, func(i, j int) bool {
		if rides[i][0] == rides[j][0] {
			return rides[i][1] < rides[j][1]
		}
		return rides[i][0] < rides[j][0]
	})
	j := 1
	for i := 0; i < len(rides); i++ { // 遍历订单
		a, b, c := rides[i][0], rides[i][1], rides[i][2]
		for j < a { // 更新指针
			j++
			dp[j] = max(dp[j], dp[j-1])
		}
		dp[b] = max(dp[b], dp[j]+(b-a+c))
	}
	for ; j <= n; j++ {
		dp[j] = max(dp[j], dp[j-1])
	}
	return int64(dp[n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
