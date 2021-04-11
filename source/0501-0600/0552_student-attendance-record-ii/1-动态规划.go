package main

func main() {

}

// leetcode552_学生出勤记录II
var mod = 1000000007

func checkRecord(n int) int {
	dp := [6]int{}
	dp[0] = 1 // 0A 0L
	dp[1] = 1 // 0A 1L
	dp[2] = 0 // 0A 2L
	dp[3] = 1 // 1A 0L
	dp[4] = 0 // 1A 1L
	dp[5] = 0 // 1A 2L
	for i := 2; i <= n; i++ {
		temp := [6]int{}
		temp[0] = (dp[0] + dp[1] + dp[2]) % mod                         // +P
		temp[1] = dp[0]                                                 // +L
		temp[2] = dp[1]                                                 // +L
		temp[3] = (dp[0] + dp[1] + dp[2] + dp[3] + dp[4] + dp[5]) % mod // 0、1、2+A，3、4、5+P
		temp[4] = dp[3]                                                 // +L
		temp[5] = dp[4]                                                 // +L
		dp = temp
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res = (res + dp[i]) % 1000000007
	}
	return res
}
