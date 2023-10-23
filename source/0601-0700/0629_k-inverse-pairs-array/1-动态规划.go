package main

import "fmt"

func main() {
	fmt.Println(kInversePairs(3, 1))
	fmt.Println(kInversePairs(3, 3))
	fmt.Println(kInversePairs(2, 2))
}

var mod = 1000000007

func kInversePairs(n int, k int) int {
	dp := make([][]int, n+1) // dp[n][k]表示1-n的排列中，包含k个逆序对
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	for i := 1; i <= n; i++ {
		dp[i][0] = 1
		// 最多i*(i-1)/2
		for j := 1; j <= k && j <= i*(i-1)/2; j++ {
			// 前面i-1个数，i个插入位
			// 插入到最后，不增加： f(i,j) = f(i,j) + f(i-1,j)
			// 插入到倒数第2个，增加：f(i,j) = f(i,j) + f(i-1,j-1)
			// ...
			// 插入到倒数第i个，增加：f(i,j) = f(i,j) + f(i-1,j-i+1)
			// => f(i,j) = f(i-1,j) + f(i-1,j-1) + ... + f(i-1,j-i+1)
			// f(i,j-1) = f(i-1, j-1)+ ...+f(i-1, j-i)
			// f(i,j) - f(i,j-1) = f(i-1,j)-f(i-1, j-i)
			// => f(i, j) = f(i,j-1)+f(i-1,j)-f(i-1,j-i)
			if j >= i {
				dp[i][j] = dp[i][j-1] + dp[i-1][j] - dp[i-1][j-i]
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
			if dp[i][j] >= 0 {
				dp[i][j] = dp[i][j] % mod
			} else {
				dp[i][j] = (dp[i][j] + mod) % mod
			}
		}
	}
	return dp[n][k]
}
