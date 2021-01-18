package main

import "fmt"

func main() {
	fmt.Println(countNumbersWithUniqueDigits(10))
}

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	dp := make([]int, n+10)
	dp[1] = 10
	prev := 9
	/*
		n = 1, 1+9
		n = 2, 1+9+9*9
		n = 3, 1+9+9*9+9*9*8
		n = 4, 1+9+9*9+9*9*8+9*9*8*7
	*/
	for i := 2; i <= 10; i++ {
		dp[i] = dp[i-1] + 9*prev
		prev = prev * (10 - i)
	}
	if n >= 10 {
		return dp[10]
	}
	return dp[n]
}
