package main

import "sort"

func main() {

}

var m map[string]bool

func longestWord(words []string) string {
	m = make(map[string]bool)
	n := len(words)
	for i := 0; i < n; i++ {
		m[words[i]] = true
	}
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return len(words[i]) > len(words[j])
	})
	// 从最长最小字典序的开始找
	for i := 0; i < n; i++ {
		m[words[i]] = false
		if judge(words[i]) == true {
			return words[i]
		}
	}
	return ""
}

// leetcode 139.单词拆分
func judge(s string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	n := len(s)
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] == true && m[s[j:i]] == true {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
