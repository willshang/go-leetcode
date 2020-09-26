package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(12))
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	arr := make([]int, 0)
	arr = append(arr, 0)
	for i := 1; i*i <= n; i++ {
		arr = append(arr, i*i)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if i < arr[j] {
				break
			}
			dp[i] = min(dp[i], dp[i-arr[j]]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
