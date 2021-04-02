package main

func main() {

}

// leetcode1223_掷骰子模拟
func dieSimulator(n int, rollMax []int) int {
	dp := make([][7][16]int, n+1) // 第i轮,以j结尾,出现k次
	for j := 1; j <= 6; j++ {
		dp[1][j][1] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 1; j <= 6; j++ {
			for k := 1; k <= rollMax[j-1] && k <= i; k++ {
				if k == 1 { // 不同情况
					for a := 1; a <= 6; a++ {
						if a != j { // 考虑不同的情况
							for b := 1; b <= rollMax[a-1] && b <= i-1; b++ {
								dp[i][j][k] = (dp[i][j][k] + dp[i-1][a][b]) % 1000000007
							}
						}
					}
				} else { // 直接同上一个
					dp[i][j][k] = dp[i-1][j][k-1]
				}
			}
		}
	}
	res := 0
	for j := 1; j <= 6; j++ {
		for k := 1; k <= 15; k++ {
			res = (res + dp[n][j][k]) % 1000000007
		}
	}
	return res
}
