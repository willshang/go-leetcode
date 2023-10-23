package main

func main() {

}

// leetcode1458_两个子序列的最大点积
func maxDotProduct(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			value := nums1[i] * nums2[j] // 单独选对应的i,j 乘积
			dp[i][j] = value             // 至少选一对
			if i > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j]) // 不选nums1[i]
			}
			if j > 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1]) // 不选nums2[j]
			}
			if i > 0 && j > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+value) // 选择nums1[i]和nums2[j]
			}
		}
	}
	return dp[n-1][m-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
