package main

import "fmt"

func main() {
	fmt.Println(cuttingRope(10))
}

var dp []int

func cuttingRope(n int) int {
	if n <= 1 {
		return 1
	}
	dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 2; i <= n; i++ {
		//dp[i] = helper
	}
}
