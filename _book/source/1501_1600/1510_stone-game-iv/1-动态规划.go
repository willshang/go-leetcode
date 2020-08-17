package main

import "fmt"

func main() {
	fmt.Println(winnerSquareGame(7))
}

// leetcode1510_石子游戏IV
func winnerSquareGame(n int) bool {
	dp := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if dp[i-j*j] == false {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
