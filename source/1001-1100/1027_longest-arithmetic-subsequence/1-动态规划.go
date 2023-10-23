package main

func main() {

}

// leetcode1027_最长等差数列
func longestArithSeqLength(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}
	res := 0
	dp := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := A[i] - A[j]
			if _, ok := dp[j][diff]; ok {
				dp[i][diff] = dp[j][diff] + 1
			} else {
				dp[j][diff] = 1
				dp[i][diff] = dp[j][diff] + 1
			}
			if dp[i][diff] > res {
				res = dp[i][diff]
			}
		}
	}
	return res
}
