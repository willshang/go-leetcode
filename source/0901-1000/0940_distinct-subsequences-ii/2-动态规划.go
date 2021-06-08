package main

import "fmt"

func main() {
	fmt.Println(distinctSubseqII("abc"))
	fmt.Println(distinctSubseqII("blljuffdyfrkqtwfyfztpdiyktrhftgtabxxoibcclbjvirnqyynkyaqlxgyybkgyzvcahmytjdqqtctirnxfjpktxmjkojlvvrr"))
}

// leetcode940_不同的子序列II
var mod = 1000000007

func distinctSubseqII(s string) int {
	n := len(s)
	dp := make([]int, n+1)  // dp[i] 表示s[0:i+1]的不同子序列数
	dp[0] = 1               // 空序列
	m := make(map[byte]int) // 保存上次出现的下标
	for i := 0; i < n; i++ {
		if index, ok := m[s[i]]; ok { // 该字符出现过，取上次坐标
			dp[i+1] = (2*dp[i] - dp[index] + mod) % mod // 减去重复的
		} else {
			// 如 ab 含有4种情况: "" "a" "b" "ab"
			// 如加入c后的abc 含有8种情况："" "a" "b" "ab" "c" "ac" "bc" "abc"
			dp[i+1] = 2 * dp[i] % mod // 不重复翻倍
		}
		m[s[i]] = i
	}
	return (dp[n] - 1) % mod // 去除空序列
}
