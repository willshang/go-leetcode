package main

func main() {

}

// leetcode1866_恰有K根木棍可以看到的排列数目
func rearrangeSticks(n int, k int) int {
	dp := make([][]int, n+1) // dp[i][j]代表i根棍子能看到j根；假设每次插入1，假设上一轮的数为：2到n
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			a := dp[i-1][j-1]                        // i-1根棍子里面已经看到了j-1根，插入到最前面
			b := ((i - 1) * dp[i-1][j]) % 1000000007 // i-1根棍子里面已经看到了j根，插入到除最前面的其它i-1个地方都行
			dp[i][j] = (a + b) % 1000000007
		}
	}
	return dp[n][k]
}
