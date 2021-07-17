package main

import "fmt"

func main() {
	fmt.Println(largestMagicSquare([][]int{
		{7, 1, 4, 5, 6},
		{2, 5, 1, 6, 4},
		{1, 5, 4, 3, 2},
		{1, 2, 7, 3, 4},
	}))
}

func largestMagicSquare(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	rowArr := make([][]int, n) // 行前缀和
	colArr := make([][]int, n) // 列前缀和
	for i := 0; i < n; i++ {
		rowArr[i] = make([]int, m)
		colArr[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		rowArr[i][0] = grid[i][0]
		for j := 1; j < m; j++ {
			rowArr[i][j] = rowArr[i][j-1] + grid[i][j]
		}
	}
	for j := 0; j < m; j++ {
		colArr[0][j] = grid[0][j]
		for i := 1; i < n; i++ {
			colArr[i][j] = colArr[i-1][j] + grid[i][j]
		}
	}
	for length := min(n, m); length >= 2; length-- { // 枚举边长
		for i := 0; i+length <= n; i++ {
			for j := 0; j+length <= m; j++ {
				var target int // 以行目标和
				if j == 0 {
					target = rowArr[i][j+length-1]
				} else {
					target = rowArr[i][j+length-1] - rowArr[i][j-1]
				}
				flag := true
				for k := i + 1; k < i+length; k++ { // 继续检查行
					var temp int
					if j == 0 {
						temp = rowArr[k][j+length-1]
					} else {
						temp = rowArr[k][j+length-1] - rowArr[k][j-1]
					}
					if temp != target {
						flag = false
						break
					}
				}
				if flag == false {
					continue
				}
				for k := j; k < j+length; k++ { // 检查列
					var temp int
					if i == 0 {
						temp = colArr[i+length-1][k]
					} else {
						temp = colArr[i+length-1][k] - colArr[i-1][k]
					}
					if temp != target {
						flag = false
						break
					}
				}
				if flag == false {
					continue
				}
				var left, right int // 左右对角线
				for k := 0; k < length; k++ {
					left = left + grid[i+k][j+k]
					right = right + grid[i+k][j+length-1-k]
				}
				if target == left && target == right {
					return length
				}
			}
		}
	}
	return 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
