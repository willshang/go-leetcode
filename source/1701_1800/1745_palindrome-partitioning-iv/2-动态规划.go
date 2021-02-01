package main

func main() {

}

func checkPartitioning(s string) bool {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = true
			} else if i+1 == j && s[i] == s[j] {
				dp[i][j] = true
			} else {
				if s[i] == s[j] && dp[i+1][j-1] == true {
					dp[i][j] = true
				}
			}
		}
	}
	for i := 0; i < n-2; i++ {
		if dp[0][i] == false {
			continue
		}
		for j := i + 1; j < n-1; j++ {
			if dp[i+1][j] && dp[j+1][n-1] {
				return true
			}
		}
	}
	return false
}
