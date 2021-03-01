package main

func main() {

}

// leetcode474_一和零
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(strs); i++ {
		a, b := getCount(strs[i])
		for j := m; j >= a; j-- {
			for k := n; k >= b; k-- {
				dp[j][k] = max(dp[j][k], dp[j-a][k-b]+1)
			}
		}
	}
	return dp[m][n]
}

func getCount(str string) (a, b int) {
	a, b = 0, 0
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			a++
		} else {
			b++
		}
	}
	return a, b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
