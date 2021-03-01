package main

import "sort"

func main() {

}

// leetcode1048_最长字符串链
func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	res := 0
	dp := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if judge(words[i], words[j]) == true && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func judge(a, b string) bool {
	if len(a)-len(b) != 1 {
		return false
	}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			j++
		}
		i++
	}
	return j == len(b)
}
