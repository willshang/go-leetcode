package main

import "fmt"

func main() {
	fmt.Println(countBits(10))
}

// 剑指OfferII003.前n个数字二进制中1的个数
func countBits(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i-1] + 1
		}
	}
	return dp
}
