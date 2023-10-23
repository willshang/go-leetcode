package main

import "fmt"

func main() {
	fmt.Println(divisorGame(10))
}

// leetcode1025_除数博弈
func divisorGame(N int) bool {
	dp := make([]bool, N+1)
	dp[1] = false // 1的时候爱丽丝没有选择，失败
	for i := 2; i <= N; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && dp[i-j] == false {
				dp[i] = true
			}
		}
	}
	return dp[N]
}
