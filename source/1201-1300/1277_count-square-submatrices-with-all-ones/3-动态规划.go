package main

func main() {

}

// leetcode1277_统计全为1的正方形子矩阵
func countSquares(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i > 0 && j > 0 && matrix[i][j] > 0 {
				matrix[i][j] = min(min(matrix[i][j-1], matrix[i-1][j]), matrix[i-1][j-1]) + 1
			}
			res = res + matrix[i][j]
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
