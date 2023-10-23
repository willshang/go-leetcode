package main

func main() {

}

// leetcode300.最长上升子序列
func canBeIncreasing(nums []int) bool {
	n := len(nums)
	dp := make([]int, n)
	res := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		res = max(res, dp[i])
	}
	return res == n || res == n-1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
