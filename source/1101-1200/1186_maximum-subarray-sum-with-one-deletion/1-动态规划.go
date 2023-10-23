package main

func main() {

}

func maximumSum(arr []int) int {
	n := len(arr)
	dp := make([][2]int, n)
	dp[0][0] = arr[0] // 不删
	dp[0][1] = 0      // 删除
	res := dp[0][0]   // 长度为1，不删除
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0]+arr[i], arr[i])
		dp[i][1] = max(dp[i-1][1]+arr[i], dp[i-1][0])
		res = max(res, max(dp[i][0], dp[i][1]))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
