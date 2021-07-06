package main

import (
	"fmt"
)

func main() {
	fmt.Println(respace([]string{"abc", "ab"}, "aaa"))
}

// 程序员面试金典17.13_恢复空格
func respace(dictionary []string, sentence string) int {
	n := len(sentence)
	m := make(map[string]bool)
	for i := 0; i < len(dictionary); i++ {
		m[dictionary[i]] = true
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1 // 上一个长度+1
		for j := i; j >= 1; j-- {
			str := sentence[j-1 : i]
			if m[str] == true {
				dp[i] = min(dp[i], dp[j-1])
			}
			if dp[i] == 0 {
				break
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
