package main

import "sort"

func main() {

}

// leetcode1738_找出第K大的异或坐标值
func kthLargestValue(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	arr := make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j > 0 {
				matrix[i][j] = matrix[i][j-1] ^ matrix[i][j]
			} else if i > 0 && j == 0 {
				matrix[i][j] = matrix[i-1][j] ^ matrix[i][j]
			} else if i > 0 && j > 0 {
				matrix[i][j] = matrix[i-1][j] ^ matrix[i][j-1] ^ matrix[i][j] ^ matrix[i-1][j-1]
			}
			arr = append(arr, matrix[i][j])
		}
	}
	sort.Ints(arr)
	return arr[len(arr)-k]
}
