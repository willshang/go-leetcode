package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minCost(7, []int{1, 3, 4, 5}))
}

// leetcode1547_切棍子的最小成本
func minCost(n int, cuts []int) int {
	m := len(cuts)
	cuts = append(cuts, 0, n)
	sort.Ints(cuts)
	dp := make([][]int, m+2) // dp[L][R]为切割以L,R为左右端点的最小成本
	for i := 0; i < m+2; i++ {
		dp[i] = make([]int, m+2)
	}
	for i := m; i >= 1; i-- {
		for j := i; j <= m; j++ {
			dp[i][j] = math.MaxInt32
			for k := i; k <= j; k++ { // 枚举切割点
				dp[i][j] = min(dp[i][j], dp[i][k-1]+dp[k+1][j])
			}
			dp[i][j] = dp[i][j] + cuts[j+1] - cuts[i-1]
		}
	}
	fmt.Println(dp)
	return dp[1][m]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
