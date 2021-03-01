package main

func main() {

}

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[arr[i]] = i
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dp[i][j] = 2
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			index, ok := m[arr[i]-arr[j]]
			if ok && arr[index] < arr[j] {
				dp[j][i] = dp[index][j] + 1
				if dp[j][i] > 2 && dp[j][i] > res {
					res = dp[j][i]
				}
			}
		}
	}
	return res
}
