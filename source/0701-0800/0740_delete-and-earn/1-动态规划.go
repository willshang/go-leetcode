package main

func main() {

}

// leetcode740_删除与获得点数
func deleteAndEarn(nums []int) int {
	count := make([]int, 10001)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	dp := make([]int, 10001)
	dp[1] = count[1]
	dp[2] = max(dp[1], count[2]*2)
	for i := 2; i < 10001; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+i*count[i])
	}
	return dp[10000]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
