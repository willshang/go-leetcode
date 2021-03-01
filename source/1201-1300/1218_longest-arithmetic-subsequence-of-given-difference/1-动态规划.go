package main

func main() {

}

// leetcode1218_最长定差子序列
func longestSubsequence(arr []int, difference int) int {
	res := 0
	dp := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		dp[arr[i]] = dp[arr[i]-difference] + 1
		if dp[arr[i]] > res {
			res = dp[arr[i]]
		}
	}
	return res
}
