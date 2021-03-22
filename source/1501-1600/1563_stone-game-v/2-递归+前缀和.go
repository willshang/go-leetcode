package main

import "fmt"

func main() {
	fmt.Println(stoneGameV([]int{6, 2, 3, 4, 5, 5}))
}

var dp [][]int
var sum []int

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	sum = make([]int, n+1)
	sum[0] = 0
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + stoneValue[i]
	}
	dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -1
		}
	}
	return dfs(1, n)
}

func dfs(left, right int) int {
	if dp[left][right] != -1 {
		return dp[left][right]
	}
	if left == right {
		dp[left][right] = 0
	} else {
		value := 0
		for i := left; i < right; i++ {
			leftSum := sum[i] - sum[left-1]
			rightSum := sum[right] - sum[i]
			if leftSum < rightSum {
				value = max(value, leftSum+dfs(left, i))
			} else if leftSum > rightSum {
				value = max(value, rightSum+dfs(i+1, right))
			} else {
				value = max(value, max(dfs(left, i), dfs(i+1, right))+leftSum)
			}
		}
		dp[left][right] = value
	}
	return dp[left][right]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
