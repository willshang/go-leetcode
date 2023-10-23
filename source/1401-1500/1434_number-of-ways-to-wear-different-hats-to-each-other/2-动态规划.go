package main

func main() {

}

var mod = 1000000007

func numberWays(hats [][]int) int {
	n := len(hats)                  // n个人
	m := make(map[int]map[int]bool) // 帽子对应人的喜欢关系
	for i := 0; i < n; i++ {
		for j := 0; j < len(hats[i]); j++ {
			id := hats[i][j]
			if m[id] == nil {
				m[id] = make(map[int]bool)
			}
			m[id][i] = true
		}
	}
	target := (1 << n) - 1      // n个1
	dp := make([]int, target+1) // dp[i][j] 表示第i个帽子状态为j的方案数
	dp[0] = 1
	for i := 1; i <= 40; i++ {
		for j := target; j >= 0; j-- { // 状态为0~2^n-1
			for k := range m[i] { // 对第i个帽子喜欢
				if (j>>k)&1 == 0 { // 判断状态j的第k位是否是0
					// j+1<<k ：第k个人没有分配帽子，然后分配帽子
					next := j + 1<<k
					dp[next] = (dp[j] + dp[next]) % mod
				}
			}
		}
	}
	return dp[target]
}
