package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	fmt.Println(calculateMinimumHP(arr))
}

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return 0
	}
	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)/2
		if judge(dungeon, mid) == true {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func judge(dungeon [][]int, hp int) bool {
	n, m := len(dungeon), len(dungeon[0])
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = math.MinInt32
		}
	}
	dp[0][1], dp[1][0] = hp, hp
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			value := max(dp[i-1][j], dp[i][j-1]) + dungeon[i-1][j-1]
			if value <= 0 {
				continue
			}
			dp[i][j] = value
		}
	}
	return dp[n][m] > 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
