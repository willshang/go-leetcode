package main

func main() {

}

func maxSideLength(mat [][]int, threshold int) int {
	n, m := len(mat), len(mat[0])
	arr := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			arr[i+1][j+1] = mat[i][j] + arr[i+1][j] + arr[i][j+1] - arr[i][j]
		}
	}
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k := min(i, j); k > res; k-- {
				value := arr[i][j] - arr[i-k][j] - arr[i][j-k] + arr[i-k][j-k]
				if value <= threshold && res < k {
					res = k
					break
				}
			}
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
