package main

func main() {

}

// leetcode_lcp47_入场安检
var mod = 1000000007

func securityCheck(capacities []int, k int) int {
	dp := make([]int, k+1) // dp[i] => 使得第i个人 第1个离开的方案数量
	dp[0] = 1
	for i := 0; i < len(capacities); i++ {
		// 先进先出不改变 入场顺序
		// 改变为 后进先出 的入场顺序：会留住 capacities[i]-1 个人
		// 把c（capacities[i]-1）看成物品大小, k看成背包容量，等价求 01背包 的方案数。
		// 从后向前遍历
		c := capacities[i] - 1
		for j := k; j >= c; j-- {
			dp[j] = (dp[j] + dp[j-c]) % mod
		}
	}
	return dp[k]
}
