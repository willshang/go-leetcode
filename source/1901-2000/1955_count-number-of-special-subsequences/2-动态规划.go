package main

func main() {

}

var mod = 1000000007

func countSpecialSubsequences(nums []int) int {
	n := len(nums)
	dp := make([]int, 3)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			dp[0] = (dp[0]*2 + 1) % mod
		} else if nums[i] == 1 {
			dp[1] = (dp[1]*2 + dp[0]) % mod
		} else if nums[i] == 2 {
			dp[2] = (dp[2]*2 + dp[1]) % mod
		}
	}
	return dp[2]
}
