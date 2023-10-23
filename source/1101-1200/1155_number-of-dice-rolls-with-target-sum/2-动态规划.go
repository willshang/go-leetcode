package main

func main() {

}

func numRollsToTarget(d int, f int, target int) int {
	dp := make([][]int, d+1)
	for i := 0; i <= d; i++ {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1
	for i := 1; i <= d; i++ {
		for j := i; j <= target; j++ {
			for k := 1; k <= f; k++ {
				if j >= k {
					dp[i][j] = (dp[i][j] + dp[i-1][j-k]) % 1000000007
				}
			}
		}
	}
	return dp[d][target]
}
