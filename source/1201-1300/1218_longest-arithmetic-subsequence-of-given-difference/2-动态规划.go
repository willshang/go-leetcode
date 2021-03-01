package main

func main() {

}

func longestSubsequence(arr []int, difference int) int {
	res := 0
	dp := make([]int, 40001)
	for i := 0; i < len(arr); i++ {
		index := arr[i] + 20000
		dp[index] = dp[index-difference] + 1
		if dp[index] > res {
			res = dp[index]
		}
	}
	return res
}
