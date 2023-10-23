package main

func main() {

}

var mod = 1000000007

func firstDayBeenInAllRooms(nextVisit []int) int {
	n := len(nextVisit)
	dp := make([]int, n) // dp[i]=>访问房间i所需要的天数
	dp[0] = 1
	// 访问到i点时，前面所有的点的访问次数必为偶数
	for i := 0; i < n-1; i++ {
		a := (dp[i] - dp[nextVisit[i]] + mod) % mod
		dp[i+1] = (dp[i] + a + 2) % mod
	}
	return (dp[n-1] - 1 + mod) % mod // 第几天从0开始，处理下标
}
