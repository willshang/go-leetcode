package main

import ()

func main() {
	//arr := [][]byte{
	//	{'1', '0', '1', '0', '0'},
	//	{'1', '0', '1', '1', '1'},
	//	{'1', '1', '1', '1', '1'},
	//	{'1', '0', '0', '1', '0'},
	//}
	//fmt.Println(maximalRectangle(arr))
}

func maximalRectangle(matrix []string) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	res := 0
	n, m := len(matrix), len(matrix[0])
	left, right, height := make([]int, m), make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		right[i] = m
	}
	for i := 0; i < n; i++ {
		curLeft, curRight := 0, m
		// 高度
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		// 左边
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				left[j] = max(left[j], curLeft)
			} else {
				left[j] = 0
				curLeft = j + 1
			}
		}
		// 右边
		for j := m - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				right[j] = min(right[j], curRight)
			} else {
				right[j] = m
				curRight = j
			}
		}
		for j := 0; j < m; j++ {
			res = max(res, height[j]*(right[j]-left[j]))
		}
	}
	return res
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
