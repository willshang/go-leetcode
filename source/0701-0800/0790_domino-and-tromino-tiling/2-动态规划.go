package main

func main() {

}

func numTilings(N int) int {
	dp := make([]int, N+3)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5
	for i := 4; i <= N; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % 1000000007
	}
	return dp[N]
}
