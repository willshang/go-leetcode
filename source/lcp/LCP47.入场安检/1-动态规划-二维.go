package main

func main() {

}

var mod = 1000000007

func securityCheck(capacities []int, k int) int {
	n := len(capacities)
	dp := make([][]int, n+1) // dp[i][j] => 使得i个实验室的情况下第j个人 第1个离开的方案数量
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		// 先进先出不改变 入场顺序
		// 改变为 后进先出 的入场顺序：会留住 capacities[i]-1 个人
		// 把c（capacities[i]-1）看成物品大小, k看成背包容量，等价求 01背包 的方案数。
		// 从后向前遍历
		c := capacities[i-1] - 1
		for j := 0; j <= k; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= c {
				dp[i][j] = (dp[i][j] + dp[i-1][j-c]) % mod
			}
		}
	}
	return dp[n][k]
}
