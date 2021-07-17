package main

import "fmt"

func main() {
	fmt.Println(kInversePairs(3, 1))
	fmt.Println(kInversePairs(3, 3))
	fmt.Println(kInversePairs(2, 2))
}

// leetcode629_K个逆序对数组
var mod = 1000000007

func kInversePairs(n int, k int) int {
	dp := make([]int, k+1) // dp[k]包含k个逆序对的方案数
	dp[0] = 1
	sum := make([]int, k+2)
	sum[1] = 1
	for i := 1; i <= n; i++ {
		// 最多i*(i-1)/2
		for j := 1; j <= k && j <= i*(i-1)/2; j++ {
			// 前面i-1个数，i个插入位
			// 插入到最后，不增加： f(i,j) = f(i,j) + f(i-1,j)
			// 插入到倒数第2个，增加：f(i,j) = f(i,j) + f(i-1,j-1)
			// ...
			// 插入到倒数第i个，增加：f(i,j) = f(i,j) + f(i-1,j-i+1)
			// => f(i,j) = f(i-1,j) + f(i-1,j-1) + ... + f(i-1,j-i+1)
			// => f(j) = sum[j+1] - sum[j-i+1]
			dp[j] = (sum[j+1] - sum[max(0, j-i+1)]) % mod
		}
		for j := 1; j <= k; j++ {
			sum[j+1] = sum[j] + dp[j]
		}
	}
	return dp[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
