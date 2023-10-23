package main

import "fmt"

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

// leetcode139_单词拆分
func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		m[wordDict[i]] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] == true && m[s[j:i]] == true {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
