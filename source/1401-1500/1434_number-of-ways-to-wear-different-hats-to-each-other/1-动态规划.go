package main

func main() {

}

var mod = 1000000007

func numberWays(hats [][]int) int {
	n := len(hats)                  // n个人
	maxValue := 0                   // 最大的帽子编号，从1开始
	m := make(map[int]map[int]bool) // 帽子对应人的喜欢关系
	for i := 0; i < n; i++ {
		for j := 0; j < len(hats[i]); j++ {
			id := hats[i][j]
			maxValue = max(maxValue, id)
			if m[id] == nil {
				m[id] = make(map[int]bool)
			}
			m[id][i] = true
		}
	}
	dp := make([][]int, maxValue+1) // dp[i][j] 表示第i个帽子状态为j的方案数
	for i := 0; i <= maxValue; i++ {
		dp[i] = make([]int, 1<<n)
	}
	dp[0][0] = 1
	target := (1 << n) - 1 // n个1
	for i := 1; i <= maxValue; i++ {
		for j := 0; j <= target; j++ { // 状态为0~2^n-1
			dp[i][j] = dp[i-1][j]
			for k := range m[i] { // 对第i个帽子喜欢
				if (j>>k)&1 == 1 { // 判断状态j的第k位是否是1
					// j ^ (1<<k) ：第k个人没有分配帽子
					dp[i][j] = (dp[i][j] + dp[i-1][j^(1<<k)]) % mod
				}
			}
		}
	}
	return dp[maxValue][target]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
