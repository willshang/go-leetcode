package main

func main() {

}

// leetcode1140_石子游戏II
func stoneGameII(piles []int) int {
	n := len(piles)
	dp := make([][]int, n+1) // dp[i][j]=>有piles[i:]，M=j的情况下得分
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	sum := 0
	for i := n - 1; i >= 0; i-- {
		sum = sum + piles[i]
		for M := 1; M <= n; M++ {
			if i+2*M >= n { // 可以全部拿走
				dp[i][M] = sum
			} else {
				for j := 1; j <= 2*M; j++ { // 尝试不同拿法，dp[i+j][max(j, M)]其中M=max(j,M)
					dp[i][M] = max(dp[i][M], sum-dp[i+j][max(j, M)])
				}
			}
		}
	}
	return dp[0][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
