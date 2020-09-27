package main

import "fmt"

func main() {
	fmt.Println(minSteps(10))
}

func minSteps(n int) int {
	dp := make([]int, n+3)
	if n <= 1 {
		return 0
	}
	dp[0] = 0
	dp[1] = 0
	dp[2] = 2
	for i := 3; i <= n; i++ {
		minValue := i
		for j := i / 2; j >= 2; j-- {
			if i%j == 0 {
				minValue = dp[j] + i/j
				break
			}
		}
		dp[i] = minValue
	}
	return dp[n]
}
