package main

func main() {

}

// leetcode1039_多边形三角剖分的最低得分
func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n) // dp[i][j]表示从i到j序列的最低分
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	// 点0和n-1，因为总有某个点j与点0和n-1构成三角形
	// 这样就把1个多边形，转换为2个或者3个区域
	// dp[0][n-1] = min(dp[0][j] + dp[j][n-1] + A[0]*A[k]*A[n-1])
	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			for k := i + 1; k < j; k++ {
				sum := dp[i][k] + dp[k][j] + values[i]*values[k]*values[j]
				if dp[i][j] == 0 {
					dp[i][j] = sum
				} else {
					dp[i][j] = min(dp[i][j], sum)
				}
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
