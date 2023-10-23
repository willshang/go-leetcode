package main

func main() {

}

func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	dp := make([][2]int, n)
	dp[0][0] = int(nums[0])
	dp[0][1] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+nums[i])
		dp[i][1] = max(dp[i-1][1], dp[i][0]-nums[i])
	}
	return int64(max(dp[n-1][0], dp[n-1][1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
