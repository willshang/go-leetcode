package main

func main() {

}

// leetcode1997_访问完所有房间的第一天
var mod = 1000000007

func firstDayBeenInAllRooms(nextVisit []int) int {
	n := len(nextVisit)
	dp := make([]int, n) // dp[i]=>访问房间i所需要的天数
	dp[0] = 0            // 下标从0开始
	// 访问到i点时，前面所有的点的访问次数必为偶数
	for i := 1; i < n; i++ {
		dp[i] = (2*dp[i-1] - dp[nextVisit[i-1]] + 2 + mod) % mod
	}
	return dp[n-1]
}
