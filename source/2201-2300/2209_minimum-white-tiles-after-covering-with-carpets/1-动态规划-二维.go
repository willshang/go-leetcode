package main

func main() {

}

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	n := numCarpets
	m := len(floor)
	dp := make([][]int, n+1) // dp[i][j]=>表示i条毛毯覆盖前j块砖块时白色砖块的最少数目
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = int(floor[0] - '0')
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + int(floor[j]-'0')
	}
	for i := 1; i <= n; i++ {
		for j := carpetLen; j < m; j++ { // 遍历覆盖终点
			// 不覆盖：dp[i][j-1]+int(floor[j]-'0')
			// 覆盖：dp[i-1][j-carpetLen]（其中j-carpetLen是起点）
			dp[i][j] = min(dp[i][j-1]+int(floor[j]-'0'), dp[i-1][j-carpetLen])
		}
	}
	return dp[n][m-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
