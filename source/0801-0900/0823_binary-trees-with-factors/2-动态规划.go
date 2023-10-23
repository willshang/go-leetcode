package main

import "sort"

func main() {

}

var mod = 1000000007

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	n := len(arr)
	dp := make(map[int]int)
	res := 0
	for i := 0; i < n; i++ {
		dp[arr[i]] = 1
		for j := 0; j < i; j++ {
			if arr[i]%arr[j] == 0 {
				c := arr[i] / arr[j]
				dp[arr[i]] = (dp[arr[i]] + dp[arr[j]]*dp[c]) % mod
			}
		}
		res = (res + dp[arr[i]]) % mod
	}
	return res
}
