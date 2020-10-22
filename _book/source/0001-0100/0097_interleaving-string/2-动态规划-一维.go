package main

import "fmt"

func main() {
	//fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("a", "", "b"))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if n+m != t {
		return false
	}
	// dp[j]表示s1的前i个元素和s2的前j个元素是否能交错组成s3的前i+j个元素
	dp := make([]bool, m+1)
	dp[0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			total := i + j - 1
			if i > 0 {
				if dp[j] == true && s1[i-1] == s3[total] {
					dp[j] = true
				} else {
					dp[j] = false
				}
			}
			if j > 0 {
				if dp[j] == true || (dp[j-1] == true && s2[j-1] == s3[total]) {
					dp[j] = true
				} else {
					dp[j] = false
				}
			}
		}
	}
	return dp[m]
}
