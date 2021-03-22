package main

func main() {

}

// leetcode1329_将矩阵按对角线排序
func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	for k := 0; k < min(m, n); k++ {
		for i := 0; i < m-1; i++ {
			for j := 0; j < n-1; j++ {
				if mat[i][j] > mat[i+1][j+1] {
					mat[i][j], mat[i+1][j+1] = mat[i+1][j+1], mat[i][j]
				}
			}
		}
	}
	return mat
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
