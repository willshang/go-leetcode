package main

func main() {

}

func countSquares(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	dp := make([]int, m)
	res := 0
	for i := 0; i < n; i++ {
		temp := make([]int, m)
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				temp[j] = matrix[i][j]
			} else if matrix[i][j] == 0 {
				temp[j] = 0
			} else {
				temp[j] = min(min(temp[j-1], dp[j]), dp[j-1]) + 1
			}
			res = res + temp[j]
		}
		copy(dp, temp)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
