package main

func main() {

}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	res := 0
	n, m := len(matrix), len(matrix[0])
	arr := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			arr[i][j] = matrix[i-1][j-1] + arr[i-1][j] + arr[i][j-1] - arr[i-1][j-1]
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for x := i; x <= n; x++ {
				for y := j; y <= m; y++ {
					if arr[x][y]-arr[i-1][y]-arr[x][j-1]+arr[i-1][j-1] == target {
						res++
					}
				}
			}
		}
	}
	return res
}
