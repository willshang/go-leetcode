package main

func main() {

}

// leetcode790_多米诺和托米诺平铺
var mod = 1000000007

func numTilings(N int) int {
	dp := [4]int{}
	dp[0] = 1
	for i := 0; i < N; i++ {
		temp := [4]int{}
		temp[0] = (dp[0] + dp[3]) % mod
		temp[1] = (dp[0] + dp[2]) % mod
		temp[2] = (dp[0] + dp[1]) % mod
		temp[3] = (dp[0] + dp[1] + dp[2]) % mod
		dp = temp
	}
	return dp[0]
}
