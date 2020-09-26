package main

import (
	"fmt"
)

func main() {
	arr := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(arr))
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	res := 0
	n, m := len(matrix), len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				res = max(res, 1)
				minLength := min(n-i, m-j)
				for k := 1; k < minLength; k++ {
					flag := true
					if matrix[i+k][j+k] == '0' {
						break
					}
					for l := 0; l < k; l++ {
						if matrix[i+k][j+l] == '0' || matrix[i+l][j+k] == '0' {
							flag = false
							break
						}
					}
					if flag == true {
						res = max(res, k+1)
					} else {
						break
					}
				}
			}
		}
	}
	return res * res
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
