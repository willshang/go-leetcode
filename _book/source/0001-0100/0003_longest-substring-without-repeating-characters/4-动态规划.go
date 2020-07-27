package main

import "fmt"

func main() {
	str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 1
	res := 1
	m := make(map[byte]int)
	m[s[0]] = 0
	for i := 1; i < len(s); i++ {
		index := -1
		if value, ok := m[s[i]]; ok {
			index = value
		}
		if i-index > dp[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = i - index
		}
		m[s[i]] = i
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
