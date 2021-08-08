package main

func main() {

}

// leetcode1388_3n块披萨
func maxSizeSlices(slices []int) int {
	a := calculate(slices[1:])             // 去除第一个数
	b := calculate(slices[:len(slices)-1]) // 去除最后一个数
	return max(a, b)
}

func calculate(slices []int) int {
	n := len(slices)
	target := (n + 1) / 3
	dp := make([][]int, n+1) // dp[i][j] => 在前i个数，选择j个不相邻的数的最大和
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			var a, b int
			if 2 <= i {
				a = dp[i-2][j-1]
			}
			a = a + slices[i-1] // 选择第i-1个数
			b = dp[i-1][j]      // 不选第i-1个数
			dp[i][j] = max(a, b)
		}
	}
	return dp[n][target]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
