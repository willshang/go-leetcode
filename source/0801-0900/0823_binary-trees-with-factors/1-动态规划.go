package main

import "sort"

func main() {

}

// leetcode823_带因子的二叉树
var mod = 1000000007

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	n := len(arr)
	m := make(map[int]int)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		m[arr[i]] = i
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i]%arr[j] == 0 {
				c := arr[i] / arr[j]
				if v, ok := m[c]; ok {
					dp[i] = (dp[i] + dp[j]*dp[v]) % mod
				}
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res = (res + dp[i]) % mod
	}
	return res
}
