package main

func main() {

}

// leetcode1434_每个人戴不同帽子的方案数
var mod = 1000000007

func numberWays(hats [][]int) int {
	n := len(hats)
	target := (1 << n) - 1
	dp := make([]int, target+1)
	dp[0] = 1
	arr := make([][]int, 41)
	for i := 0; i < n; i++ {
		for j := range hats[i] {
			arr[hats[i][j]] = append(arr[hats[i][j]], i)
		}
	}
	for i := 1; i <= 40; i++ {
		for j := target - 1; j >= 0; j-- {
			for _, k := range arr[i] {
				if j&(1<<k) == 0 {
					dp[j|(1<<k)] = dp[j|(1<<k)] + dp[j]
				}
			}
		}
	}
	return dp[target] % mod
}
