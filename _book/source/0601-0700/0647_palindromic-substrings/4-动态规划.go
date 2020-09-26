package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("abc"))
}

func countSubstrings(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	dp := make([][]bool, len(s))
	res := 0
	for r := 0; r < len(s); r++ {
		dp[r] = make([]bool, len(s))
		dp[r][r] = true
		res++
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1] == true) {
				dp[l][r] = true
			} else {
				dp[l][r] = false
			}
			if dp[l][r] == true {
				res++
			}
		}
	}
	return res
}
