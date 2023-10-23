package main

func main() {

}

func numberOfArithmeticSlices(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}
	res := 0
	dp := make([]int, n)
	for i := 2; i < n; i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			dp[i] = dp[i-1] + 1
		}
		res = res + dp[i]
	}
	return res
}
