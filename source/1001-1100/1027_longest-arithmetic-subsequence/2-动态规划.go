package main

func main() {

}

func longestArithSeqLength(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}
	res := 0
	dp := make([][20001]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := A[i] - A[j] + 10001
			if dp[j][diff] > 0 {
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
