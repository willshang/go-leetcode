package main

func main() {

}

// leetcode1641_统计字典序元音字符串的数目
func countVowelStrings(n int) int {
	dp := make([]int, 5)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := 1; j < 5; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	res := 0
	for i := 0; i < 5; i++ {
		res = res + dp[i]
	}
	return res
}
