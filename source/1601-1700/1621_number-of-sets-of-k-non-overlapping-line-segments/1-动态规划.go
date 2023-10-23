package main

func main() {

}

// leetcode1621_大小为K的不重叠线段的数目
var mod = 1000000007

func numberOfSets(n int, k int) int {
	dp := make([][][2]int, n)
	// dp[i][j][0] =>0到i的点，构造成j条线段的方案数, 第j条线段的右端点没有使用i
	// dp[i][j][1] =>0到i的点，构造成j条线段的方案数, 第j条线段的右端点使用i
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, n)
	}
	dp[0][0][0] = 1
	for i := 1; i < n; i++ {
		for j := 0; j <= k; j++ {
			dp[i][j][0] = (dp[i-1][j][0] + dp[i-1][j][1]) % mod // 没有使用i
			dp[i][j][1] = dp[i-1][j][1]                         // 使用i：扩展右侧
			if j > 0 {
				dp[i][j][1] = (dp[i][j][1] + dp[i-1][j-1][0] + dp[i-1][j-1][1]) % mod // 使用i：新开一个
			}
		}
	}
	return (dp[n-1][k][0] + dp[n-1][k][1]) % mod
}
