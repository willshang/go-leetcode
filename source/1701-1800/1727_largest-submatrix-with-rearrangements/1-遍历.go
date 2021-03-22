package main

import "sort"

func main() {

}

// leetcode1727_重新排列后的最大子矩阵
func largestSubmatrix(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 { // 统计以该行为基连续1的个数
				matrix[i][j] = matrix[i][j] + matrix[i-1][j]
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		sort.Ints(matrix[i])
		for j := m - 1; j >= 0; j-- {
			height := matrix[i][j] // 高度有序，右边最高=>高度取决于左边
			width := m - j
			if height*width > res {
				res = height * width
			}
		}
	}
	return res
}
