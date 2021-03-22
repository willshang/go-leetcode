package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minCost(7, []int{1, 3, 4, 5}))
}

var dp [][]int

func minCost(n int, cuts []int) int {
	m := len(cuts)
	cuts = append(cuts, 0, n)
	sort.Ints(cuts)
	dp = make([][]int, m+2) // dp[L][R]为切割以L,R为左右端点的最小成本
	for i := 0; i < m+2; i++ {
		dp[i] = make([]int, m+2)
	}
	return dfs(cuts, 1, m)
}

func dfs(cuts []int, i, j int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	if i > j {
		return 0
	}
	if i == j {
		return cuts[j+1] - cuts[i-1]
	}
	dp[i][j] = math.MaxInt32
	for k := i; k <= j; k++ {
		dp[i][j] = min(dp[i][j], dfs(cuts, i, k-1)+dfs(cuts, k+1, j))
	}
	dp[i][j] = dp[i][j] + cuts[j+1] - cuts[i-1]
	return dp[i][j]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
