package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	prev := make([]int, m+1)
	cur := make([]int, m+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				cur[j] = prev[j-1] + 1
			} else {
				cur[j] = max(prev[j], cur[j-1])
			}
		}
		copy(prev, cur)
	}
	return cur[m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
