package main

import "fmt"

func main() {
	fmt.Println(isScramble("great", "rgeat"))
}

// leetcode87_扰乱字符串
func isScramble(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n != m {
		return false
	}
	// dp[i][j][l]:表示s1从i开始，s2从j开始长度为l的两个子字符串是扰乱
	dp := make([][][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]bool, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]bool, n+1)
		}
	}
	// 单个字符
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j][1] = s1[i] == s2[j]
		}
	}
	for k := 2; k <= n; k++ { // 枚举长度: 2-n
		for i := 0; i <= n-k; i++ { // s1起点
			for j := 0; j <= n-k; j++ { // s2起点
				dp[i][j][k] = false
				// 长度为w，分为两部分，其中最少是1
				for w := 1; w <= k-1; w++ {
					// 划分不交换：S1->T1, S2->T2
					// 划分交换： S1->T2, S2->T1
					if (dp[i][j][w] == true && dp[i+w][j+w][k-w] == true) ||
						(dp[i][j+k-w][w] == true && dp[i+w][j][k-w] == true) {
						dp[i][j][k] = true
					}
				}
			}
		}
	}
	return dp[0][0][n]
}
