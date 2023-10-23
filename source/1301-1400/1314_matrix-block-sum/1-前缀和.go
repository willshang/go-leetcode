package main

func main() {

}

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
			a1, a2 := getIndex(n, m, i+K+1, j+K+1)
			b1, b2 := getIndex(n, m, i-K, j+K+1)
			c1, c2 := getIndex(n, m, i+K+1, j-K)
			d1, d2 := getIndex(n, m, i-K, j-K)
			res[i][j] = arr[a1][a2] - arr[b1][b2] - arr[c1][c2] + arr[d1][d2]
		}
	}
	return res
}

func getIndex(a, b, x, y int) (int, int) {
	x = max(min(a, x), 0)
	y = max(min(b, y), 0)
	return x, y
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
