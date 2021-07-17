package main

import "fmt"

func main() {
	fmt.Println(countPalindromicSubsequence("aabca"))
}

// leetcode1930_长度为3的不同回文子序列
func countPalindromicSubsequence(s string) int {
	n := len(s)
	res := 0
	for i := 0; i < 26; i++ { // 枚举2边字符
		left, right := 0, n-1
		for left < n && s[left] != byte(i+'a') {
			left++
		}
		for right >= 0 && s[right] != byte(i+'a') {
			right--
		}
		if left+2 > right {
			continue
		}
		m := make(map[byte]int)
		for k := left + 1; k < right; k++ {
			m[s[k]] = 1
		}
		res = res + len(m)
	}
	return res
}
