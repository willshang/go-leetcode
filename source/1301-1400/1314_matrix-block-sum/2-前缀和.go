package main

func main() {

}

// leetcode1314_矩阵区域和
func matrixBlockSum(mat [][]int, K int) [][]int {
	n, m := len(mat), len(mat[0])
	arr := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		arr[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			arr[i][j] = arr[i][j-1] + arr[i-1][j] - arr[i-1][j-1] + mat[i-1][j-1]
		}
	}
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
		for j := 0; j < m; j++ {
			left, right := max(0, j-K), min(m, j+K+1)
			up, down := max(0, i-K), min(n, i+K+1)
			res[i][j] = arr[down][right] - arr[down][left] - arr[up][right] + arr[up][left]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
