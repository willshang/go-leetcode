package main

import "fmt"

func main() {
	fmt.Println(stoneGameVII([]int{5, 3, 1, 4, 2}))
}

func stoneGameVII(stones []int) int {
	n := len(stones)
	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		arr[i+1] = arr[i] + stones[i]
	}
	dp := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		sum := stones[i]
		for j := i + 1; j < n; j++ {
			sum = sum + stones[j]
			left := sum - stones[i] - dp[j]    // 左边得分-得分
			right := sum - stones[j] - dp[j-1] // 右边得分-得分
			dp[j] = max(left, right)
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
