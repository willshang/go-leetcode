package main

func main() {

}

// leetcode1292_元素和小于等于阈值的正方形的最大边长
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
			for k := 0; i+k <= n && j+k <= m; k++ {
				value := arr[i+k][j+k] - arr[i-1][j+k] - arr[i+k][j-1] + arr[i-1][j-1]
				if value <= threshold && res < k+1 {
					res = k + 1
				}
				if value > threshold {
					break
				}
			}
		}
	}
	return res
}
