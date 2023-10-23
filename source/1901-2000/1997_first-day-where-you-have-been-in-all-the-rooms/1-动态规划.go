package main

func main() {

}

var mod = 1000000007

func firstDayBeenInAllRooms(nextVisit []int) int {
	n := len(nextVisit)
	dp := make([]int, n) // dp[i]=>访问房间i所需要的天数
	dp[0] = 1
	// 访问到i点时，前面所有的点的访问次数必为偶数
	for i := 1; i < n; i++ {
		// dp[i] = (2*dp[i-1] - dp[nextVisit[i-1]] + 2 + mod) % mod
		a := dp[i-1] + 1                      // 第一次：i-1 => i 等于到达dp[i-1]的时间+1
		b := dp[i-1] - dp[nextVisit[i-1]] + 1 // 第二次：第一个到达i-1会回去nextVisit[i-1]，2者相减后+1次返回
		dp[i] = (a + b + mod) % mod
	}
	return (dp[n-1] - 1 + mod) % mod // 第几天从0开始，处理下标
}
