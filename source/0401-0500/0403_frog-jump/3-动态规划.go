package main

func main() {
}

func canCross(stones []int) bool {
	n := len(stones)
	dp := make([][]bool, n+1)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			k := stones[i] - stones[j]
			if k <= i {
				dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1] // 满足其一
				if i == n-1 && dp[i][k] == true {
					return true
				}
			}
		}
	}
	return false
}
