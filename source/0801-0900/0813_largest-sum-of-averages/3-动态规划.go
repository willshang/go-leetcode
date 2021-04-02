package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestSumOfAverages([]int{9, 1, 2, 3, 9}, 3))
}

func largestSumOfAverages(A []int, K int) float64 {
	n := len(A)
	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		arr[i+1] = arr[i] + A[i]
	}
	dp := make([][]float64, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]float64, K+1)
		dp[i][1] = float64(arr[i]) / float64(i) // 划分为1组
	}
	for i := 1; i <= n; i++ {
		for k := 2; k <= K && k <= i; k++ {
			for j := 1; j < i; j++ {
				target := dp[j][k-1] + float64(arr[i]-arr[j])/float64(i-j)
				dp[i][k] = math.Max(dp[i][k], target)
			}
		}
	}
	return dp[n][K]
}
