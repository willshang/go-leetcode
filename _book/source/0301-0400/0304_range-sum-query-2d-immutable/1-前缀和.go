package main

func main() {

}

// leetcode304_二维区域和检索-矩阵不可变
type NumMatrix struct {
	arr [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if matrix == nil || len(matrix) == 0 || matrix[0] == nil || len(matrix[0]) == 0 {
		arr := make([][]int, 1)
		for i := 0; i < 1; i++ {
			arr[i] = make([]int, 1)
		}
		return NumMatrix{arr: arr}
	}
	n, m := len(matrix), len(matrix[0])
	arr := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		arr[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			arr[i][j] = arr[i][j-1] + arr[i-1][j] - arr[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{arr: arr}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.arr[row2+1][col2+1] - this.arr[row2+1][col1] - this.arr[row1][col2+1] + this.arr[row1][col1]
}
