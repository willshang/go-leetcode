package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

var dp [][]int

func minDistance(word1 string, word2 string) int {
	dp = make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	return helper(word1, word2, 0, 0)
}

func helper(word1, word2 string, i, j int) int {
	if dp[i][j] > 0 {
		return dp[i][j]
	}
	if i == len(word1) || j == len(word2) {
		return len(word1) - i + len(word2) - j
	}
	if word1[i] == word2[j] {
		return helper(word1, word2, i+1, j+1)
	}
	inserted := helper(word1, word2, i, j+1)
	deleted := helper(word1, word2, i+1, j)
	replaced := helper(word1, word2, i+1, j+1)
	dp[i][j] = min(inserted, min(deleted, replaced)) + 1
	return dp[i][j]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
