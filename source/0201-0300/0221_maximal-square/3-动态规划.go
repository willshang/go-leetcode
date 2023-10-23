package main

import "fmt"

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
				res = max(res, int(matrix[i][j]-'0'))
			}
			if i == 0 || j == 0 {
				continue
			}
			if matrix[i][j] == '1' {
				a := int(matrix[i-1][j-1] - '0')
				b := int(matrix[i-1][j] - '0')
				c := int(matrix[i][j-1] - '0')
				matrix[i][j] = byte(min(a, min(b, c)) + 1 + '0')
				res = max(res, int(matrix[i][j]-'0'))
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
