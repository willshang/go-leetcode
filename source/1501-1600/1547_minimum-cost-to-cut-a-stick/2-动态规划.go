package main

import (
	"math"
	"sort"
)

func main() {

}

func minCost(n int, cuts []int) int {
	m := len(cuts)
	cuts = append(cuts, 0, n)
	sort.Ints(cuts)
	dp := make([][]int, m+2) // dp[L][R]为切割以L,R为左右端点的最小成本
	for i := 0; i < m+2; i++ {
		dp[i] = make([]int, m+2)
	}
	for length := 2; length <= m+1; length++ { // 枚举长度
		for i := 1; i <= m; i++ {
			j := i + length - 2
			if j > m {
				break
			}
			dp[i][j] = math.MaxInt32
			for k := i; k <= j; k++ { // 枚举切割点
				dp[i][j] = min(dp[i][j], dp[i][k-1]+dp[k+1][j])
			}
			dp[i][j] = dp[i][j] + cuts[j+1] - cuts[i-1]
		}
	}
	return dp[1][m]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
