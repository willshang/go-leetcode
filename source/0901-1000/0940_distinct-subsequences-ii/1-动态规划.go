package main

import "fmt"

func main() {
	fmt.Println(distinctSubseqII("abc"))
	fmt.Println(distinctSubseqII("blljuffdyfrkqtwfyfztpdiyktrhftgtabxxoibcclbjvirnqyynkyaqlxgyybkgyzvcahmytjdqqtctirnxfjpktxmjkojlvvrr"))
}

var mod = 1000000007

func distinctSubseqII(s string) int {
	n := len(s)
	dp := make([]int, n+1)  // dp[i] 表示s[0:i+1]的不同子序列数
	m := make(map[byte]int) // 保存上次出现的下标
	for i := 0; i < n; i++ {
		if index, ok := m[s[i]]; ok { // 该字符出现过，取上次坐标
			dp[i+1] = (2*dp[i] - dp[index] + mod) % mod // 减去重复的
		} else {
			dp[i+1] = (2*dp[i] + 1) % mod // 这个不包含空序列
		}
		m[s[i]] = i
	}
	return dp[n] % mod
}
