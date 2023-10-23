package main

func main() {

}

func longestPalindrome(word1 string, word2 string) int {
	s := word1 + word2
	a := len(word1)
	n := len(s)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	res := 0
	for i := n - 1; i >= 0; i-- {
		prev := 0
		for j := i + 1; j < n; j++ {
			temp := dp[j]
			if s[i] == s[j] {
				dp[j] = prev + 2 // 内层+2
				if i < a && j >= a {
					res = max(res, dp[j])
				}
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			prev = temp
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
