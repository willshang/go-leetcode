package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

/*
dp[i]表示n=i时括号的组合
dp[i]="(" + dp[j] + ")"+dp[i-j-1] (j<i)
dp[0] = ""
*/
func generateParenthesis(n int) []string {
	dp := make([][]string, n+1)
	dp[0] = make([]string, 0)
	if n == 0 {
		return dp[0]
	}
	dp[0] = append(dp[0], "")
	for i := 1; i <= n; i++ {
		dp[i] = make([]string, 0)
		for j := 0; j < i; j++ {
			for _, a := range dp[j] {
				for _, b := range dp[i-j-1] {
					str := "(" + a + ")" + b
					dp[i] = append(dp[i], str)
				}
			}
		}
	}
	return dp[n]
}
