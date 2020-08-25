package main

import "fmt"

func main() {
	fmt.Println(countBits(10))
}

// leetcode338_比特位计数
func countBits(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i-1] + 1
		}
	}
	return dp
}
