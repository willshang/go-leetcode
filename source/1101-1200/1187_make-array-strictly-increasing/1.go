package main

import (
	"math"
	"sort"
)

func main() {

}

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	sort.Ints(arr2)
	n, m := len(arr1), len(arr2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	//dp[0][0] = 0
	//dp[0][1] =
}
